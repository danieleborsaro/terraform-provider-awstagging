package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsAutoscalingGroupDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsAutoscalingGroupDataSource{}
	_ datasource.DataSource                  = &awsAutoscalingGroupDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsAutoscalingGroupDataSource{}
)

func NewAwsAutoscalingGroupDataSource() datasource.DataSource {
	newDatasource := &awsAutoscalingGroupDataSource{}

	newDatasource.Tagging = tagging.GetAwsAutoscalingGroup()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
