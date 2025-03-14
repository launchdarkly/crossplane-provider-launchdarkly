package relayproxyconfig

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_relay_proxy_configuration", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "RelayProxyConfiguration"
	})
}
