package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafv2WebAclDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafv2WebAclDataSource{}
	_ datasource.DataSource                  = &awsWafv2WebAclDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafv2WebAclDataSource{}
)

func NewAwsWafv2WebAclDataSource() datasource.DataSource {
	newDatasource := &awsWafv2WebAclDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafv2WebAcl()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
