package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsCloudfrontDistribution struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsCloudfrontDistribution{}
)

func GetAwsCloudfrontDistribution() core.ResourceInterface {
	this := &AwsCloudfrontDistribution{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon CloudFront",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::CloudFront::Distribution",
			Short: "cfd",
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
			ResourceName: "aws_cloudfront_distribution",
		},
	}

	return this
}
