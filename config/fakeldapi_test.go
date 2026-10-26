package config

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"strings"
	"sync"

	ldapi "github.com/launchdarkly/api-client-go/v17"
)

// fakeRoute maps an LD REST endpoint to the api-client-go type it returns.
type fakeRoute struct {
	method  string
	pattern string
	resp    any
	// create routes store the response under <path>/<key or generated id>.
	create bool
	// fixup fills in what the real API derives server-side.
	fixup func(resp map[string]any)
}

// fixFlag mimics the API's default boolean variations and gives the flag a
// config for the "test" environment the cases use.
func fixFlag(resp map[string]any) {
	if v, _ := resp["variations"].([]any); len(v) == 0 {
		resp["variations"] = []any{map[string]any{"value": true}, map[string]any{"value": false}}
	}
	envs, _ := resp["environments"].(map[string]any)
	if envs == nil {
		envs = map[string]any{}
		resp["environments"] = envs
	}
	if _, ok := envs["test"]; !ok {
		envs["test"] = zeroFilled(ldapi.FeatureFlagConfig{})
	}
}

var fakeRoutes = []fakeRoute{
	{http.MethodPost, "/api/v2/projects", ldapi.ProjectRep{}, true, nil},
	{http.MethodGet, "/api/v2/projects/{projectKey}", ldapi.Project{}, false, nil},
	{http.MethodPatch, "/api/v2/projects/{projectKey}", ldapi.ProjectRep{}, false, nil},
	{http.MethodPut, "/api/v2/projects/{projectKey}/context-kinds/{key}", ldapi.UpsertResponseRep{}, false, nil},
	{http.MethodPost, "/api/v2/projects/{projectKey}/environments", ldapi.Environment{}, true, nil},
	{http.MethodGet, "/api/v2/projects/{projectKey}/environments", ldapi.Environments{}, false, nil},
	{http.MethodGet, "/api/v2/projects/{projectKey}/environments/{environmentKey}", ldapi.Environment{}, false, nil},
	{http.MethodPatch, "/api/v2/projects/{projectKey}/environments/{environmentKey}", ldapi.Environment{}, false, nil},
	{http.MethodPost, "/api/v2/flags/{projectKey}", ldapi.FeatureFlag{}, true, fixFlag},
	{http.MethodGet, "/api/v2/flags/{projectKey}/{featureFlagKey}", ldapi.FeatureFlag{}, false, fixFlag},
	{http.MethodPatch, "/api/v2/flags/{projectKey}/{featureFlagKey}", ldapi.FeatureFlag{}, false, fixFlag},
	{http.MethodPost, "/api/v2/segments/{projectKey}/{environmentKey}", ldapi.UserSegment{}, true, nil},
	{http.MethodGet, "/api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}", ldapi.UserSegment{}, false, nil},
	{http.MethodPatch, "/api/v2/segments/{projectKey}/{environmentKey}/{segmentKey}", ldapi.UserSegment{}, false, nil},
	{http.MethodPost, "/api/v2/tokens", ldapi.Token{}, true, nil},
	{http.MethodGet, "/api/v2/tokens/{id}", ldapi.Token{}, false, nil},
	{http.MethodPatch, "/api/v2/tokens/{id}", ldapi.Token{}, false, nil},
	{http.MethodPost, "/api/v2/roles", ldapi.CustomRole{}, true, nil},
	{http.MethodGet, "/api/v2/roles/{customRoleKey}", ldapi.CustomRole{}, false, nil},
	{http.MethodPatch, "/api/v2/roles/{customRoleKey}", ldapi.CustomRole{}, false, nil},
	{http.MethodPost, "/api/v2/webhooks", ldapi.Webhook{}, true, nil},
	{http.MethodGet, "/api/v2/webhooks/{id}", ldapi.Webhook{}, false, nil},
	{http.MethodPatch, "/api/v2/webhooks/{id}", ldapi.Webhook{}, false, nil},
	{http.MethodPost, "/api/v2/members", ldapi.Members{}, true, nil},
	{http.MethodGet, "/api/v2/members", ldapi.Members{}, false, nil},
	{http.MethodGet, "/api/v2/members/{id}", ldapi.Member{}, false, nil},
	{http.MethodPatch, "/api/v2/members/{id}", ldapi.Member{}, false, nil},
	{http.MethodPost, "/api/v2/teams", ldapi.Team{}, true, nil},
	{http.MethodGet, "/api/v2/teams/{teamKey}", ldapi.Team{}, false, nil},
	{http.MethodPatch, "/api/v2/teams/{teamKey}", ldapi.Team{}, false, nil},
	{http.MethodPost, "/api/v2/metrics/{projectKey}", ldapi.MetricRep{}, true, nil},
	{http.MethodGet, "/api/v2/metrics/{projectKey}/{metricKey}", ldapi.MetricRep{}, false, nil},
	{http.MethodPatch, "/api/v2/metrics/{projectKey}/{metricKey}", ldapi.MetricRep{}, false, nil},
	{http.MethodPost, "/api/v2/destinations/{projectKey}/{environmentKey}", ldapi.Destination{}, true, nil},
	{http.MethodGet, "/api/v2/destinations/{projectKey}/{environmentKey}/{id}", ldapi.Destination{}, false, nil},
	{http.MethodPatch, "/api/v2/destinations/{projectKey}/{environmentKey}/{id}", ldapi.Destination{}, false, nil},
	{http.MethodPost, "/api/v2/integrations/{integrationKey}", ldapi.Integration{}, true, nil},
	{http.MethodGet, "/api/v2/integrations/{integrationKey}/{id}", ldapi.Integration{}, false, nil},
	{http.MethodPatch, "/api/v2/integrations/{integrationKey}/{id}", ldapi.Integration{}, false, nil},
	{http.MethodPost, "/api/v2/account/relay-auto-configs", ldapi.RelayAutoConfigRep{}, true, nil},
	{http.MethodGet, "/api/v2/account/relay-auto-configs/{id}", ldapi.RelayAutoConfigRep{}, false, nil},
	{http.MethodPatch, "/api/v2/account/relay-auto-configs/{id}", ldapi.RelayAutoConfigRep{}, false, nil},
}

var pathParam = regexp.MustCompile(`\{(\w+)\}`)

// fakeLDAPI is an in-memory stand-in for the LaunchDarkly REST API. Every
// response is a fully populated instance of the api-client-go response type
// (all optional fields set to zero values), overlaid with the fields the
// provider sent, so the Terraform provider's CRUD code runs end to end.
type fakeLDAPI struct {
	mu    sync.Mutex
	store map[string]map[string]any
	ids   int
}

func newFakeLDAPI() (*fakeLDAPI, *httptest.Server) {
	f := &fakeLDAPI{store: map[string]map[string]any{}}
	return f, httptest.NewTLSServer(f)
}

func (f *fakeLDAPI) reset() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.store = map[string]map[string]any{}
}

func (f *fakeLDAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodDelete {
		delete(f.store, r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	route, params := matchRoute(r.Method, r.URL.Path)
	if route == nil {
		// Not 404: the provider retries 404s with backoff.
		http.Error(w, fmt.Sprintf(`{"message":"fake LD API: unhandled route %s %s"}`, r.Method, r.URL.Path), http.StatusBadRequest)
		return
	}
	var body any
	raw, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(raw, &body)

	path := r.URL.Path
	// Always answer with the route's own response type, overlaid with what was
	// stored, since e.g. POST and GET /projects return different shapes.
	resp := zeroFilled(route.resp)
	if n := names(route); !route.create && len(n) > 0 {
		setKey(resp, n[len(n)-1], params[n[len(n)-1]])
	}
	if stored, ok := f.store[path]; ok && !route.create {
		overlay(resp, stored)
	}
	switch {
	case route.create:
		b := body
		if items, ok := resp["items"].([]any); ok && len(items) > 0 {
			// e.g. POST /members takes and returns a list.
			if list, ok := body.([]any); ok && len(list) > 0 {
				b = list[0]
			}
			item := items[0].(map[string]any)
			f.fill(item, b)
			f.store[path+"/"+itemID(item)] = item
		} else {
			f.fill(resp, b)
			f.store[path+"/"+itemID(resp)] = resp
			// Projects are created with their environments.
			if envs, ok := resp["environments"].([]any); ok {
				for _, e := range envs {
					env := e.(map[string]any)
					f.store[path+"/"+itemID(resp)+"/environments/"+itemID(env)] = env
				}
			}
		}
	case r.Method == http.MethodGet && resp["items"] != nil:
		// List endpoints return whatever was created under them.
		items := []any{}
		for p, v := range f.store {
			if strings.HasPrefix(p, path+"/") && !strings.Contains(p[len(path)+1:], "/") {
				items = append(items, v)
			}
		}
		resp["items"] = items
		resp["totalCount"] = len(items)
	case r.Method == http.MethodPut:
		overlay(resp, body)
		f.store[path] = resp
	}
	if route.fixup != nil {
		route.fixup(resp)
	}
	_ = json.NewEncoder(w).Encode(resp)
}

func (f *fakeLDAPI) fill(resp map[string]any, body any) {
	overlay(resp, body)
	if id, ok := resp["_id"]; ok && id == "" {
		f.ids++
		resp["_id"] = fmt.Sprintf("fake-id-%d", f.ids)
	}
}

func itemID(m map[string]any) string {
	if k, ok := m["key"].(string); ok && k != "" {
		return k
	}
	id, _ := m["_id"].(string)
	return id
}

func names(rt *fakeRoute) []string {
	matches := pathParam.FindAllStringSubmatch(rt.pattern, -1)
	out := make([]string, 0, len(matches))
	for _, m := range matches {
		out = append(out, m[1])
	}
	return out
}

func matchRoute(method, path string) (*fakeRoute, map[string]string) {
	for i := range fakeRoutes {
		rt := &fakeRoutes[i]
		if rt.method != method {
			continue
		}
		names := pathParam.FindAllStringSubmatch(rt.pattern, -1)
		re := "^" + pathParam.ReplaceAllString(rt.pattern, `([^/]+)`) + "$"
		m := regexp.MustCompile(re).FindStringSubmatch(path)
		if m == nil {
			continue
		}
		params := map[string]string{}
		for j, n := range names {
			params[n[1]] = m[j+1]
		}
		return rt, params
	}
	return nil, nil
}

// setKey puts a path parameter into the matching response field, so that a GET
// of a resource that was never created still looks like that resource.
func setKey(resp map[string]any, param, v string) {
	if param == "id" {
		if _, ok := resp["_id"]; ok {
			resp["_id"] = v
		}
		return
	}
	if strings.HasSuffix(param, "Key") || param == "key" {
		if _, ok := resp["key"]; ok {
			resp["key"] = v
		}
	}
}

// overlay copies fields from src into dst where dst already has a field of the
// same JSON kind. List elements are overlaid onto a copy of dst's first element.
func overlay(dst map[string]any, src any) {
	s, ok := src.(map[string]any)
	if !ok {
		return
	}
	for k, sv := range s {
		dst[k] = overlayValue(dst[k], sv)
	}
}

func overlayValue(dv, sv any) any {
	switch d := dv.(type) {
	case map[string]any:
		switch v := sv.(type) {
		case map[string]any:
			overlay(d, v)
		case []any:
			// A plain list stored for a paginated {"items": [...]} field.
			if items, ok := d["items"].([]any); ok {
				d["items"] = overlayList(items, v)
				d["totalCount"] = len(v)
			}
		}
		return d
	case []any:
		sl, ok := sv.([]any)
		if !ok {
			return d
		}
		return overlayList(d, sl)
	case nil:
		return sv
	default:
		if reflect.TypeOf(d) == reflect.TypeOf(sv) {
			return sv
		}
		return d
	}
}

func overlayList(d, sl []any) []any {
	{
		out := make([]any, 0, len(sl))
		for _, e := range sl {
			if len(d) == 0 {
				out = append(out, e)
				continue
			}
			out = append(out, overlayValue(deepCopy(d[0]), e))
		}
		return out
	}
}

func deepCopy(v any) any {
	b, _ := json.Marshal(v)
	var out any
	_ = json.Unmarshal(b, &out)
	return out
}

// zeroFilled returns v's JSON form with every pointer, struct and list populated
// (lists with one element), mimicking a full API response.
func zeroFilled(v any) map[string]any {
	val := fillValue(reflect.TypeOf(v), 0)
	b, err := json.Marshal(val.Interface())
	if err != nil {
		panic(err)
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		panic(err)
	}
	return m
}

func fillValue(t reflect.Type, depth int) reflect.Value {
	v := reflect.New(t).Elem()
	if depth > 6 {
		return v
	}
	switch t.Kind() {
	case reflect.Pointer:
		v.Set(reflect.New(t.Elem()))
		v.Elem().Set(fillValue(t.Elem(), depth+1))
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			if t.Field(i).IsExported() {
				v.Field(i).Set(fillValue(t.Field(i).Type, depth+1))
			}
		}
	case reflect.Slice:
		// Lists of objects get one element so overlay has a template for each
		// item; lists of scalars stay empty.
		e := t.Elem()
		if e.Kind() == reflect.Pointer {
			e = e.Elem()
		}
		if e.Kind() == reflect.Struct {
			v.Set(reflect.Append(reflect.MakeSlice(t, 0, 1), fillValue(t.Elem(), depth+1)))
		} else {
			v.Set(reflect.MakeSlice(t, 0, 0))
		}
	case reflect.Map:
		v.Set(reflect.MakeMap(t))
	}
	return v
}
