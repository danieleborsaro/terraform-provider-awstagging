package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsGlacierVault struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsGlacierVault{}
)

func GetAwsGlacierVault() core.ResourceInterface {
	this := &AwsGlacierVault{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon S3 Glacier",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Glacier::Vault",
			Short: "glac",
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
			ResourceName: "aws_glacier_vault",
		},
	}

	return this
}
