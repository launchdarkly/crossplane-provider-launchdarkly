package metric

import (
	"testing"

	"github.com/launchdarkly/crossplane-provider-launchdarkly/apis/project/v1alpha1"
)

// Fields that customizeDiff validates from spec must not be late-initialized
// from LD's response, or they would be validated as if the user had set them.
func TestLateInitializeSkipsValidatedFields(t *testing.T) {
	m := &v1alpha1.Metric{}
	observed := `{"name": "n", "event_key": "e", "include_units_without_events": true, "percentile_value": 95,
		"selector": ".x", "success_criteria": "HigherThanBaseline", "unit": "ms"}`
	if _, err := m.LateInitialize([]byte(observed)); err != nil {
		t.Fatal(err)
	}
	p := m.Spec.ForProvider
	if p.Name == nil {
		t.Fatal("name was not late-initialized; the test no longer exercises LateInitialize")
	}
	for field, v := range map[string]bool{
		"event_key":                    p.EventKey != nil,
		"include_units_without_events": p.IncludeUnitsWithoutEvents != nil,
		"percentile_value":             p.PercentileValue != nil,
		"selector":                     p.Selector != nil,
		"success_criteria":             p.SuccessCriteria != nil,
		"unit":                         p.Unit != nil,
	} {
		if v {
			t.Errorf("%s was late-initialized into spec", field)
		}
	}
}
