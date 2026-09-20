package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsCognitoIdentityPoolDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsCognitoIdentityPoolDataSource{}
	_ datasource.DataSource                  = &awsCognitoIdentityPoolDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsCognitoIdentityPoolDataSource{}
)

func NewAwsCognitoIdentityPoolDataSource() datasource.DataSource {
	newDatasource := &awsCognitoIdentityPoolDataSource{}

	newDatasource.Tagging = tagging.GetAwsCognitoIdentityPool()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
