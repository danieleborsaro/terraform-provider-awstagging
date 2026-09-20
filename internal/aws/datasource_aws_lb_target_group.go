package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsLbTargetGroupDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsLbTargetGroupDataSource{}
	_ datasource.DataSource                  = &awsLbTargetGroupDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsLbTargetGroupDataSource{}
)

func NewAwsLbTargetGroupDataSource() datasource.DataSource {
	newDatasource := &awsLbTargetGroupDataSource{}

	newDatasource.Tagging = tagging.GetAwsLbTargetGroup()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
