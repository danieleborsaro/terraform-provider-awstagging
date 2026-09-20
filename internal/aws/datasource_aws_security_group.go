package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsSecurityGroupDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsSecurityGroupDataSource{}
	_ datasource.DataSource                  = &awsSecurityGroupDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsSecurityGroupDataSource{}
)

func NewAwsSecurityGroupDataSource() datasource.DataSource {
	newDatasource := &awsSecurityGroupDataSource{}

	newDatasource.Tagging = tagging.GetAwsSecurityGroup()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
