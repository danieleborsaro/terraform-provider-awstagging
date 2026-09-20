package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsKeyPairDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsKeyPairDataSource{}
	_ datasource.DataSource                  = &awsKeyPairDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsKeyPairDataSource{}
)

func NewAwsKeyPairDataSource() datasource.DataSource {
	newDatasource := &awsKeyPairDataSource{}

	newDatasource.Tagging = tagging.GetAwsKeyPair()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
