package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsLbListenerRule struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsLbListenerRule{}
)

func GetAwsLbListenerRule() core.ResourceInterface {
	this := &AwsLbListenerRule{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Elastic Load Balancing",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::ElasticLoadbalancingV2::ListenerRule",
			Short: "lblr",
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
			ResourceName: "aws_lb_listener_rule",
		},
	}

	return this
}
