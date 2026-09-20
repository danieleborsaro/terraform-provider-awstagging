package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsCloudwatchCompositeAlarmDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsCloudwatchCompositeAlarmDataSource{}
	_ datasource.DataSource                  = &awsCloudwatchCompositeAlarmDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsCloudwatchCompositeAlarmDataSource{}
)

func NewAwsCloudwatchCompositeAlarmDataSource() datasource.DataSource {
	newDatasource := &awsCloudwatchCompositeAlarmDataSource{}

	newDatasource.Tagging = tagging.GetAwsCloudwatchCompositeAlarm()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
