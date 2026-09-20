package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsDynamodbTableDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsDynamodbTableDataSource{}
	_ datasource.DataSource                  = &awsDynamodbTableDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsDynamodbTableDataSource{}
)

func NewAwsDynamodbTableDataSource() datasource.DataSource {
	newDatasource := &awsDynamodbTableDataSource{}

	newDatasource.Tagging = tagging.GetAwsDynamodbTable()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
