package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsS3Bucket struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsS3Bucket{}
)

func GetAwsS3Bucket() core.ResourceInterface {
	this := &AwsS3Bucket{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Simple Storage Service (Amazon S3)",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::S3::Bucket",
			Short: "s3",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        false,
			MaxLength:           63,
			IsLowerCase:         true,
			IsEnforceVersioning: false,
			Components: []string{"business",
				"account",
				"infra_environment",
				"project",
				"resource_set",
				"app_environment",
				"role",
				"region"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_s3_bucket",
		},
	}

	return this
}
