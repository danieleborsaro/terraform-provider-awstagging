package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsSqsQueue struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsSqsQueue{}
)

func GetAwsSqsQueue() core.ResourceInterface {
	this := &AwsSqsQueue{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Simple Queue Service",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::SQS::Queue",
			Short: "sqs",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        false,
			MaxLength:           80,
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
			ResourceName: "aws_sqs_queue",
		},
	}

	return this
}
