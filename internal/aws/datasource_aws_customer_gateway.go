package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsCustomerGatewayDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsCustomerGatewayDataSource{}
	_ datasource.DataSource                  = &awsCustomerGatewayDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsCustomerGatewayDataSource{}
)

func NewAwsCustomerGatewayDataSource() datasource.DataSource {
	newDatasource := &awsCustomerGatewayDataSource{}

	newDatasource.Tagging = tagging.GetAwsCustomerGateway()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
