//go:build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/ibmcloudshellv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the ibmcloudshellv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`IBMCloudShellV1 Integration Tests`, func() {

	const externalConfigFile = "../ibm_cloud_shell_v1.env"

	var (
		err                  error
		ibmCloudShellService *ibmcloudshellv1.IBMCloudShellV1
		serviceURL           string
		config               map[string]string
		accountID            string
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
			config, err = core.GetServiceProperties(ibmcloudshellv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}
			accountID = config["ACCOUNT_ID"]
			if accountID == "" {
				Skip("Unable to load account ID configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			ibmCloudShellServiceOptions := &ibmcloudshellv1.IBMCloudShellV1Options{}

			ibmCloudShellService, err = ibmcloudshellv1.NewIBMCloudShellV1UsingExternalConfig(ibmCloudShellServiceOptions)

			Expect(err).To(BeNil())
			Expect(ibmCloudShellService).ToNot(BeNil())
			Expect(ibmCloudShellService.Service.Options.URL).To(Equal(serviceURL))

			goLogger := log.New(GinkgoWriter, "", log.LstdFlags)
			core.SetLogger(core.NewLogger(core.LevelDebug, goLogger, goLogger))
			ibmCloudShellService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetAccountSettings - Get account settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions)`, func() {

			getAccountSettingsOptions := &ibmcloudshellv1.GetAccountSettingsOptions{
				AccountID: &accountID,
			}

			accountSettings, response, err := ibmCloudShellService.GetAccountSettings(getAccountSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

		})
	})

	Describe(`UpdateAccountSettings - Update account settings`, func() {

		var existingAccountSettings *ibmcloudshellv1.AccountSettings

		BeforeEach(func() {
			shouldSkipTest()
			getAccountSettingsOptions := &ibmcloudshellv1.GetAccountSettingsOptions{
				AccountID: &accountID,
			}
			accountSettings, response, err := ibmCloudShellService.GetAccountSettings(getAccountSettingsOptions)
			existingAccountSettings = accountSettings
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(existingAccountSettings).ToNot(BeNil())
		})
		It(`UpdateAccountSettings(updateAccountSettingsOptions *UpdateAccountSettingsOptions)`, func() {

			featureModel := []ibmcloudshellv1.Feature{
				{
					Enabled: core.BoolPtr(false),
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
					Enabled: core.BoolPtr(false),
					Key:     core.StringPtr("us-south"),
				},
			}

			updateAccountSettingsOptions := &ibmcloudshellv1.UpdateAccountSettingsOptions{
				AccountID:                &accountID,
				Rev:                      existingAccountSettings.Rev,
				DefaultEnableNewFeatures: core.BoolPtr(false),
				DefaultEnableNewRegions:  core.BoolPtr(true),
				Enabled:                  core.BoolPtr(true),
				Features:                 featureModel,
				Regions:                  regionSettingModel,
			}

			accountSettings, response, err := ibmCloudShellService.UpdateAccountSettings(updateAccountSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())
			Expect(*accountSettings.DefaultEnableNewFeatures).To(Equal(false))
			Expect(*accountSettings.DefaultEnableNewRegions).To(Equal(true))
			Expect(*accountSettings.Enabled).To(Equal(true))
			Expect(accountSettings.Features).To(Equal(featureModel))
			Expect(accountSettings.Regions).To(Equal(regionSettingModel))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
