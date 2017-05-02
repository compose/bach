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

// recipeCmd represents the recipe command
var recipeCmd = &cobra.Command{
	Use:   "recipe",
	Short: "Show details of a recipe",
	Long:  `Prints the details of a recipe based on its given id.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Need a recipe id to show")
		}
		c := getComposeAPI()
		if outputRaw {
			text, errs := c.GetRecipeJSON(args[0])
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			recipe, errs := c.GetRecipe(args[0])
			bailOnErrs(errs)

			if !outputJSON {
				printRecipe(*recipe)
			} else {
				printAsJSON(*recipe)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(recipeCmd)
}
