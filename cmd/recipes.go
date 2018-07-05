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

// recipesCmd represents the recipes command
var recipesCmd = &cobra.Command{
	Use:   "recipes [deployment id]",
	Short: "Show Recipes related to deployment",
	Long:  `Print out all the recipes that have been invoked on a particular deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Need a deployment id to show recipes for")
		}
		c := getComposeAPI()
		depid, err := resolveDepID(c, args[0])
		if err != nil {
			log.Fatal(err)
		}
		if outputRaw {
			text, errs := c.GetRecipesForDeploymentJSON(depid)
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			recipes, errs := c.GetRecipesForDeployment(depid)
			bailOnErrs(errs)
			if !outputJSON {
				for _, v := range *recipes {
					printRecipe(v)
					fmt.Println()
				}
			} else {
				printAsJSON(*recipes)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(recipesCmd)
}
