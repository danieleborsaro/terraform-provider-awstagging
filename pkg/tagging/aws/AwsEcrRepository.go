package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsEcrRepository struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsEcrRepository{}
)

func GetAwsEcrRepository() core.ResourceInterface {
	this := &AwsEcrRepository{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Elastic Container Registry",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::ECR::Repository",
			Short: "ecr",
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
			ResourceName: "aws_ecr_repository",
		},
	}

	return this
}
