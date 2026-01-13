/*
Copyright 2022 Upbound Inc.
*/

// Package v1alpha1 contains the core resources of the launchdarkly upjet provider.
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
var SchemeBuilder = &scheme.Builder{GroupVersion: schema.GroupVersion{Group: "launchdarkly.launchdarkly.com", Version: "v1alpha1"}}
