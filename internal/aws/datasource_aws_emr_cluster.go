package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEmrClusterDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEmrClusterDataSource{}
	_ datasource.DataSource                  = &awsEmrClusterDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEmrClusterDataSource{}
)

func NewAwsEmrClusterDataSource() datasource.DataSource {
	newDatasource := &awsEmrClusterDataSource{}

	newDatasource.Tagging = tagging.GetAwsEmrCluster()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
