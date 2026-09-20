package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsEfsAccessPoint struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsEfsAccessPoint{}
)

func GetAwsEfsAccessPoint() core.ResourceInterface {
	this := &AwsEfsAccessPoint{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Elastic File System",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::EFS::AccessPoint",
			Short: "efsap",
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
			ResourceName: "aws_efs_access_point",
		},
	}

	return this
}
