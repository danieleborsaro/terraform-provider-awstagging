package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafregionalRegexMatchSetDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafregionalRegexMatchSetDataSource{}
	_ datasource.DataSource                  = &awsWafregionalRegexMatchSetDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafregionalRegexMatchSetDataSource{}
)

func NewAwsWafregionalRegexMatchSetDataSource() datasource.DataSource {
	newDatasource := &awsWafregionalRegexMatchSetDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafregionalRegexMatchSet()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
