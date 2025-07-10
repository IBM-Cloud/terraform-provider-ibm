//go:build examples

/**
 * (C) Copyright IBM Corp. 2021.
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

package ibmcloudshellv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/ibmcloudshellv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the IBM Cloud Shell service.
//
// The following configuration properties are assumed to be defined:
// IBM_CLOUD_SHELL_URL=<service base url>
// IBM_CLOUD_SHELL_AUTH_TYPE=iam
// IBM_CLOUD_SHELL_APIKEY=<IAM apikey>
// IBM_CLOUD_SHELL_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
// IBM_CLOUD_SHELL_ACCOUNT_ID=<IBM Cloud account ID>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//

var _ = Describe(`IBMCloudShellV1 Examples Tests`, func() {
	const externalConfigFile = "../ibm_cloud_shell_v1.env"

	var (
		ibmCloudShellService *ibmcloudshellv1.IBMCloudShellV1
		config               map[string]string
		configLoaded         bool = false
		accountID            string
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
			config, err = core.GetServiceProperties(ibmcloudshellv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			accountID = config["ACCOUNT_ID"]
			Expect(accountID).ToNot(BeEmpty())

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

			ibmCloudShellServiceOptions := &ibmcloudshellv1.IBMCloudShellV1Options{}

			ibmCloudShellService, err = ibmcloudshellv1.NewIBMCloudShellV1UsingExternalConfig(ibmCloudShellServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(ibmCloudShellService).ToNot(BeNil())
		})
	})

	Describe(`IBMCloudShellV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountSettings request example`, func() {
			fmt.Println("\nGetAccountSettings() result:")
			// begin-get_account_settings

			getAccountSettingsOptions := ibmCloudShellService.NewGetAccountSettingsOptions(accountID)

			accountSettings, response, err := ibmCloudShellService.GetAccountSettings(getAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettings, "", "  ")
			fmt.Println(string(b))

			// end-get_account_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

		})
		It(`UpdateAccountSettings request example`, func() {
			fmt.Println("\nUpdateAccountSettings() result:")
			// begin-update_account_settings

			featureModel := []ibmcloudshellv1.Feature{
				{
					Enabled: core.BoolPtr(true),
					Key:     core.StringPtr("server.file_manager"),
				},
				{
					Enabled: core.BoolPtr(true),
					Key:     core.StringPtr("server.web_preview"),
				},
			}

			regionSettingModel := []ibmcloudshellv1.RegionSetting{
				{
					Enabled: core.BoolPtr(true),
					Key:     core.StringPtr("eu-de"),
				},
				{
					Enabled: core.BoolPtr(true),
					Key:     core.StringPtr("us-south"),
				},
			}

			updateAccountSettingsOptions := &ibmcloudshellv1.UpdateAccountSettingsOptions{
				AccountID:                &accountID,
				Rev:                      core.StringPtr(fmt.Sprintf("130-%s", accountID)),
				DefaultEnableNewFeatures: core.BoolPtr(false),
				DefaultEnableNewRegions:  core.BoolPtr(true),
				Enabled:                  core.BoolPtr(true),
				Features:                 featureModel,
				Regions:                  regionSettingModel,
			}

			accountSettings, response, err := ibmCloudShellService.UpdateAccountSettings(updateAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettings, "", "  ")
			fmt.Println(string(b))

			// end-update_account_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

		})
	})
})
