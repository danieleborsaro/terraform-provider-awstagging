package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafRuleDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafRuleDataSource{}
	_ datasource.DataSource                  = &awsWafRuleDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafRuleDataSource{}
)

func NewAwsWafRuleDataSource() datasource.DataSource {
	newDatasource := &awsWafRuleDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafRule()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
