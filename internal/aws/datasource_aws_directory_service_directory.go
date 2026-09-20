package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsDirectoryServiceDirectoryDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsDirectoryServiceDirectoryDataSource{}
	_ datasource.DataSource                  = &awsDirectoryServiceDirectoryDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsDirectoryServiceDirectoryDataSource{}
)

func NewAwsDirectoryServiceDirectoryDataSource() datasource.DataSource {
	newDatasource := &awsDirectoryServiceDirectoryDataSource{}

	newDatasource.Tagging = tagging.GetAwsDirectoryServiceDirectory()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
