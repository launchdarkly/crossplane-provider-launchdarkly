package auditlogsubscription

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_audit_log_subscription", func(r *config.Resource) {
		r.ShortGroup = "account"
		r.Kind = "AuditLogSubscription"
	})
}
