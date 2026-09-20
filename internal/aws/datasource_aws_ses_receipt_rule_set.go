package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsSesReceiptRuleSetDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsSesReceiptRuleSetDataSource{}
	_ datasource.DataSource                  = &awsSesReceiptRuleSetDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsSesReceiptRuleSetDataSource{}
)

func NewAwsSesReceiptRuleSetDataSource() datasource.DataSource {
	newDatasource := &awsSesReceiptRuleSetDataSource{}

	newDatasource.Tagging = tagging.GetAwsSesReceiptRuleSet()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
