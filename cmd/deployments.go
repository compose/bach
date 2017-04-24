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

// deploymentsCmd represents the deployments command
var deploymentsCmd = &cobra.Command{
	Use:   "deployments",
	Short: "Show Deployments attached to account",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if outputRaw {
			text, errs := c.GetDeploymentsJSON()
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			deployments, errs := c.GetDeployments()
			bailOnErrs(errs)

			if !outputJSON {
				for _, v := range *deployments {
					fmt.Printf("%15s: %s\n", "ID", v.ID)
					fmt.Printf("%15s: %s\n", "Name", v.Name)
					fmt.Printf("%15s: %s\n", "Type", v.Type)
					fmt.Printf("%15s: %s\n", "Created At", v.CreatedAt)
					fmt.Printf("%15s: %s\n", "Web UI Link", getLink(v.Links.ComposeWebUILink))
					fmt.Println()
				}
			} else {
				printAsJSON(deployments)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(deploymentsCmd)
}
