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

		// Remove deprecated fields
		delete(r.TerraformResource.Schema, "expire")
		delete(r.TerraformResource.Schema, "policy_statements")
	})
}
