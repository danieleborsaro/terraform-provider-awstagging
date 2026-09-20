package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsSesReceiptRuleDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsSesReceiptRuleDataSource{}
	_ datasource.DataSource                  = &awsSesReceiptRuleDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsSesReceiptRuleDataSource{}
)

func NewAwsSesReceiptRuleDataSource() datasource.DataSource {
	newDatasource := &awsSesReceiptRuleDataSource{}

	newDatasource.Tagging = tagging.GetAwsSesReceiptRule()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
