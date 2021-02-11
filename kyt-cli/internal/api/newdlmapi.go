/*
Copyright © 2021 Ci4Rail GmbH <engineering@ci4rail.com>

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
	"context"

	openapiclient "github.com/ci4rail/kyt/kyt-cli/openapidlm"
)

// NewDlmAPI generates a new APIClient object to issue further API calls
// It considers "server" command line parameter
func NewDlmAPI(serverURL string) (*openapiclient.APIClient, context.Context) {
	configuration := openapiclient.NewConfiguration()

	if serverURL != "" {
		configuration.Servers[0].URL = serverURL
	}

	apiClient := openapiclient.NewAPIClient(configuration)

	return apiClient, context.Background()
}

// NewDlmAPIWithToken generates a new APIClient object to issue further API calls with a specific token
// It considers "server"argument and a specified token
func NewDlmAPIWithToken(serverURL string, token string) (*openapiclient.APIClient, context.Context) {
	configuration := openapiclient.NewConfiguration()

	// configuration.AddDefaultHeader("X-Auth-Token", token)
	configuration.AddDefaultHeader("Authorization", token)
	configuration.AddDefaultHeader("Content-Type", "Content-Type")
	if serverURL != "" {
		configuration.Servers[0].URL = serverURL
	}

	apiClient := openapiclient.NewAPIClient(configuration)

	return apiClient, context.Background()
}