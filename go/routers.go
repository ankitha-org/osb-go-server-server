/*
 * Open Service Broker API
 *
 * The Open Service Broker API defines an HTTP(S) interface between Platforms and Service Brokers.
 *
 * API version: master - might contain changes that are not yet released
 * Contact: open-service-broker-api@googlegroups.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CatalogGet",
		strings.ToUpper("Get"),
		"/v2/catalog",
		CatalogGet,
	},

	Route{
		"ServiceBindingBinding",
		strings.ToUpper("Put"),
		"/v2/service_instances/{instance_id}/service_bindings/{binding_id}",
		ServiceBindingBinding,
	},

	Route{
		"ServiceBindingGet",
		strings.ToUpper("Get"),
		"/v2/service_instances/{instance_id}/service_bindings/{binding_id}",
		ServiceBindingGet,
	},

	Route{
		"ServiceBindingLastOperationGet",
		strings.ToUpper("Get"),
		"/v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation",
		ServiceBindingLastOperationGet,
	},

	Route{
		"ServiceBindingUnbinding",
		strings.ToUpper("Delete"),
		"/v2/service_instances/{instance_id}/service_bindings/{binding_id}",
		ServiceBindingUnbinding,
	},

	Route{
		"ServiceInstanceDeprovision",
		strings.ToUpper("Delete"),
		"/v2/service_instances/{instance_id}",
		ServiceInstanceDeprovision,
	},

	Route{
		"ServiceInstanceGet",
		strings.ToUpper("Get"),
		"/v2/service_instances/{instance_id}",
		ServiceInstanceGet,
	},

	Route{
		"ServiceInstanceLastOperationGet",
		strings.ToUpper("Get"),
		"/v2/service_instances/{instance_id}/last_operation",
		ServiceInstanceLastOperationGet,
	},

	Route{
		"ServiceInstanceProvision",
		strings.ToUpper("Put"),
		"/v2/service_instances/{instance_id}",
		ServiceInstanceProvision,
	},

	Route{
		"ServiceInstanceUpdate",
		strings.ToUpper("Patch"),
		"/v2/service_instances/{instance_id}",
		ServiceInstanceUpdate,
	},
}
