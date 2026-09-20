package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafregionalByteMatchSetDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafregionalByteMatchSetDataSource{}
	_ datasource.DataSource                  = &awsWafregionalByteMatchSetDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafregionalByteMatchSetDataSource{}
)

func NewAwsWafregionalByteMatchSetDataSource() datasource.DataSource {
	newDatasource := &awsWafregionalByteMatchSetDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafregionalByteMatchSet()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
