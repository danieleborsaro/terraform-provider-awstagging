package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsSesIdentityNotificationTopic struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsSesIdentityNotificationTopic{}
)

func GetAwsSesIdentityNotificationTopic() core.ResourceInterface {
	this := &AwsSesIdentityNotificationTopic{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Simple Email Service Identity Notification Topic",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::SES::IdentityNotificationTopic",
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
			ResourceName: "aws_ses_identity_notification_topic",
		},
	}

	return this
}
