package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsDirectoryServiceDirectory struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsDirectoryServiceDirectory{}
)

func GetAwsDirectoryServiceDirectory() core.ResourceInterface {
	this := &AwsDirectoryServiceDirectory{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Directory Service",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::DirectoryService::MicrosoftAD",
			Short: "dsmsad",
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
			ResourceName: "aws_directory_service_directory",
		},
	}

	return this
}
