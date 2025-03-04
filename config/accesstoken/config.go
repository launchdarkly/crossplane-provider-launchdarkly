package accesstoken

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_access_token", func(r *config.Resource) {
		r.Kind = "AccessToken"
		r.ShortGroup = "account"
	})
}
