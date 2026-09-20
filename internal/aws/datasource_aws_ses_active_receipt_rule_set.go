package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsSesActiveReceiptRuleSetDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsSesActiveReceiptRuleSetDataSource{}
	_ datasource.DataSource                  = &awsSesActiveReceiptRuleSetDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsSesActiveReceiptRuleSetDataSource{}
)

func NewAwsSesActiveReceiptRuleSetDataSource() datasource.DataSource {
	newDatasource := &awsSesActiveReceiptRuleSetDataSource{}

	newDatasource.Tagging = tagging.GetAwsSesActiveReceiptRuleSet()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
