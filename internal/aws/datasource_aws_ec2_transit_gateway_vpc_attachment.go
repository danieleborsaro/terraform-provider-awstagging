package aws

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	tagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

type awsEc2TransitGatewayVpcAttachmentDataSource struct {
	AwsDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasources.TaggingDataSourceInterface = &awsEc2TransitGatewayVpcAttachmentDataSource{}
	_ datasource.DataSource                  = &awsEc2TransitGatewayVpcAttachmentDataSource{}
	_ datasource.DataSourceWithConfigure     = &awsEc2TransitGatewayVpcAttachmentDataSource{}
)

func NewAwsEc2TransitGatewayVpcAttachmentDataSource() datasource.DataSource {
	newDatasource := &awsEc2TransitGatewayVpcAttachmentDataSource{}

	newDatasource.Tagging = tagging.GetAwsEc2TransitGatewayVpcAttachment()
	newDatasource.DatasourceType = newDatasource.Tagging.GetProperties().Id.Long

	return newDatasource
}
