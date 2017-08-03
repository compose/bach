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
	//	"log"

	"github.com/spf13/cobra"
)

// teamslistCmd represents the alerts command
var teamslistCmd = &cobra.Command{
	Use:   "list",
	Short: "Show teams",
	Long:  `Show the teams for an account`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		// if len(args) == 0 {
		// 	log.Fatal("Need a deployment id")
		// }
		if outputRaw {
			text, errs := c.GetTeamsJSON()
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			teams, errs := c.GetTeams()
			bailOnErrs(errs)

			if !outputJSON {
				for _, team := range *teams {
					printTeam(team)
					fmt.Println()
				}
			} else {
				printAsJSON(teams)
			}
		}
	},
}

func init() {
	teamsCmd.AddCommand(teamslistCmd)
}
