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
	"regexp"
	"sort"
	"strings"

	composeAPI "github.com/compose/gocomposeapi"

	"github.com/spf13/cobra"
)

var listdbtype string
var listfilter string
var bytype bool
var byname bool
var byid bool

//ByDBType sorts by type
type ByDBType []composeAPI.Deployment

func (a ByDBType) Len() int           { return len(a) }
func (a ByDBType) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDBType) Less(i, j int) bool { return a[i].Type < a[j].Type }

//ByName sorts by name
type ByName []composeAPI.Deployment

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return strings.ToLower(a[i].Name) < strings.ToLower(a[j].Name) }

//ByID sorts by ID
type ByID []composeAPI.Deployment

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }

// listCmd represents the new list deployments command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List deployments attached to account",
	Long:  `List the database deployments attached to an account.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := getComposeAPI()
		if outputRaw {
			text, errs := c.GetDeploymentsJSON()
			bailOnErrs(errs)
			fmt.Println(text)
		} else {
			deployments, errs := c.GetDeployments()
			bailOnErrs(errs)

			matcher, err := regexp.Compile(listfilter)
			if err != nil {
				log.Fatal(err)
			}

			if !outputJSON {
				fmt.Printf("%-24s %-14s %-40s\n", "ID", "Type", "Name")
				fmt.Printf("%-24s %-14s %-40s\n", strings.Repeat("-", 24), strings.Repeat("-", 14), strings.Repeat("-", 40))

				if bytype {
					sort.Sort(ByDBType(*deployments))
				} else if byid {
					sort.Sort(ByID(*deployments))
				} else if byname {
					sort.Sort(ByName(*deployments))
				}

				for _, v := range *deployments {
					if (listdbtype == "" || listdbtype == v.Type) && matcher.MatchString(v.Name) {
						fmt.Printf("%-24s %-14s %-40s\n", v.ID, v.Type, v.Name)
					}
				}
			} else {
				printAsJSON(deployments)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&listdbtype, "type", "t", "", "Only this database type")
	listCmd.Flags().StringVarP(&listfilter, "filter", "f", "", "Regular expression to filter names on")
	listCmd.Flags().BoolVarP(&byname, "sortname", "n", false, "Sort by name")
	listCmd.Flags().BoolVarP(&bytype, "sorttype", "d", false, "Sort by type")
	listCmd.Flags().BoolVarP(&byid, "sortid", "i", false, "Sort by id")
}
