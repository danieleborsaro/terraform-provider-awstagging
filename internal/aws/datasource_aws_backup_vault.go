package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsBackupVaultDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsBackupVaultDataSource{}
	_ datasource.DataSource                  = &awsBackupVaultDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsBackupVaultDataSource{}
)

func NewAwsBackupVaultDataSource() datasource.DataSource {
	newDatasource := &awsBackupVaultDataSource{}

	newDatasource.Tagging = tagging.GetAwsBackupVault()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
