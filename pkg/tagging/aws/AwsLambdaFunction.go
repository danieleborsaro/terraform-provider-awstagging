package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsLambdaFunction struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsLambdaFunction{}
)

func GetAwsLambdaFunction() core.ResourceInterface {
	this := &AwsLambdaFunction{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS Lambda",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Lambda::Function",
			Short: "lambda",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        true,
			MaxLength:           64,
			IsLowerCase:         false,
			IsEnforceVersioning: false,
			Components: []string{"account",
				"infra_environment",
				"project",
				"resource_set",
				"app_environment",
				"role"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_lambda_function",
		},
	}

	return this
}
