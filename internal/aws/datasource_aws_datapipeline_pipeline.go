package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsDatapipelinePipelineDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsDatapipelinePipelineDataSource{}
	_ datasource.DataSource                  = &awsDatapipelinePipelineDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsDatapipelinePipelineDataSource{}
)

func NewAwsDatapipelinePipelineDataSource() datasource.DataSource {
	newDatasource := &awsDatapipelinePipelineDataSource{}

	newDatasource.Tagging = tagging.GetAwsDatapipelinePipeline()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
