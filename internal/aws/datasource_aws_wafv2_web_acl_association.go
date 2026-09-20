package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsWafv2WebAclAssociationDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsWafv2WebAclAssociationDataSource{}
	_ datasource.DataSource                  = &awsWafv2WebAclAssociationDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsWafv2WebAclAssociationDataSource{}
)

func NewAwsWafv2WebAclAssociationDataSource() datasource.DataSource {
	newDatasource := &awsWafv2WebAclAssociationDataSource{}

	newDatasource.Tagging = tagging.GetAwsWafv2WebAclAssociation()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
