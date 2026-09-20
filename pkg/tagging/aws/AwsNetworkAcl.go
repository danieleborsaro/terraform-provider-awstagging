package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsNetworkAcl struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsNetworkAcl{}
)

func GetAwsNetworkAcl() core.ResourceInterface {
	this := &AwsNetworkAcl{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Virtual Private Cloud",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::VPC::NetworkACL",
			Short: "nacl",
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
				"resource_type",
				"role",
				"availability_zone"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_network_acl",
		},
	}

	return this
}
