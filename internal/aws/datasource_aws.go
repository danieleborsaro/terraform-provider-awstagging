package aws

import (
	"context"

	datasources "github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type AwsDataSource struct {
	datasources.TaggingDataSource
}

// // Read refreshes the Terraform state with the latest data.
// / NB: Overriding this function in order to consume the overriding UpdateState function below
func (d *AwsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Trace(ctx, "datasource "+d.DatasourceType+" - BEGIN reading state")

	var datasourceConfiguration, state datasources.DataSourceModel

	// Read data source input parameters into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &datasourceConfiguration)...)

	if resp.Diagnostics.HasError() {
		return
	}

	//// Check per-datasource properties

	if d.Tagging.GetProperties().Id.Key != "AwsAutoscalingGroup" && d.Tagging.GetProperties().Id.Key != "AwsLaunchTemplate" && !datasourceConfiguration.IsPropagateTagsAtLaunch.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("is_propagate_tags_at_lauch"),
			"Unexpected aws-tagging property is_propagate_tags_at_lauch",
			"The datasource "+d.Tagging.GetProperties().Id.Key+" cannot read its configuration as there is an unexpected value for the aws-tagging is_propagate_tags_at_lauch. "+
				"Remove is_propagate_tags_at_lauch from the configuration. ",
		)
	}

	if d.Tagging.GetProperties().Id.Key != "AwsInstance" && !datasourceConfiguration.PlatformName.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("platform_name"),
			"Unexpected aws-tagging property platform_name",
			"The datasource "+d.Tagging.GetProperties().Id.Key+" cannot read its configuration as there is an unexpected value for the aws-tagging platform_name. "+
				"Remove platform_name from the configuration. ",
		)
	}

	if d.Tagging.GetProperties().Id.Key == "AwsLb" && datasourceConfiguration.IsApplicationLoadBalancer.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("is_application_load_balancer"),
			"Missing aws-tagging property is_application_load_balancer",
			"The datasource "+d.Tagging.GetProperties().Id.Key+" cannot read its configuration as there is a missing or empty value for the aws-tagging is_application_load_balancer. "+
				"Set the is_application_load_balancer value in the configuration. ",
		)

		// } else if d.Tagging.GetProperties().Id.Key != "AwsLb" {
		// 	datasourceConfiguration.IsApplicationLoadBalancer = types.BoolValue(false)
	}

	if d.Tagging.GetProperties().Id.Key == "AwsRoute53ResolverEndpoint" && datasourceConfiguration.IsInbound.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("is_inbound"),
			"Missing aws-tagging property is_inbound",
			"The datasource "+d.Tagging.GetProperties().Id.Key+" cannot read its configuration as there is a missing or empty value for the aws-tagging is_inbound. "+
				"Set the is_inbound value in the configuration. ",
		)
	}

	// resp.Diagnostics.AddAttributeWarning(
	// 	path.Root("platform_name"),
	// 	"DEBUG from " + d.Tagging.GetProperties().Id.Key,
	// 	datasourceConfiguration.PlatformName.ValueString(),
	// )

	//// NB: enable this to read datasource arguments, but it will return the full configuration with each datasource
	d.MergeConfiguration(ctx, &datasourceConfiguration, req, resp)

	//// NB: Typically here the data source would query the remote client for a resource state and add it to state variable

	//// Preparing state

	//// NB: Preloading staging state with datasource configuration to avoid having to copy it entirely
	////     in UpdateState: there, we will just update state with info coming from the taggign package
	// state = datasources.DataSourceModel{}
	state = datasourceConfiguration

	d.UpdateState(ctx, &state, &datasourceConfiguration, req, resp)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "awstagging - datasource "+d.DatasourceType+" state computed")

	tflog.Trace(ctx, "datasource "+d.DatasourceType+" - END reading state")
}
