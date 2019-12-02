/*
Copyright © 2019 Sysdig

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sysdiglabs/sdc-db-unshare/pkg/types"
	"io"
	"net/http"
)

type client struct {
	token      string
	httpClient *http.Client
	url        string
}

func (c *client) request(method, url string, body io.Reader) (response *http.Response, err error) {
	request, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.url, url), body)
	if err != nil {
		err = errors.Wrap(err, "unable to create request")
		return
	}
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	response, err = c.httpClient.Do(request)
	if response == nil || response.StatusCode >= 400 {
		err = errors.Errorf("status code err: %s (%d)", response.Status, response.StatusCode)
		return
	}
	err = errors.Wrap(err, "request error")
	return
}

func (c *client) GetAllDashboards() (dashboards []*types.Dashboard, err error) {
	response, err := c.request(http.MethodGet, "api/v2/dashboards", nil)
	if err != nil {
		err = errors.Wrap(err, "error retrieving the dashboards")
		return
	}
	defer response.Body.Close()

	type dashboardsResponseType struct {
		Dashboards []*types.Dashboard `json:"dashboards"`
	}
	var dashboardsResponse dashboardsResponseType
	err = json.NewDecoder(response.Body).Decode(&dashboardsResponse)

	dashboards = dashboardsResponse.Dashboards
	err = errors.Wrap(err, "unable to unmarshal the dashboards response")
	return
}

func (c *client) GetDashboard(ID string) (dashboard *types.Dashboard, err error) {
	response, err := c.request(http.MethodGet, fmt.Sprintf("api/v2/dashboards/%s", ID), nil)
	if err != nil {
		err = errors.Wrapf(err, "error retrieving the dashboard: %s", ID)
		return
	}
	defer response.Body.Close()

	type dashboardResponseType struct {
		Dashboard *types.Dashboard `json:"dashboard"`
	}
	var dashboardResponse dashboardResponseType
	err = json.NewDecoder(response.Body).Decode(&dashboardResponse)

	dashboard = dashboardResponse.Dashboard
	err = errors.Wrap(err, "unable to unmarshal the dashboard response")
	return
}

func (c *client) getDashboardAsMap(ID string) (dashboard map[string]interface{}, err error) {
	response, err := c.request(http.MethodGet, fmt.Sprintf("api/v2/dashboards/%s", ID), nil)
	if err != nil {
		err = errors.Wrapf(err, "error retrieving the dashboard: %s", ID)
		return
	}
	defer response.Body.Close()

	type dashboardResponseType struct {
		Dashboard map[string]interface{} `json:"dashboard"`
	}
	var dashboardResponse dashboardResponseType
	err = json.NewDecoder(response.Body).Decode(&dashboardResponse)

	dashboard = dashboardResponse.Dashboard
	err = errors.Wrap(err, "unable to unmarshal the dashboard response")
	return
}

func (c *client) UnshareDashboard(ID string) (err error) {
	// we get the dashboard as map, because we don't want Go to modify the other existing fields while
	// unmarshaling and remarshaling it -- we just want to modify the "shared" field to False.
	dashboard, err := c.getDashboardAsMap(ID)
	if err != nil {
		err = errors.Wrap(err, "unable to retrieve the dashboard")
		return
	}
	dashboard["shared"] = false

	type dashboardRequestType struct {
		Dashboard map[string]interface{} `json:"dashboard"`
	}
	dashboardRequest := dashboardRequestType{Dashboard: dashboard}
	buff := &bytes.Buffer{}
	err = json.NewEncoder(buff).Encode(dashboardRequest)
	if err != nil {
		err = errors.Wrap(err, "unable to marshall the request")
	}

	response, err := c.request(http.MethodPut, fmt.Sprintf("api/v2/dashboards/%s", ID), buff)
	if err != nil {
		buff := &bytes.Buffer{}
		io.Copy(buff, response.Body)
		err = errors.Wrapf(err, "unable to update the dashboard: %s", buff.String())
		return
	}
	defer response.Body.Close()
	return
}

func New(token string) Client {
	return &client{
		token:      token,
		httpClient: http.DefaultClient,
		url:        "https://app.sysdigcloud.com",
	}
}
