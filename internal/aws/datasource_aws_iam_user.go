package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsIamUserDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsIamUserDataSource{}
	_ datasource.DataSource                  = &awsIamUserDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsIamUserDataSource{}
)

func NewAwsIamUserDataSource() datasource.DataSource {
	newDatasource := &awsIamUserDataSource{}

	newDatasource.Tagging = tagging.GetAwsIamUser()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
