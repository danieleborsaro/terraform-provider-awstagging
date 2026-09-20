package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsRoute53HealthCheckDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsRoute53HealthCheckDataSource{}
	_ datasource.DataSource                  = &awsRoute53HealthCheckDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsRoute53HealthCheckDataSource{}
)

func NewAwsRoute53HealthCheckDataSource() datasource.DataSource {
	newDatasource := &awsRoute53HealthCheckDataSource{}

	newDatasource.Tagging = tagging.GetAwsRoute53HealthCheck()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
