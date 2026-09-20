package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsDbSnapshotDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsDbSnapshotDataSource{}
	_ datasource.DataSource                  = &awsDbSnapshotDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsDbSnapshotDataSource{}
)

func NewAwsDbSnapshotDataSource() datasource.DataSource {
	newDatasource := &awsDbSnapshotDataSource{}

	newDatasource.Tagging = tagging.GetAwsDbSnapshot()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
