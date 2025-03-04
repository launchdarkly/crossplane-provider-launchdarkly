package project

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_project", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.Kind = "Project"
	})
}
