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

// backuplistCmd represents the backups list command
var backuplistCmd = &cobra.Command{
	Use:   "list [deployment id/name]",
	Short: "Show Backups for deployment",
	Long:  `Show the backups for a deployment`,
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
			text, errs := c.GetBackupsForDeploymentJSON(depid)
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			backups, errs := c.GetBackupsForDeployment(depid)
			bailOnErrs(errs)

			if !outputJSON {
				for _, backup := range *backups {
					fmt.Printf("%15s: %s\n", "Backup ID", backup.ID)
					fmt.Printf("%15s: %s\n", "Deployment ID", backup.Deploymentid)
					fmt.Printf("%15s: %s\n", "Backup Name", backup.Name)
					fmt.Printf("%15s: %s\n", "Type", backup.Type)
					fmt.Printf("%15s: %s\n", "Status", backup.Status)
					fmt.Println()
				}
			} else {
				printAsJSON(backups)
			}
		}
	},
}

func init() {
	backupsCmd.AddCommand(backuplistCmd)
}
