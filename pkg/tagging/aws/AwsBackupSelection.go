package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsBackupSelection struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsBackupSelection{}
)

func GetAwsBackupSelection() core.ResourceInterface {
	this := &AwsBackupSelection{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS Backup Selection",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Backup::BackupSelection",
			Short: "back-s",
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
			ResourceName: "aws_backup_selection",
		},
	}

	return this
}
