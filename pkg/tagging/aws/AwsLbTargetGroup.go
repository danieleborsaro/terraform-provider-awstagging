package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsLbTargetGroup struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsLbTargetGroup{}
)

func GetAwsLbTargetGroup() core.ResourceInterface {
	this := &AwsLbTargetGroup{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Elastic Load Balancing",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::ElasticLoadBalancingV2::TargetGroup",
			Short: "tg",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        false,
			MaxLength:           32,
			IsLowerCase:         false,
			IsEnforceVersioning: true,
			Components: []string{"account",
				"infra_environment",
				"project",
				"resource_set",
				"app_environment",
				"role"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_lb_target_group",
		},
	}

	return this
}
