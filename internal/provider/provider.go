package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	aws "github.com/danieleborsaro/terraform-provider-awstagging/internal/aws"
	"github.com/danieleborsaro/terraform-provider-awstagging/internal/datasources"
	models "github.com/danieleborsaro/terraform-provider-awstagging/internal/shared"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &resourceTaggingProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &resourceTaggingProvider{
			version: version,
		}
	}
}

// resourceTaggingProvider is the provider implementation.
type resourceTaggingProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *resourceTaggingProvider) Metadata(ctx context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	tflog.Trace(ctx, "provider - BEGIN configuring Metadata")

	resp.TypeName = "awstagging"
	resp.Version = p.version

	tflog.Trace(ctx, "provider - END configuring Metadata")
}

// Schema defines the provider-level schema for configuration data.
func (p *resourceTaggingProvider) Schema(ctx context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	tflog.Trace(ctx, "provider - BEGIN configuring schema")
	resp.Schema = models.ProviderSchema
	// Any override option must be defined in this schema.
	for k, v := range models.OverridesSchema.Attributes {
		resp.Schema.Attributes[k] = v
	}
	tflog.Trace(ctx, "provider - END configuring schema")
}

// // Configure prepares a aws-tagging client for data sources and resources.
func (p *resourceTaggingProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Trace(ctx, "provider - BEGIN retrieving configuration")

	// Retrieve provider data from configuration
	var config models.ProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//// Mandatory
	account := ""
	if !config.Account.IsNull() {
		account = config.Account.ValueString()
	}
	if account == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("account"),
			"Missing aws-tagging property account",
			"The provider cannot create the aws-tagging client as there is a missing or empty value for the aws-tagging account. "+
				"Set the account value in the configuration. ",
		)
	}

	//// Mandatory
	var accountsCoding = make(map[string]models.AccountCodingConfiguration)
	for k, v := range config.AccountsCoding {
		thisCoding := models.AccountCodingConfiguration{
			Class:         v.Class.ValueString(),
			Name:          v.Name.ValueString(),
			NameCanonical: v.NameCanonical.ValueString(),
			NameEncoded:   v.NameEncoded.ValueString(),
		}

		accountsCoding[k] = thisCoding
	}
	if len(accountsCoding) == 0 {
		resp.Diagnostics.AddAttributeError(
			path.Root("accounts_coding"),
			"Missing aws-tagging property accounts_coding",
			"The provider cannot create the aws-tagging client as there is a missing or empty value for the aws-tagging accounts_coding. "+
				"Set the accounts_coding value in the configuration. ",
		)
	}

	//// Optional
	appEcosystem := ""
	if !config.AppEcosystem.IsNull() {
		appEcosystem = config.AppEcosystem.ValueString()
	}

	//// Optional
	appEnvironment := ""
	if !config.AppEnvironment.IsNull() {
		appEnvironment = config.AppEnvironment.ValueString()
	}

	//// Optional
	companyNameLong := "My Company Foo"
	if !config.CompanyNameLong.IsNull() {
		companyNameLong = config.CompanyNameLong.ValueString()
	}

	//// Optional
	companyNameShort := "Foo"
	if !config.CompanyNameShort.IsNull() {
		companyNameShort = config.CompanyNameShort.ValueString()
	}

	//// Mandatory
	compliance := ""
	if !config.Compliance.IsNull() {
		compliance = config.Compliance.ValueString()
	}
	if compliance == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("compliance"),
			"Missing aws-tagging property compliance",
			"The provider cannot create the aws-tagging client as there is a missing or empty value for the aws-tagging compliance. "+
				"Set the compliance value in the configuration. ",
		)
	}

	//// Mandatory
	costCentre := ""
	if !config.CostCentre.IsNull() {
		costCentre = config.CostCentre.ValueString()
	}
	if costCentre == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("cost_centre"),
			"Missing aws-tagging property cost_centre",
			"The provider cannot create the aws-tagging client as there is a missing or empty value for the aws-tagging cost_centre. "+
				"Set the cost_centre value in the configuration. ",
		)
	}

	//// Optional
	var customTags = make(map[string]string)
	for k, v := range config.CustomTags {
		customTags[k] = v.ValueString()
	}

	//// Optional
	var customTagsVerbatim = make(map[string]string)
	for k, v := range config.CustomTagsVerbatim {
		customTagsVerbatim[k] = v.ValueString()
	}

	//// Optional
	description := ""
	if !config.Description.IsNull() {
		description = config.Description.ValueString()
	}

	//// Mandatory
	if config.InfraEnvironment.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("infra_environment"),
			"Unknown aws-tagging property infra_environment",
			"The provider cannot create the aws-tagging client as there is an unknown configuration value for the aws-tagging property infra_environment. "+
				"Either target apply the source of the value first or set the value statically in the configuration.",
		)
	}

	infraEnvironment := ""
	if !config.InfraEnvironment.IsNull() {
		infraEnvironment = config.InfraEnvironment.ValueString()
	}
	if infraEnvironment == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("infra_environment"),
			"Missing aws-tagging property infra_environment",
			"The provider cannot create the aws-tagging client as there is a missing or empty value for the aws-tagging infra_environment. "+
				"Set the infra_environment value in the configuration. ",
		)
	}

	//// Optional
	isCreateBeforeDestroy := true
	if !config.IsCreateBeforeDestroy.IsNull() {
		isCreateBeforeDestroy = config.IsCreateBeforeDestroy.ValueBool()
	}

	//// Optional
	isForceGeneratedName := false
	if !config.IsForceGeneratedName.IsNull() {
		isForceGeneratedName = config.IsForceGeneratedName.ValueBool()
	}

	//// Mandatory
	owner := ""
	if !config.Owner.IsNull() {
		owner = config.Owner.ValueString()
	}
	if owner == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("owner"),
			"Missing aws-tagging property owner",
			"The provider cannot create the aws-tagging client as there is a missing or empty value for the aws-tagging owner. "+
				"Set the owner value in the configuration. ",
		)
	}

	//// Mandatory
	projectNameLong := ""
	if !config.ProjectNameLong.IsNull() {
		projectNameLong = config.ProjectNameLong.ValueString()
	}
	if projectNameLong == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("project_name_long"),
			"Missing aws-tagging property project_name_long",
			"The provider cannot create the aws-tagging client as there is a missing or empty value for the aws-tagging project_name_long. "+
				"Set the project_name_long value in the configuration. ",
		)
	}

	//// Mandatory
	projectNameShort := ""
	if !config.ProjectNameShort.IsNull() {
		projectNameShort = config.ProjectNameShort.ValueString()
	}
	if projectNameShort == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("project_name_short"),
			"Missing aws-tagging property project_name_short",
			"The provider cannot create the aws-tagging client as there is a missing or empty value for the aws-tagging project_name_short. "+
				"Set the project_name_short value in the configuration. ",
		)
	}

	//// Mandatory
	region := ""
	if !config.Region.IsNull() {
		region = config.Region.ValueString()
	}
	if region == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("region"),
			"Missing aws-tagging property region",
			"The provider cannot create the aws-tagging client as there is a missing or empty value for the aws-tagging region. "+
				"Set the region value in the configuration. ",
		)
	}

	//// Mandatory
	resourceSetLong := ""
	if !config.ResourceSetLong.IsNull() {
		resourceSetLong = config.ResourceSetLong.ValueString()
	}
	if resourceSetLong == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("resource_set_long"),
			"Missing aws-tagging property resource_set_long",
			"The provider cannot create the aws-tagging client as there is a missing or empty value for the aws-tagging resource_set_long. "+
				"Set the resource_set_long value in the configuration. ",
		)
	}

	//// Mandatory
	resourceSetShort := ""
	if !config.ResourceSetShort.IsNull() {
		resourceSetShort = config.ResourceSetShort.ValueString()
	}
	if resourceSetShort == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("resource_set_short"),
			"Missing aws-tagging property resource_set_short",
			"The provider cannot create the aws-tagging client as there is a missing or empty value for the aws-tagging resource_set_short. "+
				"Set the resource_set_short value in the configuration. ",
		)
	}

	//// Optional
	role := ""
	if !config.Role.IsNull() {
		role = config.Role.ValueString()
	}

	//// Mandatory
	terraformModule := ""
	if !config.TerraformModule.IsNull() {
		terraformModule = config.TerraformModule.ValueString()
	}
	if terraformModule == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("terraform_module"),
			"Missing aws-tagging property terraform_module",
			"The provider cannot create the aws-tagging client as there is a missing or empty value for the aws-tagging terraform_module. "+
				"Set the terraform_module value in the configuration. ",
		)
	}

	//// Mandatory
	terraformWorkspace := ""
	if !config.TerraformWorkspace.IsNull() {
		terraformWorkspace = config.TerraformWorkspace.ValueString()
	}
	if terraformWorkspace == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("terraform_workspace"),
			"Missing aws-tagging property terraform_workspace",
			"The provider cannot create the aws-tagging client as there is a missing or empty value for the aws-tagging terraform_workspace. "+
				"Set the terraform_workspace value in the configuration. ",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	client := &models.ProviderConfiguration{
		Compliance:            compliance,
		CostCentre:            costCentre,
		CustomTags:            customTags,
		CustomTagsVerbatim:    customTagsVerbatim,
		Description:           description,
		IsCreateBeforeDestroy: isCreateBeforeDestroy,
		IsForceGeneratedName:  isForceGeneratedName,
		Owner:                 owner,
		Region:                region,
		Role:                  role,

		Account:            account,
		AccountsCoding:     accountsCoding,
		AppEcosystem:       appEcosystem,
		AppEnvironment:     appEnvironment,
		CompanyNameLong:    companyNameLong,
		CompanyNameShort:   companyNameShort,
		InfraEnvironment:   infraEnvironment,
		ProjectNameLong:    projectNameLong,
		ProjectNameShort:   projectNameShort,
		ResourceSetLong:    resourceSetLong,
		ResourceSetShort:   resourceSetShort,
		TerraformModule:    terraformModule,
		TerraformWorkspace: terraformWorkspace,
	}

	// Make the aws-tagging Client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "awstagging - provider configured")

	tflog.Trace(ctx, "provider - END retrieving configuration")
}

// DataSources defines the data sources implemented in the provider.
func (p *resourceTaggingProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewConfigurationDataSource,

		aws.NewAwsEc2SnapshotDataSource,
		aws.NewAwsCognitoUserPoolDataSource,
		aws.NewAwsIamPolicyDataSource,
		aws.NewAwsInternetGatewayDataSource,
		aws.NewAwsEksClusterDataSource,
		aws.NewAwsCloudwatchEventRuleDataSource,
		aws.NewAwsEc2ClientVpnEndpointDataSource,
		aws.NewAwsCloudwatchDashboardDataSource,
		aws.NewAwsAutoscalingGroupDataSource,
		aws.NewAwsEcsClusterDataSource,
		aws.NewAwsSpotInstanceRequestDataSource,
		aws.NewAwsAcmCertificateDataSource,
		aws.NewAwsWafregionalRegexMatchSetDataSource,
		aws.NewAwsEcrRepositoryDataSource,
		aws.NewAwsSesConfigurationSetDataSource,
		aws.NewAwsS3ObjectDataSource,
		aws.NewAwsSecurityGroupDataSource,
		aws.NewAwsCognitoIdentityPoolDataSource,
		aws.NewAwsNetworkInterfaceDataSource,
		aws.NewAwsEc2TransitGatewayVpcAttachmentDataSource,
		aws.NewAwsNetworkAclDataSource,
		aws.NewAwsDbParameterGroupDataSource,
		aws.NewAwsSubnetDataSource,
		aws.NewAwsEc2TransitGatewayRouteTableDataSource,
		aws.NewAwsRoute53HealthCheckDataSource,
		aws.NewAwsEmrClusterDataSource,
		aws.NewAwsRedshiftClusterDataSource,
		aws.NewAwsRoute53ZoneDataSource,
		aws.NewAwsKeyPairDataSource,
		aws.NewAwsRdsReservedInstanceDataSource,
		aws.NewAwsWafv2WebAclLoggingConfigurationDataSource,
		aws.NewAwsKmsAliasDataSource,
		aws.NewAwsWafregionalRuleDataSource,
		aws.NewAwsElbDataSource,
		aws.NewAwsCloudwatchMetricAlarmDataSource,
		aws.NewAwsLambdaFunctionDataSource,
		aws.NewAwsSesReceiptRuleDataSource,
		aws.NewAwsCodebuildProjectDataSource,
		aws.NewAwsCloudfrontDistributionDataSource,
		aws.NewAwsWafv2WebAclAssociationDataSource,
		aws.NewAwsRamResourceShareDataSource,
		aws.NewAwsWafRuleDataSource,
		aws.NewAwsWafv2IpSetDataSource,
		aws.NewAwsWafWebAclDataSource,
		aws.NewAwsSsmParameterDataSource,
		aws.NewAwsEipDataSource,
		aws.NewAwsVpcDhcpOptionsDataSource,
		aws.NewAwsWafv2RuleGroupDataSource,
		aws.NewAwsNatGatewayDataSource,
		aws.NewAwsSesv2ConfigurationSetEventDestinationDataSource,
		aws.NewAwsEfsAccessPointDataSource,
		aws.NewAwsDirectoryServiceDirectoryDataSource,
		aws.NewAwsSesIdentityNotificationTopicDataSource,
		aws.NewAwsLaunchConfigurationDataSource,
		aws.NewAwsLaunchTemplateDataSource,
		aws.NewAwsSesEmailIdentityDataSource,
		aws.NewAwsDatapipelinePipelineDataSource,
		aws.NewAwsEcsServiceDataSource,
		aws.NewAwsWafregionalIpsetDataSource,
		aws.NewAwsLbListenerDataSource,
		aws.NewAwsFlowLogDataSource,
		aws.NewAwsWafv2RegexPatternSetDataSource,
		aws.NewAwsEc2TransitGatewayDataSource,
		aws.NewAwsLbTargetGroupDataSource,
		aws.NewAwsEc2ManagedPrefixListDataSource,
		aws.NewAwsSesActiveReceiptRuleSetDataSource,
		aws.NewAwsWafregionalGeoMatchSetDataSource,
		aws.NewAwsVpcDhcpOptionsAssociationDataSource,
		aws.NewAwsCloudwatchLogStreamDataSource,
		aws.NewAwsCloudwatchLogGroupDataSource,
		aws.NewAwsRdsGlobalClusterDataSource,
		aws.NewAwsRouteTableDataSource,
		aws.NewAwsDbEventSubscriptionDataSource,
		aws.NewAwsConfigConfigRuleDataSource,
		aws.NewAwsS3BucketDataSource,
		aws.NewAwsCloudformationStackDataSource,
		aws.NewAwsEcsTaskDefinitionDataSource,
		aws.NewAwsWafIpsetDataSource,
		aws.NewAwsTransferServerDataSource,
		aws.NewAwsKmsKeyDataSource,
		aws.NewAwsBackupVaultDataSource,
		aws.NewAwsBackupSelectionDataSource,
		aws.NewAwsWafv2WebAclDataSource,
		aws.NewAwsBackupPlanDataSource,
		aws.NewAwsSesReceiptRuleSetDataSource,
		aws.NewAwsRoute53ResolverRuleDataSource,
		aws.NewAwsRoute53ResolverEndpointDataSource,
		aws.NewAwsReservedInstanceDataSource,
		aws.NewAwsDbSnapshotDataSource,
		aws.NewAwsDbInstanceDataSource,
		aws.NewAwsOpensearchDomainDataSource,
		aws.NewAwsRdsClusterDataSource,
		aws.NewAwsInstanceDataSource,
		aws.NewAwsElasticacheClusterDataSource,
		aws.NewAwsVpnConnectionDataSource,
		aws.NewAwsVpcDataSource,
		aws.NewAwsCloudwatchCompositeAlarmDataSource,
		aws.NewAwsCustomerGatewayDataSource,
		aws.NewAwsIamGroupDataSource,
		aws.NewAwsDbOptionGroupDataSource,
		aws.NewAwsDynamodbTableDataSource,
		aws.NewAwsSqsQueueDataSource,
		aws.NewAwsLbDataSource,
		aws.NewAwsIamRoleDataSource,
		aws.NewAwsIamUserDataSource,
		aws.NewAwsCloudformationExportDataSource,
		aws.NewAwsLbListenerRuleDataSource,
		aws.NewAwsAmiDataSource,
		aws.NewAwsWafregionalByteMatchSetDataSource,
		aws.NewAwsDbSubnetGroupDataSource,
		aws.NewAwsWafregionalRegexPatternSetDataSource,
		aws.NewAwsGlacierVaultDataSource,
		aws.NewAwsWafRateBasedRuleDataSource,
		aws.NewAwsEcsCapacityProviderDataSource,
		aws.NewAwsEbsVolumeDataSource,
		aws.NewAwsElasticacheSubnetGroupDataSource,
		aws.NewAwsWafregionalWebAclDataSource,
		aws.NewAwsWafregionalRateBasedRuleDataSource,
		aws.NewAwsSnsTopicDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *resourceTaggingProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
