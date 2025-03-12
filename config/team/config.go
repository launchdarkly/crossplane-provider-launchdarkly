package team

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_team", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "Team"

		r.References["member_ids"] = config.Reference{
			TerraformName: "launchdarkly_team_member",
		}

		r.References["maintainers"] = config.Reference{
			TerraformName: "launchdarkly_team_member",
		}

		r.References["custom_role_keys"] = config.Reference{
			TerraformName: "launchdarkly_custom_role",
		}
	})
}
