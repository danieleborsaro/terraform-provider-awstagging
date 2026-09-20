package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsCognitoIdentityPool struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsCognitoIdentityPool{}
)

func GetAwsCognitoIdentityPool() core.ResourceInterface {
	this := &AwsCognitoIdentityPool{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Cognito",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Cognito::IdentityPool",
			Short: "cgip",
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
			ResourceName: "aws_cognito_identity_pool",
		},
	}

	return this
}
