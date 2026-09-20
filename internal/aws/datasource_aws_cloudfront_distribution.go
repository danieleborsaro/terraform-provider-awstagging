package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsCloudfrontDistributionDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsCloudfrontDistributionDataSource{}
	_ datasource.DataSource                  = &awsCloudfrontDistributionDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsCloudfrontDistributionDataSource{}
)

func NewAwsCloudfrontDistributionDataSource() datasource.DataSource {
	newDatasource := &awsCloudfrontDistributionDataSource{}

	newDatasource.Tagging = tagging.GetAwsCloudfrontDistribution()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
