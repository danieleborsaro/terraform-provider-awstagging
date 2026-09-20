package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsRoute53ResolverRuleDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsRoute53ResolverRuleDataSource{}
	_ datasource.DataSource                  = &awsRoute53ResolverRuleDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsRoute53ResolverRuleDataSource{}
)

func NewAwsRoute53ResolverRuleDataSource() datasource.DataSource {
	newDatasource := &awsRoute53ResolverRuleDataSource{}

	newDatasource.Tagging = tagging.GetAwsRoute53ResolverRule()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
