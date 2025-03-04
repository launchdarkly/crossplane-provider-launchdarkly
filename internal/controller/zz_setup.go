// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accesstoken "github.com/launchdarkly/crossplane-provider-launchdarkly/internal/controller/account/accesstoken"
	auditlogsubscription "github.com/launchdarkly/crossplane-provider-launchdarkly/internal/controller/account/auditlogsubscription"
	customrole "github.com/launchdarkly/crossplane-provider-launchdarkly/internal/controller/account/customrole"
	team "github.com/launchdarkly/crossplane-provider-launchdarkly/internal/controller/account/team"
	teamrolemapping "github.com/launchdarkly/crossplane-provider-launchdarkly/internal/controller/account/teamrolemapping"
	user "github.com/launchdarkly/crossplane-provider-launchdarkly/internal/controller/account/user"
	webhook "github.com/launchdarkly/crossplane-provider-launchdarkly/internal/controller/account/webhook"
	featureflag "github.com/launchdarkly/crossplane-provider-launchdarkly/internal/controller/flag/featureflag"
	featureflagenvironment "github.com/launchdarkly/crossplane-provider-launchdarkly/internal/controller/flag/featureflagenvironment"
	environment "github.com/launchdarkly/crossplane-provider-launchdarkly/internal/controller/project/environment"
	environmentdestination "github.com/launchdarkly/crossplane-provider-launchdarkly/internal/controller/project/environmentdestination"
	environmentsegment "github.com/launchdarkly/crossplane-provider-launchdarkly/internal/controller/project/environmentsegment"
	project "github.com/launchdarkly/crossplane-provider-launchdarkly/internal/controller/project/project"
	providerconfig "github.com/launchdarkly/crossplane-provider-launchdarkly/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesstoken.Setup,
		auditlogsubscription.Setup,
		customrole.Setup,
		team.Setup,
		teamrolemapping.Setup,
		user.Setup,
		webhook.Setup,
		featureflag.Setup,
		featureflagenvironment.Setup,
		environment.Setup,
		environmentdestination.Setup,
		environmentsegment.Setup,
		project.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
