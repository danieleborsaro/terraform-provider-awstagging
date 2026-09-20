package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsIamPolicy struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsIamPolicy{}
)

func GetAwsIamPolicy() core.ResourceInterface {
	this := &AwsIamPolicy{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS IAM",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::IAM::Policy",
			Short: "iampo",
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
			ResourceName: "aws_iam_policy",
		},
	}

	return this
}
