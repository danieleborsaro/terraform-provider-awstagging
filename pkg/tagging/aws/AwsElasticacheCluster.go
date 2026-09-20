package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsElasticacheCluster struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsElasticacheCluster{}
)

func GetAwsElasticacheCluster() core.ResourceInterface {
	this := &AwsElasticacheCluster{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon ElastiCache",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::ElastiCache::Cluster",
			Short: "echc",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        false,
			MaxLength:           50,
			IsLowerCase:         true,
			IsEnforceVersioning: false,
			Components: []string{"account",
				"infra_environment",
				"project",
				"resource_set",
				"app_environment",
				"role"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_elasticache_cluster",
		},
	}

	return this
}
