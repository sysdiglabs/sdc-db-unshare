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
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sysdiglabs/sdc-db-unshare/pkg/usecases"
	"log"
)

var unshareCmd = &cobra.Command{
	Use:   "unshare [ID]",
	Short: "Unshares a dashboard from Sysdig Monitor",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dashboardID := args[0]
		factory := usecases.NewFactory(token)
		useCase := factory.UnshareDashboardUseCase(dashboardID)
		err := useCase.Execute()
		if err != nil {
			log.Fatalf("%#v\n", err)
		}
	},
}

func init() {
	dashboardCmd.AddCommand(unshareCmd)
}
