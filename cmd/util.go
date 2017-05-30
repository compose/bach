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
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	composeAPI "github.com/compose/gocomposeapi"
)

func getComposeAPI() (client *composeAPI.Client) {
	if apiToken == "Your API Token" {
		ostoken := os.Getenv("COMPOSEAPITOKEN")
		if ostoken == "" {
			log.Fatal("Token not set and COMPOSEAPITOKEN environment variable not set")
		}
		apiToken = ostoken
	}

	var err error
	client, err = composeAPI.NewClient(apiToken)
	if err != nil {
		log.Fatalf("Could not create compose client: %s", err.Error())
	}
	return client
}

func watchRecipeTillComplete(client *composeAPI.Client, recipeid string) {
	var lastRecipe *composeAPI.Recipe

	for {
		time.Sleep(time.Duration(5) * time.Second)
		recipe, errs := client.GetRecipe(recipeid)
		bailOnErrs(errs)

		if lastRecipe == nil {
			lastRecipe = recipe
			fmt.Println()
			printShortRecipe(*recipe)
		} else {
			if lastRecipe.Status == recipe.Status &&
				lastRecipe.UpdatedAt == recipe.UpdatedAt &&
				lastRecipe.StatusDetail == recipe.StatusDetail {
				fmt.Print(".")
			} else {
				lastRecipe = recipe
				fmt.Println()
				printShortRecipe(*recipe)
			}
		}

		if recipe.Status == "complete" {
			return
		}
	}
}

func bailOnErrs(errs []error) {
	if errs != nil {
		log.Fatal(errs)
	}
}

func printAsJSON(toprint interface{}) {
	jsonstr, _ := json.MarshalIndent(toprint, "", " ")
	fmt.Println(string(jsonstr))
}

func getLink(link composeAPI.Link) string {
	return strings.Replace(link.HREF, "{?embed}", "", -1) // TODO: This should mangle the HREF properly
}

var savedVersion string

//SaveVersion called from outside to retain version string
func SaveVersion(version string) {
	savedVersion = version
}

func getVersion() string {
	return savedVersion
}
