package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsRedshiftClusterDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsRedshiftClusterDataSource{}
	_ datasource.DataSource                  = &awsRedshiftClusterDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsRedshiftClusterDataSource{}
)

func NewAwsRedshiftClusterDataSource() datasource.DataSource {
	newDatasource := &awsRedshiftClusterDataSource{}

	newDatasource.Tagging = tagging.GetAwsRedshiftCluster()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
