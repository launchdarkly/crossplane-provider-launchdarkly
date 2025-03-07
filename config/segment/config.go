package segment

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/extractors"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_segment", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.Kind = "EnvironmentSegment"
		r.References["env_key"] = config.Reference{
			TerraformName: "launchdarkly_environment",
			Extractor:     extractors.FieldExtractorFnReference("key"),
		}
	})
}
