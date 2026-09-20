package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafRateBasedRuleDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafRateBasedRuleDataSource{}
	_ datasource.DataSource                  = &awsWafRateBasedRuleDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafRateBasedRuleDataSource{}
)

func NewAwsWafRateBasedRuleDataSource() datasource.DataSource {
	newDatasource := &awsWafRateBasedRuleDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafRateBasedRule()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
