package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsRoute53HealthCheck struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsRoute53HealthCheck{}
)

func GetAwsRoute53HealthCheck() core.ResourceInterface {
	this := &AwsRoute53HealthCheck{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Route 53 (Supported only in the US East (N. Virginia) Region, us-east-1.) ",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Route53::HealthCheck",
			Short: "r53hc",
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
			ResourceName: "aws_route53_health_check",
		},
	}

	return this
}
