package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsIamRoleDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsIamRoleDataSource{}
	_ datasource.DataSource                  = &awsIamRoleDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsIamRoleDataSource{}
)

func NewAwsIamRoleDataSource() datasource.DataSource {
	newDatasource := &awsIamRoleDataSource{}

	newDatasource.Tagging = tagging.GetAwsIamRole()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
