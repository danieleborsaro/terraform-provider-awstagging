package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsAutoscalingGroup struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsAutoscalingGroup{}
)

func GetAwsAutoscalingGroup() core.ResourceInterface {
	this := &AwsAutoscalingGroup{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS::AutoScaling",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::AutoScaling::AutoScalingGroup",
			Short: "asg",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        true,
			MaxLength:           255,
			IsLowerCase:         false,
			IsEnforceVersioning: false,
			Components: []string{"account",
				"infra_environment",
				"project",
				"resource_set",
				"app_environment",
				"resource_type",
				"platform",
				"role"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_autoscaling_group",
		},
	}

	return this
}
