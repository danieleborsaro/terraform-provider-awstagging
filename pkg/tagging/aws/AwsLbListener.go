package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsLbListener struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsLbListener{}
)

func GetAwsLbListener() core.ResourceInterface {
	this := &AwsLbListener{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Elastic Load Balancing",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::ElasticLoadbalancingV2::Listener",
			Short: "lbl",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        true,
			MaxLength:           255,
			IsLowerCase:         false,
			IsEnforceVersioning: true,
			Components:          []string{},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_lb_listener",
		},
	}

	return this
}
