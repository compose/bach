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
	"log"
	//	"log"

	"github.com/spf13/cobra"
)

// teamsaddCmd represents the teams add command
var teamsaddCmd = &cobra.Command{
	Use:   "add [teamid] [userid]",
	Short: "add user to a team",
	Long:  `Add a user to a team`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if len(args) != 2 {
			log.Fatal("Need a team id and user id")
		}
		teamid := args[0]
		userid := args[1]

		team, errs := c.GetTeam(teamid)
		bailOnErrs(errs)

		// Now we have the teams, we have to reduce the users to
		// a string array

		userids := make([]string, len(team.Users))
		for i, user := range team.Users {
			userids[i] = user.ID
		}

		userids = append(userids, userid)

		team, errs = c.PutTeamUsers(teamid, userids)

		printTeam(*team)
	},
}

var teamsremCmd = &cobra.Command{
	Use:   "rem [teamid] [userid]",
	Short: "removes user from a team",
	Long:  `Removes a user from a team`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if len(args) != 2 {
			log.Fatal("Need a team id and user id")
		}
		teamid := args[0]
		userid := args[1]

		team, errs := c.GetTeam(teamid)
		bailOnErrs(errs)

		// Now we have the teams, we have to reduce the users to
		// a string array

		found := false
		for _, user := range team.Users {
			if user.ID == userid {
				found = true
			}
		}

		if !found {
			log.Fatal("User is not member of given team")
		}

		userids := make([]string, len(team.Users)-1)
		i := 0
		for _, user := range team.Users {
			if user.ID != userid {
				userids[i] = user.ID
				i = i + 1
			}
		}

		team, errs = c.PutTeamUsers(teamid, userids)

		printTeam(*team)
	},
}

var teamsUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Commands for team users",
	Long:  `A selection of subcommands are available for listing, modifying and deleting team users`,
}

func init() {
	teamsCmd.AddCommand(teamsUserCmd)
	teamsUserCmd.AddCommand(teamsaddCmd)
	teamsUserCmd.AddCommand(teamsremCmd)
}
