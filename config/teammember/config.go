package teammember

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_team_member", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "TeamMember"
	})
}
