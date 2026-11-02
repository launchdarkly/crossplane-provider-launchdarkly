// Package customrole contains the configuration for the CustomRole resource.
package customrole

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_custom_role", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "CustomRole"

		// Hide deprecated fields from spec. Keep them in the schema: in no-fork mode the
		// Terraform provider code runs against this schema and reads these fields.
		config.MoveToStatus(r.TerraformResource, "policy")
	})
}
