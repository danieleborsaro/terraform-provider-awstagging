package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEbsVolumeDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEbsVolumeDataSource{}
	_ datasource.DataSource                  = &awsEbsVolumeDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEbsVolumeDataSource{}
)

func NewAwsEbsVolumeDataSource() datasource.DataSource {
	newDatasource := &awsEbsVolumeDataSource{}

	newDatasource.Tagging = tagging.GetAwsEbsVolume()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
