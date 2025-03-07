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

// ResolveReferences of this EnvironmentSegment.
func (mg *EnvironmentSegment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EnvKey),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EnvKeyRef,
		Selector:     mg.Spec.ForProvider.EnvKeySelector,
		To: reference.To{
			List:    &EnvironmentList{},
			Managed: &Environment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EnvKey")
	}
	mg.Spec.ForProvider.EnvKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EnvKeyRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EnvKey),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.EnvKeyRef,
		Selector:     mg.Spec.InitProvider.EnvKeySelector,
		To: reference.To{
			List:    &EnvironmentList{},
			Managed: &Environment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EnvKey")
	}
	mg.Spec.InitProvider.EnvKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EnvKeyRef = rsp.ResolvedReference

	return nil
}
