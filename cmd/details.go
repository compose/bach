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

	"github.com/spf13/cobra"
)

// detailsCmd represents the details command
var detailsCmd = &cobra.Command{
	Use:   "details [deployment id/name]",
	Short: "Show details for a deployment",
	Long:  `Show the details for a deployment including connection strings and certificates`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Need a deployment id/name")
		}
		c := getComposeAPI()
		depid, err := resolveDepID(c, args[0])
		if err != nil {
			log.Fatal(err)
		}
		if outputRaw {
			text, errs := c.GetDeploymentJSON(depid)
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			deployment, errs := c.GetDeployment(depid)
			bailOnErrs(errs)

			if !outputJSON {
				printDeployment(*deployment)
			} else {
				printAsJSON(*deployment)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(detailsCmd)
}
