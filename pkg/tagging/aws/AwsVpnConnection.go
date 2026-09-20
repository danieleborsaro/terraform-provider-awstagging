package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsVpnConnection struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsVpnConnection{}
)

func GetAwsVpnConnection() core.ResourceInterface {
	this := &AwsVpnConnection{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Virtual Private Cloud",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::VPC::VPNConnection",
			Short: "vpn",
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
			ResourceName: "aws_vpn_connection",
		},
	}

	return this
}
