/*
Copyright © 2021 Ci4Rail GmbH

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"net/http"

	token "github.com/ci4rail/kyt/kyt-server-common/token"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// NewRouter returns a new router.
func NewRouter() (*mux.Router, error) {

	r := mux.NewRouter()
	jwtMiddleware := token.CreateMiddleware("https://edgefarm-dev.eu.auth0.com")
	r.Handle("/v1/runtimes/", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(RuntimesGet))))

	return r, nil
}

// // Route is the information for every URI.
// type Route struct {
// 	// Name is the name of this Route.
// 	Name string
// 	// Method is the string for the HTTP method. ex) GET, POST etc..
// 	Method string
// 	// Pattern is the pattern of the URI.
// 	Pattern string
// 	// HandlerFunc is the handler function of this route.
// 	HandlerFunc gin.HandlerFunc
// }

// // Routes is the list of the generated Route.
// type Routes []Route

// // NewRouter returns a new router.
// func NewRouter() *gin.Engine {
// 	router := gin.Default()

// 	for _, route := range routes {
// 		switch route.Method {
// 		case http.MethodGet:
// 			router.GET(route.Pattern, route.HandlerFunc)
// 		case http.MethodPost:
// 			router.POST(route.Pattern, route.HandlerFunc)
// 		case http.MethodPut:
// 			router.PUT(route.Pattern, route.HandlerFunc)
// 		case http.MethodDelete:
// 			router.DELETE(route.Pattern, route.HandlerFunc)
// 		}
// 	}
// 	return router
// }

// // Index is the index handler.
// func Index(c *gin.Context) {
// 	c.String(http.StatusOK, "Hello KYT!")
// }

// var routes = Routes{
// 	{
// 		"Index",
// 		http.MethodGet,
// 		"/v1/",
// 		Index,
// 	},
// 	{
// 		"RuntimesRidGet",
// 		http.MethodGet,
// 		"/v1/runtimes/:rid",
// 		RuntimesRidGet,
// 	},
// 	{
// 		"RuntimesGet",
// 		http.MethodGet,
// 		"/v1/runtimes/",
// 		RuntimesGet,
// 	},
// 	{
// 		"DeploymentApply",
// 		http.MethodPut,
// 		"/v1/apply/",
// 		ApplyPut,
// 	},
// }
