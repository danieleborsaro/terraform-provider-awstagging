package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsLbListenerRuleDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsLbListenerRuleDataSource{}
	_ datasource.DataSource                  = &awsLbListenerRuleDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsLbListenerRuleDataSource{}
)

func NewAwsLbListenerRuleDataSource() datasource.DataSource {
	newDatasource := &awsLbListenerRuleDataSource{}

	newDatasource.Tagging = tagging.GetAwsLbListenerRule()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
