package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsKmsKey struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsKmsKey{}
)

func GetAwsKmsKey() core.ResourceInterface {
	this := &AwsKmsKey{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS Key Management Service",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::KMS::Key",
			Short: "kms",
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
			ResourceName: "aws_kms_key",
		},
	}

	return this
}
