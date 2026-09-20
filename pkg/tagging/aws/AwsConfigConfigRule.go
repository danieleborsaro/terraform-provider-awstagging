package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsConfigConfigRule struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsConfigConfigRule{}
)

func GetAwsConfigConfigRule() core.ResourceInterface {
	this := &AwsConfigConfigRule{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS Config",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Config::Rule",
			Short: "cfgr",
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
			ResourceName: "aws_config_config_rule",
		},
	}

	return this
}
