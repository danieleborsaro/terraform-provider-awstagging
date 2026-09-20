package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsBackupSelectionDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsBackupSelectionDataSource{}
	_ datasource.DataSource                  = &awsBackupSelectionDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsBackupSelectionDataSource{}
)

func NewAwsBackupSelectionDataSource() datasource.DataSource {
	newDatasource := &awsBackupSelectionDataSource{}

	newDatasource.Tagging = tagging.GetAwsBackupSelection()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
