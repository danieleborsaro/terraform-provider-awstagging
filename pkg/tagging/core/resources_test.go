package tagging_test

import (
	"context"
	"strings"
	"testing"

	awsTagging "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/aws"
	taggingdata "github.com/danieleborsaro/terraform-provider-awstagging/pkg/tagging/shared"
)

func standardInputConfig() *taggingdata.InputConfiguration {
	return &taggingdata.InputConfiguration{
		Account: "811635568629",
		AccountsCoding: map[string]taggingdata.AccountCodingConfiguration{
			"811635568629": {
				Class:         "production",
				Name:          "devopsprod",
				NameCanonical: "DevOpsProd",
				NameEncoded:   "DVP",
			},
		},
		AppEcosystem:          "dev1",
		AppEnvironment:        "dev1",
		AvailabilityZone:      "eu-west-1a",
		CompanyNameLong:       "My Company Foo",
		CompanyNameShort:      "Foo",
		Compliance:            "tbc",
		CostCentre:            "n/a",
		CustomTags:            map[string]string{"MyKey1": "My value 1", "MyKey2": "My value 2"},
		CustomTagsVerbatim:    map[string]string{"MyKeyVerbatim1": "My value verbatim 1", "MyKeyVerbatim2": "My value verbatim 2"},
		Description:           "This is my provider. I am proud of it.",
		InfraEnvironment:      "dev1",
		IsCreateBeforeDestroy: true,
		IsForceGeneratedName:  false,
		Owner:                 "daniele@borsaro.it",
		ProjectNameLong:       "My test project",
		ProjectNameShort:      "mytstprjct",
		Region:                "eu-west-1",
		ResourceSetLong:       "My test resource set",
		ResourceSetShort:      "mytstrsrst",
		Role:                  "provider",
		TerraformModule:       "/tmp/testmodule",
		TerraformWorkspace:    "default",
		VersioningSources:     nil,
	}
}

func TestResourceGenerate_EcsCluster_StandardNameAndTags(t *testing.T) {
	cfg := standardInputConfig()
	res := awsTagging.GetAwsEcsCluster()
	if err := res.Generate(context.Background(), cfg); err != nil {
		t.Fatalf("unexpected generation error: %v", err)
	}

	wantName := "P-dev1-mytstprjct-mytstrsrst-clstr-Provider-euw1-e441f8d0a471f44763e72e93f6cf9e8803f74ac0"
	if got := res.GetName().Safe; got != wantName {
		t.Fatalf("unexpected generated name: got %q, want %q", got, wantName)
	}

	if got := res.GetTags()["Name"]; got != wantName {
		t.Fatalf("unexpected Name tag: got %q, want %q", got, wantName)
	}

	if got := res.GetTags()["Foo:Business:Owner"]; got != "daniele@borsaro.it" {
		t.Fatalf("unexpected owner tag: got %q", got)
	}

	if got := res.GetTags()["Foo:Environment:Account"]; got != "DevOpsProd" {
		t.Fatalf("unexpected account tag: got %q", got)
	}

	if got := res.GetTags()["Foo:Custom:MyKey1"]; got != "My value 1" {
		t.Fatalf("unexpected custom tag: got %q", got)
	}

	if got := res.GetTags()["MyKeyVerbatim1"]; got != "My value verbatim 1" {
		t.Fatalf("unexpected verbatim tag: got %q", got)
	}
}

func TestResourceGenerate_RequiresEnvironment(t *testing.T) {
	cfg := standardInputConfig()
	cfg.InfraEnvironment = ""
	cfg.AppEnvironment = ""

	res := awsTagging.GetAwsEcsCluster()
	if err := res.Generate(context.Background(), cfg); err == nil {
		t.Fatal("expected an error when both environments are empty")
	}
}

func TestResourceGenerate_RejectsBothCustomNameOptions(t *testing.T) {
	cfg := standardInputConfig()
	cfg.CustomName = "billing-api"
	cfg.CustomNamePrefix = "billing"

	res := awsTagging.GetAwsEcsCluster()
	if err := res.Generate(context.Background(), cfg); err == nil {
		t.Fatal("expected an error when both custom name options are set")
	}
}

func TestResourceGenerate_CustomNameOverride(t *testing.T) {
	cfg := standardInputConfig()
	cfg.CustomName = "billing-api"
	cfg.CustomNamePrefix = ""
	cfg.IsForceGeneratedName = false

	res := awsTagging.GetAwsEcsCluster()
	if err := res.Generate(context.Background(), cfg); err != nil {
		t.Fatalf("unexpected generation error: %v", err)
	}

	if got := res.GetName().Safe; got != "billing-api" {
		t.Fatalf("unexpected custom-name override: got %q, want %q", got, "billing-api")
	}

	if got := res.GetTags()["Name"]; got != "billing-api" {
		t.Fatalf("unexpected Name tag for custom-name override: got %q, want %q", got, "billing-api")
	}
}

func TestResourceGenerate_VersionedNameWithCustomPrefix(t *testing.T) {
	cfg := standardInputConfig()
	cfg.CustomName = ""
	cfg.CustomNamePrefix = "prefix-demo"
	cfg.VersioningSources = []string{"build-42"}

	res := awsTagging.GetAwsEcsCluster()
	if err := res.Generate(context.Background(), cfg); err != nil {
		t.Fatalf("unexpected generation error: %v", err)
	}

	if got := res.GetName().SafePrefix; !strings.HasPrefix(got, "prefix-demo") {
		t.Fatalf("generated name prefix mismatch: got %q", got)
	}

	if got := res.GetName().SafeSuffix; got == "" {
		t.Fatalf("expected a versioned suffix when custom name prefix is set")
	}

	if got := res.GetTags()["Name"]; !strings.HasPrefix(got, "prefix-demo") {
		t.Fatalf("Name tag should be versioned with custom prefix, got %q", got)
	}
}

func TestResourceGenerate_ReservedAwsPrefixValueIsEscaped(t *testing.T) {
	cfg := standardInputConfig()
	cfg.CustomTags = map[string]string{"Reserved": "aws:reserved"}

	res := awsTagging.GetAwsEcsCluster()
	if err := res.Generate(context.Background(), cfg); err != nil {
		t.Fatalf("unexpected generation error: %v", err)
	}

	if got := res.GetTags()["Foo:Custom:Reserved"]; got != ":aws:reserved" {
		t.Fatalf("expected reserved AWS prefix to be escaped, got %q", got)
	}
}
