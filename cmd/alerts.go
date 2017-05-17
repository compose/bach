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

// alertsCmd represents the alerts command
var alertsCmd = &cobra.Command{
	Use:   "alerts [deployment id]",
	Short: "Show Alerts for deployment",
	Long:  `Show the alerts for a deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if len(args) == 0 {
			log.Fatal("Need a deployment id")
		}
		if outputRaw {
			text, errs := c.GetAlertsForDeploymentJSON(args[0])
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			alerts, errs := c.GetAlertsForDeployment(args[0])
			bailOnErrs(errs)

			if !outputJSON {
				fmt.Printf("%15s: %s\n\n", "Summary", alerts.Summary)
				for _, alert := range alerts.Embedded.Alerts {
					fmt.Printf("%15s: %s\n", "Capsule ID", alert.CapsuleID)
					fmt.Printf("%15s: %s\n", "Deployment ID", alert.DeploymentID)
					fmt.Printf("%15s: %s\n", "Message", alert.Message)
					fmt.Printf("%15s: %s\n", "Status", alert.Status)
					fmt.Println()
				}
			} else {
				printAsJSON(alerts)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(alertsCmd)
}
