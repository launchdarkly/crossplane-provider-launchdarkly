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

type DefaultClientSideAvailabilityInitParameters struct {

	// (Boolean)
	UsingEnvironmentID *bool `json:"usingEnvironmentId,omitempty" tf:"using_environment_id,omitempty"`

	// (Boolean)
	UsingMobileKey *bool `json:"usingMobileKey,omitempty" tf:"using_mobile_key,omitempty"`
}

type DefaultClientSideAvailabilityObservation struct {

	// (Boolean)
	UsingEnvironmentID *bool `json:"usingEnvironmentId,omitempty" tf:"using_environment_id,omitempty"`

	// (Boolean)
	UsingMobileKey *bool `json:"usingMobileKey,omitempty" tf:"using_mobile_key,omitempty"`
}

type DefaultClientSideAvailabilityParameters struct {

	// (Boolean)
	// +kubebuilder:validation:Optional
	UsingEnvironmentID *bool `json:"usingEnvironmentId" tf:"using_environment_id,omitempty"`

	// (Boolean)
	// +kubebuilder:validation:Optional
	UsingMobileKey *bool `json:"usingMobileKey" tf:"using_mobile_key,omitempty"`
}

type EnvironmentsApprovalSettingsInitParameters struct {

	// (Boolean) Set to true if changes can be applied as long as the min_num_approvals is met, regardless of whether any reviewers have declined a request. Defaults to true.
	// Set to `true` if changes can be applied as long as the `min_num_approvals` is met, regardless of whether any reviewers have declined a request. Defaults to `true`.
	CanApplyDeclinedChanges *bool `json:"canApplyDeclinedChanges,omitempty" tf:"can_apply_declined_changes,omitempty"`

	// (Boolean) Set to true if requesters can approve or decline their own request. They may always comment. Defaults to false.
	// Set to `true` if requesters can approve or decline their own request. They may always comment. Defaults to `false`.
	CanReviewOwnRequest *bool `json:"canReviewOwnRequest,omitempty" tf:"can_review_own_request,omitempty"`

	// (Number) The number of approvals required before an approval request can be applied. This number must be between 1 and 5. Defaults to 1.
	// The number of approvals required before an approval request can be applied. This number must be between 1 and 5. Defaults to 1.
	MinNumApprovals *float64 `json:"minNumApprovals,omitempty" tf:"min_num_approvals,omitempty"`

	// (Boolean) Set to true for changes to flags in this environment to require approval. You may only set required to true if required_approval_tags is not set and vice versa. Defaults to false.
	// Set to `true` for changes to flags in this environment to require approval. You may only set `required` to true if `required_approval_tags` is not set and vice versa. Defaults to `false`.
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	// (List of String) An array of tags used to specify which flags with those tags require approval. You may only set required_approval_tags if required is not set to true and vice versa.
	// An array of tags used to specify which flags with those tags require approval. You may only set `required_approval_tags` if `required` is not set to `true` and vice versa.
	RequiredApprovalTags []*string `json:"requiredApprovalTags,omitempty" tf:"required_approval_tags,omitempty"`

	// (Map of String) The configuration for the service associated with this approval. This is specific to each approval service. For a service_kind of servicenow, the following fields apply:
	// The configuration for the service associated with this approval. This is specific to each approval service. For a `service_kind` of `servicenow`, the following fields apply:
	//
	// - `template` (String) The sys_id of the Standard Change Request Template in ServiceNow that LaunchDarkly will use when creating the change request.
	// - `detail_column` (String) The name of the ServiceNow Change Request column LaunchDarkly uses to populate detailed approval request information.
	// +mapType=granular
	ServiceConfig map[string]*string `json:"serviceConfig,omitempty" tf:"service_config,omitempty"`

	// (String) The kind of service associated with this approval. This determines which platform is used for requesting approval. Valid values are servicenow, launchdarkly.
	// The kind of service associated with this approval. This determines which platform is used for requesting approval. Valid values are `servicenow`, `launchdarkly`.
	ServiceKind *string `json:"serviceKind,omitempty" tf:"service_kind,omitempty"`
}

type EnvironmentsApprovalSettingsObservation struct {

	// (Boolean) Set to true if changes can be applied as long as the min_num_approvals is met, regardless of whether any reviewers have declined a request. Defaults to true.
	// Set to `true` if changes can be applied as long as the `min_num_approvals` is met, regardless of whether any reviewers have declined a request. Defaults to `true`.
	CanApplyDeclinedChanges *bool `json:"canApplyDeclinedChanges,omitempty" tf:"can_apply_declined_changes,omitempty"`

	// (Boolean) Set to true if requesters can approve or decline their own request. They may always comment. Defaults to false.
	// Set to `true` if requesters can approve or decline their own request. They may always comment. Defaults to `false`.
	CanReviewOwnRequest *bool `json:"canReviewOwnRequest,omitempty" tf:"can_review_own_request,omitempty"`

	// (Number) The number of approvals required before an approval request can be applied. This number must be between 1 and 5. Defaults to 1.
	// The number of approvals required before an approval request can be applied. This number must be between 1 and 5. Defaults to 1.
	MinNumApprovals *float64 `json:"minNumApprovals,omitempty" tf:"min_num_approvals,omitempty"`

	// (Boolean) Set to true for changes to flags in this environment to require approval. You may only set required to true if required_approval_tags is not set and vice versa. Defaults to false.
	// Set to `true` for changes to flags in this environment to require approval. You may only set `required` to true if `required_approval_tags` is not set and vice versa. Defaults to `false`.
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	// (List of String) An array of tags used to specify which flags with those tags require approval. You may only set required_approval_tags if required is not set to true and vice versa.
	// An array of tags used to specify which flags with those tags require approval. You may only set `required_approval_tags` if `required` is not set to `true` and vice versa.
	RequiredApprovalTags []*string `json:"requiredApprovalTags,omitempty" tf:"required_approval_tags,omitempty"`

	// (Map of String) The configuration for the service associated with this approval. This is specific to each approval service. For a service_kind of servicenow, the following fields apply:
	// The configuration for the service associated with this approval. This is specific to each approval service. For a `service_kind` of `servicenow`, the following fields apply:
	//
	// - `template` (String) The sys_id of the Standard Change Request Template in ServiceNow that LaunchDarkly will use when creating the change request.
	// - `detail_column` (String) The name of the ServiceNow Change Request column LaunchDarkly uses to populate detailed approval request information.
	// +mapType=granular
	ServiceConfig map[string]*string `json:"serviceConfig,omitempty" tf:"service_config,omitempty"`

	// (String) The kind of service associated with this approval. This determines which platform is used for requesting approval. Valid values are servicenow, launchdarkly.
	// The kind of service associated with this approval. This determines which platform is used for requesting approval. Valid values are `servicenow`, `launchdarkly`.
	ServiceKind *string `json:"serviceKind,omitempty" tf:"service_kind,omitempty"`
}

type EnvironmentsApprovalSettingsParameters struct {

	// (Boolean) Set to true if changes can be applied as long as the min_num_approvals is met, regardless of whether any reviewers have declined a request. Defaults to true.
	// Set to `true` if changes can be applied as long as the `min_num_approvals` is met, regardless of whether any reviewers have declined a request. Defaults to `true`.
	// +kubebuilder:validation:Optional
	CanApplyDeclinedChanges *bool `json:"canApplyDeclinedChanges,omitempty" tf:"can_apply_declined_changes,omitempty"`

	// (Boolean) Set to true if requesters can approve or decline their own request. They may always comment. Defaults to false.
	// Set to `true` if requesters can approve or decline their own request. They may always comment. Defaults to `false`.
	// +kubebuilder:validation:Optional
	CanReviewOwnRequest *bool `json:"canReviewOwnRequest,omitempty" tf:"can_review_own_request,omitempty"`

	// (Number) The number of approvals required before an approval request can be applied. This number must be between 1 and 5. Defaults to 1.
	// The number of approvals required before an approval request can be applied. This number must be between 1 and 5. Defaults to 1.
	// +kubebuilder:validation:Optional
	MinNumApprovals *float64 `json:"minNumApprovals,omitempty" tf:"min_num_approvals,omitempty"`

	// (Boolean) Set to true for changes to flags in this environment to require approval. You may only set required to true if required_approval_tags is not set and vice versa. Defaults to false.
	// Set to `true` for changes to flags in this environment to require approval. You may only set `required` to true if `required_approval_tags` is not set and vice versa. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Required *bool `json:"required,omitempty" tf:"required,omitempty"`

	// (List of String) An array of tags used to specify which flags with those tags require approval. You may only set required_approval_tags if required is not set to true and vice versa.
	// An array of tags used to specify which flags with those tags require approval. You may only set `required_approval_tags` if `required` is not set to `true` and vice versa.
	// +kubebuilder:validation:Optional
	RequiredApprovalTags []*string `json:"requiredApprovalTags,omitempty" tf:"required_approval_tags,omitempty"`

	// (Map of String) The configuration for the service associated with this approval. This is specific to each approval service. For a service_kind of servicenow, the following fields apply:
	// The configuration for the service associated with this approval. This is specific to each approval service. For a `service_kind` of `servicenow`, the following fields apply:
	//
	// - `template` (String) The sys_id of the Standard Change Request Template in ServiceNow that LaunchDarkly will use when creating the change request.
	// - `detail_column` (String) The name of the ServiceNow Change Request column LaunchDarkly uses to populate detailed approval request information.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ServiceConfig map[string]*string `json:"serviceConfig,omitempty" tf:"service_config,omitempty"`

	// (String) The kind of service associated with this approval. This determines which platform is used for requesting approval. Valid values are servicenow, launchdarkly.
	// The kind of service associated with this approval. This determines which platform is used for requesting approval. Valid values are `servicenow`, `launchdarkly`.
	// +kubebuilder:validation:Optional
	ServiceKind *string `json:"serviceKind,omitempty" tf:"service_kind,omitempty"`
}

type EnvironmentsInitParameters struct {

	// (Block List) (see below for nested schema)
	ApprovalSettings []EnvironmentsApprovalSettingsInitParameters `json:"approvalSettings,omitempty" tf:"approval_settings,omitempty"`

	// (String) The color swatch as an RGB hex value with no leading #. For example: 000000
	// The color swatch as an RGB hex value with no leading `#`. For example: `000000`
	Color *string `json:"color,omitempty" tf:"color,omitempty"`

	// (Boolean) Set to true if this environment requires confirmation for flag and segment changes. This field will default to false when not set.
	// Set to `true` if this environment requires confirmation for flag and segment changes. This field will default to `false` when not set.
	ConfirmChanges *bool `json:"confirmChanges,omitempty" tf:"confirm_changes,omitempty"`

	// (Boolean) Denotes whether the environment is critical.
	// Denotes whether the environment is critical.
	Critical *bool `json:"critical,omitempty" tf:"critical,omitempty"`

	// (Number) The TTL for the environment. This must be between 0 and 60 minutes. The TTL setting only applies to environments using the PHP SDK. This field will default to 0 when not set. To learn more, read TTL settings.
	// The TTL for the environment. This must be between 0 and 60 minutes. The TTL setting only applies to environments using the PHP SDK. This field will default to `0` when not set. To learn more, read [TTL settings](https://docs.launchdarkly.com/home/organize/environments#ttl-settings).
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// (Boolean) Set to true to enable data export for every flag created in this environment after you configure this argument. This field will default to false when not set. To learn more, read Data Export.
	// Set to `true` to enable data export for every flag created in this environment after you configure this argument. This field will default to `false` when not set. To learn more, read [Data Export](https://docs.launchdarkly.com/home/data-export).
	DefaultTrackEvents *bool `json:"defaultTrackEvents,omitempty" tf:"default_track_events,omitempty"`

	// (String) The project's unique key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The project-unique key for the environment. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) The project's name.
	// The name of the environment.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Set to true if this environment requires comments for flag and segment changes. This field will default to false when not set.
	// Set to `true` if this environment requires comments for flag and segment changes. This field will default to `false` when not set.
	RequireComments *bool `json:"requireComments,omitempty" tf:"require_comments,omitempty"`

	// side SDK cannot impersonate another user. This field will default to false when not set.
	// Set to `true` to ensure a user of the client-side SDK cannot impersonate another user. This field will default to `false` when not set.
	SecureMode *bool `json:"secureMode,omitempty" tf:"secure_mode,omitempty"`

	// (Set of String) Tags associated with your resource.
	// Tags associated with your resource.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EnvironmentsObservation struct {

	// (Block List) (see below for nested schema)
	ApprovalSettings []EnvironmentsApprovalSettingsObservation `json:"approvalSettings,omitempty" tf:"approval_settings,omitempty"`

	// (String) The color swatch as an RGB hex value with no leading #. For example: 000000
	// The color swatch as an RGB hex value with no leading `#`. For example: `000000`
	Color *string `json:"color,omitempty" tf:"color,omitempty"`

	// (Boolean) Set to true if this environment requires confirmation for flag and segment changes. This field will default to false when not set.
	// Set to `true` if this environment requires confirmation for flag and segment changes. This field will default to `false` when not set.
	ConfirmChanges *bool `json:"confirmChanges,omitempty" tf:"confirm_changes,omitempty"`

	// (Boolean) Denotes whether the environment is critical.
	// Denotes whether the environment is critical.
	Critical *bool `json:"critical,omitempty" tf:"critical,omitempty"`

	// (Number) The TTL for the environment. This must be between 0 and 60 minutes. The TTL setting only applies to environments using the PHP SDK. This field will default to 0 when not set. To learn more, read TTL settings.
	// The TTL for the environment. This must be between 0 and 60 minutes. The TTL setting only applies to environments using the PHP SDK. This field will default to `0` when not set. To learn more, read [TTL settings](https://docs.launchdarkly.com/home/organize/environments#ttl-settings).
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// (Boolean) Set to true to enable data export for every flag created in this environment after you configure this argument. This field will default to false when not set. To learn more, read Data Export.
	// Set to `true` to enable data export for every flag created in this environment after you configure this argument. This field will default to `false` when not set. To learn more, read [Data Export](https://docs.launchdarkly.com/home/data-export).
	DefaultTrackEvents *bool `json:"defaultTrackEvents,omitempty" tf:"default_track_events,omitempty"`

	// (String) The project's unique key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The project-unique key for the environment. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) The project's name.
	// The name of the environment.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Boolean) Set to true if this environment requires comments for flag and segment changes. This field will default to false when not set.
	// Set to `true` if this environment requires comments for flag and segment changes. This field will default to `false` when not set.
	RequireComments *bool `json:"requireComments,omitempty" tf:"require_comments,omitempty"`

	// side SDK cannot impersonate another user. This field will default to false when not set.
	// Set to `true` to ensure a user of the client-side SDK cannot impersonate another user. This field will default to `false` when not set.
	SecureMode *bool `json:"secureMode,omitempty" tf:"secure_mode,omitempty"`

	// (Set of String) Tags associated with your resource.
	// Tags associated with your resource.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EnvironmentsParameters struct {

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	ApprovalSettings []EnvironmentsApprovalSettingsParameters `json:"approvalSettings,omitempty" tf:"approval_settings,omitempty"`

	// (String) The color swatch as an RGB hex value with no leading #. For example: 000000
	// The color swatch as an RGB hex value with no leading `#`. For example: `000000`
	// +kubebuilder:validation:Optional
	Color *string `json:"color" tf:"color,omitempty"`

	// (Boolean) Set to true if this environment requires confirmation for flag and segment changes. This field will default to false when not set.
	// Set to `true` if this environment requires confirmation for flag and segment changes. This field will default to `false` when not set.
	// +kubebuilder:validation:Optional
	ConfirmChanges *bool `json:"confirmChanges,omitempty" tf:"confirm_changes,omitempty"`

	// (Boolean) Denotes whether the environment is critical.
	// Denotes whether the environment is critical.
	// +kubebuilder:validation:Optional
	Critical *bool `json:"critical,omitempty" tf:"critical,omitempty"`

	// (Number) The TTL for the environment. This must be between 0 and 60 minutes. The TTL setting only applies to environments using the PHP SDK. This field will default to 0 when not set. To learn more, read TTL settings.
	// The TTL for the environment. This must be between 0 and 60 minutes. The TTL setting only applies to environments using the PHP SDK. This field will default to `0` when not set. To learn more, read [TTL settings](https://docs.launchdarkly.com/home/organize/environments#ttl-settings).
	// +kubebuilder:validation:Optional
	DefaultTTL *float64 `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`

	// (Boolean) Set to true to enable data export for every flag created in this environment after you configure this argument. This field will default to false when not set. To learn more, read Data Export.
	// Set to `true` to enable data export for every flag created in this environment after you configure this argument. This field will default to `false` when not set. To learn more, read [Data Export](https://docs.launchdarkly.com/home/data-export).
	// +kubebuilder:validation:Optional
	DefaultTrackEvents *bool `json:"defaultTrackEvents,omitempty" tf:"default_track_events,omitempty"`

	// (String) The project's unique key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The project-unique key for the environment. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// (String) The project's name.
	// The name of the environment.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (Boolean) Set to true if this environment requires comments for flag and segment changes. This field will default to false when not set.
	// Set to `true` if this environment requires comments for flag and segment changes. This field will default to `false` when not set.
	// +kubebuilder:validation:Optional
	RequireComments *bool `json:"requireComments,omitempty" tf:"require_comments,omitempty"`

	// side SDK cannot impersonate another user. This field will default to false when not set.
	// Set to `true` to ensure a user of the client-side SDK cannot impersonate another user. This field will default to `false` when not set.
	// +kubebuilder:validation:Optional
	SecureMode *bool `json:"secureMode,omitempty" tf:"secure_mode,omitempty"`

	// (Set of String) Tags associated with your resource.
	// Tags associated with your resource.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ProjectInitParameters struct {

	// side SDKs can use new flags by default. (see below for nested schema)
	// A block describing which client-side SDKs can use new flags by default.
	DefaultClientSideAvailability []DefaultClientSideAvailabilityInitParameters `json:"defaultClientSideAvailability,omitempty" tf:"default_client_side_availability,omitempty"`

	// (Block List, Min: 1) List of nested environments blocks describing LaunchDarkly environments that belong to the project.
	// List of nested `environments` blocks describing LaunchDarkly environments that belong to the project.
	//
	// -> **Note:** Mixing the use of nested `environments` blocks and [`launchdarkly_environment`](/docs/providers/launchdarkly/r/environment.html) resources is not recommended.
	Environments []EnvironmentsInitParameters `json:"environments,omitempty" tf:"environments,omitempty"`

	// unique key for the environment. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The project's unique key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) The name of the environment.
	// The project's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) Tags associated with your resource.
	// Tags associated with your resource.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ProjectObservation struct {

	// side SDKs can use new flags by default. (see below for nested schema)
	// A block describing which client-side SDKs can use new flags by default.
	DefaultClientSideAvailability []DefaultClientSideAvailabilityObservation `json:"defaultClientSideAvailability,omitempty" tf:"default_client_side_availability,omitempty"`

	// (Block List, Min: 1) List of nested environments blocks describing LaunchDarkly environments that belong to the project.
	// List of nested `environments` blocks describing LaunchDarkly environments that belong to the project.
	//
	// -> **Note:** Mixing the use of nested `environments` blocks and [`launchdarkly_environment`](/docs/providers/launchdarkly/r/environment.html) resources is not recommended.
	Environments []EnvironmentsObservation `json:"environments,omitempty" tf:"environments,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// unique key for the environment. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The project's unique key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) The name of the environment.
	// The project's name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) Tags associated with your resource.
	// Tags associated with your resource.
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ProjectParameters struct {

	// side SDKs can use new flags by default. (see below for nested schema)
	// A block describing which client-side SDKs can use new flags by default.
	// +kubebuilder:validation:Optional
	DefaultClientSideAvailability []DefaultClientSideAvailabilityParameters `json:"defaultClientSideAvailability,omitempty" tf:"default_client_side_availability,omitempty"`

	// (Block List, Min: 1) List of nested environments blocks describing LaunchDarkly environments that belong to the project.
	// List of nested `environments` blocks describing LaunchDarkly environments that belong to the project.
	//
	// -> **Note:** Mixing the use of nested `environments` blocks and [`launchdarkly_environment`](/docs/providers/launchdarkly/r/environment.html) resources is not recommended.
	// +kubebuilder:validation:Optional
	Environments []EnvironmentsParameters `json:"environments,omitempty" tf:"environments,omitempty"`

	// unique key for the environment. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The project's unique key. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (String) The name of the environment.
	// The project's name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) Tags associated with your resource.
	// Tags associated with your resource.
	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ProjectSpec defines the desired state of Project
type ProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectParameters `json:"forProvider"`
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
	InitProvider ProjectInitParameters `json:"initProvider,omitempty"`
}

// ProjectStatus defines the observed state of Project.
type ProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Project is the Schema for the Projects API. Provides a LaunchDarkly project resource. This resource allows you to create and manage projects within your LaunchDarkly organization.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,launchdarkly}
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.environments) || (has(self.initProvider) && has(self.initProvider.environments))",message="spec.forProvider.environments is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ProjectSpec   `json:"spec"`
	Status ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectList contains a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// Repository type metadata.
var (
	Project_Kind             = "Project"
	Project_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Project_Kind}.String()
	Project_KindAPIVersion   = Project_Kind + "." + CRDGroupVersion.String()
	Project_GroupVersionKind = CRDGroupVersion.WithKind(Project_Kind)
)

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
