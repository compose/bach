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
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch [recipe id]",
	Short: "Watch a recipe status",
	Long: `Polls a recipe, given as an id, at 5 second rate (or custom rate)
	
		bach watch 58fe03c15c6efc0014000034 10
		
	.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()

		if len(args) < 1 {
			log.Fatal("Need a recipe id ")
		}
		recipeid := args[0]

		refreshrate := 5

		if len(args) == 2 {
			rate, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatal("Rate must be an integer")
			}
			refreshrate = rate
		}

		recipe, errs := c.GetRecipe(args[0])
		bailOnErrs(errs)

		printRecipe(*recipe)

		if recipe.Status == "complete" {
			return
		}

		for {
			time.Sleep(time.Duration(refreshrate) * time.Second)
			recipe, errs = c.GetRecipe(recipeid)
			bailOnErrs(errs)

			fmt.Println()
			printRecipe(*recipe)

			if recipe.Status == "complete" || recipe.Status == "failed" {
				return
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(watchCmd)
}
