package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsElbDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsElbDataSource{}
	_ datasource.DataSource                  = &awsElbDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsElbDataSource{}
)

func NewAwsElbDataSource() datasource.DataSource {
	newDatasource := &awsElbDataSource{}

	newDatasource.Tagging = tagging.GetAwsElb()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
