package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafv2IpSetDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafv2IpSetDataSource{}
	_ datasource.DataSource                  = &awsWafv2IpSetDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafv2IpSetDataSource{}
)

func NewAwsWafv2IpSetDataSource() datasource.DataSource {
	newDatasource := &awsWafv2IpSetDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafv2IpSet()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
