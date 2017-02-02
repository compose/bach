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
	"encoding/base64"
	"encoding/json"
	"fmt"
	composeAPI "github.com/compose/gocomposeapi"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"strings"
	"time"
)

var (
	app = kingpin.New("bach", "A Compose CLI application")

	outputRawFlag  = app.Flag("raw", "Output raw JSON responses").Default("false").Short('r').Bool()
	outputJSONFlag = app.Flag("json", "Output post-processed JSON results").Default("false").Short('j').Bool()
	showFullCAFlag = app.Flag("fullca", "Show all of CA Certificates").Default("false").Short('f').Bool()
	noDecodeCAFlag = app.Flag("nodecodeca", "Do not Decode base64 CA Certificates").Default("false").Bool()
	apiToken       = app.Flag("token", "Set API token").Default("yourAPItoken").OverrideDefaultFromEnvar("COMPOSEAPITOKEN").Short('t').String()

	accountCmd = app.Command("account", "Show account details")

	deploymentsCmd = app.Command("deployments", "Show deployments")

	recipeCmd = app.Command("recipe", "Show recipe")
	recipeID  = recipeCmd.Arg("recid", "Recipe ID").String()

	recipesCmd = app.Command("recipes", "Show deployment recipes")
	recipesID  = recipesCmd.Arg("deployment id", "Deployment ID").Required().String()

	versionsCmd          = app.Command("versions", "Show version and upgrades")
	versionsDeploymentID = versionsCmd.Arg("deployment id", "Deployment ID").Required().String()

	detailsCmd          = app.Command("details", "Show deployment information")
	detailsDeploymentID = detailsCmd.Arg("deployment id", "Deployment ID").Required().String()

	scaleCmd          = app.Command("scale", "Get scale of deployment")
	scaleDeploymentID = scaleCmd.Arg("deployment id", "Deployment ID").Required().String()

	clustersCmd    = app.Command("clusters", "Show available clusters")
	userCmd        = app.Command("user", "Show current associated user")
	datacentersCmd = app.Command("datacenters", "Show available datacenters")
	databasesCmd   = app.Command("databases", "Show available database types")

	createCmd                  = app.Command("create", "Create deployment")
	createdeploymentname       = createCmd.Arg("name", "New Deployment Name").String()
	createdeploymenttype       = createCmd.Arg("type", "New Deployment Type").String()
	createdeploymentcluster    = createCmd.Flag("cluster", "Cluster ID").String()
	createdeploymentdatacenter = createCmd.Flag("datacenter", "Datacenter location").String()

	setCmd               = app.Command("set", "Set...")
	setScaleCmd          = setCmd.Command("scale", "Set scale of deployment")
	setScaleDeploymentID = setScaleCmd.Arg("deployment id", "Deployment ID to scale").Required().String()
	setScaleLevel        = setScaleCmd.Arg("units", "New scale units").Required().Int()

	watchCmd         = app.Command("watch", "Watch recipe")
	watcRecipeCmd    = watchCmd.Arg("recipe id", "recipeid").Required().String()
	watchRefreshRate = watchCmd.Flag("refresh", "Refresh rate in seconds").Default("10").Int()

	//apiToken = os.Getenv("COMPOSEAPITOKEN")
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

	parsed := kingpin.MustParse(app.Parse(os.Args[1:]))
	var client *composeAPI.Client

	if *apiToken == "yourAPItoken" {
		log.Fatal("Token not set and COMPOSEAPITOKEN environment variable not set")
	} else {
		var err error
		client, err = composeAPI.NewClient(*apiToken)
		if err != nil {
			log.Fatalf("Could not create compose client: %s", err.Error())
		}
	}

	switch parsed {
	case "account":
		showAccount(client)
	case "deployments":
		showDeployments(client)
	case "recipes":
		showRecipes(client)
	case "versions":
		showVersions(client)
	case "details":
		showDeployment(client)
	case "recipe":
		showRecipe(client)
	case "clusters":
		showClusters(client)
	case "user":
		showUser(client)
	case "datacenters":
		showDatacenters(client)
	case "databases":
		showDatabases(client)
	case "create":
		createDeployment(client)
	case "watch":
		watchRecipe(client)
	case "scale":
		getScaleDeployment(client)
	case "set scale":
		setScaleDeployment(client)
	}
}

func showAccount(c *composeAPI.Client) {
	if *outputRawFlag {
		text, errs := c.GetAccountJSON()
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		account, errs := c.GetAccount()
		bailOnErrs(errs)

		if !*outputJSONFlag {
			fmt.Printf("%15s: %s\n", "ID", account.ID)
			fmt.Printf("%15s: %s\n", "Name", account.Name)
			fmt.Printf("%15s: %s\n", "Slug", account.Slug)
			fmt.Println()
		} else {
			printAsJSON(account)
		}
	}
}

func showDeployments(c *composeAPI.Client) {
	if *outputRawFlag {
		text, errs := c.GetDeploymentsJSON()
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		deployments, errs := c.GetDeployments()
		bailOnErrs(errs)

		if !*outputJSONFlag {
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

func showDeployment(c *composeAPI.Client) {
	if *outputRawFlag {
		text, errs := c.GetDeploymentJSON(*detailsDeploymentID)
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		deployment, errs := c.GetDeployment(*detailsDeploymentID)
		bailOnErrs(errs)

		if !*outputJSONFlag {
			printDeployment(*deployment)
		} else {
			printAsJSON(*deployment)
		}
	}
}

func showRecipe(c *composeAPI.Client) {
	if *outputRawFlag {
		text, errs := c.GetRecipeJSON(*recipeID)
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		recipe, errs := c.GetRecipe(*recipeID)
		bailOnErrs(errs)

		if !*outputJSONFlag {
			printRecipe(*recipe)
		} else {
			printAsJSON(*recipe)
		}
	}
}

func watchRecipe(c *composeAPI.Client) {
	recipe, errs := c.GetRecipe(*watcRecipeCmd)
	bailOnErrs(errs)
	printRecipe(*recipe)

	if recipe.Status == "complete" {
		return
	}

	for {
		time.Sleep(time.Duration(*watchRefreshRate) * time.Second)
		recipe, errs = c.GetRecipe(*watcRecipeCmd)
		bailOnErrs(errs)

		fmt.Println()
		printRecipe(*recipe)

		if recipe.Status == "complete" {
			return
		}
	}
}

func showRecipes(c *composeAPI.Client) {
	if *outputRawFlag {
		text, errs := c.GetRecipesForDeploymentJSON(*recipesID)
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		recipes, errs := c.GetRecipesForDeployment(*recipesID)
		bailOnErrs(errs)
		if !*outputJSONFlag {
			for _, v := range *recipes {
				printRecipe(v)
				fmt.Println()
			}
		} else {
			printAsJSON(*recipes)
		}
	}
}

func showVersions(c *composeAPI.Client) {
	if *outputRawFlag {
		text, errs := c.GetVersionsForDeploymentJSON(*versionsDeploymentID)
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		versions, errs := c.GetVersionsForDeployment(*versionsDeploymentID)
		bailOnErrs(errs)
		if !*outputJSONFlag {
			for _, v := range *versions {
				printVersionTransitions(v)
				fmt.Println()
			}
		} else {
			printAsJSON(*versions)
		}
	}
}

func getScaleDeployment(c *composeAPI.Client) {
	if *outputRawFlag {
		text, errs := c.GetScalingsJSON(*scaleDeploymentID)
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		scalings, errs := c.GetScalings(*scaleDeploymentID)
		bailOnErrs(errs)

		if !*outputJSONFlag {
			printScalings(*scalings)
			fmt.Println()
		} else {
			printAsJSON(*scalings)
		}

	}
}

func setScaleDeployment(c *composeAPI.Client) {
	params := composeAPI.ScalingsParams{}
	params.DeploymentID = *setScaleDeploymentID
	params.Deployment.Units = *setScaleLevel

	recipe, errs := c.SetScalings(params)
	bailOnErrs(errs)
	if !*outputJSONFlag {
		printRecipe(*recipe)
		fmt.Println()
	} else {
		printAsJSON(*recipe)
	}
}

func showClusters(c *composeAPI.Client) {
	if *outputRawFlag {
		text, errs := c.GetClustersJSON()
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		clusters, errs := c.GetClusters()
		bailOnErrs(errs)

		if !*outputJSONFlag {
			for _, v := range *clusters {
				printCluster(v)
				fmt.Println()
			}
		} else {
			printAsJSON(clusters)
		}
	}
}

func showUser(c *composeAPI.Client) {
	if *outputRawFlag {
		text, errs := c.GetUserJSON()
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		user, errs := c.GetUser()
		bailOnErrs(errs)
		if !*outputJSONFlag {
			fmt.Printf("%15s: %s\n", "ID", user.ID)
			fmt.Println()
		} else {
			printAsJSON(user)
		}
	}
}

func showDatacenters(c *composeAPI.Client) {
	if *outputRawFlag {
		text, errs := c.GetDatacentersJSON()
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		datacenters, errs := c.GetDatacenters()
		bailOnErrs(errs)

		if !*outputJSONFlag {
			for _, v := range *datacenters {
				printDatacenter(v)
				fmt.Println()
			}
		} else {
			printAsJSON(datacenters)
		}
	}
}

func showDatabases(c *composeAPI.Client) {
	if *outputRawFlag {
		text, errs := c.GetDatabasesJSON()
		bailOnErrs(errs)
		fmt.Println(text)
	} else {
		databases, errs := c.GetDatabases()
		bailOnErrs(errs)

		if !*outputJSONFlag {
			for _, v := range *databases {
				printDatabase(v)
				fmt.Println()
			}
		} else {
			printAsJSON(databases)
		}
	}
}

func createDeployment(c *composeAPI.Client) {
	if *outputRawFlag {
		log.Fatal("Raw mode not supported for createDeployment")
	}

	account, errs := c.GetAccount()
	bailOnErrs(errs)

	if *createdeploymentdatacenter == "" && *createdeploymentcluster == "" {
		log.Fatal("Must supply either a --cluster id or --datacenter region")
	}

	params := composeAPI.CreateDeploymentParams{
		Name:         *createdeploymentname,
		AccountID:    account.ID,
		DatabaseType: *createdeploymenttype,
		Datacenter:   *createdeploymentdatacenter,
		ClusterID:    *createdeploymentcluster,
	}

	deployment, errs := c.CreateDeployment(params)
	bailOnErrs(errs)

	if !*outputJSONFlag {
		printDeployment(*deployment)
	} else {
		printAsJSON(*deployment)
	}

}

func getLink(link composeAPI.Link) string {
	return strings.Replace(link.HREF, "{?embed}", "", -1) // TODO: This should mangle the HREF properly
}

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
		if *showFullCAFlag {
			if *noDecodeCAFlag {
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
