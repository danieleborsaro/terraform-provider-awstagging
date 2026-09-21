package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsBackupPlan struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsBackupPlan{}
)

func GetAwsBackupPlan() core.ResourceInterface {
	this := &AwsBackupPlan{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS Backup Plan",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Backup::BackupPlan",
			Short: "back-p",
		},
		Tags: data.ResourcePropertiesTags{
			Max:        50,
			CustomList: []string{},
		},
		Name: data.ResourcePropertiesName{
			IsLongPrefix:        false,
			MaxLength:           50,
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
			ResourceName: "aws_backup_plan",
		},
	}

	return this
}
