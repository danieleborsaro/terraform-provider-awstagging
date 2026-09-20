package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafregionalRegexPatternSetDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafregionalRegexPatternSetDataSource{}
	_ datasource.DataSource                  = &awsWafregionalRegexPatternSetDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafregionalRegexPatternSetDataSource{}
)

func NewAwsWafregionalRegexPatternSetDataSource() datasource.DataSource {
	newDatasource := &awsWafregionalRegexPatternSetDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafregionalRegexPatternSet()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
