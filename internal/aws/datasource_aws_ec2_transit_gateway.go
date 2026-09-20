package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEc2TransitGatewayDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEc2TransitGatewayDataSource{}
	_ datasource.DataSource                  = &awsEc2TransitGatewayDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEc2TransitGatewayDataSource{}
)

func NewAwsEc2TransitGatewayDataSource() datasource.DataSource {
	newDatasource := &awsEc2TransitGatewayDataSource{}

	newDatasource.Tagging = tagging.GetAwsEc2TransitGateway()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
