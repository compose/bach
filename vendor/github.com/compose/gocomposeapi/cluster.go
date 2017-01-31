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

// Cluster structure
type Cluster struct {
	ID          string    `json:"id"`
	AccountID   string    `json:"account_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Provider    string    `json:"provider"`
	Region      string    `json:"region"`
	Multitenant bool      `json:"multitenant"`
	AccountSlug string    `json:"account_slug"`
	CreatedAt   time.Time `json:"created_at"`
	Subdomain   string    `json:"subdomain"`
}

type clusterResponse struct {
	Embedded struct {
		Clusters []Cluster `json:"clusters"`
	} `json:"_embedded"`
}

//GetClustersJSON gets clusters available
func (c *Client) GetClustersJSON() (string, []error) {
	return c.getJSON("clusters")
}

//GetClusters gets clusters available
func (c *Client) GetClusters() (*[]Cluster, []error) {
	body, errs := c.GetClustersJSON()

	if errs != nil {
		return nil, errs
	}

	clustersResponse := clusterResponse{}
	json.Unmarshal([]byte(body), &clustersResponse)
	clusters := clustersResponse.Embedded.Clusters

	return &clusters, nil
}
