package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafregionalGeoMatchSetDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafregionalGeoMatchSetDataSource{}
	_ datasource.DataSource                  = &awsWafregionalGeoMatchSetDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafregionalGeoMatchSetDataSource{}
)

func NewAwsWafregionalGeoMatchSetDataSource() datasource.DataSource {
	newDatasource := &awsWafregionalGeoMatchSetDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafregionalGeoMatchSet()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
