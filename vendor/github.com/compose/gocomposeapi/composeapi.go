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
//

// Package composeapi provides an idiomatic Go wrapper around the Compose
// API for database platform for deployment, management and monitoring.
package composeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/parnurzeal/gorequest"
)

const (
	apibase = "https://api.compose.io/2016-07/"
)

// Client is a structure that holds session information for the API
type Client struct {
	apiToken      string
	logger        *log.Logger
	enableLogging bool
}

// NewClient returns a Client for further interaction with the API
func NewClient(apiToken string) (*Client, error) {
	return &Client{
		apiToken: apiToken,
		logger:   log.New(ioutil.Discard, "", 0),
	}, nil
}

// SetLogger can enable or disable http logging to and from the Compose
// API endpoint using the provided io.Writer for the provided client.
func (c *Client) SetLogger(enableLogging bool, logger io.Writer) *Client {
	c.logger = log.New(logger, "[composeapi]", log.LstdFlags)
	c.enableLogging = enableLogging
	return c
}

func (c *Client) newRequest(method, targetURL string) *gorequest.SuperAgent {
	return gorequest.New().
		CustomMethod(method, targetURL).
		Set("Authorization", "Bearer "+c.apiToken).
		Set("Content-type", "application/json; charset=utf-8").
		SetLogger(c.logger).
		SetDebug(c.enableLogging)
}

// Link structure for JSON+HAL links
type Link struct {
	HREF      string `json:"href"`
	Templated bool   `json:"templated"`
}

//Errors struct for parsing error returns
type Errors struct {
	Error map[string][]string `json:"errors,omitempty"`
}

func printJSON(jsontext string) {
	var tempholder map[string]interface{}

	if err := json.Unmarshal([]byte(jsontext), &tempholder); err != nil {
		log.Fatal(err)
	}
	indentedjson, err := json.MarshalIndent(tempholder, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(indentedjson))
}

//SetAPIToken overrides the API token
func (c *Client) SetAPIToken(newtoken string) {
	c.apiToken = newtoken
}

//GetJSON Gets JSON string of content at an endpoint
func (c *Client) getJSON(endpoint string) (string, []error) {
	response, body, errs := c.newRequest("GET", apibase+endpoint).End()

	if response.StatusCode != 200 {
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
