package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsEc2ClientVpnEndpoint struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsEc2ClientVpnEndpoint{}
)

func GetAwsEc2ClientVpnEndpoint() core.ResourceInterface {
	this := &AwsEc2ClientVpnEndpoint{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Virtual Private Cloud",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::VPC::VPNCustomerEndpoint",
			Short: "vpnep",
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
			ResourceName: "aws_ec2_client_vpn_endpoint",
		},
	}

	return this
}
