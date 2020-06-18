/*
Copyright 2020 Adobe. All rights reserved.
This file is licensed to you under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License. You may obtain a copy
of the License at http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR REPRESENTATIONS
OF ANY KIND, either express or implied. See the License for the specific language
governing permissions and limitations under the License.
*/

package analytics

// Response types

// CalculatedMetricDefinition represents a calculated metric definition
type CalculatedMetricDefinition struct {
}

// CalculatedMetric represents a calculated metric
type CalculatedMetric struct {
	ID              string                      `json:"id,omitempty"`
	Name            string                      `json:"name,omitempty"`
	Description     string                      `json:"description,omitempty"`
	RSID            string                      `json:"rsid,omitempty"`
	ReportSuiteName string                      `json:"reportSuiteName,omitempty"`
	Owner           *Owner                      `json:"owner,omitempty"`
	Polarity        string                      `json:"polarity,omitempty"`
	Precision       int                         `json:"precision"`
	Type            string                      `json:"type,omitempty"`
	Definition      *CalculatedMetricDefinition `json:"definition,omitempty"`
	Categories      []string                    `json:"categories,omitempty"`
	Tags            *[]Tag                      `json:"tags,omitempty"`
	SiteTitle       string                      `json:"siteTitle,omitempty"`
	Modified        string                      `json:"modified,omitempty"`
	Created         string                      `json:"created,omitempty"`
}

// CalculatedMetrics represents a page of calculated metrics
type CalculatedMetrics struct {
	Content          *[]CalculatedMetric `json:"content,omitempty"`
	Number           int                 `json:"number"`
	Size             int                 `json:"size"`
	NumberOfElements int                 `json:"numberOfElements"`
	TotalElements    int                 `json:"totalElements"`
	PreviousPage     bool                `json:"previousPage"`
	FirstPage        bool                `json:"firstPage"`
	NextPage         bool                `json:"nextPage"`
	LastPage         bool                `json:"lastPage"`
	Sort             *[]Sort             `json:"sort,omitempty"`
	TotalPages       int                 `json:"totalPages"`
}
