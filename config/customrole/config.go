// Package customrole contains the configuration for the CustomRole resource.
package customrole

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_custom_role", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "CustomRole"

		// Remove deprecated fields
		delete(r.TerraformResource.Schema, "policy")
	})
}
