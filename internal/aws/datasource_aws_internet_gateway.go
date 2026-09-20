package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsInternetGatewayDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsInternetGatewayDataSource{}
	_ datasource.DataSource                  = &awsInternetGatewayDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsInternetGatewayDataSource{}
)

func NewAwsInternetGatewayDataSource() datasource.DataSource {
	newDatasource := &awsInternetGatewayDataSource{}

	newDatasource.Tagging = tagging.GetAwsInternetGateway()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
