// Package relayproxyconfig contains the configuration for the RelayProxyConfiguration resource.
package relayproxyconfig

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_relay_proxy_configuration", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "RelayProxyConfiguration"
	})
}
