package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsBackupPlanDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsBackupPlanDataSource{}
	_ datasource.DataSource                  = &awsBackupPlanDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsBackupPlanDataSource{}
)

func NewAwsBackupPlanDataSource() datasource.DataSource {
	newDatasource := &awsBackupPlanDataSource{}

	newDatasource.Tagging = tagging.GetAwsBackupPlan()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
