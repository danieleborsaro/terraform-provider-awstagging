package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEcsClusterDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEcsClusterDataSource{}
	_ datasource.DataSource                  = &awsEcsClusterDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEcsClusterDataSource{}
)

func NewAwsEcsClusterDataSource() datasource.DataSource {
	newDatasource := &awsEcsClusterDataSource{}

	newDatasource.Tagging = tagging.GetAwsEcsCluster()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
