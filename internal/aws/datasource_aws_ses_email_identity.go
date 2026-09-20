package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsSesEmailIdentityDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsSesEmailIdentityDataSource{}
	_ datasource.DataSource                  = &awsSesEmailIdentityDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsSesEmailIdentityDataSource{}
)

func NewAwsSesEmailIdentityDataSource() datasource.DataSource {
	newDatasource := &awsSesEmailIdentityDataSource{}

	newDatasource.Tagging = tagging.GetAwsSesEmailIdentity()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
