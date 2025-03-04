package teamrolemapping

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_team_role_mapping", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "TeamRoleMapping"
		r.References["team"] = config.Reference{
			TerraformName: "launchdarkly_team",
		}
	})
}
