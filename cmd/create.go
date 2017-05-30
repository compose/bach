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

var datacenterid string
var clusterid string
var version string
var ssl bool
var wiredtiger bool

// createCmd represents the deployment command
var createCmd = &cobra.Command{
	Use:   "create [deployment name] [database type]",
	Short: "Create a deployment",
	Long:  `Creates a deployment. Requires deployment name and database type.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if outputRaw {
			log.Fatal("Raw mode not supported for createDeployment")
		}

		if len(args) != 2 {
			log.Fatal("Need deployment name and deployment type")
		}

		account, errs := c.GetAccount()
		bailOnErrs(errs)

		if datacenterid == "" && clusterid == "" {
			log.Fatal("Must supply either a --cluster id or --datacenter slug")
		}

		deploymentname := args[0]
		dbtype := args[1]

		params :=
			composeAPI.DeploymentParams{
				Name:         deploymentname,
				AccountID:    account.ID,
				DatabaseType: dbtype,
				Datacenter:   datacenterid,
				ClusterID:    clusterid,
				Version:      version,
				SSL:          ssl,
				WiredTiger:   wiredtiger,
			}

		deployment, errs := c.CreateDeployment(params)
		bailOnErrs(errs)

		if !outputJSON {
			printDeployment(*deployment)
		} else {
			printAsJSON(*deployment)
		}

		if recipewatch {
			watchRecipeTillComplete(c, deployment.ProvisionRecipeID)
		}
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVar(&clusterid, "cluster", "", "Cluster Id")
	createCmd.Flags().StringVar(&datacenterid, "datacenter", "", "Datacenter region")
	createCmd.Flags().StringVar(&version, "version", "", "Database version required")
	createCmd.Flags().BoolVar(&ssl, "ssl", false, "SSL required (where supported)")
	createCmd.Flags().BoolVar(&wiredtiger, "wiredtiger", false, "Use WiredTiger storage (MongoDB only)")
}
