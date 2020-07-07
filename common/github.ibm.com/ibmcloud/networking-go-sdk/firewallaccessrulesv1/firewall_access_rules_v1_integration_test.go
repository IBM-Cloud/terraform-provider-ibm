/*
 * (C) Copyright IBM Corp. 2020.
 */

package firewallaccessrulesv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/firewallaccessrulesv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`AccountFirewallAccessRulesV1`, func() {
	if _, err := os.Stat(configFile); err != nil {
		configLoaded = false
	}

	err := godotenv.Load(configFile)
	if err != nil {
		configLoaded = false
	}

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("CIS_SERVICES_APIKEY"),
		URL:    os.Getenv("CIS_SERVICES_AUTH_URL"),
	}
	serviceURL := os.Getenv("API_ENDPOINT")
	crn := os.Getenv("CRN")
	globalOptions := &FirewallAccessRulesV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
	}
	testService, testServiceErr := NewFirewallAccessRulesV1(globalOptions)
	if testServiceErr != nil {
		fmt.Println(testServiceErr)
	}

	// CIS_Frontend_API_Spec-Account_Firewall_AccessRules.yaml file integration test

	Describe(`AccessRuleActions`, func() {
		Context("AccountAccessRuleActions", func() {
			BeforeEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllAccountAccessRules(testService.NewListAllAccountAccessRulesOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				for _, rule := range result.Result {
					option := testService.NewDeleteAccountAccessRuleOptions(*rule.ID)
					delResult, response, err := testService.DeleteAccountAccessRule(option)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllAccountAccessRules(testService.NewListAllAccountAccessRulesOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				for _, rule := range result.Result {
					option := testService.NewDeleteAccountAccessRuleOptions(*rule.ID)
					delResult, response, err := testService.DeleteAccountAccessRule(option)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())
				}
			})
			It(`create/update/delete/get account access rules`, func() {
				shouldSkipTest()
				modes := [4]string{
					CreateAccountAccessRuleOptions_Mode_Block,
					CreateAccountAccessRuleOptions_Mode_Challenge,
					CreateAccountAccessRuleOptions_Mode_JsChallenge,
					CreateAccountAccessRuleOptions_Mode_Whitelist,
				}
				targets := [4]string{
					AccountAccessRuleObjectConfiguration_Target_Ip,
					AccountAccessRuleObjectConfiguration_Target_IpRange,
					AccountAccessRuleObjectConfiguration_Target_Asn,
					AccountAccessRuleObjectConfiguration_Target_Country,
				}
				configData := [4]string{
					"172.168.1.1",
					"172.168.1.0/24",
					"AS12345",
					"US",
				}

				newModes := [4]string{
					CreateAccountAccessRuleOptions_Mode_Block,
					CreateAccountAccessRuleOptions_Mode_Challenge,
					CreateAccountAccessRuleOptions_Mode_JsChallenge,
					CreateAccountAccessRuleOptions_Mode_Whitelist,
				}

				for i, mode := range modes {
					options := testService.NewCreateAccountAccessRuleOptions()
					options.SetMode(mode)

					configOpt, err := testService.NewAccountAccessRuleInputConfiguration(targets[i], configData[i])
					Expect(err).To(BeNil())
					Expect(configOpt).ToNot(BeNil())

					options.SetConfiguration(configOpt)
					options.SetNotes("This rule is added because of event X that occurred on date xyz")
					// create zone access rule
					result, response, err := testService.CreateAccountAccessRule(options)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())

					// get account access rule
					getOption := testService.NewGetAccountAccessRuleOptions(*result.Result.ID)
					result, response, err = testService.GetAccountAccessRule(getOption)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())

					// update account access rule
					updateOpt := testService.NewUpdateAccountAccessRuleOptions(*result.Result.ID)
					updateOpt.SetMode(newModes[i])
					result, response, err = testService.UpdateAccountAccessRule(updateOpt)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())

					// delete zone access rule
					deleteOpt := testService.NewDeleteAccountAccessRuleOptions(*result.Result.ID)
					delResult, response, err := testService.DeleteAccountAccessRule(deleteOpt)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())
				}
			})
		})

		It(`List All Account access rules`, func() {
			shouldSkipTest()
			modes := [4]string{
				CreateAccountAccessRuleOptions_Mode_Block,
				CreateAccountAccessRuleOptions_Mode_Challenge,
				CreateAccountAccessRuleOptions_Mode_JsChallenge,
				CreateAccountAccessRuleOptions_Mode_Whitelist,
			}
			targets := [4]string{
				AccountAccessRuleObjectConfiguration_Target_Ip,
				AccountAccessRuleObjectConfiguration_Target_IpRange,
				AccountAccessRuleObjectConfiguration_Target_Asn,
				AccountAccessRuleObjectConfiguration_Target_Country,
			}
			configData := [4]string{
				"172.168.1.1",
				"172.168.1.0/24",
				"AS12345",
				"US",
			}

			for i, mode := range modes {
				options := testService.NewCreateAccountAccessRuleOptions()
				options.SetMode(mode)

				configOpt, err := testService.NewAccountAccessRuleInputConfiguration(targets[i], configData[i])
				Expect(err).To(BeNil())
				Expect(configOpt).ToNot(BeNil())

				options.SetConfiguration(configOpt)
				options.SetNotes("This rule is added because of event X that occurred on date xyz")
				// create account access rule
				result, response, err := testService.CreateAccountAccessRule(options)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
			}

			result, response, operationErr := testService.ListAllAccountAccessRules(testService.NewListAllAccountAccessRulesOptions())
			Expect(operationErr).To(BeNil())
			Expect(response).ToNot(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(*result.Success).Should(BeTrue())
			for _, rule := range result.Result {
				option := testService.NewDeleteAccountAccessRuleOptions(*rule.ID)
				delResult, response, err := testService.DeleteAccountAccessRule(option)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
			}
		})
	})
})
