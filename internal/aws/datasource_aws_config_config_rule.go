package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsConfigConfigRuleDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsConfigConfigRuleDataSource{}
	_ datasource.DataSource                  = &awsConfigConfigRuleDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsConfigConfigRuleDataSource{}
)

func NewAwsConfigConfigRuleDataSource() datasource.DataSource {
	newDatasource := &awsConfigConfigRuleDataSource{}

	newDatasource.Tagging = tagging.GetAwsConfigConfigRule()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
