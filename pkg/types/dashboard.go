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
package types

type GridConfiguration struct {
	Col   int `json:"col,omitempty"`
	Row   int `json:"row,omitempty"`
	SizeX int `json:"size_x,omitempty"`
	SizeY int `json:"size_y,omitempty"`
}

type ValueLimit struct {
	Count     int    `json:"count,omitempty"`
	Direction string `json:"direction,omitempty"`
}

type Histogram struct {
	NumberOfBuckets int `json:"numberOfBuckets,omitempty"`
}

type YAxisLeftDomain struct {
	From int `json:"from,omitempty"`
	To   int `json:"to,omitempty"`
}

type YAxisRightDomain struct {
	From int `json:"from,omitempty"`
	To   int `json:"to,omitempty"`
}

type XAxis struct {
	From int `json:"from,omitempty"`
	To   int `json:"to,omitempty"`
}

type CustomDisplayOptions struct {
	ValueLimit       ValueLimit       `json:"valueLimit,omitempty"`
	Histogram        Histogram        `json:"histogram,omitempty"`
	YAxisScale       string           `json:"yAxisScale,omitempty"`
	YAxisLeftDomain  YAxisLeftDomain  `json:"yAxisLeftDomain,omitempty"`
	YAxisRightDomain YAxisRightDomain `json:"yAxisRightDomain,omitempty"`
	XAxis            XAxis            `json:"xAxis,omitempty"`
}

type Metrics struct {
	ID               string `json:"id,omitempty"`
	PropertyName     string `json:"propertyName,omitempty"`
	TimeAggregation  string `json:"timeAggregation,omitempty"`
	GroupAggregation string `json:"groupAggregation,omitempty"`
}

type CompareToConfig struct {
	PeriodCount  int    `json:"periodCount,omitempty"`
	PeriodTypeID string `json:"periodTypeId,omitempty"`
}
type Widgets struct {
	ShowAs               string               `json:"showAs,omitempty"`
	Name                 string               `json:"name,omitempty"`
	GridConfiguration    GridConfiguration    `json:"gridConfiguration,omitempty"`
	CustomDisplayOptions CustomDisplayOptions `json:"customDisplayOptions,omitempty"`
	Scope                string               `json:"scope,omitempty"`
	OverrideScope        bool                 `json:"overrideScope,omitempty"`
	Metrics              []Metrics            `json:"metrics,omitempty"`

	MarkdownSource        string `json:"markdownSource,omitempty"`
	PanelTitleVisible     bool   `json:"panelTitleVisible,omitempty"`
	TextAutosized         bool   `json:"textAutosized,omitempty"`
	TransparentBackground bool   `json:"transparentBackground,omitempty"`

	CompareToConfig interface{} `json:"compareToConfig,omitempty"`
	// ColorCoding     ColorCoding `json:"colorCoding"`
}

type Thresholds struct {
	Worst int `json:"worst"`
	Best  int `json:"best"`
}

type ColorCoding struct {
	Active     bool       `json:"active"`
	Thresholds Thresholds `json:"thresholds"`
}

type ScopeExpressionList struct {
	Operand     string   `json:"operand,omitempty"`
	Operator    string   `json:"operator,omitempty"`
	DisplayName string   `json:"displayName,omitempty"`
	Value       []string `json:"value"`
	Variable    bool     `json:"variable,omitempty"`
	IsVariable  bool     `json:"isVariable,omitempty"`
}

type EventsOverlaySettings struct {
	ShowNotificationsEnabled           bool   `json:"showNotificationsEnabled,omitempty"`
	FilterNotificationsScopeFilter     bool   `json:"filterNotificationsScopeFilter,omitempty"`
	FilterNotificationsUserInputFilter string `json:"filterNotificationsUserInputFilter,omitempty"`
	EventOverlayLimit                  int    `json:"eventOverlayLimit,omitempty"`
	FilterNotificationsTypeFilter      string `json:"filterNotificationsTypeFilter,omitempty"`
	FilterNotificationsStateFilter     string `json:"filterNotificationsStateFilter,omitempty"`
	FilterNotificationsSeverityFilter  string `json:"filterNotificationsSeverityFilter,omitempty"`
	FilterNotificationsResolvedFilter  string `json:"filterNotificationsResolvedFilter,omitempty"`
}

type Dashboard struct {
	Widgets               []Widgets             `json:"widgets"`
	ScopeExpressionList   []ScopeExpressionList `json:"scopeExpressionList"`
	EventsOverlaySettings EventsOverlaySettings `json:"eventsOverlaySettings,omitempty"`
	Name                  string                `json:"name,omitempty"`
	ID                    int                   `json:"id,omitempty"`
	Public                bool                  `json:"public,omitempty"`
	Version               int                   `json:"version"`
	Shared                bool                  `json:"shared"`
	AutoCreated           bool                  `json:"autoCreated,omitempty"`
	Username              string                `json:"username,omitempty"`
	PublicToken           string                `json:"publicToken,omitempty"`
	TeamID                int                   `json:"teamId,omitempty"`
	Favorite              bool                  `json:"favorite,omitempty"`
}
