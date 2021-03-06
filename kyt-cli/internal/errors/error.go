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

package errors

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Er logs the error on stderr and terminates with exit code 1
func Er(msg interface{}) {
	fmt.Fprintf(os.Stderr, "Error: %v", msg)
	os.Exit(1)
}

// TokenConfigCheck checks if a token is present int the config file
func TokenConfigCheck() {
	if !viper.IsSet("alm_token") || !viper.IsSet("dlm_token") {
		fmt.Println("Required access token not found. Please run `login` command.")
		os.Exit(1)
	}
}
