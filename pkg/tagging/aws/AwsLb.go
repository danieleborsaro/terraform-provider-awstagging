package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsLb struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsLb{}
)

func GetAwsLb() core.ResourceInterface {
	this := &AwsLb{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Elastic Load Balancing",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::ElasticLoadBalancingV2::LoadBalancer",
			Short: "lb",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        true,
			MaxLength:           32,
			IsLowerCase:         false,
			IsEnforceVersioning: true,
			Components: []string{"account",
				"infra_environment",
				"project",
				"resource_set",
				"app_environment",
				"resource_type",
				"role"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_lb",
		},
	}

	return this
}
