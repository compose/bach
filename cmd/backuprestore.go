// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"log"

	composeAPI "github.com/compose/gocomposeapi"
	"github.com/spf13/cobra"
)

var restoreclusterid string
var restoredatacenterid string
var restoressl bool

// backuprestoreCmd represents the backups restore command
var backuprestoreCmd = &cobra.Command{
	Use:   "restore [deployment id/name] [backup id] [new deployment name]",
	Short: "Restore a deployment",
	Long:  `Restores a deployment. Requires deployment id/name, backup id, and new deployment name.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if outputRaw {
			log.Fatal("Raw mode not supported for backup restore")
		}

		if len(args) != 3 {
			log.Fatal("Need deployment id/name, backup id and new deployment name")
		}

		if restoredatacenterid == "" && restoreclusterid == "" {
			log.Fatal("Must supply either a --cluster id or --datacenter region")
		}

		deploymentid, err := resolveDepID(c, args[0])
		if err != nil {
			log.Fatal(err)
		}
		backupid := args[1]
		deploymentname := args[2]

		params := composeAPI.RestoreBackupParams{
			DeploymentID: deploymentid,
			BackupID:     backupid,
			Name:         deploymentname,
			Datacenter:   restoredatacenterid,
			ClusterID:    restoreclusterid,
			SSL:          ssl,
		}

		deployment, errs := c.RestoreBackup(params)
		bailOnErrs(errs)

		if !outputJSON {
			printDeployment(*deployment)
		} else {
			printAsJSON(*deployment)
		}

		if recipewatch || recipewait {
			watchRecipeTillComplete(c, deployment.ProvisionRecipeID)
		}

	},
}

func init() {
	backupsCmd.AddCommand(backuprestoreCmd)
	backuprestoreCmd.Flags().StringVar(&restoreclusterid, "cluster", "", "Cluster Id")
	backuprestoreCmd.Flags().StringVar(&restoredatacenterid, "datacenter", "", "Datacenter region")
	backuprestoreCmd.Flags().BoolVar(&ssl, "ssl", false, "SSL required (where supported)")
}
