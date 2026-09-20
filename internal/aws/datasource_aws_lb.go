package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsLbDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsLbDataSource{}
	_ datasource.DataSource                  = &awsLbDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsLbDataSource{}
)

func NewAwsLbDataSource() datasource.DataSource {
	newDatasource := &awsLbDataSource{}

	newDatasource.Tagging = tagging.GetAwsLb()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
