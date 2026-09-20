package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsCloudwatchMetricAlarmDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsCloudwatchMetricAlarmDataSource{}
	_ datasource.DataSource                  = &awsCloudwatchMetricAlarmDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsCloudwatchMetricAlarmDataSource{}
)

func NewAwsCloudwatchMetricAlarmDataSource() datasource.DataSource {
	newDatasource := &awsCloudwatchMetricAlarmDataSource{}

	newDatasource.Tagging = tagging.GetAwsCloudwatchMetricAlarm()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
