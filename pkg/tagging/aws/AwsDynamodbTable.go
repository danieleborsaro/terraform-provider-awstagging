package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsDynamodbTable struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsDynamodbTable{}
)

func GetAwsDynamodbTable() core.ResourceInterface {
	this := &AwsDynamodbTable{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon DynamoDB",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::DynamoDB::Table",
			Short: "dyn",
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
				"role"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_dynamodb_table",
		},
	}

	return this
}
