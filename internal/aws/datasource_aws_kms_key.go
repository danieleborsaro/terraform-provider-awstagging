package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsKmsKeyDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsKmsKeyDataSource{}
	_ datasource.DataSource                  = &awsKmsKeyDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsKmsKeyDataSource{}
)

func NewAwsKmsKeyDataSource() datasource.DataSource {
	newDatasource := &awsKmsKeyDataSource{}

	newDatasource.Tagging = tagging.GetAwsKmsKey()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
