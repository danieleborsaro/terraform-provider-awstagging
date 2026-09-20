package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsOpensearchDomainDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsOpensearchDomainDataSource{}
	_ datasource.DataSource                  = &awsOpensearchDomainDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsOpensearchDomainDataSource{}
)

func NewAwsOpensearchDomainDataSource() datasource.DataSource {
	newDatasource := &awsOpensearchDomainDataSource{}

	newDatasource.Tagging = tagging.GetAwsOpensearchDomain()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
