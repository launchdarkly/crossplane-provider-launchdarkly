// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AccessToken.
func (mg *AccessToken) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CustomRoles),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.CustomRolesRefs,
		Selector:      mg.Spec.ForProvider.CustomRolesSelector,
		To: reference.To{
			List:    &CustomRoleList{},
			Managed: &CustomRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomRoles")
	}
	mg.Spec.ForProvider.CustomRoles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.CustomRolesRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.CustomRoles),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.CustomRolesRefs,
		Selector:      mg.Spec.InitProvider.CustomRolesSelector,
		To: reference.To{
			List:    &CustomRoleList{},
			Managed: &CustomRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CustomRoles")
	}
	mg.Spec.InitProvider.CustomRoles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.CustomRolesRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this Team.
func (mg *Team) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CustomRoleKeys),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.CustomRoleKeysRefs,
		Selector:      mg.Spec.ForProvider.CustomRoleKeysSelector,
		To: reference.To{
			List:    &CustomRoleList{},
			Managed: &CustomRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomRoleKeys")
	}
	mg.Spec.ForProvider.CustomRoleKeys = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.CustomRoleKeysRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Maintainers),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.MaintainersRefs,
		Selector:      mg.Spec.ForProvider.MaintainersSelector,
		To: reference.To{
			List:    &TeamMemberList{},
			Managed: &TeamMember{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Maintainers")
	}
	mg.Spec.ForProvider.Maintainers = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.MaintainersRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.MemberIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.MemberIdsRefs,
		Selector:      mg.Spec.ForProvider.MemberIdsSelector,
		To: reference.To{
			List:    &TeamMemberList{},
			Managed: &TeamMember{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MemberIds")
	}
	mg.Spec.ForProvider.MemberIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.MemberIdsRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.CustomRoleKeys),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.CustomRoleKeysRefs,
		Selector:      mg.Spec.InitProvider.CustomRoleKeysSelector,
		To: reference.To{
			List:    &CustomRoleList{},
			Managed: &CustomRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CustomRoleKeys")
	}
	mg.Spec.InitProvider.CustomRoleKeys = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.CustomRoleKeysRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Maintainers),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.MaintainersRefs,
		Selector:      mg.Spec.InitProvider.MaintainersSelector,
		To: reference.To{
			List:    &TeamMemberList{},
			Managed: &TeamMember{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Maintainers")
	}
	mg.Spec.InitProvider.Maintainers = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.MaintainersRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.MemberIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.MemberIdsRefs,
		Selector:      mg.Spec.InitProvider.MemberIdsSelector,
		To: reference.To{
			List:    &TeamMemberList{},
			Managed: &TeamMember{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MemberIds")
	}
	mg.Spec.InitProvider.MemberIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.MemberIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this TeamRoleMapping.
func (mg *TeamRoleMapping) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CustomRoleKeys),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.CustomRoleKeysRefs,
		Selector:      mg.Spec.ForProvider.CustomRoleKeysSelector,
		To: reference.To{
			List:    &CustomRoleList{},
			Managed: &CustomRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomRoleKeys")
	}
	mg.Spec.ForProvider.CustomRoleKeys = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.CustomRoleKeysRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TeamKey),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TeamKeyRef,
		Selector:     mg.Spec.ForProvider.TeamKeySelector,
		To: reference.To{
			List:    &TeamList{},
			Managed: &Team{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TeamKey")
	}
	mg.Spec.ForProvider.TeamKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TeamKeyRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.CustomRoleKeys),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.CustomRoleKeysRefs,
		Selector:      mg.Spec.InitProvider.CustomRoleKeysSelector,
		To: reference.To{
			List:    &CustomRoleList{},
			Managed: &CustomRole{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CustomRoleKeys")
	}
	mg.Spec.InitProvider.CustomRoleKeys = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.CustomRoleKeysRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TeamKey),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.TeamKeyRef,
		Selector:     mg.Spec.InitProvider.TeamKeySelector,
		To: reference.To{
			List:    &TeamList{},
			Managed: &Team{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TeamKey")
	}
	mg.Spec.InitProvider.TeamKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TeamKeyRef = rsp.ResolvedReference

	return nil
}
