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

// logfilesListCmd represents the logfiles list command
var logfilesListCmd = &cobra.Command{
	Use:   "list [deployment id/name]",
	Short: "Show Logfiles for deployment",
	Long:  `Show the Logfiles for a deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if len(args) == 0 {
			log.Fatal("Need a deployment id/name")
		}
		depid, err := resolveDepID(c, args[0])
		if err != nil {
			log.Fatal(err)
		}
		if outputRaw {
			text, errs := c.GetLogfilesForDeploymentJSON(depid)
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			logfiles, errs := c.GetLogfilesForDeployment(depid)
			bailOnErrs(errs)

			if !outputJSON {
				for _, logfile := range *logfiles {
					fmt.Printf("%15s: %s\n", "Logfile ID", logfile.ID)
					fmt.Printf("%15s: %s\n", "Deployment ID", logfile.Deploymentid)
					fmt.Printf("%15s: %s\n", "Capsule ID", logfile.Capsuleid)
					fmt.Printf("%15s: %s\n", "Logfile Name", logfile.Name)
					fmt.Printf("%15s: %s\n", "Status", logfile.Status)
					fmt.Println()
				}
			} else {
				printAsJSON(logfiles)
			}
		}
	},
}

func init() {
	logfilesCmd.AddCommand(logfilesListCmd)
}
