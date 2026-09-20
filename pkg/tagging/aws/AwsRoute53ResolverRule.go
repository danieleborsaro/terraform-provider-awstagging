package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsRoute53ResolverRule struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsRoute53ResolverRule{}
)

func GetAwsRoute53ResolverRule() core.ResourceInterface {
	this := &AwsRoute53ResolverRule{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Route 53 (Supported only in the US East (N. Virginia) Region, us-east-1.) ",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Route53::ResolverRule",
			Short: "r53rr",
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
			ResourceName: "aws_route53_resolver_rule",
		},
	}

	return this
}
