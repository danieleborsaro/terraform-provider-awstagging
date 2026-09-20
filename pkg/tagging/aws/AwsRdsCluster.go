package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsRdsCluster struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsRdsCluster{}
)

func GetAwsRdsCluster() core.ResourceInterface {
	this := &AwsRdsCluster{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Relational Database Service (Amazon RDS)",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::RDS::DBCluster",
			Short: "rdsc",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        false,
			MaxLength:           60,
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
			ResourceName: "aws_rds_cluster",
		},
	}

	return this
}
