package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsRoute53ResolverEndpointDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsRoute53ResolverEndpointDataSource{}
	_ datasource.DataSource                  = &awsRoute53ResolverEndpointDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsRoute53ResolverEndpointDataSource{}
)

func NewAwsRoute53ResolverEndpointDataSource() datasource.DataSource {
	newDatasource := &awsRoute53ResolverEndpointDataSource{}

	newDatasource.Tagging = tagging.GetAwsRoute53ResolverEndpoint()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
