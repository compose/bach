// Copyright © 2017 Compose, an IBM company
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// userCmd represents the user command
var usershowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current user information",
	Long:  `Show current user information`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if outputRaw {
			text, errs := c.GetUserJSON()
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			user, errs := c.GetUser()
			bailOnErrs(errs)
			if !outputJSON {
				fmt.Printf("%15s: %s\n", "ID", user.ID)
				fmt.Printf("%15s: %s\n", "Name", user.Name)
				fmt.Println()
			} else {
				printAsJSON(user)
			}
		}
	},
}

func init() {
	userCmd.AddCommand(usershowCmd)
}
