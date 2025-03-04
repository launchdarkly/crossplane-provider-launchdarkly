package featureflagenvironment

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_feature_flag_environment", func(r *config.Resource) {
		r.ShortGroup = "flag"
		r.Kind = "FeatureFlagEnvironment"
		r.References["environment"] = config.Reference{
			TerraformName: "launchdarkly_environment",
		}
	})
}
