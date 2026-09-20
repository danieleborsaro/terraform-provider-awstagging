package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsReservedInstanceDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsReservedInstanceDataSource{}
	_ datasource.DataSource                  = &awsReservedInstanceDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsReservedInstanceDataSource{}
)

func NewAwsReservedInstanceDataSource() datasource.DataSource {
	newDatasource := &awsReservedInstanceDataSource{}

	newDatasource.Tagging = tagging.GetAwsReservedInstance()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
