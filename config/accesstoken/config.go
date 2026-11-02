// Package accesstoken contains the configuration for the AccessToken resource.
package accesstoken

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_access_token", func(r *config.Resource) {
		r.Kind = "AccessToken"
		r.ShortGroup = "account"

		r.References["custom_roles"] = config.Reference{
			TerraformName: "launchdarkly_custom_role",
		}

		// Hide deprecated fields from spec. Keep them in the schema: in no-fork mode the
		// Terraform provider code runs against this schema and reads these fields.
		config.MoveToStatus(r.TerraformResource, "expire", "policy_statements")
	})
}
