package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	models "github.com/danieleborsaro/terraform-provider-awstagging/internal/shared"
	taggingcore "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/core"
	taggingdata "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

type TaggingDataSourceInterface interface {
	Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse)
}

type TaggingDataSource struct {
	TaggingDataSourceInterface

	DatasourceType        string
	providerConfiguration *models.ProviderConfiguration
	Tagging               taggingcore.ResourceInterface
}

func (d *TaggingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Trace(ctx, "datasource "+d.DatasourceType+" - BEGIN configuring metadata")
	resp.TypeName = req.ProviderTypeName + "_" + d.Tagging.GetProperties().Terraform.ResourceName
	tflog.Trace(ctx, "datasource "+d.DatasourceType+" - END configuring metadata")
}

func (d *TaggingDataSource) Schema(ctx context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Trace(ctx, "datasource "+d.DatasourceType+" - BEGIN configuring schema")

	// Schema must be instantiated on child datasource struct to avoid race conditions
	// even when using it a global variable
	tempSchema := datasourceschema.Schema{
		Attributes: map[string]datasourceschema.Attribute{

			"availability_zone": datasourceschema.StringAttribute{
				Description: "AWS availability zone where the resource is being provisioned in",
				Optional:    true,
				Required:    false,
			},

			"compute_type": datasourceschema.StringAttribute{
				Description: "Compute type for compute resources, e.g. EC2, RDS, ECS, Lambda, ...",
				Optional:    true,
				Required:    false,
			},

			"custom_fqdn": datasourceschema.StringAttribute{
				Description: "Custom FQDN for the resource, e.g. a Route53 alias",
				Optional:    true,
				Required:    false,
			},

			"custom_name": datasourceschema.StringAttribute{
				Description: "Custom name, when relying on automatically generated ones is not feasible",
				Optional:    true,
				Required:    false,
			},

			"custom_name_prefix": datasourceschema.StringAttribute{
				Description: "Custom name prefix, when relying on automatically generated ones is not feasible",
				Optional:    true,
				Required:    false,
			},

			"is_application_load_balancer": datasourceschema.BoolAttribute{
				Description: "Is this an application or network LoadBalancerV2?",
				Optional:    true,
				Required:    false,
			},

			"is_inbound": datasourceschema.BoolAttribute{
				Description: "Is this an inbound or outbound Route53 resolver endpoint?",
				Optional:    true,
				Required:    false,
			},

			"is_propagate_tags_at_lauch": datasourceschema.BoolAttribute{
				Description: "Propagate tags at launch? This is for building the AWS-formatted map of tags with the 'propagate_at_launch' property (relevant only for autoscaling groups?)",
				Optional:    true,
				Required:    false,
			},

			"platform_name": datasourceschema.StringAttribute{
				Description: "Runtime ID a resource is running, e.g. Linux, Windows, Python, Nodejs... Useful for naming AGS, Lambda functions, Launch Templates, ...",
				Optional:    true,
				Required:    false,
			},

			"role_extend": datasourceschema.StringAttribute{
				Description: "Role played by AWS resource in the context of its environment, e.g. database, security, ...",
				Optional:    true,
				Required:    false,
			},

			"versioning_sources": datasourceschema.ListAttribute{
				Description: "Values used to generate the random suffix. If not specified, the static portion of the resource name will be used",
				Optional:    true,
				Required:    false,
				ElementType: types.StringType,
			},

			"name": datasourceschema.StringAttribute{
				Computed: true,
			},
			"name_hash": datasourceschema.StringAttribute{
				Computed: true,
			},
			"name_prefix": datasourceschema.StringAttribute{
				Computed: true,
			},
			"tags": datasourceschema.MapAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"tags_as_maps": datasourceschema.ListNestedAttribute{
				Computed: true,
				NestedObject: datasourceschema.NestedAttributeObject{
					Attributes: map[string]datasourceschema.Attribute{
						"key": datasourceschema.StringAttribute{
							Computed: true,
						},
						"value": datasourceschema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"type_name": datasourceschema.StringAttribute{
				Computed: true,
			},
			"type_prefix": datasourceschema.StringAttribute{
				Computed: true,
			},
		},
	}

	// Any override option must be defined in this schema.
	// Not sure whether this is the best way to share the provider base schema.
	// Not even sure this is how it's supposed to work, but if I don't merge the two schemas
	// then I cannot specify parameters on datasources
	// for k, v := range models.ProviderSchema.Attributes {
	for k, v := range models.OverridesSchema.Attributes {
		tempSchema.Attributes[k] = v
	}

	resp.Schema = tempSchema

	tflog.Trace(ctx, "datasource "+d.DatasourceType+" - END configuring schema")
}

// Configure adds the provider configuration to the data source.
func (d *TaggingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Trace(ctx, "datasource "+d.DatasourceType+" - BEGIN retrieving configuration")

	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	providerConfiguration, ok := req.ProviderData.(*models.ProviderConfiguration)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *models.ProviderConfiguration, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.providerConfiguration = providerConfiguration

	tflog.Info(ctx, "awstagging - datasource "+d.DatasourceType+" configured")

	tflog.Trace(ctx, "datasource "+d.DatasourceType+" - END retrieving configuration")
}

// // DataSourceModel maps the data source schema data.
// // This datasource must be the join of datasource specific fields and provider fields
type DataSourceModel struct {
	//// Common parameters
	//// NB: the following unfortunately won't work because tags from ProviderModel are not automatically expanded
	// models.ProviderModel
	Compliance            types.String            `tfsdk:"compliance"`
	CostCentre            types.String            `tfsdk:"cost_centre"`
	CustomTags            map[string]types.String `tfsdk:"custom_tags"`
	CustomTagsVerbatim    map[string]types.String `tfsdk:"custom_tags_verbatim"`
	Description           types.String            `tfsdk:"description"`
	IsCreateBeforeDestroy types.Bool              `tfsdk:"is_create_before_destroy"`
	IsForceGeneratedName  types.Bool              `tfsdk:"is_force_generated_name"`
	Owner                 types.String            `tfsdk:"owner"`
	Region                types.String            `tfsdk:"region"`
	Role                  types.String            `tfsdk:"role"`

	//// Datasource parameters
	AvailabilityZone          types.String   `tfsdk:"availability_zone"`
	ComputeType               types.String   `tfsdk:"compute_type"`
	CustomFQDN                types.String   `tfsdk:"custom_fqdn"`
	CustomName                types.String   `tfsdk:"custom_name"`
	CustomNamePrefix          types.String   `tfsdk:"custom_name_prefix"`
	IsApplicationLoadBalancer types.Bool     `tfsdk:"is_application_load_balancer"`
	IsInbound                 types.Bool     `tfsdk:"is_inbound"`
	IsPropagateTagsAtLaunch   types.Bool     `tfsdk:"is_propagate_tags_at_lauch"`
	PlatformName              types.String   `tfsdk:"platform_name"`
	RoleExtend                types.String   `tfsdk:"role_extend"`
	VersioningSources         []types.String `tfsdk:"versioning_sources"`

	//// Datasource output
	Name       types.String      `tfsdk:"name"`
	NameHash   types.String      `tfsdk:"name_hash"`
	NamePrefix types.String      `tfsdk:"name_prefix"`
	Tags       map[string]string `tfsdk:"tags"`
	TagsAsMap  []TagsAsMapModel  `tfsdk:"tags_as_maps"`
	TypeName   types.String      `tfsdk:"type_name"`
	TypePrefix types.String      `tfsdk:"type_prefix"`
}

// tagsModel maps coffees schema data.
type TagsAsMapModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func (d *TaggingDataSource) MergeConfiguration(ctx context.Context, datasourceConfiguration *DataSourceModel, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Trace(ctx, "datasource "+d.DatasourceType+" - BEGIN consolidating configuration")

	//// Optional overrides

	compliance := d.providerConfiguration.Compliance
	if !datasourceConfiguration.Compliance.IsNull() {
		compliance = datasourceConfiguration.Compliance.ValueString()
	}
	datasourceConfiguration.Compliance = types.StringValue(compliance)

	costCentre := d.providerConfiguration.CostCentre
	if !datasourceConfiguration.CostCentre.IsNull() {
		costCentre = datasourceConfiguration.CostCentre.ValueString()
	}
	datasourceConfiguration.CostCentre = types.StringValue(costCentre)

	customTags := d.providerConfiguration.CustomTags
	if len(datasourceConfiguration.CustomTags) > 0 {
		customTags = make(map[string]string)
		for k, v := range datasourceConfiguration.CustomTags {
			customTags[k] = v.ValueString()
		}
	}
	datasourceConfiguration.CustomTags = make(map[string]types.String)
	for k, v := range customTags {
		datasourceConfiguration.CustomTags[k] = types.StringValue(v)
	}

	customTagsVerbatim := d.providerConfiguration.CustomTagsVerbatim
	if len(datasourceConfiguration.CustomTagsVerbatim) > 0 {
		customTagsVerbatim = make(map[string]string)
		for k, v := range datasourceConfiguration.CustomTagsVerbatim {
			customTagsVerbatim[k] = v.ValueString()
		}
	}
	datasourceConfiguration.CustomTagsVerbatim = make(map[string]types.String)
	for k, v := range customTagsVerbatim {
		datasourceConfiguration.CustomTagsVerbatim[k] = types.StringValue(v)
	}

	description := d.providerConfiguration.Description
	if !datasourceConfiguration.Description.IsNull() {
		description = datasourceConfiguration.Description.ValueString()
	}
	datasourceConfiguration.Description = types.StringValue(description)

	isCreateBeforeDestroy := d.providerConfiguration.IsCreateBeforeDestroy
	if !datasourceConfiguration.IsCreateBeforeDestroy.IsNull() {
		isCreateBeforeDestroy = datasourceConfiguration.IsCreateBeforeDestroy.ValueBool()
	}
	datasourceConfiguration.IsCreateBeforeDestroy = types.BoolValue(isCreateBeforeDestroy)

	isForceGeneratedName := d.providerConfiguration.IsForceGeneratedName
	if !datasourceConfiguration.IsForceGeneratedName.IsNull() {
		isForceGeneratedName = datasourceConfiguration.IsForceGeneratedName.ValueBool()
	}
	datasourceConfiguration.IsForceGeneratedName = types.BoolValue(isForceGeneratedName)

	owner := d.providerConfiguration.Owner
	if !datasourceConfiguration.Owner.IsNull() {
		owner = datasourceConfiguration.Owner.ValueString()
	}
	datasourceConfiguration.Owner = types.StringValue(owner)

	region := d.providerConfiguration.Region
	if !datasourceConfiguration.Region.IsNull() {
		region = datasourceConfiguration.Region.ValueString()
	}
	datasourceConfiguration.Region = types.StringValue(region)

	role := d.providerConfiguration.Role
	if !datasourceConfiguration.Role.IsNull() {
		role = datasourceConfiguration.Role.ValueString()
	}
	datasourceConfiguration.Role = types.StringValue(role)

	tflog.Debug(ctx, "awstagging - datasource "+d.DatasourceType+" configuration merged")
	tflog.Trace(ctx, "datasource "+d.DatasourceType+" - END consolidating configuration")
}

// Read refreshes the Terraform state with the latest data.
func (d *TaggingDataSource) UpdateState(ctx context.Context, state *DataSourceModel, datasourceConfiguration *DataSourceModel, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Trace(ctx, "datasource "+d.DatasourceType+" - BEGIN updating state")

	tagConfig := taggingdata.InputConfiguration{
		Account:                   d.providerConfiguration.Account,
		AccountsCoding:            make(map[string]taggingdata.AccountCodingConfiguration),
		AppEcosystem:              d.providerConfiguration.AppEcosystem,
		AppEnvironment:            d.providerConfiguration.AppEnvironment,
		AvailabilityZone:          datasourceConfiguration.AvailabilityZone.ValueString(),
		CompanyNameLong:           d.providerConfiguration.CompanyNameLong,
		CompanyNameShort:          d.providerConfiguration.CompanyNameShort,
		Compliance:                datasourceConfiguration.Compliance.ValueString(),
		ComputeType:               datasourceConfiguration.ComputeType.ValueString(),
		CostCentre:                datasourceConfiguration.CostCentre.ValueString(),
		CustomFQDN:                datasourceConfiguration.CustomFQDN.ValueString(),
		CustomName:                datasourceConfiguration.CustomName.ValueString(),
		CustomNamePrefix:          datasourceConfiguration.CustomNamePrefix.ValueString(),
		CustomTags:                map[string]string{},
		CustomTagsVerbatim:        map[string]string{},
		Description:               datasourceConfiguration.Description.ValueString(),
		InfraEnvironment:          d.providerConfiguration.InfraEnvironment,
		IsApplicationLoadBalancer: datasourceConfiguration.IsApplicationLoadBalancer.ValueBool(),
		IsCreateBeforeDestroy:     datasourceConfiguration.IsCreateBeforeDestroy.ValueBool(),
		IsForceGeneratedName:      datasourceConfiguration.IsForceGeneratedName.ValueBool(),
		IsInbound:                 datasourceConfiguration.IsInbound.ValueBool(),
		IsPropagateTagsAtLaunch:   datasourceConfiguration.IsPropagateTagsAtLaunch.ValueBool(),
		Owner:                     datasourceConfiguration.Owner.ValueString(),
		PlatformName:              datasourceConfiguration.PlatformName.ValueString(),
		ProjectNameLong:           d.providerConfiguration.ProjectNameLong,
		ProjectNameShort:          d.providerConfiguration.ProjectNameShort,
		Region:                    datasourceConfiguration.Region.ValueString(),
		ResourceSetLong:           d.providerConfiguration.ResourceSetLong,
		ResourceSetShort:          d.providerConfiguration.ResourceSetShort,
		Role:                      datasourceConfiguration.Role.ValueString(),
		RoleExtend:                datasourceConfiguration.RoleExtend.ValueString(),
		TerraformModule:           d.providerConfiguration.TerraformModule,
		TerraformWorkspace:        d.providerConfiguration.TerraformWorkspace,
		VersioningSources:         []string{},
	}

	for k, v := range d.providerConfiguration.AccountsCoding {
		tagConfig.AccountsCoding[k] = taggingdata.AccountCodingConfiguration{
			Class:         v.Class,
			Name:          v.Name,
			NameCanonical: v.NameCanonical,
			NameEncoded:   v.NameEncoded,
		}
	}

	for k, v := range datasourceConfiguration.CustomTags {
		tagConfig.CustomTags[k] = v.ValueString()
	}

	for k, v := range datasourceConfiguration.CustomTagsVerbatim {
		tagConfig.CustomTagsVerbatim[k] = v.ValueString()
	}

	for _, v := range datasourceConfiguration.VersioningSources {
		tagConfig.VersioningSources = append(tagConfig.VersioningSources, v.ValueString())
	}

	if err := d.Tagging.Generate(ctx, &tagConfig); err != nil {
		resp.Diagnostics.AddError(
			"Invalid tagging configuration",
			err.Error(),
		)
		return
	}

	state.Tags = d.Tagging.GetTags()

	for _, m := range d.Tagging.GetTagsAsMap() {
		state.TagsAsMap = append(state.TagsAsMap, TagsAsMapModel{
			Key:   types.StringValue(m["Key"]),
			Value: types.StringValue(m["Value"]),
		})
	}

	resName := d.Tagging.GetName().Safe
	resNameHash := d.Tagging.GetName().SafeSuffix
	resNamePrefix := d.Tagging.GetName().SafePrefix
	resResourceType := d.Tagging.GetProperties().Id.Long
	resTypePrefix := d.Tagging.GetName().TypePrefix

	state.Name = types.StringValue(resName)
	state.NameHash = types.StringValue(resNameHash)
	state.NamePrefix = types.StringValue(resNamePrefix)
	state.TypeName = types.StringValue(resResourceType)
	state.TypePrefix = types.StringValue(resTypePrefix)

	tflog.Debug(ctx, "awstagging - datasource "+d.DatasourceType+" state computed")

	tflog.Trace(ctx, "datasource "+d.DatasourceType+" - END updating state")
}
