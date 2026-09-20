package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsSesv2ConfigurationSetEventDestinationDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsSesv2ConfigurationSetEventDestinationDataSource{}
	_ datasource.DataSource                  = &awsSesv2ConfigurationSetEventDestinationDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsSesv2ConfigurationSetEventDestinationDataSource{}
)

func NewAwsSesv2ConfigurationSetEventDestinationDataSource() datasource.DataSource {
	newDatasource := &awsSesv2ConfigurationSetEventDestinationDataSource{}

	newDatasource.Tagging = tagging.GetAwsSesv2ConfigurationSetEventDestination()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
