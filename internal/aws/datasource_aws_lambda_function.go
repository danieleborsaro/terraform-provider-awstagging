package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsLambdaFunctionDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsLambdaFunctionDataSource{}
	_ datasource.DataSource                  = &awsLambdaFunctionDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsLambdaFunctionDataSource{}
)

func NewAwsLambdaFunctionDataSource() datasource.DataSource {
	newDatasource := &awsLambdaFunctionDataSource{}

	newDatasource.Tagging = tagging.GetAwsLambdaFunction()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
