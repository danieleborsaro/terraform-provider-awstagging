package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsVpcDhcpOptionsAssociation struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsVpcDhcpOptionsAssociation{}
)

func GetAwsVpcDhcpOptionsAssociation() core.ResourceInterface {
	this := &AwsVpcDhcpOptionsAssociation{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Virtual Private Cloud",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::EC2::VPCDHCPOptionsAssociation",
			Short: "dhcpa",
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
				"resource_type",
				"role",
				"region"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_vpc_dhcp_options_association",
		},
	}

	return this
}
