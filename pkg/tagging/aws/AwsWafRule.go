package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsWafRule struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsWafRule{}
)

func GetAwsWafRule() core.ResourceInterface {
	this := &AwsWafRule{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Web Application Firewall",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::WAF::Rule",
			Short: "wafrule",
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
			ResourceName: "aws_waf_rule",
		},
	}

	return this
}
