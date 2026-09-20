package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEcsTaskDefinitionDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEcsTaskDefinitionDataSource{}
	_ datasource.DataSource                  = &awsEcsTaskDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEcsTaskDefinitionDataSource{}
)

func NewAwsEcsTaskDefinitionDataSource() datasource.DataSource {
	newDatasource := &awsEcsTaskDefinitionDataSource{}

	newDatasource.Tagging = tagging.GetAwsEcsTaskDefinition()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
