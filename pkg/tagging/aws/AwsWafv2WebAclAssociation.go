package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsWafv2WebAclAssociation struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsWafv2WebAclAssociation{}
)

func GetAwsWafv2WebAclAssociation() core.ResourceInterface {
	this := &AwsWafv2WebAclAssociation{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Web Application Firewall v2",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::WAFv2::WebACLAssociation",
			Short: "waf2asc",
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
			ResourceName: "aws_wafv2_web_acl_association",
		},
	}

	return this
}
