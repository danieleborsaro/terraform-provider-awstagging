package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsWafv2RegexPatternSet struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsWafv2RegexPatternSet{}
)

func GetAwsWafv2RegexPatternSet() core.ResourceInterface {
	this := &AwsWafv2RegexPatternSet{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Web Application Firewall v2",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::WAFv2::RegexPatternSet",
			Short: "waf2regx",
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
			ResourceName: "aws_wafv2_regex_pattern_set",
		},
	}

	return this
}
