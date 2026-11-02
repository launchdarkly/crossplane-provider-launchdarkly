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

		// Hide deprecated fields from spec. Keep them in the schema: in no-fork mode the
		// Terraform provider code runs against this schema and reads these fields.
		config.MoveToStatus(r.TerraformResource, "include_in_snippet")
	})
}
