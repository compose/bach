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
	"net/url"
	"path/filepath"

	"github.com/spf13/cobra"
)

var downloadBackupNow bool

// backupgetCmd represents the backups get command
var backupgetCmd = &cobra.Command{
	Use:   "get [deployment id] [backup id/name]",
	Short: "Show Backup details for deployment",
	Long:  `Show the backup details for a deployment's backup`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if len(args) != 2 {
			log.Fatal("Need a deployment id/name and a backup id")
		}
		depid, err := resolveDepID(c, args[0])
		if err != nil {
			log.Fatal(err)
		}
		backupid := args[1]
		if outputRaw {
			text, errs := c.GetBackupDetailsForDeploymentJSON(depid, backupid)
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			backup, errs := c.GetBackupDetailsForDeployment(depid, backupid)
			bailOnErrs(errs)

			if !outputJSON {
				fmt.Printf("%15s: %s\n", "Backup ID", backup.ID)
				fmt.Printf("%15s: %s\n", "Deployment ID", backup.Deploymentid)
				fmt.Printf("%15s: %s\n", "Backup Name", backup.Name)
				fmt.Printf("%15s: %s\n", "Type", backup.Type)
				fmt.Printf("%15s: %s\n", "Status", backup.Status)
				fmt.Printf("%15s: %s\n", "Download Link", backup.DownloadLink)

				if downloadBackupNow && backup.IsDownloadable {
					u, err := url.Parse(backup.DownloadLink)
					err = DownloadFile(filepath.Base(u.Path), backup.DownloadLink)
					if err != nil {
						fmt.Printf("Error downloading: %s\n", err)
					} else {
						fmt.Printf("Downloaded: %s\n", backup.Name)
					}
				} else if downloadBackupNow && !backup.IsDownloadable {
					fmt.Println("Backup is not downloadable - Download skipped")
				}
			} else {
				printAsJSON(backup)
			}
		}
	},
}

func init() {
	backupsCmd.AddCommand(backupgetCmd)
	backupgetCmd.Flags().BoolVarP(&downloadBackupNow, "download", "d", false, "Download the selected backup")
}
