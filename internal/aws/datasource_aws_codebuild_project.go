package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsCodebuildProjectDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsCodebuildProjectDataSource{}
	_ datasource.DataSource                  = &awsCodebuildProjectDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsCodebuildProjectDataSource{}
)

func NewAwsCodebuildProjectDataSource() datasource.DataSource {
	newDatasource := &awsCodebuildProjectDataSource{}

	newDatasource.Tagging = tagging.GetAwsCodebuildProject()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
