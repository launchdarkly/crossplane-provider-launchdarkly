/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/accesstoken"
	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/auditlogsubscription"
	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/customrole"
	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/destination"
	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/environment"
	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/featureflag"
	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/featureflagenvironment"
	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/project"
	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/segment"
	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/team"
	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/teammember"
	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/teamrolemapping"
	"github.com/launchdarkly/crossplane-provider-launchdarkly/config/webhook"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "launchdarkly"
	modulePath     = "github.com/launchdarkly/crossplane-provider-launchdarkly"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("launchdarkly.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		accesstoken.Configure,
		auditlogsubscription.Configure,
		customrole.Configure,
		destination.Configure,
		environment.Configure,
		featureflag.Configure,
		featureflagenvironment.Configure,
		project.Configure,
		segment.Configure,
		team.Configure,
		teammember.Configure,
		teamrolemapping.Configure,
		webhook.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
