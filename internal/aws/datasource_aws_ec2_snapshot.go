package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEc2SnapshotDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEc2SnapshotDataSource{}
	_ datasource.DataSource                  = &awsEc2SnapshotDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEc2SnapshotDataSource{}
)

func NewAwsEc2SnapshotDataSource() datasource.DataSource {
	newDatasource := &awsEc2SnapshotDataSource{}

	newDatasource.Tagging = tagging.GetAwsEc2Snapshot()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
