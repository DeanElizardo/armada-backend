package routers

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

var routes = []route{
	newRoute("POST", "/cohort/create", cohortCreate),
	newRoute("GET", "/cohort/all", cohortGetAll),
	newRoute("GET", "/cohort/([0-9]+)", cohortGetSingle),
	newRoute("GET", "/cohort/([0-9]+)/courses", coursesGetForCohort),
	newRoute("PUT", "/cohort/([0-9]+)/update", cohortUpdate),
	newRoute("DELETE", "/cohort/([0-9]+)/delete", cohortDelete),
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type ctxKey struct{}

func getField(r *http.Request, index int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[index]
}

func Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			// Omit the first match, because it will be the full pattern for a route
			// Elements in match[1:] will be the path parameters, passed as a context
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		// The allow header must be set to send 405; indicates what methods are NOT allowed in the response headers
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}
