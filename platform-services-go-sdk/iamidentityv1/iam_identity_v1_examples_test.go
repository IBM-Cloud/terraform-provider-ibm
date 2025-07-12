//go:build examples

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
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the IAM Identity service.
//
// The following configuration properties are assumed to be defined:
//
// IAM_IDENTITY_URL=<service url>
// IAM_IDENTITY_AUTHTYPE=iam
// IAM_IDENTITY_AUTH_URL=<IAM Token Service url>
// IAM_IDENTITY_APIKEY=<IAM APIKEY for the User>
// IAM_IDENTITY_ACCOUNT_ID=<AccountID which is unique to the User>
// IAM_IDENTITY_IAM_ID=<IAM ID which is unique to the User account>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//

var _ = Describe(`IamIdentityV1 Examples Tests`, func() {
	const externalConfigFile = "../iam_identity.env"

	var (
		iamIdentityService *iamidentityv1.IamIdentityV1
		config             map[string]string
		configLoaded       bool = false

		serviceURL string

		apikeyName    string = "Example-ApiKey"
		serviceIDName string = "Example-ServiceId"
		profileName   string = "Example-Profile"
		accountID     string
		iamID         string
		iamIDMember   string
		iamAPIKey     string

		apikeyID   string
		apikeyEtag string

		svcID     string
		svcIDEtag string

		profileId     string
		profileEtag   string
		claimRuleId   string
		claimRuleEtag string
		claimRuleType string = "Profile-SAML"
		realmName     string = "https://sdk.test.realm/1234"
		linkId        string

		accountSettingEtag string

		enterpriseAccountID                   string
		enterpriseSubAccountID                string
		profileTemplateName                   string = "Example-Profile-Template"
		profileTemplateProfileName            string = "Example-Profile-From-Template"
		profileTemplateId                     string
		profileTemplateVersion                int64
		profileTemplateEtag                   string
		profileTemplateAssignmentId           string
		profileTemplateAssignmentEtag         string
		accountSettingsTemplateName           string = "Example-AccountSettings-Template"
		accountSettingsTemplateId             string
		accountSettingsTemplateVersion        int64
		accountSettingsTemplateEtag           string
		accountSettingsTemplateAssignmentId   string
		accountSettingsTemplateAssignmentEtag string

		service             string = "console"
		valueString         string = "/billing"
		preferenceID1       string = "landing_page"
		iamIDForPreferences string
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
			config, err = core.GetServiceProperties(iamidentityv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0

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

			enterpriseAccountID = config["ENTERPRISE_ACCOUNT_ID"]
			Expect(enterpriseAccountID).ToNot(BeEmpty())

			enterpriseSubAccountID = config["ENTERPRISE_SUBACCOUNT_ID"]
			Expect(enterpriseSubAccountID).ToNot(BeEmpty())

			iamIDForPreferences = config["IAM_ID_FOR_PREFERENCES"]
			Expect(enterpriseSubAccountID).ToNot(BeEmpty())

			fmt.Printf("Service URL: %s\n", serviceURL)
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			iamIdentityServiceOptions := &iamidentityv1.IamIdentityV1Options{}

			iamIdentityService, err = iamidentityv1.NewIamIdentityV1UsingExternalConfig(iamIdentityServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(err).To(BeNil())
			Expect(iamIdentityService).ToNot(BeNil())
			Expect(iamIdentityService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`IamIdentityV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAPIKey request example`, func() {
			fmt.Println("\nCreateAPIKey() result:")
			// begin-create_api_key

			createAPIKeyOptions := iamIdentityService.NewCreateAPIKeyOptions(apikeyName, iamID)
			createAPIKeyOptions.SetDescription("Example ApiKey")

			apiKey, response, err := iamIdentityService.CreateAPIKey(createAPIKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiKey, "", "  ")
			fmt.Println(string(b))
			apikeyID = *apiKey.ID

			// end-create_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(apiKey).ToNot(BeNil())
			Expect(apikeyID).ToNot(BeNil())
		})
		It(`ListAPIKeys request example`, func() {
			fmt.Println("\nListAPIKeys() result:")
			// begin-list_api_keys

			listAPIKeysOptions := iamIdentityService.NewListAPIKeysOptions()
			listAPIKeysOptions.SetAccountID(accountID)
			listAPIKeysOptions.SetIamID(iamID)
			listAPIKeysOptions.SetIncludeHistory(true)

			apiKeyList, response, err := iamIdentityService.ListAPIKeys(listAPIKeysOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiKeyList, "", "  ")
			fmt.Println(string(b))

			// end-list_api_keys

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKeyList).ToNot(BeNil())
		})
		It(`GetAPIKeysDetails request example`, func() {
			fmt.Println("\nGetAPIKeysDetails() result:")
			// begin-get_api_keys_details

			getAPIKeysDetailsOptions := iamIdentityService.NewGetAPIKeysDetailsOptions()
			getAPIKeysDetailsOptions.SetIamAPIKey(iamAPIKey)
			getAPIKeysDetailsOptions.SetIncludeHistory(false)

			apiKey, response, err := iamIdentityService.GetAPIKeysDetails(getAPIKeysDetailsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiKey, "", "  ")
			fmt.Println(string(b))

			// end-get_api_keys_details

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
		})
		It(`GetAPIKey request example`, func() {
			fmt.Println("\nGetAPIKey() result:")
			// begin-get_api_key

			getAPIKeyOptions := iamIdentityService.NewGetAPIKeyOptions(apikeyID)

			getAPIKeyOptions.SetIncludeHistory(false)
			getAPIKeyOptions.SetIncludeActivity(false)

			apiKey, response, err := iamIdentityService.GetAPIKey(getAPIKeyOptions)
			if err != nil {
				panic(err)
			}
			apikeyEtag = response.GetHeaders().Get("Etag")
			b, _ := json.MarshalIndent(apiKey, "", "  ")
			fmt.Println(string(b))

			// end-get_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
			Expect(apikeyEtag).ToNot(BeEmpty())
		})
		It(`UpdateAPIKey request example`, func() {
			fmt.Println("\nUpdateAPIKey() result:")
			// begin-update_api_key

			updateAPIKeyOptions := iamIdentityService.NewUpdateAPIKeyOptions(apikeyID, apikeyEtag)
			updateAPIKeyOptions.SetDescription("This is an updated description")

			apiKey, response, err := iamIdentityService.UpdateAPIKey(updateAPIKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(apiKey, "", "  ")
			fmt.Println(string(b))

			// end-update_api_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(apiKey).ToNot(BeNil())
		})
		It(`LockAPIKey request example`, func() {
			// begin-lock_api_key

			lockAPIKeyOptions := iamIdentityService.NewLockAPIKeyOptions(apikeyID)

			response, err := iamIdentityService.LockAPIKey(lockAPIKeyOptions)
			if err != nil {
				panic(err)
			}

			// end-lock_api_key
			fmt.Printf("\nLockAPIKey() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`UnlockAPIKey request example`, func() {
			// begin-unlock_api_key

			unlockAPIKeyOptions := iamIdentityService.NewUnlockAPIKeyOptions(apikeyID)

			response, err := iamIdentityService.UnlockAPIKey(unlockAPIKeyOptions)
			if err != nil {
				panic(err)
			}

			// end-unlock_api_key
			fmt.Printf("\nUnlockAPIKey() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DisableAPIKey request example`, func() {
			// begin-disable_api_key

			disableAPIKeyOptions := iamIdentityService.NewDisableAPIKeyOptions(apikeyID)

			response, err := iamIdentityService.DisableAPIKey(disableAPIKeyOptions)
			if err != nil {
				panic(err)
			}

			// end-disable_api_key
			fmt.Printf("\nDisableAPIKey() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`EnableAPIKey request example`, func() {
			// begin-enable_api_key

			enableAPIKeyOptions := iamIdentityService.NewEnableAPIKeyOptions(apikeyID)

			response, err := iamIdentityService.EnableAPIKey(enableAPIKeyOptions)
			if err != nil {
				panic(err)
			}

			// end-enable_api_key
			fmt.Printf("\nEnableAPIKey() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteAPIKey request example`, func() {
			// begin-delete_api_key

			deleteAPIKeyOptions := iamIdentityService.NewDeleteAPIKeyOptions(apikeyID)

			response, err := iamIdentityService.DeleteAPIKey(deleteAPIKeyOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_api_key
			fmt.Printf("\nDeleteAPIKey() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CreateServiceID request example`, func() {
			fmt.Println("\nCreateServiceID() result:")
			// begin-create_service_id

			createServiceIDOptions := iamIdentityService.NewCreateServiceIDOptions(accountID, serviceIDName)
			createServiceIDOptions.SetDescription("Example ServiceId")

			serviceID, response, err := iamIdentityService.CreateServiceID(createServiceIDOptions)
			if err != nil {
				panic(err)
			}
			svcID = *serviceID.ID
			b, _ := json.MarshalIndent(serviceID, "", "  ")
			fmt.Println(string(b))

			// end-create_service_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(serviceID).ToNot(BeNil())
			Expect(svcID).ToNot(BeEmpty())
		})
		It(`GetServiceID request example`, func() {
			fmt.Println("\nGetServiceID() result:")
			// begin-get_service_id

			getServiceIDOptions := iamIdentityService.NewGetServiceIDOptions(svcID)

			getServiceIDOptions.SetIncludeActivity(false)

			serviceID, response, err := iamIdentityService.GetServiceID(getServiceIDOptions)
			if err != nil {
				panic(err)
			}
			svcIDEtag = response.GetHeaders().Get("Etag")
			b, _ := json.MarshalIndent(serviceID, "", "  ")
			fmt.Println(string(b))

			// end-get_service_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
			Expect(svcIDEtag).ToNot(BeEmpty())
		})
		It(`ListServiceIds request example`, func() {
			fmt.Println("\nListServiceIds() result:")
			// begin-list_service_ids

			listServiceIdsOptions := iamIdentityService.NewListServiceIdsOptions()
			listServiceIdsOptions.SetAccountID(accountID)
			listServiceIdsOptions.SetName(serviceIDName)

			serviceIDList, response, err := iamIdentityService.ListServiceIds(listServiceIdsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceIDList, "", "  ")
			fmt.Println(string(b))

			// end-list_service_ids

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceIDList).ToNot(BeNil())
		})
		It(`UpdateServiceID request example`, func() {
			fmt.Println("\nUpdateServiceID() result:")
			// begin-update_service_id

			updateServiceIDOptions := iamIdentityService.NewUpdateServiceIDOptions(svcID, svcIDEtag)
			updateServiceIDOptions.SetDescription("This is an updated description")

			serviceID, response, err := iamIdentityService.UpdateServiceID(updateServiceIDOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceID, "", "  ")
			fmt.Println(string(b))

			// end-update_service_id

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceID).ToNot(BeNil())
		})
		It(`LockServiceID request example`, func() {
			// begin-lock_service_id

			lockServiceIDOptions := iamIdentityService.NewLockServiceIDOptions(svcID)

			response, err := iamIdentityService.LockServiceID(lockServiceIDOptions)
			if err != nil {
				panic(err)
			}

			// end-lock_service_id
			fmt.Printf("\nLockServiceID() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`UnlockServiceID request example`, func() {
			// begin-unlock_service_id

			unlockServiceIDOptions := iamIdentityService.NewUnlockServiceIDOptions(svcID)

			response, err := iamIdentityService.UnlockServiceID(unlockServiceIDOptions)
			if err != nil {
				panic(err)
			}

			// end-unlock_service_id
			fmt.Printf("\nUnlockServiceID() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteServiceID request example`, func() {
			// begin-delete_service_id

			deleteServiceIDOptions := iamIdentityService.NewDeleteServiceIDOptions(svcID)

			response, err := iamIdentityService.DeleteServiceID(deleteServiceIDOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_service_id
			fmt.Printf("\nDeleteServiceID() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CreateProfile request example`, func() {
			fmt.Println("\nCreateProfile() result:")
			// begin-create_profile

			createProfileOptions := iamIdentityService.NewCreateProfileOptions(profileName, accountID)
			createProfileOptions.SetDescription("Example Profile")

			profile, response, err := iamIdentityService.CreateProfile(createProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))
			profileId = *profile.ID

			// end-create_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(profile).ToNot(BeNil())
			Expect(profileId).ToNot(BeNil())
		})
		It(`GetProfile request example`, func() {
			fmt.Println("\nGetProfile() result:")
			// begin-get_profile

			getProfileOptions := iamIdentityService.NewGetProfileOptions(profileId)

			getProfileOptions.SetIncludeActivity(false)

			profile, response, err := iamIdentityService.GetProfile(getProfileOptions)
			if err != nil {
				panic(err)
			}
			profileEtag = response.GetHeaders().Get("Etag")
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-get_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
			Expect(profileEtag).ToNot(BeEmpty())
		})
		It(`ListProfiles request example`, func() {
			fmt.Println("\nListProfiles() result:")
			// begin-list_profiles

			listProfilesOptions := iamIdentityService.NewListProfilesOptions(accountID)
			listProfilesOptions.SetIncludeHistory(false)

			trustedProfiles, response, err := iamIdentityService.ListProfiles(listProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trustedProfiles, "", "  ")
			fmt.Println(string(b))

			// end-list_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trustedProfiles).ToNot(BeNil())
		})
		It(`UpdateProfile request example`, func() {
			fmt.Println("\nUpdateProfile() result:")
			// begin-update_profile

			updateProfileOptions := iamIdentityService.NewUpdateProfileOptions(profileId, profileEtag)
			updateProfileOptions.SetDescription("This is an updated description")

			profile, response, err := iamIdentityService.UpdateProfile(updateProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profile, "", "  ")
			fmt.Println(string(b))

			// end-update_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())
		})
		It(`CreateClaimRule request example`, func() {
			fmt.Println("\nCreateClaimRule() result:")
			// begin-create_claim_rule

			profileClaimRuleConditions := new(iamidentityv1.ProfileClaimRuleConditions)
			profileClaimRuleConditions.Claim = core.StringPtr("blueGroups")
			profileClaimRuleConditions.Operator = core.StringPtr("EQUALS")
			profileClaimRuleConditions.Value = core.StringPtr("\"cloud-docs-dev\"")

			createClaimRuleOptions := iamIdentityService.NewCreateClaimRuleOptions(profileId, claimRuleType, []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditions})
			createClaimRuleOptions.SetName("claimRule")
			createClaimRuleOptions.SetRealmName(realmName)
			createClaimRuleOptions.SetExpiration(int64(43200))

			claimRule, response, err := iamIdentityService.CreateClaimRule(createClaimRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(claimRule, "", "  ")
			fmt.Println(string(b))
			claimRuleId = *claimRule.ID

			// end-create_claim_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(claimRule).ToNot(BeNil())
			Expect(claimRuleId).ToNot(BeNil())
		})
		It(`GetClaimRule request example`, func() {
			fmt.Println("\nGetClaimRule() result:")
			// begin-get_claim_rule

			getClaimRuleOptions := iamIdentityService.NewGetClaimRuleOptions(profileId, claimRuleId)

			claimRule, response, err := iamIdentityService.GetClaimRule(getClaimRuleOptions)
			if err != nil {
				panic(err)
			}
			claimRuleEtag = response.GetHeaders().Get("Etag")
			b, _ := json.MarshalIndent(claimRule, "", "  ")
			fmt.Println(string(b))

			// end-get_claim_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(claimRule).ToNot(BeNil())
			Expect(claimRuleEtag).ToNot(BeEmpty())
		})
		It(`ListClaimRules request example`, func() {
			fmt.Println("\nListClaimRules() result:")
			// begin-list_claim_rules

			listClaimRulesOptions := iamIdentityService.NewListClaimRulesOptions(profileId)

			claimRulesList, response, err := iamIdentityService.ListClaimRules(listClaimRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(claimRulesList, "", "  ")
			fmt.Println(string(b))

			// end-list_claim_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(claimRulesList).ToNot(BeNil())
		})
		It(`UpdateClaimRule request example`, func() {
			fmt.Println("\nUpdateClaimRule() result:")
			// begin-update_claim_rule

			profileClaimRuleConditions := new(iamidentityv1.ProfileClaimRuleConditions)
			profileClaimRuleConditions.Claim = core.StringPtr("blueGroups")
			profileClaimRuleConditions.Operator = core.StringPtr("EQUALS")
			profileClaimRuleConditions.Value = core.StringPtr("\"Europe_Group\"")

			updateClaimRuleOptions := iamIdentityService.NewUpdateClaimRuleOptions(profileId, claimRuleId, claimRuleEtag, claimRuleType, []iamidentityv1.ProfileClaimRuleConditions{*profileClaimRuleConditions})
			updateClaimRuleOptions.SetRealmName(realmName)
			updateClaimRuleOptions.SetExpiration(int64(33200))

			claimRule, response, err := iamIdentityService.UpdateClaimRule(updateClaimRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(claimRule, "", "  ")
			fmt.Println(string(b))

			// end-update_claim_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(claimRule).ToNot(BeNil())
		})
		It(`DeleteClaimRule request example`, func() {
			// begin-delete_claim_rule

			deleteClaimRuleOptions := iamIdentityService.NewDeleteClaimRuleOptions(profileId, claimRuleId)

			response, err := iamIdentityService.DeleteClaimRule(deleteClaimRuleOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_claim_rule
			fmt.Printf("\nDeleteClaimRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CreateLink request example`, func() {
			fmt.Println("\nCreateLink() result:")
			// begin-create_link

			createProfileLinkRequestLink := new(iamidentityv1.CreateProfileLinkRequestLink)
			createProfileLinkRequestLink.CRN = core.StringPtr("crn:v1:staging:public:iam-identity::a/" + accountID + "::computeresource:Fake-Compute-Resource")
			createProfileLinkRequestLink.Namespace = core.StringPtr("default")
			createProfileLinkRequestLink.Name = core.StringPtr("niceName")

			createLinkOptions := iamIdentityService.NewCreateLinkOptions(profileId, "ROKS_SA", createProfileLinkRequestLink)
			createLinkOptions.SetName("niceLink")

			link, response, err := iamIdentityService.CreateLink(createLinkOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(link, "", "  ")
			fmt.Println(string(b))
			linkId = *link.ID

			// end-create_link

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(link).ToNot(BeNil())
			Expect(linkId).ToNot(BeNil())
		})
		It(`GetLink request example`, func() {
			fmt.Println("\nGetLink() result:")
			// begin-get_link

			getLinkOptions := iamIdentityService.NewGetLinkOptions(profileId, linkId)

			link, response, err := iamIdentityService.GetLink(getLinkOptions)
			if err != nil {
				panic(err)
			}

			// end-get_link

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(link).ToNot(BeNil())
		})
		It(`ListLinks request example`, func() {
			fmt.Println("\nListLinks() result:")
			// begin-list_links

			listLinksOptions := iamIdentityService.NewListLinksOptions(profileId)

			linkList, response, err := iamIdentityService.ListLinks(listLinksOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(linkList, "", "  ")
			fmt.Println(string(b))

			// end-list_links

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(linkList).ToNot(BeNil())
		})
		It(`DeleteLink request example`, func() {
			// begin-delete_link

			deleteLinkOptions := iamIdentityService.NewDeleteLinkOptions(profileId, linkId)

			response, err := iamIdentityService.DeleteLink(deleteLinkOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_link
			fmt.Printf("\nDeleteLink() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`GetProfileIdentities request example`, func() {
			fmt.Println("\nGetProfileIdentities() result:")
			// begin-get_profile_identities

			getProfileIdentitiesOptions := iamidentityv1.GetProfileIdentitiesOptions{
				ProfileID: &profileId,
			}

			profileIdentities, response, err := iamIdentityService.GetProfileIdentities(&getProfileIdentitiesOptions)

			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileIdentities, "", "  ")
			fmt.Println(string(b))

			// end-get_profile_identities
			fmt.Printf("\nSetProfileIdentities() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileIdentities).ToNot(BeNil())
			Expect(profileEtag).ToNot(BeEmpty())
			profileEtag = *profileIdentities.EntityTag
		})
		It(`SetProfileIdentities request example`, func() {
			fmt.Println("\nSetProfileIdentities() result:")
			// begin-set_profile_identities

			accounts := []string{accountID}
			identity := &iamidentityv1.ProfileIdentityRequest{
				Identifier:  &iamID,
				Accounts:    accounts,
				Type:        core.StringPtr("user"),
				Description: core.StringPtr("Identity description"),
			}
			listProfileIdentity := []iamidentityv1.ProfileIdentityRequest{*identity}
			setProfileIdentitiesOptions := iamidentityv1.SetProfileIdentitiesOptions{
				ProfileID:  &profileId,
				Identities: listProfileIdentity,
				IfMatch:    &profileEtag,
			}

			profileIdnetities, response, err := iamIdentityService.SetProfileIdentities(&setProfileIdentitiesOptions)

			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileIdnetities, "", "  ")
			fmt.Println(string(b))

			// end-set_profile_identities
			fmt.Printf("\nSetProfileIdentities() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileIdnetities).ToNot(BeNil())
		})
		It(`SetProfileIdentity request example`, func() {
			fmt.Println("\nSetProfileIdentity() result:")
			// begin-set_profile_identity

			accounts := []string{accountID}
			setProfileIdentityOptions := iamidentityv1.SetProfileIdentityOptions{
				ProfileID:    &profileId,
				IdentityType: core.StringPtr("user"),
				Identifier:   &iamIDMember,
				Accounts:     accounts,
				Type:         core.StringPtr("user"),
				Description:  core.StringPtr("Identity description"),
			}

			profileIdnetity, response, err := iamIdentityService.SetProfileIdentity(&setProfileIdentityOptions)

			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileIdnetity, "", "  ")
			fmt.Println(string(b))

			// end-set_profile_identity
			fmt.Printf("\nSetProfileIdentity() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileIdnetity).ToNot(BeNil())
		})
		It(`GetProfileIdentity request example`, func() {
			fmt.Println("\nGetProfileIdentity() result:")
			// begin-get_profile_identity

			getProfileIdentityOptions := iamidentityv1.GetProfileIdentityOptions{
				ProfileID:    &profileId,
				IdentityType: core.StringPtr("user"),
				IdentifierID: &iamIDMember,
			}

			profileIdnetity, response, err := iamIdentityService.GetProfileIdentity(&getProfileIdentityOptions)

			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(profileIdnetity, "", "  ")
			fmt.Println(string(b))

			// end-get_profile_identity
			fmt.Printf("\nGetProfileIdentity() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profileIdnetity).ToNot(BeNil())
		})
		It(`DeleteProfileIdentity request example`, func() {
			fmt.Println("\nDeleteProfileIdentity() result:")
			// begin-delete_profile_identity

			deleteProfileIdentityOptions := iamidentityv1.DeleteProfileIdentityOptions{
				ProfileID:    &profileId,
				IdentityType: core.StringPtr("user"),
				IdentifierID: &iamIDMember,
			}

			response, err := iamIdentityService.DeleteProfileIdentity(&deleteProfileIdentityOptions)

			if err != nil {
				panic(err)
			}

			// end-delete_profile_identity
			fmt.Printf("\nGetProfileIdentity() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteProfile request example`, func() {
			// begin-delete_profile

			deleteProfileOptions := iamIdentityService.NewDeleteProfileOptions(profileId)

			response, err := iamIdentityService.DeleteProfile(deleteProfileOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_profile
			fmt.Printf("\nDeleteProfile() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`GetAccountSettings request example`, func() {
			fmt.Println("\nGetAccountSettings() result:")
			// begin-getAccountSettings

			getAccountSettingsOptions := iamIdentityService.NewGetAccountSettingsOptions(accountID)

			accountSettingsResponse, response, err := iamIdentityService.GetAccountSettings(getAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettingsResponse, "", "  ")
			fmt.Println(string(b))

			// end-getAccountSettings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettingsResponse).ToNot(BeNil())

			accountSettingEtag = response.GetHeaders().Get("Etag")
			Expect(accountSettingEtag).ToNot(BeEmpty())
		})
		It(`UpdateAccountSettings request example`, func() {
			fmt.Println("\nUpdateAccountSettings() result:")
			// begin-updateAccountSettings

			accountSettingsUserMFA := new(iamidentityv1.AccountSettingsUserMfa)
			accountSettingsUserMFA.IamID = core.StringPtr(iamIDMember)
			accountSettingsUserMFA.Mfa = core.StringPtr("NONE")

			updateAccountSettingsOptions := iamIdentityService.NewUpdateAccountSettingsOptions(
				accountSettingEtag,
				accountID,
			)
			updateAccountSettingsOptions.SetSessionExpirationInSeconds("86400")
			updateAccountSettingsOptions.SetSessionInvalidationInSeconds("7200")
			updateAccountSettingsOptions.SetMfa("NONE")
			updateAccountSettingsOptions.SetUserMfa([]iamidentityv1.AccountSettingsUserMfa{*accountSettingsUserMFA})
			updateAccountSettingsOptions.SetRestrictCreatePlatformApikey("NOT_RESTRICTED")
			updateAccountSettingsOptions.SetRestrictCreatePlatformApikey("NOT_RESTRICTED")
			updateAccountSettingsOptions.SetSystemAccessTokenExpirationInSeconds("3600")
			updateAccountSettingsOptions.SetSystemRefreshTokenExpirationInSeconds("259200")

			accountSettingsResponse, response, err := iamIdentityService.UpdateAccountSettings(updateAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettingsResponse, "", "  ")
			fmt.Println(string(b))

			// end-updateAccountSettings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettingsResponse).ToNot(BeNil())
		})
		It(`GetEffectiveAccountSettings request example`, func() {
			fmt.Println("\nGetEffectiveAccountSettings() result:")
			// begin-getEffectiveAccountSettings

			getEffectiveAccountSettingsOptions := iamIdentityService.NewGetEffectiveAccountSettingsOptions(accountID)

			effectiveAccountSettingsResponse, response, err := iamIdentityService.GetEffectiveAccountSettings(getEffectiveAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(effectiveAccountSettingsResponse, "", "  ")
			fmt.Println(string(b))

			// end-getEffectiveAccountSettings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(effectiveAccountSettingsResponse).ToNot(BeNil())
		})
		It(`CreateReport request example`, func() {
			fmt.Println("\nCreateReport() result:")
			// begin-create_report

			createReportOptions := iamIdentityService.NewCreateReportOptions(accountID)
			createReportOptions.SetType("inactive")
			createReportOptions.SetDuration("120")

			report, response, err := iamIdentityService.CreateReport(createReportOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(report, "", "  ")
			fmt.Println(string(b))

			// end-create_report

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(report).ToNot(BeNil())
		})
		It(`GetReport request example`, func() {
			fmt.Println("\nGetReport() result:")
			// begin-get_report

			getReportOptions := iamIdentityService.NewGetReportOptions(accountID, "latest")

			report, response, err := iamIdentityService.GetReport(getReportOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(report, "", "  ")
			fmt.Println(string(b))

			// end-get_report

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(report).ToNot(BeNil())
		})
		It(`CreateMfaReport request example`, func() {
			fmt.Println("\nCreateMfaReport() result:")
			// begin-create_mfa_report

			createMfaReportOptions := iamIdentityService.NewCreateMfaReportOptions(accountID)
			createMfaReportOptions.SetType("mfa_status")

			report, response, err := iamIdentityService.CreateMfaReport(createMfaReportOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(report, "", "  ")
			fmt.Println(string(b))

			// end-create_mfa_report

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(report).ToNot(BeNil())
		})
		It(`GetMfaReport request example`, func() {
			fmt.Println("\nGetMfaReport() result:")
			// begin-get_mfa_report

			getMfaReportOptions := iamIdentityService.NewGetMfaReportOptions(accountID, "latest")

			report, response, err := iamIdentityService.GetMfaReport(getMfaReportOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(report, "", "  ")
			fmt.Println(string(b))

			// end-get_mfa_report

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(report).ToNot(BeNil())
		})
		It(`GetMfaStatus request example`, func() {
			fmt.Println("\nGetMfaStatus() result:")
			// begin-get_mfa_status

			getMfaStatusOptions := iamIdentityService.NewGetMfaStatusOptions(accountID, iamID)

			mfaStatusResponse, response, err := iamIdentityService.GetMfaStatus(getMfaStatusOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(mfaStatusResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_mfa_status

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(mfaStatusResponse).ToNot(BeNil())
		})
		It(`CreateProfileTemplate request example`, func() {
			fmt.Println("\nCreateProfileTemplate() result:")
			// begin-create_profile_template
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
			profile.Description = core.StringPtr("Example Profile created from Profile Template")
			profile.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*profileTemplateClaimRule}

			createOptions := &iamidentityv1.CreateProfileTemplateOptions{
				Name:        &profileTemplateName,
				Description: core.StringPtr("Example Profile Template"),
				AccountID:   &enterpriseAccountID,
				Profile:     profile,
			}

			createResponse, response, err := iamIdentityService.CreateProfileTemplate(createOptions)

			b, _ := json.MarshalIndent(createResponse, "", "  ")
			fmt.Println(string(b))

			// Grab the ID and Etag value from the response for use in the update operation
			profileTemplateId = *createResponse.ID
			profileTemplateVersion = *createResponse.Version
			profileTemplateEtag = response.GetHeaders().Get("Etag")

			// end-create_profile_template
			Expect(response.StatusCode).To(Equal(201))
			Expect(err).To(BeNil())
			Expect(createResponse).ToNot(BeNil())
			Expect(profileTemplateId).ToNot(BeNil())
			Expect(profileTemplateEtag).ToNot(BeEmpty())
		})
		It(`GetProfileTemplateVersion request example`, func() {
			fmt.Println("\nGetProfileTemplateVersion() result:")
			// begin-get_profile_template_version

			getOptions := &iamidentityv1.GetProfileTemplateVersionOptions{
				TemplateID: &profileTemplateId,
				Version:    core.StringPtr(strconv.FormatInt(profileTemplateVersion, 10)),
			}
			getResponse, response, err := iamIdentityService.GetProfileTemplateVersion(getOptions)

			b, _ := json.MarshalIndent(getResponse, "", "  ")
			fmt.Println(string(b))

			// Grab the Etag value from the response for use in the update operation
			profileTemplateEtag = response.GetHeaders().Get("Etag")

			// end-get_profile_template_version
			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(getResponse).ToNot(BeNil())
			Expect(profileTemplateEtag).ToNot(BeEmpty())
		})
		It(`ListProfileTemplates request example`, func() {
			fmt.Println("\nListProfileTemplates() result:")
			// begin-list_profile_templates
			listOptions := &iamidentityv1.ListProfileTemplatesOptions{
				AccountID: &enterpriseAccountID,
			}
			listResponse, response, err := iamIdentityService.ListProfileTemplates(listOptions)

			b, _ := json.MarshalIndent(listResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_profile_templates
			Expect(response.StatusCode).To(Equal(200))
			Expect(err).To(BeNil())
			Expect(listResponse).ToNot(BeNil())
		})
		It(`UpdateProfileTemplateVersion request example`, func() {
			fmt.Println("\nUpdateProfileTemplateVersion() result:")
			// begin-update_profile_template_version

			updateOptions := &iamidentityv1.UpdateProfileTemplateVersionOptions{
				AccountID:   &enterpriseAccountID,
				TemplateID:  &profileTemplateId,
				Version:     core.StringPtr(strconv.FormatInt(profileTemplateVersion, 10)),
				IfMatch:     &profileTemplateEtag,
				Name:        &profileTemplateName,
				Description: core.StringPtr("Example Profile Template - updated"),
			}
			updateResponse, response, err := iamIdentityService.UpdateProfileTemplateVersion(updateOptions)

			b, _ := json.MarshalIndent(updateResponse, "", "  ")
			fmt.Println(string(b))

			// Grab the Etag value from the response for use in the update operation.
			profileTemplateEtag = response.GetHeaders().Get("Etag")

			// end-update_profile_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updateResponse).ToNot(BeNil())
			Expect(profileTemplateEtag).ToNot(BeEmpty())
		})
		It(`CommitProfileTemplate request example`, func() {
			fmt.Println("\nCommitProfileTemplate() result:")
			// begin-commit_profile_template

			commitOptions := &iamidentityv1.CommitProfileTemplateOptions{
				TemplateID: &profileTemplateId,
				Version:    core.StringPtr(strconv.FormatInt(profileTemplateVersion, 10)),
			}

			response, err := iamIdentityService.CommitProfileTemplate(commitOptions)

			// end-commit_profile_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CreateProfileTemplateAssignment request example`, func() {
			fmt.Println("\nCreateProfileTemplateAssignment() result:")
			// begin-create_trusted_profile_assignment

			assignOptions := &iamidentityv1.CreateTrustedProfileAssignmentOptions{
				TemplateID:      &profileTemplateId,
				TemplateVersion: &profileTemplateVersion,
				TargetType:      core.StringPtr("Account"),
				Target:          &enterpriseSubAccountID,
			}

			assignResponse, response, err := iamIdentityService.CreateTrustedProfileAssignment(assignOptions)

			b, _ := json.MarshalIndent(assignResponse, "", "  ")
			fmt.Println(string(b))

			// Grab the Etag and id for use by other test methods.
			profileTemplateAssignmentEtag = response.GetHeaders().Get("Etag")
			profileTemplateAssignmentId = *assignResponse.ID

			// end-create_trusted_profile_assignment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(assignResponse).ToNot(BeNil())
			Expect(profileTemplateAssignmentId).ToNot(BeNil())
			Expect(profileTemplateAssignmentEtag).ToNot(BeEmpty())
		})
		It(`GetProfileTemplateAssignment request example`, func() {
			fmt.Println("\nGetProfileTemplateAssignment() result:")
			// begin-get_trusted_profile_assignment

			getAssignmentOptions := &iamidentityv1.GetTrustedProfileAssignmentOptions{
				AssignmentID: &profileTemplateAssignmentId,
			}

			assignment, response, err := iamIdentityService.GetTrustedProfileAssignment(getAssignmentOptions)

			b, _ := json.MarshalIndent(assignment, "", "  ")
			fmt.Println(string(b))

			// end-get_trusted_profile_assignment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(assignment).ToNot(BeNil())
		})
		It(`ListTrustedProfileAssignments request example`, func() {
			fmt.Println("\nListTrustedProfileAssignments() result:")
			// begin-list_trusted_profile_assignments

			listOptions := &iamidentityv1.ListTrustedProfileAssignmentsOptions{
				AccountID:  &enterpriseAccountID,
				TemplateID: &profileTemplateId,
			}

			listResponse, response, err := iamIdentityService.ListTrustedProfileAssignments(listOptions)

			b, _ := json.MarshalIndent(listResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_trusted_profile_assignments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listResponse).ToNot(BeNil())
		})
		It(`CreateProfileTemplateVersion request example`, func() {
			fmt.Println("\nCreateProfileTemplateVersion() result:")
			// begin-create_profile_template_version

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
			profile.Description = core.StringPtr("Example Profile created from Profile Template - new version")
			profile.Rules = []iamidentityv1.TrustedProfileTemplateClaimRule{*profileTemplateClaimRule}

			createOptions := &iamidentityv1.CreateProfileTemplateVersionOptions{
				Name:        &profileTemplateName,
				Description: core.StringPtr("Example Profile Template - new version"),
				AccountID:   &enterpriseAccountID,
				TemplateID:  &profileTemplateId,
				Profile:     profile,
			}

			createResponse, response, err := iamIdentityService.CreateProfileTemplateVersion(createOptions)

			b, _ := json.MarshalIndent(createResponse, "", "  ")
			fmt.Println(string(b))

			// save the new version to be used in subsequent calls
			profileTemplateVersion = *createResponse.Version

			// end-create_profile_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createResponse).ToNot(BeNil())
		})
		It(`GetLatestProfileTemplateVersion request example`, func() {
			fmt.Println("\nGetLatestProfileTemplateVersion() result:")
			// begin-get_latest_profile_template_version

			getOptions := &iamidentityv1.GetLatestProfileTemplateVersionOptions{
				TemplateID: &profileTemplateId,
			}

			getResponse, response, err := iamIdentityService.GetLatestProfileTemplateVersion(getOptions)

			b, _ := json.MarshalIndent(getResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_latest_profile_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getResponse).ToNot(BeNil())
		})
		It(`ListVersionsOfProfileTemplate request example`, func() {
			fmt.Println("\nListVersionsOfProfileTemplate() result:")
			// begin-list_versions_of_profile_template

			listOptions := &iamidentityv1.ListVersionsOfProfileTemplateOptions{
				TemplateID: &profileTemplateId,
			}
			listResponse, response, err := iamIdentityService.ListVersionsOfProfileTemplate(listOptions)

			b, _ := json.MarshalIndent(listResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_versions_of_profile_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listResponse).ToNot(BeNil())
		})
		It(`UpdateTrustedProfileAssignment request example`, func() {

			commitOptions := &iamidentityv1.CommitProfileTemplateOptions{
				TemplateID: &profileTemplateId,
				Version:    core.StringPtr(strconv.FormatInt(profileTemplateVersion, 10)),
			}
			cResponse, cErr := iamIdentityService.CommitProfileTemplate(commitOptions)
			Expect(cResponse.StatusCode).To(Equal(204))
			Expect(cErr).To(BeNil())

			waitUntilTrustedProfileAssignmentFinished(iamIdentityService, &profileTemplateAssignmentId, &profileTemplateAssignmentEtag)

			fmt.Println("\nUpdateTrustedProfileAssignment() result:")
			// begin-update_trusted_profile_assignment

			updateOptions := &iamidentityv1.UpdateTrustedProfileAssignmentOptions{
				AssignmentID:    &profileTemplateAssignmentId,
				TemplateVersion: &profileTemplateVersion,
				IfMatch:         &profileTemplateAssignmentEtag,
			}

			updateResponse, response, err := iamIdentityService.UpdateTrustedProfileAssignment(updateOptions)

			b, _ := json.MarshalIndent(updateResponse, "", "  ")
			fmt.Println(string(b))

			// Grab the Etag and id for use by other test methods.
			profileTemplateAssignmentEtag = response.GetHeaders().Get("Etag")

			// end-update_trusted_profile_assignment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(updateResponse).ToNot(BeNil())
			Expect(profileTemplateAssignmentEtag).ToNot(BeEmpty())
		})
		It(`DeleteTrustedProfileAssignment request example`, func() {
			waitUntilTrustedProfileAssignmentFinished(iamIdentityService, &profileTemplateAssignmentId, &profileTemplateAssignmentEtag)

			fmt.Println("\nDeleteTrustedProfileAssignmentx() result:")
			// begin-delete_trusted_profile_assignment

			deleteOptions := &iamidentityv1.DeleteTrustedProfileAssignmentOptions{
				AssignmentID: &profileTemplateAssignmentId,
			}
			excResponse, response, err := iamIdentityService.DeleteTrustedProfileAssignment(deleteOptions)

			// end-delete_trusted_profile_assignment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(excResponse).To(BeNil())
		})
		It(`DeleteProfileTemplateVersion request example`, func() {
			fmt.Println("\nDeleteProfileTemplateVersion() result:")
			// begin-delete_profile_template_version

			deleteOptions := &iamidentityv1.DeleteProfileTemplateVersionOptions{
				TemplateID: &profileTemplateId,
				Version:    core.StringPtr("1"),
			}

			response, err := iamIdentityService.DeleteProfileTemplateVersion(deleteOptions)

			// end-delete_profile_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteProfileTemplateAllVersions request example`, func() {
			waitUntilTrustedProfileAssignmentFinished(iamIdentityService, &profileTemplateAssignmentId, &profileTemplateAssignmentEtag)

			fmt.Println("\nDeleteProfileTemplateAllVersions() result:")
			// begin-delete_all_versions_of_profile_template

			deleteOptions := &iamidentityv1.DeleteAllVersionsOfProfileTemplateOptions{
				TemplateID: &profileTemplateId,
			}

			response, err := iamIdentityService.DeleteAllVersionsOfProfileTemplate(deleteOptions)

			// end-delete_all_versions_of_profile_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`createAccountSettingsTemplate request example`, func() {

			fmt.Println("\ncreateAccountSettingsTemplate() result:")
			// begin-create_account_settings_template

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

			b, _ := json.MarshalIndent(createResponse, "", "  ")
			fmt.Println(string(b))

			// Grab the ID and Etag value from the response for use in the update operation.
			accountSettingsTemplateId = *createResponse.ID
			accountSettingsTemplateVersion = *createResponse.Version
			accountSettingsTemplateEtag = response.GetHeaders().Get("Etag")

			// end-create_account_settings_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createResponse).ToNot(BeNil())
			Expect(accountSettingsTemplateId).ToNot(BeNil())
			Expect(accountSettingsTemplateEtag).ToNot(BeEmpty())
		})
		It(`getAccountSettingsTemplateVersion request example`, func() {

			fmt.Println("\ngetAccountSettingsTemplateVersion() result:")
			// begin-get_account_settings_template_version

			getOptions := &iamidentityv1.GetAccountSettingsTemplateVersionOptions{
				TemplateID: &accountSettingsTemplateId,
				Version:    core.StringPtr(strconv.FormatInt(accountSettingsTemplateVersion, 10)),
			}

			getResponse, response, err := iamIdentityService.GetAccountSettingsTemplateVersion(getOptions)

			b, _ := json.MarshalIndent(getResponse, "", "  ")
			fmt.Println(string(b))

			// Grab the Etag value from the response for use in the update operation.
			accountSettingsTemplateEtag = response.GetHeaders().Get("Etag")

			// end-get_account_settings_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getResponse).ToNot(BeNil())
			Expect(accountSettingsTemplateEtag).ToNot(BeEmpty())
		})
		It(`listAccountSettingsTemplates request example`, func() {

			fmt.Println("\nlistAccountSettingsTemplates() result:")
			// begin-list_account_settings_templates

			listOptions := &iamidentityv1.ListAccountSettingsTemplatesOptions{
				AccountID: &enterpriseAccountID,
			}

			listResponse, response, err := iamIdentityService.ListAccountSettingsTemplates(listOptions)

			b, _ := json.MarshalIndent(listResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_account_settings_templates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listResponse).ToNot(BeNil())
		})
		It(`updateAccountSettingsTemplateVersion request example`, func() {

			fmt.Println("\nupdateAccountSettingsTemplateVersion() result:")
			// begin-update_account_settings_template_version

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

			b, _ := json.MarshalIndent(updateResponse, "", "  ")
			fmt.Println(string(b))

			// Grab the Etag value from the response for use in the update operation.
			accountSettingsTemplateEtag = response.GetHeaders().Get("Etag")

			// end-update_account_settings_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updateResponse).ToNot(BeNil())
			Expect(accountSettingsTemplateEtag).ToNot(BeEmpty())
		})
		It(`commitAccountSettingsTemplate request example`, func() {

			fmt.Println("\ncommitAccountSettingsTemplate() result:")
			// begin-commit_account_settings_template

			commitOptions := &iamidentityv1.CommitAccountSettingsTemplateOptions{
				TemplateID: &accountSettingsTemplateId,
				Version:    core.StringPtr(strconv.FormatInt(accountSettingsTemplateVersion, 10)),
			}

			response, err := iamIdentityService.CommitAccountSettingsTemplate(commitOptions)

			// end-commit_account_settings_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`createAccountSettingsAssignment request example`, func() {

			fmt.Println("\ncreateAccountSettingsAssignment() result:")
			// begin-create_account_settings_assignment

			assignOptions := &iamidentityv1.CreateAccountSettingsAssignmentOptions{
				TemplateID:      &accountSettingsTemplateId,
				TemplateVersion: &accountSettingsTemplateVersion,
				TargetType:      core.StringPtr("Account"),
				Target:          &enterpriseSubAccountID,
			}

			assignResponse, response, err := iamIdentityService.CreateAccountSettingsAssignment(assignOptions)

			b, _ := json.MarshalIndent(assignResponse, "", "  ")
			fmt.Println(string(b))

			// Grab the Etag and id for use by other test methods.
			accountSettingsTemplateAssignmentEtag = response.GetHeaders().Get("Etag")
			accountSettingsTemplateAssignmentId = *assignResponse.ID

			// end-create_account_settings_assignment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(assignResponse).ToNot(BeNil())
			Expect(accountSettingsTemplateAssignmentId).ToNot(BeNil())
			Expect(accountSettingsTemplateAssignmentEtag).ToNot(BeEmpty())
		})
		It(`listAccountSettingsAssignments request example`, func() {

			fmt.Println("\nlistAccountSettingsAssignments() result:")
			// begin-list_account_settings_assignments

			listOptions := &iamidentityv1.ListAccountSettingsAssignmentsOptions{
				AccountID:  &enterpriseAccountID,
				TemplateID: &accountSettingsTemplateId,
			}

			listResponse, response, err := iamIdentityService.ListAccountSettingsAssignments(listOptions)

			b, _ := json.MarshalIndent(listResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_account_settings_assignments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listResponse).ToNot(BeNil())
		})
		It(`getAccountSettingsAssignment request example`, func() {

			fmt.Println("\ngetAccountSettingsAssignment() result:")
			// begin-get_account_settings_assignment

			getAssignmentOptions := &iamidentityv1.GetAccountSettingsAssignmentOptions{
				AssignmentID: &accountSettingsTemplateAssignmentId,
			}

			assignment, response, err := iamIdentityService.GetAccountSettingsAssignment(getAssignmentOptions)

			b, _ := json.MarshalIndent(assignment, "", "  ")
			fmt.Println(string(b))

			// end-get_account_settings_assignment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(assignment).ToNot(BeNil())
		})
		It(`createAccountSettingsTemplateVersion request example`, func() {

			fmt.Println("\ncreateAccountSettingsTemplateVersion() result:")
			// begin-create_account_settings_template_version

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

			b, _ := json.MarshalIndent(createResponse, "", "  ")
			fmt.Println(string(b))

			// save the new version to be used in subsequent calls
			accountSettingsTemplateVersion = *createResponse.Version

			// end-create_account_settings_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createResponse).ToNot(BeNil())
			Expect(accountSettingsTemplateVersion).ToNot(BeNil())
		})
		It(`getLatestAccountSettingsTemplateVersion request example`, func() {

			fmt.Println("\ngetLatestAccountSettingsTemplateVersion() result:")
			// begin-get_latest_account_settings_template_version

			getOptions := &iamidentityv1.GetLatestAccountSettingsTemplateVersionOptions{
				TemplateID: &accountSettingsTemplateId,
			}

			getResponse, response, err := iamIdentityService.GetLatestAccountSettingsTemplateVersion(getOptions)

			b, _ := json.MarshalIndent(getResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_latest_account_settings_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getResponse).ToNot(BeNil())
		})
		It(`listVersionsOfAccountSettingsTemplate request example`, func() {

			fmt.Println("\nlistVersionsOfAccountSettingsTemplate() result:")
			// begin-list_versions_of_account_settings_template

			listOptions := &iamidentityv1.ListVersionsOfAccountSettingsTemplateOptions{
				TemplateID: &accountSettingsTemplateId,
			}

			listResponse, response, err := iamIdentityService.ListVersionsOfAccountSettingsTemplate(listOptions)

			b, _ := json.MarshalIndent(listResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_versions_of_account_settings_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listResponse).ToNot(BeNil())
		})
		It(`updateAccountSettingsAssignment request example`, func() {

			commitOptions := &iamidentityv1.CommitAccountSettingsTemplateOptions{
				TemplateID: &accountSettingsTemplateId,
				Version:    core.StringPtr(strconv.FormatInt(accountSettingsTemplateVersion, 10)),
			}
			cResponse, cErr := iamIdentityService.CommitAccountSettingsTemplate(commitOptions)
			Expect(cResponse.StatusCode).To(Equal(204))
			Expect(cErr).To(BeNil())

			waitUntilAccountSettingsAssignmentFinished(iamIdentityService, &accountSettingsTemplateAssignmentId, &accountSettingsTemplateAssignmentEtag)

			fmt.Println("\nupdateAccountSettingsAssignment() result:")
			// begin-update_account_settings_assignment

			updateOptions := &iamidentityv1.UpdateAccountSettingsAssignmentOptions{
				AssignmentID:    &accountSettingsTemplateAssignmentId,
				TemplateVersion: &accountSettingsTemplateVersion,
				IfMatch:         &accountSettingsTemplateAssignmentEtag,
			}

			updateResponse, response, err := iamIdentityService.UpdateAccountSettingsAssignment(updateOptions)

			b, _ := json.MarshalIndent(updateResponse, "", "  ")
			fmt.Println(string(b))

			// end-update_account_settings_assignment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(updateResponse).ToNot(BeNil())
		})
		It(`deleteAccountSettingsAssignment request example`, func() {

			waitUntilAccountSettingsAssignmentFinished(iamIdentityService, &accountSettingsTemplateAssignmentId, &accountSettingsTemplateAssignmentEtag)

			fmt.Println("\ndeleteAccountSettingsAssignment() result:")
			// begin-delete_account_settings_assignment

			deleteOptions := &iamidentityv1.DeleteAccountSettingsAssignmentOptions{
				AssignmentID: &accountSettingsTemplateAssignmentId,
			}

			excResponse, response, err := iamIdentityService.DeleteAccountSettingsAssignment(deleteOptions)

			// end-delete_account_settings_assignment

			Expect(response.StatusCode).To(Equal(202))
			Expect(err).To(BeNil())
			Expect(excResponse).To(BeNil())
		})
		It(`deleteAccountSettingsTemplateVersion request example`, func() {

			fmt.Println("\ndeleteAccountSettingsTemplateVersion() result:")
			// begin-delete_account_settings_template_version

			deleteOptions := &iamidentityv1.DeleteAccountSettingsTemplateVersionOptions{
				TemplateID: &accountSettingsTemplateId,
				Version:    core.StringPtr("1"),
			}

			response, err := iamIdentityService.DeleteAccountSettingsTemplateVersion(deleteOptions)

			// end-delete_account_settings_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`deleteAllVersionsOfAccountSettingsTemplate request example`, func() {

			waitUntilAccountSettingsAssignmentFinished(iamIdentityService, &accountSettingsTemplateAssignmentId, &accountSettingsTemplateAssignmentEtag)

			fmt.Println("\ndeleteAllVersionsOfAccountSettingsTemplate() result:")
			// begin-delete_all_versions_of_account_settings_template

			deleteOptions := &iamidentityv1.DeleteAllVersionsOfAccountSettingsTemplateOptions{
				TemplateID: &accountSettingsTemplateId,
			}

			response, err := iamIdentityService.DeleteAllVersionsOfAccountSettingsTemplate(deleteOptions)

			// end-delete_all_versions_of_account_settings_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`UpdatePreferenceOnScopeAccount request example`, func() {

			// begin-update_preference_on_scope_account

			updatePreferenceOnScopeAccountOptions := &iamidentityv1.UpdatePreferenceOnScopeAccountOptions{
				AccountID:    &accountID,
				IamID:        &iamIDForPreferences,
				Service:      &service,
				PreferenceID: &preferenceID1,
				ValueString:  &valueString,
			}

			preference, response, err := iamIdentityService.UpdatePreferenceOnScopeAccount(updatePreferenceOnScopeAccountOptions)

			// end-update_preference_on_scope_account

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(preference).ToNot(BeNil())
		})
		It(`GetPreferencesOnScopeAccount request example`, func() {

			// begin-get_preferences_on_scope_account

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

			// end-get_preferences_on_scope_account
		})
		It(`GetAllPreferencesOnScopeAccount request example`, func() {

			// begin-get_all_preferences_on_scope_account

			getAllPreferencesOnScopeAccountOptions := &iamidentityv1.GetAllPreferencesOnScopeAccountOptions{
				AccountID: &accountID,
				IamID:     &iamIDForPreferences,
			}

			preference, response, err := iamIdentityService.GetAllPreferencesOnScopeAccount(getAllPreferencesOnScopeAccountOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(preference).ToNot(BeNil())

			// end-get_all_preferences_on_scope_account
		})
		It(`DeletePreferencesOnScopeAccount request example`, func() {

			// begin-delete_preferences_on_scope_account

			deletePreferencesOnScopeAccountOptions := &iamidentityv1.DeletePreferencesOnScopeAccountOptions{
				AccountID:    &accountID,
				IamID:        &iamIDForPreferences,
				Service:      &service,
				PreferenceID: &preferenceID1,
			}

			response, err := iamIdentityService.DeletePreferencesOnScopeAccount(deletePreferencesOnScopeAccountOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

			// end-delete_preferences_on_scope_account
		})
	})
})

func isFinished(status *string) bool {
	var finished = false
	if strings.EqualFold(*status, "succeeded") || strings.EqualFold(*status, "failed") {
		finished = true
	}
	return finished
}

func waitUntilTrustedProfileAssignmentFinished(service *iamidentityv1.IamIdentityV1, assignmentId *string, profileTemplateAssignmentEtag *string) {
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
			finished = isFinished(assignment.Status)
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

func waitUntilAccountSettingsAssignmentFinished(service *iamidentityv1.IamIdentityV1, assignmentId *string, accountSettingsTemplateAssignmentEtag *string) {
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
			finished = isFinished(assignment.Status)
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
