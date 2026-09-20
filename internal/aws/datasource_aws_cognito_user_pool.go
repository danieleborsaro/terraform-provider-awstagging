package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsCognitoUserPoolDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsCognitoUserPoolDataSource{}
	_ datasource.DataSource                  = &awsCognitoUserPoolDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsCognitoUserPoolDataSource{}
)

func NewAwsCognitoUserPoolDataSource() datasource.DataSource {
	newDatasource := &awsCognitoUserPoolDataSource{}

	newDatasource.Tagging = tagging.GetAwsCognitoUserPool()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
