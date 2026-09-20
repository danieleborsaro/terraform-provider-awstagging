package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsCloudformationExport struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsCloudformationExport{}
)

func GetAwsCloudformationExport() core.ResourceInterface {
	this := &AwsCloudformationExport{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS CloudFormation",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::CloudFormation::Export",
			Short: "cfnex",
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
			ResourceName: "aws_cloudformation_export",
		},
	}

	return this
}
