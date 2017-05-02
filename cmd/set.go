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

	composeAPI "github.com/compose/gocomposeapi"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set scale for a deployment",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal("Need a deployment id and new units value")
		}
		c := getComposeAPI()

		params := composeAPI.ScalingsParams{}
		params.DeploymentID = args[0]
		scaleval, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("Scale units must be integer")
		}

		params.Units = scaleval
		recipe, errs := c.SetScalings(params)
		bailOnErrs(errs)
		if !outputJSON {
			printRecipe(*recipe)
			fmt.Println()
		} else {
			printAsJSON(*recipe)
		}
	},
}

func init() {
	scaleCmd.AddCommand(setCmd)
}
