package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsRdsReservedInstanceDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsRdsReservedInstanceDataSource{}
	_ datasource.DataSource                  = &awsRdsReservedInstanceDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsRdsReservedInstanceDataSource{}
)

func NewAwsRdsReservedInstanceDataSource() datasource.DataSource {
	newDatasource := &awsRdsReservedInstanceDataSource{}

	newDatasource.Tagging = tagging.GetAwsRdsReservedInstance()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
