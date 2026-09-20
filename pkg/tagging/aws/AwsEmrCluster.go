package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsEmrCluster struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsEmrCluster{}
)

func GetAwsEmrCluster() core.ResourceInterface {
	this := &AwsEmrCluster{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon EMR",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::EMR::Cluster",
			Short: "emr",
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
			ResourceName: "aws_emr_cluster",
		},
	}

	return this
}
