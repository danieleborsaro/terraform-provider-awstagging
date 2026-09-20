package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsRoute53ZoneDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsRoute53ZoneDataSource{}
	_ datasource.DataSource                  = &awsRoute53ZoneDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsRoute53ZoneDataSource{}
)

func NewAwsRoute53ZoneDataSource() datasource.DataSource {
	newDatasource := &awsRoute53ZoneDataSource{}

	newDatasource.Tagging = tagging.GetAwsRoute53Zone()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
