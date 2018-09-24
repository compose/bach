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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	composeAPI "github.com/compose/gocomposeapi"
	"github.com/dustin/go-humanize"
)

func getComposeAPI() (client *composeAPI.Client) {
	if apiToken == "Your API Token" {
		ostoken := os.Getenv("COMPOSEAPITOKEN")
		if ostoken == "" {
			log.Fatal("Token not set and COMPOSEAPITOKEN environment variable not set - Get your token at https://app.compose.io/oauth/api_tokens")
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

func resolveDepID(client *composeAPI.Client, arg string) (depid string, err error) {
	// Test for being just deployment id
	if len(arg) == 24 && isHexString(arg) {
		return arg, nil
	}

	// Get all the deployments and search
	deployments, errs := client.GetDeployments()

	if errs != nil {
		bailOnErrs(errs)
		return "", errs[0]
	}

	for _, deployment := range *deployments {
		if deployment.Name == arg {
			return deployment.ID, nil
		}
	}

	return "", fmt.Errorf("deployment not found: %s", arg)
}

func isHexString(s string) bool {
	_, err := hex.DecodeString(s)
	if err == nil {
		return true
	}
	return false
}

func watchRecipeTillComplete(client *composeAPI.Client, recipeid string) {
	var lastRecipe *composeAPI.Recipe

	for {
		time.Sleep(time.Duration(5) * time.Second)
		recipe, errs := client.GetRecipe(recipeid)
		bailOnErrs(errs)

		if lastRecipe == nil {
			lastRecipe = recipe
			if !recipewait {
				fmt.Println()
				printShortRecipe(*recipe)
			}
		} else {
			if lastRecipe.Status == recipe.Status &&
				lastRecipe.UpdatedAt == recipe.UpdatedAt &&
				lastRecipe.StatusDetail == recipe.StatusDetail {
				if !recipewait {
					fmt.Print(".")
				}
			} else {
				lastRecipe = recipe
				if !recipewait {
					fmt.Println()
					printShortRecipe(*recipe)
				}
			}
		}

		if recipe.Status == "complete" || recipe.Status == "failed" {
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

//WriteCounter is a type to track a download in progress
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

//PrintProgress will display the progress of the download
func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

//DownloadFile is a utility based on the example at https://golangcode.com/download-a-file-with-progress/
func DownloadFile(filepath string, url string) error {

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")

	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		return err
	}

	return nil
}
