package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsDbOptionGroupDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsDbOptionGroupDataSource{}
	_ datasource.DataSource                  = &awsDbOptionGroupDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsDbOptionGroupDataSource{}
)

func NewAwsDbOptionGroupDataSource() datasource.DataSource {
	newDatasource := &awsDbOptionGroupDataSource{}

	newDatasource.Tagging = tagging.GetAwsDbOptionGroup()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
