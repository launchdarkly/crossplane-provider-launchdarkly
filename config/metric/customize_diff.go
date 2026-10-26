package metric

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	kindClick    = "click"
	kindCustom   = "custom"
	kindPageview = "pageview"
	percentile   = "percentile"
	urlsField    = "urls"
)

// customizeDiff is a port of the upstream customizeMetricDiff that does not
// panic under upjet. Upstream assumes Terraform's raw config, where an omitted
// block is an empty list; upjet passes it as null, and sometimes passes no raw
// config at all. Without a raw config, only the config-independent parts run.
func customizeDiff(_ context.Context, diff *schema.ResourceDiff, _ any) error {
	cfg := diff.GetRawConfig()
	if hasConfig(cfg) {
		if err := validateConfig(diff, cfg); err != nil {
			return err
		}
	}

	// If anything is changed at all, expect a new value will be computed for "version".
	if len(diff.GetChangedKeysPrefix("")) > 0 {
		if err := diff.SetNewComputed("version"); err != nil {
			return err
		}
	}

	if hasConfig(cfg) && attr(cfg, "include_units_without_events").IsNull() {
		return diff.SetNew("include_units_without_events", diff.Get("analysis_type").(string) != percentile)
	}
	return nil
}

func validateConfig(diff *schema.ResourceDiff, cfg cty.Value) error { //nolint:gocyclo // mirrors upstream
	kind := diff.Get("kind").(string)
	isSet := func(name string) bool { return !attr(cfg, name).IsNull() }
	urls := attr(cfg, urlsField)

	switch kind {
	case kindClick, kindPageview:
		if kind == kindClick && !isSet("selector") {
			return fmt.Errorf("click metrics require 'selector' to be set")
		}
		if length(urls) == 0 {
			return fmt.Errorf("%s metrics require an 'urls' block to be set", kind)
		}
		if urls.IsKnown() && urls.ForEachElement(invalidURL) {
			return fmt.Errorf("'urls' block is misconfigured, please check documentation for required fields")
		}
		for _, f := range []string{"success_criteria", "unit", "event_key"} {
			if isSet(f) {
				return fmt.Errorf("%s metrics do not accept '%s'", kind, f)
			}
		}
		if kind == kindPageview && isSet("selector") {
			return fmt.Errorf("pageview metrics do not accept 'selector'")
		}
	case kindCustom:
		if !isSet("success_criteria") {
			if err := diff.SetNew("success_criteria", "HigherThanBaseline"); err != nil {
				return err
			}
		}
		if n := attr(cfg, "is_numeric"); n.IsKnown() && !n.IsNull() && n.True() {
			if !isSet("success_criteria") {
				return fmt.Errorf("numeric custom metrics require 'success_criteria' to be set")
			}
			if !isSet("unit") {
				return fmt.Errorf("numeric custom metrics require 'unit' to be set")
			}
		}
		if !isSet("event_key") {
			return fmt.Errorf("custom meterics require 'event_key' to be set")
		}
		if length(urls) != 0 {
			return fmt.Errorf("custom metrics do not accept a 'urls' block")
		}
		if isSet("selector") {
			return fmt.Errorf("custom metrics do not accept 'selector'")
		}
	}

	analysisType := diff.Get("analysis_type").(string)
	if analysisType == percentile {
		if !isSet("percentile_value") {
			return fmt.Errorf("percentile_value is required when analysis_type is percentile")
		}
		if u := attr(cfg, "include_units_without_events"); u.IsKnown() && !u.IsNull() && u.True() {
			return fmt.Errorf("include_units_without_events is not supported for percentile metrics")
		}
	} else if isSet("percentile_value") {
		return fmt.Errorf("%s type metrics can not have percentile values", analysisType)
	}
	return nil
}

// invalidURL reports whether a urls block is missing the field its kind
// requires or sets a field that belongs to another kind.
func invalidURL(_, v cty.Value) bool {
	required := map[string]string{"canonical": "url", "exact": "url", "substring": "substring", "regex": "pattern", "pattern": "pattern"}
	k := attr(v, "kind")
	if k.IsNull() || !k.IsKnown() {
		return false
	}
	want, ok := required[k.AsString()]
	if !ok {
		return false
	}
	for _, f := range []string{"url", "substring", "pattern"} {
		if attr(v, f).IsNull() == (f == want) {
			return true
		}
	}
	return false
}

func hasConfig(v cty.Value) bool {
	return !v.IsNull() && v.IsKnown() && v.Type().IsObjectType()
}

func attr(v cty.Value, name string) cty.Value {
	if !hasConfig(v) || !v.Type().HasAttribute(name) {
		return cty.NullVal(cty.DynamicPseudoType)
	}
	return v.GetAttr(name)
}

func length(v cty.Value) int {
	if v.IsNull() || !v.IsKnown() {
		return 0
	}
	return v.LengthInt()
}
