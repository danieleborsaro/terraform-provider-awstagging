package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsSsmParameter struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsSsmParameter{}
)

func GetAwsSsmParameter() core.ResourceInterface {
	this := &AwsSsmParameter{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS Systems Manager",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::SSM::Parameter",
			Short: "sm",
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
			ResourceName: "aws_ssm_parameter",
		},
	}

	return this
}
