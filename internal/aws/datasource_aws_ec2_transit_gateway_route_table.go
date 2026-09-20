package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEc2TransitGatewayRouteTableDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEc2TransitGatewayRouteTableDataSource{}
	_ datasource.DataSource                  = &awsEc2TransitGatewayRouteTableDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEc2TransitGatewayRouteTableDataSource{}
)

func NewAwsEc2TransitGatewayRouteTableDataSource() datasource.DataSource {
	newDatasource := &awsEc2TransitGatewayRouteTableDataSource{}

	newDatasource.Tagging = tagging.GetAwsEc2TransitGatewayRouteTable()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
