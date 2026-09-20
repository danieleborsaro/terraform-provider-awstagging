package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsSsmParameterDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsSsmParameterDataSource{}
	_ datasource.DataSource                  = &awsSsmParameterDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsSsmParameterDataSource{}
)

func NewAwsSsmParameterDataSource() datasource.DataSource {
	newDatasource := &awsSsmParameterDataSource{}

	newDatasource.Tagging = tagging.GetAwsSsmParameter()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
