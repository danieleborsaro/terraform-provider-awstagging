package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsElasticacheSubnetGroupDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsElasticacheSubnetGroupDataSource{}
	_ datasource.DataSource                  = &awsElasticacheSubnetGroupDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsElasticacheSubnetGroupDataSource{}
)

func NewAwsElasticacheSubnetGroupDataSource() datasource.DataSource {
	newDatasource := &awsElasticacheSubnetGroupDataSource{}

	newDatasource.Tagging = tagging.GetAwsElasticacheSubnetGroup()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
