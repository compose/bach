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

// backupgetCmd represents the backups get command
var backupgetCmd = &cobra.Command{
	Use:   "get [deployment id] [backup id]",
	Short: "Show Backup details for deployment",
	Long:  `Show the backup details for a deployment's backup`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if len(args) != 2 {
			log.Fatal("Need a deployment id and a backup id")
		}
		if outputRaw {
			text, errs := c.GetBackupDetailsForDeploymentJSON(args[0], args[1])
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			backup, errs := c.GetBackupDetailsForDeployment(args[0], args[1])
			bailOnErrs(errs)

			if !outputJSON {
				fmt.Printf("%15s: %s\n", "Backup ID", backup.ID)
				fmt.Printf("%15s: %s\n", "Deployment ID", backup.Deploymentid)
				fmt.Printf("%15s: %s\n", "Backup Name", backup.Name)
				fmt.Printf("%15s: %s\n", "Type", backup.Type)
				fmt.Printf("%15s: %s\n", "Status", backup.Status)
				fmt.Printf("%15s: %s\n", "Download Link", backup.DownloadLink)
			} else {
				printAsJSON(backup)
			}
		}
	},
}

func init() {
	backupsCmd.AddCommand(backupgetCmd)
}
