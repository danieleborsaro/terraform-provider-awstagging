package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafv2RegexPatternSetDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafv2RegexPatternSetDataSource{}
	_ datasource.DataSource                  = &awsWafv2RegexPatternSetDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafv2RegexPatternSetDataSource{}
)

func NewAwsWafv2RegexPatternSetDataSource() datasource.DataSource {
	newDatasource := &awsWafv2RegexPatternSetDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafv2RegexPatternSet()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
