package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsVpcDhcpOptionsDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsVpcDhcpOptionsDataSource{}
	_ datasource.DataSource                  = &awsVpcDhcpOptionsDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsVpcDhcpOptionsDataSource{}
)

func NewAwsVpcDhcpOptionsDataSource() datasource.DataSource {
	newDatasource := &awsVpcDhcpOptionsDataSource{}

	newDatasource.Tagging = tagging.GetAwsVpcDhcpOptions()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
