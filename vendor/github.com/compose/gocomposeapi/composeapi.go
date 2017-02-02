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
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
	"strconv"
)

var ()

const (
	apibase = "https://api.compose.io/2016-07/"
)

type Client struct {
	apiToken string
}

func NewClient(apiToken string) (*Client, error) {
	return &Client{
		apiToken: apiToken,
	}, nil
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
	response, body, errs := gorequest.New().Get(apibase+endpoint).
		Set("Authorization", "Bearer "+c.apiToken).
		Set("Content-type", "json").
		End()
	if response.StatusCode != 200 {
		myerrors := Errors{}
		err := json.Unmarshal([]byte(body), &myerrors)
		if err != nil {
			errs = append(errs, errors.New("Unable to parse error - status code "+strconv.Itoa(response.StatusCode)))
		} else {
			errs = append(errs, errors.New(fmt.Sprintf("%v", myerrors.Error)))
		}
	}
	return body, errs
}
