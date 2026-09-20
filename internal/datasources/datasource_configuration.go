package datasources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	// datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	// datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	// models "github.com/danieleborsaro/terraform-provider-awstagging/internal/shared"
	// taggingdata "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type configurationDataSource struct {
	TaggingDataSource
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ TaggingDataSourceInterface         = &configurationDataSource{}
	_ datasource.DataSource              = &configurationDataSource{}
	_ datasource.DataSourceWithConfigure = &configurationDataSource{}
)

func NewConfigurationDataSource() datasource.DataSource {
	newDatasource := &configurationDataSource{}

	newDatasource.DatasourceType = "Provider Configuration"
	// newDatasource.Tagging = tagging.GetAwsAutoscalingGroup()

	return newDatasource
}

func (d *configurationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Trace(ctx, "datasource-configuration"+d.DatasourceType+" - BEGIN configuring metadata")
	resp.TypeName = req.ProviderTypeName + "_configuration"
	tflog.Trace(ctx, "datasource-configuration"+d.DatasourceType+" - END configuring metadata")
}

// // Read refreshes the Terraform state with the latest data.
// / NB: Overriding this function in order to consume the overriding UpdateState function below
func (d *configurationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Trace(ctx, "datasource-configuration"+d.DatasourceType+" - BEGIN reading state")

	var datasourceConfiguration, state DataSourceModel

	// Read data source input parameters into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &datasourceConfiguration)...)

	if resp.Diagnostics.HasError() {
		return
	}

	d.MergeConfiguration(ctx, &datasourceConfiguration, req, resp)

	//// NB: Typically here the data source would query the remote client for a resource state and add it to state variable

	//// Preparing state

	state = datasourceConfiguration

	d.UpdateState(ctx, &state, &datasourceConfiguration, req, resp)

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "awstagging - datasource "+d.DatasourceType+" state updated")

	tflog.Trace(ctx, "datasource-configuration"+d.DatasourceType+" - END reading state")
}

// Read refreshes the Terraform state with the latest data.
func (d *configurationDataSource) UpdateState(ctx context.Context, state *DataSourceModel, datasourceConfiguration *DataSourceModel, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Trace(ctx, "datasource-configuration"+d.DatasourceType+" - BEGIN updating state")

	//  Pass provider configuration as is

	tflog.Trace(ctx, "datasource-configuration"+d.DatasourceType+" - END updating state")
}
