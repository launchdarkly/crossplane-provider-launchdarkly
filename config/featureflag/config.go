package featureflag

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_feature_flag", func(r *config.Resource) {
		r.ShortGroup = "flag"
		r.Kind = "FeatureFlag"
		r.References["project"] = config.Reference{
			TerraformName: "launchdarkly_project",
		}
	})
}
