package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsTransferServerDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsTransferServerDataSource{}
	_ datasource.DataSource                  = &awsTransferServerDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsTransferServerDataSource{}
)

func NewAwsTransferServerDataSource() datasource.DataSource {
	newDatasource := &awsTransferServerDataSource{}

	newDatasource.Tagging = tagging.GetAwsTransferServer()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
