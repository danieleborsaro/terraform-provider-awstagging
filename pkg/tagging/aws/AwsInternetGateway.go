package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsInternetGateway struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsInternetGateway{}
)

func GetAwsInternetGateway() core.ResourceInterface {
	this := &AwsInternetGateway{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Virtual Private Cloud",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::VPC::InternetGateway",
			Short: "igw",
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
			ResourceName: "aws_internet_gateway",
		},
	}

	return this
}
