package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsDbSubnetGroupDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsDbSubnetGroupDataSource{}
	_ datasource.DataSource                  = &awsDbSubnetGroupDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsDbSubnetGroupDataSource{}
)

func NewAwsDbSubnetGroupDataSource() datasource.DataSource {
	newDatasource := &awsDbSubnetGroupDataSource{}

	newDatasource.Tagging = tagging.GetAwsDbSubnetGroup()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
