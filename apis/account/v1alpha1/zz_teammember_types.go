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

type TeamMemberInitParameters struct {

	// (Set of String) The list of custom roles keys associated with the team member. Custom roles are only available to customers on an Enterprise plan. To learn more, read about our pricing. To upgrade your plan, contact LaunchDarkly Sales.
	// The list of custom roles keys associated with the team member. Custom roles are only available to customers on an Enterprise plan. To learn more, [read about our pricing](https://launchdarkly.com/pricing/). To upgrade your plan, [contact LaunchDarkly Sales](https://launchdarkly.com/contact-sales/).
	//
	// -> **Note:** each `launchdarkly_team_member` must have either a `role` or `custom_roles` argument.
	// +listType=set
	CustomRoles []*string `json:"customRoles,omitempty" tf:"custom_roles,omitempty"`

	// (String) The unique email address associated with the team member. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The unique email address associated with the team member. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) The team member's given name. Once created, this cannot be updated except by the team member.
	// The team member's given name. Once created, this cannot be updated except by the team member.
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// (String) TThe team member's family name. Once created, this cannot be updated except by the team member.
	// TThe team member's family name. Once created, this cannot be updated except by the team member.
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// (String) The role associated with team member. Supported roles are reader, writer, no_access, or admin. If you don't specify a role, reader is assigned by default.
	// The role associated with team member. Supported roles are `reader`, `writer`, `no_access`, or `admin`. If you don't specify a role, `reader` is assigned by default.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// (Block Set) A role attributes block. One block must be defined per role attribute. The key is the role attribute key and the value is a string array of resource keys that apply. (see below for nested schema)
	// A role attributes block. One block must be defined per role attribute. The key is the role attribute key and the value is a string array of resource keys that apply.
	RoleAttributes []TeamMemberRoleAttributesInitParameters `json:"roleAttributes,omitempty" tf:"role_attributes,omitempty"`
}

type TeamMemberObservation struct {

	// (Set of String) The list of custom roles keys associated with the team member. Custom roles are only available to customers on an Enterprise plan. To learn more, read about our pricing. To upgrade your plan, contact LaunchDarkly Sales.
	// The list of custom roles keys associated with the team member. Custom roles are only available to customers on an Enterprise plan. To learn more, [read about our pricing](https://launchdarkly.com/pricing/). To upgrade your plan, [contact LaunchDarkly Sales](https://launchdarkly.com/contact-sales/).
	//
	// -> **Note:** each `launchdarkly_team_member` must have either a `role` or `custom_roles` argument.
	// +listType=set
	CustomRoles []*string `json:"customRoles,omitempty" tf:"custom_roles,omitempty"`

	// (String) The unique email address associated with the team member. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The unique email address associated with the team member. A change in this field will force the destruction of the existing resource and the creation of a new one.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) The team member's given name. Once created, this cannot be updated except by the team member.
	// The team member's given name. Once created, this cannot be updated except by the team member.
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// (String) The 24 character alphanumeric ID of the team member.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) TThe team member's family name. Once created, this cannot be updated except by the team member.
	// TThe team member's family name. Once created, this cannot be updated except by the team member.
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// (String) The role associated with team member. Supported roles are reader, writer, no_access, or admin. If you don't specify a role, reader is assigned by default.
	// The role associated with team member. Supported roles are `reader`, `writer`, `no_access`, or `admin`. If you don't specify a role, `reader` is assigned by default.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// (Block Set) A role attributes block. One block must be defined per role attribute. The key is the role attribute key and the value is a string array of resource keys that apply. (see below for nested schema)
	// A role attributes block. One block must be defined per role attribute. The key is the role attribute key and the value is a string array of resource keys that apply.
	RoleAttributes []TeamMemberRoleAttributesObservation `json:"roleAttributes,omitempty" tf:"role_attributes,omitempty"`
}

type TeamMemberParameters struct {

	// (Set of String) The list of custom roles keys associated with the team member. Custom roles are only available to customers on an Enterprise plan. To learn more, read about our pricing. To upgrade your plan, contact LaunchDarkly Sales.
	// The list of custom roles keys associated with the team member. Custom roles are only available to customers on an Enterprise plan. To learn more, [read about our pricing](https://launchdarkly.com/pricing/). To upgrade your plan, [contact LaunchDarkly Sales](https://launchdarkly.com/contact-sales/).
	//
	// -> **Note:** each `launchdarkly_team_member` must have either a `role` or `custom_roles` argument.
	// +kubebuilder:validation:Optional
	// +listType=set
	CustomRoles []*string `json:"customRoles,omitempty" tf:"custom_roles,omitempty"`

	// (String) The unique email address associated with the team member. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// The unique email address associated with the team member. A change in this field will force the destruction of the existing resource and the creation of a new one.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) The team member's given name. Once created, this cannot be updated except by the team member.
	// The team member's given name. Once created, this cannot be updated except by the team member.
	// +kubebuilder:validation:Optional
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// (String) TThe team member's family name. Once created, this cannot be updated except by the team member.
	// TThe team member's family name. Once created, this cannot be updated except by the team member.
	// +kubebuilder:validation:Optional
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// (String) The role associated with team member. Supported roles are reader, writer, no_access, or admin. If you don't specify a role, reader is assigned by default.
	// The role associated with team member. Supported roles are `reader`, `writer`, `no_access`, or `admin`. If you don't specify a role, `reader` is assigned by default.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// (Block Set) A role attributes block. One block must be defined per role attribute. The key is the role attribute key and the value is a string array of resource keys that apply. (see below for nested schema)
	// A role attributes block. One block must be defined per role attribute. The key is the role attribute key and the value is a string array of resource keys that apply.
	// +kubebuilder:validation:Optional
	RoleAttributes []TeamMemberRoleAttributesParameters `json:"roleAttributes,omitempty" tf:"role_attributes,omitempty"`
}

type TeamMemberRoleAttributesInitParameters struct {

	// (String) The key / name of your role attribute. In the example $${roleAttribute/testAttribute}, the key is testAttribute.
	// The key / name of your role attribute. In the example `$${roleAttribute/testAttribute}`, the key is `testAttribute`.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (List of String) A list of values for your role attribute. For example, if your policy statement defines the resource "proj/$${roleAttribute/testAttribute}", the values would be the keys of the projects you wanted to assign access to.
	// A list of values for your role attribute. For example, if your policy statement defines the resource `"proj/$${roleAttribute/testAttribute}"`, the values would be the keys of the projects you wanted to assign access to.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type TeamMemberRoleAttributesObservation struct {

	// (String) The key / name of your role attribute. In the example $${roleAttribute/testAttribute}, the key is testAttribute.
	// The key / name of your role attribute. In the example `$${roleAttribute/testAttribute}`, the key is `testAttribute`.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// (List of String) A list of values for your role attribute. For example, if your policy statement defines the resource "proj/$${roleAttribute/testAttribute}", the values would be the keys of the projects you wanted to assign access to.
	// A list of values for your role attribute. For example, if your policy statement defines the resource `"proj/$${roleAttribute/testAttribute}"`, the values would be the keys of the projects you wanted to assign access to.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type TeamMemberRoleAttributesParameters struct {

	// (String) The key / name of your role attribute. In the example $${roleAttribute/testAttribute}, the key is testAttribute.
	// The key / name of your role attribute. In the example `$${roleAttribute/testAttribute}`, the key is `testAttribute`.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// (List of String) A list of values for your role attribute. For example, if your policy statement defines the resource "proj/$${roleAttribute/testAttribute}", the values would be the keys of the projects you wanted to assign access to.
	// A list of values for your role attribute. For example, if your policy statement defines the resource `"proj/$${roleAttribute/testAttribute}"`, the values would be the keys of the projects you wanted to assign access to.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

// TeamMemberSpec defines the desired state of TeamMember
type TeamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TeamMemberParameters `json:"forProvider"`
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
	InitProvider TeamMemberInitParameters `json:"initProvider,omitempty"`
}

// TeamMemberStatus defines the observed state of TeamMember.
type TeamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TeamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TeamMember is the Schema for the TeamMembers API. Provides a LaunchDarkly team member resource. This resource allows you to create and manage team members within your LaunchDarkly organization. -> Note: You can only manage team members with "admin" level personal access tokens. To learn more, read Managing Teams https://docs.launchdarkly.com/home/teams/managing.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,launchdarkly}
type TeamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || (has(self.initProvider) && has(self.initProvider.email))",message="spec.forProvider.email is a required parameter"
	Spec   TeamMemberSpec   `json:"spec"`
	Status TeamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TeamMemberList contains a list of TeamMembers
type TeamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TeamMember `json:"items"`
}

// Repository type metadata.
var (
	TeamMember_Kind             = "TeamMember"
	TeamMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TeamMember_Kind}.String()
	TeamMember_KindAPIVersion   = TeamMember_Kind + "." + CRDGroupVersion.String()
	TeamMember_GroupVersionKind = CRDGroupVersion.WithKind(TeamMember_Kind)
)

func init() {
	SchemeBuilder.Register(&TeamMember{}, &TeamMemberList{})
}
