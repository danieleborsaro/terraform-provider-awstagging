package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsEksCluster struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsEksCluster{}
)

func GetAwsEksCluster() core.ResourceInterface {
	this := &AwsEksCluster{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Elastic Kubernetes Service",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::EKS::Cluster",
			Short: "eks",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        false,
			MaxLength:           100,
			IsLowerCase:         false,
			IsEnforceVersioning: false,
			Components: []string{"account",
				"infra_environment",
				"project",
				"resource_set",
				"app_environment",
				"resource_type",
				"role",
				"region"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_eks_cluster",
		},
	}

	return this
}
