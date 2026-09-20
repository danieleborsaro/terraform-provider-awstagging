package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsDbEventSubscriptionDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsDbEventSubscriptionDataSource{}
	_ datasource.DataSource                  = &awsDbEventSubscriptionDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsDbEventSubscriptionDataSource{}
)

func NewAwsDbEventSubscriptionDataSource() datasource.DataSource {
	newDatasource := &awsDbEventSubscriptionDataSource{}

	newDatasource.Tagging = tagging.GetAwsDbEventSubscription()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
