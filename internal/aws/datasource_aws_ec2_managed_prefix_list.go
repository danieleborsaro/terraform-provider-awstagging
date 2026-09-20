package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEc2ManagedPrefixListDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEc2ManagedPrefixListDataSource{}
	_ datasource.DataSource                  = &awsEc2ManagedPrefixListDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEc2ManagedPrefixListDataSource{}
)

func NewAwsEc2ManagedPrefixListDataSource() datasource.DataSource {
	newDatasource := &awsEc2ManagedPrefixListDataSource{}

	newDatasource.Tagging = tagging.GetAwsEc2ManagedPrefixList()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
