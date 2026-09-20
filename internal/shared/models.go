package shared

import (
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// // This is the 'object'
var OverridesSchema = providerschema.Schema{
	Attributes: map[string]providerschema.Attribute{

		"compliance": providerschema.StringAttribute{
			Description: "Compliance level for a resource",
			Optional:    true,
			Required:    false,
		},

		"cost_centre": providerschema.StringAttribute{
			Description: "Cost centre identifier for billing AWS resources",
			Optional:    true,
			Required:    false,
		},

		"custom_tags": datasourceschema.MapAttribute{
			Description: "Map of optional tags, in the form: 'TagName = TagValue'. Names are prefixed with custom prefix",
			Optional:    true,
			ElementType: types.StringType,
		},

		"custom_tags_verbatim": datasourceschema.MapAttribute{
			Description: "Map of optional tags, in the form: 'TagName = TagValue'. Names are used as specified, without custom prefix",
			Optional:    true,
			ElementType: types.StringType,
		},

		"description": providerschema.StringAttribute{
			Description: "Free-text description for the resource being tagged",
			Optional:    true,
			Required:    false,
		},

		"is_create_before_destroy": providerschema.BoolAttribute{
			Description: "If we want to recycle a resource by creating a new instance before destroying the existing one, then names must be different otherwise the process will fail. We do this by appending a hashed value to a common name-prefix. This is an override flag of the resource type specific property",
			Optional:    true,
			Required:    false,
		},

		"is_force_generated_name": providerschema.BoolAttribute{
			Description: "Expliticly decide to ignore var.custom_name and automatically generate a new name (used to ease transition of resources with legacy names)",
			Optional:    true,
			Required:    false,
		},

		"owner": providerschema.StringAttribute{
			Description: "Owner/maintainer of the AWS resource being tagged",
			Optional:    true,
			Required:    false,
		},

		"region": providerschema.StringAttribute{
			Description: "AWS region where the resource is being provisioned in",
			Optional:    true,
			Required:    false,
		},

		"role": providerschema.StringAttribute{
			Description: "Role played by AWS resource in the context of its environment, e.g. database, security, ...",
			Optional:    true,
			Required:    false,
		},
	},
}

// // This is the 'object'
var ProviderSchema = providerschema.Schema{
	Attributes: map[string]providerschema.Attribute{

		"account": providerschema.StringAttribute{
			Description: "Explicit AWS account ID to use to compute tags",
			Optional:    true,
			Required:    false,
		},

		"accounts_coding": datasourceschema.MapNestedAttribute{
			Description: "Configuration for supported AWS accounts to allow tags and names to be correctly populated. If not specified, default configuration will be used",
			Optional:    true,
			Required:    false,
			NestedObject: datasourceschema.NestedAttributeObject{
				Attributes: map[string]datasourceschema.Attribute{
					"class": datasourceschema.StringAttribute{
						Required: true,
					},
					"name": datasourceschema.StringAttribute{
						Required: true,
					},
					"name_canonical": datasourceschema.StringAttribute{
						Required: true,
					},
					"name_encoded": datasourceschema.StringAttribute{
						Required: true,
					},
				},
			},
		},

		"app_ecosystem": providerschema.StringAttribute{
			Description: "Application/service ecosystem name. This is a coherent collection of application environments, e.g. all the microservices making up a distributed system",
			Optional:    true,
			Required:    false,
		},

		"app_environment": providerschema.StringAttribute{
			Description: "Application/service environment name. This is a logical environment, multiple logical environments can live on the same AWS account. By default gets the same value from infra_environment",
			Optional:    true,
			Required:    false,
		},

		"infra_environment": providerschema.StringAttribute{
			Description: "Infrastructure environment name. This typically coincides with a VPC: this is our environment container, mutliple VPC/environments can live on the same AWS account",
			Optional:    true,
			Required:    false,
		},

		"company_name_long": providerschema.StringAttribute{
			Description: "Name of the company, extended descripting version. This could probably become a constant...",
			Optional:    true,
			Required:    false,
		},

		"company_name_short": providerschema.StringAttribute{
			Description: "Name of the company, compact version. This could probably become a constant...",
			Optional:    true,
			Required:    false,
		},

		"project_name_long": providerschema.StringAttribute{
			Description: "Project name, extended descripting version",
			Optional:    true,
			Required:    false,
		},

		"project_name_short": providerschema.StringAttribute{
			Description: "Project name, compact version",
			Optional:    true,
			Required:    false,
		},

		"resource_set_long": providerschema.StringAttribute{
			Description: "Logical name of the resource set (all the resources provisioned by a single Terraform root module), extended descripting version",
			Optional:    true,
			Required:    false,
		},

		"resource_set_short": providerschema.StringAttribute{
			Description: "Logical name of the resource set (all the resources provisioned by a single Terraform root module), compact version",
			Optional:    true,
			Required:    false,
		},

		"terraform_module": providerschema.StringAttribute{
			Description: "Terraform module name used to provision this resource being tagged.",
			Optional:    true,
			Required:    false,
		},

		"terraform_workspace": providerschema.StringAttribute{
			Description: "Terraform workspace name used to provision this resource being tagged.",
			Optional:    true,
			Required:    false,
		},
	},
}

// // resourceTaggingroviderModel maps provider schema data to a Go type.
// // This is the 'struct'
type ProviderModel struct {
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

	//// Provider parameters
	Account            types.String                  `tfsdk:"account"`
	AccountsCoding     map[string]AccountCodingModel `tfsdk:"accounts_coding"`
	AppEcosystem       types.String                  `tfsdk:"app_ecosystem"`
	AppEnvironment     types.String                  `tfsdk:"app_environment"`
	InfraEnvironment   types.String                  `tfsdk:"infra_environment"`
	CompanyNameLong    types.String                  `tfsdk:"company_name_long"`
	CompanyNameShort   types.String                  `tfsdk:"company_name_short"`
	ProjectNameLong    types.String                  `tfsdk:"project_name_long"`
	ProjectNameShort   types.String                  `tfsdk:"project_name_short"`
	ResourceSetLong    types.String                  `tfsdk:"resource_set_long"`
	ResourceSetShort   types.String                  `tfsdk:"resource_set_short"`
	TerraformModule    types.String                  `tfsdk:"terraform_module"`
	TerraformWorkspace types.String                  `tfsdk:"terraform_workspace"`
}

type AccountCodingModel struct {
	Class         types.String `tfsdk:"class"`
	Name          types.String `tfsdk:"name"`
	NameCanonical types.String `tfsdk:"name_canonical"`
	NameEncoded   types.String `tfsdk:"name_encoded"`
}

type ProviderConfiguration struct {
	Compliance              string
	CostCentre              string
	CustomFQDN              string
	CustomTags              map[string]string
	CustomTagsVerbatim      map[string]string
	Description             string
	IsCreateBeforeDestroy   bool
	IsForceGeneratedName    bool
	IsPropagateTagsAtLaunch bool
	Owner                   string
	Region                  string
	Role                    string

	Account            string
	AccountsCoding     map[string]AccountCodingConfiguration
	AppEcosystem       string
	AppEnvironment     string
	CompanyNameLong    string
	CompanyNameShort   string
	InfraEnvironment   string
	ProjectNameLong    string
	ProjectNameShort   string
	ResourceSetLong    string
	ResourceSetShort   string
	TerraformModule    string
	TerraformWorkspace string
}

type AccountCodingConfiguration struct {
	Class         string
	Name          string
	NameCanonical string
	NameEncoded   string
}
