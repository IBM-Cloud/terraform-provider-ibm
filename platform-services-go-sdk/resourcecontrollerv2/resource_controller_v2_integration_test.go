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
package resourcecontrollerv2_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const externalConfigFile = "../resource_controller.env"

const resultsPerPage = 20

var (
	resourceControllerService *resourcecontrollerv2.ResourceControllerV2
	err                       error
	configLoaded              bool = false
	config                    map[string]string

	instanceNames map[string]string
	aliasNames    map[string]string
	bindingNames  map[string]string
	keyNames      map[string]string

	testAccountID                string
	testResourceGroupGUID        string
	testOrgGUID                  string
	testSpaceGUID                string
	testAppGUID                  string
	testRegionID1                string
	testPlanID1                  string
	testRegionID2                string
	testPlanID2                  string
	testReclaimInstanceName      string
	testLockedInstanceNameUpdate string

	//result info
	testInstanceCRN         string
	testInstanceGUID        string
	testAliasCRN            string
	testAliasGUID           string
	testBindingCRN          string
	testBindingGUID         string
	testInstanceKeyCRN      string
	testInstanceKeyGUID     string
	testAliasKeyCRN         string
	testAliasKeyGUID        string
	aliasTargetCRN          string
	bindTargetCRN           string
	testReclaimInstanceCRN  string
	testReclaimInstanceGUID string
	testReclamationID1      string
	testReclamationID2      string

	transactionID string = uuid.New().String()
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe("Resource Controller - Integration Tests", func() {

	fmt.Fprintln(GinkgoWriter, "Transaction ID for this test run: ", transactionID)

	It("Successfully load the configuration", func() {
		var err error
		_, err = os.Stat(externalConfigFile)
		if err != nil {
			Skip("External configuration file not found, skipping tests: " + err.Error())
		}

		os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
		config, err = core.GetServiceProperties(resourcecontrollerv2.DefaultServiceName)
		if err != nil {
			Skip("Error loading service properties, skipping tests: " + err.Error())
		}

		configLoaded = len(config) > 0

		testResourceGroupGUID = config["RESOURCE_GROUP"]
		Expect(testResourceGroupGUID).ToNot(BeEmpty())

		testPlanID2 = config["RECLAMATION_PLAN_ID"]
		Expect(testPlanID2).ToNot(BeEmpty())

		testAccountID = config["ACCOUNT_ID"]
		Expect(testAccountID).ToNot(BeEmpty())

		testOrgGUID = config["ORGANIZATION_GUID"]
		Expect(testOrgGUID).ToNot(BeEmpty())

		testSpaceGUID = config["SPACE_GUID"]
		Expect(testSpaceGUID).ToNot(BeEmpty())

		testAppGUID = config["APPLICATION_GUID"]
		Expect(testAppGUID).ToNot(BeEmpty())

		testPlanID1 = config["PLAN_ID"]
		Expect(testPlanID1).ToNot(BeEmpty())

		testRegionID1 = "global"
		testRegionID2 = "global"
		testReclaimInstanceName = "RcSdkReclaimInstance1"
		testLockedInstanceNameUpdate = "RcSdkLockedInstanceUpdate1"

		instanceNames = map[string]string{
			"name":   "RcSdkInstance1Go",
			"update": "RcSdkUpdateInstance1Go",
		}
		aliasNames = map[string]string{
			"name":   "RcSdkAlias1Go",
			"update": "RcSdkAliasUpdate1Go",
		}
		bindingNames = map[string]string{
			"name":   "RcSdkBinding1Go",
			"update": "RcSdkBindingUpdate1Go",
		}
		keyNames = map[string]string{
			"name":    "RcSdkKey1Go",
			"update":  "RcSdkKeyUpdate1Go",
			"name2":   "RcSdkKey2Go",
			"update2": "RcSdkKeyUpdate2Go",
		}
	})

	It(`Successfully created ResourceControllerV2 service instances`, func() {
		shouldSkipTest()

		resourceControllerService, err = resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(
			&resourcecontrollerv2.ResourceControllerV2Options{},
		)

		Expect(err).To(BeNil())
		Expect(resourceControllerService).ToNot(BeNil())

		core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
		resourceControllerService.EnableRetries(4, 30*time.Second)

		//setting timeout to 1 minute
		resourceControllerService.Service.Client.Timeout = 1 * time.Minute
		fmt.Fprintln(GinkgoWriter, "Timeout set to: ", resourceControllerService.Service.Client.Timeout)
	})

	Describe("Create, Retrieve, and Update Resource Instance", func() {
		It("00 - Create Resource Instance", func() {
			shouldSkipTest()

			options := resourceControllerService.NewCreateResourceInstanceOptions(
				instanceNames["name"],
				testRegionID1,
				testResourceGroupGUID,
				testPlanID1,
			)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test00-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.CreateResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "CreateResourceInstance() result:\n%s\n", common.ToJSON(result))

			Expect(result.ID).NotTo(BeNil())
			Expect(result.GUID).NotTo(BeNil())
			Expect(result.CRN).NotTo(BeNil())
			Expect(*result.ID).To(Equal(*result.CRN))
			Expect(*result.Name).To(Equal(instanceNames["name"]))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGUID))
			Expect(*result.ResourcePlanID).To(Equal(testPlanID1))
			Expect(*result.State).To(Equal("active"))
			Expect(*result.Locked).Should(BeFalse())
			Expect(*result.LastOperation.Type).To(Equal("create"))
			Expect(*result.LastOperation.Async).Should(BeFalse())
			Expect(*result.LastOperation.State).To(Equal("succeeded"))

			testInstanceCRN = *result.ID
			testInstanceGUID = *result.GUID
		})

		It("01 - Get A Resource Instance", func() {
			shouldSkipTest()

			options := resourceControllerService.NewGetResourceInstanceOptions(testInstanceGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test01-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.GetResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "GetResourceInstance() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ID).To(Equal(testInstanceCRN))
			Expect(*result.OnetimeCredentials).ToNot(BeNil())
			Expect(*result.GUID).To(Equal(testInstanceGUID))
			Expect(*result.CRN).To(Equal(testInstanceCRN))
			Expect(*result.Name).To(Equal(instanceNames["name"]))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGUID))
			Expect(*result.ResourcePlanID).To(Equal(testPlanID1))
			Expect(*result.State).To(Equal("active"))
			Expect(*result.Locked).Should(BeFalse())
			Expect(*result.LastOperation.Type).To(Equal("create"))
			Expect(*result.LastOperation.Async).Should(BeFalse())
			Expect(*result.LastOperation.State).To(Equal("succeeded"))
		})

		It("02 - Update A Resource Instance", func() {
			shouldSkipTest()

			options := resourceControllerService.NewUpdateResourceInstanceOptions(testInstanceGUID)
			options.SetName(instanceNames["update"])

			params := make(map[string]interface{}, 0)
			params["hello"] = "bye"
			options.SetParameters(params)

			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test02-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.UpdateResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "UpdateResourceInstance() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ID).To(Equal(testInstanceCRN))
			Expect(*result.Name).To(Equal(instanceNames["update"]))
			Expect(*result.State).To(Equal("active"))
			Expect(*result.LastOperation.Type).To(Equal("update"))
			Expect(*result.LastOperation.Async).Should(BeFalse())
			Expect(*result.LastOperation.State).To(Equal("succeeded"))
		})

		Describe("List Resource Instances", func() {
			It("03 - List Resource Instances With No Filter", func() {
				shouldSkipTest()

				options := resourceControllerService.NewListResourceInstancesOptions()
				options.SetLimit(resultsPerPage)
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test03-" + transactionID,
				}
				options.SetHeaders(headers)

				results := []resourcecontrollerv2.ResourceInstance{}

				for {
					result, resp, err := resourceControllerService.ListResourceInstances(options)

					//should return one or more instances
					Expect(err).To(BeNil())
					Expect(resp.StatusCode).To(Equal(200))
					Expect(result).ToNot(BeNil())

					Expect(*result.RowsCount).Should(BeNumerically(">=", int64(1)))
					Expect(*result.RowsCount).Should(BeNumerically("<=", int64(resultsPerPage)))
					Expect(len(result.Resources)).Should(BeNumerically(">=", 1))
					Expect(len(result.Resources)).Should(BeNumerically("<=", resultsPerPage))

					results = append(results, result.Resources...)

					start, err := core.GetQueryParam(result.NextURL, "start")
					Expect(err).To(BeNil())

					if start == nil {
						break
					}

					options.SetStart(*start)
				}

				fmt.Fprintf(GinkgoWriter, "ListResourceInstances() result:\n%s\n", common.ToJSON(results))

			})

			It("04 - List Resource Instances With GUID Filter", func() {
				shouldSkipTest()

				options := resourceControllerService.NewListResourceInstancesOptions()
				options.SetGUID(testInstanceGUID)
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test04-" + transactionID,
				}
				options.SetHeaders(headers)
				result, resp, err := resourceControllerService.ListResourceInstances(options)

				//should return list with only newly created instance
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

				fmt.Fprintf(GinkgoWriter, "ListResourceInstances() result:\n%s\n", common.ToJSON(result))

				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
				Expect(*result.Resources[0].ID).To(Equal(testInstanceCRN))
				Expect(*result.Resources[0].GUID).To(Equal(testInstanceGUID))
				Expect(*result.Resources[0].Name).To(Equal(instanceNames["update"]))
				Expect(*result.Resources[0].State).To(Equal("active"))
				Expect(*result.Resources[0].LastOperation.Type).To(Equal("update"))
				Expect(*result.Resources[0].LastOperation.Async).Should(BeFalse())
				Expect(*result.Resources[0].LastOperation.State).To(Equal("succeeded"))
			})

			It("05 - List Resource Instances With Name Filter", func() {
				shouldSkipTest()

				options := resourceControllerService.NewListResourceInstancesOptions()
				options.SetName(instanceNames["update"])
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test05-" + transactionID,
				}
				options.SetHeaders(headers)
				result, resp, err := resourceControllerService.ListResourceInstances(options)

				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

				fmt.Fprintf(GinkgoWriter, "ListResourceInstances() result:\n%s\n", common.ToJSON(result))

				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
			})
			It(`ListResourceInstances(listResourceInstancesOptions *ListResourceInstancesOptions) using ResourceInstancesPager`, func() {
				listResourceInstancesOptions := &resourcecontrollerv2.ListResourceInstancesOptions{}

				// Test GetNext().
				pager, err := resourceControllerService.NewResourceInstancesPager(listResourceInstancesOptions)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourcecontrollerv2.ResourceInstance
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}

				// Test GetAll().
				pager, err = resourceControllerService.NewResourceInstancesPager(listResourceInstancesOptions)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allItems, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allItems).ToNot(BeNil())

				Expect(len(allItems)).To(Equal(len(allResults)))
				fmt.Fprintf(GinkgoWriter, "ListResourceInstances() returned a total of %d item(s) using ResourceInstancesPager.\n", len(allResults))
			})
		})
	})

	Describe("Create, Retrieve, and Update Resource Alias", func() {
		It("06 - Create Resource Alias", func() {
			shouldSkipTest()

			target := "crn:v1:bluemix:public:bluemix:us-south:o/" + testOrgGUID + "::cf-space:" + testSpaceGUID
			aliasTargetCRN = "crn:v1:bluemix:public:cf:us-south:o/" + testOrgGUID + "::cf-space:" + testSpaceGUID
			options := resourceControllerService.NewCreateResourceAliasOptions(aliasNames["name"], testInstanceGUID, target)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test06-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.CreateResourceAlias(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "CreateResourceAlias() result:\n%s\n", common.ToJSON(result))

			Expect(result.ID).NotTo(BeNil())
			Expect(result.GUID).NotTo(BeNil())
			Expect(result.CRN).NotTo(BeNil())
			Expect(*result.ID).To(Equal(*result.CRN))
			Expect(*result.Name).To(Equal(aliasNames["name"]))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGUID))
			Expect(*result.TargetCRN).To(Equal(aliasTargetCRN))
			Expect(*result.State).To(Equal("active"))
			Expect(*result.ResourceInstanceID).To(Equal(testInstanceCRN))

			testAliasCRN = *result.ID
			testAliasGUID = *result.GUID
		})

		It("07 - Get A Resource Alias", func() {
			shouldSkipTest()

			Expect(testAliasGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewGetResourceAliasOptions(testAliasGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test07-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.GetResourceAlias(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "GetResourceAlias() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ID).To(Equal(testAliasCRN))
			Expect(*result.GUID).To(Equal(testAliasGUID))
			Expect(*result.CRN).To(Equal(testAliasCRN))
			Expect(*result.Name).To(Equal(aliasNames["name"]))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGUID))
			Expect(*result.TargetCRN).To(Equal(aliasTargetCRN))
			Expect(*result.State).To(Equal("active"))
			Expect(*result.ResourceInstanceID).To(Equal(testInstanceCRN))
		})

		It("08 - Update A Resource Alias", func() {
			shouldSkipTest()

			Expect(testAliasGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewUpdateResourceAliasOptions(testAliasGUID, aliasNames["update"])
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test08-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.UpdateResourceAlias(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "UpdateResourceAlias() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ID).To(Equal(testAliasCRN))
			Expect(*result.Name).To(Equal(aliasNames["update"]))
			Expect(*result.State).To(Equal("active"))
		})

		Describe("List Resource Aliases", func() {
			It("09 - List Resource Aliases With No Filter", func() {
				shouldSkipTest()

				options := resourceControllerService.NewListResourceAliasesOptions()
				options.SetLimit(resultsPerPage)
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test09-" + transactionID,
				}
				options.SetHeaders(headers)

				results := []resourcecontrollerv2.ResourceAlias{}

				for {
					result, resp, err := resourceControllerService.ListResourceAliases(options)

					//should return one or more aliases
					Expect(err).To(BeNil())
					Expect(resp.StatusCode).To(Equal(200))
					Expect(result).ToNot(BeNil())

					Expect(*result.RowsCount).Should(BeNumerically(">=", int64(1)))
					Expect(*result.RowsCount).Should(BeNumerically("<=", int64(resultsPerPage)))
					Expect(len(result.Resources)).Should(BeNumerically(">=", 1))
					Expect(len(result.Resources)).Should(BeNumerically("<=", resultsPerPage))

					results = append(results, result.Resources...)

					start, err := core.GetQueryParam(result.NextURL, "start")
					Expect(err).To(BeNil())

					if start == nil {
						break
					}

					options.SetStart(*start)

				}

				fmt.Fprintf(GinkgoWriter, "ListResourceAliases() result:\n%s\n", common.ToJSON(results))
			})

			It("10 - List Resource Aliases With GUID Filter", func() {
				shouldSkipTest()

				Expect(testAliasGUID).ToNot(BeEmpty())

				options := resourceControllerService.NewListResourceAliasesOptions()
				options.SetGUID(testAliasGUID)
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test10-" + transactionID,
				}
				options.SetHeaders(headers)
				result, resp, err := resourceControllerService.ListResourceAliases(options)

				//should return list with only newly created alias
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

				fmt.Fprintf(GinkgoWriter, "ListResourceAliases() result:\n%s\n", common.ToJSON(result))

				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
				Expect(*result.Resources[0].ID).To(Equal(testAliasCRN))
				Expect(*result.Resources[0].Name).To(Equal(aliasNames["update"]))
				Expect(*result.Resources[0].ResourceGroupID).To(Equal(testResourceGroupGUID))
				Expect(*result.Resources[0].TargetCRN).To(Equal(aliasTargetCRN))
				Expect(*result.Resources[0].State).To(Equal("active"))
				Expect(*result.Resources[0].ResourceInstanceID).To(Equal(testInstanceCRN))
			})

			It("11 - List Resource Aliases With Name Filter", func() {
				shouldSkipTest()

				options := resourceControllerService.NewListResourceAliasesOptions()
				options.SetName(aliasNames["update"])
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test11-" + transactionID,
				}
				options.SetHeaders(headers)
				result, resp, err := resourceControllerService.ListResourceAliases(options)

				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

				fmt.Fprintf(GinkgoWriter, "ListResourceAliases() result:\n%s\n", common.ToJSON(result))

				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
			})
			It(`ListResourceAliases(listResourceAliasesOptions *ListResourceAliasesOptions) using ResourceAliasesPager`, func() {
				listResourceAliasesOptions := &resourcecontrollerv2.ListResourceAliasesOptions{}

				// Test GetNext().
				pager, err := resourceControllerService.NewResourceAliasesPager(listResourceAliasesOptions)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourcecontrollerv2.ResourceAlias
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}

				// Test GetAll().
				pager, err = resourceControllerService.NewResourceAliasesPager(listResourceAliasesOptions)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allItems, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allItems).ToNot(BeNil())

				Expect(len(allItems)).To(Equal(len(allResults)))
				fmt.Fprintf(GinkgoWriter, "ListResourceAliases() returned a total of %d item(s) using ResourceAliasesPager.\n", len(allResults))
			})

			It(`11a - List Resource Aliases For Instance`, func() {
				shouldSkipTest()

				Expect(testInstanceGUID).ToNot(BeEmpty())

				listResourceAliasesForInstanceOptions := &resourcecontrollerv2.ListResourceAliasesForInstanceOptions{
					ID:    &testInstanceGUID,
					Limit: core.Int64Ptr(resultsPerPage),
				}

				results := []resourcecontrollerv2.ResourceAlias{}

				for {
					resourceAliasesList, response, err := resourceControllerService.ListResourceAliasesForInstance(listResourceAliasesForInstanceOptions)

					Expect(err).To(BeNil())
					Expect(response.StatusCode).To(Equal(200))
					Expect(resourceAliasesList).ToNot(BeNil())

					results = append(results, resourceAliasesList.Resources...)

					Expect(*resourceAliasesList.RowsCount).To(Equal(int64(1)))
					Expect(len(resourceAliasesList.Resources)).To(Equal(1))

					start, err := core.GetQueryParam(resourceAliasesList.NextURL, "start")
					Expect(err).To(BeNil())

					if start == nil {
						break
					}

					listResourceAliasesForInstanceOptions.Start = start
				}

				fmt.Fprintf(GinkgoWriter, "ListResourceAliasesForInstance() result:\n%s\n", common.ToJSON(results))
			})
			It(`ListResourceAliasesForInstance(listResourceAliasesForInstanceOptions *ListResourceAliasesForInstanceOptions) using ResourceAliasesForInstancePager`, func() {
				listResourceAliasesForInstanceOptions := &resourcecontrollerv2.ListResourceAliasesForInstanceOptions{
					ID: &testInstanceGUID,
				}

				// Test GetNext().
				pager, err := resourceControllerService.NewResourceAliasesForInstancePager(listResourceAliasesForInstanceOptions)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourcecontrollerv2.ResourceAlias
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}

				// Test GetAll().
				pager, err = resourceControllerService.NewResourceAliasesForInstancePager(listResourceAliasesForInstanceOptions)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allItems, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allItems).ToNot(BeNil())

				Expect(len(allItems)).To(Equal(len(allResults)))
				fmt.Fprintf(GinkgoWriter, "ListResourceAliasesForInstance() returned a total of %d item(s) using ResourceAliasesForInstancePager.\n", len(allResults))
			})
		})
	})

	Describe("Create, Retrieve, and Update Resource Binding", func() {
		It("12 - Create Resource Binding", func() {
			shouldSkipTest()

			Expect(testAliasCRN).ToNot(BeEmpty())

			target := "crn:v1:staging:public:bluemix:us-south:s/" + testSpaceGUID + "::cf-application:" + testAppGUID
			bindTargetCRN = "crn:v1:staging:public:cf:us-south:s/" + testSpaceGUID + "::cf-application:" + testAppGUID
			options := resourceControllerService.NewCreateResourceBindingOptions(testAliasCRN, target)
			options.SetName(bindingNames["name"])

			parameters := &resourcecontrollerv2.ResourceBindingPostParameters{}
			parameters.SetProperty("parameter1", "value1")
			parameters.SetProperty("parameter2", "value2")
			options.SetParameters(parameters)

			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test12-" + transactionID,
			}
			options.SetHeaders(headers)

			result, resp, err := resourceControllerService.CreateResourceBinding(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "CreateResourceBinding() result:\n%s\n", common.ToJSON(result))

			Expect(result.ID).NotTo(BeNil())
			Expect(result.GUID).NotTo(BeNil())
			Expect(result.CRN).NotTo(BeNil())
			Expect(*result.ID).To(Equal(*result.CRN))
			Expect(*result.Name).To(Equal(bindingNames["name"]))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGUID))
			Expect(*result.SourceCRN).To(Equal(testAliasCRN))
			Expect(*result.TargetCRN).To(Equal(bindTargetCRN))
			Expect(*result.State).To(Equal("active"))

			testBindingCRN = *result.ID
			testBindingGUID = *result.GUID
		})

		It("13 - Get A Resource Binding", func() {
			shouldSkipTest()

			Expect(testBindingGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewGetResourceBindingOptions(testBindingGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test13-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.GetResourceBinding(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "GetResourceBinding() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ID).To(Equal(testBindingCRN))
			Expect(*result.GUID).To(Equal(testBindingGUID))
			Expect(*result.CRN).To(Equal(testBindingCRN))
			Expect(*result.Name).To(Equal(bindingNames["name"]))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGUID))
			Expect(*result.SourceCRN).To(Equal(testAliasCRN))
			Expect(*result.TargetCRN).To(Equal(bindTargetCRN))
			Expect(*result.State).To(Equal("active"))
		})

		It("14 - Update A Resource Binding", func() {
			shouldSkipTest()

			Expect(testBindingGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewUpdateResourceBindingOptions(testBindingGUID, bindingNames["update"])
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test14-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.UpdateResourceBinding(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "UpdateResourceBinding() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ID).To(Equal(testBindingCRN))
			Expect(*result.Name).To(Equal(bindingNames["update"]))
			Expect(*result.State).To(Equal("active"))
		})

		Describe("List Resource Bindings", func() {
			It("15 - List Resource Bindings With No Filter", func() {
				shouldSkipTest()

				options := resourceControllerService.NewListResourceBindingsOptions()
				options.SetLimit(resultsPerPage)
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test15-" + transactionID,
				}
				options.SetHeaders(headers)

				results := []resourcecontrollerv2.ResourceBinding{}

				for {
					result, resp, err := resourceControllerService.ListResourceBindings(options)

					//should return one or more aliases
					Expect(err).To(BeNil())
					Expect(resp.StatusCode).To(Equal(200))
					Expect(result).ToNot(BeNil())

					Expect(*result.RowsCount).Should(BeNumerically(">=", int64(1)))
					Expect(*result.RowsCount).Should(BeNumerically("<=", int64(resultsPerPage)))
					Expect(len(result.Resources)).Should(BeNumerically(">=", 1))
					Expect(len(result.Resources)).Should(BeNumerically("<=", resultsPerPage))

					results = append(results, result.Resources...)

					start, err := core.GetQueryParam(result.NextURL, "start")
					Expect(err).To(BeNil())

					if start == nil {
						break
					}

					options.SetStart(*start)

				}

				fmt.Fprintf(GinkgoWriter, "ListResourceBindings() result:\n%s\n", common.ToJSON(results))
			})

			It("16 - List Resource Bindings With GUID Filter", func() {
				shouldSkipTest()

				Expect(testBindingGUID).ToNot(BeEmpty())

				options := resourceControllerService.NewListResourceBindingsOptions()
				options.SetGUID(testBindingGUID)
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test16-" + transactionID,
				}
				options.SetHeaders(headers)
				result, resp, err := resourceControllerService.ListResourceBindings(options)

				//should return list with only newly created binding
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

				fmt.Fprintf(GinkgoWriter, "ListResourceBindings() result:\n%s\n", common.ToJSON(result))

				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
				Expect(*result.Resources[0].ID).To(Equal(testBindingCRN))
				Expect(*result.Resources[0].Name).To(Equal(bindingNames["update"]))
				Expect(*result.Resources[0].ResourceGroupID).To(Equal(testResourceGroupGUID))
				Expect(*result.Resources[0].SourceCRN).To(Equal(testAliasCRN))
				Expect(*result.Resources[0].TargetCRN).To(Equal(bindTargetCRN))
				Expect(*result.Resources[0].State).To(Equal("active"))
			})

			It("17 - List Resource Bindings With Name Filter", func() {
				shouldSkipTest()

				options := resourceControllerService.NewListResourceBindingsOptions()
				options.SetName(bindingNames["update"])
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test17-" + transactionID,
				}
				options.SetHeaders(headers)
				result, resp, err := resourceControllerService.ListResourceBindings(options)

				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

				fmt.Fprintf(GinkgoWriter, "ListResourceBindings() result:\n%s\n", common.ToJSON(result))

				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
			})
			It(`ListResourceBindings(listResourceBindingsOptions *ListResourceBindingsOptions) using ResourceBindingsPager`, func() {
				listResourceBindingsOptions := &resourcecontrollerv2.ListResourceBindingsOptions{}

				// Test GetNext().
				pager, err := resourceControllerService.NewResourceBindingsPager(listResourceBindingsOptions)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourcecontrollerv2.ResourceBinding
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}

				// Test GetAll().
				pager, err = resourceControllerService.NewResourceBindingsPager(listResourceBindingsOptions)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allItems, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allItems).ToNot(BeNil())

				Expect(len(allItems)).To(Equal(len(allResults)))
				fmt.Fprintf(GinkgoWriter, "ListResourceBindings() returned a total of %d item(s) using ResourceBindingsPager.\n", len(allResults))
			})

			It(`17a - List Resource Bindings For Alias`, func() {
				shouldSkipTest()

				Expect(testAliasCRN).ToNot(BeEmpty())

				listResourceBindingsForAliasOptions := &resourcecontrollerv2.ListResourceBindingsForAliasOptions{
					ID:    &testAliasCRN,
					Limit: core.Int64Ptr(resultsPerPage),
				}

				results := []resourcecontrollerv2.ResourceBinding{}

				for {
					resourceBindingsList, response, err := resourceControllerService.ListResourceBindingsForAlias(listResourceBindingsForAliasOptions)

					Expect(err).To(BeNil())
					Expect(response.StatusCode).To(Equal(200))
					Expect(resourceBindingsList).ToNot(BeNil())

					results = append(results, resourceBindingsList.Resources...)

					Expect(*resourceBindingsList.RowsCount).To(Equal(int64(1)))
					Expect(len(resourceBindingsList.Resources)).To(Equal(1))

					start, err := core.GetQueryParam(resourceBindingsList.NextURL, "start")
					Expect(err).To(BeNil())

					if start == nil {
						break
					}

					listResourceBindingsForAliasOptions.Start = start
				}

				fmt.Fprintf(GinkgoWriter, "ListResourceBindingsForAlias() result:\n%s\n", common.ToJSON(results))
			})
			It(`ListResourceBindingsForAlias(listResourceBindingsForAliasOptions *ListResourceBindingsForAliasOptions) using ResourceBindingsForAliasPager`, func() {
				listResourceBindingsForAliasOptions := &resourcecontrollerv2.ListResourceBindingsForAliasOptions{
					ID: &testAliasCRN,
				}

				// Test GetNext().
				pager, err := resourceControllerService.NewResourceBindingsForAliasPager(listResourceBindingsForAliasOptions)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourcecontrollerv2.ResourceBinding
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}

				// Test GetAll().
				pager, err = resourceControllerService.NewResourceBindingsForAliasPager(listResourceBindingsForAliasOptions)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allItems, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allItems).ToNot(BeNil())

				Expect(len(allItems)).To(Equal(len(allResults)))
				fmt.Fprintf(GinkgoWriter, "ListResourceBindingsForAlias() returned a total of %d item(s) using ResourceBindingsForAliasPager.\n", len(allResults))
			})
		})
	})

	Describe("Create, Retrieve, and Update Resource Key With Instance Source", func() {
		It("18 - Create Resource Key For Instance", func() {
			shouldSkipTest()

			options := resourceControllerService.NewCreateResourceKeyOptions(keyNames["name"], testInstanceGUID)

			parameters := &resourcecontrollerv2.ResourceKeyPostParameters{}
			parameters.SetProperty("parameter1", "value1")
			parameters.SetProperty("parameter2", "value2")
			options.SetParameters(parameters)

			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test18-" + transactionID,
			}
			options.SetHeaders(headers)

			result, resp, err := resourceControllerService.CreateResourceKey(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "CreateResourceKey() result:\n%s\n", common.ToJSON(result))

			Expect(result.ID).NotTo(BeNil())
			Expect(result.GUID).NotTo(BeNil())
			Expect(result.CRN).NotTo(BeNil())
			Expect(*result.ID).To(Equal(*result.CRN))
			Expect(*result.Name).To(Equal(keyNames["name"]))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGUID))
			Expect(*result.SourceCRN).To(Equal(testInstanceCRN))
			Expect(*result.State).To(Equal("active"))

			testInstanceKeyCRN = *result.ID
			testInstanceKeyGUID = *result.GUID
		})

		It("19 - Get A Resource Key", func() {
			shouldSkipTest()

			options := resourceControllerService.NewGetResourceKeyOptions(testInstanceKeyGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test19-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.GetResourceKey(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "GetResourceKey() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ID).To(Equal(testInstanceKeyCRN))
			Expect(*result.OnetimeCredentials).ToNot(BeNil())
			Expect(*result.GUID).To(Equal(testInstanceKeyGUID))
			Expect(*result.CRN).To(Equal(testInstanceKeyCRN))
			Expect(*result.Name).To(Equal(keyNames["name"]))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGUID))
			Expect(*result.SourceCRN).To(Equal(testInstanceCRN))
			Expect(*result.State).To(Equal("active"))
		})

		It("20 - Update A Resource Key", func() {
			shouldSkipTest()

			options := resourceControllerService.NewUpdateResourceKeyOptions(testInstanceKeyGUID, keyNames["update"])
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test20-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.UpdateResourceKey(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "UpdateResourceKey() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ID).To(Equal(testInstanceKeyCRN))
			Expect(*result.Name).To(Equal(keyNames["update"]))
			Expect(*result.State).To(Equal("active"))
		})

		Describe("List Resource Keys", func() {
			It("21 - List Resource Keys With No Filter", func() {
				shouldSkipTest()

				options := resourceControllerService.NewListResourceKeysOptions()
				options.SetLimit(resultsPerPage)
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test21-" + transactionID,
				}
				options.SetHeaders(headers)

				results := []resourcecontrollerv2.ResourceKey{}

				for {
					result, resp, err := resourceControllerService.ListResourceKeys(options)

					//should return one or more aliases
					Expect(err).To(BeNil())
					Expect(resp.StatusCode).To(Equal(200))
					Expect(result).ToNot(BeNil())
					results = append(results, result.Resources...)

					start, err := core.GetQueryParam(result.NextURL, "start")
					Expect(err).To(BeNil())

					if start == nil {
						break
					}

					options.SetStart(*start)

				}
				Expect(len(results)).Should(BeNumerically(">=", 1))

				fmt.Fprintf(GinkgoWriter, "ListResourceKeys() result:\n%s\n", common.ToJSON(results))
			})

			It("22 - List Resource Keys With GUID Filter", func() {
				shouldSkipTest()

				options := resourceControllerService.NewListResourceKeysOptions()
				options.SetGUID(testInstanceKeyGUID)
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test22-" + transactionID,
				}
				options.SetHeaders(headers)
				result, resp, err := resourceControllerService.ListResourceKeys(options)

				//should return list with only newly created key
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

				fmt.Fprintf(GinkgoWriter, "ListResourceKeys() result:\n%s\n", common.ToJSON(result))

				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
				Expect(*result.Resources[0].ID).To(Equal(testInstanceKeyCRN))
				Expect(*result.Resources[0].Name).To(Equal(keyNames["update"]))
				Expect(*result.Resources[0].ResourceGroupID).To(Equal(testResourceGroupGUID))
				Expect(*result.Resources[0].SourceCRN).To(Equal(testInstanceCRN))
				Expect(*result.Resources[0].State).To(Equal("active"))
			})

			It("23 - List Resource Keys With Name Filter", func() {
				shouldSkipTest()

				options := resourceControllerService.NewListResourceKeysOptions()
				options.SetName(keyNames["update"])
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test23-" + transactionID,
				}
				options.SetHeaders(headers)
				result, resp, err := resourceControllerService.ListResourceKeys(options)

				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

				fmt.Fprintf(GinkgoWriter, "ListResourceKeys() result:\n%s\n", common.ToJSON(result))

				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
			})
			It(`ListResourceKeys(listResourceKeysOptions *ListResourceKeysOptions) using ResourceKeysPager`, func() {
				listResourceKeysOptions := &resourcecontrollerv2.ListResourceKeysOptions{}

				// Test GetNext().
				pager, err := resourceControllerService.NewResourceKeysPager(listResourceKeysOptions)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourcecontrollerv2.ResourceKey
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}

				// Test GetAll().
				pager, err = resourceControllerService.NewResourceKeysPager(listResourceKeysOptions)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allItems, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allItems).ToNot(BeNil())

				Expect(len(allItems)).To(Equal(len(allResults)))
				fmt.Fprintf(GinkgoWriter, "ListResourceKeys() returned a total of %d item(s) using ResourceKeysPager.\n", len(allResults))
			})

			It(`23a - List Resource Keys For Instance`, func() {
				shouldSkipTest()

				Expect(testInstanceGUID).ToNot(BeEmpty())

				listResourceKeysForInstanceOptions := &resourcecontrollerv2.ListResourceKeysForInstanceOptions{
					ID:    &testInstanceGUID,
					Limit: core.Int64Ptr(resultsPerPage),
				}

				results := []resourcecontrollerv2.ResourceKey{}

				for {
					resourceKeysList, response, err := resourceControllerService.ListResourceKeysForInstance(listResourceKeysForInstanceOptions)

					Expect(err).To(BeNil())
					Expect(response.StatusCode).To(Equal(200))
					Expect(resourceKeysList).ToNot(BeNil())

					results = append(results, resourceKeysList.Resources...)

					Expect(*resourceKeysList.RowsCount).To(Equal(int64(1)))
					Expect(len(resourceKeysList.Resources)).To(Equal(1))

					start, err := core.GetQueryParam(resourceKeysList.NextURL, "start")
					Expect(err).To(BeNil())

					if start == nil {
						break
					}

					listResourceKeysForInstanceOptions.Start = start
				}

				fmt.Fprintf(GinkgoWriter, "ListResourceKeysForInstance() result:\n%s\n", common.ToJSON(results))
			})
			It(`ListResourceKeysForInstance(listResourceKeysForInstanceOptions *ListResourceKeysForInstanceOptions) using ResourceKeysForInstancePager`, func() {
				listResourceKeysForInstanceOptions := &resourcecontrollerv2.ListResourceKeysForInstanceOptions{
					ID: &testInstanceGUID,
				}

				// Test GetNext().
				pager, err := resourceControllerService.NewResourceKeysForInstancePager(listResourceKeysForInstanceOptions)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []resourcecontrollerv2.ResourceKey
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}

				// Test GetAll().
				pager, err = resourceControllerService.NewResourceKeysForInstancePager(listResourceKeysForInstanceOptions)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allItems, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allItems).ToNot(BeNil())

				Expect(len(allItems)).To(Equal(len(allResults)))
				fmt.Fprintf(GinkgoWriter, "ListResourceKeysForInstance() returned a total of %d item(s) using ResourceKeysForInstancePager.\n", len(allResults))
			})
		})
	})

	Describe("Create, Retrieve, and Update Resource Key With Alias Source", func() {
		It("24 - Create Resource Key For Alias", func() {
			shouldSkipTest()

			options := resourceControllerService.NewCreateResourceKeyOptions(keyNames["name2"], testAliasCRN)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test24-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.CreateResourceKey(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "CreateResourceKey() result:\n%s\n", common.ToJSON(result))

			Expect(result.ID).NotTo(BeNil())
			Expect(result.GUID).NotTo(BeNil())
			Expect(result.CRN).NotTo(BeNil())
			Expect(*result.ID).To(Equal(*result.CRN))
			Expect(*result.Name).To(Equal(keyNames["name2"]))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGUID))
			Expect(*result.SourceCRN).To(Equal(testAliasCRN))
			Expect(*result.State).To(Equal("active"))

			testAliasKeyCRN = *result.ID
			testAliasKeyGUID = *result.GUID
		})

		It("25 - Get A Resource Key", func() {
			shouldSkipTest()

			Expect(testAliasKeyGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewGetResourceKeyOptions(testAliasKeyGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test25-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.GetResourceKey(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "GetResourceKey() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ID).To(Equal(testAliasKeyCRN))
			Expect(*result.GUID).To(Equal(testAliasKeyGUID))
			Expect(*result.CRN).To(Equal(testAliasKeyCRN))
			Expect(*result.Name).To(Equal(keyNames["name2"]))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGUID))
			Expect(*result.SourceCRN).To(Equal(testAliasCRN))
			Expect(*result.State).To(Equal("active"))
		})

		It("26 - Update A Resource Key", func() {
			shouldSkipTest()

			Expect(testAliasKeyGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewUpdateResourceKeyOptions(testAliasKeyGUID, keyNames["update2"])
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test26-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.UpdateResourceKey(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "UpdateResourceKey() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ID).To(Equal(testAliasKeyCRN))
			Expect(*result.Name).To(Equal(keyNames["update2"]))
			Expect(*result.State).To(Equal("active"))
		})

		Describe("List Resource Keys", func() {
			It("27 - List Resource Keys With No Filter", func() {
				shouldSkipTest()

				options := resourceControllerService.NewListResourceKeysOptions()
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test27-" + transactionID,
				}
				options.SetHeaders(headers)
				result, resp, err := resourceControllerService.ListResourceKeys(options)

				//should return two or more keys
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

				fmt.Fprintf(GinkgoWriter, "ListResourceKeys() result:\n%s\n", common.ToJSON(result))

				Expect(*result.RowsCount).Should(BeNumerically(">=", int64(2)))
				Expect(len(result.Resources)).Should(BeNumerically(">=", 2))
			})

			It("28 - List Resource Keys With GUID Filter", func() {
				shouldSkipTest()

				Expect(testAliasKeyGUID).ToNot(BeEmpty())

				options := resourceControllerService.NewListResourceKeysOptions()
				options.SetGUID(testAliasKeyGUID)
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test28-" + transactionID,
				}
				options.SetHeaders(headers)
				result, resp, err := resourceControllerService.ListResourceKeys(options)

				//should return list with only newly created key
				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

				fmt.Fprintf(GinkgoWriter, "ListResourceKeys() result:\n%s\n", common.ToJSON(result))

				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
				Expect(*result.Resources[0].ID).To(Equal(testAliasKeyCRN))
				Expect(*result.Resources[0].Name).To(Equal(keyNames["update2"]))
				Expect(*result.Resources[0].ResourceGroupID).To(Equal(testResourceGroupGUID))
				Expect(*result.Resources[0].SourceCRN).To(Equal(testAliasCRN))
				Expect(*result.Resources[0].State).To(Equal("active"))
			})

			It("29 - List Resource Keys With Name Filter", func() {
				shouldSkipTest()

				options := resourceControllerService.NewListResourceKeysOptions()
				options.SetName(keyNames["update2"])
				headers := map[string]string{
					"Transaction-ID": "rc-sdk-go-test29-" + transactionID,
				}
				options.SetHeaders(headers)
				result, resp, err := resourceControllerService.ListResourceKeys(options)

				Expect(err).To(BeNil())
				Expect(resp.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())

				fmt.Fprintf(GinkgoWriter, "ListResourceKeys() result:\n%s\n", common.ToJSON(result))

				Expect(*result.RowsCount).To(Equal(int64(1)))
				Expect(result.Resources).Should(HaveLen(1))
			})
		})
	})

	Describe("Delete All Resources", func() {
		It("30 - Delete A Resource Alias With Dependencies - Fail", func() {
			shouldSkipTest()

			options := resourceControllerService.NewDeleteResourceAliasOptions(testAliasCRN)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test30-" + transactionID,
			}
			options.SetHeaders(headers)
			resp, err := resourceControllerService.DeleteResourceAlias(options)

			Expect(resp.StatusCode).To(Equal(400))
			Expect(err).NotTo(BeNil())
		})

		It("31 - Delete A Resource Instance With Dependencies - Fail", func() {
			shouldSkipTest()

			options := resourceControllerService.NewDeleteResourceInstanceOptions(testInstanceGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test31-" + transactionID,
			}
			options.SetHeaders(headers)
			resp, err := resourceControllerService.DeleteResourceInstance(options)

			Expect(resp.StatusCode).To(Equal(400))
			Expect(err).NotTo(BeNil())
		})

		It("32 - Delete A Resource Binding", func() {
			shouldSkipTest()

			Expect(testBindingGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewDeleteResourceBindingOptions(testBindingGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test32-" + transactionID,
			}
			options.SetHeaders(headers)
			resp, err := resourceControllerService.DeleteResourceBinding(options)

			Expect(resp.StatusCode).To(Equal(204))
			Expect(err).To(BeNil())
		})

		It("33 - Verify Resource Binding Was Deleted", func() {
			shouldSkipTest()

			Expect(testBindingGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewGetResourceBindingOptions(testBindingGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test33-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.GetResourceBinding(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testBindingCRN))
			Expect(*result.State).To(Equal("removed"))
		})

		It("34 - Delete Resource Keys", func() {
			shouldSkipTest()

			Expect(testInstanceKeyGUID).ToNot(BeEmpty())

			options1 := resourceControllerService.NewDeleteResourceKeyOptions(testInstanceKeyGUID)
			headers1 := map[string]string{
				"Transaction-ID": "rc-sdk-go-test34-" + transactionID,
			}
			options1 = options1.SetHeaders(headers1)
			resp1, err1 := resourceControllerService.DeleteResourceKey(options1)

			Expect(resp1.StatusCode).To(Equal(204))
			Expect(err1).To(BeNil())

			options2 := resourceControllerService.NewDeleteResourceKeyOptions(testAliasKeyGUID)
			headers2 := map[string]string{
				"Transaction-ID": "rc-sdk-go-test34-" + transactionID,
			}
			options2 = options2.SetHeaders(headers2)
			resp2, err2 := resourceControllerService.DeleteResourceKey(options2)

			Expect(resp2.StatusCode).To(Equal(204))
			Expect(err2).To(BeNil())
		})

		It("35 - Verify Resource Keys Were Deleted", func() {
			shouldSkipTest()

			Expect(testInstanceKeyGUID).ToNot(BeEmpty())

			options1 := resourceControllerService.NewGetResourceKeyOptions(testInstanceKeyGUID)
			headers1 := map[string]string{
				"Transaction-ID": "rc-sdk-go-test35-" + transactionID,
			}
			options1 = options1.SetHeaders(headers1)
			result1, resp1, err1 := resourceControllerService.GetResourceKey(options1)

			Expect(err1).To(BeNil())
			Expect(resp1.StatusCode).To(Equal(200))
			Expect(*result1.ID).To(Equal(testInstanceKeyCRN))
			Expect(*result1.State).To(Equal("removed"))

			Expect(testAliasKeyGUID).ToNot(BeEmpty())

			options2 := resourceControllerService.NewGetResourceKeyOptions(testAliasKeyGUID)
			headers2 := map[string]string{
				"Transaction-ID": "rc-sdk-go-test35-" + transactionID,
			}
			options2 = options2.SetHeaders(headers2)
			result2, resp2, err2 := resourceControllerService.GetResourceKey(options2)

			Expect(err2).To(BeNil())
			Expect(resp2.StatusCode).To(Equal(200))
			Expect(*result2.ID).To(Equal(testAliasKeyCRN))
			Expect(*result2.State).To(Equal("removed"))
		})

		It("36 - Delete A Resource Alias", func() {
			shouldSkipTest()

			Expect(testAliasCRN).ToNot(BeEmpty())

			options := resourceControllerService.NewDeleteResourceAliasOptions(testAliasCRN)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test36-" + transactionID,
			}
			options.SetHeaders(headers)
			resp, err := resourceControllerService.DeleteResourceAlias(options)

			Expect(resp.StatusCode).To(Equal(204))
			Expect(err).To(BeNil())
		})

		It("37 - Verify Resource Alias Was Deleted", func() {
			shouldSkipTest()

			Expect(testAliasCRN).ToNot(BeEmpty())

			options := resourceControllerService.NewGetResourceAliasOptions(testAliasCRN)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test37-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.GetResourceAlias(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "GetResourceAlias() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ID).To(Equal(testAliasCRN))
			Expect(*result.State).To(Equal("removed"))
		})
	})

	Describe("Locking and Unlocking Resource Instance", func() {
		It("38 - Lock A Resource Instance", func() {
			shouldSkipTest()

			Expect(testInstanceGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewLockResourceInstanceOptions(testInstanceGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test38-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.LockResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "LockResourceInstance() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ID).To(Equal(testInstanceCRN))
			Expect(*result.Locked).To(BeTrue())
			Expect(*result.LastOperation.Type).To(Equal("lock"))
			Expect(*result.LastOperation.Async).Should(BeFalse())
			Expect(*result.LastOperation.State).To(Equal("succeeded"))
		})

		It("39 - Update A Locked Resource Instance - Fail", func() {
			shouldSkipTest()

			Expect(testInstanceGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewUpdateResourceInstanceOptions(testInstanceGUID)
			options.SetName(testLockedInstanceNameUpdate)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test39-" + transactionID,
			}
			options.SetHeaders(headers)
			_, resp, err := resourceControllerService.UpdateResourceInstance(options)

			Expect(err).NotTo(BeNil())
			Expect(resp.StatusCode).To(Equal(422))
		})

		It("40 - Delete A Locked Resource Instance - Fail", func() {
			shouldSkipTest()

			Expect(testInstanceGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewDeleteResourceInstanceOptions(testInstanceGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test40-" + transactionID,
			}
			options.SetHeaders(headers)
			resp, err := resourceControllerService.DeleteResourceInstance(options)

			Expect(err).NotTo(BeNil())
			Expect(resp.StatusCode).To(Equal(422))
		})

		It("41 - Unlock A Resource Instance", func() {
			shouldSkipTest()

			Expect(testInstanceGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewUnlockResourceInstanceOptions(testInstanceGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test41-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.UnlockResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "UnlockResourceInstance() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ID).To(Equal(testInstanceCRN))
			Expect(*result.Locked).To(BeFalse())
			Expect(*result.LastOperation.Type).To(Equal("unlock"))
			Expect(*result.LastOperation.Async).Should(BeFalse())
			Expect(*result.LastOperation.State).To(Equal("succeeded"))
		})
	})

	Describe("Delete Resource Instance", func() {
		It("42 - Delete A Resource Instance", func() {
			shouldSkipTest()

			Expect(testInstanceGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewDeleteResourceInstanceOptions(testInstanceGUID)
			options.SetRecursive(false)

			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test42-" + transactionID,
			}
			options.SetHeaders(headers)
			resp, err := resourceControllerService.DeleteResourceInstance(options)

			Expect(resp.StatusCode).To(Equal(204))
			Expect(err).To(BeNil())
		})

		It("43 - Verify Resource Instance Was Deleted", func() {
			shouldSkipTest()

			Expect(testInstanceGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewGetResourceInstanceOptions(testInstanceGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test43-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.GetResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "GetResourceInstance() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ID).To(Equal(testInstanceCRN))
			Expect(*result.State).To(Equal("removed"))
			Expect(*result.LastOperation.Type).To(Equal("delete"))
			Expect(*result.LastOperation.Async).Should(BeFalse())
			Expect(*result.LastOperation.State).To(Equal("succeeded"))
		})
	})

	Describe("Resource Reclamation", func() {
		It("44 - Create Resource Instance For Reclamation Enabled Plan", func() {
			shouldSkipTest()

			options := resourceControllerService.NewCreateResourceInstanceOptions(
				testReclaimInstanceName,
				testRegionID2,
				testResourceGroupGUID,
				testPlanID2,
			)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test44-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.CreateResourceInstance(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "CreateResourceInstance() result:\n%s\n", common.ToJSON(result))

			Expect(result.ID).NotTo(BeNil())
			Expect(result.GUID).NotTo(BeNil())
			Expect(result.CRN).NotTo(BeNil())
			Expect(*result.ID).To(Equal(*result.CRN))
			Expect(*result.Name).To(Equal(testReclaimInstanceName))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGUID))
			Expect(*result.ResourcePlanID).To(Equal(testPlanID2))
			Expect(*result.State).To(Equal("active"))
			Expect(*result.Locked).Should(BeFalse())
			Expect(*result.LastOperation.Type).To(Equal("create"))
			Expect(*result.LastOperation.Async).Should(BeFalse())
			Expect(*result.LastOperation.State).To(Equal("succeeded"))

			testReclaimInstanceCRN = *result.ID
			testReclaimInstanceGUID = *result.GUID
		})

		It("45 - Schedule The Resource Instance For Reclamation", func() {
			shouldSkipTest()

			Expect(testReclaimInstanceGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewDeleteResourceInstanceOptions(testReclaimInstanceGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test45-" + transactionID,
			}
			options.SetHeaders(headers)
			resp, err := resourceControllerService.DeleteResourceInstance(options)

			Expect(resp.StatusCode).To(Equal(204))
			Expect(err).To(BeNil())

			//wait for reclamation object to be created
			time.Sleep(20 * time.Second)
		})

		// Commented because redis timeouts cause intermittent failure

		// It("46 - Verify The Resource Instance Is Pending Reclamation", func() {
		// 	shouldSkipTest()

		// 	options := resourceControllerService.NewGetResourceInstanceOptions(testReclaimInstanceGUID)
		// 	headers := map[string]string{
		// 		"Transaction-ID": "rc-sdk-go-test46-" + transactionID,
		// 	}
		// 	options.SetHeaders(headers)
		// 	result, resp, err := resourceControllerService.GetResourceInstance(options)

		// 	Expect(err).To(BeNil())
		// 	Expect(resp.StatusCode).To(Equal(200))
		// 	Expect(*result.ID).To(Equal(testReclaimInstanceCRN))
		// 	Expect(*result.State).To(Equal("pending_reclamation"))
		// 	Expect(*result.LastOperation.Type).To(Equal("reclamation"))
		// 	Expect(*result.LastOperation.Async).Should(BeFalse())
		// 	Expect(*result.LastOperation.State).To(Equal("succeeded"))
		// })

		It("47 - List Reclamations For Account ID", func() {
			shouldSkipTest()

			Expect(testReclaimInstanceGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewListReclamationsOptions()
			// options.SetAccountID(testAccountID)
			options.SetResourceInstanceID(testReclaimInstanceGUID) //checking reclamations with instance guid to make it more reliable
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test47-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.ListReclamations(options)

			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "ListReclamations() result:\n%s\n", common.ToJSON(result))

			Expect(len(result.Resources)).Should(BeNumerically(">=", 1))
			Expect(err).To(BeNil())

			foundReclamation := false
			for _, res := range result.Resources {
				if *res.ResourceInstanceID == testReclaimInstanceGUID {
					Expect(*res.ResourceInstanceID).To(Equal(testReclaimInstanceGUID))
					Expect(*res.AccountID).To(Equal(testAccountID))
					Expect(*res.ResourceGroupID).To(Equal(testResourceGroupGUID))
					Expect(*res.State).To(Equal("SCHEDULED"))

					foundReclamation = true
					testReclamationID1 = *res.ID
				}
			}

			Expect(foundReclamation).To(BeTrue())
		})

		It("48 - Restore A Resource Instance", func() {
			shouldSkipTest()

			Expect(testReclamationID1).ToNot(BeEmpty())

			options := resourceControllerService.NewRunReclamationActionOptions(testReclamationID1, "restore")
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test48-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.RunReclamationAction(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "RunReclamationAction() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ResourceInstanceID).To(Equal(testReclaimInstanceGUID))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGUID))
			Expect(*result.State).To(Equal("RESTORING"))

			//wait for reclamation object to be created
			time.Sleep(20 * time.Second)
		})

		// Commented because redis timeouts cause intermittent failure

		// It("49 - Verify The Resource Instance Is Restored", func() {
		// 	shouldSkipTest()

		// 	options := resourceControllerService.NewGetResourceInstanceOptions(testReclaimInstanceGUID)
		// 	headers := map[string]string{
		// 		"Transaction-ID": "rc-sdk-go-test49-" + transactionID,
		// 	}
		// 	options.SetHeaders(headers)
		// 	result, resp, err := resourceControllerService.GetResourceInstance(options)

		// 	Expect(err).To(BeNil())
		// 	Expect(resp.StatusCode).To(Equal(200))
		// 	Expect(*result.ID).To(Equal(testReclaimInstanceCRN))
		// 	Expect(*result.State).To(Equal("active"))
		// 	Expect(*result.LastOperation.Type).To(Equal("reclamation"))
		// 	Expect(*result.LastOperation.Async).Should(BeFalse())
		// 	Expect(*result.LastOperation.State).To(Equal("succeeded"))
		// })

		It("50 - Schedule The Resource Instance For Reclamation 2", func() {
			shouldSkipTest()

			Expect(testReclaimInstanceGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewDeleteResourceInstanceOptions(testReclaimInstanceGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test50-" + transactionID,
			}
			options.SetHeaders(headers)
			resp, err := resourceControllerService.DeleteResourceInstance(options)

			Expect(resp.StatusCode).To(Equal(204))
			Expect(err).To(BeNil())

			//wait for reclamation object to be created
			time.Sleep(20 * time.Second)
		})

		It("51 - List Reclamations For Account and Resource Instance ID", func() {
			shouldSkipTest()

			Expect(testAccountID).ToNot(BeEmpty())
			Expect(testReclaimInstanceGUID).ToNot(BeEmpty())

			options := resourceControllerService.NewListReclamationsOptions()
			options.SetAccountID(testAccountID)
			options.SetResourceInstanceID(testReclaimInstanceGUID)
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test51-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.ListReclamations(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "ListReclamations() result:\n%s\n", common.ToJSON(result))

			Expect(result.Resources).Should(HaveLen(1))
			Expect(*result.Resources[0].ResourceInstanceID).To(Equal(testReclaimInstanceGUID))
			Expect(*result.Resources[0].AccountID).To(Equal(testAccountID))
			Expect(*result.Resources[0].ResourceGroupID).To(Equal(testResourceGroupGUID))
			Expect(*result.Resources[0].State).To(Equal("SCHEDULED"))

			testReclamationID2 = *result.Resources[0].ID
		})

		It("52 - Reclaim A Resource Instance", func() {
			shouldSkipTest()

			Expect(testReclamationID2).ToNot(BeEmpty())

			options := resourceControllerService.NewRunReclamationActionOptions(testReclamationID2, "reclaim")
			headers := map[string]string{
				"Transaction-ID": "rc-sdk-go-test52-" + transactionID,
			}
			options.SetHeaders(headers)
			result, resp, err := resourceControllerService.RunReclamationAction(options)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "RunReclamationAction() result:\n%s\n", common.ToJSON(result))

			Expect(*result.ResourceInstanceID).To(Equal(testReclaimInstanceGUID))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ResourceGroupID).To(Equal(testResourceGroupGUID))
			Expect(*result.State).To(Equal("RECLAIMING"))

			//wait for reclamation object to be created
			time.Sleep(20 * time.Second)
		})

		Describe(`CancelLastopResourceInstance - Cancel the in progress last operation of the resource instance`, func() {
			BeforeEach(func() {
				shouldSkipTest()
			})
			It(`53 - CancelLastopResourceInstance(cancelLastopResourceInstanceOptions *CancelLastopResourceInstanceOptions)`, func() {
				Expect(testInstanceGUID).ToNot(BeEmpty())
				cancelLastopResourceInstanceOptions := &resourcecontrollerv2.CancelLastopResourceInstanceOptions{
					ID: &testInstanceGUID,
				}

				resourceInstance, response, err := resourceControllerService.CancelLastopResourceInstance(cancelLastopResourceInstanceOptions)
				// Expect(err).To(BeNil())
				// Expect(response.StatusCode).To(Equal(200))
				// Expect(resourceInstance).ToNot(BeNil())
				Expect(err.Error()).To(Equal("The instance is not cancelable."))
				Expect(response.StatusCode).To(Equal(422))
				Expect(resourceInstance).To(BeNil())
			})
		})

		// Commented because redis timeouts cause intermittent failure

		// It("53 - Verify The Resource Instance Is Reclaimed", func() {
		// 	shouldSkipTest()

		// 	options := resourceControllerService.NewGetResourceInstanceOptions(testReclaimInstanceGUID)
		// 	headers := map[string]string{
		// 		"Transaction-ID": "rc-sdk-go-test53-" + transactionID,
		// 	}
		// 	options.SetHeaders(headers)
		// 	result, resp, err := resourceControllerService.GetResourceInstance(options)

		// 	//printing info for debugging
		// 	fmt.Fprintln(GinkgoWriter, "\nDEBUGGING - testReclaimInstanceGUID: %s\n", testReclaimInstanceGUID)
		// 	fmt.Fprintln(GinkgoWriter, "\nDEBUGGING - Transaction-ID: rc-sdk-go-test53-%s\n", transactionID)

		// 	Expect(err).To(BeNil())
		// 	Expect(resp.StatusCode).To(Equal(200))
		// 	Expect(*result.ID).To(Equal(testReclaimInstanceCRN))
		// 	Expect(*result.State).To(Equal("removed"))
		// 	Expect(*result.LastOperation.Type).To(Equal("reclamation"))
		// 	Expect(*result.LastOperation.Async).Should(BeFalse())
		// 	Expect(*result.LastOperation.State).To(Equal("succeeded"))
		// })
	})
})

// clean up resources
var _ = AfterSuite(func() {
	if !configLoaded {
		return
	}

	fmt.Fprintln(GinkgoWriter, "After tests: cleaning up test resources...")
	cleanupResources()
	if testReclaimInstanceGUID != "" {
		cleanupReclamationInstance()
	} else {
		fmt.Fprintln(GinkgoWriter, "Reclamation instance was not created. No cleanup needed.")
	}
	cleanupByName()
})

func cleanupByName() {
	fmt.Fprintln(GinkgoWriter, "Begin cleanup by name")

	//Resource Key
	for _, name := range keyNames {
		listKeyOptions := resourceControllerService.NewListResourceKeysOptions()
		listKeyOptions = listKeyOptions.SetName(name)
		listKeyHeaders := map[string]string{
			"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
		}
		listKeyOptions = listKeyOptions.SetHeaders(listKeyHeaders)
		keyResult, _, keyListErr := resourceControllerService.ListResourceKeys(listKeyOptions)

		if keyListErr != nil {
			fmt.Fprintln(GinkgoWriter, "Failed to retrieve key with name ", name, " for cleanup.")
			return
		}

		if len(keyResult.Resources) > 0 {
			keyResources := &keyResult.Resources
			for _, res := range *keyResources {
				keyGUID := *res.GUID

				deleteKeyOptions := resourceControllerService.NewDeleteResourceKeyOptions(keyGUID)
				deleteKeyHeaders := map[string]string{
					"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
				}
				deleteKeyOptions = deleteKeyOptions.SetHeaders(deleteKeyHeaders)
				keyDelResp, keyDelErr := resourceControllerService.DeleteResourceKey(deleteKeyOptions)
				if keyDelResp.StatusCode == 204 {
					fmt.Fprintln(GinkgoWriter, "Successful cleanup of key ", keyGUID)
				} else if keyDelResp.StatusCode == 410 {
					fmt.Fprintln(GinkgoWriter, "Key ", keyGUID, " was already deleted by the tests.")
				} else {
					fmt.Fprintln(GinkgoWriter, "Failed to cleanup key ", keyGUID, ". Error: ", keyDelErr.Error())
				}
			}
		} else {
			fmt.Fprintln(GinkgoWriter, "No keys found with name ", name, " for cleanup.")
		}
	}

	//Resource Instance
	for _, name := range instanceNames {
		listInstanceOptions := resourceControllerService.NewListResourceInstancesOptions()
		listInstanceOptions = listInstanceOptions.SetName(name)
		listInstanceHeaders := map[string]string{
			"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
		}
		listInstanceOptions = listInstanceOptions.SetHeaders(listInstanceHeaders)
		instanceResult, _, instanceListErr := resourceControllerService.ListResourceInstances(listInstanceOptions)

		if instanceListErr != nil {
			fmt.Fprintln(GinkgoWriter, "Failed to retrieve instance with name ", name, " for cleanup.")
			return
		}

		if len(instanceResult.Resources) > 0 {
			instanceResources := &instanceResult.Resources
			for _, res := range *instanceResources {
				instanceGUID := *res.GUID

				if *res.State == "active" && *res.Locked {
					instanceUnlockOptions := resourceControllerService.NewUnlockResourceInstanceOptions(instanceGUID)
					instanceUnlockHeaders := map[string]string{
						"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
					}
					instanceUnlockOptions = instanceUnlockOptions.SetHeaders(instanceUnlockHeaders)
					_, _, instanceUnlockErr := resourceControllerService.UnlockResourceInstance(instanceUnlockOptions)
					if instanceUnlockErr != nil {
						fmt.Fprintln(GinkgoWriter, "Failed to unlock instance ", instanceGUID, " for cleanup. Error: ", instanceUnlockErr.Error())
						return
					}
				}

				deleteInstanceOptions := resourceControllerService.NewDeleteResourceInstanceOptions(instanceGUID)
				deleteInstanceHeaders := map[string]string{
					"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
				}
				deleteInstanceOptions = deleteInstanceOptions.SetHeaders(deleteInstanceHeaders)
				instanceDelResp, instanceDelErr := resourceControllerService.DeleteResourceInstance(deleteInstanceOptions)
				if instanceDelResp.StatusCode == 204 {
					fmt.Fprintln(GinkgoWriter, "Successful cleanup of instance ", instanceGUID)
				} else if instanceDelResp.StatusCode == 410 {
					fmt.Fprintln(GinkgoWriter, "Instance ", instanceGUID, " was already deleted by the tests.")
				} else {
					fmt.Fprintln(GinkgoWriter, "Failed to cleanup instance ", instanceGUID, ". Error: ", instanceDelErr.Error())
				}
			}
		} else {
			fmt.Fprintln(GinkgoWriter, "No instances found with name ", name, " for cleanup.")
		}
	}

	//Resource Binding
	for _, name := range bindingNames {
		listBindingOptions := resourceControllerService.NewListResourceBindingsOptions()
		listBindingOptions = listBindingOptions.SetName(name)
		listBindingHeaders := map[string]string{
			"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
		}
		listBindingOptions = listBindingOptions.SetHeaders(listBindingHeaders)
		bindingResult, _, bindingListErr := resourceControllerService.ListResourceBindings(listBindingOptions)

		if bindingListErr != nil {
			fmt.Fprintln(GinkgoWriter, "Failed to retrieve binding with name ", name, " for cleanup.")
			return
		}

		if len(bindingResult.Resources) > 0 {
			bindingResources := &bindingResult.Resources
			for _, res := range *bindingResources {
				bindingGUID := *res.GUID

				deleteBindingOptions := resourceControllerService.NewDeleteResourceBindingOptions(bindingGUID)
				deleteBindingHeaders := map[string]string{
					"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
				}
				deleteBindingOptions = deleteBindingOptions.SetHeaders(deleteBindingHeaders)
				bindingDelResp, bindingDelErr := resourceControllerService.DeleteResourceBinding(deleteBindingOptions)
				if bindingDelResp.StatusCode == 204 {
					fmt.Fprintln(GinkgoWriter, "Successful cleanup of binding ", bindingGUID)
				} else if bindingDelResp.StatusCode == 410 {
					fmt.Fprintln(GinkgoWriter, "Binding ", bindingGUID, " was already deleted by the tests.")
				} else {
					fmt.Fprintln(GinkgoWriter, "Failed to cleanup binding ", bindingGUID, ". Error: ", bindingDelErr.Error())
				}
			}
		} else {
			fmt.Fprintln(GinkgoWriter, "No bindings found with name ", name, " for cleanup.")
		}
	}

	//Resource Alias
	for _, name := range aliasNames {
		listAliasOptions := resourceControllerService.NewListResourceAliasesOptions()
		listAliasOptions = listAliasOptions.SetName(name)
		listAliasHeaders := map[string]string{
			"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
		}
		listAliasOptions = listAliasOptions.SetHeaders(listAliasHeaders)
		aliasResult, _, aliasListErr := resourceControllerService.ListResourceAliases(listAliasOptions)

		if aliasListErr != nil {
			fmt.Fprintln(GinkgoWriter, "Failed to retrieve alias with name ", name, " for cleanup.")
			return
		}

		if len(aliasResult.Resources) > 0 {
			aliasResources := &aliasResult.Resources
			for _, res := range *aliasResources {
				aliasGUID := *res.GUID

				deleteAliasOptions := resourceControllerService.NewDeleteResourceAliasOptions(aliasGUID)
				deleteAliasHeaders := map[string]string{
					"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
				}
				deleteAliasOptions = deleteAliasOptions.SetHeaders(deleteAliasHeaders)
				aliasDelResp, aliasDelErr := resourceControllerService.DeleteResourceAlias(deleteAliasOptions)
				if aliasDelResp.StatusCode == 204 {
					fmt.Fprintln(GinkgoWriter, "Successful cleanup of alias ", aliasGUID)
				} else if aliasDelResp.StatusCode == 410 {
					fmt.Fprintln(GinkgoWriter, "Alias ", aliasGUID, " was already deleted by the tests.")
				} else {
					fmt.Fprintln(GinkgoWriter, "Failed to cleanup alias ", aliasGUID, ". Error: ", aliasDelErr.Error())
				}
			}
		} else {
			fmt.Fprintln(GinkgoWriter, "No aliases found with name ", name, " for cleanup.")
		}
	}
}

func cleanupResources() {
	if testInstanceKeyGUID != "" {
		options := resourceControllerService.NewDeleteResourceKeyOptions(testInstanceKeyGUID)
		headers := map[string]string{
			"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
		}
		options.SetHeaders(headers)
		resp, err := resourceControllerService.DeleteResourceKey(options)
		if resp.StatusCode == 204 {
			fmt.Fprintf(GinkgoWriter, "Successful cleanup of key %s.\n", testInstanceKeyGUID)
		} else if resp.StatusCode == 410 {
			fmt.Fprintf(GinkgoWriter, "Key %s was already deleted by the tests.\n", testInstanceKeyGUID)
		} else {
			fmt.Fprintf(GinkgoWriter, "Failed to cleanup key %s. Error: %s\n", testInstanceKeyGUID, err.Error())
		}
	} else {
		fmt.Fprintln(GinkgoWriter, "Key for instance was not created. No cleanup needed.")
	}

	if testAliasKeyGUID != "" {
		options := resourceControllerService.NewDeleteResourceKeyOptions(testAliasKeyGUID)
		headers := map[string]string{
			"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
		}
		options.SetHeaders(headers)
		resp, err := resourceControllerService.DeleteResourceKey(options)
		if resp.StatusCode == 204 {
			fmt.Fprintf(GinkgoWriter, "Successful cleanup of key %s.\n", testAliasKeyGUID)
		} else if resp.StatusCode == 410 {
			fmt.Fprintf(GinkgoWriter, "Key %s was already deleted by the tests.\n", testAliasKeyGUID)
		} else {
			fmt.Fprintf(GinkgoWriter, "Failed to cleanup key %s. Error: %s\n", testAliasKeyGUID, err.Error())
		}
	} else {
		fmt.Fprintln(GinkgoWriter, "Key for alias was not created. No cleanup needed.")
	}

	if testBindingGUID != "" {
		options := resourceControllerService.NewDeleteResourceBindingOptions(testBindingGUID)
		headers := map[string]string{
			"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
		}
		options.SetHeaders(headers)
		resp, err := resourceControllerService.DeleteResourceBinding(options)
		if resp.StatusCode == 204 {
			fmt.Fprintf(GinkgoWriter, "Successful cleanup of binding %s.\n", testBindingGUID)
		} else if resp.StatusCode == 410 {
			fmt.Fprintf(GinkgoWriter, "Binding %s was already deleted by the tests.\n", testBindingGUID)
		} else {
			fmt.Fprintf(GinkgoWriter, "Failed to cleanup binding %s. Error: %s\n", testBindingGUID, err.Error())
		}
	} else {
		fmt.Fprintln(GinkgoWriter, "Binding was not created. No cleanup needed.")
	}

	if testAliasGUID != "" {
		options := resourceControllerService.NewDeleteResourceAliasOptions(testAliasGUID)
		headers := map[string]string{
			"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
		}
		options.SetHeaders(headers)
		resp, err := resourceControllerService.DeleteResourceAlias(options)
		if resp.StatusCode == 204 {
			fmt.Fprintf(GinkgoWriter, "Successful cleanup of alias %s.\n", testAliasGUID)
		} else if resp.StatusCode == 410 {
			fmt.Fprintf(GinkgoWriter, "Alias %s was already deleted by the tests.\n", testAliasGUID)
		} else {
			fmt.Fprintf(GinkgoWriter, "Failed to cleanup alias %s. Error: %s\n", testAliasGUID, err.Error())
		}
	} else {
		fmt.Fprintln(GinkgoWriter, "Alias was not created. No cleanup needed.")
	}

	if testInstanceGUID != "" {
		cleanupInstance()
	} else {
		fmt.Fprintln(GinkgoWriter, "Instance was not created. No cleanup needed.")
	}
}

func cleanupInstance() {
	options := resourceControllerService.NewGetResourceInstanceOptions(testInstanceGUID)
	headers := map[string]string{
		"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
	}
	options.SetHeaders(headers)
	result, _, err := resourceControllerService.GetResourceInstance(options)
	if err != nil {
		fmt.Fprintf(GinkgoWriter, "Failed to retrieve instance %s for cleanup.\n", testInstanceGUID)
		return
	}

	if *result.State == "active" && *result.Locked {
		options2 := resourceControllerService.NewUnlockResourceInstanceOptions(testInstanceGUID)
		headers2 := map[string]string{
			"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
		}
		options2 = options2.SetHeaders(headers2)
		_, _, err2 := resourceControllerService.UnlockResourceInstance(options2)
		if err2 != nil {
			fmt.Fprintf(GinkgoWriter, "Failed to unlock instance %s for cleanup. Error: %s\n", testInstanceGUID, err2.Error())
			return
		}
	}

	options3 := resourceControllerService.NewDeleteResourceInstanceOptions(testInstanceGUID)
	headers3 := map[string]string{
		"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
	}
	options3 = options3.SetHeaders(headers3)
	resp3, err3 := resourceControllerService.DeleteResourceInstance(options3)
	if resp3.StatusCode == 204 {
		fmt.Fprintf(GinkgoWriter, "Successful cleanup of instance %s.\n", testInstanceGUID)
	} else if resp3.StatusCode == 410 {
		fmt.Fprintf(GinkgoWriter, "Instance %s was already deleted by the tests.\n", testInstanceGUID)
	} else {
		fmt.Fprintf(GinkgoWriter, "Failed to cleanup instance %s. Error: %s\n", testInstanceGUID, err3.Error())
	}
}

func cleanupReclamationInstance() {
	options1 := resourceControllerService.NewGetResourceInstanceOptions(testReclaimInstanceGUID)
	headers1 := map[string]string{
		"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
	}
	options1 = options1.SetHeaders(headers1)
	result1, _, err1 := resourceControllerService.GetResourceInstance(options1)
	if err1 != nil {
		fmt.Fprintf(GinkgoWriter, "Failed to retrieve instance %s for cleanup.\n", testReclaimInstanceGUID)
		return
	}

	if *result1.State == "removed" {
		fmt.Fprintf(GinkgoWriter, "Instance %s was already reclaimed by the tests.\n", testReclaimInstanceGUID)
	} else if *result1.State == "pending_reclamation" {
		cleanupInstancePendingReclamation()
	} else {
		options2 := resourceControllerService.NewDeleteResourceInstanceOptions(testReclaimInstanceGUID)
		headers2 := map[string]string{
			"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
		}
		options2 = options2.SetHeaders(headers2)
		resp2, err2 := resourceControllerService.DeleteResourceInstance(options2)
		if resp2.StatusCode == 204 {
			fmt.Fprintf(GinkgoWriter, "Successfully scheduled instance %s for reclamation.\n", testReclaimInstanceGUID)
			time.Sleep(20 * time.Second)
			cleanupInstancePendingReclamation()
		} else {
			fmt.Fprintf(GinkgoWriter, "Failed to schedule active instance %s for reclamation. Error: %s\n", testReclaimInstanceGUID, err2.Error())
		}
	}
}

func cleanupInstancePendingReclamation() {
	options1 := resourceControllerService.NewListReclamationsOptions()
	options1 = options1.SetAccountID(testAccountID)
	options1 = options1.SetResourceInstanceID(testReclaimInstanceGUID)
	headers1 := map[string]string{
		"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
	}
	options1 = options1.SetHeaders(headers1)
	result1, _, err1 := resourceControllerService.ListReclamations(options1)
	if err1 != nil {
		fmt.Fprintf(GinkgoWriter, "Failed to retrieve reclamation to process to reclaim instance %s. Error: %s\n", testReclaimInstanceGUID, err1.Error())
		return
	}

	if len(result1.Resources) == 0 {
		fmt.Fprintf(GinkgoWriter, "Failed to retrieve reclamation to process to reclaim instance %s.\n", testReclaimInstanceGUID)
		return
	}

	reclamationID := *result1.Resources[0].ID
	if *result1.Resources[0].State != "RECLAIMING" {
		options2 := resourceControllerService.NewRunReclamationActionOptions(reclamationID, "reclaim")
		headers2 := map[string]string{
			"Transaction-ID": "rc-sdk-cleanup-" + transactionID,
		}
		options2 = options2.SetHeaders(headers2)
		_, _, err2 := resourceControllerService.RunReclamationAction(options2)
		if err2 != nil {
			fmt.Fprintf(GinkgoWriter, "Failed to process reclamation %s for instance %s. Error: %s\n", reclamationID, testReclaimInstanceGUID, err2.Error())
		} else {
			fmt.Fprintf(GinkgoWriter, "Successfully reclaimed instance %s.\n", testReclaimInstanceGUID)
		}
	} else {
		fmt.Fprintf(GinkgoWriter, "Instance %s was already reclaimed by the tests.\n", testReclaimInstanceGUID)
	}
}
