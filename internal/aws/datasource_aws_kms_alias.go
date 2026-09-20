package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsKmsAliasDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsKmsAliasDataSource{}
	_ datasource.DataSource                  = &awsKmsAliasDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsKmsAliasDataSource{}
)

func NewAwsKmsAliasDataSource() datasource.DataSource {
	newDatasource := &awsKmsAliasDataSource{}

	newDatasource.Tagging = tagging.GetAwsKmsAlias()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
