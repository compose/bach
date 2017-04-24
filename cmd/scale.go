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

// scaleCmd represents the scale command
var scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Show scale information for a deployment",
	Long:  `Show Scale information (including unit size) for a deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Need deployment id")
		}
		c := getComposeAPI()
		if outputRaw {
			text, errs := c.GetScalingsJSON(args[0])
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			scalings, errs := c.GetScalings(args[0])
			bailOnErrs(errs)

			if !outputJSON {
				printScalings(*scalings)
				fmt.Println()
			} else {
				printAsJSON(*scalings)
			}

		}
	},
}

func init() {
	RootCmd.AddCommand(scaleCmd)
}
