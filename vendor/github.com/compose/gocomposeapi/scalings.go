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
	"errors"
	"github.com/parnurzeal/gorequest"
	"strconv"
)

// Scalings represents the used, allocated, starting and minimum unit scale
// of a deployment
type Scalings struct {
	AllocatedUnits int `json:"allocated_units"`
	UsedUnits      int `json:"used_units"`
	StartingUnits  int `json:"starting_units"`
	MinimumUnits   int `json:"minimum_units"`
}

//ScalingsParams represents the parameters needed to scale a deployment
type ScalingsParams struct {
	DeploymentID string `json:"-"`
	Deployment   struct {
		Units int `json:"units"`
	} `json:"deployment"`
}

//GetScalingsJSON returns raw scalings
func (c *Client) GetScalingsJSON(deploymentid string) (string, []error) {
	return c.getJSON("deployments/" + deploymentid + "/scalings")
}

//GetScalings returns deployment structure
func (c *Client) GetScalings(deploymentid string) (*Scalings, []error) {
	body, errs := c.GetScalingsJSON(deploymentid)

	if errs != nil {
		return nil, errs
	}

	scalings := Scalings{}
	json.Unmarshal([]byte(body), &scalings)

	return &scalings, nil
}

//SetScalingsJSON sets JSON scaling and returns string respones
func (c *Client) SetScalingsJSON(params ScalingsParams) (string, []error) {
	response, body, errs := gorequest.New().Post(apibase+"deployments/"+params.DeploymentID+"/scalings").
		Set("Authorization", "Bearer "+c.apiToken).
		Set("Content-type", "application/json; charset=utf-8").
		Send(params).
		End()

	if response.StatusCode != 200 { // Expect Accepted on success - assume error on anything else
		myerrors := Errors{}
		err := json.Unmarshal([]byte(body), &myerrors)
		if err != nil {
			errs = append(errs, errors.New("Unable to parse error - status code "+strconv.Itoa(response.StatusCode)))
		} else {
			errs = append(errs, errors.New(myerrors.Error))
		}
	}

	return body, errs
}

//SetScalings sets scale and returns recipe for scaling
func (c *Client) SetScalings(scalingsParams ScalingsParams) (*Recipe, []error) {
	body, errs := c.SetScalingsJSON(scalingsParams)
	if errs != nil {
		return nil, errs
	}

	recipe := Recipe{}
	json.Unmarshal([]byte(body), &recipe)

	return &recipe, nil
}
