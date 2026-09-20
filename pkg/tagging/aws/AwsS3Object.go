package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsS3Object struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsS3Object{}
)

func GetAwsS3Object() core.ResourceInterface {
	this := &AwsS3Object{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Simple Storage Service (Amazon S3)",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::S3::BucketObject",
			Short: "s3",
		},
		Tags: data.ResourcePropertiesTags{
			Max: 10,
			CustomList: []string{"AppEnvironment",
				"AppEcosystem",
				"InfraEnvironment",
				"Name",
				"Owner"},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        false,
			MaxLength:           63,
			IsLowerCase:         true,
			IsEnforceVersioning: false,
			Components:          []string{},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_s3_object",
		},
	}

	return this
}
