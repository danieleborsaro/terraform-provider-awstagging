package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEcsCapacityProviderDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEcsCapacityProviderDataSource{}
	_ datasource.DataSource                  = &awsEcsCapacityProviderDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEcsCapacityProviderDataSource{}
)

func NewAwsEcsCapacityProviderDataSource() datasource.DataSource {
	newDatasource := &awsEcsCapacityProviderDataSource{}

	newDatasource.Tagging = tagging.GetAwsEcsCapacityProvider()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
