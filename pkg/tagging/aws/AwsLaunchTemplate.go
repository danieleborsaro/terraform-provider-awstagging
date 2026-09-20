package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsLaunchTemplate struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsLaunchTemplate{}
)

func GetAwsLaunchTemplate() core.ResourceInterface {
	this := &AwsLaunchTemplate{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS::AutoScaling",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::AutoScaling::Launchtemplate",
			Short: "asglt",
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
				"platform",
				"role"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_launch_template",
		},
	}

	return this
}
