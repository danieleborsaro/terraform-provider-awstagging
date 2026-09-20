package datasources

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	models "github.com/danieleborsaro/terraform-provider-awstagging/internal/shared"
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
