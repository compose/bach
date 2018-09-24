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

package composeapi

import (
	"encoding/json"
	"fmt"
)

// DeploymentWhitelistParams - construct to pass to CreateDeploymentWhitelist
type DeploymentWhitelistParams struct {
	IP          string `json:"ip"`
	Description string `json:"description"`
}

// DeploymentWhitelist - representation of an applied whitelist entry
type DeploymentWhitelist struct {
	DeploymentWhitelistID string `json:"id"`
	Description           string `json:"description"`
	IP                    string `json:"ip"`
}

type createDeploymentWhitelistWrapper struct {
	Deployment createDeploymentWhitelistParams `json:"deployment"`
}

type createDeploymentWhitelistParams struct {
	Whitelist DeploymentWhitelistParams `json:"whitelist"`
}

type deploymentWhitelistResponse struct {
	Embedded struct {
		Whitelist []DeploymentWhitelist `json:"whitelist"`
	} `json:"_embedded"`
}

func (c *Client) createDeploymentWhitelistJSON(deploymentID string, params DeploymentWhitelistParams) (string, []error) {
	whitelistParams := createDeploymentWhitelistWrapper{
		Deployment: createDeploymentWhitelistParams{
			Whitelist: params,
		},
	}
	url := apibase + "deployments/" + deploymentID + "/whitelist"

	response, body, errs := c.newRequest("POST", url).
		Send(whitelistParams).
		End()

	if response.StatusCode != 202 {
		myerrors := Errors{}
		err := json.Unmarshal([]byte(body), &myerrors)
		if err != nil {
			errs = append(errs, fmt.Errorf("Unable to parse error - status code %d - body %s", response.StatusCode, response.Body))
		} else {
			errs = append(errs, fmt.Errorf("%v", myerrors.Error))
		}
	}

	return body, errs
}

// CreateDeploymentWhitelist creates a single deployment whitelist entry for a CIDR range
func (c *Client) CreateDeploymentWhitelist(deploymentID string, params DeploymentWhitelistParams) (*Recipe, []error) {
	body, errs := c.createDeploymentWhitelistJSON(deploymentID, params)
	if errs != nil {
		return nil, errs
	}

	deployed := Recipe{}
	err := json.Unmarshal([]byte(body), &deployed)
	if err != nil {
		return nil, []error{err}
	}

	return &deployed, nil
}

func (c *Client) getWhitelistForDeploymentJSON(deploymentid string) (string, []error) {
	return c.getJSON("deployments/" + deploymentid + "/whitelist")
}

// GetWhitelistForDeployment gets all whitelist entries for a given deployment ID
func (c *Client) GetWhitelistForDeployment(deploymentid string) ([]DeploymentWhitelist, []error) {
	body, errs := c.getWhitelistForDeploymentJSON(deploymentid)
	if errs != nil {
		return nil, errs
	}

	whitelistResponse := deploymentWhitelistResponse{}
	err := json.Unmarshal([]byte(body), &whitelistResponse)
	if err != nil {
		return nil, []error{err}
	}

	return whitelistResponse.Embedded.Whitelist, nil
}
