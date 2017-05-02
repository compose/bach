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

	"github.com/spf13/cobra"
)

// clustersCmd represents the clusters command
var clustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "Show clusters",
	Long:  `Lists any Compose Enterprise Clusters associated with this account.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if outputRaw {
			text, errs := c.GetClustersJSON()
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			clusters, errs := c.GetClusters()
			bailOnErrs(errs)

			if !outputJSON {
				for _, v := range *clusters {
					printCluster(v)
					fmt.Println()
				}
			} else {
				printAsJSON(clusters)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(clustersCmd)
}
