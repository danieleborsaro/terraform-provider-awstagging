package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsWafv2WebAclLoggingConfiguration struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsWafv2WebAclLoggingConfiguration{}
)

func GetAwsWafv2WebAclLoggingConfiguration() core.ResourceInterface {
	this := &AwsWafv2WebAclLoggingConfiguration{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Web Application Firewall v2",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::WAFv2::Logging Configuration",
			Short: "waf2log",
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
			ResourceName: "aws_wafv2_web_acl_logging_configuration",
		},
	}

	return this
}
