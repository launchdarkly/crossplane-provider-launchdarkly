// Package featureflagenvironment contains the configuration for the FeatureFlagEnvironment resource.
package featureflagenvironment

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_feature_flag_environment", func(r *config.Resource) {
		r.ShortGroup = "flag"
		r.Kind = "FeatureFlagEnvironment"
		r.References["env_key"] = config.Reference{
			TerraformName: "launchdarkly_environment",

			// See: https://github.com/crossplane/upjet/blob/main/pkg/resource/extractor.go
			Extractor: `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("key", false)`,
		}

		r.References["flag_id"] = config.Reference{
			TerraformName: "launchdarkly_feature_flag",
		}
	})
}
