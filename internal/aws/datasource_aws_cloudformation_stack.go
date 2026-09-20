package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsCloudformationStackDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsCloudformationStackDataSource{}
	_ datasource.DataSource                  = &awsCloudformationStackDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsCloudformationStackDataSource{}
)

func NewAwsCloudformationStackDataSource() datasource.DataSource {
	newDatasource := &awsCloudformationStackDataSource{}

	newDatasource.Tagging = tagging.GetAwsCloudformationStack()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
