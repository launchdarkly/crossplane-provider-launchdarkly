package environment

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_environment", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.Kind = "Environment"
		r.References["project_key"] = config.Reference{
			TerraformName: "launchdarkly_project",
		}
	})
}
