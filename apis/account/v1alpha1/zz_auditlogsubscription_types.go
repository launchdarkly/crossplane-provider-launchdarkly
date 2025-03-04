// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AuditLogSubscriptionInitParameters struct {

	// (Map of String) The set of configuration fields corresponding to the value defined for integration_key. Refer to the formVariables field in the corresponding integrations/<integration_key>/manifest.json file in this repo for a full list of fields for the integration you wish to configure.
	// The set of configuration fields corresponding to the value defined for `integration_key`. Refer to the `formVariables` field in the corresponding `integrations/<integration_key>/manifest.json` file in [this repo](https://github.com/launchdarkly/integration-framework/tree/master/integrations) for a full list of fields for the integration you wish to configure.
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// relic-apm, signalfx, slack, and splunk. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	IntegrationKey *string `json:"integrationKey,omitempty" tf:"integration_key,omitempty"`

	// friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
	// A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether or not you want your subscription enabled, i.e. to actively send events.
	// Whether or not you want your subscription enabled, i.e. to actively send events.
	On *bool `json:"on,omitempty" tf:"on,omitempty"`

	// (Block List, Min: 1) A block representing the resources to which you wish to subscribe. (see below for nested schema)
	// A block representing the resources to which you wish to subscribe.
	Statements []StatementsInitParameters `json:"statements,omitempty" tf:"statements,omitempty"`

	// (Set of String) Tags associated with your resource.
	// Tags associated with your resource.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AuditLogSubscriptionObservation struct {

	// (Map of String) The set of configuration fields corresponding to the value defined for integration_key. Refer to the formVariables field in the corresponding integrations/<integration_key>/manifest.json file in this repo for a full list of fields for the integration you wish to configure.
	// The set of configuration fields corresponding to the value defined for `integration_key`. Refer to the `formVariables` field in the corresponding `integrations/<integration_key>/manifest.json` file in [this repo](https://github.com/launchdarkly/integration-framework/tree/master/integrations) for a full list of fields for the integration you wish to configure.
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// relic-apm, signalfx, slack, and splunk. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	IntegrationKey *string `json:"integrationKey,omitempty" tf:"integration_key,omitempty"`

	// friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
	// A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether or not you want your subscription enabled, i.e. to actively send events.
	// Whether or not you want your subscription enabled, i.e. to actively send events.
	On *bool `json:"on,omitempty" tf:"on,omitempty"`

	// (Block List, Min: 1) A block representing the resources to which you wish to subscribe. (see below for nested schema)
	// A block representing the resources to which you wish to subscribe.
	Statements []StatementsObservation `json:"statements,omitempty" tf:"statements,omitempty"`

	// (Set of String) Tags associated with your resource.
	// Tags associated with your resource.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AuditLogSubscriptionParameters struct {

	// (Map of String) The set of configuration fields corresponding to the value defined for integration_key. Refer to the formVariables field in the corresponding integrations/<integration_key>/manifest.json file in this repo for a full list of fields for the integration you wish to configure.
	// The set of configuration fields corresponding to the value defined for `integration_key`. Refer to the `formVariables` field in the corresponding `integrations/<integration_key>/manifest.json` file in [this repo](https://github.com/launchdarkly/integration-framework/tree/master/integrations) for a full list of fields for the integration you wish to configure.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// relic-apm, signalfx, slack, and splunk. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The integration key. Supported integration keys are `chronosphere`, `cloudtrail`, `datadog`, `dynatrace`, `elastic`, `grafana`, `honeycomb`, `kosli`, `last9`, `logdna`, `msteams`, `new-relic-apm`, `signalfx`, `slack`, and `splunk`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// +kubebuilder:validation:Optional
	IntegrationKey *string `json:"integrationKey,omitempty" tf:"integration_key,omitempty"`

	// friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
	// A human-friendly name for your audit log subscription viewable from within the LaunchDarkly Integrations page.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether or not you want your subscription enabled, i.e. to actively send events.
	// Whether or not you want your subscription enabled, i.e. to actively send events.
	// +kubebuilder:validation:Optional
	On *bool `json:"on,omitempty" tf:"on,omitempty"`

	// (Block List, Min: 1) A block representing the resources to which you wish to subscribe. (see below for nested schema)
	// A block representing the resources to which you wish to subscribe.
	// +kubebuilder:validation:Optional
	Statements []StatementsParameters `json:"statements,omitempty" tf:"statements,omitempty"`

	// (Set of String) Tags associated with your resource.
	// Tags associated with your resource.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StatementsInitParameters struct {

	// (List of String) The list of action specifiers defining the actions to which the statement applies.
	// Either actions or not_actions must be specified. For a list of available actions read Actions reference.
	// The list of action specifiers defining the actions to which the statement applies.
	// Either `actions` or `not_actions` must be specified. For a list of available actions read [Actions reference](https://docs.launchdarkly.com/home/account-security/custom-roles/actions#actions-reference).
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// (String) Either allow or deny. This argument defines whether the statement allows or denies access to the named resources and actions.
	// Either `allow` or `deny`. This argument defines whether the statement allows or denies access to the named resources and actions.
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// (List of String) The list of action specifiers defining the actions to which the statement does not apply.
	// The list of action specifiers defining the actions to which the statement does not apply.
	NotActions []*string `json:"notActions,omitempty" tf:"not_actions,omitempty"`

	// (List of String) The list of resource specifiers defining the resources to which the statement does not apply.
	// The list of resource specifiers defining the resources to which the statement does not apply.
	NotResources []*string `json:"notResources,omitempty" tf:"not_resources,omitempty"`

	// (List of String) The list of resource specifiers defining the resources to which the statement applies.
	// The list of resource specifiers defining the resources to which the statement applies.
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`
}

type StatementsObservation struct {

	// (List of String) The list of action specifiers defining the actions to which the statement applies.
	// Either actions or not_actions must be specified. For a list of available actions read Actions reference.
	// The list of action specifiers defining the actions to which the statement applies.
	// Either `actions` or `not_actions` must be specified. For a list of available actions read [Actions reference](https://docs.launchdarkly.com/home/account-security/custom-roles/actions#actions-reference).
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// (String) Either allow or deny. This argument defines whether the statement allows or denies access to the named resources and actions.
	// Either `allow` or `deny`. This argument defines whether the statement allows or denies access to the named resources and actions.
	Effect *string `json:"effect,omitempty" tf:"effect,omitempty"`

	// (List of String) The list of action specifiers defining the actions to which the statement does not apply.
	// The list of action specifiers defining the actions to which the statement does not apply.
	NotActions []*string `json:"notActions,omitempty" tf:"not_actions,omitempty"`

	// (List of String) The list of resource specifiers defining the resources to which the statement does not apply.
	// The list of resource specifiers defining the resources to which the statement does not apply.
	NotResources []*string `json:"notResources,omitempty" tf:"not_resources,omitempty"`

	// (List of String) The list of resource specifiers defining the resources to which the statement applies.
	// The list of resource specifiers defining the resources to which the statement applies.
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`
}

type StatementsParameters struct {

	// (List of String) The list of action specifiers defining the actions to which the statement applies.
	// Either actions or not_actions must be specified. For a list of available actions read Actions reference.
	// The list of action specifiers defining the actions to which the statement applies.
	// Either `actions` or `not_actions` must be specified. For a list of available actions read [Actions reference](https://docs.launchdarkly.com/home/account-security/custom-roles/actions#actions-reference).
	// +kubebuilder:validation:Optional
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// (String) Either allow or deny. This argument defines whether the statement allows or denies access to the named resources and actions.
	// Either `allow` or `deny`. This argument defines whether the statement allows or denies access to the named resources and actions.
	// +kubebuilder:validation:Optional
	Effect *string `json:"effect" tf:"effect,omitempty"`

	// (List of String) The list of action specifiers defining the actions to which the statement does not apply.
	// The list of action specifiers defining the actions to which the statement does not apply.
	// +kubebuilder:validation:Optional
	NotActions []*string `json:"notActions,omitempty" tf:"not_actions,omitempty"`

	// (List of String) The list of resource specifiers defining the resources to which the statement does not apply.
	// The list of resource specifiers defining the resources to which the statement does not apply.
	// +kubebuilder:validation:Optional
	NotResources []*string `json:"notResources,omitempty" tf:"not_resources,omitempty"`

	// (List of String) The list of resource specifiers defining the resources to which the statement applies.
	// The list of resource specifiers defining the resources to which the statement applies.
	// +kubebuilder:validation:Optional
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`
}

// AuditLogSubscriptionSpec defines the desired state of AuditLogSubscription
type AuditLogSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AuditLogSubscriptionParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AuditLogSubscriptionInitParameters `json:"initProvider,omitempty"`
}

// AuditLogSubscriptionStatus defines the observed state of AuditLogSubscription.
type AuditLogSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AuditLogSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AuditLogSubscription is the Schema for the AuditLogSubscriptions API. Provides a LaunchDarkly audit log subscription resource. This resource allows you to create and manage LaunchDarkly audit log subscriptions.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,launchdarkly}
type AuditLogSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.config) || (has(self.initProvider) && has(self.initProvider.config))",message="spec.forProvider.config is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.integrationKey) || (has(self.initProvider) && has(self.initProvider.integrationKey))",message="spec.forProvider.integrationKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.on) || (has(self.initProvider) && has(self.initProvider.on))",message="spec.forProvider.on is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.statements) || (has(self.initProvider) && has(self.initProvider.statements))",message="spec.forProvider.statements is a required parameter"
	Spec   AuditLogSubscriptionSpec   `json:"spec"`
	Status AuditLogSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuditLogSubscriptionList contains a list of AuditLogSubscriptions
type AuditLogSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AuditLogSubscription `json:"items"`
}

// Repository type metadata.
var (
	AuditLogSubscription_Kind             = "AuditLogSubscription"
	AuditLogSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuditLogSubscription_Kind}.String()
	AuditLogSubscription_KindAPIVersion   = AuditLogSubscription_Kind + "." + CRDGroupVersion.String()
	AuditLogSubscription_GroupVersionKind = CRDGroupVersion.WithKind(AuditLogSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&AuditLogSubscription{}, &AuditLogSubscriptionList{})
}
