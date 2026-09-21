package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsWafWebAcl struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsWafWebAcl{}
)

func GetAwsWafWebAcl() core.ResourceInterface {
	this := &AwsWafWebAcl{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Web Application Firewall",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::WAF::WebACL",
			Short: "wafacl",
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
			ResourceName: "aws_waf_web_acl",
		},
	}

	return this
}
