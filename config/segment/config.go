package segment

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_segment", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.Kind = "EnvironmentSegment"
		r.References["env_key"] = config.Reference{
			TerraformName: "launchdarkly_environment",

			// See: https://github.com/crossplane/upjet/blob/main/pkg/resource/extractor.go
			Extractor: `github.com/crossplane/upjet/pkg/resource.ExtractParamPath("key", false)`,
		}

		r.References["project_key"] = config.Reference{
			TerraformName: "launchdarkly_project",
		}
	})
}
