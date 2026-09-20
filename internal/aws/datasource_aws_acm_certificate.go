package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsAcmCertificateDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsAcmCertificateDataSource{}
	_ datasource.DataSource                  = &awsAcmCertificateDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsAcmCertificateDataSource{}
)

func NewAwsAcmCertificateDataSource() datasource.DataSource {
	newDatasource := &awsAcmCertificateDataSource{}

	newDatasource.Tagging = tagging.GetAwsAcmCertificate()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
