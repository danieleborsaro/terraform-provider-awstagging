package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsIamGroup struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsIamGroup{}
)

func GetAwsIamGroup() core.ResourceInterface {
	this := &AwsIamGroup{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS IAM",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::IAM::Group",
			Short: "iamg",
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
				"role"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_iam_group",
		},
	}

	return this
}
