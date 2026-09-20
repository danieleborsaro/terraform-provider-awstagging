package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsSecurityGroup struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsSecurityGroup{}
)

func GetAwsSecurityGroup() core.ResourceInterface {
	this := &AwsSecurityGroup{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Elastic Compute Cloud (Amazon EC2)",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::EC2::SecurityGroup",
			Short: "sg",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        true,
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
			ResourceName: "aws_security_group",
		},
	}

	return this
}
