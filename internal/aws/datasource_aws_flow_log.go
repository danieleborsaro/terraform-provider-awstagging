package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsFlowLogDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsFlowLogDataSource{}
	_ datasource.DataSource                  = &awsFlowLogDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsFlowLogDataSource{}
)

func NewAwsFlowLogDataSource() datasource.DataSource {
	newDatasource := &awsFlowLogDataSource{}

	newDatasource.Tagging = tagging.GetAwsFlowLog()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
