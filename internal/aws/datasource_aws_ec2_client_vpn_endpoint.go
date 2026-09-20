package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEc2ClientVpnEndpointDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEc2ClientVpnEndpointDataSource{}
	_ datasource.DataSource                  = &awsEc2ClientVpnEndpointDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEc2ClientVpnEndpointDataSource{}
)

func NewAwsEc2ClientVpnEndpointDataSource() datasource.DataSource {
	newDatasource := &awsEc2ClientVpnEndpointDataSource{}

	newDatasource.Tagging = tagging.GetAwsEc2ClientVpnEndpoint()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
