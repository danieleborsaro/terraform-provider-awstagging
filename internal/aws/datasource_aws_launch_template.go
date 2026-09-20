package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsLaunchTemplateDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsLaunchTemplateDataSource{}
	_ datasource.DataSource                  = &awsLaunchTemplateDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsLaunchTemplateDataSource{}
)

func NewAwsLaunchTemplateDataSource() datasource.DataSource {
	newDatasource := &awsLaunchTemplateDataSource{}

	newDatasource.Tagging = tagging.GetAwsLaunchTemplate()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
