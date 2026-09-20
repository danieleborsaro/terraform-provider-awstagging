package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsEbsVolume struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsEbsVolume{}
)

func GetAwsEbsVolume() core.ResourceInterface {
	this := &AwsEbsVolume{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Elastic Block Storage",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::EBS::Volume",
			Short: "ebs",
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
				"availability_zone"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_ebs_volume",
		},
	}

	return this
}
