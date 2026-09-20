package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsLaunchConfigurationDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsLaunchConfigurationDataSource{}
	_ datasource.DataSource                  = &awsLaunchConfigurationDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsLaunchConfigurationDataSource{}
)

func NewAwsLaunchConfigurationDataSource() datasource.DataSource {
	newDatasource := &awsLaunchConfigurationDataSource{}

	newDatasource.Tagging = tagging.GetAwsLaunchConfiguration()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
