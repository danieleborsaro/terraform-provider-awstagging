package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEfsAccessPointDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEfsAccessPointDataSource{}
	_ datasource.DataSource                  = &awsEfsAccessPointDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEfsAccessPointDataSource{}
)

func NewAwsEfsAccessPointDataSource() datasource.DataSource {
	newDatasource := &awsEfsAccessPointDataSource{}

	newDatasource.Tagging = tagging.GetAwsEfsAccessPoint()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
