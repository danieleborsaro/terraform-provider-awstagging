package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEcsServiceDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEcsServiceDataSource{}
	_ datasource.DataSource                  = &awsEcsServiceDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEcsServiceDataSource{}
)

func NewAwsEcsServiceDataSource() datasource.DataSource {
	newDatasource := &awsEcsServiceDataSource{}

	newDatasource.Tagging = tagging.GetAwsEcsService()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
