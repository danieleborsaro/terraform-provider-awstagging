package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsOpensearchDomain struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsOpensearchDomain{}
)

func GetAwsOpensearchDomain() core.ResourceInterface {
	this := &AwsOpensearchDomain{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon OpenSearch Service",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::OpenSearchService::Domain",
			Short: "os",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        false,
			MaxLength:           28,
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
			ResourceName: "aws_opensearch_domain",
		},
	}

	return this
}
