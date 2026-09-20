package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsCloudwatchEventRule struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsCloudwatchEventRule{}
)

func GetAwsCloudwatchEventRule() core.ResourceInterface {
	this := &AwsCloudwatchEventRule{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon CloudWatch Events",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Events::Rule",
			Short: "cwr",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        false,
			MaxLength:           64,
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
			ResourceName: "aws_cloudwatch_event_rule",
		},
	}

	return this
}
