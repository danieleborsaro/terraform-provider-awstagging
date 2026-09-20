package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsRoute53Zone struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsRoute53Zone{}
)

func GetAwsRoute53Zone() core.ResourceInterface {
	this := &AwsRoute53Zone{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Route 53 (Supported only in the US East (N. Virginia) Region, us-east-1.) ",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Route53::HostedZone",
			Short: "r53hz",
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
			ResourceName: "aws_route53_zone",
		},
	}

	return this
}
