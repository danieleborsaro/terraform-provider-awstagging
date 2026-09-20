package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsElasticacheSubnetGroup struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsElasticacheSubnetGroup{}
)

func GetAwsElasticacheSubnetGroup() core.ResourceInterface {
	this := &AwsElasticacheSubnetGroup{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon ElastiCache Subnet group",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::ElastiCache::Subnet group",
			Short: "ecsg",
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
			ResourceName: "aws_elasticache_subnet_group",
		},
	}

	return this
}
