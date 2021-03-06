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

package main

import (
	"log"

	"github.com/adobe/aa-client-go/examples/config"
	"github.com/adobe/aa-client-go/examples/utils"
)

func main() {
	// Read configuration
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Create an Analytics client
	client, err := utils.NewAnalyticsClient(&cfg.Analytics)
	if err != nil {
		log.Fatal(err)
	}

	// Get dimensions
	dimensions, err := client.Dimensions.GetByID(cfg.Analytics.ReportSuiteID, "evar1", "en_US", []string{})
	if err != nil {
		log.Fatal(err)
	}

	utils.PrintJSON(dimensions)

	// Get dimension
	dimension, err := client.Dimensions.GetAll(cfg.Analytics.ReportSuiteID, "en_US", false, false, false, []string{})
	if err != nil {
		log.Fatal(err)
	}

	utils.PrintJSON(dimension)
}
