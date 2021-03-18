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
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// NewRouter returns a new router.
func NewRouter() (*mux.Router, error) {

	r := mux.NewRouter()
	jwtMiddleware := token.CreateMiddleware("https://edgefarm-dev.eu.auth0.com")

	r.Handle("/v1/devices/", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(DevicesGet))))

	r.Handle("/v1/devices/{did}", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(DevicesDidGet))))

	return r, nil
}
