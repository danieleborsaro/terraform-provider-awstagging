package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsWafregionalRegexPatternSet struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsWafregionalRegexPatternSet{}
)

func GetAwsWafregionalRegexPatternSet() core.ResourceInterface {
	this := &AwsWafregionalRegexPatternSet{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Web Application Firewall Regional",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::WAFRegional::RegexPatternSet",
			Short: "wafrset",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        false,
			MaxLength:           255,
			IsLowerCase:         false,
			IsEnforceVersioning: false,
			Components: []string{"account",
				"infra_environment",
				"project",
				"resource_set",
				"app_environment",
				"resource_type",
				"role"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_wafregional_regex_pattern_set",
		},
	}

	return this
}
