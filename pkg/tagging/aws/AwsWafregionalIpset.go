package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsWafregionalIpset struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsWafregionalIpset{}
)

func GetAwsWafregionalIpset() core.ResourceInterface {
	this := &AwsWafregionalIpset{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Web Application Firewall",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::WAFRegional::IPSet",
			Short: "wafipset",
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
				"role",
				"region"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_wafregional_ipset",
		},
	}

	return this
}
