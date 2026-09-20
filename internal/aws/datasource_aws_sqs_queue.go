package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsSqsQueueDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsSqsQueueDataSource{}
	_ datasource.DataSource                  = &awsSqsQueueDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsSqsQueueDataSource{}
)

func NewAwsSqsQueueDataSource() datasource.DataSource {
	newDatasource := &awsSqsQueueDataSource{}

	newDatasource.Tagging = tagging.GetAwsSqsQueue()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
