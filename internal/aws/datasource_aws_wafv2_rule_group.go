package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafv2RuleGroupDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafv2RuleGroupDataSource{}
	_ datasource.DataSource                  = &awsWafv2RuleGroupDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafv2RuleGroupDataSource{}
)

func NewAwsWafv2RuleGroupDataSource() datasource.DataSource {
	newDatasource := &awsWafv2RuleGroupDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafv2RuleGroup()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
