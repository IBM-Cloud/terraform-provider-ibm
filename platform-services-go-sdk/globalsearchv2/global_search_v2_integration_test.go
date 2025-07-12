//go:build integration

/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package globalsearchv2_test

import (
	"fmt"
	"log"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/IBM/platform-services-go-sdk/globalsearchv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the globalsearchv2 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`GlobalSearchV2 Integration Tests`, func() {

	const externalConfigFile = "../global_search.env"

	var (
		err                 error
		globalSearchService *globalsearchv2.GlobalSearchV2
		serviceURL          string
		config              map[string]string
		gstQuery            = "GST-sdk*"
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(globalsearchv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			globalSearchServiceOptions := &globalsearchv2.GlobalSearchV2Options{}

			globalSearchService, err = globalsearchv2.NewGlobalSearchV2UsingExternalConfig(globalSearchServiceOptions)

			Expect(err).To(BeNil())
			Expect(globalSearchService).ToNot(BeNil())
			Expect(globalSearchService.Service.Options.URL).To(Equal(serviceURL))

			goLogger := log.New(GinkgoWriter, "", log.LstdFlags)
			core.SetLogger(core.NewLogger(core.LevelDebug, goLogger, goLogger))
		})
	})

	Describe(`Search - Find instances of resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Search(searchOptions *SearchOptions)`, func() {

			searchResults := []globalsearchv2.ResultItem{}

			// Search for resources 1 item at a time to exercise pagination.
			var searchCursor *string = nil
			var limit int64 = 1
			var moreResults bool = true

			for moreResults {
				searchOptions := &globalsearchv2.SearchOptions{
					Query:        &gstQuery,
					Fields:       []string{"*"},
					SearchCursor: searchCursor,
					Limit:        &limit,
				}

				scanResult, response, err := globalSearchService.Search(searchOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(scanResult).ToNot(BeNil())
				fmt.Fprintf(GinkgoWriter, "Search() result:\n%s\n", common.ToJSON(scanResult))

				if len(scanResult.Items) > 0 {
					moreResults = true
					searchCursor = scanResult.SearchCursor
					searchResults = append(searchResults, scanResult.Items...)
				} else {
					moreResults = false
				}
			}

			fmt.Fprintf(GinkgoWriter, "Total results returned by Search(): %d\n", len(searchResults))
		})
	})
})
