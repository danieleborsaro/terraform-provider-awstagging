package tagging

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type Namer struct {
	configuration *data.InputConfiguration
	// components    NameComponents
	components             map[string]string
	properties             *data.ResourceProperties
	sanitisedConfiguration *data.SanitisedConfiguration
	Name                   data.Name
	Type                   data.Type
}

// type NameComponents struct {
// 	// account           string
// 	Account           string
// 	AppEnvironment   string
// 	AvailabilityZone string
// 	Business          string
// 	ComputeType      string
// 	InfraEnvironment string
// 	Platform          string
// 	Project           string
// 	Region            string
// 	ResourceSet      string
// 	ResourceType     string
// 	Role              string
// }

func (thisResource *Namer) init(properties *data.ResourceProperties, config *data.InputConfiguration, sanitisedConfig *data.SanitisedConfiguration) {
	thisResource.configuration = config
	thisResource.properties = properties
	thisResource.sanitisedConfiguration = sanitisedConfig
}

func (thisResource *Namer) setName(ctx context.Context) {
	tflog.Trace(ctx, "namer - BEGIN name generation")

	var account = thisResource.sanitisedConfiguration.EnvironmentClasses[thisResource.sanitisedConfiguration.AccountObject.Class].NameEncoded
	if account == "" {
		panic("Account ID '" + thisResource.configuration.Account + "' not recognised, aborting")
	}

	infraEnvironment := thisResource.sanitisedConfiguration.Keys.InfraEnvironment[0 : 0+min(uint(len(thisResource.sanitisedConfiguration.Keys.InfraEnvironment)), thisResource.sanitisedConfiguration.Constraints.EnvNameMaxLength)]
	appEnvironment := thisResource.sanitisedConfiguration.Keys.AppEnvironment[0 : 0+min(uint(len(thisResource.sanitisedConfiguration.Keys.AppEnvironment)), thisResource.sanitisedConfiguration.Constraints.EnvNameMaxLength)]
	if strings.EqualFold(appEnvironment, infraEnvironment) {
		appEnvironment = ""
	} //// If app_environment and infra_environment are the same, don't add app_environment to save length

	resourceSet := ""
	if !strings.EqualFold(thisResource.configuration.ResourceSetShort, thisResource.configuration.ProjectNameShort) {
		resourceSet = strings.ToLower(thisResource.configuration.ResourceSetShort)
	} //// If resource_set and project are the same, don't add resource_set to save length

	resourceType := ""
	resourceIdShort := thisResource.properties.Id.Short
	if thisResource.properties.Id.Key == "AwsLb" && thisResource.configuration.IsApplicationLoadBalancer {
		resourceIdShort = "alb"
	} else if thisResource.properties.Id.Key == "AwsLb" && !thisResource.configuration.IsApplicationLoadBalancer {
		resourceIdShort = "nlb"

	} else if thisResource.properties.Id.Key == "AwsRoute53ResolverEndpoint" && thisResource.configuration.IsInbound {
		resourceIdShort = "r53rei"
	} else if thisResource.properties.Id.Key == "AwsRoute53ResolverEndpoint" && !thisResource.configuration.IsInbound {
		resourceIdShort = "r53reo"
	}

	if !strings.EqualFold(resourceIdShort, thisResource.configuration.ResourceSetShort) {
		resourceType = strings.ToLower(resourceIdShort)
	} //// If resource_type and resource_set are the same, don't add resource_type to save length

	role := ""
	if !strings.EqualFold(thisResource.configuration.ResourceSetShort, thisResource.sanitisedConfiguration.CanonicalRole) {
		role = thisResource.sanitisedConfiguration.CanonicalRole
	} //// If role and resource_set are the same, don't add role to save length

	thisResource.components = map[string]string{
		"account":           account,
		"app_environment":   appEnvironment,
		"availability_zone": thisResource.sanitisedConfiguration.CanonicalAvailabilityZone,
		"business":          thisResource.configuration.CompanyNameShort,
		"compute_type":      strings.ToLower(thisResource.configuration.ComputeType),
		"infra_environment": infraEnvironment,
		"platform":          strings.ToLower(thisResource.configuration.PlatformName),
		"project":           strings.ToLower(thisResource.configuration.ProjectNameShort),
		"region":            thisResource.sanitisedConfiguration.CanonicalRegion,
		"resource_set":      resourceSet,
		"resource_type":     resourceType,
		"role":              role,
	}

	//// Assemble name from resource components specification (cattle)
	selectedComponents := []string{}
	for _, c := range thisResource.properties.Name.Components {
		if thisResource.components[c] != "" {
			//// Filter out empty components
			selectedComponents = append(selectedComponents, thisResource.components[c])
		}
	}

	assembledName := strings.Join(selectedComponents, thisResource.sanitisedConfiguration.Separators.NameComponent)
	//// Check whether we want to use a custom name instead (pet)
	resourceNameStaticUnsafe := thisResource.sanitisedConfiguration.CustomName
	if resourceNameStaticUnsafe == "" {
		resourceNameStaticUnsafe = assembledName
	}

	//// Does resource wants an all lowercase name?
	resourceNameStaticUnsafeFormatted := ""
	if thisResource.properties.Name.IsLowerCase {
		resourceNameStaticUnsafeFormatted = strings.ToLower(resourceNameStaticUnsafe)
	} else {
		resourceNameStaticUnsafeFormatted = resourceNameStaticUnsafe
	}
	resourceNameStaticUnsafeFormatted = strings.Trim(resourceNameStaticUnsafeFormatted, thisResource.sanitisedConfiguration.Separators.NameComponent)

	//// Compute versioned hash suffix for explicit disambiguation
	versioningList := []string{}
	// for _, v := range thisResource.configuration.VersioningSources {
	// 	versioningList = append(versioningList, v)
	// }
	versioningList = append(versioningList, thisResource.configuration.VersioningSources...)

	versioningList = append(versioningList, resourceNameStaticUnsafeFormatted)
	resourceVersion := ""
	if thisResource.sanitisedConfiguration.IsVersionedName {
		hasher := sha1.New()
		hasher.Write([]byte(strings.Join(versioningList, thisResource.sanitisedConfiguration.Separators.NameComponent)))
		resourceVersion = hex.EncodeToString(hasher.Sum(nil))
	}

	//// Compute lengths for hashed suffix and prefix, related to name max lenght for resource
	//// NB: we are using uints, negative bumbers overflow and become very big numbers!
	resourceNameVersionedLengthLeft := int(thisResource.properties.Name.MaxLength) - len(resourceNameStaticUnsafeFormatted)

	tflog.Debug(ctx, thisResource.properties.Id.Key+" Unsafe prefix: "+resourceNameStaticUnsafeFormatted)
	tflog.Debug(ctx, thisResource.properties.Id.Key+" Length unsafe prefix: "+strconv.FormatUint(uint64(len(resourceNameStaticUnsafeFormatted)), 10))
	tflog.Debug(ctx, thisResource.properties.Id.Key+" Unsafe suffix: "+resourceVersion)
	tflog.Debug(ctx, thisResource.properties.Id.Key+" Length unsafe suffix: "+strconv.FormatUint(uint64(len(resourceVersion)), 10))
	tflog.Debug(ctx, thisResource.properties.Id.Key+" Length max: "+strconv.FormatUint(uint64(thisResource.properties.Name.MaxLength), 10))
	tflog.Debug(ctx, thisResource.properties.Id.Key+" Length left: "+strconv.FormatInt(int64(resourceNameVersionedLengthLeft), 10))

	resourceNameVersionedPrefixLength := int(0)
	if resourceNameVersionedLengthLeft >= int(thisResource.sanitisedConfiguration.Constraints.VersionStringMinLength) {
		//// Enough length left for the hashed suffix (of what length still available)
		resourceNameVersionedPrefixLength = len(resourceNameStaticUnsafeFormatted)
		tflog.Debug(ctx, thisResource.properties.Id.Key+" Plenty Prefix left, new length: "+strconv.FormatInt(int64(resourceNameVersionedPrefixLength), 10))
	} else {
		//// Not enough space for the hashed suffix, must shorten the name to make room!
		resourceNameVersionedPrefixLength = int(thisResource.properties.Name.MaxLength) - int(thisResource.sanitisedConfiguration.Constraints.VersionStringMinLength)
		tflog.Debug(ctx, thisResource.properties.Id.Key+" Not Enough Prefix left, new length: "+strconv.FormatInt(int64(resourceNameVersionedPrefixLength), 10))
	}

	resourceNameVersionedSuffixLength := int(0)
	if resourceNameVersionedLengthLeft >= int(thisResource.sanitisedConfiguration.Constraints.VersionStringMinLength) {
		//// Enough length left for the hashed suffix, compute exactly how long it will be
		//// NB: resourceNameVersionedLengthLeft coulb ne negative, but we reach this branch only if positive - we tweak anyway for extr asafety
		resourceNameVersionedSuffixLength = min(len(resourceVersion), resourceNameVersionedLengthLeft)
		tflog.Debug(ctx, thisResource.properties.Id.Key+" Plenty Suffix left, new length: "+strconv.FormatInt(int64(resourceNameVersionedSuffixLength), 10))
	} else {
		//// Not enough space for the hashed suffix, but we need one of minimal length so force it
		//// This length matches what we shortened the prefix with
		resourceNameVersionedSuffixLength = int(thisResource.sanitisedConfiguration.Constraints.VersionStringMinLength)
		tflog.Debug(ctx, thisResource.properties.Id.Key+" Not Enough Suffix left, new length: "+strconv.FormatInt(int64(resourceNameVersionedSuffixLength), 10))
	}

	//// Names composition for resources with common constraints
	prefix := ""
	if thisResource.sanitisedConfiguration.IsVersionedName {
		prefix = strings.Trim(resourceNameStaticUnsafeFormatted[0:0+resourceNameVersionedPrefixLength], thisResource.sanitisedConfiguration.Separators.NameComponent)
		tflog.Debug(ctx, thisResource.properties.Id.Key+" Shortening Prefix: "+prefix)
		tflog.Debug(ctx, thisResource.properties.Id.Key+" Shortening Prefix length: "+strconv.FormatUint(uint64(len(prefix)), 10))
	} else {
		prefix = resourceNameStaticUnsafeFormatted
		tflog.Debug(ctx, thisResource.properties.Id.Key+" Non Shortening Prefix: "+prefix)
		tflog.Debug(ctx, thisResource.properties.Id.Key+" Non Shortening Prefix length: "+strconv.FormatUint(uint64(len(prefix)), 10))
	}

	suffix := ""
	if thisResource.sanitisedConfiguration.IsVersionedName {
		suffix = strings.Trim(resourceVersion[0:0+resourceNameVersionedSuffixLength], thisResource.sanitisedConfiguration.Separators.NameComponent)

		tflog.Debug(ctx, thisResource.properties.Id.Key+" Shortening Suffix: "+suffix)
		tflog.Debug(ctx, thisResource.properties.Id.Key+" Shortening Suffix length: "+strconv.FormatUint(uint64(len(suffix)), 10))
	}
	versioned := map[string]string{
		"prefix": prefix,
		"suffix": suffix,
	}

	//// We now have suffix and prefix, compose the name
	unsafe := versioned["prefix"]
	if thisResource.sanitisedConfiguration.IsVersionedName {
		unsafe = versioned["prefix"] + "-" + versioned["suffix"]
	}

	// //// For extra safety, further shorten the ovweall name
	unsafe = unsafe[0 : 0+min(uint(len(unsafe)), thisResource.properties.Name.MaxLength)] //// Make sure we fit in name max length

	// tflog.Debug( ctx, thisResource.properties.Id.Key + " IsVersioned: " + strconv.FormatBool(thisResource.sanitisedConfiguration.IsVersionedName) )
	// tflog.Debug( ctx, thisResource.properties.Id.Key + " Max length: " + strconv.FormatUint(uint64(thisResource.properties.Name.MaxLength), 10) )
	// tflog.Debug( ctx, thisResource.properties.Id.Key + " Prefix: " + versioned["prefix"] )
	// tflog.Debug( ctx, thisResource.properties.Id.Key + " Prefix length: " + strconv.FormatUint(uint64(len(versioned["prefix"])), 10) )
	// tflog.Debug( ctx, thisResource.properties.Id.Key + " Suffix: " + versioned["suffix"] )
	// tflog.Debug( ctx, thisResource.properties.Id.Key + " Suffix length: " + strconv.FormatUint(uint64(len(versioned["suffix"])), 10) )
	// tflog.Debug( ctx, thisResource.properties.Id.Key + " Safe name: " + nameConstrained["name"]["safe"] )
	// tflog.Debug( ctx, thisResource.properties.Id.Key + " Safe name length: " + strconv.FormatUint(uint64(len(nameConstrained["name"]["safe"])), 10) )

	//// END - leaving this block here for backward compatibility

	//// With the names composed, we now compute additional properties that depend on them
	standardTypePrefix := strings.ToLower(thisResource.properties.Id.Short)[0 : 0+min(uint(len(strings.ToLower(thisResource.properties.Id.Short))), thisResource.sanitisedConfiguration.Constraints.TypePrefixMaxLength)]
	if thisResource.properties.Name.IsLongPrefix {
		//// TODO: this is hardly used, but worth double checking it's right
		////       It should at least use the safe name/prefix
		// // standardTypePrefix = nameCore["name"]["unsafe"]
		// standardTypePrefix = nameCore["name"]["safe"]
		standardTypePrefix = versioned["prefix"]
	}

	//// BEGIN - leaving this block here for backward compatibility
	// standardResource := map[string]interface{}{
	//   "name": data.Name{
	//     Safe        : nameConstrained["name"]["safe"],
	//     SafePrefix: versioned["prefix"],
	//     SafeSuffix: nameConstrained["name"]["safe"][len(versioned["prefix"]):len(versioned["prefix"]) + len(nameConstrained["name"]["safe"]) - len(versioned["prefix"])],
	//     TypePrefix: standardTypePrefix,
	//   },
	//   "type": data.Type{
	//     Long: thisResource.properties.Id.Long,
	//   },
	// }
	//// END - leaving this block here for backward compatibility

	standardResource := map[string]interface{}{
		"name": data.Name{
			Safe:       unsafe,
			SafePrefix: versioned["prefix"],
			SafeSuffix: versioned["suffix"],
			TypePrefix: standardTypePrefix,
		},
		"type": data.Type{
			Long: thisResource.properties.Id.Long,
		},
	}

	tflog.Debug(ctx, thisResource.properties.Id.Key+" Standard IsVersioned: "+strconv.FormatBool(thisResource.sanitisedConfiguration.IsVersionedName))
	tflog.Debug(ctx, thisResource.properties.Id.Key+" Standard  Max length: "+strconv.FormatUint(uint64(thisResource.properties.Name.MaxLength), 10))
	tflog.Debug(ctx, thisResource.properties.Id.Key+" Standard  Prefix: "+standardResource["name"].(data.Name).SafePrefix)
	tflog.Debug(ctx, thisResource.properties.Id.Key+" Standard  Prefix length: "+strconv.FormatUint(uint64(len(standardResource["name"].(data.Name).SafePrefix)), 10))
	tflog.Debug(ctx, thisResource.properties.Id.Key+" Standard  Suffix: "+standardResource["name"].(data.Name).SafeSuffix)
	tflog.Debug(ctx, thisResource.properties.Id.Key+" Standard  Suffix length: "+strconv.FormatUint(uint64(len(standardResource["name"].(data.Name).SafeSuffix)), 10))
	tflog.Debug(ctx, thisResource.properties.Id.Key+" Standard  Safe name: "+standardResource["name"].(data.Name).Safe)
	tflog.Debug(ctx, thisResource.properties.Id.Key+" Standard  Safe name length: "+strconv.FormatUint(uint64(len(standardResource["name"].(data.Name).Safe)), 10))

	//// Names composition for resources with specific constraints
	r := regexp.MustCompile(`^alias`)

	kmsAliasNameSafe := ""
	if len(r.FindAllString(standardResource["name"].(data.Name).Safe, -1)) > 0 {
		kmsAliasNameSafe = standardResource["name"].(data.Name).Safe
	} else {
		kmsAliasNameSafe = strings.ReplaceAll("alias/"+standardResource["name"].(data.Name).Safe, "////", "/")
	}

	kmsAliasTypePrefix := ""
	if len(r.FindAllString(standardResource["name"].(data.Name).TypePrefix, -1)) > 0 {
		kmsAliasTypePrefix = standardResource["name"].(data.Name).Safe
	} else {
		kmsAliasTypePrefix = strings.ReplaceAll("alias/"+standardResource["name"].(data.Name).TypePrefix, "////", "/")
	}

	//// NB: resource IDs to match what set with reflect.TypeOf(*this).Name()
	nonStandardResources := map[string]map[string]interface{}{
		"AwsKmsAlias": {
			"name": data.Name{
				//// Only KMS keys are subject to rules, if they are aliases instead then keep as is
				Safe:       strings.ReplaceAll(kmsAliasNameSafe, ".", "-"),
				SafePrefix: strings.ReplaceAll(standardResource["name"].(data.Name).SafePrefix, ".", "-"),
				SafeSuffix: strings.ReplaceAll(standardResource["name"].(data.Name).SafeSuffix, ".", "-"),
				TypePrefix: strings.ReplaceAll(kmsAliasTypePrefix, ".", "-"),
			},

			"type": data.Type{
				Long: standardResource["type"].(data.Type).Long,
			},
		},

		"AwsCloudwatchLogStream": {
			"name": data.Name{
				//// Logstream name separator is different from the one of other resource names, let's patch what we build so far
				Safe:       strings.ReplaceAll(standardResource["name"].(data.Name).Safe, thisResource.sanitisedConfiguration.Separators.NameComponent, thisResource.sanitisedConfiguration.Separators.LogstreamComponent),
				SafePrefix: standardResource["name"].(data.Name).SafePrefix,
				SafeSuffix: standardResource["name"].(data.Name).SafeSuffix,
				TypePrefix: standardResource["name"].(data.Name).TypePrefix,
			},

			"type": data.Type{
				Long: standardResource["type"].(data.Type).Long,
			},
		},

		"AwsLbTargetGroup": {
			"name": data.Name{
				Safe:       standardResource["name"].(data.Name).Safe,
				SafePrefix: standardResource["name"].(data.Name).SafePrefix,
				SafeSuffix: standardResource["name"].(data.Name).SafeSuffix,
				//// Target Groups encoded types better carry their app environment for better disambiguation,
				TypePrefix: standardResource["name"].(data.Name).TypePrefix + "-" + thisResource.sanitisedConfiguration.Keys.AppEnvironment,
			},

			"type": data.Type{
				Long: standardResource["type"].(data.Type).Long,
			},
		},

		"AwsDbParameterGroup": {
			"name": data.Name{
				//// Only accepts lower alphanumeric characters with hyphens
				Safe:       strings.ToLower(standardResource["name"].(data.Name).Safe),
				SafePrefix: strings.ToLower(standardResource["name"].(data.Name).SafePrefix),
				SafeSuffix: strings.ToLower(standardResource["name"].(data.Name).SafeSuffix),
				TypePrefix: strings.ToLower(standardResource["name"].(data.Name).TypePrefix),
			},

			"type": data.Type{
				Long: standardResource["type"].(data.Type).Long,
			},
		},
	}

	//// All resources examined, now apply global constraints
	//// Assemble standard and non standard resource properties on a single object, for convenience
	nonLintedSafe := standardResource["name"].(data.Name).Safe
	if _, ok := nonStandardResources[thisResource.properties.Id.Key]; ok {
		nonLintedSafe = nonStandardResources[thisResource.properties.Id.Key]["name"].(data.Name).Safe
	}

	nonLintedPrefix := standardResource["name"].(data.Name).SafePrefix
	if _, ok := nonStandardResources[thisResource.properties.Id.Key]; ok {
		nonLintedPrefix = nonStandardResources[thisResource.properties.Id.Key]["name"].(data.Name).SafePrefix
	}

	nonLintedSuffix := standardResource["name"].(data.Name).SafeSuffix
	if _, ok := nonStandardResources[thisResource.properties.Id.Key]; ok {
		nonLintedSuffix = nonStandardResources[thisResource.properties.Id.Key]["name"].(data.Name).SafeSuffix
	}

	nonLintedType := standardResource["name"].(data.Name).TypePrefix
	if _, ok := nonStandardResources[thisResource.properties.Id.Key]; ok {
		nonLintedType = nonStandardResources[thisResource.properties.Id.Key]["name"].(data.Name).TypePrefix
	}

	nameSafeNotLinted := data.Name{
		Safe:       nonLintedSafe,
		SafePrefix: nonLintedPrefix,
		SafeSuffix: nonLintedSuffix,
		TypePrefix: nonLintedType,
	}

	//// Linting
	nameLintNoHyphen := data.Name{
		Safe:       strings.Trim(nameSafeNotLinted.Safe, thisResource.sanitisedConfiguration.Separators.NameComponent),
		SafePrefix: strings.Trim(nameSafeNotLinted.SafePrefix, thisResource.sanitisedConfiguration.Separators.NameComponent),
		SafeSuffix: strings.Trim(nameSafeNotLinted.SafeSuffix, thisResource.sanitisedConfiguration.Separators.NameComponent),
		TypePrefix: nameSafeNotLinted.TypePrefix,
	}

	//// Finally, name properties for all resources!
	resName := data.Name{
		Safe:       nameLintNoHyphen.Safe,
		SafePrefix: nameLintNoHyphen.SafePrefix,
		SafeSuffix: nameLintNoHyphen.SafeSuffix,
		TypePrefix: nameLintNoHyphen.TypePrefix + "" + thisResource.sanitisedConfiguration.Separators.NameComponent,
	}
	thisResource.Name = resName

	resTypeLong := thisResource.properties.Id.Long
	if _, ok := nonStandardResources[thisResource.properties.Id.Key]; ok {
		resTypeLong = nonStandardResources[thisResource.properties.Id.Key]["type"].(data.Type).Long
	}
	resType := data.Type{
		Long: resTypeLong,
	}
	thisResource.Type = resType

	tflog.Debug(ctx, "namer - name generated")

	tflog.Trace(ctx, "namer - END name generation")
}
