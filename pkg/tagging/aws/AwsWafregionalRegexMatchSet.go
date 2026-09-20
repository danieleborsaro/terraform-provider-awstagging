package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsWafregionalRegexMatchSet struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsWafregionalRegexMatchSet{}
)

func GetAwsWafregionalRegexMatchSet() core.ResourceInterface {
	this := &AwsWafregionalRegexMatchSet{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Web Application Firewall Regional",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::WAFRegional::RegexMatch",
			Short: "wafrmatch",
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
			ResourceName: "aws_wafregional_regex_match_set",
		},
	}

	return this
}
