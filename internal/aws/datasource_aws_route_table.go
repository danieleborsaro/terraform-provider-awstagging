package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsRouteTableDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsRouteTableDataSource{}
	_ datasource.DataSource                  = &awsRouteTableDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsRouteTableDataSource{}
)

func NewAwsRouteTableDataSource() datasource.DataSource {
	newDatasource := &awsRouteTableDataSource{}

	newDatasource.Tagging = tagging.GetAwsRouteTable()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
