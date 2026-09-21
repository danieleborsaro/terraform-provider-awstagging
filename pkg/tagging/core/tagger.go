package tagging

import (
	"context"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

//nolint:unused
type automation struct {
	TerraformModule    string
	TerraformWorkspace string
}

//nolint:unused
type business struct {
	CostCentre string
	Owner      string
	Project    string
}

//nolint:unused
type security struct {
	Compliance string
}

//nolint:unused
type environment struct {
	Account          string
	AppEcosystem     string
	AppEnvironment   string
	Description      string
	InfraEnvironment string
	Name             string
	ResourceSet      string
	Role             string
	ResourceType     string
}

type Tagger struct {
	configuration          *data.InputConfiguration
	properties             *data.ResourceProperties
	sanitisedConfiguration *data.SanitisedConfiguration
	resourceName           data.Name
	resourceType           data.Type
	Tags                   map[string]string
	TagsAsMap              []map[string]string
	DroppedTags            []string
}

func (thisResource *Tagger) init(properties *data.ResourceProperties, config *data.InputConfiguration, sanitisedConfig *data.SanitisedConfiguration, resourceName data.Name, resourceType data.Type) {
	thisResource.configuration = config
	thisResource.properties = properties
	thisResource.resourceName = resourceName
	thisResource.resourceType = resourceType
	thisResource.sanitisedConfiguration = sanitisedConfig

	//// NB: Here on from Resource?
	// thisResource.generateTags()
	// thisResource.generateTagsAsMap()
}

func (thisResource *Tagger) generateTags(ctx context.Context) {
	tflog.Trace(ctx, "tagger - BEGIN tag generation")

	moduleAbs, _ := filepath.Abs(thisResource.configuration.TerraformModule)
	moduleName := filepath.Base(moduleAbs)

	caser := cases.Title(language.English, cases.NoLower)

	allTags := map[string]map[string]string{
		"automation": map[string]string{
			"TerraformModule":    moduleName,
			"TerraformWorkspace": thisResource.configuration.TerraformWorkspace,
		},

		"business": map[string]string{
			"CostCentre": strings.Trim(thisResource.configuration.CostCentre, " "),
			"Owner":      strings.Trim(thisResource.configuration.Owner, " "),
			"Project":    caser.String(thisResource.configuration.ProjectNameLong),
		},

		"security": map[string]string{
			"Compliance": strings.Trim(thisResource.configuration.Compliance, " "),
		},

		"environment": map[string]string{
			"Account":          thisResource.sanitisedConfiguration.AccountObject.NameCanonical,
			"AppEcosystem":     caser.String(thisResource.sanitisedConfiguration.Keys.AppEcosystem),
			"AppEnvironment":   caser.String(thisResource.sanitisedConfiguration.Keys.AppEnvironment),
			"Description":      strings.Trim(thisResource.configuration.Description, " "),
			"InfraEnvironment": caser.String(thisResource.sanitisedConfiguration.Keys.InfraEnvironment),
			"Name":             thisResource.resourceName.Safe,
			"ResourceSet":      caser.String(thisResource.configuration.ResourceSetLong),
			"Role":             caser.String(strings.Trim(strings.TrimPrefix(thisResource.sanitisedConfiguration.CanonicalRole, thisResource.sanitisedConfiguration.Separators.NameComponent), " ")),
			"ResourceType":     thisResource.resourceType.Long,
		},
	}

	flattenedTags := map[string]map[string]string{}
	for allK, allV := range allTags {
		for catK, catV := range allV {
			thisMap := make(map[string]string)
			thisMap["value"] = catV
			thisMap["section"] = allK
			flattenedTags[catK] = thisMap
		}
	}

	resourceSpecificTags := map[string]map[string]string{}
	if len(thisResource.properties.Tags.CustomList) == 0 {
		for k, v := range flattenedTags {
			resourceSpecificTags[k] = v
		}
	} else {
		for _, k := range thisResource.properties.Tags.CustomList {
			resourceSpecificTags[k] = flattenedTags[k]
		}
	}

	//// Assemble all tags as a single, plain map
	//// Two cases:
	//// 1. if desired, we need to prefix tag names with their section name (custom tags get 'custom')
	//// 2. otherwise just use tags
	//// In both cases we flatten the map of maps to a plan map containing tag => value pairs
	allRawTags := map[string]string{}
	if thisResource.sanitisedConfiguration.Constraints.IsUseSectionsInNames {
		for ks, vs := range thisResource.configuration.CustomTags {
			allRawTags[thisResource.sanitisedConfiguration.Constraints.CustomTagSection+""+thisResource.sanitisedConfiguration.Separators.TagComponent+""+caser.String(ks)] = vs
		}
		for kt, vt := range resourceSpecificTags {
			if kt != thisResource.sanitisedConfiguration.Constraints.NameTagKey {
				allRawTags[vt["section"]+""+thisResource.sanitisedConfiguration.Separators.TagComponent+""+caser.String(kt)] = vt["value"]
			} else {
				allRawTags[kt] = vt["value"]
			}
		}

	} else {
		for ks, vs := range thisResource.configuration.CustomTags {
			allRawTags[ks] = vs
		}
		for kt, vt := range resourceSpecificTags {
			allRawTags[kt] = vt["value"]
		}
	}

	//// Filter out tags with empty values
	allNonEmptyTags := map[string]string{}
	for k, v := range allRawTags {
		if v != "" {
			allNonEmptyTags[k] = v
		}
	}

	//// Apply case beautification to tag keys
	beautifiedKeys := make([]string, 0, len(allNonEmptyTags))
	for k := range allNonEmptyTags {
		beautifiedKeys = append(beautifiedKeys, k)
	}
	sort.Strings(beautifiedKeys)

	allBeautifiedTags := map[string]string{}
	for _, k := range beautifiedKeys {
		allBeautifiedTags[caser.String(k)] = allNonEmptyTags[k]
	}

	//// Prepend 'Foo:' to all tag keys, but not 'Name' so that it will show up nicely in AWS Console
	//// NB: 'Name' might be empty and already filtered out, so we have to check
	allPrefixedTags := map[string]string{}
	for k, v := range allBeautifiedTags {
		if k != thisResource.sanitisedConfiguration.Constraints.NameTagKey {
			allPrefixedTags[caser.String(thisResource.configuration.CompanyNameShort+""+thisResource.sanitisedConfiguration.Separators.TagComponent+""+k)] = v

		} else {
			allPrefixedTags[caser.String(thisResource.sanitisedConfiguration.Constraints.NameTagKey)] = v
		}
	}

	//// Add verbatim tags to map
	allTagsUnrestricted := map[string]string{}
	for k, v := range allPrefixedTags {
		allTagsUnrestricted[k] = v
	}
	for k, v := range thisResource.configuration.CustomTagsVerbatim {
		allTagsUnrestricted[k] = v
	}

	//// Apply AWS restrictions
	//// NB: see https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	allTagsRestrictedPrefixes := map[string]string{}
	escape := func(s string) string {
		//// If our tag key starts with an AWS-reserved prefix, then we need to tweak it by pre-ending a set string, e.g. aws:my-custom-tag --> :aws:my-custom-tag
		if strings.HasPrefix(strings.ToLower(s), strings.ToLower(thisResource.sanitisedConfiguration.Constraints.TagValueReservedPrefix)) {
			return thisResource.sanitisedConfiguration.Prefixes.ReservedTagKey + s
		}
		return s
	}
	for tk, tv := range allTagsUnrestricted {
		k := escape(tk)
		if _, ok := allTagsUnrestricted[k]; ok && k != tk {
			//// A key written as :aws:x wins over aws:x escaped to the same key
			continue
		}
		allTagsRestrictedPrefixes[k] = escape(tv)
	}

	allTagsMaxLength := map[string]string{}
	for tk, tv := range allTagsRestrictedPrefixes {
		allTagsMaxLength[tk[0:0+min(len(tk), int(thisResource.sanitisedConfiguration.Constraints.TagKeyMaxLength))]] = tv[0 : 0+min(len(tv), int(thisResource.sanitisedConfiguration.Constraints.TagValueMaxLength))]
	}

	//// This is quite rough: just take the first n tags off the ordered list of tag names... (i.e. S3 objects accept max 10 tags)
	generatedPrefix := caser.String(thisResource.configuration.CompanyNameShort + thisResource.sanitisedConfiguration.Separators.TagComponent)
	customPrefix := generatedPrefix + thisResource.sanitisedConfiguration.Constraints.CustomTagSection + thisResource.sanitisedConfiguration.Separators.TagComponent
	rank := func(k string) int {
		switch {
		case k == thisResource.sanitisedConfiguration.Constraints.NameTagKey:
			return 0
		case strings.HasPrefix(k, customPrefix):
			return 2
		case strings.HasPrefix(k, generatedPrefix):
			return 1
		default:
			return 3
		}
	}

	orderedKeys := make([]string, 0, len(allTagsMaxLength))
	for k := range allTagsMaxLength {
		orderedKeys = append(orderedKeys, k)
	}
	sort.Slice(orderedKeys, func(i, j int) bool {
		if rank(orderedKeys[i]) != rank(orderedKeys[j]) {
			return rank(orderedKeys[i]) < rank(orderedKeys[j])
		}
		return orderedKeys[i] < orderedKeys[j]
	})

	allTagsMaxNumber := map[string]string{}
	for i, k := range orderedKeys {
		if uint(i) < thisResource.properties.Tags.Max {
			allTagsMaxNumber[k] = allTagsMaxLength[k]
		} else {
			thisResource.DroppedTags = append(thisResource.DroppedTags, k)
		}
	}

	thisResource.Tags = allTagsMaxNumber

	tflog.Debug(ctx, "tagger - tags generated")

	tflog.Trace(ctx, "tagger - END tag generation")
}

func (thisResource *Tagger) generateTagsAsMap(ctx context.Context) {
	tflog.Trace(ctx, "tagger - BEGIN tag as map generation")

	thisResource.TagsAsMap = []map[string]string{}
	for k, v := range thisResource.Tags {
		mapBUffer := make(map[string]string)
		mapBUffer["Key"] = k
		mapBUffer["Value"] = v
		thisResource.TagsAsMap = append(thisResource.TagsAsMap, mapBUffer)
	}

	tflog.Debug(ctx, "tagger - tags as map generated")

	tflog.Trace(ctx, "tagger - END tag as map generation")
}
