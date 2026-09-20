package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsSesv2ConfigurationSetEventDestination struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsSesv2ConfigurationSetEventDestination{}
)

func GetAwsSesv2ConfigurationSetEventDestination() core.ResourceInterface {
	this := &AwsSesv2ConfigurationSetEventDestination{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Simple Email Service Configuration Set Event Destination",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::SES::ConfigurationSetEventDestination",
			Short: "ses",
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
			ResourceName: "aws_sesv2_configuration_set_event_destination",
		},
	}

	return this
}
