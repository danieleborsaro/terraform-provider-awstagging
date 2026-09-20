package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafregionalWebAclDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafregionalWebAclDataSource{}
	_ datasource.DataSource                  = &awsWafregionalWebAclDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafregionalWebAclDataSource{}
)

func NewAwsWafregionalWebAclDataSource() datasource.DataSource {
	newDatasource := &awsWafregionalWebAclDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafregionalWebAcl()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
