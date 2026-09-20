package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsEc2TransitGatewayVpcAttachment struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsEc2TransitGatewayVpcAttachment{}
)

func GetAwsEc2TransitGatewayVpcAttachment() core.ResourceInterface {
	this := &AwsEc2TransitGatewayVpcAttachment{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Transit Gateway Attachment",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::EC2::TransitGatewayAttachment",
			Short: "tgwat",
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
			ResourceName: "aws_ec2_transit_gateway_vpc_attachment",
		},
	}

	return this
}
