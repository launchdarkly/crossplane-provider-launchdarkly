// Package project contains the configuration for the Project resource.
package project

import (
	"context"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const defaultClientSideAvailability = "default_client_side_availability"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_project", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.Kind = "Project"

		// Hide deprecated fields from spec. Keep them in the schema: in no-fork mode the
		// Terraform provider code runs against this schema and reads these fields.
		config.MoveToStatus(r.TerraformResource, "include_in_snippet")

		r.TerraformResource.CustomizeDiff = customizeDiff
	})
}

// customizeDiff replaces the upstream CustomizeDiff, which assumes Terraform's
// raw config and panics under upjet: upjet passes an omitted block as null, and
// sometimes passes no raw config at all. Like upstream, it falls back to LD's
// defaults when default_client_side_availability is not set.
func customizeDiff(_ context.Context, diff *schema.ResourceDiff, _ any) error {
	cfg := diff.GetRawConfig()
	if cfg.IsNull() || !cfg.IsKnown() || !cfg.Type().IsObjectType() || !cfg.Type().HasAttribute(defaultClientSideAvailability) {
		return nil
	}
	if csa := cfg.GetAttr(defaultClientSideAvailability); !csa.IsKnown() || (!csa.IsNull() && csa.LengthInt() > 0) {
		return nil
	}
	if err := diff.SetNew("include_in_snippet", false); err != nil {
		return err
	}
	return diff.SetNew(defaultClientSideAvailability, []map[string]any{{
		"using_environment_id": false,
		"using_mobile_key":     true,
	}})
}
