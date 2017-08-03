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
	//	"log"

	"github.com/spf13/cobra"
)

// teamslistCmd represents the alerts command
var teamsgetCmd = &cobra.Command{
	Use:   "get",
	Short: "get team details",
	Long:  `List the details of a team`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if len(args) == 0 {
			log.Fatal("Need a team id")
		}
		if outputRaw {
			text, errs := c.GetTeamJSON(args[0])
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			team, errs := c.GetTeam(args[0])
			bailOnErrs(errs)

			if !outputJSON {
				printTeam(*team)
			} else {
				printAsJSON(team)
			}
		}
	},
}

func init() {
	teamsCmd.AddCommand(teamsgetCmd)
}
