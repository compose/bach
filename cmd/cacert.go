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
	"encoding/base64"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// cacertCmd represents the cacert command
var cacertCmd = &cobra.Command{
	Use:   "cacert [deployment id/name]",
	Short: "Returns the self-signed cert for the deployment",
	Long:  `Returns the self-signed cert for the deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Need a deployment id/name")
		}
		c := getComposeAPI()
		depid, err := resolveDepID(c, args[0])
		if err != nil {
			log.Fatal(err)
		}
		deployment, errs := c.GetDeployment(depid)
		bailOnErrs(errs)
		decoded, err := base64.StdEncoding.DecodeString(deployment.CACertificateBase64)
		if err != nil {
			log.Fatal("Problem decoding cacert")
		}

		fmt.Printf("%s\n", decoded)
	},
}

func init() {
	RootCmd.AddCommand(cacertCmd)
}
