package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsSesActiveReceiptRuleSet struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsSesActiveReceiptRuleSet{}
)

func GetAwsSesActiveReceiptRuleSet() core.ResourceInterface {
	this := &AwsSesActiveReceiptRuleSet{}

	this.Properties = &data.ResourceProperties{
		AwsService: "Amazon Simple Email Service Active Receipt Rule Set",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::SES::ActiveReceiptRuleSet",
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
			ResourceName: "aws_ses_active_receipt_rule_set",
		},
	}

	return this
}
