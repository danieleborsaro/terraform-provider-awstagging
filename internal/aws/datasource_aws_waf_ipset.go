package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafIpsetDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafIpsetDataSource{}
	_ datasource.DataSource                  = &awsWafIpsetDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafIpsetDataSource{}
)

func NewAwsWafIpsetDataSource() datasource.DataSource {
	newDatasource := &awsWafIpsetDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafIpset()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
