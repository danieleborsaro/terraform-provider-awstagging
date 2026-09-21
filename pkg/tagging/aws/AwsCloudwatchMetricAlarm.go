package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsCloudwatchMetricAlarm struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsCloudwatchMetricAlarm{}
)

func GetAwsCloudwatchMetricAlarm() core.ResourceInterface {
	this := &AwsCloudwatchMetricAlarm{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon CloudWatch",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::CloudWatch::Alarm",
			Short: "cwma",
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
			ResourceName: "aws_cloudwatch_metric_alarm",
		},
	}

	return this
}
