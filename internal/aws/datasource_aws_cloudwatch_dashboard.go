package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsCloudwatchDashboardDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsCloudwatchDashboardDataSource{}
	_ datasource.DataSource                  = &awsCloudwatchDashboardDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsCloudwatchDashboardDataSource{}
)

func NewAwsCloudwatchDashboardDataSource() datasource.DataSource {
	newDatasource := &awsCloudwatchDashboardDataSource{}

	newDatasource.Tagging = tagging.GetAwsCloudwatchDashboard()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
