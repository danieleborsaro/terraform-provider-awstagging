package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsSubnetDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsSubnetDataSource{}
	_ datasource.DataSource                  = &awsSubnetDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsSubnetDataSource{}
)

func NewAwsSubnetDataSource() datasource.DataSource {
	newDatasource := &awsSubnetDataSource{}

	newDatasource.Tagging = tagging.GetAwsSubnet()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
