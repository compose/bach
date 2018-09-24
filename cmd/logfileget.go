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

var downloadLogfileNow bool

// logfilegetCmd represents the logfiles get command
var logfilegetCmd = &cobra.Command{
	Use:   "get [deployment id]  logfile id/name]",
	Short: "Show Logfile details for deployment",
	Long:  `Show the Logfile details for a deployment's Logfile`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if len(args) != 2 {
			log.Fatal("Need a deployment id/name and a logfile id")
		}
		depid, err := resolveDepID(c, args[0])
		if err != nil {
			log.Fatal(err)
		}
		logfileid := args[1]
		if outputRaw {
			text, errs := c.GetLogfileDetailsForDeploymentJSON(depid, logfileid)
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			logfile, errs := c.GetLogfileDetailsForDeployment(depid, logfileid)
			bailOnErrs(errs)

			if !outputJSON {
				fmt.Printf("%15s: %s\n", "Logfile ID", logfile.ID)
				fmt.Printf("%15s: %s\n", "Deployment ID", logfile.Deploymentid)
				fmt.Printf("%15s: %s\n", "Capsule ID", logfile.Capsuleid)
				fmt.Printf("%15s: %s\n", "Logfile Name", logfile.Name)
				fmt.Printf("%15s: %s\n", "Status", logfile.Status)
				fmt.Printf("%15s: %s\n", "Download Link", logfile.DownloadLink)

				if downloadLogfileNow {
					err := DownloadFile(logfile.Name, logfile.DownloadLink)
					if err != nil {
						fmt.Printf("Error downloading: %s\n", err)
					} else {
						fmt.Printf("Downloaded: %s\n", logfile.Name)
					}
				}
			} else {
				printAsJSON(logfile)
			}
		}
	},
}

func init() {
	logfilesCmd.AddCommand(logfilegetCmd)
	logfilegetCmd.Flags().BoolVarP(&downloadLogfileNow, "download", "d", false, "Download the selected logfile")
}
