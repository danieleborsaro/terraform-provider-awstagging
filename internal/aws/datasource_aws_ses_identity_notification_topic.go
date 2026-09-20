package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsSesIdentityNotificationTopicDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsSesIdentityNotificationTopicDataSource{}
	_ datasource.DataSource                  = &awsSesIdentityNotificationTopicDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsSesIdentityNotificationTopicDataSource{}
)

func NewAwsSesIdentityNotificationTopicDataSource() datasource.DataSource {
	newDatasource := &awsSesIdentityNotificationTopicDataSource{}

	newDatasource.Tagging = tagging.GetAwsSesIdentityNotificationTopic()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
