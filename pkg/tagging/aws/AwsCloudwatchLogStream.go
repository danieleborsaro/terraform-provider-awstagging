package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsCloudwatchLogStream struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsCloudwatchLogStream{}
)

func GetAwsCloudwatchLogStream() core.ResourceInterface {
	this := &AwsCloudwatchLogStream{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Logs",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Logs::LogStream",
			Short: "logstrm",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        true,
			MaxLength:           255,
			IsLowerCase:         false,
			IsEnforceVersioning: false,
			Components: []string{"resource_set",
				"role",
				"platform",
				"compute_type"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_cloudwatch_log_stream",
		},
	}

	return this
}
