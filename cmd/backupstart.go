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

// backupstartCmd represents the backups start command
var backupstartCmd = &cobra.Command{
	Use:   "start [deployment id]",
	Short: "Start backups for a deployment",
	Long:  `Start an on-demand backups for a deployment. Will return the recipe performing the backup.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if len(args) == 0 {
			log.Fatal("Need a deployment id")
		}
		if outputRaw {
			text, errs := c.StartBackupForDeploymentJSON(args[0])
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			recipe, errs := c.StartBackupForDeployment(args[0])
			bailOnErrs(errs)

			if !outputJSON {
				printRecipe(*recipe)
			} else {
				printAsJSON(recipe)
			}

			if recipewatch || recipewait {
				watchRecipeTillComplete(c, recipe.ID)
			}
		}
	},
}

func init() {
	backupsCmd.AddCommand(backupstartCmd)
}
