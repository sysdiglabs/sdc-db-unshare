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
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/sysdiglabs/sdc-db-unshare/pkg/usecases"
	"log"
	"os"
	"strconv"
	"strings"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the dashboards",
	Run: func(cmd *cobra.Command, args []string) {
		factory := usecases.NewFactory(token)
		useCase := factory.RetrieveAllDashboardsUseCase()
		dashboards, err := useCase.Execute()
		if err != nil {
			log.Fatalf("%#v\n", err)
		}

		table := simpleTable()
		table.SetHeader([]string{"ID", "Dashboard", "User"})

		for _, dashboard := range dashboards {
			if dashboard.Name == "" {
				continue
			}
			table.Append([]string{strconv.Itoa(dashboard.ID), strings.ReplaceAll(dashboard.Name, "\n", ""), dashboard.Username})
		}

		table.Render()
	},
}

func simpleTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding(" ") // pad with tabs
	table.SetNoWhiteSpace(true)
	return table
}

func init() {
	dashboardCmd.AddCommand(listCmd)
}
