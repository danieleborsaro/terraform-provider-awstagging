package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsRamResourceShare struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsRamResourceShare{}
)

func GetAwsRamResourceShare() core.ResourceInterface {
	this := &AwsRamResourceShare{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS Resource Access Manager",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::RAM::ResourceShare",
			Short: "rs",
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
				"role",
				"region"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_ram_resource_share",
		},
	}

	return this
}
