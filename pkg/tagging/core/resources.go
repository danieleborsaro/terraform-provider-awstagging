package tagging

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type ResourceInterface interface {
	Generate(context.Context, *data.InputConfiguration) error
	GetProperties() *data.ResourceProperties
	GetName() data.Name
	GetTags() map[string]string
	GetTagsAsMap() []map[string]string
	GetType() data.Type
}

type Resource struct {
	Properties             *data.ResourceProperties
	Configuration          *data.InputConfiguration
	Name                   data.Name
	SanitisedConfiguration *data.SanitisedConfiguration
	Tags                   map[string]string
	TagsAsMap              []map[string]string
	Type                   data.Type

	AccountsCodingById map[string]data.AccountCodingConfiguration
	Prefixes           map[string]string
	Separators         map[string]string
}

func (thisResource *Resource) GetProperties() *data.ResourceProperties {
	return thisResource.Properties
}

func (thisResource *Resource) Generate(ctx context.Context, config *data.InputConfiguration) error {
	tflog.Trace(ctx, "resources - BEGIN info generation")

	thisResource.Configuration = config

	if thisResource.Configuration.InfraEnvironment == "" && thisResource.Configuration.AppEnvironment == "" {
		return fmt.Errorf("[ERROR] Specify at least one between InfraEnvironment and AppEnvironment. If only one is specified, they will both get the same value")
	}

	if thisResource.Configuration.CustomName != "" && thisResource.Configuration.CustomNamePrefix != "" {
		return fmt.Errorf("[ERROR] Specify only one between CustomName and CustomNamePrefix")
	}

	environmentClasses := map[string]data.EnvironmentClassProperties{
		"development": data.EnvironmentClassProperties{
			NameEncoded: "D",
		},

		"preproduction": data.EnvironmentClassProperties{
			NameEncoded: "E",
		},

		"production": data.EnvironmentClassProperties{
			NameEncoded: "P",
		},

		"management": data.EnvironmentClassProperties{
			NameEncoded: "M",
		},

		"undefined": data.EnvironmentClassProperties{
			NameEncoded: "U",
		},
	}

	//// NB this falls back to an 'undefined' default, it's dangerous
	// undefinedEncoding := AccountCodingConfiguration{
	// 	Class          : "undefined",
	// 	Name           : "undefined",
	// 	NameCanonical : "Undefined",
	// 	NameEncoded   : "na",
	// }
	// thisResource.Configuration.AccountsCoding["noop"] = undefinedEncoding
	thisResource.AccountsCodingById = thisResource.Configuration.AccountsCoding

	// see https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	// excerpt: '... Allowed characters are letters, numbers, spaces representable in UTF-8, and the following characters: . : + = @ _ / - (hyphen) ...'
	prefixes := data.Prefixes{
		ReservedTagKey: ":", // Vvalues with 'aws' prefix are not allowed
	}

	separators := data.Separators{
		TagComponent:       ":",
		NameComponent:      "-",
		ComponentWords:     "",
		LogstreamComponent: "/",
	}

	constraints := data.Constraints{
		// NB: https://docs.aws.amazon.com/acm/latest/userguide/tags-restrictions.html
		TagKeyMaxLength:        127,
		TagValueMaxLength:      255,
		TagValueReservedPrefix: "aws:",
		// These may completely aritrary values
		EnvNameMaxLength:       20,
		VersionStringMinLength: (4 + 1), // Extra '+1' to account for the hyphen separator when assemblying later
		TypePrefixMaxLength:    6,

		NameTagKey:       "Name",
		CustomTagSection: "Custom",
		// Flip following switch for enabling or disabling section name in tag keys:
		// 1. Foo:Service
		// 2. Foo:Environment:Service
		IsUseSectionsInNames: true,
	}

	targetInfraEnv := ""
	if thisResource.Configuration.InfraEnvironment != "" {
		targetInfraEnv = thisResource.Configuration.InfraEnvironment
	} else {
		targetInfraEnv = thisResource.Configuration.AppEnvironment
	}

	targetAppEnv := ""
	if thisResource.Configuration.AppEnvironment != "" {
		targetAppEnv = thisResource.Configuration.AppEnvironment
	} else {
		targetAppEnv = thisResource.Configuration.InfraEnvironment
	}

	keys := data.Keys{
		InfraEnvironment: strings.ToLower(targetInfraEnv),
		AppEcosystem:     strings.ToLower(thisResource.Configuration.AppEcosystem),
		AppEnvironment:   strings.ToLower(targetAppEnv),
		AwsRegion:        strings.ToLower(thisResource.Configuration.Region),
		AwsZone:          strings.ToLower(thisResource.Configuration.AvailabilityZone),
	}

	placementRegion, err := parsePlacement("Region", keys.AwsRegion)
	if err != nil {
		return err
	}

	placementAz, err := parsePlacement("Availability zone", keys.AwsZone)
	if err != nil {
		return err
	}

	var awsPlacements = data.AwsPlacements{
		Region:           placementRegion,
		AvailabilityZone: placementAz,
	}

	accountIdUnsafe := thisResource.Configuration.Account
	// NB: removing leading zero because of terraform issue https://github.com/hashicorp/terraform/issues/28619
	canonicalAccountId := strings.TrimPrefix(accountIdUnsafe, "0")

	//// NB this falls back to an 'undefined' default, it's dangerous
	// if accountObj, ok := thisResource.AccountsCodingById[canonicalAccountId]; !ok {
	// 	accountObj := thisResource.AccountsCodingById["noop"]
	// }

	// This is preferrable: early failure if input config is not up to date
	accountObj, ok := thisResource.AccountsCodingById[canonicalAccountId]
	if !ok {
		return fmt.Errorf("[ERROR] Account ID '%s' is not in accounts_coding", accountIdUnsafe)
	}
	if _, ok := environmentClasses[accountObj.Class]; !ok {
		return fmt.Errorf("[ERROR] Account ID '%s' has class '%s' in accounts_coding. Use one of development, preproduction, production, management, undefined", accountIdUnsafe, accountObj.Class)
	}
	// resourceObj := this

	// We have both custom_name and custom_name_prefix, but want to use only one
	customNameToUse := thisResource.Configuration.CustomNamePrefix
	if customNameToUse == "" {
		customNameToUse = thisResource.Configuration.CustomName
	}

	//// If we specified custom_name, then don't use the hashed suffix (use custom_name as is, or use custom_name_prefix to enable the hashed suffix)
	var isCreateBeforeDestroy bool
	if customNameToUse != "" {
		isCreateBeforeDestroy = false
	} else {
		isCreateBeforeDestroy = thisResource.Configuration.IsCreateBeforeDestroy
	}

	/// We use the disambihuating hash prefix only if:
	//// 1. explicitly requested
	//// 2a. supported by resource type --> this gets in the way...
	// isVersionedName := (isCreateBeforeDestroy || thisResource.Properties.Name.IsEnforceVersioning)
	//// 2b. spcified a custom name prefix --> a prefix obviously needs a suffix, let's generate one
	isVersionedName := (isCreateBeforeDestroy || thisResource.Configuration.CustomNamePrefix != "")

	workspace := thisResource.Configuration.TerraformWorkspace
	if workspace == "default" {
		workspace = ""
	}

	//// Use custom_name if enabled and passed
	customName := ""
	if !thisResource.Configuration.IsForceGeneratedName && strings.Trim(customNameToUse, " ") != "" {
		customName = strings.TrimSpace(customNameToUse + " " + workspace)
	}

	canonicalRegion := ""
	if keys.AwsRegion != "" {
		canonicalRegion = awsPlacements.Region.Region + "" + awsPlacements.Region.Direction[0:1] + "" + awsPlacements.Region.Sequence
	}

	canonicalAz := ""
	if keys.AwsZone != "" {
		canonicalAz = awsPlacements.AvailabilityZone.Region + "" + awsPlacements.AvailabilityZone.Direction[0:1] + "" + awsPlacements.AvailabilityZone.Sequence
	}

	caser := cases.Title(language.English, cases.NoLower)
	//// Impose consistent word separator so that we can use it in resource names
	canonicalRole := caser.String(strings.Trim(strings.TrimPrefix(strings.ReplaceAll(thisResource.Configuration.Role, " ", separators.ComponentWords), separators.NameComponent), " "))
	canonicalRoleExtend := caser.String(strings.Trim(strings.TrimPrefix(strings.ReplaceAll(thisResource.Configuration.RoleExtend, " ", separators.ComponentWords), separators.NameComponent), " "))

	if canonicalRoleExtend != "" {
		canonicalRole = canonicalRole + separators.NameComponent + canonicalRoleExtend
	}

	//// From name module
	infraEnvironmentSafe := keys.InfraEnvironment[0:min(utf8.RuneCountInString(keys.InfraEnvironment), int(constraints.EnvNameMaxLength))]
	appEnvironmentSafe := keys.AppEnvironment[0:min(utf8.RuneCountInString(keys.AppEnvironment), int(constraints.EnvNameMaxLength))]

	thisResource.SanitisedConfiguration = &data.SanitisedConfiguration{
		AccountObject:             accountObj,
		AccountId:                 canonicalAccountId,
		AppEnvironment:            appEnvironmentSafe,
		CanonicalAvailabilityZone: canonicalAz,
		CanonicalRegion:           canonicalRegion,
		CanonicalRole:             canonicalRole,
		Constraints:               constraints,
		CustomName:                customName,
		EnvironmentClasses:        environmentClasses,
		InfraEnvironment:          infraEnvironmentSafe,
		IsVersionedName:           isVersionedName,
		Keys:                      keys,
		Prefixes:                  prefixes,
		Separators:                separators,
	}

	namer := new(Namer)
	namer.init(thisResource.Properties, thisResource.Configuration, thisResource.SanitisedConfiguration)
	namer.setName(ctx)

	thisResource.Name = namer.Name
	thisResource.Type = namer.Type

	tagger := new(Tagger)
	tagger.init(thisResource.Properties, thisResource.Configuration, thisResource.SanitisedConfiguration, namer.Name, namer.Type)
	tagger.generateTags(ctx)
	tagger.generateTagsAsMap(ctx)

	thisResource.Tags = tagger.Tags
	thisResource.TagsAsMap = tagger.TagsAsMap

	tflog.Debug(ctx, "resources - configuration generated")

	tflog.Trace(ctx, "resources - END info generation")

	return nil
}

func (thisResource *Resource) GetName() data.Name {
	return thisResource.Name
}

func (thisResource *Resource) GetType() data.Type {
	return thisResource.Type
}

func (thisResource *Resource) GetTags() map[string]string {
	return thisResource.Tags
}

func (thisResource *Resource) GetTagsAsMap() []map[string]string {
	return thisResource.TagsAsMap
}

func parsePlacement(label string, value string) (data.AwsPlacemenetEntity, error) {
	if value == "" {
		return data.AwsPlacemenetEntity{}, nil
	}

	parts := strings.Split(value, "-")
	if len(parts) < 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return data.AwsPlacemenetEntity{}, fmt.Errorf("[ERROR] %s '%s' is not in the form eu-west-1", label, value)
	}

	return data.AwsPlacemenetEntity{
		Region:    parts[0],
		Direction: parts[1],
		Sequence:  parts[2],
	}, nil
}
