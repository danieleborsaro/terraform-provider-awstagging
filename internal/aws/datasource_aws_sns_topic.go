package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsSnsTopicDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsSnsTopicDataSource{}
	_ datasource.DataSource                  = &awsSnsTopicDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsSnsTopicDataSource{}
)

func NewAwsSnsTopicDataSource() datasource.DataSource {
	newDatasource := &awsSnsTopicDataSource{}

	newDatasource.Tagging = tagging.GetAwsSnsTopic()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
