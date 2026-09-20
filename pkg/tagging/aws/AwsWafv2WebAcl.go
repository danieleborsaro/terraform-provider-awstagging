package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsWafv2WebAcl struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsWafv2WebAcl{}
)

func GetAwsWafv2WebAcl() core.ResourceInterface {
	this := &AwsWafv2WebAcl{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Web Application Firewall v2",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::WAFv2::WebACL",
			Short: "waf2acl",
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
			ResourceName: "aws_wafv2_web_acl",
		},
	}

	return this
}
