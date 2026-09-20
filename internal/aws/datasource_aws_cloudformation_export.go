package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsCloudformationExportDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsCloudformationExportDataSource{}
	_ datasource.DataSource                  = &awsCloudformationExportDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsCloudformationExportDataSource{}
)

func NewAwsCloudformationExportDataSource() datasource.DataSource {
	newDatasource := &awsCloudformationExportDataSource{}

	newDatasource.Tagging = tagging.GetAwsCloudformationExport()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
