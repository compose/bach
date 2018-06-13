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
	"log"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var listdbtype string
var listfilter string

// listCmd represents the new list deployments command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List deployments attached to account",
	Long:  `List the database deployments attached to an account.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if outputRaw {
			text, errs := c.GetDeploymentsJSON()
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			deployments, errs := c.GetDeployments()
			bailOnErrs(errs)

			matcher, err := regexp.Compile(listfilter)
			if err != nil {
				log.Fatal(err)
			}

			if !outputJSON {
				fmt.Printf("%-24s %-14s %-40s\n", "ID", "Type", "Name")
				fmt.Printf("%-24s %-14s %-40s\n", strings.Repeat("-", 24), strings.Repeat("-", 14), strings.Repeat("-", 40))

				for _, v := range *deployments {
					if (listdbtype == "" || listdbtype == v.Type) && matcher.MatchString(v.Name) {
						fmt.Printf("%-24s %-14s %-40s\n", v.ID, v.Type, v.Name)
					}
				}
			} else {
				printAsJSON(deployments)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&listdbtype, "type", "t", "", "Only this database type")
	listCmd.Flags().StringVarP(&listfilter, "filter", "f", "", "Regular expression to filter names on")
}
