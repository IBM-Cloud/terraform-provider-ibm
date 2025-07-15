//go:build integration

/**
 * (C) Copyright IBM Corp. 2020, 2024.
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

package iamidentityv1_test

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the iamidentityv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`IamIdentityV1 Integration Tests`, func() {

	const externalConfigFile = "../iam_identity.env"

	var (
		apikeyName    string = "Go-SDK-IT-APIKey"
		serviceIDName string = "Go-SDK-IT-ServiceId"
		profileName1  string = "Go-SDK-IT-Profile-1"
		profileName2  string = "Go-SDK-IT-Profile-2"
		accountID     string
		iamID         string
		iamIDMember   string
		iamAPIKey     string
		claimRuleType string = "Profile-SAML"
		realmName     string = "https://sdk.test.realm/1234"

		iamIdentityService *iamidentityv1.IamIdentityV1
		err                error
		serviceURL         string
		config             map[string]string

		apikeyId1   string
		apikeyId2   string
		apikeyEtag1 string

		serviceId1     string
		serviceIdEtag1 string
		newDescription string = "This is an updated description"

		profileId1   string
		profileId2   string
		profileIamId string
		profileEtag  string

		claimRuleId1  string
		claimRuleId2  string
		claimRuleEtag string

		linkId string

		accountSettingEtag string

		reportId           string
		reportReferenceMfa string

		enterpriseAccountID                   string
		enterpriseSubAccountID                string
		profileTemplateName                   string = "Go-SDK-IT-Profile-Template"
		profileTemplateProfileName            string = "Go-SDK-IT-Profile-From-Template"
		profileTemplateId                     string
		profileTemplateVersion                int64
		profileTemplateEtag                   string
		profileTemplateAssignmentId           string
		profileTemplateAssignmentEtag         string
		accountSettingsTemplateName           string = "Go-SDK-IT-AccountSettings-Template"
		accountSettingsTemplateId             string
		accountSettingsTemplateVersion        int64
		accountSettingsTemplateEtag           string
		accountSettingsTemplateAssignmentId   string
		accountSettingsTemplateAssignmentEtag string

		service             string = "console"
		valueString         string = "/billing"
		preferenceID1       string = "landing_page"
		iamIDForPreferences string
		filter              string = "name co \"Go\""
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
			config, err = core.GetServiceProperties(iamidentityv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			accountID = config["ACCOUNT_ID"]
			Expect(accountID).ToNot(BeEmpty())

			iamID = config["IAM_ID"]
			Expect(iamID).ToNot(BeEmpty())

			iamIDMember = config["IAM_ID_MEMBER"]
			Expect(iamIDMember).ToNot(BeEmpty())

			iamAPIKey = config["APIKEY"]
			Expect(iamAPIKey).ToNot(BeEmpty())

			fmt.Fprintf(GinkgoWriter, "Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}

			enterpriseAccountID = config["ENTERPRISE_ACCOUNT_ID"]
			Expect(enterpriseAccountID).ToNot(BeEmpty())

			enterpriseSubAccountID = config["ENTERPRISE_SUBACCOUNT_ID"]
			Expect(enterpriseSubAccountID).ToNot(BeEmpty())

			iamIDForPreferences = config["IAM_ID_FOR_PREFERENCES"]
			Expect(iamIDForPreferences).ToNot(BeEmpty())

		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			iamIdentityServiceOptions := &iamidentityv1.IamIdentityV1Options{}

			iamIdentityService, err = iamidentityv1.NewIamIdentityV1UsingExternalConfig(iamIdentityServiceOptions)

			Expect(err).To(BeNil())
			Expect(iamIdentityService).ToNot(BeNil())
			Expect(iamIdentityService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			iamIdentityService.EnableRetries(4, 30*time.Second)
		})
		It("Successfully setup the environment for tests", func() {
			fmt.Fprintln(GinkgoWriter, "Setup...")
			cleanupResources(iamIdentityService, accountID, iamID, apikeyName, serviceIDName, profileName1, profileName2, enterpriseAccountID, profileTemplateName, &profileTemplateAssignmentEtag, accountSettingsTemplateName, &accountSettingsTemplateAssignmentEtag)
			fmt.Fprintln(GinkgoWriter, "Finished setup.")
		})
	})

	Describe(`CreateAPIKey1 - Create API key #1`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAPIKey(createAPIKeyOptions *CreateAPIKeyOptions)`, func() {

			createAPIKeyOptions := &iamidentityv1.CreateAPIKeyOptions{
				Name:        &apikeyName,
				IamID:       &iamID,
				Description: core.StringPtr("GoSDK test apikey #1"),
			}

			apiKey, response, err := iamIdentityService.CreateAPIKey(createAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKey).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateAPIKey #1 response:\n%s\n", common.ToJSON(apiKey))

			apikeyId1 = *apiKey.ID
			Expect(apikeyId1).ToNot(BeNil())
		})
	})

	Describe(`CreateAPIKey2 - Create API key #2`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAPIKey(createAPIKeyOptions *CreateAPIKeyOptions)`, func() {

			createAPIKeyOptions := &iamidentityv1.CreateAPIKeyOptions{
				Name:        &apikeyName,
				IamID:       &iamID,
				Description: core.StringPtr("GoSDK test apikey #2"),
			}

			apiKey, response, err := iamIdentityService.CreateAPIKey(createAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKey).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateAPIKey #2 response:\n%s\n", common.ToJSON(apiKey))

			apikeyId2 = *apiKey.ID
			Expect(apikeyId2).ToNot(BeNil())
		})
	})

	Describe(`GetAPIKey - Get details of an API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAPIKey(getAPIKeyOptions *GetAPIKeyOptions)`, func() {
			Expect(apikeyId1).ToNot(BeNil())

			getAPIKeyOptions := &iamidentityv1.GetAPIKeyOptions{
				ID:              &apikeyId1,
				IncludeHistory:  core.BoolPtr(true),
				IncludeActivity: core.BoolPtr(true),
			}

			apiKey, response, err := iamIdentityService.GetAPIKey(getAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetAPIKey response:\n%s\n", common.ToJSON(apiKey))

			Expect(*apiKey.ID).To(Equal(apikeyId1))
			Expect(*apiKey.Name).To(Equal(apikeyName))
			Expect(*apiKey.IamID).To(Equal(iamID))
			Expect(*apiKey.AccountID).To(Equal(accountID))
			Expect(*apiKey.CreatedBy).To(Equal(iamID))
			Expect(*apiKey.CreatedAt).ToNot(BeNil())
			Expect(*apiKey.Locked).To(BeFalse())
			Expect(*apiKey.CRN).ToNot(BeNil())
			Expect(apiKey.History).ToNot(BeEmpty())
			Expect(apiKey.Activity).ToNot(BeNil())
			Expect(apiKey.Activity.AuthnCount).ToNot(BeNil())

			// Grab the Etag value from the response for use in the update operation.
			apikeyEtag1 = response.GetHeaders().Get("Etag")
			Expect(apikeyEtag1).ToNot(BeEmpty())
		})
	})

	Describe(`GetAPIKeysDetails - Get details of an API key by its value`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAPIKeysDetails(getAPIKeysDetailsOptions *GetAPIKeysDetailsOptions)`, func() {

			getAPIKeysDetailsOptions := &iamidentityv1.GetAPIKeysDetailsOptions{
				IamAPIKey:      &iamAPIKey,
				IncludeHistory: core.BoolPtr(true),
			}

			apiKey, response, err := iamIdentityService.GetAPIKeysDetails(getAPIKeysDetailsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetAPIKeyDetails response:\n%s\n", common.ToJSON(apiKey))

			Expect(*apiKey.AccountID).To(Equal(accountID))
			Expect(*apiKey.IamID).To(Equal(iamID))
			Expect(*apiKey.Locked).To(BeFalse())
			Expect(apiKey.History).ToNot(BeEmpty())
		})
	})

	Describe(`ListAPIKeys - Get API keys for a given service or user IAM ID and account ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAPIKeys(listAPIKeysOptions *ListAPIKeysOptions)`, func() {

			apikeys := []iamidentityv1.APIKey{}

			// var pageToken *string = nil
			var pageTokenPresent bool = true
			var pageToken *string = nil

			// for ok := true; ok; ok = (pageToken != nil) {
			for pageTokenPresent {
				listAPIKeysOptions := &iamidentityv1.ListAPIKeysOptions{
					AccountID: &accountID,
					IamID:     &iamID,
					Pagetoken: pageToken,
					Pagesize:  core.Int64Ptr(int64(1)),
					Filter:    &filter,
				}

				apiKeyList, response, err := iamIdentityService.ListAPIKeys(listAPIKeysOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(apiKeyList).ToNot(BeNil())
				fmt.Fprintf(GinkgoWriter, "ListAPIKeys response:\n%s\n", common.ToJSON(apiKeyList))

				// Walk through the returned results and save off the apikeys that we created earlier.
				for _, apikey := range apiKeyList.Apikeys {
					if apikeyName == *apikey.Name {
						apikeys = append(apikeys, apikey)
					}
				}

				pageToken = getPageTokenFromURL(apiKeyList.Next)
				pageTokenPresent = (pageToken != nil)
			}

			// Make sure we got back two apikeys.
			Expect(len(apikeys)).To(Equal(2))
		})
	})

	Describe(`UpdateAPIKey - Updates an API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAPIKey(updateAPIKeyOptions *UpdateAPIKeyOptions)`, func() {
			Expect(apikeyId1).ToNot(BeEmpty())
			Expect(apikeyEtag1).ToNot(BeEmpty())

			updateAPIKeyOptions := &iamidentityv1.UpdateAPIKeyOptions{
				ID:          &apikeyId1,
				IfMatch:     &apikeyEtag1,
				Description: &newDescription,
			}

			apiKey, response, err := iamIdentityService.UpdateAPIKey(updateAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "UpdateAPIKey response:\n%s\n", common.ToJSON(apiKey))

			Expect(*apiKey.ID).To(Equal(apikeyId1))
			Expect(*apiKey.Description).To(Equal(newDescription))
		})
	})

	Describe(`LockAPIKey - Lock the API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`LockAPIKey(lockAPIKeyOptions *LockAPIKeyOptions)`, func() {
			Expect(apikeyId2).ToNot(BeEmpty())

			lockAPIKeyOptions := &iamidentityv1.LockAPIKeyOptions{
				ID: &apikeyId2,
			}

			response, err := iamIdentityService.LockAPIKey(lockAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			apiKey := getAPIkey(iamIdentityService, apikeyId2)
			Expect(apiKey).ToNot(BeNil())
			Expect(*apiKey.Locked).To(BeTrue())
		})
	})

	Describe(`UnlockAPIKey - Unlock the API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UnlockAPIKey(unlockAPIKeyOptions *UnlockAPIKeyOptions)`, func() {
			Expect(apikeyId2).ToNot(BeEmpty())

			unlockAPIKeyOptions := &iamidentityv1.UnlockAPIKeyOptions{
				ID: &apikeyId2,
			}

			response, err := iamIdentityService.UnlockAPIKey(unlockAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			apiKey := getAPIkey(iamIdentityService, apikeyId2)
			Expect(apiKey).ToNot(BeNil())
			Expect(*apiKey.Locked).To(BeFalse())
		})
	})

	Describe(`DisableAPIKey - Enable the API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DisableAPIKey(disableAPIKeyOptions *DisableAPIKeyOptions)`, func() {
			Expect(apikeyId2).ToNot(BeEmpty())

			disableAPIKeyOptions := &iamidentityv1.DisableAPIKeyOptions{
				ID: &apikeyId2,
			}

			response, err := iamIdentityService.DisableAPIKey(disableAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			apiKey := getAPIkey(iamIdentityService, apikeyId2)
			Expect(apiKey).ToNot(BeNil())
			Expect(*apiKey.Disabled).To(BeTrue())
		})
	})

	Describe(`EnableAPIKey - Enable the API key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`EnableAPIKey(enableAPIKeyOptions *EnableAPIKeyOptions)`, func() {
			Expect(apikeyId2).ToNot(BeEmpty())

			enableAPIKeyOptions := &iamidentityv1.EnableAPIKeyOptions{
				ID: &apikeyId2,
			}

			response, err := iamIdentityService.EnableAPIKey(enableAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			apiKey := getAPIkey(iamIdentityService, apikeyId2)
			Expect(apiKey).ToNot(BeNil())
			Expect(*apiKey.Disabled).To(BeFalse())
		})
	})

	Describe(`DeleteAPIKey1 - Deletes an API key1`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAPIKey(deleteAPIKeyOptions *DeleteAPIKeyOptions)`, func() {
			Expect(apikeyId1).ToNot(BeEmpty())

			deleteAPIKeyOptions := &iamidentityv1.DeleteAPIKeyOptions{
				ID: &apikeyId1,
			}

			response, err := iamIdentityService.DeleteAPIKey(deleteAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			apiKey := getAPIkey(iamIdentityService, apikeyId1)
			Expect(apiKey).To(BeNil())
		})
	})

	Describe(`DeleteAPIKey2 - Deletes an API key2`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAPIKey(deleteAPIKeyOptions *DeleteAPIKeyOptions)`, func() {
			Expect(apikeyId2).ToNot(BeEmpty())

			deleteAPIKeyOptions := &iamidentityv1.DeleteAPIKeyOptions{
				ID: &apikeyId2,
			}

			response, err := iamIdentityService.DeleteAPIKey(deleteAPIKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			apiKey := getAPIkey(iamIdentityService, apikeyId2)
			Expect(apiKey).To(BeNil())
		})
	})

	Describe(`CreateServiceID - Create a service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateServiceID(createServiceIdOptions *CreateServiceIdOptions)`, func() {

			createServiceIDOptions := &iamidentityv1.CreateServiceIDOptions{
				AccountID:   &accountID,
				Name:        &serviceIDName,
				Description: core.StringPtr("GoSDK test serviceId"),
			}

			serviceID, response, err := iamIdentityService.CreateServiceID(createServiceIDOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(serviceID).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateServiceID response:\n%s\n", common.ToJSON(serviceID))

			serviceId1 = *serviceID.ID
			Expect(serviceId1).ToNot(BeNil())
		})
	})

	Describe(`GetServiceID - Get details of a service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetServiceID(getServiceIdOptions *GetServiceIdOptions)`, func() {
			Expect(serviceId1).ToNot(BeEmpty())
			getServiceIDOptions := &iamidentityv1.GetServiceIDOptions{
				ID:              &serviceId1,
				IncludeHistory:  core.BoolPtr(true),
				IncludeActivity: core.BoolPtr(true),
			}

			serviceID, response, err := iamIdentityService.GetServiceID(getServiceIDOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetServiceID response:\n%s\n", common.ToJSON(serviceID))

			Expect(*serviceID.Name).To(Equal(serviceIDName))
			Expect(*serviceID.Description).To(Equal("GoSDK test serviceId"))
			Expect(serviceID.History).ToNot(BeEmpty())
			Expect(serviceID.Activity).ToNot(BeNil())
			Expect(serviceID.Activity.AuthnCount).ToNot(BeNil())

			// Grab the Etag value from the response for use in the update operation.
			serviceIdEtag1 = response.GetHeaders().Get("Etag")
			Expect(serviceIdEtag1).ToNot(BeEmpty())
		})
	})

	Describe(`ListServiceIds - List service IDs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListServiceIds(listServiceIdsOptions *ListServiceIdsOptions)`, func() {

			listServiceIdsOptions := &iamidentityv1.ListServiceIdsOptions{
				AccountID: &accountID,
				Name:      &serviceIDName,
				Pagesize:  core.Int64Ptr(int64(100)),
			}

			serviceIdList, response, err := iamIdentityService.ListServiceIds(listServiceIdsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceIdList).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ListServiceIds response:\n%s\n", common.ToJSON(serviceIdList))

			Expect(len(serviceIdList.Serviceids)).To(Equal(1))
			Expect(serviceIdList.Offset).ToNot(BeNil())
			Expect(serviceIdList.Next).To(BeNil())
			Expect(*serviceIdList.Serviceids[0].Name).To(Equal(serviceIDName))
		})
	})

	Describe(`UpdateServiceID - Update service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateServiceID(updateServiceIdOptions *UpdateServiceIdOptions)`, func() {
			Expect(serviceId1).ToNot(BeEmpty())
			Expect(serviceIdEtag1).ToNot(BeEmpty())

			updateServiceIDOptions := &iamidentityv1.UpdateServiceIDOptions{
				ID:          &serviceId1,
				IfMatch:     &serviceIdEtag1,
				Description: &newDescription,
			}

			serviceID, response, err := iamIdentityService.UpdateServiceID(updateServiceIDOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "UpdateServiceID response:\n%s\n", common.ToJSON(serviceID))

			Expect(*serviceID.Description).To(Equal(newDescription))
		})
	})

	Describe(`LockServiceID - Lock the service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`LockServiceID(lockServiceIdOptions *LockServiceIdOptions)`, func() {
			Expect(serviceId1).ToNot(BeEmpty())

			lockServiceIDOptions := &iamidentityv1.LockServiceIDOptions{
				ID: &serviceId1,
			}

			response, err := iamIdentityService.LockServiceID(lockServiceIDOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
			fmt.Fprintf(GinkgoWriter, "LockServiceID response:\n%v\n", response)

			serviceID := getServiceID(iamIdentityService, serviceId1)
			Expect(serviceID).ToNot(BeNil())
			Expect(*serviceID.Locked).To(BeTrue())
		})
	})

	Describe(`UnlockServiceID - Unlock the service ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UnlockServiceID(unlockServiceIdOptions *UnlockServiceIdOptions)`, func() {
			Expect(serviceId1).ToNot(BeEmpty())

			unlockServiceIDOptions := &iamidentityv1.UnlockServiceIDOptions{
				ID: &serviceId1,
			}

			response, err := iamIdentityService.UnlockServiceID(unlockServiceIDOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
			fmt.Fprintf(GinkgoWriter, "UnlockServiceID response:\n%v\n", response)

			serviceID := getServiceID(iamIdentityService, serviceId1)
			Expect(serviceID).ToNot(BeNil())
			Expect(*serviceID.Locked).To(BeFalse())
		})
	})

	Describe(`DeleteServiceID - Deletes a service ID and associated API keys`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteServiceID(deleteServiceIdOptions *DeleteServiceIdOptions)`, func() {
			Expect(serviceId1).ToNot(BeEmpty())

			deleteServiceIDOptions := &iamidentityv1.DeleteServiceIDOptions{
				ID: &serviceId1,
			}

			response, err := iamIdentityService.DeleteServiceID(deleteServiceIDOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			serviceID := getServiceID(iamIdentityService, serviceId1)
			Expect(serviceID).To(BeNil())
		})
	})

	Describe(`CreateProfile1 - Create trusted profile #1`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProfile(createProfileOptions *CreateProfileOptions)`, func() {

			createProfileOptions := &iamidentityv1.CreateProfileOptions{
				Name:        &profileName1,
				Description: core.StringPtr("GoSDK test profile #1"),
				AccountID:   &accountID,
			}

			trustedProfile, response, err := iamIdentityService.CreateProfile(createProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(trustedProfile).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateProfile #1 response:\n%s\n", common.ToJSON(trustedProfile))

			profileId1 = *trustedProfile.ID
			profileIamId = *trustedProfile.IamID
			Expect(profileId1).ToNot(BeNil())
		})
	})

	Describe(`CreateProfile2 - Create trusted profile #2`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProfile(createProfileOptions *CreateProfileOptions)`, func() {

			createProfileOptions := &iamidentityv1.CreateProfileOptions{
				Name:        &profileName2,
				Description: core.StringPtr("GoSDK test profile #2"),
				AccountID:   &accountID,
			}

			trustedProfile, response, err := iamIdentityService.CreateProfile(createProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(trustedProfile).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateProfile #1 response:\n%s\n", common.ToJSON(trustedProfile))

			profileId2 = *trustedProfile.ID
			Expect(profileId2).ToNot(BeNil())
		})
	})

	Describe(`GetProfile - Get trusted profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfile(getProfileOptions *GetProfileOptions)`, func() {

			getProfileOptions := &iamidentityv1.GetProfileOptions{
				ProfileID: &profileId1,
			}

			trustedProfile, response, err := iamIdentityService.GetProfile(getProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trustedProfile).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetProfile #1 response:\n%s\n", common.ToJSON(trustedProfile))

			Expect(trustedProfile.ID).To(Equal(&profileId1))
			Expect(trustedProfile.IamID).To(Equal(&profileIamId))
			Expect(trustedProfile.AccountID).To(Equal(&accountID))
			Expect(trustedProfile.Name).To(Equal(&profileName1))
			Expect(trustedProfile.CRN).ToNot(BeNil())

			profileEtag = response.GetHeaders().Get("Etag")
			Expect(profileEtag).ToNot(BeNil())
		})
	})

	Describe(`ListProfiles - List trusted profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfiles(listProfilesOptions *ListProfilesOptions)`, func() {

			profiles := []iamidentityv1.TrustedProfile{}

			var pageTokenPresent bool = true
			var pageToken *string = nil

			for pageTokenPresent {

				listProfilesOptions := &iamidentityv1.ListProfilesOptions{
					AccountID:      &accountID,
					Pagetoken:      pageToken,
					Pagesize:       core.Int64Ptr(int64(1)),
					IncludeHistory: core.BoolPtr(false),
				}

				trustedProfiles, response, err := iamIdentityService.ListProfiles(listProfilesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(trustedProfiles).ToNot(BeNil())
				fmt.Fprintf(GinkgoWriter, "ListProfiles #1 response:\n%s\n", common.ToJSON(trustedProfiles))

				for _, trustedProfile := range trustedProfiles.Profiles {
					if profileName1 == *trustedProfile.Name || profileName2 == *trustedProfile.Name {
						profiles = append(profiles, trustedProfile)
					}
				}

				pageToken = getPageTokenFromURL(trustedProfiles.Next)
				pageTokenPresent = (pageToken != nil)
			}

			Expect(len(profiles)).To(Equal(2))
		})
	})

	Describe(`UpdateProfile - Update trusted profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateProfile(updateProfileOptions *UpdateProfileOptions)`, func() {

			updateProfileOptions := &iamidentityv1.UpdateProfileOptions{
				ProfileID:   &profileId1,
				IfMatch:     &profileEtag,
				Description: &newDescription,
			}

			trustedProfile, response, err := iamIdentityService.UpdateProfile(updateProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trustedProfile).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "UpdateProfile #1 response:\n%s\n", common.ToJSON(trustedProfile))

			Expect(*trustedProfile.ID).To(Equal(profileId1))
			Expect(*trustedProfile.Description).To(Equal(newDescription))
		})
	})

	Describe(`DeleteProfile1 - Delete trusted profile #1`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProfile(deleteProfileOptions *DeleteProfileOptions)`, func() {

			deleteProfileOptions := &iamidentityv1.DeleteProfileOptions{
				ProfileID: &profileId1,
			}

			response, err := iamIdentityService.DeleteProfile(deleteProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			profile := getProfile(iamIdentityService, profileId1)
			Expect(profile).To(BeNil())
		})
	})

	Describe(`CreateClaimRule1 - Create claim rule #1`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateClaimRule(createClaimRuleOptions *CreateClaimRuleOptions)`, func() {

			profileClaimRuleConditions := new(iamidentityv1.ProfileClaimRuleConditions)
			profileClaimRuleConditions.Claim = core.StringPtr("blueGroups")
			profileClaimRuleConditions.Operator = core.StringPtr("EQUALS")
			profileClaimRuleConditions.Value = core.StringPtr("\"cloud-docs-dev\"")

			createClaimRuleOptions := &iamidentityv1.CreateClaimRuleOptions{
				ProfileID:  &profileId2,
				Type:       &claimRuleType,
				RealmName:  &realmName,
				Expiration: core.Int64Ptr(int64(43200)),
				Conditions: []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditions},
			}

			claimRule, response, err := iamIdentityService.CreateClaimRule(createClaimRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(claimRule).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateClaimRule #1 response:\n%s\n", common.ToJSON(claimRule))

			claimRuleId1 = *claimRule.ID
			Expect(claimRuleId1).ToNot(BeNil())
		})
	})

	Describe(`CreateClaimRule2 - Create claim rule #2`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateClaimRule(createClaimRuleOptions *CreateClaimRuleOptions)`, func() {

			profileClaimRuleConditions := new(iamidentityv1.ProfileClaimRuleConditions)
			profileClaimRuleConditions.Claim = core.StringPtr("blueGroups")
			profileClaimRuleConditions.Operator = core.StringPtr("EQUALS")
			profileClaimRuleConditions.Value = core.StringPtr("\"Europe_Group\"")

			createClaimRuleOptions := &iamidentityv1.CreateClaimRuleOptions{
				ProfileID:  &profileId2,
				Type:       &claimRuleType,
				RealmName:  &realmName,
				Expiration: core.Int64Ptr(int64(43200)),
				Conditions: []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditions},
			}

			claimRule, response, err := iamIdentityService.CreateClaimRule(createClaimRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(claimRule).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateClaimRule #1 response:\n%s\n", common.ToJSON(claimRule))

			claimRuleId2 = *claimRule.ID
			Expect(claimRuleId2).ToNot(BeNil())
		})
	})

	Describe(`GetClaimRule - Get claim rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetClaimRule(getClaimRuleOptions *GetClaimRuleOptions)`, func() {

			getClaimRuleOptions := &iamidentityv1.GetClaimRuleOptions{
				ProfileID: &profileId2,
				RuleID:    &claimRuleId1,
			}

			claimRule, response, err := iamIdentityService.GetClaimRule(getClaimRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(claimRule).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetClaimRule #1 response:\n%s\n", common.ToJSON(claimRule))

			claimRuleEtag = response.GetHeaders().Get("Etag")
			Expect(claimRuleEtag).ToNot(BeNil())
		})
	})

	Describe(`ListClaimRules - List claim rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListClaimRules(listClaimRulesOptions *ListClaimRulesOptions)`, func() {

			claimRules := []iamidentityv1.ProfileClaimRule{}

			listClaimRulesOptions := &iamidentityv1.ListClaimRulesOptions{
				ProfileID: &profileId2,
			}

			claimRulesList, response, err := iamIdentityService.ListClaimRules(listClaimRulesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(claimRulesList).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ListClaimRules #1 response:\n%s\n", common.ToJSON(claimRulesList))

			for _, claimRule := range claimRulesList.Rules {
				if claimRuleId1 == *claimRule.ID || claimRuleId2 == *claimRule.ID {
					claimRules = append(claimRules, claimRule)
				}
			}

			Expect(len(claimRules)).To(Equal(2))
		})
	})

	Describe(`UpdateClaimRule - Update claim rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateClaimRule(updateClaimRuleOptions *UpdateClaimRuleOptions)`, func() {

			profileClaimRuleConditions := new(iamidentityv1.ProfileClaimRuleConditions)
			profileClaimRuleConditions.Claim = core.StringPtr("blueGroups")
			profileClaimRuleConditions.Operator = core.StringPtr("EQUALS")
			profileClaimRuleConditions.Value = core.StringPtr("\"Europe_Group\"")

			updateClaimRuleOptions := &iamidentityv1.UpdateClaimRuleOptions{
				ProfileID:  &profileId2,
				RuleID:     &claimRuleId1,
				IfMatch:    &claimRuleEtag,
				Expiration: core.Int64Ptr(int64(33200)),
				Type:       &claimRuleType,
				RealmName:  &realmName,
				Conditions: []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditions},
			}

			claimRule, response, err := iamIdentityService.UpdateClaimRule(updateClaimRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(claimRule).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "UpdateProfile #1 response:\n%s\n", common.ToJSON(claimRule))

		})
	})

	Describe(`DeleteClaimRule1 - Delete claim rule #1`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteClaimRule(deleteClaimRuleOptions *DeleteClaimRuleOptions)`, func() {

			deleteClaimRuleOptions := &iamidentityv1.DeleteClaimRuleOptions{
				ProfileID: &profileId2,
				RuleID:    &claimRuleId1,
			}

			response, err := iamIdentityService.DeleteClaimRule(deleteClaimRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			profile := getClaimRule(iamIdentityService, profileId2, claimRuleId1)
			Expect(profile).To(BeNil())
		})
	})

	Describe(`DeleteClaimRule2 - Delete claim rule #2`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteClaimRule(deleteClaimRuleOptions *DeleteClaimRuleOptions)`, func() {

			deleteClaimRuleOptions := &iamidentityv1.DeleteClaimRuleOptions{
				ProfileID: &profileId2,
				RuleID:    &claimRuleId2,
			}

			response, err := iamIdentityService.DeleteClaimRule(deleteClaimRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			profile := getClaimRule(iamIdentityService, profileId2, claimRuleId2)
			Expect(profile).To(BeNil())
		})
	})

	Describe(`CreateLink - Create link #2`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateLink(createLinkOptions *CreateLinkOptions)`, func() {

			createProfileLinkRequestLink := new(iamidentityv1.CreateProfileLinkRequestLink)
			createProfileLinkRequestLink.CRN = core.StringPtr("crn:v1:staging:public:iam-identity::a/" + accountID + "::computeresource:Fake-Compute-Resource")
			createProfileLinkRequestLink.Namespace = core.StringPtr("default")
			createProfileLinkRequestLink.Name = core.StringPtr("nice name")

			createLinkOptions := &iamidentityv1.CreateLinkOptions{
				ProfileID: &profileId2,
				Name:      core.StringPtr("niceLink"),
				CrType:    core.StringPtr("ROKS_SA"),
				Link:      createProfileLinkRequestLink,
			}

			link, response, err := iamIdentityService.CreateLink(createLinkOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(link).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateClaimRule #1 response:\n%s\n", common.ToJSON(link))

			linkId = *link.ID
			Expect(linkId).ToNot(BeNil())
		})
	})

	Describe(`GetLink - Get link`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLink(getLinkOptions *GetLinkOptions)`, func() {

			getLinkOptions := &iamidentityv1.GetLinkOptions{
				ProfileID: &profileId2,
				LinkID:    &linkId,
			}

			link, response, err := iamIdentityService.GetLink(getLinkOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(link).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetLink #1 response:\n%s\n", common.ToJSON(link))

			Expect(link.ID).To(Equal(&linkId))
			Expect(link.Link).ToNot(BeNil())
		})
	})

	Describe(`ListLinks - List link`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLinks(listLinksOptions *ListLinksOptions)`, func() {

			links := []iamidentityv1.ProfileLink{}

			listLinksOptions := &iamidentityv1.ListLinksOptions{
				ProfileID: &profileId2,
			}

			linkList, response, err := iamIdentityService.ListLinks(listLinksOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(linkList).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ListLinks response:\n%s\n", common.ToJSON(linkList))

			for _, link := range linkList.Links {
				if linkId == *link.ID {
					links = append(links, link)
				}
			}
			Expect(len(links)).To(Equal(1))
		})
	})

	Describe(`DeleteLink - Delete link`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteLink(deleteLinkOptions *DeleteLinkOptions)`, func() {

			deleteLinkOptions := &iamidentityv1.DeleteLinkOptions{
				ProfileID: &profileId2,
				LinkID:    &linkId,
			}

			response, err := iamIdentityService.DeleteLink(deleteLinkOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			link := getLink(iamIdentityService, profileId2, linkId)
			Expect(link).To(BeNil())
		})
	})

	Describe(`SetProfileIdentities - Set Profile Identities`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetProfileIdentities(setProfileIdentitiesOptions *SetProfileIdentitiesOptions)`, func() {

			accounts := []string{accountID}
			identity := &iamidentityv1.ProfileIdentityRequest{
				Identifier:  &iamID,
				Accounts:    accounts,
				Type:        core.StringPtr("user"),
				Description: core.StringPtr("Identity description"),
			}
			listProfileIdentity := []iamidentityv1.ProfileIdentityRequest{*identity}
			setProfileIdentitiesOptions := iamidentityv1.SetProfileIdentitiesOptions{
				ProfileID:  &profileId2,
				Identities: listProfileIdentity,
				IfMatch:    core.StringPtr("*"),
			}

			profileIdentities, response, err := iamIdentityService.SetProfileIdentities(&setProfileIdentitiesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

			Expect(profileIdentities.Identities).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "SetProfileIdentities #1 response:\n%s\n", common.ToJSON(profileIdentities))
		})
	})

	Describe(`GetProfileIdentities - Get Profile Identities`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfileIdentities(getProfileIdentitiesOptions *GetProfileIdentitiesOptions)`, func() {

			getProfileIdentitiesOptions := iamidentityv1.GetProfileIdentitiesOptions{
				ProfileID: &profileId2,
			}

			profileIdentities, response, err := iamIdentityService.GetProfileIdentities(&getProfileIdentitiesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

			Expect(profileIdentities.Identities).ToNot(BeNil())
			Expect(profileIdentities.Identities[0].Identifier).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "SetProfileIdentities #1 response:\n%s\n", common.ToJSON(profileIdentities))
		})
	})

	Describe(`SetProfileIdentity - Set single identity on Trusted Profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetProfileIdentity(setProfileIdentityOptions *SetProfileIdentityOptions)`, func() {

			accounts := []string{accountID}

			setProfileIdentityOptions := iamidentityv1.SetProfileIdentityOptions{
				ProfileID:    &profileId2,
				IdentityType: core.StringPtr("user"),
				Identifier:   &iamIDMember,
				Accounts:     accounts,
				Type:         core.StringPtr("user"),
				Description:  core.StringPtr("Identity description"),
			}

			profileIdnetity, response, err := iamIdentityService.SetProfileIdentity(&setProfileIdentityOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

			Expect(profileIdnetity).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "SetProfileIdentities #1 response:\n%s\n", common.ToJSON(profileIdnetity))
		})
	})

	Describe(`GetProfileIdentity - Get specific identity of the profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfileIdentity(getProfileIdentityOptions *GetProfileIdentityOptions)`, func() {

			getProfileIdentityOptions := iamidentityv1.GetProfileIdentityOptions{
				ProfileID:    &profileId2,
				IdentityType: core.StringPtr("user"),
				IdentifierID: &iamIDMember,
			}

			profileIdnetity, response, err := iamIdentityService.GetProfileIdentity(&getProfileIdentityOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

			Expect(profileIdnetity.Identifier).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "SetProfileIdentities #1 response:\n%s\n", common.ToJSON(profileIdnetity))
		})
	})

	Describe(`DeleteProfileIdentity - Delete specific identity of the profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProfileIdentity(deleteProfileIdentityOptions *DeleteProfileIdentityOptions)`, func() {

			deleteProfileIdentityOptions := iamidentityv1.DeleteProfileIdentityOptions{
				ProfileID:    &profileId2,
				IdentityType: core.StringPtr("user"),
				IdentifierID: &iamIDMember,
			}

			response, err := iamIdentityService.DeleteProfileIdentity(&deleteProfileIdentityOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteProfile2 - Delete trusted profile #2`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProfile(deleteProfileOptions *DeleteProfileOptions)`, func() {

			deleteProfileOptions := &iamidentityv1.DeleteProfileOptions{
				ProfileID: &profileId2,
			}

			response, err := iamIdentityService.DeleteProfile(deleteProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			profile := getProfile(iamIdentityService, profileId2)
			Expect(profile).To(BeNil())
		})
	})

	Describe(`CreateProfileBadRequest`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProfileBadRequest(createProfileOptions *CreateProfileOptions)`, func() {

			createProfileOptions := &iamidentityv1.CreateProfileOptions{
				Name:        &profileName1,
				Description: core.StringPtr("GoSDK test profile #1"),
				AccountID:   core.StringPtr("InvalidID"),
			}

			trustedProfile, response, err := iamIdentityService.CreateProfile(createProfileOptions)

			Expect(err).ToNot(BeNil())
			Expect(trustedProfile).To(BeNil())
			Expect(response.StatusCode).To(Equal(400))

		})
	})

	Describe(`GetProfileNotFound`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfileNotFound(getProfileOptions *GetProfileOptions)`, func() {

			getProfileOptions := &iamidentityv1.GetProfileOptions{
				ProfileID: core.StringPtr("InvalidID"),
			}

			trustedProfile, response, err := iamIdentityService.GetProfile(getProfileOptions)

			Expect(trustedProfile).To(BeNil())
			Expect(response.StatusCode).To(Equal(404))
			Expect(err).ToNot(BeNil())
		})
	})

	Describe(`UpdateProfileNotFound`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateProfileNotFound(updateProfileOptions *UpdateProfileOptions)`, func() {

			updateProfileOptions := &iamidentityv1.UpdateProfileOptions{
				ProfileID:   core.StringPtr("InvalidID"),
				IfMatch:     core.StringPtr("dummy"),
				Description: core.StringPtr("dummy"),
			}

			trustedProfile, response, err := iamIdentityService.UpdateProfile(updateProfileOptions)

			Expect(trustedProfile).To(BeNil())
			Expect(response.StatusCode).To(Equal(404))
			Expect(err).ToNot(BeNil())
		})
	})

	Describe(`DeleteProfileNotFound`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProfileNotFound(deleteProfileOptions *DeleteProfileOptions)`, func() {

			deleteProfileOptions := &iamidentityv1.DeleteProfileOptions{
				ProfileID: core.StringPtr("InvalidID"),
			}

			response, err := iamIdentityService.DeleteProfile(deleteProfileOptions)

			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`CreateClaimRuleNotFound`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateClaimRuleNotFound(createClaimRuleOptions *CreateClaimRuleOptions)`, func() {

			profileClaimRuleConditions := new(iamidentityv1.ProfileClaimRuleConditions)
			profileClaimRuleConditions.Claim = core.StringPtr("dummy")
			profileClaimRuleConditions.Operator = core.StringPtr("EQUALS")
			profileClaimRuleConditions.Value = core.StringPtr("\"dummy\"")

			createClaimRuleOptions := &iamidentityv1.CreateClaimRuleOptions{
				ProfileID:  core.StringPtr("InvalidID"),
				Type:       &claimRuleType,
				RealmName:  &realmName,
				Expiration: core.Int64Ptr(int64(43200)),
				Conditions: []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditions},
			}

			claimRule, response, err := iamIdentityService.CreateClaimRule(createClaimRuleOptions)

			Expect(claimRule).To(BeNil())
			Expect(response.StatusCode).To(Equal(404))
			Expect(err).ToNot(BeNil())
		})
	})

	Describe(`GetClaimRuleNotFound`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetClaimRuleNotFound(getClaimRuleOptions *GetClaimRuleOptions)`, func() {

			getClaimRuleOptions := &iamidentityv1.GetClaimRuleOptions{
				ProfileID: core.StringPtr("InvalidID"),
				RuleID:    core.StringPtr("InvalidID"),
			}

			claimRule, response, err := iamIdentityService.GetClaimRule(getClaimRuleOptions)

			Expect(claimRule).To(BeNil())
			Expect(response.StatusCode).To(Equal(404))
			Expect(err).ToNot(BeNil())
		})
	})

	Describe(`UpdateClaimRuleBadRequest`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateClaimRuleBadRequest(updateClaimRuleOptions *UpdateClaimRuleOptions)`, func() {

			profileClaimRuleConditions := new(iamidentityv1.ProfileClaimRuleConditions)
			profileClaimRuleConditions.Claim = core.StringPtr("blueGroups")
			profileClaimRuleConditions.Operator = core.StringPtr("EQUALS")
			profileClaimRuleConditions.Value = core.StringPtr("\"Europe_Group\"")

			updateClaimRuleOptions := &iamidentityv1.UpdateClaimRuleOptions{
				ProfileID:  core.StringPtr("InvalidID"),
				RuleID:     core.StringPtr("InvalidID"),
				IfMatch:    core.StringPtr("dummy"),
				Expiration: core.Int64Ptr(int64(33200)),
				Type:       core.StringPtr(""),
				RealmName:  core.StringPtr(""),
				Conditions: []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditions},
			}

			claimRule, response, err := iamIdentityService.UpdateClaimRule(updateClaimRuleOptions)

			Expect(claimRule).To(BeNil())
			Expect(response.StatusCode).To(Equal(400))
			Expect(err).ToNot(BeNil())
		})
	})

	Describe(`DeleteClaimRuleNotFound`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteClaimRuleNotFound(deleteClaimRuleOptions *DeleteClaimRuleOptions)`, func() {

			deleteClaimRuleOptions := &iamidentityv1.DeleteClaimRuleOptions{
				ProfileID: core.StringPtr("InvalidID"),
				RuleID:    core.StringPtr("InvalidID"),
			}

			response, err := iamIdentityService.DeleteClaimRule(deleteClaimRuleOptions)

			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`CreateLinkBadRequest`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateLinkBadRequest(createLinkOptions *CreateLinkOptions)`, func() {

			createProfileLinkRequestLink := new(iamidentityv1.CreateProfileLinkRequestLink)
			createProfileLinkRequestLink.CRN = core.StringPtr("crn:v1:staging:public:iam-identity::a/" + accountID + "::computeresource:Fake-Compute-Resource")
			createProfileLinkRequestLink.Namespace = core.StringPtr("default")
			createProfileLinkRequestLink.Name = core.StringPtr("nice name")

			createLinkOptions := &iamidentityv1.CreateLinkOptions{
				ProfileID: core.StringPtr("invalidId"),
				Name:      core.StringPtr("dummy"),
				CrType:    core.StringPtr("dummy"),
				Link:      createProfileLinkRequestLink,
			}

			link, response, err := iamIdentityService.CreateLink(createLinkOptions)

			Expect(link).To(BeNil())
			Expect(response.StatusCode).To(Equal(400))
			Expect(err).ToNot(BeNil())
		})
	})

	Describe(`GetLinkBadRequest`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLinkBadRequest(getLinkOptions *GetLinkOptions)`, func() {

			getLinkOptions := &iamidentityv1.GetLinkOptions{
				ProfileID: core.StringPtr("invalidId"),
				LinkID:    core.StringPtr("invalidId"),
			}

			link, response, err := iamIdentityService.GetLink(getLinkOptions)

			Expect(link).To(BeNil())
			Expect(response.StatusCode).To(Equal(404))
			Expect(err).ToNot(BeNil())
		})
	})

	Describe(`DeleteLinkNotFound`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteLinkNotFound(deleteLinkOptions *DeleteLinkOptions)`, func() {

			deleteLinkOptions := &iamidentityv1.DeleteLinkOptions{
				ProfileID: core.StringPtr("invalidId"),
				LinkID:    core.StringPtr("invalidId"),
			}

			response, err := iamIdentityService.DeleteLink(deleteLinkOptions)

			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`GetAccountSettings - Get account configurations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions)`, func() {

			getAccountSettingsOptions := &iamidentityv1.GetAccountSettingsOptions{
				AccountID:      core.StringPtr(accountID),
				IncludeHistory: core.BoolPtr(true),
			}

			accountSettingsResponse, response, err := iamIdentityService.GetAccountSettings(getAccountSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettingsResponse).ToNot(BeNil())
			Expect(accountSettingsResponse.History).ToNot(BeNil())
			Expect(accountSettingsResponse.EntityTag).ToNot(BeNil())
			Expect(accountSettingsResponse.RestrictCreateServiceID).ToNot(BeNil())
			Expect(accountSettingsResponse.RestrictCreatePlatformApikey).ToNot(BeNil())
			Expect(accountSettingsResponse.SessionExpirationInSeconds).ToNot(BeNil())
			Expect(accountSettingsResponse.SessionInvalidationInSeconds).ToNot(BeNil())
			Expect(accountSettingsResponse.SystemAccessTokenExpirationInSeconds).ToNot(BeNil())
			Expect(accountSettingsResponse.SystemRefreshTokenExpirationInSeconds).ToNot(BeNil())
			Expect(accountSettingsResponse.Mfa).ToNot(BeNil())
			Expect(accountSettingsResponse.UserMfa).ToNot(BeNil())

			accountSettingEtag = response.GetHeaders().Get("Etag")
			Expect(accountSettingEtag).ToNot(BeEmpty())
		})
	})

	Describe(`UpdateAccountSettings - Update account configurations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAccountSettings(updateAccountSettingsOptions *UpdateAccountSettingsOptions)`, func() {

			accountSettingsUserMFA := new(iamidentityv1.AccountSettingsUserMfa)
			accountSettingsUserMFA.IamID = core.StringPtr(iamIDMember)
			accountSettingsUserMFA.Mfa = core.StringPtr("NONE")

			accountSettingsRequestOptions := &iamidentityv1.UpdateAccountSettingsOptions{
				IfMatch:                      core.StringPtr(accountSettingEtag),
				AccountID:                    core.StringPtr(accountID),
				RestrictCreateServiceID:      core.StringPtr("NOT_RESTRICTED"),
				RestrictCreatePlatformApikey: core.StringPtr("NOT_RESTRICTED"),
				//AllowedIPAddresses:                  core.StringPtr("testString"),
				Mfa:                                   core.StringPtr("NONE"),
				UserMfa:                               []iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMFA},
				SessionExpirationInSeconds:            core.StringPtr("86400"),
				SessionInvalidationInSeconds:          core.StringPtr("7200"),
				MaxSessionsPerIdentity:                core.StringPtr("10"),
				SystemAccessTokenExpirationInSeconds:  core.StringPtr("3600"),
				SystemRefreshTokenExpirationInSeconds: core.StringPtr("259200"),
			}

			accountSettingsResponse, response, err := iamIdentityService.UpdateAccountSettings(accountSettingsRequestOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettingsResponse).ToNot(BeNil())
			Expect(accountSettingsResponse.History).ToNot(BeNil())
			Expect(accountSettingsResponse.EntityTag).ToNot(Equal(accountSettingEtag))
			Expect(accountSettingsResponse.Mfa).To(Equal(accountSettingsRequestOptions.Mfa))
			Expect(accountSettingsResponse.UserMfa).To(Equal(accountSettingsRequestOptions.UserMfa))
			Expect(accountSettingsResponse.AccountID).To(Equal(accountSettingsRequestOptions.AccountID))
			Expect(accountSettingsResponse.RestrictCreateServiceID).To(Equal(accountSettingsRequestOptions.RestrictCreateServiceID))
			Expect(accountSettingsResponse.RestrictCreatePlatformApikey).To(Equal(accountSettingsRequestOptions.RestrictCreatePlatformApikey))
			Expect(accountSettingsResponse.SessionInvalidationInSeconds).To(Equal(accountSettingsRequestOptions.SessionInvalidationInSeconds))
			Expect(accountSettingsResponse.SessionExpirationInSeconds).To(Equal(accountSettingsRequestOptions.SessionExpirationInSeconds))
			Expect(accountSettingsResponse.SystemAccessTokenExpirationInSeconds).To(Equal(accountSettingsRequestOptions.SystemAccessTokenExpirationInSeconds))
			Expect(accountSettingsResponse.SystemRefreshTokenExpirationInSeconds).To(Equal(accountSettingsRequestOptions.SystemRefreshTokenExpirationInSeconds))
			fmt.Fprintf(GinkgoWriter, "UpdateAccountSettings response:\n%s\n", common.ToJSON(accountSettingsResponse))
		})
	})

	Describe(`GetEffectiveAccountSettings - Get effective account settings configurations`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEffectiveAccountSettings(getEffectiveAccountSettingsOptions *GetEffectiveAccountSettingsOptions)`, func() {

			getEffectiveAccountSettingsOptions := &iamidentityv1.GetEffectiveAccountSettingsOptions{
				AccountID:      core.StringPtr(accountID),
				IncludeHistory: core.BoolPtr(true),
			}

			effectiveAccountSettingsResponse, response, err := iamIdentityService.GetEffectiveAccountSettings(getEffectiveAccountSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(effectiveAccountSettingsResponse).ToNot(BeNil())
			Expect(effectiveAccountSettingsResponse.AccountID).ToNot(BeNil())
			Expect(effectiveAccountSettingsResponse.Effective).ToNot(BeNil())
			Expect(effectiveAccountSettingsResponse.Account).ToNot(BeNil())
		})
	})

	Describe(`CreateInactivityReport - Create an inactivity report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateReport(createReportOptions *CreateReportOptions)`, func() {

			createReportOptions := &iamidentityv1.CreateReportOptions{
				AccountID: &accountID,
				Type:      core.StringPtr("inactive"),
				Duration:  core.StringPtr("120"),
			}

			reportRef, response, err := iamIdentityService.CreateReport(createReportOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(reportRef).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateReport response:\n%s\n", common.ToJSON(reportRef))

			reportId = *reportRef.Reference
			Expect(reportId).ToNot(BeNil())
		})
	})

	Describe(`GetInactivityReportIncomplete - Get an incomplete inactivity report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReport(getReportOptions *GetReportOptions)`, func() {
			Expect(reportId).ToNot(BeEmpty())
			getReportOptions := &iamidentityv1.GetReportOptions{
				AccountID: &accountID,
				Reference: &reportId,
			}

			report, response, err := iamIdentityService.GetReport(getReportOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
			Expect(report).To(BeNil())
		})
	})

	Describe(`GetInactivityReportComplete - Get a complete inactivity report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReport(getReportOptions *GetReportOptions)`, func() {
			Expect(reportId).ToNot(BeEmpty())
			getReportOptions := &iamidentityv1.GetReportOptions{
				AccountID: &accountID,
				Reference: &reportId,
			}

			for i := 0; i < 30; i++ {
				report, response, err := iamIdentityService.GetReport(getReportOptions)
				Expect(err).To(BeNil())
				if response.StatusCode != 204 {
					Expect(response.StatusCode).To(Equal(200))
					Expect(report).ToNot(BeNil())
					Expect(report.CreatedBy).ToNot(BeNil())
					Expect(*report.CreatedBy).To(Equal(iamID))
					Expect(report.Reference).ToNot(BeNil())
					Expect(*report.Reference).To(Equal(reportId))
					Expect(report.ReportDuration).ToNot(BeNil())
					Expect(report.ReportStartTime).ToNot(BeNil())
					Expect(report.ReportEndTime).ToNot(BeNil())
					Expect(report.Users).ToNot(BeNil())
					Expect(report.Apikeys).ToNot(BeNil())
					Expect(report.Serviceids).ToNot(BeNil())
					Expect(report.Profiles).ToNot(BeNil())
					break
				}
				time.Sleep(1 * time.Second)
			}
		})
	})

	Describe(`GetInactivityReportNotFound`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportNotFound(getReportOptions *GetReportOptions)`, func() {

			getReportOptions := &iamidentityv1.GetReportOptions{
				AccountID: &accountID,
				Reference: core.StringPtr("1234567890"),
			}

			report, response, err := iamIdentityService.GetReport(getReportOptions)

			Expect(report).To(BeNil())
			Expect(response.StatusCode).To(Equal(404))
			Expect(err).ToNot(BeNil())
		})
	})

	Describe(`CreateReportMfa - Create an mfa report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateReportMfa(createMfReportOptions *createMfReportOptions)`, func() {

			createMfaReportOptions := &iamidentityv1.CreateMfaReportOptions{
				AccountID: &accountID,
				Type:      core.StringPtr("mfa_status"),
			}

			reportRef, response, err := iamIdentityService.CreateMfaReport(createMfaReportOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(reportRef).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateMfaReport response:\n%s\n", common.ToJSON(reportRef))

			reportReferenceMfa = *reportRef.Reference
			Expect(reportReferenceMfa).ToNot(BeNil())
		})
	})

	Describe(`GetReportMfaComplete - Get a complete mfa report`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetMfaReport(getMfaReportOptions *GetMfaReportOptions)`, func() {
			Expect(reportId).ToNot(BeEmpty())
			getMfaReportOptions := &iamidentityv1.GetMfaReportOptions{
				AccountID: &accountID,
				Reference: &reportReferenceMfa,
			}

			for i := 0; i < 30; i++ {
				report, response, err := iamIdentityService.GetMfaReport(getMfaReportOptions)
				Expect(err).To(BeNil())
				if response.StatusCode != 204 {
					Expect(response.StatusCode).To(Equal(200))
					Expect(report).ToNot(BeNil())
					Expect(report.CreatedBy).ToNot(BeNil())
					Expect(*report.CreatedBy).To(Equal(iamID))
					Expect(report.Reference).ToNot(BeNil())
					Expect(*report.Reference).To(Equal(reportReferenceMfa))
					break
				}
				time.Sleep(1 * time.Second)
			}
		})
	})

	Describe(`GetReportMfaNotFound`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportMfaNotFound(getMfaReportOptions *GetMfaReportOptions)`, func() {

			getMfaReportOptions := &iamidentityv1.GetMfaReportOptions{
				AccountID: &accountID,
				Reference: core.StringPtr("1234567890"),
			}

			report, response, err := iamIdentityService.GetMfaReport(getMfaReportOptions)

			Expect(report).To(BeNil())
			Expect(response.StatusCode).To(Equal(404))
			Expect(err).ToNot(BeNil())
		})
	})

	Describe(`GetMfaStatus`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetMfaStatus(getMfaStatusOptions *getMfaStatusOptions)`, func() {

			getMfaStatusOptions := &iamidentityv1.GetMfaStatusOptions{
				AccountID: &accountID,
				IamID:     &iamID,
			}

			mfaStatusResponse, response, err := iamIdentityService.GetMfaStatus(getMfaStatusOptions)

			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(mfaStatusResponse).ToNot(BeNil())
		})
	})

	Describe(`CreateProfileTemplateIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateProfileTemplateIT`, func() {
			profileClaimRuleConditions := new(iamidentityv1.ProfileClaimRuleConditions)
			profileClaimRuleConditions.Claim = core.StringPtr("blueGroups")
			profileClaimRuleConditions.Operator = core.StringPtr("EQUALS")
			profileClaimRuleConditions.Value = core.StringPtr("\"cloud-docs-dev\"")

			profileTemplateClaimRule := new(iamidentityv1.TrustedProfileTemplateClaimRule)
			profileTemplateClaimRule.Name = core.StringPtr("My Rule")
			profileTemplateClaimRule.RealmName = &realmName
			profileTemplateClaimRule.Type = &claimRuleType
			profileTemplateClaimRule.Expiration = core.Int64Ptr(int64(43200))
			profileTemplateClaimRule.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditions}

			profile := new(iamidentityv1.TemplateProfileComponentRequest)
			profile.Name = &profileTemplateProfileName
			profile.Description = core.StringPtr("GoSDK test Profile cretaed from Profile Template")
			profile.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*profileTemplateClaimRule}

			createOptions := &iamidentityv1.CreateProfileTemplateOptions{
				Name:        &profileTemplateName,
				Description: core.StringPtr("GoSDK test Profile Template"),
				AccountID:   &enterpriseAccountID,
				Profile:     profile,
			}

			createResponse, response, err := iamIdentityService.CreateProfileTemplate(createOptions)
			Expect(response.StatusCode).To(Equal(201))
			Expect(err).To(BeNil())
			Expect(createResponse).ToNot(BeNil())

			// Grab the ID and Etag value from the response for use in the update operation.
			profileTemplateId = *createResponse.ID
			Expect(profileTemplateId).ToNot(BeNil())
			profileTemplateVersion = *createResponse.Version
			Expect(response.GetHeaders()).ToNot(BeNil())
			profileTemplateEtag = response.GetHeaders().Get("Etag")
			Expect(profileTemplateEtag).ToNot(BeEmpty())
		})
	})

	Describe(`GetProfileTemplateIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetProfileTemplateIT`, func() {
			getOptions := &iamidentityv1.GetProfileTemplateVersionOptions{
				TemplateID: &profileTemplateId,
				Version:    core.StringPtr(strconv.FormatInt(profileTemplateVersion, 10)),
			}
			getResponse, response, err := iamIdentityService.GetProfileTemplateVersion(getOptions)
			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(getResponse).ToNot(BeNil())
			// Grab the Etag value from the response for use in the update operation.
			Expect(response.GetHeaders()).ToNot(BeNil())
			profileTemplateEtag = response.GetHeaders().Get("Etag")
			Expect(profileTemplateEtag).ToNot(BeEmpty())
		})
	})

	Describe(`ListProfileTemplatesIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfileTemplatesIT`, func() {
			listOptions := &iamidentityv1.ListProfileTemplatesOptions{
				AccountID: &enterpriseAccountID,
			}
			listResponse, response, err := iamIdentityService.ListProfileTemplates(listOptions)
			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(listResponse).ToNot(BeNil())
		})
	})

	Describe(`UpdateProfileTemplateIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateProfileTemplateIT`, func() {
			updateOptions := &iamidentityv1.UpdateProfileTemplateVersionOptions{
				AccountID:   &enterpriseAccountID,
				TemplateID:  &profileTemplateId,
				Version:     core.StringPtr(strconv.FormatInt(profileTemplateVersion, 10)),
				IfMatch:     &profileTemplateEtag,
				Name:        &profileTemplateName,
				Description: core.StringPtr("GoSDK test Profile Template - updated"),
			}
			updateResponse, response, err := iamIdentityService.UpdateProfileTemplateVersion(updateOptions)
			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(updateResponse).ToNot(BeNil())
			// Grab the Etag value from the response for use in the update operation.
			Expect(response.GetHeaders()).ToNot(BeNil())
			profileTemplateEtag = response.GetHeaders().Get("Etag")
			Expect(profileTemplateEtag).ToNot(BeEmpty())
		})
	})

	Describe(`AssignProfileTemplateIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AssignProfileTemplateIT`, func() {
			commitOptions := &iamidentityv1.CommitProfileTemplateOptions{
				TemplateID: &profileTemplateId,
				Version:    core.StringPtr(strconv.FormatInt(profileTemplateVersion, 10)),
			}
			cResponse, cErr := iamIdentityService.CommitProfileTemplate(commitOptions)
			Expect(cResponse.StatusCode).To(Equal(204))
			Expect(cErr).To(BeNil())

			assignOptions := &iamidentityv1.CreateTrustedProfileAssignmentOptions{
				TemplateID:      &profileTemplateId,
				TemplateVersion: &profileTemplateVersion,
				TargetType:      core.StringPtr("Account"),
				Target:          &enterpriseSubAccountID,
			}
			assignResponse, aResponse, aErr := iamIdentityService.CreateTrustedProfileAssignment(assignOptions)
			Expect(aResponse.StatusCode).To(Equal(202))
			Expect(aErr).To(BeNil())
			Expect(assignResponse).ToNot(BeNil())
			// Grab the Etag and id for use by other test methods.
			Expect(aResponse.GetHeaders()).ToNot(BeNil())
			profileTemplateAssignmentEtag = aResponse.GetHeaders().Get("Etag")
			Expect(profileTemplateAssignmentEtag).ToNot(BeEmpty())
			profileTemplateAssignmentId = *assignResponse.ID
			Expect(profileTemplateAssignmentId).ToNot(BeNil())
		})
	})

	Describe(`ListProfileTemplateAssignmentsIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfileTemplateAssignmentsIT`, func() {
			listOptions := &iamidentityv1.ListTrustedProfileAssignmentsOptions{
				AccountID:  &enterpriseAccountID,
				TemplateID: &profileTemplateId,
			}
			listResponse, response, err := iamIdentityService.ListTrustedProfileAssignments(listOptions)
			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(listResponse).ToNot(BeNil())
			Expect(listResponse.Assignments).ToNot(BeNil())
			numAssignments := len(listResponse.Assignments)
			Expect(numAssignments).ToNot(BeZero())
		})
	})

	Describe(`CreateNewProfileTemplateVersionIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNewProfileTemplateVersionIT`, func() {
			profileClaimRuleConditions := new(iamidentityv1.ProfileClaimRuleConditions)
			profileClaimRuleConditions.Claim = core.StringPtr("blueGroups")
			profileClaimRuleConditions.Operator = core.StringPtr("EQUALS")
			profileClaimRuleConditions.Value = core.StringPtr("\"cloud-docs-dev\"")

			profileTemplateClaimRule := new(iamidentityv1.TrustedProfileTemplateClaimRule)
			profileTemplateClaimRule.Name = core.StringPtr("My Rule")
			profileTemplateClaimRule.RealmName = &realmName
			profileTemplateClaimRule.Type = &claimRuleType
			profileTemplateClaimRule.Expiration = core.Int64Ptr(int64(43200))
			profileTemplateClaimRule.Conditions = []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditions}

			profile := new(iamidentityv1.TemplateProfileComponentRequest)
			profile.Name = &profileTemplateProfileName
			profile.Description = core.StringPtr("GoSDK test Profile cretaed from Profile Template - new version")
			profile.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*profileTemplateClaimRule}

			createOptions := &iamidentityv1.CreateProfileTemplateVersionOptions{
				Name:        &profileTemplateName,
				Description: core.StringPtr("GoSDK test Profile Template - new version"),
				AccountID:   &enterpriseAccountID,
				TemplateID:  &profileTemplateId,
				Profile:     profile,
			}

			createResponse, response, err := iamIdentityService.CreateProfileTemplateVersion(createOptions)
			Expect(response.StatusCode).To(Equal(201))
			Expect(err).To(BeNil())
			Expect(createResponse).ToNot(BeNil())
			// save the new version to be used in subsequent calls
			profileTemplateVersion = *createResponse.Version
			Expect(response.GetHeaders()).ToNot(BeNil())
		})
	})

	Describe(`GetLatestProfileTemplateVersionIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLatestProfileTemplateVersionIT`, func() {
			getOptions := &iamidentityv1.GetLatestProfileTemplateVersionOptions{
				TemplateID: &profileTemplateId,
			}
			getResponse, response, err := iamIdentityService.GetLatestProfileTemplateVersion(getOptions)
			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(getResponse).ToNot(BeNil())
		})
	})

	Describe(`ListProfileTemplateVersionsIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListProfileTemplateVersionsIT`, func() {
			listOptions := &iamidentityv1.ListVersionsOfProfileTemplateOptions{
				TemplateID: &profileTemplateId,
			}
			listResponse, response, err := iamIdentityService.ListVersionsOfProfileTemplate(listOptions)
			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(listResponse).ToNot(BeNil())
			Expect(listResponse.ProfileTemplates).ToNot(BeNil())
			numTemplates := len(listResponse.ProfileTemplates)
			Expect(numTemplates).ToNot(BeZero())
		})
	})

	Describe(`UpdateProfileTemplateAssignmentIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateProfileTemplateAssignmentIT`, func() {
			commitOptions := &iamidentityv1.CommitProfileTemplateOptions{
				TemplateID: &profileTemplateId,
				Version:    core.StringPtr(strconv.FormatInt(profileTemplateVersion, 10)),
			}
			cResponse, cErr := iamIdentityService.CommitProfileTemplate(commitOptions)
			Expect(cResponse.StatusCode).To(Equal(204))
			Expect(cErr).To(BeNil())

			waitUntilTrustedProfileAssignmentFinishedIT(iamIdentityService, &profileTemplateAssignmentId, &profileTemplateAssignmentEtag)

			updateOptions := &iamidentityv1.UpdateTrustedProfileAssignmentOptions{
				AssignmentID:    &profileTemplateAssignmentId,
				TemplateVersion: &profileTemplateVersion,
				IfMatch:         &profileTemplateAssignmentEtag,
			}
			updateResponse, uResponse, uErr := iamIdentityService.UpdateTrustedProfileAssignment(updateOptions)
			Expect(uResponse.StatusCode).To(Equal(202))
			Expect(uErr).To(BeNil())
			Expect(updateResponse).ToNot(BeNil())
			// Grab the Etag and id for use by other test methods.
			Expect(uResponse.GetHeaders()).ToNot(BeNil())
			profileTemplateAssignmentEtag = uResponse.GetHeaders().Get("Etag")
			Expect(profileTemplateAssignmentEtag).ToNot(BeEmpty())
		})
	})

	Describe(`DeleteProfileTemplateAssignmentIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProfileTemplateAssignmentIT`, func() {
			waitUntilTrustedProfileAssignmentFinishedIT(iamIdentityService, &profileTemplateAssignmentId, &profileTemplateAssignmentEtag)

			deleteOptions := &iamidentityv1.DeleteTrustedProfileAssignmentOptions{
				AssignmentID: &profileTemplateAssignmentId,
			}
			excResponse, response, err := iamIdentityService.DeleteTrustedProfileAssignment(deleteOptions)
			Expect(response.StatusCode).To(Equal(202))
			Expect(err).To(BeNil())
			Expect(excResponse).To(BeNil())

			waitUntilTrustedProfileAssignmentFinishedIT(iamIdentityService, &profileTemplateAssignmentId, &profileTemplateAssignmentEtag)
		})
	})

	Describe(`DeleteProfileTemplateVersionIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProfileTemplateVersionIT`, func() {
			deleteOptions := &iamidentityv1.DeleteProfileTemplateVersionOptions{
				TemplateID: &profileTemplateId,
				Version:    core.StringPtr("1"),
			}
			response, err := iamIdentityService.DeleteProfileTemplateVersion(deleteOptions)
			Expect(response.StatusCode).To(Equal(204))
			Expect(err).To(BeNil())
		})
	})

	Describe(`DeleteProfileTemplateIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteProfileTemplateIT`, func() {
			waitUntilTrustedProfileAssignmentFinishedIT(iamIdentityService, &profileTemplateAssignmentId, &profileTemplateAssignmentEtag)

			deleteOptions := &iamidentityv1.DeleteAllVersionsOfProfileTemplateOptions{
				TemplateID: &profileTemplateId,
			}
			response, err := iamIdentityService.DeleteAllVersionsOfProfileTemplate(deleteOptions)
			Expect(response.StatusCode).To(Equal(204))
			Expect(err).To(BeNil())
		})
	})

	Describe(`CreateAccountSettingsTemplateIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAccountSettingsTemplateIT`, func() {

			settings := &iamidentityv1.AccountSettingsComponent{
				Mfa:                                  core.StringPtr("LEVEL1"),
				SystemAccessTokenExpirationInSeconds: core.StringPtr("3000"),
			}
			createOptions := &iamidentityv1.CreateAccountSettingsTemplateOptions{
				Name:            &accountSettingsTemplateName,
				Description:     core.StringPtr("GoSDK test Account Settings Template"),
				AccountID:       &enterpriseAccountID,
				AccountSettings: settings,
			}

			createResponse, response, err := iamIdentityService.CreateAccountSettingsTemplate(createOptions)
			Expect(response.StatusCode).To(Equal(201))
			Expect(err).To(BeNil())
			Expect(createResponse).ToNot(BeNil())

			// Grab the ID and Etag value from the response for use in the update operation.
			accountSettingsTemplateId = *createResponse.ID
			Expect(accountSettingsTemplateId).ToNot(BeNil())
			accountSettingsTemplateVersion = *createResponse.Version
			Expect(response.GetHeaders()).ToNot(BeNil())
			accountSettingsTemplateEtag = response.GetHeaders().Get("Etag")
			Expect(accountSettingsTemplateEtag).ToNot(BeEmpty())
		})
	})

	Describe(`GetAccountSettingsTemplateIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountSettingsTemplateIT`, func() {
			getOptions := &iamidentityv1.GetAccountSettingsTemplateVersionOptions{
				TemplateID: &accountSettingsTemplateId,
				Version:    core.StringPtr(strconv.FormatInt(accountSettingsTemplateVersion, 10)),
			}
			getResponse, response, err := iamIdentityService.GetAccountSettingsTemplateVersion(getOptions)
			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(getResponse).ToNot(BeNil())
			// Grab the Etag value from the response for use in the update operation.
			Expect(response.GetHeaders()).ToNot(BeNil())
			accountSettingsTemplateEtag = response.GetHeaders().Get("Etag")
			Expect(accountSettingsTemplateEtag).ToNot(BeEmpty())
		})
	})

	Describe(`ListAccountSettingsTemplatesIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAccountSettingsTemplatesIT`, func() {
			listOptions := &iamidentityv1.ListAccountSettingsTemplatesOptions{
				AccountID: &enterpriseAccountID,
			}
			listResponse, response, err := iamIdentityService.ListAccountSettingsTemplates(listOptions)
			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(listResponse).ToNot(BeNil())
		})
	})

	Describe(`UpdateAccountSettingsTemplateIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAccountSettingsTemplateIT`, func() {
			settings := &iamidentityv1.AccountSettingsComponent{
				Mfa:                                  core.StringPtr("LEVEL1"),
				SystemAccessTokenExpirationInSeconds: core.StringPtr("3000"),
			}
			updateOptions := &iamidentityv1.UpdateAccountSettingsTemplateVersionOptions{
				AccountID:       &enterpriseAccountID,
				TemplateID:      &accountSettingsTemplateId,
				Version:         core.StringPtr(strconv.FormatInt(accountSettingsTemplateVersion, 10)),
				IfMatch:         &accountSettingsTemplateEtag,
				Name:            &accountSettingsTemplateName,
				Description:     core.StringPtr("GoSDK test Account Settings Template - updated"),
				AccountSettings: settings,
			}
			updateResponse, response, err := iamIdentityService.UpdateAccountSettingsTemplateVersion(updateOptions)
			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(updateResponse).ToNot(BeNil())
			// Grab the Etag value from the response for use in the update operation.
			Expect(response.GetHeaders()).ToNot(BeNil())
			accountSettingsTemplateEtag = response.GetHeaders().Get("Etag")
			Expect(accountSettingsTemplateEtag).ToNot(BeEmpty())
		})
	})

	Describe(`AssignAccountSettingsTemplateIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AssignAccountSettingsTemplateIT`, func() {
			commitOptions := &iamidentityv1.CommitAccountSettingsTemplateOptions{
				TemplateID: &accountSettingsTemplateId,
				Version:    core.StringPtr(strconv.FormatInt(accountSettingsTemplateVersion, 10)),
			}
			cResponse, cErr := iamIdentityService.CommitAccountSettingsTemplate(commitOptions)
			Expect(cResponse.StatusCode).To(Equal(204))
			Expect(cErr).To(BeNil())

			assignOptions := &iamidentityv1.CreateAccountSettingsAssignmentOptions{
				TemplateID:      &accountSettingsTemplateId,
				TemplateVersion: &accountSettingsTemplateVersion,
				TargetType:      core.StringPtr("Account"),
				Target:          &enterpriseSubAccountID,
			}
			assignResponse, aResponse, aErr := iamIdentityService.CreateAccountSettingsAssignment(assignOptions)
			Expect(aResponse.StatusCode).To(Equal(202))
			Expect(aErr).To(BeNil())
			Expect(assignResponse).ToNot(BeNil())
			// Grab the Etag and id for use by other test methods.
			Expect(aResponse.GetHeaders()).ToNot(BeNil())
			accountSettingsTemplateAssignmentEtag = aResponse.GetHeaders().Get("Etag")
			Expect(accountSettingsTemplateAssignmentEtag).ToNot(BeEmpty())
			accountSettingsTemplateAssignmentId = *assignResponse.ID
			Expect(accountSettingsTemplateAssignmentId).ToNot(BeNil())
		})
	})

	Describe(`ListAccountSettingsTemplateAssignmentsIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAccountSettingsTemplateAssignmentsIT`, func() {
			listOptions := &iamidentityv1.ListAccountSettingsAssignmentsOptions{
				AccountID:  &enterpriseAccountID,
				TemplateID: &accountSettingsTemplateId,
			}
			listResponse, response, err := iamIdentityService.ListAccountSettingsAssignments(listOptions)
			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(listResponse).ToNot(BeNil())
			Expect(listResponse.Assignments).ToNot(BeNil())
			numAssignments := len(listResponse.Assignments)
			Expect(numAssignments).ToNot(BeZero())
		})
	})

	Describe(`CreateNewAccountSettingsTemplateVersionIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNewAccountSettingsTemplateVersionIT`, func() {
			settings := &iamidentityv1.AccountSettingsComponent{
				Mfa:                                  core.StringPtr("LEVEL1"),
				SystemAccessTokenExpirationInSeconds: core.StringPtr("2600"),
				RestrictCreatePlatformApikey:         core.StringPtr("RESTRICTED"),
				RestrictCreateServiceID:              core.StringPtr("RESTRICTED"),
			}
			createOptions := &iamidentityv1.CreateAccountSettingsTemplateVersionOptions{
				Name:            &accountSettingsTemplateName,
				Description:     core.StringPtr("GoSDK test Account Settings Template - new version"),
				AccountID:       &enterpriseAccountID,
				TemplateID:      &accountSettingsTemplateId,
				AccountSettings: settings,
			}

			createResponse, response, err := iamIdentityService.CreateAccountSettingsTemplateVersion(createOptions)
			Expect(response.StatusCode).To(Equal(201))
			Expect(err).To(BeNil())
			Expect(createResponse).ToNot(BeNil())
			// save the new version to be used in subsequent calls
			accountSettingsTemplateVersion = *createResponse.Version
			Expect(response.GetHeaders()).ToNot(BeNil())
		})
	})

	Describe(`GetLatestAccountSettingsTemplateVersionIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLatestAccountSettingsTemplateVersionIT`, func() {
			getOptions := &iamidentityv1.GetLatestAccountSettingsTemplateVersionOptions{
				TemplateID: &accountSettingsTemplateId,
			}
			getResponse, response, err := iamIdentityService.GetLatestAccountSettingsTemplateVersion(getOptions)
			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(getResponse).ToNot(BeNil())
		})
	})

	Describe(`ListAccountSettingsTemplateVersionsIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAccountSettingsTemplateVersionsIT`, func() {
			listOptions := &iamidentityv1.ListVersionsOfAccountSettingsTemplateOptions{
				TemplateID: &accountSettingsTemplateId,
			}
			listResponse, response, err := iamIdentityService.ListVersionsOfAccountSettingsTemplate(listOptions)
			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(listResponse).ToNot(BeNil())
			Expect(listResponse.AccountSettingsTemplates).ToNot(BeNil())
			numTemplates := len(listResponse.AccountSettingsTemplates)
			Expect(numTemplates).ToNot(BeZero())
		})
	})

	Describe(`UpdateAccountSettingsTemplateAssignmentIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAccountSettingsTemplateAssignmentIT`, func() {
			commitOptions := &iamidentityv1.CommitAccountSettingsTemplateOptions{
				TemplateID: &accountSettingsTemplateId,
				Version:    core.StringPtr(strconv.FormatInt(accountSettingsTemplateVersion, 10)),
			}
			cResponse, cErr := iamIdentityService.CommitAccountSettingsTemplate(commitOptions)
			Expect(cResponse.StatusCode).To(Equal(204))
			Expect(cErr).To(BeNil())

			waitUntilAccountSettingsAssignmentFinishedIT(iamIdentityService, &accountSettingsTemplateAssignmentId, &accountSettingsTemplateAssignmentEtag)

			updateOptions := &iamidentityv1.UpdateAccountSettingsAssignmentOptions{
				AssignmentID:    &accountSettingsTemplateAssignmentId,
				TemplateVersion: &accountSettingsTemplateVersion,
				IfMatch:         &accountSettingsTemplateAssignmentEtag,
			}
			updateResponse, uResponse, uErr := iamIdentityService.UpdateAccountSettingsAssignment(updateOptions)
			Expect(uResponse.StatusCode).To(Equal(202))
			Expect(uErr).To(BeNil())
			Expect(updateResponse).ToNot(BeNil())
			// Grab the Etag and id for use by other test methods.
			Expect(uResponse.GetHeaders()).ToNot(BeNil())
			accountSettingsTemplateAssignmentEtag = uResponse.GetHeaders().Get("Etag")
			Expect(accountSettingsTemplateAssignmentEtag).ToNot(BeEmpty())
		})
	})

	Describe(`DeleteAccountSettingsTemplateAssignmentIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAccountSettingsTemplateAssignmentIT`, func() {
			waitUntilAccountSettingsAssignmentFinishedIT(iamIdentityService, &accountSettingsTemplateAssignmentId, &accountSettingsTemplateAssignmentEtag)

			deleteOptions := &iamidentityv1.DeleteAccountSettingsAssignmentOptions{
				AssignmentID: &accountSettingsTemplateAssignmentId,
			}
			excResponse, response, err := iamIdentityService.DeleteAccountSettingsAssignment(deleteOptions)
			Expect(response.StatusCode).To(Equal(202))
			Expect(err).To(BeNil())
			Expect(excResponse).To(BeNil())

			waitUntilAccountSettingsAssignmentFinishedIT(iamIdentityService, &accountSettingsTemplateAssignmentId, &accountSettingsTemplateAssignmentEtag)
		})
	})

	Describe(`DeleteAccountSettingsTemplateVersionIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAccountSettingsTemplateVersionIT`, func() {
			deleteOptions := &iamidentityv1.DeleteAccountSettingsTemplateVersionOptions{
				TemplateID: &accountSettingsTemplateId,
				Version:    core.StringPtr("1"),
			}
			response, err := iamIdentityService.DeleteAccountSettingsTemplateVersion(deleteOptions)
			Expect(response.StatusCode).To(Equal(204))
			Expect(err).To(BeNil())
		})
	})

	Describe(`DeleteAccountSettingsTemplateIT`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAccountSettingsTemplateIT`, func() {
			waitUntilAccountSettingsAssignmentFinishedIT(iamIdentityService, &accountSettingsTemplateAssignmentId, &accountSettingsTemplateAssignmentEtag)

			deleteOptions := &iamidentityv1.DeleteAllVersionsOfAccountSettingsTemplateOptions{
				TemplateID: &accountSettingsTemplateId,
			}
			response, err := iamIdentityService.DeleteAllVersionsOfAccountSettingsTemplate(deleteOptions)
			Expect(response.StatusCode).To(Equal(204))
			Expect(err).To(BeNil())
		})
	})

	Describe(`UpdatePreferenceOnScopeAccount`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePreferenceOnScopeAccount(updatePreferenceOnScopeAccountOptions *UpdatePreferenceOnScopeAccountOptions)`, func() {
			Expect(accountID).ToNot(BeEmpty())
			Expect(iamIDForPreferences).ToNot(BeEmpty())

			updatePreferenceOnScopeAccountOptions := &iamidentityv1.UpdatePreferenceOnScopeAccountOptions{
				AccountID:    &accountID,
				IamID:        &iamIDForPreferences,
				Service:      &service,
				PreferenceID: &preferenceID1,
				ValueString:  &valueString,
			}

			preference, response, err := iamIdentityService.UpdatePreferenceOnScopeAccount(updatePreferenceOnScopeAccountOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(preference).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "UpdatePreferenceOnScopeAccount response:\n%s\n", common.ToJSON(preference))
		})
	})

	Describe(`GetPreferencesOnScopeAccount`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPreferencesOnScopeAccount(getPreferencesOnScopeAccountOptions *GetPreferencesOnScopeAccountOptions)`, func() {

			getPreferencesOnScopeAccountOptions := &iamidentityv1.GetPreferencesOnScopeAccountOptions{
				AccountID:    &accountID,
				IamID:        &iamIDForPreferences,
				Service:      &service,
				PreferenceID: &preferenceID1,
			}

			preference, response, err := iamIdentityService.GetPreferencesOnScopeAccount(getPreferencesOnScopeAccountOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(preference).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "getPreferencesOnScopeAccount response:\n%s\n", common.ToJSON(preference))
		})
	})

	Describe(`GetAllPreferencesOnScopeAccount`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAllPreferencesOnScopeAccount(getAllPreferencesOnScopeAccountOptions *GetAllPreferencesOnScopeAccount)`, func() {

			getAllPreferencesOnScopeAccountOptions := &iamidentityv1.GetAllPreferencesOnScopeAccountOptions{
				AccountID: &accountID,
				IamID:     &iamIDForPreferences,
			}

			preference, response, err := iamIdentityService.GetAllPreferencesOnScopeAccount(getAllPreferencesOnScopeAccountOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(preference).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "getAllPreferencesOnScopeAccount response:\n%s\n", common.ToJSON(preference))
		})
	})

	Describe(`DeletePreferencesOnScopeAccount`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePreferencesOnScopeAccount(deletePreferencesOnScopeAccountOptions *DeletePreferencesOnScopeAccount)`, func() {

			deletePreferencesOnScopeAccountOptions := &iamidentityv1.DeletePreferencesOnScopeAccountOptions{
				AccountID:    &accountID,
				IamID:        &iamIDForPreferences,
				Service:      &service,
				PreferenceID: &preferenceID1,
			}

			response, err := iamIdentityService.DeletePreferencesOnScopeAccount(deletePreferencesOnScopeAccountOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	AfterSuite(func() {
		fmt.Println("\nBeginning teardown.")
		cleanupResources(iamIdentityService, accountID, iamID, apikeyName, serviceIDName, profileName1, profileName2, enterpriseAccountID, profileTemplateName, &profileTemplateAssignmentEtag, accountSettingsTemplateName, &accountSettingsTemplateAssignmentEtag)
		fmt.Println("Finished teardown.")
	})
})

func getAPIkey(service *iamidentityv1.IamIdentityV1, apikeyID string) *iamidentityv1.APIKey {
	getAPIKeyOptions := &iamidentityv1.GetAPIKeyOptions{
		ID: &apikeyID,
	}
	apiKey, _, _ := service.GetAPIKey(getAPIKeyOptions)
	return apiKey
}

func getServiceID(iamIdentityService *iamidentityv1.IamIdentityV1, serviceID string) *iamidentityv1.ServiceID {
	getServiceIDOptions := &iamidentityv1.GetServiceIDOptions{
		ID: &serviceID,
	}
	result, _, _ := iamIdentityService.GetServiceID(getServiceIDOptions)
	return result
}

func getProfile(service *iamidentityv1.IamIdentityV1, profileID string) *iamidentityv1.TrustedProfile {
	getProfileOptions := &iamidentityv1.GetProfileOptions{
		ProfileID: &profileID,
	}
	profile, _, _ := service.GetProfile(getProfileOptions)
	return profile
}

func getClaimRule(service *iamidentityv1.IamIdentityV1, profileID string, claimRuleID string) *iamidentityv1.ProfileClaimRule {
	getClaimRuleOptions := &iamidentityv1.GetClaimRuleOptions{
		ProfileID: &profileID,
		RuleID:    &claimRuleID,
	}
	claimRule, _, _ := service.GetClaimRule(getClaimRuleOptions)
	return claimRule
}

func getLink(service *iamidentityv1.IamIdentityV1, profileID string, linkID string) *iamidentityv1.ProfileLink {
	getLinkOptions := &iamidentityv1.GetLinkOptions{
		ProfileID: &profileID,
		LinkID:    &linkID,
	}
	link, _, _ := service.GetLink(getLinkOptions)
	return link
}

func isFinishedIT(status *string) bool {
	var finished = false
	if strings.EqualFold(*status, "succeeded") || strings.EqualFold(*status, "failed") {
		finished = true
	}
	return finished
}

func waitUntilTrustedProfileAssignmentFinishedIT(service *iamidentityv1.IamIdentityV1, assignmentId *string, profileTemplateAssignmentEtag *string) {
	getAssignmentOptions := &iamidentityv1.GetTrustedProfileAssignmentOptions{
		AssignmentID: assignmentId,
	}

	var finished = true
	for i := 0; i < 30; i++ {
		assignment, response, err := service.GetTrustedProfileAssignment(getAssignmentOptions)
		if response.StatusCode == 404 {
			Expect(err).ToNot(BeNil())
			finished = true // assignment removed
			break
		} else {
			finished = isFinishedIT(assignment.Status)
			if finished {
				// Grab the Etag value from the response for use in the update operation.
				Expect(response.GetHeaders()).ToNot(BeNil())
				*profileTemplateAssignmentEtag = response.GetHeaders().Get("Etag")
				Expect(*profileTemplateAssignmentEtag).ToNot(BeEmpty())
				break
			}
		}
		time.Sleep(10 * time.Second)
	}
	Expect(finished).To(BeTrue())
}

func waitUntilAccountSettingsAssignmentFinishedIT(service *iamidentityv1.IamIdentityV1, assignmentId *string, accountSettingsTemplateAssignmentEtag *string) {
	getAssignmentOptions := &iamidentityv1.GetAccountSettingsAssignmentOptions{
		AssignmentID: assignmentId,
	}

	var finished = true
	for i := 0; i < 30; i++ {
		assignment, response, err := service.GetAccountSettingsAssignment(getAssignmentOptions)
		if response.StatusCode == 404 {
			Expect(err).ToNot(BeNil())
			finished = true // assignment removed
			break
		} else {
			finished = isFinishedIT(assignment.Status)
			if finished {
				// Grab the Etag value from the response for use in the update operation.
				Expect(response.GetHeaders()).ToNot(BeNil())
				*accountSettingsTemplateAssignmentEtag = response.GetHeaders().Get("Etag")
				Expect(*accountSettingsTemplateAssignmentEtag).ToNot(BeEmpty())
				break
			}
		}
		time.Sleep(10 * time.Second)
	}
	Expect(finished).To(BeTrue())
}

func getPageTokenFromURL(sptr *string) *string {
	if sptr == nil {
		return nil
	}

	s := *sptr
	if s == "" {
		return nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return nil
	}

	if u.RawQuery == "" {
		return nil
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil
	}

	token := q.Get("pagetoken")
	if token == "" {
		return nil
	}
	return &token
}

func cleanupResources(service *iamidentityv1.IamIdentityV1, accountID string, iamID string, apikeyName string, serviceIDName string, profileName1 string, profileName2 string, enterpriseAccountID string, profileTemplateName string, profileTemplateAssignmentEtag *string, accountSettingsTemplateName string, accountSettingsTemplateAssignmentEtag *string) {
	if service == nil {
		panic("'service' cannot be nil!")
	}

	listAPIKeysOptions := &iamidentityv1.ListAPIKeysOptions{
		AccountID: &accountID,
		IamID:     &iamID,
		Pagesize:  core.Int64Ptr(int64(100)),
	}

	apiKeyList, response, err := service.ListAPIKeys(listAPIKeysOptions)
	Expect(err).To(BeNil())
	Expect(response.StatusCode).To(Equal(200))

	numAPIKeys := len(apiKeyList.Apikeys)
	fmt.Fprintf(GinkgoWriter, ">>> Cleanup found %d apikeys.\n", numAPIKeys)

	if numAPIKeys > 0 {
		for _, element := range apiKeyList.Apikeys {
			if *element.Name == apikeyName {
				fmt.Fprintf(GinkgoWriter, ">>> Deleting apikey: %s\n", *element.ID)
				deleteAPIKeyOptions := &iamidentityv1.DeleteAPIKeyOptions{
					ID: element.ID,
				}
				response, err := service.DeleteAPIKey(deleteAPIKeyOptions)
				Expect(response).ToNot(BeNil())
				Expect(err).To(BeNil())
			}
		}
	}

	listServiceIdsOptions := &iamidentityv1.ListServiceIdsOptions{
		AccountID: &accountID,
		Name:      &serviceIDName,
		Pagesize:  core.Int64Ptr(int64(100)),
	}

	serviceIDList, _, _ := service.ListServiceIds(listServiceIdsOptions)

	numServiceIds := len(serviceIDList.Serviceids)
	fmt.Fprintf(GinkgoWriter, ">>> Cleanup found %d serviceIDs.\n", numServiceIds)

	if numServiceIds > 0 {
		for _, element := range serviceIDList.Serviceids {
			fmt.Fprintf(GinkgoWriter, ">>> Deleting serviceId: %s\n", *element.ID)
			deleteServiceIDOptions := &iamidentityv1.DeleteServiceIDOptions{
				ID: element.ID,
			}
			response, err := service.DeleteServiceID(deleteServiceIDOptions)
			Expect(response).ToNot(BeNil())
			Expect(err).To(BeNil())
		}
	}

	listProfilesOptions := &iamidentityv1.ListProfilesOptions{
		AccountID: &accountID,
	}

	profileList, response, err := service.ListProfiles(listProfilesOptions)
	Expect(err).To(BeNil())
	Expect(response.StatusCode).To(Equal(200))

	numProfiles := len(profileList.Profiles)
	fmt.Fprintf(GinkgoWriter, ">>> Cleanup found %d profiles.\n", numProfiles)

	if numProfiles > 0 {
		for _, element := range profileList.Profiles {
			if *element.Name == profileName1 || *element.Name == profileName2 {
				fmt.Fprintf(GinkgoWriter, ">>> Deleting profile: %s\n", *element.ID)
				deleteProfileOptions := &iamidentityv1.DeleteProfileOptions{
					ProfileID: element.ID,
				}
				response, err := service.DeleteProfile(deleteProfileOptions)
				Expect(response).ToNot(BeNil())
				Expect(err).To(BeNil())
			}
		}
	}

	listProfileTemplateOptions := &iamidentityv1.ListProfileTemplatesOptions{
		AccountID: &enterpriseAccountID,
	}
	profileTemplateList, response, err := service.ListProfileTemplates(listProfileTemplateOptions)
	Expect(err).To(BeNil())
	Expect(response.StatusCode).To(Equal(200))

	numProfileTemplates := len(profileTemplateList.ProfileTemplates)
	fmt.Fprintf(GinkgoWriter, ">>> Cleanup found %d profile templates.\n", numProfileTemplates)

	if numProfileTemplates > 0 {
		for _, element := range profileTemplateList.ProfileTemplates {
			if *element.Name == profileTemplateName {
				fmt.Fprintf(GinkgoWriter, ">>> Deleting profile template: %s\n", *element.ID)

				listAssignmentsOptions := &iamidentityv1.ListTrustedProfileAssignmentsOptions{
					AccountID:  &enterpriseAccountID,
					TemplateID: element.ID,
				}
				assignmentsList, alResponse, alErr := service.ListTrustedProfileAssignments(listAssignmentsOptions)
				Expect(alErr).To(BeNil())
				Expect(alResponse.StatusCode).To(Equal(200))
				Expect(assignmentsList).ToNot(BeNil())
				numAssignments := len(assignmentsList.Assignments)
				if numAssignments > 0 {
					for _, assignment := range assignmentsList.Assignments {
						if !isFinishedIT(assignment.Status) {
							waitUntilTrustedProfileAssignmentFinishedIT(service, assignment.ID, profileTemplateAssignmentEtag)
						}
						deleteAssignmentOptions := &iamidentityv1.DeleteTrustedProfileAssignmentOptions{
							AssignmentID: assignment.ID,
						}
						exceptionRsp, daResponse, daErr := service.DeleteTrustedProfileAssignment(deleteAssignmentOptions)
						Expect(exceptionRsp).To(BeNil())
						Expect(daResponse).ToNot(BeNil())
						Expect(daErr).To(BeNil())
						waitUntilTrustedProfileAssignmentFinishedIT(service, assignment.ID, profileTemplateAssignmentEtag)
					}
				}

				deleteProfileTemplateOptions := &iamidentityv1.DeleteAllVersionsOfProfileTemplateOptions{
					TemplateID: element.ID,
				}
				response, err := service.DeleteAllVersionsOfProfileTemplate(deleteProfileTemplateOptions)
				Expect(response).ToNot(BeNil())
				Expect(err).To(BeNil())
			}
		}
	}

	listAccountSettingsTemplateOptions := &iamidentityv1.ListAccountSettingsTemplatesOptions{
		AccountID: &enterpriseAccountID,
	}
	AccountSettingsTemplateList, response, err := service.ListAccountSettingsTemplates(listAccountSettingsTemplateOptions)
	Expect(err).To(BeNil())
	Expect(response.StatusCode).To(Equal(200))

	numAccountSettingsTemplates := len(AccountSettingsTemplateList.AccountSettingsTemplates)
	fmt.Fprintf(GinkgoWriter, ">>> Cleanup found %d account settings templates.\n", numAccountSettingsTemplates)

	if numAccountSettingsTemplates > 0 {
		for _, element := range AccountSettingsTemplateList.AccountSettingsTemplates {
			if *element.Name == accountSettingsTemplateName {
				fmt.Fprintf(GinkgoWriter, ">>> Deleting account settings template: %s\n", *element.ID)

				listAssignmentsOptions := &iamidentityv1.ListAccountSettingsAssignmentsOptions{
					AccountID:  &enterpriseAccountID,
					TemplateID: element.ID,
				}
				assignmentsList, alResponse, alErr := service.ListAccountSettingsAssignments(listAssignmentsOptions)
				Expect(alErr).To(BeNil())
				Expect(alResponse.StatusCode).To(Equal(200))
				Expect(assignmentsList).ToNot(BeNil())
				numAssignments := len(assignmentsList.Assignments)
				if numAssignments > 0 {
					for _, assignment := range assignmentsList.Assignments {
						if !isFinishedIT(assignment.Status) {
							waitUntilAccountSettingsAssignmentFinishedIT(service, assignment.ID, accountSettingsTemplateAssignmentEtag)
						}
						deleteAssignmentOptions := &iamidentityv1.DeleteAccountSettingsAssignmentOptions{
							AssignmentID: assignment.ID,
						}
						exceptionRsp, daResponse, daErr := service.DeleteAccountSettingsAssignment(deleteAssignmentOptions)
						Expect(exceptionRsp).To(BeNil())
						Expect(daResponse).ToNot(BeNil())
						Expect(daErr).To(BeNil())
						waitUntilAccountSettingsAssignmentFinishedIT(service, assignment.ID, accountSettingsTemplateAssignmentEtag)
					}
				}

				deleteAccountSettingsTemplateOptions := &iamidentityv1.DeleteAllVersionsOfAccountSettingsTemplateOptions{
					TemplateID: element.ID,
				}
				response, err := service.DeleteAllVersionsOfAccountSettingsTemplate(deleteAccountSettingsTemplateOptions)
				Expect(response).ToNot(BeNil())
				Expect(err).To(BeNil())
			}
		}
	}
}
