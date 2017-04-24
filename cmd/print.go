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

	composeAPI "github.com/compose/gocomposeapi"
)

func printRecipe(recipe composeAPI.Recipe) {
	fmt.Printf("%15s: %s\n", "ID", recipe.ID)
	fmt.Printf("%15s: %s\n", "Template", recipe.Template)
	fmt.Printf("%15s: %s\n", "Status", recipe.Status)
	fmt.Printf("%15s: %s\n", "Status Detail", recipe.StatusDetail)
	fmt.Printf("%15s: %s\n", "Account ID", recipe.AccountID)
	fmt.Printf("%15s: %s\n", "Deployment ID", recipe.DeploymentID)
	fmt.Printf("%15s: %s\n", "Name", recipe.Name)
	fmt.Printf("%15s: %s\n", "Created At", recipe.CreatedAt)
	fmt.Printf("%15s: %s\n", "Updated At", recipe.UpdatedAt)
	fmt.Printf("%15s: %d\n", "Child Recipes", len(recipe.Embedded.Recipes))
}

func printVersionTransitions(version composeAPI.VersionTransition) {
	fmt.Printf("%15s: %s\n", "Application", version.Application)
	fmt.Printf("%15s: %s\n", "Method", version.Method)
	fmt.Printf("%15s: %s\n", "From Version", version.FromVersion)
	fmt.Printf("%15s: %s\n", "To Version", version.ToVersion)
}
func printCluster(cluster composeAPI.Cluster) {
	fmt.Printf("%15s: %s\n", "ID", cluster.ID)
	fmt.Printf("%15s: %s\n", "Account ID", cluster.AccountID)
	fmt.Printf("%15s: %s\n", "Account Slug", cluster.AccountSlug)
	fmt.Printf("%15s: %s\n", "Name", cluster.Name)
	fmt.Printf("%15s: %s\n", "Type", cluster.Type)
	fmt.Printf("%15s: %t\n", "Multitenant", cluster.Multitenant)
	fmt.Printf("%15s: %s\n", "Provider", cluster.Provider)
	fmt.Printf("%15s: %s\n", "Region", cluster.Region)
	fmt.Printf("%15s: %s\n", "Created At", cluster.CreatedAt)
	fmt.Printf("%15s: %s\n", "Subdomain", cluster.Subdomain)
}

func printDatacenter(datacenter composeAPI.Datacenter) {
	fmt.Printf("%15s: %s\n", "Region", datacenter.Region)
	fmt.Printf("%15s: %s\n", "Provider", datacenter.Provider)
	fmt.Printf("%15s: %s\n", "Slug", datacenter.Slug)
}

func printDeployment(deployment composeAPI.Deployment) {
	fmt.Printf("%15s: %s\n", "ID", deployment.ID)
	fmt.Printf("%15s: %s\n", "Name", deployment.Name)
	fmt.Printf("%15s: %s\n", "Type", deployment.Type)
	fmt.Printf("%15s: %s\n", "Created At", deployment.CreatedAt)
	if deployment.ProvisionRecipeID != "" {
		fmt.Printf("%15s: %s\n", "Prov Recipe ID", deployment.ProvisionRecipeID)
	}
	if deployment.CACertificateBase64 != "" {
		if showFullCA {
			if noDecodeCA {
				fmt.Printf("%15s: %s\n", "CA Certificate", deployment.CACertificateBase64)
			} else {
				decoded, err := base64.StdEncoding.DecodeString(deployment.CACertificateBase64)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%15s:\n%s", "CA Certificate", decoded)
			}
		} else {
			fmt.Printf("%15s: %s... (Use --fullca for certificate)\n", "CA Certificate", deployment.CACertificateBase64[0:32])
		}
	}
	fmt.Printf("%15s: %s\n", "Web UI Link", getLink(deployment.Links.ComposeWebUILink))
	fmt.Printf("%15s: %s\n", "Health", deployment.Connection.Health)
	fmt.Printf("%15s: %s\n", "SSH", deployment.Connection.SSH)
	fmt.Printf("%15s: %s\n", "Admin", deployment.Connection.Admin)
	fmt.Printf("%15s: %s\n", "SSHAdmin", deployment.Connection.SSHAdmin)
	fmt.Printf("%15s: %s\n", "CLI Connect", deployment.Connection.CLI)
	fmt.Printf("%15s: %s\n", "Direct Connect", deployment.Connection.Direct)

}

func printDatabase(database composeAPI.Database) {
	fmt.Printf("%15s: %s\n", "Type", database.DatabaseType)
	fmt.Printf("%15s: %s\n", "Status", database.Status)

	for _, v := range database.Embedded.Versions {
		if v.Status != "deprecated" {
			if v.Preferred {
				fmt.Printf("%15s: %s (%s) Preferred\n", "Version", v.Version, v.Status)
			} else {
				fmt.Printf("%15s: %s (%s)\n", "Version", v.Version, v.Status)
			}
		}
	}
}

func printScalings(scalings composeAPI.Scalings) {
	fmt.Printf("%15s: %d\n", "Allocated Units", scalings.AllocatedUnits)
	fmt.Printf("%15s: %d\n", "Used Units", scalings.UsedUnits)
	fmt.Printf("%15s: %d\n", "Starting Units", scalings.StartingUnits)
	fmt.Printf("%15s: %d\n", "Minimum Units", scalings.MinimumUnits)
}
