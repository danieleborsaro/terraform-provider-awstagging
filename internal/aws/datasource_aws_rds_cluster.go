package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsRdsClusterDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsRdsClusterDataSource{}
	_ datasource.DataSource                  = &awsRdsClusterDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsRdsClusterDataSource{}
)

func NewAwsRdsClusterDataSource() datasource.DataSource {
	newDatasource := &awsRdsClusterDataSource{}

	newDatasource.Tagging = tagging.GetAwsRdsCluster()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
