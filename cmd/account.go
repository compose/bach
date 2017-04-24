// Copyright © 2017 Compose, an IBM Company
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

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Show Account Details",
	Long:  `Show the details for the Compose account registered with the $COMPOSEAPITOKEN`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if outputRaw {
			text, errs := c.GetAccountJSON()
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			account, errs := c.GetAccount()
			bailOnErrs(errs)

			if !outputJSON {
				fmt.Printf("%15s: %s\n", "ID", account.ID)
				fmt.Printf("%15s: %s\n", "Name", account.Name)
				fmt.Printf("%15s: %s\n", "Slug", account.Slug)
				fmt.Println()
			} else {
				printAsJSON(account)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(accountCmd)
}
