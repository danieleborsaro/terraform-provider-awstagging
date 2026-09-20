package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsS3BucketDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsS3BucketDataSource{}
	_ datasource.DataSource                  = &awsS3BucketDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsS3BucketDataSource{}
)

func NewAwsS3BucketDataSource() datasource.DataSource {
	newDatasource := &awsS3BucketDataSource{}

	newDatasource.Tagging = tagging.GetAwsS3Bucket()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
