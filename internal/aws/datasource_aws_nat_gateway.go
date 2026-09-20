package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsNatGatewayDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsNatGatewayDataSource{}
	_ datasource.DataSource                  = &awsNatGatewayDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsNatGatewayDataSource{}
)

func NewAwsNatGatewayDataSource() datasource.DataSource {
	newDatasource := &awsNatGatewayDataSource{}

	newDatasource.Tagging = tagging.GetAwsNatGateway()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
