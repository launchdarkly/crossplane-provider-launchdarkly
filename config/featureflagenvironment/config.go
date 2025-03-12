package featureflagenvironment

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/extractors"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_feature_flag_environment", func(r *config.Resource) {
		r.ShortGroup = "flag"
		r.Kind = "FeatureFlagEnvironment"
		r.References["env_key"] = config.Reference{
			TerraformName: "launchdarkly_environment",
			Extractor:     extractors.FieldExtractorFnReference("key"),
		}

		r.References["flag_id"] = config.Reference{
			TerraformName: "launchdarkly_feature_flag",
		}
	})
}
