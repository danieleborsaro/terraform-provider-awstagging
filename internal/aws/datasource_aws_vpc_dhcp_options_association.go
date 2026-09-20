package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsVpcDhcpOptionsAssociationDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsVpcDhcpOptionsAssociationDataSource{}
	_ datasource.DataSource                  = &awsVpcDhcpOptionsAssociationDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsVpcDhcpOptionsAssociationDataSource{}
)

func NewAwsVpcDhcpOptionsAssociationDataSource() datasource.DataSource {
	newDatasource := &awsVpcDhcpOptionsAssociationDataSource{}

	newDatasource.Tagging = tagging.GetAwsVpcDhcpOptionsAssociation()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
