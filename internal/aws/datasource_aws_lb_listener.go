package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsLbListenerDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsLbListenerDataSource{}
	_ datasource.DataSource                  = &awsLbListenerDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsLbListenerDataSource{}
)

func NewAwsLbListenerDataSource() datasource.DataSource {
	newDatasource := &awsLbListenerDataSource{}

	newDatasource.Tagging = tagging.GetAwsLbListener()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
