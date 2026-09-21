package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsRouteTable struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsRouteTable{}
)

func GetAwsRouteTable() core.ResourceInterface {
	this := &AwsRouteTable{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Virtual Private Cloud",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::EC2::RouteTable",
			Short: "rt",
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
				"availability_zone"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_route_table",
		},
	}

	return this
}
