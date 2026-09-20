package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsNetworkInterfaceDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsNetworkInterfaceDataSource{}
	_ datasource.DataSource                  = &awsNetworkInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsNetworkInterfaceDataSource{}
)

func NewAwsNetworkInterfaceDataSource() datasource.DataSource {
	newDatasource := &awsNetworkInterfaceDataSource{}

	newDatasource.Tagging = tagging.GetAwsNetworkInterface()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
