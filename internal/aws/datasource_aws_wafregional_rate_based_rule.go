package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafregionalRateBasedRuleDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafregionalRateBasedRuleDataSource{}
	_ datasource.DataSource                  = &awsWafregionalRateBasedRuleDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafregionalRateBasedRuleDataSource{}
)

func NewAwsWafregionalRateBasedRuleDataSource() datasource.DataSource {
	newDatasource := &awsWafregionalRateBasedRuleDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafregionalRateBasedRule()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
