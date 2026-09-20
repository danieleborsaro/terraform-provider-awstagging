package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsSesEmailIdentity struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsSesEmailIdentity{}
)

func GetAwsSesEmailIdentity() core.ResourceInterface {
	this := &AwsSesEmailIdentity{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Simple Email Service Email Identity",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::SES::EmailIdentity",
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
			Components:          []string{},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_ses_email_identity",
		},
	}

	return this
}
