package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafv2WebAclLoggingConfigurationDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafv2WebAclLoggingConfigurationDataSource{}
	_ datasource.DataSource                  = &awsWafv2WebAclLoggingConfigurationDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafv2WebAclLoggingConfigurationDataSource{}
)

func NewAwsWafv2WebAclLoggingConfigurationDataSource() datasource.DataSource {
	newDatasource := &awsWafv2WebAclLoggingConfigurationDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafv2WebAclLoggingConfiguration()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
