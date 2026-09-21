package aws

import (
	"reflect"

	core "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	data "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type AwsBackupVault struct{ core.Resource }

var (
	_ core.ResourceInterface = &AwsBackupVault{}
)

func GetAwsBackupVault() core.ResourceInterface {
	this := &AwsBackupVault{}

	this.Properties = &data.ResourceProperties{
		AwsService: "AWS Backup Vault",
		Id: data.ResourcePropertiesId{
			Key:   reflect.TypeOf(*this).Name(),
			Long:  "AWS::Backup::BackupVault",
			Short: "back-v",
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
			Components: []string{"infra_environment",
				"project",
				"resource_set",
				"app_environment",
				"resource_type",
				"role",
				"region"},
		},
		Terraform: data.ResourcePropertiesTerraform{
			ResourceName: "aws_backup_vault",
		},
	}

	return this
}
