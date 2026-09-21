package datasources

import (
	"context"
	"fmt"
	"maps"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	models "github.com/danieleborsaro/terraform-provider-awstagging/internal/shared"
	awsTagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
)

func TestTaggingDataSource_MergeConfiguration(t *testing.T) {
	t.Run("uses provider defaults when datasource override is unset", func(t *testing.T) {
		ds := &TaggingDataSource{
			DatasourceType: "test",
			providerConfiguration: &models.ProviderConfiguration{
				Compliance:            "provider-compliance",
				CostCentre:            "provider-cost-centre",
				CustomTags:            map[string]string{"ProviderOnly": "provider-value"},
				CustomTagsVerbatim:    map[string]string{"ProviderOnlyVerbatim": "provider-verbatim-value"},
				Description:           "provider-description",
				IsCreateBeforeDestroy: true,
				IsForceGeneratedName:  true,
				Owner:                 "provider-owner",
				Region:                "eu-west-1",
				Role:                  "provider-role",
			},
		}

		cfg := &DataSourceModel{
			Compliance:            types.StringNull(),
			CostCentre:            types.StringNull(),
			CustomTags:            nil,
			CustomTagsVerbatim:    nil,
			Description:           types.StringNull(),
			IsCreateBeforeDestroy: types.BoolNull(),
			IsForceGeneratedName:  types.BoolNull(),
			Owner:                 types.StringNull(),
			Region:                types.StringNull(),
			Role:                  types.StringNull(),
		}

		ds.MergeConfiguration(context.Background(), cfg, datasource.ReadRequest{}, &datasource.ReadResponse{})

		if got := cfg.Compliance.ValueString(); got != "provider-compliance" {
			t.Fatalf("expected provider compliance default, got %q", got)
		}
		if got := cfg.CostCentre.ValueString(); got != "provider-cost-centre" {
			t.Fatalf("expected provider cost centre default, got %q", got)
		}
		if got := cfg.CustomTags["ProviderOnly"].ValueString(); got != "provider-value" {
			t.Fatalf("expected provider custom tag to be retained, got %q", got)
		}
		if got := cfg.CustomTagsVerbatim["ProviderOnlyVerbatim"].ValueString(); got != "provider-verbatim-value" {
			t.Fatalf("expected provider verbatim tag to be retained, got %q", got)
		}
		if got := cfg.IsCreateBeforeDestroy.ValueBool(); !got {
			t.Fatalf("expected provider is_create_before_destroy default, got %t", got)
		}
	})

	t.Run("datasource values override provider defaults", func(t *testing.T) {
		ds := &TaggingDataSource{
			DatasourceType: "test",
			providerConfiguration: &models.ProviderConfiguration{
				Compliance:            "provider-compliance",
				CostCentre:            "provider-cost-centre",
				CustomTags:            map[string]string{"ProviderOnly": "provider-value"},
				CustomTagsVerbatim:    map[string]string{"ProviderOnlyVerbatim": "provider-verbatim-value"},
				Description:           "provider-description",
				IsCreateBeforeDestroy: true,
				IsForceGeneratedName:  true,
				Owner:                 "provider-owner",
				Region:                "eu-west-1",
				Role:                  "provider-role",
			},
		}

		cfg := &DataSourceModel{
			Compliance:            types.StringValue("override-compliance"),
			CostCentre:            types.StringValue("override-cost-centre"),
			CustomTags:            map[string]types.String{"OverrideOnly": types.StringValue("override-value")},
			CustomTagsVerbatim:    map[string]types.String{"OverrideOnlyVerbatim": types.StringValue("override-verbatim-value")},
			Description:           types.StringValue("override-description"),
			IsCreateBeforeDestroy: types.BoolValue(false),
			IsForceGeneratedName:  types.BoolValue(false),
			Owner:                 types.StringValue("override-owner"),
			Region:                types.StringValue("eu-central-1"),
			Role:                  types.StringValue("override-role"),
		}

		ds.MergeConfiguration(context.Background(), cfg, datasource.ReadRequest{}, &datasource.ReadResponse{})

		if got := cfg.Compliance.ValueString(); got != "override-compliance" {
			t.Fatalf("expected datasource compliance override, got %q", got)
		}
		if got := cfg.CostCentre.ValueString(); got != "override-cost-centre" {
			t.Fatalf("expected datasource cost centre override, got %q", got)
		}
		if got := cfg.CustomTags["OverrideOnly"].ValueString(); got != "override-value" {
			t.Fatalf("expected datasource custom tag override, got %q", got)
		}
		if got := cfg.CustomTagsVerbatim["OverrideOnlyVerbatim"].ValueString(); got != "override-verbatim-value" {
			t.Fatalf("expected datasource verbatim tag override, got %q", got)
		}
		if got := cfg.IsCreateBeforeDestroy.ValueBool(); got {
			t.Fatalf("expected datasource bool override to win, got %t", got)
		}
	})
}

func testProviderConfiguration() *models.ProviderConfiguration {
	return &models.ProviderConfiguration{
		Account: "123456789012",
		AccountsCoding: map[string]models.AccountCodingConfiguration{
			"123456789012": {Class: "production", Name: "prod", NameCanonical: "Prod", NameEncoded: "P"},
		},
		CompanyNameShort:      "Acme",
		Compliance:            "internal",
		CostCentre:            "platform",
		InfraEnvironment:      "prod",
		IsCreateBeforeDestroy: true,
		Owner:                 "team@example.com",
		ProjectNameLong:       "Example project",
		ProjectNameShort:      "example",
		Region:                "eu-west-1",
		ResourceSetLong:       "Example resource set",
		ResourceSetShort:      "example",
		Role:                  "app",
		TerraformModule:       "/tmp/example",
		TerraformWorkspace:    "default",
	}
}

func TestConfigurationDataSource_TagsForDefaultTags(t *testing.T) {
	ds := NewConfigurationDataSource().(*configurationDataSource)
	ds.providerConfiguration = testProviderConfiguration()

	for name, set := range map[string]func(*DataSourceModel){
		"no name settings":   func(*DataSourceModel) {},
		"custom_name":        func(c *DataSourceModel) { c.CustomName = types.StringValue("my-name") },
		"custom_name_prefix": func(c *DataSourceModel) { c.CustomNamePrefix = types.StringValue("my-prefix") },
	} {
		t.Run(name, func(t *testing.T) {
			cfg := &DataSourceModel{}
			set(cfg)
			resp := &datasource.ReadResponse{}
			ds.MergeConfiguration(context.Background(), cfg, datasource.ReadRequest{}, resp)
			state := *cfg
			ds.UpdateState(context.Background(), &state, cfg, datasource.ReadRequest{}, resp)

			if resp.Diagnostics.HasError() {
				t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
			}
			if got := state.Tags["Acme:Business:Owner"]; got != "team@example.com" {
				t.Fatalf("expected the common tags, got owner %q in %v", got, state.Tags)
			}
			if _, ok := state.Tags["Name"]; ok {
				t.Fatalf("default tags should have no Name, got %v", state.Tags)
			}
			if _, ok := state.Tags["Acme:Environment:ResourceType"]; ok {
				t.Fatalf("default tags should have no ResourceType, got %v", state.Tags)
			}
			if got := state.IsCreateBeforeDestroy.ValueBool(); !got {
				t.Fatalf("is_create_before_destroy in state should keep the configured value, got %t", got)
			}
		})
	}
}

func TestTaggingDataSource_WarnsAboutDroppedTags(t *testing.T) {
	pc := testProviderConfiguration()
	pc.CustomTagsVerbatim = map[string]string{"Extra1": "1", "Extra2": "2", "Extra3": "3", "Extra4": "4", "Extra5": "5", "Extra6": "6", "Extra7": "7", "Extra8": "8"}
	ds := &TaggingDataSource{DatasourceType: "test", Tagging: awsTagging.GetAwsS3Object(), providerConfiguration: pc}

	cfg := &DataSourceModel{}
	resp := &datasource.ReadResponse{}
	ds.MergeConfiguration(context.Background(), cfg, datasource.ReadRequest{}, resp)
	state := *cfg
	ds.UpdateState(context.Background(), &state, cfg, datasource.ReadRequest{}, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
	}
	if got := resp.Diagnostics.WarningsCount(); got != 1 {
		t.Fatalf("expected one warning about dropped tags, got %d: %v", got, resp.Diagnostics)
	}
	if got := len(state.Tags); got != 10 {
		t.Fatalf("expected 10 tags, got %d", got)
	}
}

func TestConfigurationDataSource_DefaultAndResourceTagsFitTheLimit(t *testing.T) {
	read := func(t *testing.T, ds *TaggingDataSource, update func(context.Context, *DataSourceModel, *DataSourceModel, datasource.ReadRequest, *datasource.ReadResponse)) map[string]string {
		cfg := &DataSourceModel{}
		resp := &datasource.ReadResponse{}
		ds.MergeConfiguration(context.Background(), cfg, datasource.ReadRequest{}, resp)
		state := *cfg
		update(context.Background(), &state, cfg, datasource.ReadRequest{}, resp)
		if resp.Diagnostics.HasError() {
			t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
		}
		return state.Tags
	}

	for _, n := range []int{0, 37, 45, 48, 60} {
		t.Run(fmt.Sprintf("%d custom tags", n), func(t *testing.T) {
			pc := testProviderConfiguration()
			pc.CustomTags = map[string]string{}
			for i := range n {
				pc.CustomTags[fmt.Sprintf("Key%02d", i)] = "v"
			}

			config := NewConfigurationDataSource().(*configurationDataSource)
			config.providerConfiguration = pc
			defaults := read(t, &config.TaggingDataSource, config.UpdateState)

			bucket := &TaggingDataSource{DatasourceType: "test", Tagging: awsTagging.GetAwsS3Bucket(), providerConfiguration: pc}
			resource := read(t, bucket, bucket.UpdateState)

			merged := maps.Clone(defaults)
			maps.Copy(merged, resource)
			if len(merged) > 50 {
				t.Fatalf("default and resource tags merge to %d tags, over the limit of 50", len(merged))
			}
		})
	}
}
