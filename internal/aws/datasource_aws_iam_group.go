package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsIamGroupDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsIamGroupDataSource{}
	_ datasource.DataSource                  = &awsIamGroupDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsIamGroupDataSource{}
)

func NewAwsIamGroupDataSource() datasource.DataSource {
	newDatasource := &awsIamGroupDataSource{}

	newDatasource.Tagging = tagging.GetAwsIamGroup()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
