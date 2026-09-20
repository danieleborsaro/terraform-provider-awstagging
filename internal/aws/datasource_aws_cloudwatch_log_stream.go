package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsCloudwatchLogStreamDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsCloudwatchLogStreamDataSource{}
	_ datasource.DataSource                  = &awsCloudwatchLogStreamDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsCloudwatchLogStreamDataSource{}
)

func NewAwsCloudwatchLogStreamDataSource() datasource.DataSource {
	newDatasource := &awsCloudwatchLogStreamDataSource{}

	newDatasource.Tagging = tagging.GetAwsCloudwatchLogStream()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
