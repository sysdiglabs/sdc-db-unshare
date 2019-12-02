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
package usecases

import (
	"github.com/sysdiglabs/sdc-db-unshare/pkg/client"
	"github.com/sysdiglabs/sdc-db-unshare/pkg/types"
)

type RetrieveOneDashboard struct {
	sysdigClient client.Client
	dashboardID  string
}

func (r RetrieveOneDashboard) Execute() (dashboard *types.Dashboard, err error) {
	return r.sysdigClient.GetDashboard(r.dashboardID)
}
