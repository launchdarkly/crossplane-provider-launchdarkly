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

type EnvironmentDestinationInitParameters struct {

	// specific configuration. To learn more, read Destination-Specific Configs
	// The destination-specific configuration. To learn more, read [Destination-Specific Configs](#destination-specific-configs)
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// (String) The environment key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The environment key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	EnvKey *string `json:"envKey,omitempty" tf:"env_key,omitempty"`

	// pubsub, mparticle, azure-event-hubs, and segment. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The data export destination type. Available choices are `kinesis`, `google-pubsub`, `mparticle`, `azure-event-hubs`, and `segment`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// readable name for your data export destination.
	// A human-readable name for your data export destination.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether the data export destination is on or not.
	// Whether the data export destination is on or not.
	On *bool `json:"on,omitempty" tf:"on,omitempty"`

	// (String) The LaunchDarkly project key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The LaunchDarkly project key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey *string `json:"projectKey,omitempty" tf:"project_key,omitempty"`

	// (Set of String) Tags associated with your resource.
	// Tags associated with your resource.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EnvironmentDestinationObservation struct {

	// specific configuration. To learn more, read Destination-Specific Configs
	// The destination-specific configuration. To learn more, read [Destination-Specific Configs](#destination-specific-configs)
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// (String) The environment key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The environment key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	EnvKey *string `json:"envKey,omitempty" tf:"env_key,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// pubsub, mparticle, azure-event-hubs, and segment. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The data export destination type. Available choices are `kinesis`, `google-pubsub`, `mparticle`, `azure-event-hubs`, and `segment`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// readable name for your data export destination.
	// A human-readable name for your data export destination.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether the data export destination is on or not.
	// Whether the data export destination is on or not.
	On *bool `json:"on,omitempty" tf:"on,omitempty"`

	// (String) The LaunchDarkly project key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The LaunchDarkly project key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	ProjectKey *string `json:"projectKey,omitempty" tf:"project_key,omitempty"`

	// (Set of String) Tags associated with your resource.
	// Tags associated with your resource.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EnvironmentDestinationParameters struct {

	// specific configuration. To learn more, read Destination-Specific Configs
	// The destination-specific configuration. To learn more, read [Destination-Specific Configs](#destination-specific-configs)
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// (String) The environment key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The environment key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// +kubebuilder:validation:Optional
	EnvKey *string `json:"envKey,omitempty" tf:"env_key,omitempty"`

	// pubsub, mparticle, azure-event-hubs, and segment. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The data export destination type. Available choices are `kinesis`, `google-pubsub`, `mparticle`, `azure-event-hubs`, and `segment`. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// readable name for your data export destination.
	// A human-readable name for your data export destination.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Whether the data export destination is on or not.
	// Whether the data export destination is on or not.
	// +kubebuilder:validation:Optional
	On *bool `json:"on,omitempty" tf:"on,omitempty"`

	// (String) The LaunchDarkly project key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The LaunchDarkly project key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// +kubebuilder:validation:Optional
	ProjectKey *string `json:"projectKey,omitempty" tf:"project_key,omitempty"`

	// (Set of String) Tags associated with your resource.
	// Tags associated with your resource.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// EnvironmentDestinationSpec defines the desired state of EnvironmentDestination
type EnvironmentDestinationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentDestinationParameters `json:"forProvider"`
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
	InitProvider EnvironmentDestinationInitParameters `json:"initProvider,omitempty"`
}

// EnvironmentDestinationStatus defines the observed state of EnvironmentDestination.
type EnvironmentDestinationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentDestinationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EnvironmentDestination is the Schema for the EnvironmentDestinations API. Provides a LaunchDarkly Data Export Destination resource. -> Note: Data Export is available to customers on an Enterprise LaunchDarkly plan. To learn more, read about our pricing https://launchdarkly.com/pricing/. To upgrade your plan, contact LaunchDarkly Sales https://launchdarkly.com/contact-sales/. Data Export Destinations are locations that receive exported data. This resource allows you to configure destinations for the export of raw analytics data, including feature flag requests, analytics events, custom events, and more. To learn more about data export, read Data Export Documentation https://docs.launchdarkly.com/integrations/data-export.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,launchdarkly}
type EnvironmentDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.config) || (has(self.initProvider) && has(self.initProvider.config))",message="spec.forProvider.config is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.envKey) || (has(self.initProvider) && has(self.initProvider.envKey))",message="spec.forProvider.envKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.kind) || (has(self.initProvider) && has(self.initProvider.kind))",message="spec.forProvider.kind is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.projectKey) || (has(self.initProvider) && has(self.initProvider.projectKey))",message="spec.forProvider.projectKey is a required parameter"
	Spec   EnvironmentDestinationSpec   `json:"spec"`
	Status EnvironmentDestinationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentDestinationList contains a list of EnvironmentDestinations
type EnvironmentDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnvironmentDestination `json:"items"`
}

// Repository type metadata.
var (
	EnvironmentDestination_Kind             = "EnvironmentDestination"
	EnvironmentDestination_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EnvironmentDestination_Kind}.String()
	EnvironmentDestination_KindAPIVersion   = EnvironmentDestination_Kind + "." + CRDGroupVersion.String()
	EnvironmentDestination_GroupVersionKind = CRDGroupVersion.WithKind(EnvironmentDestination_Kind)
)

func init() {
	SchemeBuilder.Register(&EnvironmentDestination{}, &EnvironmentDestinationList{})
}
