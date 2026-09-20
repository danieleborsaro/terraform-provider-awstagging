package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsS3ObjectDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsS3ObjectDataSource{}
	_ datasource.DataSource                  = &awsS3ObjectDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsS3ObjectDataSource{}
)

func NewAwsS3ObjectDataSource() datasource.DataSource {
	newDatasource := &awsS3ObjectDataSource{}

	newDatasource.Tagging = tagging.GetAwsS3Object()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
