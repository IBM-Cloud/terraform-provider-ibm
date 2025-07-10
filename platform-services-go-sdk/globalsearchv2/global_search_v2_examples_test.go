//go:build examples

/**
 * (C) Copyright IBM Corp. 2020, 2021.
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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globalsearchv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Global Search service.
//
// The following configuration properties are assumed to be defined:
//
// The following configuration properties are assumed to be defined in the external configuration file:
// GLOBAL_SEARCH_URL=<service url>
// GLOBAL_SEARCH_AUTHTYPE=iam
// GLOBAL_SEARCH_APIKEY=<IAM api key>
// GLOBAL_SEARCH_AUTH_URL=<IAM token service URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//

var _ = Describe(`GlobalSearchV2 Examples Tests`, func() {
	const externalConfigFile = "../global_search.env"

	var (
		globalSearchService *globalsearchv2.GlobalSearchV2
		config              map[string]string
		configLoaded        bool = false
	)

	var shouldSkipTest = func() {
		if !configLoaded {
			Skip("External configuration is not available, skipping tests...")
		}
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(globalsearchv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			globalSearchServiceOptions := &globalsearchv2.GlobalSearchV2Options{}

			globalSearchService, err = globalsearchv2.NewGlobalSearchV2UsingExternalConfig(globalSearchServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(globalSearchService).ToNot(BeNil())
		})

		It("Successfully construct the service client instance programmatically", func() {
			// This block is here merely to provide a working example for the front-matter
			// associated with the search API reference.
			var err error

			iamApiKey := config["APIKEY"]
			serviceURL := config["URL"]

			// begin example
			authenticator := &core.IamAuthenticator{
				ApiKey: iamApiKey,
			}

			serviceOptions := &globalsearchv2.GlobalSearchV2Options{
				Authenticator: authenticator,
				URL:           serviceURL,
			}

			searchService, err := globalsearchv2.NewGlobalSearchV2(serviceOptions)

			if err != nil {
				panic(err)
			}
			// end example

			Expect(searchService).ToNot(BeNil())
		})
	})

	Describe(`GlobalSearchV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Search request example`, func() {
			fmt.Println("\nSearch() result:")
			// begin-search

			searchOptions := globalSearchService.NewSearchOptions()
			searchOptions.SetLimit(10)
			searchOptions.SetQuery("GST-sdk-*")
			searchOptions.SetFields([]string{"*"})

			scanResult, response, err := globalSearchService.Search(searchOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scanResult, "", "  ")
			fmt.Println(string(b))

			// end-search

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(scanResult).ToNot(BeNil())

		})
	})
})
