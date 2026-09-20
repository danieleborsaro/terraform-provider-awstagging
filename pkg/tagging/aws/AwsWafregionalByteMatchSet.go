package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsWafregionalByteMatchSet struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsWafregionalByteMatchSet{}
)

func GetAwsWafregionalByteMatchSet() core.ResourceInterface {
	this := &AwsWafregionalByteMatchSet{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Web Application Firewall Regional",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::WAFRegional::ByteMatchRule",
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
			ResourceName: "aws_wafregional_byte_match_set",
		},
	}

	return this
}
