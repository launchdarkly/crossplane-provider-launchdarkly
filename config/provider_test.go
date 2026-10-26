package config

import (
	"testing"

	ldProvider "github.com/launchdarkly/terraform-provider-launchdarkly/launchdarkly"
)

// In no-fork mode the Terraform provider's CRUD functions run against
// r.TerraformResource, so removing a field there makes the upstream code panic
// when it reads that field (e.g. d.Get("include_in_snippet").(bool)).
func TestRuntimeSchemaKeepsAllTerraformFields(t *testing.T) {
	upstream := ldProvider.Provider().ResourcesMap
	for name, r := range GetProvider().Resources {
		u, ok := upstream[name]
		if !ok || r.TerraformResource == nil {
			continue
		}
		for field := range u.Schema {
			if _, ok := r.TerraformResource.Schema[field]; !ok {
				t.Errorf("%s: field %q was removed from the runtime Terraform schema", name, field)
			}
		}
	}
}

func TestDeprecatedFieldsAreNotInSpec(t *testing.T) {
	deprecated := map[string][]string{
		"launchdarkly_feature_flag": {"include_in_snippet"},
		"launchdarkly_project":      {"include_in_snippet"},
		"launchdarkly_access_token": {"expire", "policy_statements"},
		"launchdarkly_custom_role":  {"policy"},
	}
	resources := GetProvider().Resources
	for name, fields := range deprecated {
		for _, field := range fields {
			s, ok := resources[name].TerraformResource.Schema[field]
			if !ok {
				t.Errorf("%s: field %q missing from the runtime Terraform schema", name, field)
				continue
			}
			if s.Optional || s.Required {
				t.Errorf("%s: deprecated field %q is still settable in spec", name, field)
			}
		}
	}
}
