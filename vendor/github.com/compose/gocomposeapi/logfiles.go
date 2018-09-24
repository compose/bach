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
	"time"
)

// Logfile structure
type Logfile struct {
	ID           string    `json:"id"`
	Deploymentid string    `json:"deployment_id"`
	Capsuleid    string    `json:"capsule_id"`
	Name         string    `json:"name"`
	Region       string    `json:"region"`
	Status       string    `json:"status"`
	Date         time.Time `json:"created_at"`
	DownloadLink string    `json:"download_link"`
}

// logfilesResponse is used to represent and remove the JSON+HAL Embedded wrapper
type logfilesResponse struct {
	Embedded struct {
		Logfiles []Logfile `json:"logfiles"`
	} `json:"_embedded"`
}

//GetLogfilesForDeploymentJSON returns logfile details for deployment
func (c *Client) GetLogfilesForDeploymentJSON(deploymentid string) (string, []error) {
	return c.getJSON("deployments/" + deploymentid + "/logfiles")
}

//GetLogfilesForDeployment returns logfile details for deployment
func (c *Client) GetLogfilesForDeployment(deploymentid string) (*[]Logfile, []error) {
	body, errs := c.GetLogfilesForDeploymentJSON(deploymentid)

	if errs != nil {
		return nil, errs
	}

	logfilesResponse := logfilesResponse{}
	json.Unmarshal([]byte(body), &logfilesResponse)
	Logfiles := logfilesResponse.Embedded.Logfiles

	return &Logfiles, nil
}

//GetLogfileDetailsForDeploymentJSON returns the details and download link for a logfile
func (c *Client) GetLogfileDetailsForDeploymentJSON(deploymentid string, logfileid string) (string, []error) {
	return c.getJSON("deployments/" + deploymentid + "/logfiles/" + logfileid)
}

//GetLogfileDetailsForDeployment returns logfile details for deployment
func (c *Client) GetLogfileDetailsForDeployment(deploymentid string, logfileid string) (*Logfile, []error) {
	body, errs := c.GetLogfileDetailsForDeploymentJSON(deploymentid, logfileid)

	if errs != nil {
		return nil, errs
	}

	logfile := Logfile{}
	json.Unmarshal([]byte(body), &logfile)

	return &logfile, nil
}
