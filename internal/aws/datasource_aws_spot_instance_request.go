package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsSpotInstanceRequestDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsSpotInstanceRequestDataSource{}
	_ datasource.DataSource                  = &awsSpotInstanceRequestDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsSpotInstanceRequestDataSource{}
)

func NewAwsSpotInstanceRequestDataSource() datasource.DataSource {
	newDatasource := &awsSpotInstanceRequestDataSource{}

	newDatasource.Tagging = tagging.GetAwsSpotInstanceRequest()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
