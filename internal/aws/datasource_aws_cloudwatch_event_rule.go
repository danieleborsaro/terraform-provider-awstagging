package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsCloudwatchEventRuleDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsCloudwatchEventRuleDataSource{}
	_ datasource.DataSource                  = &awsCloudwatchEventRuleDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsCloudwatchEventRuleDataSource{}
)

func NewAwsCloudwatchEventRuleDataSource() datasource.DataSource {
	newDatasource := &awsCloudwatchEventRuleDataSource{}

	newDatasource.Tagging = tagging.GetAwsCloudwatchEventRule()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
