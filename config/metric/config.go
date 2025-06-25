package metric

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_metric", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.Kind = "Metric"
		r.References["project_key"] = config.Reference{
			TerraformName: "launchdarkly_project",
		}
		r.References["maintainer_id"] = config.Reference{
			TerraformName: "launchdarkly_team_member",
		}
	})
}
