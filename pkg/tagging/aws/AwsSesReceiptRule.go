package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsSesReceiptRule struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsSesReceiptRule{}
)

func GetAwsSesReceiptRule() core.ResourceInterface {
	this := &AwsSesReceiptRule{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Simple Email Service Receipt Rule",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::SES::ReceiptRule",
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
			ResourceName: "aws_ses_receipt_rule",
		},
	}

	return this
}
