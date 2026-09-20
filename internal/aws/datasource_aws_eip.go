package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEipDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEipDataSource{}
	_ datasource.DataSource                  = &awsEipDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEipDataSource{}
)

func NewAwsEipDataSource() datasource.DataSource {
	newDatasource := &awsEipDataSource{}

	newDatasource.Tagging = tagging.GetAwsEip()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
