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

	"github.com/spf13/cobra"
)

var dbtype string
var longoutput bool
var filter string

// deploymentsCmd represents the list deployments command
var deploymentsCmd = &cobra.Command{
	Use:   "deployments",
	Short: "Show deployments attached to account",
	Long:  `Show the database deployments attached to an account.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if outputRaw {
			text, errs := c.GetDeploymentsJSON()
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			deployments, errs := c.GetDeployments()
			bailOnErrs(errs)

			matcher, err := regexp.Compile(filter)
			if err != nil {
				log.Fatal(err)
			}

			if !outputJSON {
				for _, v := range *deployments {
					if (dbtype == "" || dbtype == v.Type) && matcher.MatchString(v.Name) {
						fmt.Printf("%15s: %s\n", "ID", v.ID)
						fmt.Printf("%15s: %s\n", "Name", v.Name)
						fmt.Printf("%15s: %s\n", "Type", v.Type)
						if longoutput {
							fmt.Printf("%15s: %s\n", "Created At", v.CreatedAt)
							fmt.Printf("%15s: %s\n", "Web UI Link", getLink(v.Links.ComposeWebUILink))
						}
						fmt.Println()
					}
				}
			} else {
				printAsJSON(deployments)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(deploymentsCmd)
	deploymentsCmd.Flags().StringVarP(&dbtype, "type", "t", "", "Only this database type")
	deploymentsCmd.Flags().BoolVarP(&longoutput, "long", "l", false, "Show all details")
	deploymentsCmd.Flags().StringVarP(&filter, "filter", "f", "", "Regular expression to filter names on")
}
