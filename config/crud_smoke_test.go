package config

import (
	"context"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http/httptest"
	"os"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	ldProvider "github.com/launchdarkly/terraform-provider-launchdarkly/launchdarkly"
)

type crudCase struct {
	resource string
	// Terraform arguments (snake_case) as JSON, as upjet passes them from spec.forProvider.
	create, update string
}

var crudCases = map[string]crudCase{
	"project-minimal": {"launchdarkly_project",
		`{"key": "proj", "name": "Project", "environments": [{"key": "test", "name": "Test", "color": "000000"}]}`,
		`{"key": "proj", "name": "Project 2", "environments": [{"key": "test", "name": "Test", "color": "000000"}]}`},
	"project": {"launchdarkly_project",
		`{"key": "proj", "name": "Project", "environments": [{"key": "test", "name": "Test", "color": "000000"}],
		  "default_client_side_availability": [{"using_environment_id": false, "using_mobile_key": true}]}`,
		`{"key": "proj", "name": "Project 2", "environments": [{"key": "test", "name": "Test", "color": "000000"}],
		  "default_client_side_availability": [{"using_environment_id": true, "using_mobile_key": true}]}`},
	"environment": {"launchdarkly_environment",
		`{"project_key": "proj", "key": "staging", "name": "Staging", "color": "000000"}`,
		`{"project_key": "proj", "key": "staging", "name": "Staging 2", "color": "ffffff"}`},
	"flag-minimal": {tfFeatureFlag,
		`{"project_key": "proj", "key": "minimal", "name": "Minimal", "variation_type": "boolean"}`,
		`{"project_key": "proj", "key": "minimal", "name": "Minimal", "variation_type": "boolean", "description": "after"}`},
	"flag-boolean": {tfFeatureFlag,
		`{"project_key": "proj", "key": "bool", "name": "Bool", "variation_type": "boolean",
		  "variations": [{"value": "true"}, {"value": "false"}],
		  "defaults": [{"on_variation": 0, "off_variation": 1}], "tags": ["a"]}`,
		`{"project_key": "proj", "key": "bool", "name": "Bool", "variation_type": "boolean",
		  "variations": [{"value": "true"}, {"value": "false"}],
		  "defaults": [{"on_variation": 0, "off_variation": 1}], "tags": ["a", "b"], "temporary": true}`},
	"flag-string": {tfFeatureFlag,
		`{"project_key": "proj", "key": "str", "name": "Str", "variation_type": "string",
		  "variations": [{"value": "a"}, {"value": "b"}]}`,
		`{"project_key": "proj", "key": "str", "name": "Str", "variation_type": "string",
		  "variations": [{"value": "a"}, {"value": "c"}]}`},
	"flag-number": {tfFeatureFlag,
		`{"project_key": "proj", "key": "num", "name": "Num", "variation_type": "number",
		  "variations": [{"value": "1"}, {"value": "2.5"}]}`,
		`{"project_key": "proj", "key": "num", "name": "Num 2", "variation_type": "number",
		  "variations": [{"value": "1"}, {"value": "2.5"}]}`},
	"flag-json": {tfFeatureFlag,
		`{"project_key": "proj", "key": "json", "name": "JSON", "variation_type": "json",
		  "variations": [{"value": "{\"a\":1}"}, {"value": "{\"a\":2}"}],
		  "client_side_availability": [{"using_environment_id": true, "using_mobile_key": false}]}`,
		`{"project_key": "proj", "key": "json", "name": "JSON", "variation_type": "json",
		  "variations": [{"value": "{\"a\":1}"}, {"value": "{\"a\":3}"}],
		  "client_side_availability": [{"using_environment_id": false, "using_mobile_key": true}]}`},
	"flag-environment": {"launchdarkly_feature_flag_environment",
		`{"flag_id": "proj/bool", "env_key": "test", "off_variation": 1, "on": true, "fallthrough": [{"variation": 0}]}`,
		`{"flag_id": "proj/bool", "env_key": "test", "off_variation": 1, "on": false, "fallthrough": [{"variation": 1}]}`},
	"segment": {"launchdarkly_segment",
		`{"project_key": "proj", "env_key": "test", "key": "seg", "name": "Seg", "included": ["u1"]}`,
		`{"project_key": "proj", "env_key": "test", "key": "seg", "name": "Seg", "included": ["u1", "u2"]}`},
	"access-token": {"launchdarkly_access_token",
		`{"name": "token", "role": "reader"}`,
		`{"name": "token 2", "role": "reader"}`},
	"custom-role": {"launchdarkly_custom_role",
		`{"key": "role", "name": "Role", "policy_statements": [{"effect": "allow", "actions": ["*"], "resources": ["proj/*"]}]}`,
		`{"key": "role", "name": "Role 2", "policy_statements": [{"effect": "allow", "actions": ["*"], "resources": ["proj/*"]}]}`},
	"webhook": {"launchdarkly_webhook",
		`{"url": "https://example.com/hook", "name": "hook", "on": true}`,
		`{"url": "https://example.com/hook", "name": "hook 2", "on": false}`},
	"team-member": {"launchdarkly_team_member",
		`{"email": "a@example.com", "role": "reader"}`,
		`{"email": "a@example.com", "role": "writer"}`},
	"team": {"launchdarkly_team",
		`{"key": "team", "name": "Team"}`,
		`{"key": "team", "name": "Team 2", "description": "after"}`},
	"relay-proxy": {"launchdarkly_relay_proxy_configuration",
		`{"name": "relay", "policy": [{"effect": "allow", "actions": ["*"], "resources": ["proj/*:env/*"]}]}`,
		`{"name": "relay 2", "policy": [{"effect": "allow", "actions": ["*"], "resources": ["proj/*:env/*"]}]}`},
}

// Cases that currently fail because of a known bug. They are skipped, and the
// test fails once they pass so the entry gets removed.
var crudKnownBugs = map[string]string{
	"project-minimal": "upstream customizeProjectDiff panics (\"value is null\") when default_client_side_availability is omitted",
}

// Resources without a CRUD smoke case. Add a case to crudCases instead of
// growing this list where possible.
var crudSkipped = map[string]string{
	"launchdarkly_audit_log_subscription": "needs third-party integration config",
	"launchdarkly_destination":            "needs third-party integration config",
	"launchdarkly_metric":                 "not covered yet",
}

func TestMain(m *testing.M) {
	var srv *httptest.Server
	fakeLD, srv = newFakeLDAPI()
	fakeLDHost = strings.TrimPrefix(srv.URL, "https://")
	// The provider builds its own HTTP client, so trust the fake API's
	// certificate via the system pool before anything loads it.
	f, err := os.CreateTemp("", "fake-ld-ca-*.pem")
	if err != nil {
		panic(err)
	}
	_ = pem.Encode(f, &pem.Block{Type: "CERTIFICATE", Bytes: srv.Certificate().Raw})
	_ = f.Close()
	_ = os.Setenv("SSL_CERT_FILE", f.Name())
	code := m.Run()
	srv.Close()
	_ = os.Remove(f.Name())
	os.Exit(code)
}

var (
	fakeLD     *fakeLDAPI
	fakeLDHost string
)

const tfFeatureFlag = "launchdarkly_feature_flag"

// TestCRUDSmoke runs each resource's Terraform create, read, update and delete
// against an in-memory fake of the LD API, using the runtime schema the
// provider runs with in no-fork mode. It catches panics and schema/config
// mismatches without a cluster or LD credentials.
func TestCRUDSmoke(t *testing.T) {
	p := ldProvider.Provider()
	if d := p.Configure(context.Background(), terraform.NewResourceConfigRaw(map[string]any{
		"access_token": "fake", "api_host": fakeLDHost,
	})); d.HasError() {
		t.Fatalf("configure provider: %v", d)
	}
	resources := GetProvider().Resources
	for name, c := range crudCases {
		t.Run(name, func(t *testing.T) {
			fakeLD.reset()
			r := resources[c.resource]
			if r == nil {
				t.Fatalf("unknown resource %s", c.resource)
			}
			err := runCRUD(r.TerraformResource, p.Meta(), c)
			switch bug := crudKnownBugs[name]; {
			case bug != "" && err == nil:
				t.Fatalf("known bug is fixed, remove %q from crudKnownBugs", name)
			case bug != "":
				t.Skipf("known bug: %s\n%v", bug, err)
			case err != nil:
				t.Fatal(err)
			}
		})
	}
}

func TestCRUDSmokeCoversAllResources(t *testing.T) {
	covered := map[string]bool{}
	for _, c := range crudCases {
		covered[c.resource] = true
	}
	for name, r := range GetProvider().Resources {
		if r.TerraformResource == nil || PluginFrameworkResources[name] {
			continue
		}
		if !covered[name] && crudSkipped[name] == "" {
			t.Errorf("%s has no CRUD smoke case: add one to crudCases", name)
		}
	}
}

// runCRUD returns the first failing step's error, or a panic as an error.
func runCRUD(r *schema.Resource, meta any, c crudCase) error {
	ctx := context.Background()
	var state *terraform.InstanceState
	step := func(op string, fn func() error) (err error) {
		defer func() {
			if v := recover(); v != nil {
				err = fmt.Errorf("%s panicked: %v\n%s", op, v, debug.Stack())
			}
		}()
		if err := fn(); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		return nil
	}
	apply := func(op, cfgJSON string) error {
		return step(op, func() error {
			var cfg map[string]any
			if err := json.Unmarshal([]byte(cfgJSON), &cfg); err != nil {
				return fmt.Errorf("bad test config: %w", err)
			}
			rc := terraform.NewResourceConfigRaw(cfg)
			if d := r.Validate(rc); d.HasError() {
				return fmt.Errorf("validate: %v", d)
			}
			// Same calls as upjet's no-fork external client.
			rawConfig, err := schema.JSONMapToStateValue(cfg, r.CoreConfigSchema())
			if err != nil {
				return fmt.Errorf("raw config: %w", err)
			}
			if state == nil {
				state = &terraform.InstanceState{}
			}
			state.RawConfig = rawConfig
			diff, err := schema.InternalMap(r.Schema).Diff(ctx, state, rc, r.CustomizeDiff, meta, false)
			if err != nil {
				return fmt.Errorf("diff: %w", err)
			}
			if diff == nil || diff.Empty() {
				return nil
			}
			if diff.RawPlan, err = diff.ApplyToValue(cty.EmptyObjectVal, r.CoreConfigSchema()); err != nil {
				return fmt.Errorf("plan: %w", err)
			}
			diff.RawConfig = rawConfig
			s, d := r.Apply(ctx, state, diff, meta)
			if d.HasError() {
				return fmt.Errorf("apply: %v", d)
			}
			if s == nil || s.ID == "" {
				return fmt.Errorf("resource has no ID")
			}
			state = s
			return nil
		})
	}
	read := func() error {
		return step("read", func() error {
			s, d := r.RefreshWithoutUpgrade(ctx, state, meta)
			if d.HasError() {
				return fmt.Errorf("%v", d)
			}
			if s == nil || s.ID == "" {
				return fmt.Errorf("resource disappeared on read")
			}
			state = s
			return nil
		})
	}
	del := func() error {
		return step("delete", func() error {
			if _, d := r.Apply(ctx, state, &terraform.InstanceDiff{Destroy: true}, meta); d.HasError() {
				return fmt.Errorf("%v", d)
			}
			return nil
		})
	}
	for _, fn := range []func() error{
		func() error { return apply("create", c.create) },
		read,
		func() error { return apply("update", c.update) },
		del,
	} {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}
