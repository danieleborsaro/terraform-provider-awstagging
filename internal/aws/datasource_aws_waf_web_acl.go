package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafWebAclDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafWebAclDataSource{}
	_ datasource.DataSource                  = &awsWafWebAclDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafWebAclDataSource{}
)

func NewAwsWafWebAclDataSource() datasource.DataSource {
	newDatasource := &awsWafWebAclDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafWebAcl()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
