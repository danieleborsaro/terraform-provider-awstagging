package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsInstanceDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsInstanceDataSource{}
	_ datasource.DataSource                  = &awsInstanceDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsInstanceDataSource{}
)

func NewAwsInstanceDataSource() datasource.DataSource {
	newDatasource := &awsInstanceDataSource{}

	newDatasource.Tagging = tagging.GetAwsInstance()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
