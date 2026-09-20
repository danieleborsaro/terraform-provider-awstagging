package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEksClusterDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEksClusterDataSource{}
	_ datasource.DataSource                  = &awsEksClusterDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEksClusterDataSource{}
)

func NewAwsEksClusterDataSource() datasource.DataSource {
	newDatasource := &awsEksClusterDataSource{}

	newDatasource.Tagging = tagging.GetAwsEksCluster()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
