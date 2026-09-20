package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsEc2ManagedPrefixList struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsEc2ManagedPrefixList{}
)

func GetAwsEc2ManagedPrefixList() core.ResourceInterface {
	this := &AwsEc2ManagedPrefixList{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Elastic Compute Cloud (Amazon EC2)",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::EC2::ManagedPrefixList",
			Short: "mpl",
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
				"resource_type",
				"role",
				"region"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_ec2_managed_prefix_list",
		},
	}

	return this
}
