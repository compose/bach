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
	"log"

	composeAPI "github.com/compose/gocomposeapi"
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var useraddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add user",
	Long:  `Add user`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if outputRaw {
			if outputRaw {
				log.Fatal("Raw mode not supported for user add")
			}
		} else {
			params := composeAPI.UserParams{}
			if len(args) >= 2 {
				params.Name = args[0]
				params.Email = args[1]
				if len(args) == 3 {
					params.Phone = args[2]
				}
			} else {
				log.Fatal("Need at least a name and an email ")
			}
			account, errs := c.GetAccount()
			bailOnErrs(errs)
			user, errs := c.CreateAccountUser(account.ID, params)
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
	userCmd.AddCommand(useraddCmd)
}
