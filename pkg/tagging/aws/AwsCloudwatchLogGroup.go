package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsCloudwatchLogGroup struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsCloudwatchLogGroup{}
)

func GetAwsCloudwatchLogGroup() core.ResourceInterface {
	this := &AwsCloudwatchLogGroup{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Logs",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Logs::LogGroup",
			Short: "loggrp",
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
			ResourceName: "aws_cloudwatch_log_group",
		},
	}

	return this
}
