package provider

import (
	"context"
	"testing"

	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

func TestNew_Metadata(t *testing.T) {
	p := New("test-version")()

	resp := &provider.MetadataResponse{}
	p.Metadata(context.Background(), provider.MetadataRequest{}, resp)

	if resp.TypeName != "awstagging" {
		t.Fatalf("unexpected provider type name: got %q, want %q", resp.TypeName, "awstagging")
	}
	if resp.Version != "test-version" {
		t.Fatalf("unexpected provider version: got %q, want %q", resp.Version, "test-version")
	}
}

func TestProviderSchema_ContainsProviderAndOverrideAttributes(t *testing.T) {
	p := New("test")()
	resp := &provider.SchemaResponse{}
	p.Schema(context.Background(), provider.SchemaRequest{}, resp)

	for _, name := range []string{
		"account",
		"accounts_coding",
		"infra_environment",
		"project_name_long",
		"project_name_short",
		"resource_set_long",
		"resource_set_short",
		"terraform_module",
		"terraform_workspace",
		"compliance",
		"cost_centre",
		"custom_tags",
		"custom_tags_verbatim",
		"owner",
		"region",
		"role",
	} {
		if _, ok := resp.Schema.Attributes[name]; !ok {
			t.Errorf("provider schema is missing attribute %q", name)
		}
	}

	assertOptional := func(name string) {
		t.Helper()
		attribute, ok := resp.Schema.Attributes[name]
		if !ok {
			t.Fatalf("provider schema is missing attribute %q", name)
		}
		if !isOptionalProviderAttribute(attribute) {
			t.Errorf("provider schema attribute %q is not optional", name)
		}
	}

	for _, name := range []string{
		"account",
		"accounts_coding",
		"infra_environment",
		"project_name_long",
		"project_name_short",
		"resource_set_long",
		"resource_set_short",
		"terraform_module",
		"terraform_workspace",
	} {
		assertOptional(name)
	}
}

func isOptionalProviderAttribute(attribute providerschema.Attribute) bool {
	switch value := attribute.(type) {
	case providerschema.StringAttribute:
		return value.Optional
	case providerschema.BoolAttribute:
		return value.Optional
	case providerschema.MapAttribute:
		return value.Optional
	case providerschema.MapNestedAttribute:
		return value.Optional
	case datasourceschema.MapAttribute:
		return value.Optional
	case datasourceschema.MapNestedAttribute:
		return value.Optional
	default:
		return false
	}
}
