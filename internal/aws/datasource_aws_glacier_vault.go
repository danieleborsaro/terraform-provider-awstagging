package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsGlacierVaultDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsGlacierVaultDataSource{}
	_ datasource.DataSource                  = &awsGlacierVaultDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsGlacierVaultDataSource{}
)

func NewAwsGlacierVaultDataSource() datasource.DataSource {
	newDatasource := &awsGlacierVaultDataSource{}

	newDatasource.Tagging = tagging.GetAwsGlacierVault()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
