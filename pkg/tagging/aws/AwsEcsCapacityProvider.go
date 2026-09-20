package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsEcsCapacityProvider struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsEcsCapacityProvider{}
)

func GetAwsEcsCapacityProvider() core.ResourceInterface {
	this := &AwsEcsCapacityProvider{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Elastic Container Service",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::ECS::CapacityProvider",
			Short: "cap",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        false,
			MaxLength:           255,
			IsLowerCase:         false,
			IsEnforceVersioning: true,
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
			ResourceName: "aws_ecs_capacity_provider",
		},
	}

	return this
}
