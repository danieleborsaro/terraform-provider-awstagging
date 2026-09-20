package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsNetworkInterface struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsNetworkInterface{}
)

func GetAwsNetworkInterface() core.ResourceInterface {
	this := &AwsNetworkInterface{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Elastic Compute Cloud (Amazon EC2)",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::EC2::NetworkInterface",
			Short: "ni",
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
			Components:          []string{},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_network_interface",
		},
	}

	return this
}
