package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsIamPolicyDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsIamPolicyDataSource{}
	_ datasource.DataSource                  = &awsIamPolicyDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsIamPolicyDataSource{}
)

func NewAwsIamPolicyDataSource() datasource.DataSource {
	newDatasource := &awsIamPolicyDataSource{}

	newDatasource.Tagging = tagging.GetAwsIamPolicy()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
