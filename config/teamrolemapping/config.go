// Package teamrolemapping contains the configuration for the TeamRoleMapping resource.
package teamrolemapping

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_team_role_mapping", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "TeamRoleMapping"
		r.References["team_key"] = config.Reference{
			TerraformName: "launchdarkly_team",
		}

		r.References["custom_role_keys"] = config.Reference{
			TerraformName: "launchdarkly_custom_role",
		}
	})
}
