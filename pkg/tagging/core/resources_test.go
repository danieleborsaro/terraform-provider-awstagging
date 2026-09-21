package tagging_test

import (
	"context"
	"maps"
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

func TestResourceGenerate_ReservedAwsPrefixValueIsKept(t *testing.T) {
	cfg := standardInputConfig()
	cfg.CustomTags = map[string]string{"Reserved": "aws:reserved"}

	res := awsTagging.GetAwsEcsCluster()
	if err := res.Generate(context.Background(), cfg); err != nil {
		t.Fatalf("unexpected generation error: %v", err)
	}

	if got := res.GetTags()["Foo:Custom:Reserved"]; got != "aws:reserved" {
		t.Fatalf("expected a value starting with aws: to be kept, got %q", got)
	}

	if got := res.GetTags()["Foo:Environment:ResourceType"]; got != "AWS::ECS::Cluster" {
		t.Fatalf("unexpected resource type tag: got %q", got)
	}
}

func TestResourceGenerate_ReservedAwsPrefixKeyIsEscaped(t *testing.T) {
	cfg := standardInputConfig()
	cfg.CustomTagsVerbatim = map[string]string{"aws:reserved": "value"}

	res := awsTagging.GetAwsEcsCluster()
	if err := res.Generate(context.Background(), cfg); err != nil {
		t.Fatalf("unexpected generation error: %v", err)
	}

	if _, ok := res.GetTags()["aws:reserved"]; ok {
		t.Fatal("a key starting with aws: should not be passed through")
	}

	if got := res.GetTags()[":aws:reserved"]; got != "value" {
		t.Fatalf("expected the key to be escaped, got value %q", got)
	}
}

func TestResourceGenerate_InvalidConfigurationIsAnError(t *testing.T) {
	cases := map[string]func(*taggingdata.InputConfiguration){
		"account not in accounts_coding": func(c *taggingdata.InputConfiguration) { c.Account = "123456789012" },
		"unknown account class": func(c *taggingdata.InputConfiguration) {
			c.AccountsCoding["811635568629"] = taggingdata.AccountCodingConfiguration{Class: "staging"}
		},
		"region without three parts":  func(c *taggingdata.InputConfiguration) { c.Region = "euwest1" },
		"region with an empty part":   func(c *taggingdata.InputConfiguration) { c.Region = "eu--1" },
		"availability zone too short": func(c *taggingdata.InputConfiguration) { c.AvailabilityZone = "eu-west" },
	}

	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			cfg := standardInputConfig()
			mutate(cfg)

			res := awsTagging.GetAwsEcsCluster()
			if err := res.Generate(context.Background(), cfg); err == nil {
				t.Fatal("expected an error, got none")
			}
		})
	}
}

func TestResourceGenerate_TagLimitDropsInFixedOrder(t *testing.T) {
	cfg := standardInputConfig()
	cfg.CustomTags = map[string]string{"Key1": "1", "Key2": "2", "Key3": "3", "Key4": "4", "Key5": "5", "Key6": "6"}
	cfg.CustomTagsVerbatim = map[string]string{"Extra1": "1", "Extra2": "2", "Extra3": "3"}

	var first map[string]string
	for range 50 {
		res := awsTagging.GetAwsS3Object()
		if err := res.Generate(context.Background(), cfg); err != nil {
			t.Fatalf("unexpected generation error: %v", err)
		}

		if first == nil {
			first = res.GetTags()
			if got := strings.Join(res.GetDroppedTags(), ","); got != "Foo:Custom:Key6,Extra1,Extra2,Extra3" {
				t.Fatalf("unexpected dropped tags: %s", got)
			}
			continue
		}

		if !maps.Equal(first, res.GetTags()) {
			t.Fatalf("tags differ between runs: %v and %v", first, res.GetTags())
		}
	}

	if len(first) != 10 {
		t.Fatalf("expected 10 tags, got %d: %v", len(first), first)
	}
	for _, k := range []string{"Name", "Foo:Business:Owner", "Foo:Custom:Key1", "Foo:Custom:Key5"} {
		if _, ok := first[k]; !ok {
			t.Fatalf("expected %s to be kept, got %v", k, first)
		}
	}
}
