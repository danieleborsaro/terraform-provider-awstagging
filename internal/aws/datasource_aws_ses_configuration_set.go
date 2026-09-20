package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsSesConfigurationSetDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsSesConfigurationSetDataSource{}
	_ datasource.DataSource                  = &awsSesConfigurationSetDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsSesConfigurationSetDataSource{}
)

func NewAwsSesConfigurationSetDataSource() datasource.DataSource {
	newDatasource := &awsSesConfigurationSetDataSource{}

	newDatasource.Tagging = tagging.GetAwsSesConfigurationSet()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
