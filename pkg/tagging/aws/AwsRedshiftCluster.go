package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsRedshiftCluster struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsRedshiftCluster{}
)

func GetAwsRedshiftCluster() core.ResourceInterface {
	this := &AwsRedshiftCluster{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Redshift",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Redshift::Cluster",
			Short: "rshc",
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
			ResourceName: "aws_redshift_cluster",
		},
	}

	return this
}
