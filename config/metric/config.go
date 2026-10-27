// Package metric contains the configuration for the Metric resource.
package metric

import "github.com/crossplane/upjet/v2/pkg/config"

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
		r.TerraformResource.CustomizeDiff = customizeDiff
		// customizeDiff validates spec.forProvider as the raw config, so a value
		// late-initialized from LD's response would be checked as if the user had
		// set it. For example, include_units_without_events=true (the default for
		// mean metrics) would then reject a switch to percentile.
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"event_key",
				"include_units_without_events",
				"percentile_value",
				"selector",
				"success_criteria",
				"unit",
			},
		}
	})
}
