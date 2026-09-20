package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsCloudwatchLogGroupDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsCloudwatchLogGroupDataSource{}
	_ datasource.DataSource                  = &awsCloudwatchLogGroupDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsCloudwatchLogGroupDataSource{}
)

func NewAwsCloudwatchLogGroupDataSource() datasource.DataSource {
	newDatasource := &awsCloudwatchLogGroupDataSource{}

	newDatasource.Tagging = tagging.GetAwsCloudwatchLogGroup()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
