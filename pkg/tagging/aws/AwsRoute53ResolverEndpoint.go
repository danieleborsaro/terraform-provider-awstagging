package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsRoute53ResolverEndpoint struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsRoute53ResolverEndpoint{}
)

func GetAwsRoute53ResolverEndpoint() core.ResourceInterface {
	this := &AwsRoute53ResolverEndpoint{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Route 53 (Supported only in the US East (N. Virginia) Region, us-east-1.) ",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Route53Resolver::ResolverEndpoint",
			Short: "r53re",
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
				"role"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_route53_resolver_endpoint",
		},
	}

	return this
}
