package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsRdsGlobalClusterDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsRdsGlobalClusterDataSource{}
	_ datasource.DataSource                  = &awsRdsGlobalClusterDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsRdsGlobalClusterDataSource{}
)

func NewAwsRdsGlobalClusterDataSource() datasource.DataSource {
	newDatasource := &awsRdsGlobalClusterDataSource{}

	newDatasource.Tagging = tagging.GetAwsRdsGlobalCluster()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
