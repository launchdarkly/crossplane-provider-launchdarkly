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

type AccessTokenInitParameters struct {

	// (Set of String) A list of custom role IDs to use as access limits for the access token.
	// A list of custom role IDs to use as access limits for the access token.
	// +crossplane:generate:reference:type=github.com/launchdarkly/crossplane-provider-launchdarkly/apis/account/v1alpha1.CustomRole
	// +listType=set
	CustomRoles []*string `json:"customRoles,omitempty" tf:"custom_roles,omitempty"`

	// References to CustomRole in account to populate customRoles.
	// +kubebuilder:validation:Optional
	CustomRolesRefs []v1.Reference `json:"customRolesRefs,omitempty" tf:"-"`

	// Selector for a list of CustomRole in account to populate customRoles.
	// +kubebuilder:validation:Optional
	CustomRolesSelector *v1.Selector `json:"customRolesSelector,omitempty" tf:"-"`

	// (Number) The default API version for this token. Defaults to the latest API version. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The default API version for this token. Defaults to the latest API version. A change in this field will force the destruction of the existing resource and the creation of a new one.
	DefaultAPIVersion *float64 `json:"defaultApiVersion,omitempty" tf:"default_api_version,omitempty"`

	// in or custom role. Using polices. May be specified more than once. (see below for nested schema)
	// Define inline custom roles. An array of statements represented as config blocks with three attributes: effect, resources, actions. May be used in place of a built-in or custom role. [Using polices](https://docs.launchdarkly.com/home/members/role-policies). May be specified more than once.
	InlineRoles []InlineRolesInitParameters `json:"inlineRoles,omitempty" tf:"inline_roles,omitempty"`

	// friendly name for the access token.
	// A human-friendly name for the access token.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// in LaunchDarkly role. Can be reader, writer, or admin
	// A built-in LaunchDarkly role. Can be `reader`, `writer`, or `admin`
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// (Boolean) Whether the token will be a service token. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// Whether the token will be a [service token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#service-tokens). A change in this field will force the destruction of the existing resource and the creation of a new one.
	ServiceToken *bool `json:"serviceToken,omitempty" tf:"service_token,omitempty"`
}

type AccessTokenObservation struct {

	// (Set of String) A list of custom role IDs to use as access limits for the access token.
	// A list of custom role IDs to use as access limits for the access token.
	// +listType=set
	CustomRoles []*string `json:"customRoles,omitempty" tf:"custom_roles,omitempty"`

	// (Number) The default API version for this token. Defaults to the latest API version. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The default API version for this token. Defaults to the latest API version. A change in this field will force the destruction of the existing resource and the creation of a new one.
	DefaultAPIVersion *float64 `json:"defaultApiVersion,omitempty" tf:"default_api_version,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// in or custom role. Using polices. May be specified more than once. (see below for nested schema)
	// Define inline custom roles. An array of statements represented as config blocks with three attributes: effect, resources, actions. May be used in place of a built-in or custom role. [Using polices](https://docs.launchdarkly.com/home/members/role-policies). May be specified more than once.
	InlineRoles []InlineRolesObservation `json:"inlineRoles,omitempty" tf:"inline_roles,omitempty"`

	// friendly name for the access token.
	// A human-friendly name for the access token.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// in LaunchDarkly role. Can be reader, writer, or admin
	// A built-in LaunchDarkly role. Can be `reader`, `writer`, or `admin`
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// (Boolean) Whether the token will be a service token. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// Whether the token will be a [service token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#service-tokens). A change in this field will force the destruction of the existing resource and the creation of a new one.
	ServiceToken *bool `json:"serviceToken,omitempty" tf:"service_token,omitempty"`
}

type AccessTokenParameters struct {

	// (Set of String) A list of custom role IDs to use as access limits for the access token.
	// A list of custom role IDs to use as access limits for the access token.
	// +crossplane:generate:reference:type=github.com/launchdarkly/crossplane-provider-launchdarkly/apis/account/v1alpha1.CustomRole
	// +kubebuilder:validation:Optional
	// +listType=set
	CustomRoles []*string `json:"customRoles,omitempty" tf:"custom_roles,omitempty"`

	// References to CustomRole in account to populate customRoles.
	// +kubebuilder:validation:Optional
	CustomRolesRefs []v1.Reference `json:"customRolesRefs,omitempty" tf:"-"`

	// Selector for a list of CustomRole in account to populate customRoles.
	// +kubebuilder:validation:Optional
	CustomRolesSelector *v1.Selector `json:"customRolesSelector,omitempty" tf:"-"`

	// (Number) The default API version for this token. Defaults to the latest API version. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The default API version for this token. Defaults to the latest API version. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// +kubebuilder:validation:Optional
	DefaultAPIVersion *float64 `json:"defaultApiVersion,omitempty" tf:"default_api_version,omitempty"`

	// in or custom role. Using polices. May be specified more than once. (see below for nested schema)
	// Define inline custom roles. An array of statements represented as config blocks with three attributes: effect, resources, actions. May be used in place of a built-in or custom role. [Using polices](https://docs.launchdarkly.com/home/members/role-policies). May be specified more than once.
	// +kubebuilder:validation:Optional
	InlineRoles []InlineRolesParameters `json:"inlineRoles,omitempty" tf:"inline_roles,omitempty"`

	// friendly name for the access token.
	// A human-friendly name for the access token.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// in LaunchDarkly role. Can be reader, writer, or admin
	// A built-in LaunchDarkly role. Can be `reader`, `writer`, or `admin`
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// (Boolean) Whether the token will be a service token. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// Whether the token will be a [service token](https://docs.launchdarkly.com/home/account-security/api-access-tokens#service-tokens). A change in this field will force the destruction of the existing resource and the creation of a new one.
	// +kubebuilder:validation:Optional
	ServiceToken *bool `json:"serviceToken,omitempty" tf:"service_token,omitempty"`
}

type InlineRolesInitParameters struct {

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

type InlineRolesObservation struct {

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

type InlineRolesParameters struct {

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

// AccessTokenSpec defines the desired state of AccessToken
type AccessTokenSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessTokenParameters `json:"forProvider"`
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
	InitProvider AccessTokenInitParameters `json:"initProvider,omitempty"`
}

// AccessTokenStatus defines the observed state of AccessToken.
type AccessTokenStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessTokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AccessToken is the Schema for the AccessTokens API. Provides a LaunchDarkly access token resource. This resource allows you to create and manage access tokens within your LaunchDarkly organization. Be sure your state is configured securely before using this resource. See https://www.io/docs/state/sensitive-data.html for more details. The resource must contain either a "role", "custom_role" or an "inline_roles" (previously "policy_statements") block. As of v1.7.0, "policy_statements" has been deprecated in favor of "inline_roles".
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,launchdarkly}
type AccessToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessTokenSpec   `json:"spec"`
	Status            AccessTokenStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessTokenList contains a list of AccessTokens
type AccessTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessToken `json:"items"`
}

// Repository type metadata.
var (
	AccessToken_Kind             = "AccessToken"
	AccessToken_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessToken_Kind}.String()
	AccessToken_KindAPIVersion   = AccessToken_Kind + "." + CRDGroupVersion.String()
	AccessToken_GroupVersionKind = CRDGroupVersion.WithKind(AccessToken_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessToken{}, &AccessTokenList{})
}
