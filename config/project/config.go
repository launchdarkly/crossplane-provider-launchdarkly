// Package project contains the configuration for the Project resource.
package project

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("launchdarkly_project", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.Kind = "Project"

		// Hide deprecated fields from spec. Keep them in the schema: in no-fork mode the
		// Terraform provider code runs against this schema and reads these fields.
		config.MoveToStatus(r.TerraformResource, "include_in_snippet")
	})
}
