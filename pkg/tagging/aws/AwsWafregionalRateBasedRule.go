package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsWafregionalRateBasedRule struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsWafregionalRateBasedRule{}
)

func GetAwsWafregionalRateBasedRule() core.ResourceInterface {
	this := &AwsWafregionalRateBasedRule{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Web Application Firewall Regional",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::WAFRegional::RateBasedRule",
			Short: "wafrrate",
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
			ResourceName: "aws_wafregional_rate_based_rule",
		},
	}

	return this
}
