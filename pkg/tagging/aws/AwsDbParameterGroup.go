package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsDbParameterGroup struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsDbParameterGroup{}
)

func GetAwsDbParameterGroup() core.ResourceInterface {
	this := &AwsDbParameterGroup{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Relational Database Service (Amazon RDS)",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::RDS::DBParameterGroup",
			Short: "rdsp",
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
			ResourceName: "aws_db_parameter_group",
		},
	}

	return this
}
