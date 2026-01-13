// Package featureflag contains the configuration for the FeatureFlag resource.
package featureflag

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_feature_flag", func(r *config.Resource) {
		r.ShortGroup = "flag"
		r.Kind = "FeatureFlag"
		r.References["project_key"] = config.Reference{
			TerraformName: "launchdarkly_project",
		}

		// Remove deprecated fields
		delete(r.TerraformResource.Schema, "include_in_snippet")
	})
}
