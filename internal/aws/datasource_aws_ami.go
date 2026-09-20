package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsAmiDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsAmiDataSource{}
	_ datasource.DataSource                  = &awsAmiDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsAmiDataSource{}
)

func NewAwsAmiDataSource() datasource.DataSource {
	newDatasource := &awsAmiDataSource{}

	newDatasource.Tagging = tagging.GetAwsAmi()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
