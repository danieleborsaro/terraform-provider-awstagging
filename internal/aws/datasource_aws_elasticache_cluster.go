package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsElasticacheClusterDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsElasticacheClusterDataSource{}
	_ datasource.DataSource                  = &awsElasticacheClusterDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsElasticacheClusterDataSource{}
)

func NewAwsElasticacheClusterDataSource() datasource.DataSource {
	newDatasource := &awsElasticacheClusterDataSource{}

	newDatasource.Tagging = tagging.GetAwsElasticacheCluster()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
