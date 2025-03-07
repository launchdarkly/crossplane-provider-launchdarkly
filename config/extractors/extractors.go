package extractors

import (
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/launchdarkly/crossplane-provider-launchdarkly/config/extractors"
)

// FieldExtractorFnReference returns a reference to function that extracts a field from a resource.
func FieldExtractorFnReference(field string) string {
	return fmt.Sprintf("%s.FieldExtractor(%q)", SelfPackagePath, field)
}

// FieldExtractor returns a function that extracts a field from a resource.
func FieldExtractor(field string) reference.ExtractValueFn {
	return func(mg xpresource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		r, err := paved.GetString("spec.forProvider." + field)
		if err != nil {
			return ""
		}
		return r
	}
}
