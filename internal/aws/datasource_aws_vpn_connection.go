package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsVpnConnectionDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsVpnConnectionDataSource{}
	_ datasource.DataSource                  = &awsVpnConnectionDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsVpnConnectionDataSource{}
)

func NewAwsVpnConnectionDataSource() datasource.DataSource {
	newDatasource := &awsVpnConnectionDataSource{}

	newDatasource.Tagging = tagging.GetAwsVpnConnection()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
