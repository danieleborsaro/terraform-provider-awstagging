package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsDbParameterGroupDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsDbParameterGroupDataSource{}
	_ datasource.DataSource                  = &awsDbParameterGroupDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsDbParameterGroupDataSource{}
)

func NewAwsDbParameterGroupDataSource() datasource.DataSource {
	newDatasource := &awsDbParameterGroupDataSource{}

	newDatasource.Tagging = tagging.GetAwsDbParameterGroup()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
