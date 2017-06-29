// Copyright © 2017 Compose, an IBM company
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

	composeAPI "github.com/compose/gocomposeapi"
	"github.com/spf13/cobra"
)

// setNotesCmd represents the set command
var setNotesCmd = &cobra.Command{
	Use:   "notes [deployment id] [note]",
	Short: "Set notes for a deployment",
	Long:  `Sets the customer notes for a deployment.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal("Need a deployment id and new notes value")
		}
		c := getComposeAPI()

		params := composeAPI.PatchDeploymentParams{}
		params.DeploymentID = args[0]
		params.Notes = args[1]
		deployment, errs := c.PatchDeployment(params)

		bailOnErrs(errs)
		if !outputJSON {
			printDeployment(*deployment)
			fmt.Println()
		} else {
			printAsJSON(*deployment)
		}
	},
}

func init() {
	setCmd.AddCommand(setNotesCmd)
}
