package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsCodebuildProject struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsCodebuildProject{}
)

func GetAwsCodebuildProject() core.ResourceInterface {
	this := &AwsCodebuildProject{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS CodeBuild",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::CodeBuild::Project",
			Short: "cb",
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
			Components:          []string{},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_codebuild_project",
		},
	}

	return this
}
