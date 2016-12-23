// Copyright 2016 Compose, an IBM Company
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/compose/gocomposeapi"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"strings"
)

var (
	app = kingpin.New("bach", "A Compose CLI application")

	rawmodeflag = app.Flag("raw", "Output raw JSON responses").Default("false").Bool()
	formatflag  = app.Flag("fmt", "Format output for readability").Default("false").Bool()
	fullcaflag  = app.Flag("fullca", "Show all of CA Certificates").Default("false").Bool()

	showcmd            = app.Command("show", "Show attribute")
	showaccountcmd     = showcmd.Command("account", "Show account details")
	showdeploymentscmd = showcmd.Command("deployments", "Show deployments")
	showrecipecmd      = showcmd.Command("recipe", "Show recipe")
	showrecipeid       = showrecipecmd.Arg("recid", "Recipe ID").String()

	showdeploymentcmd         = showcmd.Command("deployment", "Show deployment")
	showdeploymentrecipescmd  = showdeploymentcmd.Command("recipes", "Show deployment recipes")
	showdeploymentversionscmd = showdeploymentcmd.Command("versions", "Show version and upgrades")
	showrecipesdepid          = showdeploymentrecipescmd.Arg("depid", "Deployment ID").String()
	showversionsdepid         = showdeploymentversionscmd.Arg("depid", "Deployment ID").String()

	showrecipescmd  = showcmd.Command("recipes", "Show recipes for a deployment")
	showclusterscmd = showcmd.Command("clusters", "Show available clusters")
	showuser        = showcmd.Command("user", "Show current associated user")

	showdatacenters = showcmd.Command("datacenters", "Show available datacenters")
	showdatabases   = showcmd.Command("databases", "Show available database types")

	createcmd                  = app.Command("create", "Create...")
	createdeploymentcmd        = createcmd.Command("deployment", "Create deployment")
	createdeploymentname       = createdeploymentcmd.Arg("name", "New Deployment Name").String()
	createdeploymenttype       = createdeploymentcmd.Arg("type", "New Deployment Type").String()
	createdeploymentcluster    = createdeploymentcmd.Flag("cluster", "Cluster ID").String()
	createdeploymentdatacenter = createdeploymentcmd.Flag("datacenter", "Datacenter location").String()

	apitoken = os.Getenv("COMPOSEAPITOKEN")
)

const (
	apibase = "https://api.compose.io/2016-07/"
)

func bailOnErrs(errs []error) {
	if errs != nil {
		log.Fatal(errs)
	}
}

func printAsJSON(toprint interface{}) {
	jsonstr, _ := json.MarshalIndent(toprint, "", " ")
	fmt.Println(string(jsonstr))
}
func main() {
	if apitoken == "" {
		log.Fatal("COMPOSEAPITOKEN environment variable not set")
	}

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case "show account":
		showAccount()
	case "show deployments":
		showDeployments()
	case "show deployment recipes":
		showRecipes()
	case "show deployment versions":
		showVersions()
	case "show recipe":
		showRecipe()
	case "show clusters":
		showClusters()
	case "show user":
		showUser()
	case "show datacenters":
		showDatacenters()
	case "show databases":
		showDatabases()
	case "create deployment":
		createDeployment()
	}
}

func showAccount() {
	if *rawmodeflag {
		text, errs := composeapi.GetAccountJSON()
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		account, errs := composeapi.GetAccount()
		bailOnErrs(errs)

		if *formatflag {
			fmt.Printf("%15s: %s\n", "ID", account.ID)
			fmt.Printf("%15s: %s\n", "Name", account.Name)
			fmt.Printf("%15s: %s\n", "Slug", account.Slug)
			fmt.Println()
		} else {
			printAsJSON(account)
		}
	}
}

func showDeployments() {
	if *rawmodeflag {
		text, errs := composeapi.GetDeploymentsJSON()
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		deployments, errs := composeapi.GetDeployments()
		bailOnErrs(errs)

		if *formatflag {
			for _, v := range *deployments {
				fmt.Printf("%15s: %s\n", "ID", v.ID)
				fmt.Printf("%15s: %s\n", "Name", v.Name)
				fmt.Printf("%15s: %s\n", "Type", v.Type)
				fmt.Printf("%15s: %s\n", "Created At", v.CreatedAt)
				fmt.Printf("%15s: %s\n", "Web UI Link", getLink(v.Links.ComposeWebUILink))
				fmt.Println()
			}
		} else {
			printAsJSON(deployments)
		}
	}
}

func showRecipe() {
	if *rawmodeflag {
		text, errs := composeapi.GetRecipeJSON(*showrecipeid)
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		recipe, errs := composeapi.GetRecipe(*rawmodeflag, *showrecipeid)
		bailOnErrs(errs)

		if *formatflag {
			printRecipe(*recipe)
		} else {
			printAsJSON(*recipe)
		}
	}
}

func showRecipes() {
	if *rawmodeflag {
		fmt.Println(*showrecipesdepid)
		text, errs := composeapi.GetRecipesForDeploymentJSON(*showrecipesdepid)
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		recipes, errs := composeapi.GetRecipesForDeployment(*showrecipesdepid)
		bailOnErrs(errs)
		if *formatflag {
			for _, v := range *recipes {
				printRecipe(v)
				fmt.Println()
			}
		} else {
			printAsJSON(*recipes)
		}
	}
}

func showVersions() {
	if *rawmodeflag {
		text, errs := composeapi.GetVersionsForDeploymentJSON(*showversionsdepid)
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		versions, errs := composeapi.GetVersionsForDeployment(*showversionsdepid)
		bailOnErrs(errs)
		if *formatflag {
			for _, v := range *versions {
				printVersionTransitions(v)
				fmt.Println()
			}
		} else {
			printAsJSON(*versions)
		}
	}
}

func showClusters() {
	if *rawmodeflag {
		text, errs := composeapi.GetClustersJSON()
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		clusters, errs := composeapi.GetClusters()
		bailOnErrs(errs)

		if *formatflag {
			for _, v := range *clusters {
				printCluster(v)
				fmt.Println()
			}
		} else {
			printAsJSON(clusters)
		}
	}
}

func showUser() {
	if *rawmodeflag {
		text, errs := composeapi.GetUserJSON()
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		user, errs := composeapi.GetUser()
		bailOnErrs(errs)
		if *formatflag {
			fmt.Printf("%15s: %s\n", "ID", user.ID)
			fmt.Println()
		} else {
			printAsJSON(user)
		}
	}
}

func showDatacenters() {
	if *rawmodeflag {
		text, errs := composeapi.GetDatacentersJSON()
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		datacenters, errs := composeapi.GetDatacenters()
		bailOnErrs(errs)

		if *formatflag {
			for _, v := range *datacenters {
				printDatacenter(v)
				fmt.Println()
			}
		} else {
			printAsJSON(datacenters)
		}
	}
}

func showDatabases() {
	if *rawmodeflag {
		text, errs := composeapi.GetDatabasesJSON()
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		databases, errs := composeapi.GetDatabases()
		bailOnErrs(errs)

		if *formatflag {
			for _, v := range *databases {
				printDatabase(v)
				fmt.Println()
			}
		} else {
			printAsJSON(databases)
		}
	}
}

func createDeployment() {
	if *rawmodeflag {
		log.Fatal("Raw mode not supported for createDeployment")
	}

	account, errs := composeapi.GetAccount()
	bailOnErrs(errs)

	if *createdeploymentdatacenter == "" && *createdeploymentcluster == "" {
		log.Fatal("Must supply either a --cluster id or --datacenter region")
	}

	params := composeapi.CreateDeploymentParams{
		Name:         *createdeploymentname,
		AccountID:    account.ID,
		DatabaseType: *createdeploymenttype,
		Datacenter:   *createdeploymentdatacenter,
		ClusterID:    *createdeploymentcluster,
	}

	deployment, errs := composeapi.CreateDeployment(params)
	bailOnErrs(errs)

	if *formatflag {
		printDeployment(*deployment)
	} else {
		printAsJSON(*deployment)
	}

}
func getLink(link composeapi.Link) string {
	return strings.Replace(link.HREF, "{?embed}", "", -1) // TODO: This should mangle the HREF properly
}

func printRecipe(recipe composeapi.Recipe) {
	fmt.Printf("%15s: %s\n", "ID", recipe.ID)
	fmt.Printf("%15s: %s\n", "Template", recipe.Template)
	fmt.Printf("%15s: %s\n", "Status", recipe.Status)
	fmt.Printf("%15s: %s\n", "Status Detail", recipe.StatusDetail)
	fmt.Printf("%15s: %s\n", "Account ID", recipe.AccountID)
	fmt.Printf("%15s: %s\n", "Deployment ID", recipe.DeploymentID)
	fmt.Printf("%15s: %s\n", "Name", recipe.Name)
	fmt.Printf("%15s: %d\n", "Child Recipes", len(recipe.Embedded.Recipes))

}

func printVersionTransitions(version composeapi.VersionTransition) {
	fmt.Printf("%15s: %s\n", "Application", version.Application)
	fmt.Printf("%15s: %s\n", "Method", version.Method)
	fmt.Printf("%15s: %s\n", "From Version", version.FromVersion)
	fmt.Printf("%15s: %s\n", "To Version", version.ToVersion)
}
func printCluster(cluster composeapi.Cluster) {
	fmt.Printf("%15s: %s\n", "ID", cluster.ID)
	fmt.Printf("%15s: %s\n", "Account ID", cluster.AccountID)
	fmt.Printf("%15s: %s\n", "Account Slug", cluster.AccountSlug)
	fmt.Printf("%15s: %s\n", "Name", cluster.Name)
	fmt.Printf("%15s: %s\n", "Type", cluster.Type)
	fmt.Printf("%15s: %t\n", "Multitenant", cluster.Multitenant)
	fmt.Printf("%15s: %s\n", "Provider", cluster.Provider)
	fmt.Printf("%15s: %s\n", "Region", cluster.Region)
	fmt.Printf("%15s: %s\n", "Created Ad", cluster.CreatedAt)
	fmt.Printf("%15s: %s\n", "Subdomain", cluster.Subdomain)
}

func printDatacenter(datacenter composeapi.Datacenter) {
	fmt.Printf("%15s: %s\n", "Region", datacenter.Region)
	fmt.Printf("%15s: %s\n", "Provider", datacenter.Provider)
	fmt.Printf("%15s: %s\n", "Slug", datacenter.Slug)
}

func printDeployment(deployment composeapi.Deployment) {
	fmt.Printf("%15s: %s\n", "ID", deployment.ID)
	fmt.Printf("%15s: %s\n", "Name", deployment.Name)
	fmt.Printf("%15s: %s\n", "Type", deployment.Type)
	fmt.Printf("%15s: %s\n", "Created At", deployment.CreatedAt)
	if deployment.ProvisionRecipeID != "" {
		fmt.Printf("%15s: %s\n", "Prov Recipe ID", deployment.ProvisionRecipeID)
	}
	if deployment.CACertificateBase64 != "" {
		if *fullcaflag {
			fmt.Printf("%15s: %s\n", "CA Certificate", deployment.CACertificateBase64)
		} else {
			fmt.Printf("%15s: %s...\n", "CA Certificate", deployment.CACertificateBase64[0:32])
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

func printDatabase(database composeapi.Database) {
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
