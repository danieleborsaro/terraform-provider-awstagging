package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsCloudwatchCompositeAlarm struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsCloudwatchCompositeAlarm{}
)

func GetAwsCloudwatchCompositeAlarm() core.ResourceInterface {
	this := &AwsCloudwatchCompositeAlarm{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon CloudWatch",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::CloudWatch::CompositeAlarm",
			Short: "cwca",
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
				"role",
				"region"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_cloudwatch_composite_alarm",
		},
	}

	return this
}
