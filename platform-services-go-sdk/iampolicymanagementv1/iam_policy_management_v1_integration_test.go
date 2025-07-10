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

package iampolicymanagementv1_test

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

var _ = Describe("IAM Policy Management - Integration Tests", func() {
	const externalConfigFile = "../iam_policy_management.env"

	var (
		service *iampolicymanagementv1.IamPolicyManagementV1

		err          error
		config       map[string]string
		configLoaded bool = false

		testAccountID       string
		etagHeader          string = "ETag"
		testPolicyETag      string = ""
		testV2PolicyETag    string = ""
		testPolicyId        string = ""
		testV2PolicyId      string = ""
		testUserId          string = "IBMid-GoSDK" + strconv.Itoa(rand.Intn(100000))
		testViewerRoleCrn   string = "crn:v1:bluemix:public:iam::::role:Viewer"
		testOperatorRoleCrn string = "crn:v1:bluemix:public:iam::::role:Operator"
		testEditorRoleCrn   string = "crn:v1:bluemix:public:iam::::role:Editor"
		testServiceName     string = "iam-groups"

		testCustomRoleId                string = ""
		testCustomRoleETag              string = ""
		testCustomRoleName              string = "TestGoRole" + strconv.Itoa(rand.Intn(100000))
		testServiceRoleCrn              string = "crn:v1:bluemix:public:iam-identity::::serviceRole:ServiceIdCreator"
		testPolicyTemplateID            string = ""
		testPolicyOnlyTypeTemplateID    string = ""
		testPolicyS2STemplateID         string = ""
		testPolicyS2SOnlyTypeTemplateID string = ""

		testPolicyS2STemplateVersion                string = ""
		testPolicyS2SOnlyTypeTemplateVersions       string = ""
		testPolicyS2SUpdateTemplateVersion          string = ""
		testPolicyTemplateETag                      string = ""
		testPolicyOnlyPolicyTemplateETag            string = ""
		testPolicyTemplatePolicyTypeETag            string = ""
		testPolicyTemplateVersion                   string = ""
		testPolicyTemplatePolicyTypeVersion         string = ""
		testPolicyAssignmentId                      string = ""
		examplePolicyTemplateName                   string = "PolicySampleTemplateTestV1"
		TestPolicyType                              string = "TestPolicyType"
		assignmentPolicyID                          string
		testTargetAccountID                         string = ""
		testPolicyAssignmentETag                    string = ""
		testTargetType                              string = "Account"
		testAcountSettingsETag                      string = ""
		exampleActionControlTemplateName            string = "ActionControlTemplateGoSDK"
		exampleBasicActionControlTemplateName       string = "BasicActionControlTemplateGoSDK"
		exampleBasicActionControlUpdateTemplateName string = "BasicActionControlTemplateUpdateGoSDK"
		testActionControlTemplateID                 string = ""
		testBasicActionControlTemplateID            string = ""
		testActionControlTemplateVersion            string = ""
		testBasicActionControlTemplateETag          string = ""
		testBasicActionControlTemplateVersions      string = ""
		testActionControlUpdateTemplateVersion      string = ""
		testActionControlTemplateVersionETag        string = ""
		testActionControlAssignmentETag             string = ""
		testActionControlAssignmentId               string = ""
	)

	var shouldSkipTest = func() {
		if !configLoaded {
			Skip("External configuration is not available, skipping...")
		}
	}

	It("Successfully load the configuration", func() {
		err = os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
		if err != nil {
			Skip("Could not set IBM_CREDENTIALS_FILE environment variable: " + err.Error())
		}

		config, err = core.GetServiceProperties(iampolicymanagementv1.DefaultServiceName)
		if err == nil {
			testAccountID = config["TEST_ACCOUNT_ID"]
			testTargetAccountID = config["TEST_TARGET_ACCOUNT_ID"]
			if testAccountID != "" && testTargetAccountID != "" {
				configLoaded = true
			}
		}
		if !configLoaded {
			Skip("External configuration could not be loaded, skipping...")
		}
	})

	It(`Successfully created IamPolicyManagementV1 service instance`, func() {
		shouldSkipTest()

		service, err = iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(
			&iampolicymanagementv1.IamPolicyManagementV1Options{},
		)

		Expect(err).To(BeNil())
		Expect(service).ToNot(BeNil())

		core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
		service.EnableRetries(4, 30*time.Second)
	})

	Describe("Create an access policy", func() {

		It("Successfully created an access policy", func() {
			shouldSkipTest()

			// Construct an instance of the ResourceAttribute model
			accountIdResourceAttribute := new(iampolicymanagementv1.ResourceAttribute)
			accountIdResourceAttribute.Name = core.StringPtr("accountId")
			accountIdResourceAttribute.Value = core.StringPtr(testAccountID)
			accountIdResourceAttribute.Operator = core.StringPtr("stringEquals")

			serviceNameResourceAttribute := new(iampolicymanagementv1.ResourceAttribute)
			serviceNameResourceAttribute.Name = core.StringPtr("serviceType")
			serviceNameResourceAttribute.Value = core.StringPtr("service")
			serviceNameResourceAttribute.Operator = core.StringPtr("stringEquals")

			policyResourceTag := new(iampolicymanagementv1.ResourceTag)
			policyResourceTag.Name = core.StringPtr("project")
			policyResourceTag.Value = core.StringPtr("prototype")
			policyResourceTag.Operator = core.StringPtr("stringEquals")

			// Construct an instance of the SubjectAttribute model
			subjectAttribute := new(iampolicymanagementv1.SubjectAttribute)
			subjectAttribute.Name = core.StringPtr("iam_id")
			subjectAttribute.Value = core.StringPtr(testUserId)

			// Construct an instance of the PolicyResource model
			policyResource := new(iampolicymanagementv1.PolicyResource)
			policyResource.Attributes = []iampolicymanagementv1.ResourceAttribute{*accountIdResourceAttribute, *serviceNameResourceAttribute}
			policyResource.Tags = []iampolicymanagementv1.ResourceTag{*policyResourceTag}

			// Construct an instance of the PolicyRole model
			policyRole := new(iampolicymanagementv1.PolicyRole)
			policyRole.RoleID = core.StringPtr(testViewerRoleCrn)

			// Construct an instance of the PolicySubject model
			policySubject := new(iampolicymanagementv1.PolicySubject)
			policySubject.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttribute}

			// Construct an instance of the CreatePolicyOptions model
			options := new(iampolicymanagementv1.CreatePolicyOptions)
			options.Type = core.StringPtr("access")
			options.Subjects = []iampolicymanagementv1.PolicySubject{*policySubject}
			options.Roles = []iampolicymanagementv1.PolicyRole{*policyRole}
			options.Resources = []iampolicymanagementv1.PolicyResource{*policyResource}
			options.AcceptLanguage = core.StringPtr("en")

			policy, detailedResponse, err := service.CreatePolicy(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreatePolicy() result:\n%s\n", common.ToJSON(policy))
			Expect(policy.Type).To(Equal(options.Type))
			Expect(policy.Subjects).To(Equal(options.Subjects))
			Expect(policy.Roles[0].RoleID).To(Equal(options.Roles[0].RoleID))
			Expect(policy.Resources).To(Equal(options.Resources))

			testPolicyId = *policy.ID
		})
	})

	Describe("Get an access policy", func() {

		It("Successfully retrieved an access policy", func() {
			shouldSkipTest()
			Expect(testPolicyId).To(Not(BeNil()))

			options := service.NewGetPolicyOptions(testPolicyId)
			policy, detailedResponse, err := service.GetPolicy(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetPolicy() result:\n%s\n", common.ToJSON(policy))
			Expect(*policy.ID).To(Equal(testPolicyId))

			testPolicyETag = detailedResponse.GetHeaders().Get(etagHeader)
		})
	})

	Describe("Update an access policy", func() {

		It("Successfully updated an access policy", func() {
			shouldSkipTest()
			Expect(testPolicyId).To(Not(BeNil()))
			Expect(testPolicyETag).To(Not(BeNil()))

			// Construct an instance of the ResourceAttribute model
			accountIdResourceAttribute := new(iampolicymanagementv1.ResourceAttribute)
			accountIdResourceAttribute.Name = core.StringPtr("accountId")
			accountIdResourceAttribute.Value = core.StringPtr(testAccountID)
			accountIdResourceAttribute.Operator = core.StringPtr("stringEquals")

			serviceNameResourceAttribute := new(iampolicymanagementv1.ResourceAttribute)
			serviceNameResourceAttribute.Name = core.StringPtr("serviceType")
			serviceNameResourceAttribute.Value = core.StringPtr("service")
			serviceNameResourceAttribute.Operator = core.StringPtr("stringEquals")

			// Construct an instance of the SubjectAttribute model
			subjectAttribute := new(iampolicymanagementv1.SubjectAttribute)
			subjectAttribute.Name = core.StringPtr("iam_id")
			subjectAttribute.Value = core.StringPtr(testUserId)

			policyResourceTag := new(iampolicymanagementv1.ResourceTag)
			policyResourceTag.Name = core.StringPtr("project")
			policyResourceTag.Value = core.StringPtr("prototype")
			policyResourceTag.Operator = core.StringPtr("stringEquals")

			// Construct an instance of the PolicyResource model
			policyResource := new(iampolicymanagementv1.PolicyResource)
			policyResource.Attributes = []iampolicymanagementv1.ResourceAttribute{*accountIdResourceAttribute, *serviceNameResourceAttribute}
			policyResource.Tags = []iampolicymanagementv1.ResourceTag{*policyResourceTag}

			// Construct an instance of the PolicyRole model
			policyRole := new(iampolicymanagementv1.PolicyRole)
			policyRole.RoleID = core.StringPtr(testEditorRoleCrn)

			// Construct an instance of the PolicySubject model
			policySubject := new(iampolicymanagementv1.PolicySubject)
			policySubject.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttribute}

			// Construct an instance of the CreatePolicyOptions model
			options := new(iampolicymanagementv1.ReplacePolicyOptions)
			options.PolicyID = core.StringPtr(testPolicyId)
			options.IfMatch = core.StringPtr(testPolicyETag)
			options.Type = core.StringPtr("access")
			options.Subjects = []iampolicymanagementv1.PolicySubject{*policySubject}
			options.Roles = []iampolicymanagementv1.PolicyRole{*policyRole}
			options.Resources = []iampolicymanagementv1.PolicyResource{*policyResource}

			policy, detailedResponse, err := service.ReplacePolicy(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ReplacePolicy() result:\n%s\n", common.ToJSON(policy))
			Expect(*policy.ID).To(Equal(testPolicyId))
			Expect(policy.Type).To(Equal(options.Type))
			Expect(policy.Subjects).To(Equal(options.Subjects))
			Expect(policy.Roles[0].RoleID).To(Equal(options.Roles[0].RoleID))
			Expect(policy.Resources).To(Equal(options.Resources))

			testPolicyETag = detailedResponse.GetHeaders().Get(etagHeader)

		})
	})

	Describe("Patch an access policy", func() {

		It("Successfully patched an access policy", func() {
			shouldSkipTest()
			Expect(testPolicyId).To(Not(BeNil()))
			Expect(testPolicyETag).To(Not(BeNil()))

			// Construct an instance of the UpdatePolicyStateOptions model
			options := new(iampolicymanagementv1.UpdatePolicyStateOptions)
			options.PolicyID = &testPolicyId
			options.IfMatch = core.StringPtr(testPolicyETag)
			options.State = core.StringPtr("active")

			policy, detailedResponse, err := service.UpdatePolicyState(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "UpdatePolicyState() result:\n%s\n", common.ToJSON(policy))
			Expect(*policy.ID).To(Equal(testPolicyId))
			Expect(policy.State).To(Equal(options.State))

		})
	})

	Describe("List access policies", func() {

		It("Successfully listed the account's access policies", func() {
			shouldSkipTest()
			Expect(testPolicyId).To(Not(BeNil()))

			options := service.NewListPoliciesOptions(testAccountID)
			options.SetIamID(testUserId)
			result, detailedResponse, err := service.ListPolicies(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ListPolicies() result:\n%s\n", common.ToJSON(result))

			// confirm the test policy is present
			testPolicyPresent := false
			for _, policy := range result.Policies {
				if *policy.ID == testPolicyId {
					testPolicyPresent = true
				}
			}
			Expect(testPolicyPresent).To(BeTrue())
		})
	})

	Describe("Create a v2 access policy", func() {

		It("Successfully created a v2 access policy", func() {
			shouldSkipTest()

			// Construct an instance of the ResourceAttribute model
			accountIdResourceAttribute := new(iampolicymanagementv1.V2PolicyResourceAttribute)
			accountIdResourceAttribute.Key = core.StringPtr("accountId")
			accountIdResourceAttribute.Value = core.StringPtr(testAccountID)
			accountIdResourceAttribute.Operator = core.StringPtr("stringEquals")

			serviceNameResourceAttribute := new(iampolicymanagementv1.V2PolicyResourceAttribute)
			serviceNameResourceAttribute.Key = core.StringPtr("serviceType")
			serviceNameResourceAttribute.Value = core.StringPtr("service")
			serviceNameResourceAttribute.Operator = core.StringPtr("stringEquals")

			policyResourceTag := new(iampolicymanagementv1.V2PolicyResourceTag)
			policyResourceTag.Key = core.StringPtr("project")
			policyResourceTag.Value = core.StringPtr("prototype")
			policyResourceTag.Operator = core.StringPtr("stringEquals")

			// Construct an instance of the SubjectAttribute model
			subjectAttribute := new(iampolicymanagementv1.V2PolicySubjectAttribute)
			subjectAttribute.Key = core.StringPtr("iam_id")
			subjectAttribute.Operator = core.StringPtr("stringEquals")
			subjectAttribute.Value = core.StringPtr(testUserId)

			// Construct an instance of the V2PolicyResource model
			policyResource := new(iampolicymanagementv1.V2PolicyResource)
			policyResource.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*accountIdResourceAttribute, *serviceNameResourceAttribute}
			policyResource.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*policyResourceTag}

			// Construct an instance of the Roles model
			policyRole := new(iampolicymanagementv1.Roles)
			policyRole.RoleID = core.StringPtr(testViewerRoleCrn)

			// Construct an instance of the PolicySubject model
			policySubject := new(iampolicymanagementv1.V2PolicySubject)
			policySubject.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*subjectAttribute}

			// Contruct and instance of PolicyControl model
			control := new(iampolicymanagementv1.Control)
			grant := new(iampolicymanagementv1.Grant)
			grant.Roles = []iampolicymanagementv1.Roles{*policyRole}
			control.Grant = grant

			// Construct an instance of Policy Rule Attribute
			weeklyConditionAttribute := new(iampolicymanagementv1.NestedCondition)
			weeklyConditionAttribute.Key = core.StringPtr("{{environment.attributes.day_of_week}}")
			weeklyConditionAttribute.Operator = core.StringPtr("dayOfWeekAnyOf")
			weeklyConditionAttribute.Value = []string{"1+00:00", "2+00:00", "3+00:00", "4+00:00", "5+00:00"}

			startConditionAttribute := new(iampolicymanagementv1.NestedCondition)
			startConditionAttribute.Key = core.StringPtr("{{environment.attributes.current_time}}")
			startConditionAttribute.Operator = core.StringPtr("timeGreaterThanOrEquals")
			startConditionAttribute.Value = core.StringPtr("09:00:00+00:00")

			endConditionAttribute := new(iampolicymanagementv1.NestedCondition)
			endConditionAttribute.Key = core.StringPtr("{{environment.attributes.current_time}}")
			endConditionAttribute.Operator = core.StringPtr("timeLessThanOrEquals")
			endConditionAttribute.Value = core.StringPtr("17:00:00+00:00")

			policyRule := new(iampolicymanagementv1.V2PolicyRule)
			policyRule.Operator = core.StringPtr("and")
			policyRule.Conditions = []iampolicymanagementv1.NestedConditionIntf{weeklyConditionAttribute, startConditionAttribute, endConditionAttribute}

			// Construct an instance of the CreateV2PolicyOptions model
			options := new(iampolicymanagementv1.CreateV2PolicyOptions)
			options.Type = core.StringPtr("access")
			options.Subject = policySubject
			options.Control = control
			options.Resource = policyResource
			options.Pattern = core.StringPtr("time-based-conditions:weekly:custom-hours")
			options.Rule = policyRule
			options.AcceptLanguage = core.StringPtr("en")

			policy, detailedResponse, err := service.CreateV2Policy(options)
			controlResponse := new(iampolicymanagementv1.ControlResponse)
			controlResponse.Grant = grant
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateV2Policy() result:\n%s\n", common.ToJSON(policy))
			Expect(policy.Type).To(Equal(options.Type))
			Expect(policy.Subject.Attributes[0].Value).To(Equal(testUserId))
			Expect(policy.Control).To(Equal(controlResponse))
			Expect(policy.Resource.Attributes[0].Value).To(Equal(testAccountID))

			testV2PolicyId = *policy.ID
		})
	})

	Describe("Get a v2 access policy", func() {

		It("Successfully retrieved a v2 access policy", func() {
			shouldSkipTest()
			Expect(testPolicyId).To(Not(BeNil()))

			options := service.NewGetV2PolicyOptions(testV2PolicyId)
			policy, detailedResponse, err := service.GetV2Policy(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetV2Policy() result:\n%s\n", common.ToJSON(policy))
			Expect(*policy.ID).To(Equal(testV2PolicyId))

			testV2PolicyETag = detailedResponse.GetHeaders().Get(etagHeader)
		})
	})

	Describe("Update a v2 access policy", func() {

		It("Successfully updated a v2 access policy", func() {
			shouldSkipTest()
			Expect(testV2PolicyId).To(Not(BeNil()))
			Expect(testV2PolicyETag).To(Not(BeNil()))

			// Construct an instance of the ResourceAttribute model
			accountIdResourceAttribute := new(iampolicymanagementv1.V2PolicyResourceAttribute)
			accountIdResourceAttribute.Key = core.StringPtr("accountId")
			accountIdResourceAttribute.Value = core.StringPtr(testAccountID)
			accountIdResourceAttribute.Operator = core.StringPtr("stringEquals")

			serviceNameResourceAttribute := new(iampolicymanagementv1.V2PolicyResourceAttribute)
			serviceNameResourceAttribute.Key = core.StringPtr("serviceType")
			serviceNameResourceAttribute.Value = core.StringPtr("service")
			serviceNameResourceAttribute.Operator = core.StringPtr("stringEquals")

			// Construct an instance of the SubjectAttribute model
			subjectAttribute := new(iampolicymanagementv1.V2PolicySubjectAttribute)
			subjectAttribute.Key = core.StringPtr("iam_id")
			subjectAttribute.Operator = core.StringPtr("stringEquals")
			subjectAttribute.Value = core.StringPtr(testUserId)

			// Construct an instance of the V2PolicyBaseResource model
			policyResource := new(iampolicymanagementv1.V2PolicyResource)
			policyResource.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*accountIdResourceAttribute, *serviceNameResourceAttribute}

			// Construct an instance of the Roles model
			policyRole := new(iampolicymanagementv1.Roles)
			policyRole.RoleID = core.StringPtr(testViewerRoleCrn)

			// Construct an instance of the PolicySubject model
			policySubject := new(iampolicymanagementv1.V2PolicySubject)
			policySubject.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*subjectAttribute}

			// Contruct and instance of PolicyControl model
			control := new(iampolicymanagementv1.Control)
			grant := new(iampolicymanagementv1.Grant)
			grant.Roles = []iampolicymanagementv1.Roles{*policyRole}
			control.Grant = grant

			// Construct an instance of Policy Rule Attribute
			weeklyConditionAttribute := new(iampolicymanagementv1.NestedCondition)
			weeklyConditionAttribute.Key = core.StringPtr("{{environment.attributes.day_of_week}}")
			weeklyConditionAttribute.Operator = core.StringPtr("dayOfWeekAnyOf")
			weeklyConditionAttribute.Value = []string{"1+00:00", "2+00:00", "3+00:00", "4+00:00", "5+00:00"}

			startConditionAttribute := new(iampolicymanagementv1.NestedCondition)
			startConditionAttribute.Key = core.StringPtr("{{environment.attributes.current_time}}")
			startConditionAttribute.Operator = core.StringPtr("timeGreaterThanOrEquals")
			startConditionAttribute.Value = core.StringPtr("09:00:00+00:00")

			endConditionAttribute := new(iampolicymanagementv1.NestedCondition)
			endConditionAttribute.Key = core.StringPtr("{{environment.attributes.current_time}}")
			endConditionAttribute.Operator = core.StringPtr("timeLessThanOrEquals")
			endConditionAttribute.Value = core.StringPtr("17:00:00+00:00")

			policyRule := new(iampolicymanagementv1.V2PolicyRule)
			policyRule.Operator = core.StringPtr("and")
			policyRule.Conditions = []iampolicymanagementv1.NestedConditionIntf{weeklyConditionAttribute, startConditionAttribute, endConditionAttribute}

			// Construct an instance of the ReplaceV2PolicyOptions model
			options := new(iampolicymanagementv1.ReplaceV2PolicyOptions)
			options.ID = core.StringPtr(testV2PolicyId)
			options.IfMatch = core.StringPtr(testV2PolicyETag)
			options.Type = core.StringPtr("access")
			options.Subject = policySubject
			options.Control = control
			options.Resource = policyResource
			options.Pattern = core.StringPtr("time-based-conditions:weekly:custom-hours")
			options.Rule = policyRule

			policy, detailedResponse, err := service.ReplaceV2Policy(options)
			controlResponse := new(iampolicymanagementv1.ControlResponse)
			controlResponse.Grant = grant
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ReplaceV2Policy() result:\n%s\n", common.ToJSON(policy))
			Expect(*policy.ID).To(Equal(testV2PolicyId))
			Expect(policy.Type).To(Equal(options.Type))
			Expect(policy.Subject.Attributes[0].Value).To(Equal(testUserId))
			Expect(policy.Control).To(Equal(controlResponse))
			Expect(policy.Resource.Attributes[0].Value).To(Equal(testAccountID))

			newV2PolicyEtag := detailedResponse.GetHeaders().Get(etagHeader)
			Expect(newV2PolicyEtag).ToNot(Equal(testV2PolicyETag))

		})
	})

	Describe("List v2 access policies", func() {

		It("Successfully listed the account's v2 access policies", func() {
			shouldSkipTest()
			Expect(testV2PolicyId).To(Not(BeNil()))

			options := service.NewListV2PoliciesOptions(testAccountID)
			options.SetIamID(testUserId)
			options.SetSort("-id")
			result, detailedResponse, err := service.ListV2Policies(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ListV2Policies() result:\n%s\n", common.ToJSON(result))

			// confirm the test policy is present
			testPolicyPresent := false
			for _, policy := range result.Policies {
				if *policy.ID == testV2PolicyId {
					testPolicyPresent = true
				}
			}
			Expect(testPolicyPresent).To(BeTrue())
		})
	})

	Describe("Create custom role", func() {
		It("Successfully created custom role", func() {
			shouldSkipTest()

			actions := []string{"iam-groups.groups.read"}
			options := service.NewCreateRoleOptions(
				testCustomRoleName,
				actions,
				testCustomRoleName,
				testAccountID,
				testServiceName)
			options.SetDescription("GO SDK test role")
			result, detailedResponse, err := service.CreateRole(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "CreateRole() result:\n%s\n", common.ToJSON(result))

			testCustomRoleId = *result.ID
		})
	})

	Describe("Get a custom role", func() {
		It("Successfully retrieved a custom role", func() {
			shouldSkipTest()
			Expect(testCustomRoleId).To(Not(BeNil()))

			options := service.NewGetRoleOptions(testCustomRoleId)
			result, detailedResponse, err := service.GetRole(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetRole() result:\n%s\n", common.ToJSON(result))
			Expect(*result.ID).To(Equal(testCustomRoleId))

			testCustomRoleETag = detailedResponse.GetHeaders().Get(etagHeader)
		})
	})

	Describe("Update custom roles", func() {
		It("Successfully updated a custom role", func() {
			shouldSkipTest()
			Expect(testCustomRoleId).To(Not(BeNil()))
			Expect(testPolicyETag).To(Not(BeNil()))

			actions := []string{"iam-groups.groups.read"}
			options := service.NewReplaceRoleOptions(
				testCustomRoleId,
				testCustomRoleETag,
				testCustomRoleName,
				actions,
			)
			options.SetDescription("GO SDK test role udpated")
			options.SetDisplayName("GO SDK test role udpated")
			result, detailedResponse, err := service.ReplaceRole(options)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ReplaceRole() result:\n%s\n", common.ToJSON(result))
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testCustomRoleId))

		})
	})

	Describe("List custom roles", func() {
		It("Successfully listed the account's custom roles", func() {
			shouldSkipTest()
			Expect(testCustomRoleId).To(Not(BeNil()))

			options := service.NewListRolesOptions()
			options.SetAccountID(testAccountID)
			result, detailedResponse, err := service.ListRoles(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ListRoles() result:\n%s\n", common.ToJSON(result))

			// confirm the test policy is present
			testRolePresent := false
			for _, role := range result.CustomRoles {
				if *role.ID == testCustomRoleId {
					testRolePresent = true
				}
			}
			Expect(testRolePresent).To(BeTrue())
		})
	})

	Describe("List V2 roles", func() {
		It("Successfully listed the roles when account_id and service_group_id present", func() {
			shouldSkipTest()

			options := service.NewListRolesOptions()
			options.SetAccountID(testAccountID)
			options.SetServiceGroupID("IAM")
			result, detailedResponse, err := service.ListRoles(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "ListRoles() result:\n%s\n", common.ToJSON(result))

			// confirm the system's viewer and service roles are present
			testSystemRolePresent := false
			testServiceRolePresent := false
			for _, role := range result.SystemRoles {
				if *role.CRN == testViewerRoleCrn {
					testSystemRolePresent = true
				}
			}

			for _, role := range result.ServiceRoles {
				if *role.CRN == testServiceRoleCrn {
					testServiceRolePresent = true
				}
			}

			Expect(testSystemRolePresent).To(BeTrue())
			Expect(testServiceRolePresent).To(BeTrue())
		})
	})

	Describe(`CreatePolicyTemplate - Create a policy template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePolicyTemplate(createPolicyTemplateOptions *CreatePolicyTemplateOptions)`, func() {
			v2PolicyResourceAttributeModel := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("iam-access-management"),
			}

			v2PolicyResourceModel := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel},
			}

			rolesModel := &iampolicymanagementv1.Roles{
				RoleID: core.StringPtr(testViewerRoleCrn),
			}

			grantModel := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*rolesModel},
			}

			controlModel := &iampolicymanagementv1.Control{
				Grant: grantModel,
			}

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("access"),
				Description: core.StringPtr("Test Policy For Template"),
				Resource:    v2PolicyResourceModel,
				Control:     controlModel,
			}

			createPolicyTemplateOptions := &iampolicymanagementv1.CreatePolicyTemplateOptions{
				Name:           &examplePolicyTemplateName,
				AccountID:      &testAccountID,
				Policy:         templatePolicyModel,
				Description:    core.StringPtr("Test PolicySampleTemplate"),
				Committed:      core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("default"),
			}

			policyTemplate, response, err := service.CreatePolicyTemplate(createPolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyTemplate).ToNot(BeNil())
			Expect(policyTemplate.Name).To(Equal(core.StringPtr(examplePolicyTemplateName)))
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplate.AccountID).To(Equal(core.StringPtr(testAccountID)))
			Expect(policyTemplate.State).To(Equal(core.StringPtr("active")))

			testPolicyTemplateID = *policyTemplate.ID
		})
	})

	Describe(`CreatePolicyTemplate - Create a policy base template without Resource and Control`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePolicyTemplate(createPolicyTemplateOptions *CreatePolicyTemplateOptions testing)`, func() {
			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type: core.StringPtr("access"),
			}

			createPolicyTemplateOptions := &iampolicymanagementv1.CreatePolicyTemplateOptions{
				Name:           &TestPolicyType,
				AccountID:      &testAccountID,
				Policy:         templatePolicyModel,
				Description:    core.StringPtr("Test PolicySampleTemplate"),
				AcceptLanguage: core.StringPtr("default"),
			}

			policyTemplate, response, err := service.CreatePolicyTemplate(createPolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyTemplate).ToNot(BeNil())
			Expect(policyTemplate.Name).To(Equal(core.StringPtr(TestPolicyType)))
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplate.AccountID).To(Equal(core.StringPtr(testAccountID)))
			Expect(policyTemplate.State).To(Equal(core.StringPtr("active")))

			testPolicyOnlyTypeTemplateID = *policyTemplate.ID
			testPolicyTemplatePolicyTypeETag = response.GetHeaders().Get(etagHeader)
			testPolicyTemplatePolicyTypeVersion = *policyTemplate.Version
		})
	})

	Describe(`UpdatePolicyTemplate - Update a policy template with description and policy type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePolicyTemplate(createPolicyTemplateOptions *UpdatePolicyTemplateOptions)`, func() {

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type: core.StringPtr("access"),
			}

			replacePolicyTemplateOptions := &iampolicymanagementv1.ReplacePolicyTemplateOptions{
				PolicyTemplateID: &testPolicyOnlyTypeTemplateID,
				IfMatch:          &testPolicyTemplatePolicyTypeETag,
				Version:          &testPolicyTemplatePolicyTypeVersion,
				Policy:           templatePolicyModel,
				Description:      core.StringPtr("Template version update"),
			}

			policyTemplate, response, err := service.ReplacePolicyTemplate(replacePolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplate).ToNot(BeNil())

			Expect(policyTemplate.Version).To(Equal(core.StringPtr("1")))
			Expect(policyTemplate.Name).To(Equal(&TestPolicyType))
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplate.AccountID).To(Equal(&testAccountID))
			Expect(policyTemplate.State).To(Equal(core.StringPtr("active")))
			testPolicyTemplatePolicyTypeETag = response.GetHeaders().Get(etagHeader)
			testPolicyOnlyTypeTemplateID = *policyTemplate.ID
		})
	})

	Describe(`CreatePolicyS2STemplate - Create a s2s policy template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePolicyTemplate(createPolicyTemplateOptions *CreatePolicyTemplateOptions)`, func() {
			v2PolicyResourceAttributeModel := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("cloud-object-storage"),
			}

			v2PolicyResourceModel := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel},
			}

			v2PolicySubjectAttributeModel := &iampolicymanagementv1.V2PolicySubjectAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("compliance"),
			}

			v2PolicySubjectModel := &iampolicymanagementv1.V2PolicySubject{
				Attributes: []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel},
			}

			rolesModel := &iampolicymanagementv1.Roles{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Writer"),
			}

			grantModel := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*rolesModel},
			}

			controlModel := &iampolicymanagementv1.Control{
				Grant: grantModel,
			}

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("authorization"),
				Description: core.StringPtr("Test Policy For S2S Template"),
				Resource:    v2PolicyResourceModel,
				Subject:     v2PolicySubjectModel,
				Control:     controlModel,
			}

			createPolicyTemplateOptions := &iampolicymanagementv1.CreatePolicyTemplateOptions{
				Name:           core.StringPtr("S2S-Test"),
				AccountID:      &testAccountID,
				Policy:         templatePolicyModel,
				Description:    core.StringPtr("Test PolicySampleTemplate"),
				Committed:      core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("default"),
			}

			policyTemplate, response, err := service.CreatePolicyTemplate(createPolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyTemplate).ToNot(BeNil())
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("authorization")))
			Expect(policyTemplate.AccountID).To(Equal(core.StringPtr(testAccountID)))
			Expect(policyTemplate.State).To(Equal(core.StringPtr("active")))

			testPolicyS2STemplateID = *policyTemplate.ID
			testPolicyS2STemplateVersion = *policyTemplate.Version
		})
	})

	Describe(`CreatePolicyS2STemplate - Create a s2s policy template version without control, resource and subject`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePolicyTemplate(createPolicyTemplateOptions *CreatePolicyTemplateOptions) without control, resource and subject`, func() {

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("access"),
				Description: core.StringPtr("Watson Policy Template"),
			}

			createPolicyTemplateOptions := &iampolicymanagementv1.CreatePolicyTemplateOptions{
				Name:           core.StringPtr("S2S-Testing"),
				AccountID:      &testAccountID,
				Policy:         templatePolicyModel,
				Description:    core.StringPtr("Test PolicySampleTemplate"),
				AcceptLanguage: core.StringPtr("default"),
			}

			policyTemplate, response, err := service.CreatePolicyTemplate(createPolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyTemplate).ToNot(BeNil())
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplate.AccountID).To(Equal(core.StringPtr(testAccountID)))
			Expect(policyTemplate.State).To(Equal(core.StringPtr("active")))

			testPolicyS2SOnlyTypeTemplateID = *policyTemplate.ID
			testPolicyOnlyPolicyTemplateETag = response.GetHeaders().Get(etagHeader)
			testPolicyS2SOnlyTypeTemplateVersions = *policyTemplate.Version
		})
	})

	Describe(`ReplacePolicyTemplate - Update a policy template version with only type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplacePolicyTemplate(replacePolicyTemplateOptions *ReplacePolicyTemplateOptions)`, func() {
			v2PolicyResourceAttributeModel := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("watson"),
			}

			v2PolicyResourceModel := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel},
			}

			rolesModel := &iampolicymanagementv1.Roles{
				RoleID: &testOperatorRoleCrn,
			}

			grantModel := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*rolesModel},
			}

			controlModel := &iampolicymanagementv1.Control{
				Grant: grantModel,
			}

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("access"),
				Description: core.StringPtr("Version Update"),
				Resource:    v2PolicyResourceModel,
				Control:     controlModel,
			}

			replacePolicyTemplateOptions := &iampolicymanagementv1.ReplacePolicyTemplateOptions{
				PolicyTemplateID: &testPolicyS2SOnlyTypeTemplateID,
				Version:          &testPolicyS2SOnlyTypeTemplateVersions,
				IfMatch:          &testPolicyOnlyPolicyTemplateETag,
				Policy:           templatePolicyModel,
				Description:      core.StringPtr("Template version update"),
			}

			policyTemplate, response, err := service.ReplacePolicyTemplate(replacePolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplate).ToNot(BeNil())

			Expect(policyTemplate.Version).To(Equal(core.StringPtr("1")))
			Expect(policyTemplate.Name).To(Equal(core.StringPtr("S2S-Testing")))
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplate.AccountID).To(Equal(core.StringPtr(testAccountID)))
			Expect(policyTemplate.State).To(Equal(core.StringPtr("active")))
			testPolicyOnlyPolicyTemplateETag = response.GetHeaders().Get(etagHeader)
			testPolicyS2SOnlyTypeTemplateID = *policyTemplate.ID
		})
	})

	Describe(`ListPolicyTemplates - Get policy templates by attributes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPolicyTemplates(listPolicyTemplatesOptions *ListPolicyTemplatesOptions)`, func() {
			listPolicyTemplatesOptions := &iampolicymanagementv1.ListPolicyTemplatesOptions{
				AccountID:      &testAccountID,
				AcceptLanguage: core.StringPtr("default"),
			}

			policyTemplateCollection, response, err := service.ListPolicyTemplates(listPolicyTemplatesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplateCollection).ToNot(BeNil())

			Expect(policyTemplateCollection.PolicyTemplates[0].Policy.Type).ToNot(BeNil())
			Expect(policyTemplateCollection.PolicyTemplates[0].AccountID).To(Equal(&testAccountID))
			Expect(policyTemplateCollection.PolicyTemplates[0].State).To(Equal(core.StringPtr("active")))
		})
	})

	Describe(`GetPolicyTemplate - Retrieve latest policy template version by template ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPolicyTemplate(getPolicyTemplateOptions *GetPolicyTemplateOptions)`, func() {
			getPolicyTemplateOptions := &iampolicymanagementv1.GetPolicyTemplateOptions{
				PolicyTemplateID: &testPolicyTemplateID,
			}

			policyTemplate, response, err := service.GetPolicyTemplate(getPolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplate).ToNot(BeNil())

			Expect(policyTemplate.Name).To(Equal(&examplePolicyTemplateName))
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplate.AccountID).To(Equal(&testAccountID))
			Expect(policyTemplate.State).To(Equal(core.StringPtr("active")))
		})
	})

	Describe(`CreatePolicyTemplateVersion - Create a new policy template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePolicyTemplateVersion(createPolicyTemplateVersionOptions *CreatePolicyTemplateVersionOptions)`, func() {
			v2PolicyResourceAttributeModel := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("watson"),
			}

			v2PolicyResourceModel := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel},
			}

			rolesModel := &iampolicymanagementv1.Roles{
				RoleID: &testEditorRoleCrn,
			}

			grantModel := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*rolesModel},
			}

			controlModel := &iampolicymanagementv1.Control{
				Grant: grantModel,
			}

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("access"),
				Description: core.StringPtr("Watson Policy Template"),
				Resource:    v2PolicyResourceModel,
				Control:     controlModel,
			}

			createPolicyTemplateVersionOptions := &iampolicymanagementv1.CreatePolicyTemplateVersionOptions{
				PolicyTemplateID: &testPolicyTemplateID,
				Policy:           templatePolicyModel,
				Description:      core.StringPtr("Watson Policy Template version"),
			}

			policyTemplate, response, err := service.CreatePolicyTemplateVersion(createPolicyTemplateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyTemplate).ToNot(BeNil())

			Expect(policyTemplate.Name).To(Equal(&examplePolicyTemplateName))
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplate.AccountID).To(Equal(&testAccountID))
			Expect(policyTemplate.State).To(Equal(core.StringPtr("active")))

			testPolicyTemplateVersion = *policyTemplate.Version
			testPolicyTemplateETag = response.GetHeaders().Get(etagHeader)
		})
	})

	Describe(`CreatePolicyS2STemplateVersion - Create a new policy s2s template version`, func() {
		It(`CreatePolicyTemplateVersion(createPolicyTemplateVersionOptions *CreatePolicyTemplateVersionOptions)`, func() {
			v2PolicyResourceAttributeModel := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("kms"),
			}

			v2PolicyResourceModel := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel},
			}

			v2PolicySubjectAttributeModel := &iampolicymanagementv1.V2PolicySubjectAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("compliance"),
			}

			v2PolicySubjectModel := &iampolicymanagementv1.V2PolicySubject{
				Attributes: []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel},
			}

			rolesModel := &iampolicymanagementv1.Roles{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader"),
			}

			grantModel := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*rolesModel},
			}

			controlModel := &iampolicymanagementv1.Control{
				Grant: grantModel,
			}

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("authorization"),
				Description: core.StringPtr("Test Policy For S2S Template"),
				Resource:    v2PolicyResourceModel,
				Subject:     v2PolicySubjectModel,
				Control:     controlModel,
			}

			createPolicyTemplateOptions := &iampolicymanagementv1.CreatePolicyTemplateVersionOptions{
				Policy:           templatePolicyModel,
				PolicyTemplateID: core.StringPtr(testPolicyS2STemplateID),
				Description:      core.StringPtr("Test PolicySampleTemplate"),
				Committed:        core.BoolPtr(true),
			}

			policyTemplate, response, err := service.CreatePolicyTemplateVersion(createPolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyTemplate).ToNot(BeNil())
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("authorization")))
			Expect(policyTemplate.AccountID).To(Equal(core.StringPtr(testAccountID)))
			Expect(policyTemplate.State).To(Equal(core.StringPtr("active")))

			testPolicyS2SUpdateTemplateVersion = *policyTemplate.Version
		})
	})

	Describe(`ListPolicyTemplateVersions - Retrieve policy template versions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPolicyTemplateVersions(listPolicyTemplateVersionsOptions *ListPolicyTemplateVersionsOptions)`, func() {
			listPolicyTemplateVersionsOptions := &iampolicymanagementv1.ListPolicyTemplateVersionsOptions{
				PolicyTemplateID: &testPolicyTemplateID,
			}

			policyTemplateVersionsCollection, response, err := service.ListPolicyTemplateVersions(listPolicyTemplateVersionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplateVersionsCollection).ToNot(BeNil())
		})
	})

	Describe(`ReplacePolicyTemplate - Update a policy template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplacePolicyTemplate(replacePolicyTemplateOptions *ReplacePolicyTemplateOptions)`, func() {
			v2PolicyResourceAttributeModel := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("watson"),
			}

			v2PolicyResourceModel := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel},
			}

			rolesModel := &iampolicymanagementv1.Roles{
				RoleID: &testViewerRoleCrn,
			}

			grantModel := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*rolesModel},
			}

			controlModel := &iampolicymanagementv1.Control{
				Grant: grantModel,
			}

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("access"),
				Description: core.StringPtr("Version Update"),
				Resource:    v2PolicyResourceModel,
				Control:     controlModel,
			}

			replacePolicyTemplateOptions := &iampolicymanagementv1.ReplacePolicyTemplateOptions{
				PolicyTemplateID: &testPolicyTemplateID,
				Version:          &testPolicyTemplateVersion,
				IfMatch:          &testPolicyTemplateETag,
				Policy:           templatePolicyModel,
				Description:      core.StringPtr("Template version update"),
			}

			policyTemplate, response, err := service.ReplacePolicyTemplate(replacePolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplate).ToNot(BeNil())

			Expect(policyTemplate.Version).To(Equal(core.StringPtr("2")))
			Expect(policyTemplate.Name).To(Equal(&examplePolicyTemplateName))
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplate.AccountID).To(Equal(&testAccountID))
			Expect(policyTemplate.State).To(Equal(core.StringPtr("active")))
			testPolicyTemplateETag = response.GetHeaders().Get(etagHeader)

		})
	})

	Describe(`GetPolicyTemplateVersion - Retrieve a policy template version by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPolicyTemplateVersion(getPolicyTemplateVersionOptions *GetPolicyTemplateVersionOptions)`, func() {
			getPolicyTemplateVersionOptions := &iampolicymanagementv1.GetPolicyTemplateVersionOptions{
				PolicyTemplateID: &testPolicyTemplateID,
				Version:          &testPolicyTemplateVersion,
			}

			policyTemplate, response, err := service.GetPolicyTemplateVersion(getPolicyTemplateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplate).ToNot(BeNil())

			Expect(policyTemplate.Version).To(Equal(core.StringPtr("2")))
			Expect(policyTemplate.Policy.Type).To(Equal(core.StringPtr("access")))
			Expect(policyTemplate.AccountID).To(Equal(&testAccountID))
			Expect(policyTemplate.State).To(Equal(core.StringPtr("active")))
		})
	})

	Describe(`CommitPolicyTemplate - Commit a policy template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CommitPolicyTemplate(commitPolicyTemplateOptions *CommitPolicyTemplateOptions)`, func() {
			commitPolicyTemplateOptions := &iampolicymanagementv1.CommitPolicyTemplateOptions{
				PolicyTemplateID: &testPolicyTemplateID,
				Version:          &testPolicyTemplateVersion,
			}

			response, err := service.CommitPolicyTemplate(commitPolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`CreatePolicyAssignments - Create policy assignments by templates`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePolicyTemplateAssignment(createPolicyTemplateAssignmentOptions *CreatePolicyTemplateAssignmentOptions)`, func() {
			template := iampolicymanagementv1.AssignmentTemplateDetails{
				ID:      &testPolicyS2STemplateID,
				Version: &testPolicyS2STemplateVersion,
			}
			templates := []iampolicymanagementv1.AssignmentTemplateDetails{
				template,
			}

			target := &iampolicymanagementv1.AssignmentTargetDetails{
				Type: &testTargetType,
				ID:   &testTargetAccountID,
			}

			createPolicyTemplateVersionOptions := &iampolicymanagementv1.CreatePolicyTemplateAssignmentOptions{
				Version:   core.StringPtr("1.0"),
				Target:    target,
				Templates: templates,
			}

			policyAssignment, response, err := service.CreatePolicyTemplateAssignment(createPolicyTemplateVersionOptions)
			var assignmentDetails = policyAssignment.Assignments[0]
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(*assignmentDetails.Resources[0].Policy.ResourceCreated.ID).ToNot(BeNil())
			Expect(*assignmentDetails.Resources[0].Target.ID).ToNot(BeNil())
			Expect(*assignmentDetails.Resources[0].Target.ID).To(Equal(testTargetAccountID))
			Expect(*assignmentDetails.Resources[0].Target.Type).ToNot(BeNil())
			Expect(*assignmentDetails.Resources[0].Target.Type).To(Equal(testTargetType))
			testPolicyAssignmentETag = response.GetHeaders().Get(etagHeader)
			testPolicyAssignmentId = *assignmentDetails.ID
		})
	})

	Describe(`UpdatePolicyAssignment - update a policy assignment by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		Expect(testPolicyAssignmentETag).To(Not(BeNil()))
		It(`UpdatePolicyAssignment(updatePolicyAssignmentOptions *UpdatePolicyAssignmentOptions))`, func() {
			updatePolicyAssignmentOptions := &iampolicymanagementv1.UpdatePolicyAssignmentOptions{
				AssignmentID:    core.StringPtr(testPolicyAssignmentId),
				Version:         core.StringPtr("1.0"),
				TemplateVersion: core.StringPtr(testPolicyS2SUpdateTemplateVersion),
				IfMatch:         core.StringPtr(testPolicyAssignmentETag),
			}

			policyAssignment, response, err := service.UpdatePolicyAssignment(updatePolicyAssignmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(*policyAssignment.Resources[0].Policy.ResourceCreated.ID).ToNot(BeNil())
			Expect(*policyAssignment.ID).To(Equal(testPolicyAssignmentId))
		})
	})

	Describe(`ListPolicyAssignments - Get policies template assignments by attributes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPolicyAssignments(listPolicyAssignmentsOptions *ListPolicyAssignmentsOptions)`, func() {
			listPolicyAssignmentsOptions := &iampolicymanagementv1.ListPolicyAssignmentsOptions{
				AccountID:      core.StringPtr(testAccountID),
				AcceptLanguage: core.StringPtr("default"),
				Version:        core.StringPtr("1.0"),
			}

			policyTemplateAssignmentCollection, response, err := service.ListPolicyAssignments(listPolicyAssignmentsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplateAssignmentCollection).ToNot(BeNil())
			var assignmentDetails = policyTemplateAssignmentCollection.Assignments[0].(*iampolicymanagementv1.PolicyTemplateAssignmentItems)

			Expect(*assignmentDetails).ToNot(BeNil())
			Expect(*assignmentDetails.Template.ID).ToNot(BeNil())
			Expect(*assignmentDetails.Target.Type).ToNot(BeNil())
			Expect(*assignmentDetails.Template.Version).ToNot(BeNil())
			Expect(*assignmentDetails.Target.ID).ToNot(BeNil())
			Expect(*assignmentDetails.Status).ToNot(BeNil())
			Expect(*assignmentDetails.AccountID).ToNot(BeNil())
			Expect(assignmentDetails.Resources).ToNot(BeNil())
			Expect(*assignmentDetails.CreatedAt).ToNot(BeNil())
			Expect(*assignmentDetails.CreatedByID).ToNot(BeNil())
			Expect(*assignmentDetails.LastModifiedAt).ToNot(BeNil())
			Expect(*assignmentDetails.LastModifiedByID).ToNot(BeNil())
			Expect(*assignmentDetails.Href).ToNot(BeNil())
		})
	})

	Describe(`GetPolicyAssignment - Retrieve a policy assignment by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPolicyAssignment(getPolicyAssignmentOptions *GetPolicyAssignmentOptions)`, func() {
			getPolicyAssignmentOptions := &iampolicymanagementv1.GetPolicyAssignmentOptions{
				AssignmentID: core.StringPtr(testPolicyAssignmentId),
				Version:      core.StringPtr("1.0"),
			}

			policyAssignmentRecord, response, err := service.GetPolicyAssignment(getPolicyAssignmentOptions)
			Expect(err).To(BeNil())
			var assignmentDetails = policyAssignmentRecord.(*iampolicymanagementv1.PolicyTemplateAssignmentItems)
			Expect(response.StatusCode).To(Equal(200))
			Expect(*assignmentDetails).ToNot(BeNil())
			Expect(*assignmentDetails.Template.ID).ToNot(BeNil())
			Expect(*assignmentDetails.Target.Type).ToNot(BeNil())
			Expect(*assignmentDetails.Template.Version).ToNot(BeNil())
			Expect(*assignmentDetails.Target.ID).ToNot(BeNil())
			Expect(*assignmentDetails.Status).ToNot(BeNil())
			Expect(*assignmentDetails.AccountID).ToNot(BeNil())
			Expect(*assignmentDetails.CreatedAt).ToNot(BeNil())
			Expect(*assignmentDetails.CreatedByID).ToNot(BeNil())
			Expect(*assignmentDetails.LastModifiedAt).ToNot(BeNil())
			Expect(*assignmentDetails.LastModifiedByID).ToNot(BeNil())
			Expect(*assignmentDetails.Href).ToNot(BeNil())
			assignmentPolicyID = *assignmentDetails.Resources[0].Policy.ResourceCreated.ID
		})
	})

	Describe("GetPolicyV2 - Retrieve Policy Template MetaData created from assignment", func() {

		It("Successfully retrieved a v2 access policy", func() {
			shouldSkipTest()
			Expect(testPolicyId).To(Not(BeNil()))

			options := service.NewGetV2PolicyOptions(assignmentPolicyID)
			policy, detailedResponse, err := service.GetV2Policy(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "GetV2Policy() result:\n%s\n", common.ToJSON(policy))
			Expect(*policy.ID).To(Equal(assignmentPolicyID))
			Expect(policy.Template).ToNot(BeNil())
			Expect(policy.Template.ID).ToNot(BeNil())
			Expect(policy.Template.Version).ToNot(BeNil())
			Expect(policy.Template.AssignmentID).ToNot(BeNil())
		})
	})

	Describe(`DeletePolicyAssignment - Delete a policy assignment by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePolicyAssignment(deletePolicyAssignmentOptions *DeletePolicyAssignmentOptions)`, func() {
			deletePolicyAssignmentOptions := &iampolicymanagementv1.DeletePolicyAssignmentOptions{
				AssignmentID: core.StringPtr(testPolicyAssignmentId),
			}

			response, err := service.DeletePolicyAssignment(deletePolicyAssignmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePolicyTemplateVersion - Delete a policy template version by ID and version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePolicyTemplateVersion(deletePolicyTemplateVersionOptions *DeletePolicyTemplateVersionOptions)`, func() {
			deletePolicyTemplateVersionOptions := &iampolicymanagementv1.DeletePolicyTemplateVersionOptions{
				PolicyTemplateID: &testPolicyTemplateID,
				Version:          &testPolicyTemplateVersion,
			}

			response, err := service.DeletePolicyTemplateVersion(deletePolicyTemplateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePolicyTemplate - Delete a policy template ID Only Type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePolicyTemplate(deletePolicyTemplateOptions *DeletePolicyTemplateOptions)`, func() {
			deletePolicyTemplateOptions := &iampolicymanagementv1.DeletePolicyTemplateOptions{
				PolicyTemplateID: &testPolicyOnlyTypeTemplateID,
			}

			response, err := service.DeletePolicyTemplate(deletePolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePolicyTemplate - Delete a policy template by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePolicyTemplate(deletePolicyTemplateOptions *DeletePolicyTemplateOptions)`, func() {
			deletePolicyTemplateOptions := &iampolicymanagementv1.DeletePolicyTemplateOptions{
				PolicyTemplateID: &testPolicyTemplateID,
			}

			response, err := service.DeletePolicyTemplate(deletePolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePolicyS2STemplate - Delete a policy s2s template by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePolicyTemplate(deletePolicyTemplateOptions *DeletePolicyTemplateOptions)`, func() {
			deletePolicyTemplateOptions := &iampolicymanagementv1.DeletePolicyTemplateOptions{
				PolicyTemplateID: &testPolicyS2STemplateID,
			}

			response, err := service.DeletePolicyTemplate(deletePolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeletePolicyS2STemplate - Delete a policy s2s template by ID only Type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePolicyTemplate(deletePolicyTemplateOptions *DeletePolicyTemplateOptions)`, func() {
			deletePolicyTemplateOptions := &iampolicymanagementv1.DeletePolicyTemplateOptions{
				PolicyTemplateID: &testPolicyS2SOnlyTypeTemplateID,
			}

			response, err := service.DeletePolicyTemplate(deletePolicyTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`GetSettings - Retrieve Access Management account settings by account ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
			getSettingsOptions := &iampolicymanagementv1.GetSettingsOptions{
				AccountID:      &testAccountID,
				AcceptLanguage: core.StringPtr("default"),
			}

			accountSettingsAccessManagement, response, err := service.GetSettings(getSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettingsAccessManagement).ToNot(BeNil())
			testAcountSettingsETag = response.GetHeaders().Get(etagHeader)
		})
	})

	Describe(`UpdateSettings - Updates Access Management account settings by account ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions)`, func() {
			identityTypesBaseModel := &iampolicymanagementv1.IdentityTypesBase{
				State:                   core.StringPtr("monitor"),
				ExternalAllowedAccounts: []string{},
			}

			identityTypesPatchModel := &iampolicymanagementv1.IdentityTypesPatch{
				User:      identityTypesBaseModel,
				ServiceID: identityTypesBaseModel,
				Service:   identityTypesBaseModel,
			}

			externalAccountIdentityInteractionPatchModel := &iampolicymanagementv1.ExternalAccountIdentityInteractionPatch{
				IdentityTypes: identityTypesPatchModel,
			}

			updateSettingsOptions := &iampolicymanagementv1.UpdateSettingsOptions{
				AccountID:                          &testAccountID,
				IfMatch:                            &testAcountSettingsETag,
				ExternalAccountIdentityInteraction: externalAccountIdentityInteractionPatchModel,
				AcceptLanguage:                     core.StringPtr("default"),
			}

			accountSettingsAccessManagement, response, err := service.UpdateSettings(updateSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettingsAccessManagement).ToNot(BeNil())
		})
	})

	Describe(`CreateBasicActionControlTemplate - Create a basic action template version without action_control`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateBasicActionControlTemplate(createActionControlTemplateOptions *CreateActionControlTemplateOptions) without action_control`, func() {
			createActionControlTemplateOptions := &iampolicymanagementv1.CreateActionControlTemplateOptions{
				Name:           core.StringPtr(exampleBasicActionControlTemplateName),
				AccountID:      &testAccountID,
				Description:    core.StringPtr("Test Basic ActionControl Template from GO SDK"),
				AcceptLanguage: core.StringPtr("default"),
			}

			actionControlTemplate, response, err := service.CreateActionControlTemplate(createActionControlTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(actionControlTemplate).ToNot(BeNil())
			Expect(actionControlTemplate.Name).To(Equal(core.StringPtr(exampleBasicActionControlTemplateName)))
			Expect(actionControlTemplate.AccountID).To(Equal(core.StringPtr(testAccountID)))
			Expect(actionControlTemplate.State).To(Equal(core.StringPtr("active")))

			testBasicActionControlTemplateID = *actionControlTemplate.ID
			testBasicActionControlTemplateETag = response.GetHeaders().Get(etagHeader)
			testBasicActionControlTemplateVersions = *actionControlTemplate.Version
		})
	})

	Describe(`ReplaceBasicActionControlTemplate - Update a basic action control template version without action_control`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceActionControlTemplate(replaceActionControlTemplateOptions *ReplaceActionControlTemplateOptions)`, func() {
			replaceActionControlTemplateOptions := &iampolicymanagementv1.ReplaceActionControlTemplateOptions{
				Description:             core.StringPtr("Test ActionControl Template from GO SDK"),
				Name:                    core.StringPtr(exampleBasicActionControlUpdateTemplateName),
				ActionControlTemplateID: core.StringPtr(testBasicActionControlTemplateID),
				IfMatch:                 core.StringPtr(testBasicActionControlTemplateETag),
				Version:                 core.StringPtr(testBasicActionControlTemplateVersions),
			}

			actionControlTemplate, response, err := service.ReplaceActionControlTemplate(replaceActionControlTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionControlTemplate).ToNot(BeNil())
			testPolicyOnlyPolicyTemplateETag = response.GetHeaders().Get(etagHeader)
			testBasicActionControlTemplateID = *actionControlTemplate.ID
			testBasicActionControlTemplateETag = response.GetHeaders().Get(etagHeader)
			testBasicActionControlTemplateVersions = *actionControlTemplate.Version

			Expect(actionControlTemplate.Version).To(Equal(core.StringPtr("1")))
			Expect(actionControlTemplate.Name).To(Equal(core.StringPtr(exampleBasicActionControlUpdateTemplateName)))
			Expect(actionControlTemplate.AccountID).To(Equal(core.StringPtr(testAccountID)))
			Expect(actionControlTemplate.State).To(Equal(core.StringPtr("active")))
		})
	})

	Describe(`ReplaceBasicActionControlTemplate - Update a basic action control template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceActionControlTemplate(replaceActionControlTemplateOptions *ReplaceActionControlTemplateOptions)`, func() {
			templateActionControl := &iampolicymanagementv1.TemplateActionControl{
				ServiceName: core.StringPtr("am-test-service"),
				Description: core.StringPtr("am-test-service service actionControl"),
				Actions:     []string{"am-test-service.test.delete"},
			}

			replaceActionControlTemplateOptions := &iampolicymanagementv1.ReplaceActionControlTemplateOptions{
				ActionControl:           templateActionControl,
				Description:             core.StringPtr("Test ActionControl Template from GO SDK"),
				Committed:               core.BoolPtr(true),
				ActionControlTemplateID: core.StringPtr(testBasicActionControlTemplateID),
				IfMatch:                 core.StringPtr(testBasicActionControlTemplateETag),
				Version:                 core.StringPtr(testBasicActionControlTemplateVersions),
			}

			actionControlTemplate, response, err := service.ReplaceActionControlTemplate(replaceActionControlTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionControlTemplate).ToNot(BeNil())

			Expect(actionControlTemplate.Version).To(Equal(core.StringPtr("1")))
			Expect(actionControlTemplate.Name).To(Equal(core.StringPtr(exampleBasicActionControlUpdateTemplateName)))
			Expect(actionControlTemplate.AccountID).To(Equal(core.StringPtr(testAccountID)))
			Expect(actionControlTemplate.State).To(Equal(core.StringPtr("active")))
		})
	})

	Describe(`GetBasicActionControlTemplate - Retrieve action control template version by template ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetActionControlTemplate(getActionControlTemplateOptions *GetActionControlTemplateOptions)`, func() {
			getActionControlTemplateOptions := &iampolicymanagementv1.GetActionControlTemplateOptions{
				ActionControlTemplateID: &testBasicActionControlTemplateID,
			}

			actionControlTemplate, response, err := service.GetActionControlTemplate(getActionControlTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionControlTemplate).ToNot(BeNil())

			Expect(actionControlTemplate.Name).To(Equal(core.StringPtr(exampleBasicActionControlUpdateTemplateName)))
			Expect(actionControlTemplate.AccountID).To(Equal(&testAccountID))
			Expect(actionControlTemplate.State).To(Equal(core.StringPtr("active")))
		})
	})

	Describe(`DeleteBasicActionControlTemplate - Delete an action control template by ID and version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteActionControlTemplateVersion(deleteActionControlTemplateVersionOptions *DeleteActionControlTemplateVersionOptions`, func() {
			deleteActionControlTemplateOptions := &iampolicymanagementv1.DeleteActionControlTemplateVersionOptions{
				ActionControlTemplateID: &testBasicActionControlTemplateID,
				Version:                 &testBasicActionControlTemplateVersions,
			}

			response, err := service.DeleteActionControlTemplateVersion(deleteActionControlTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`CreateTestActionControlTemplate - Create an action control template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTestActionControlTemplate(createActionControlTemplateOptions *CreateActionControlTemplateOptions)`, func() {
			templateActionControl := &iampolicymanagementv1.TemplateActionControl{
				ServiceName: core.StringPtr("am-test-service"),
				Description: core.StringPtr("am-test-service service actionControl"),
				Actions:     []string{"am-test-service.test.create"},
			}

			createActionControlTemplateOptions := &iampolicymanagementv1.CreateActionControlTemplateOptions{
				Name:           &exampleActionControlTemplateName,
				AccountID:      &testAccountID,
				ActionControl:  templateActionControl,
				Description:    core.StringPtr("Test ActionControl Template from GO SDK"),
				Committed:      core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("default"),
			}

			actionControlTemplate, response, err := service.CreateActionControlTemplate(createActionControlTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(actionControlTemplate).ToNot(BeNil())
			Expect(actionControlTemplate.Name).To(Equal(core.StringPtr(exampleActionControlTemplateName)))
			Expect(actionControlTemplate.AccountID).To(Equal(core.StringPtr(testAccountID)))
			Expect(actionControlTemplate.State).To(Equal(core.StringPtr("active")))

			testActionControlTemplateID = *actionControlTemplate.ID
			testActionControlTemplateVersion = *actionControlTemplate.Version
		})
	})

	Describe(`CreateTestActionControlTemplateVersion - Create a new action control template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`CreateTestActionControlTemplateVersion(createActionControlTemplateVersionOptions *CreateActionControlTemplateVersionOptions)`, func() {
			templateActionControl := &iampolicymanagementv1.TemplateActionControl{
				ServiceName: core.StringPtr("am-test-service"),
				Description: core.StringPtr("am-test-service service actionControl"),
				Actions:     []string{"am-test-service.test.delete"},
			}

			updateActionControlTemplateVersionOptions := &iampolicymanagementv1.CreateActionControlTemplateVersionOptions{
				ActionControl:           templateActionControl,
				Description:             core.StringPtr("Test of ActionControl Template version from GO SDK"),
				ActionControlTemplateID: &testActionControlTemplateID,
			}

			actionControlTemplate, response, err := service.CreateActionControlTemplateVersion(updateActionControlTemplateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(actionControlTemplate).ToNot(BeNil())
			Expect(actionControlTemplate.AccountID).To(Equal(core.StringPtr(testAccountID)))
			Expect(actionControlTemplate.Name).To(Equal(core.StringPtr(exampleActionControlTemplateName)))
			Expect(actionControlTemplate.State).To(Equal(core.StringPtr("active")))

			testActionControlUpdateTemplateVersion = *actionControlTemplate.Version
			testActionControlTemplateVersionETag = response.GetHeaders().Get(etagHeader)
		})
	})

	Describe(`ReplaceTestActionControlTemplateVersion - Update an action control template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`ReplaceTestActionControlTemplate(replaceActionControlTemplateOptions *ReplaceActionControlTemplateOptions)`, func() {

			templateActionControl := &iampolicymanagementv1.TemplateActionControl{
				ServiceName: core.StringPtr("am-test-service"),
				Description: core.StringPtr("am-test-service service actionControl"),
				Actions:     []string{"am-test-service.test.delete", "am-test-service.test.create"},
			}

			replaceActionControlTemplateVersionOptions := &iampolicymanagementv1.ReplaceActionControlTemplateOptions{
				ActionControl:           templateActionControl,
				Description:             core.StringPtr("Test update of ActionControl Template from GO SDK"),
				Committed:               core.BoolPtr(true),
				IfMatch:                 core.StringPtr(testActionControlTemplateVersionETag),
				Version:                 core.StringPtr(testActionControlUpdateTemplateVersion),
				ActionControlTemplateID: &testActionControlTemplateID,
			}

			actionControlTemplate, response, err := service.ReplaceActionControlTemplate(replaceActionControlTemplateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionControlTemplate).ToNot(BeNil())
			Expect(actionControlTemplate.AccountID).To(Equal(core.StringPtr(testAccountID)))
			Expect(actionControlTemplate.Name).To(Equal(core.StringPtr(exampleActionControlTemplateName)))
			Expect(actionControlTemplate.State).To(Equal(core.StringPtr("active")))

			testActionControlUpdateTemplateVersion = *actionControlTemplate.Version
		})
	})

	Describe(`ListActionControlTemplates - Get action control templates by attributes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListActionControlTemplates(listActionControlTemplatesOptions *ListActionControlTemplatesOptions)`, func() {
			listActionControlTemplatesOptions := &iampolicymanagementv1.ListActionControlTemplatesOptions{
				AccountID:      &testAccountID,
				AcceptLanguage: core.StringPtr("default"),
			}

			actionControlTemplateCollection, response, err := service.ListActionControlTemplates(listActionControlTemplatesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionControlTemplateCollection).ToNot(BeNil())

			Expect(actionControlTemplateCollection.ActionControlTemplates[0].AccountID).To(Equal(&testAccountID))
			Expect(actionControlTemplateCollection.ActionControlTemplates[0].State).To(Equal(core.StringPtr("active")))
		})
	})

	Describe(`GetActionControlTemplateVersion - Retrieve an action control template version by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetActionControlTemplateVersion(getActionControlTemplateVersionOptions *GetActionControlTemplateVersionOptions)`, func() {
			getActionControlTemplateVersionOptions := &iampolicymanagementv1.GetActionControlTemplateVersionOptions{
				ActionControlTemplateID: &testActionControlTemplateID,
				Version:                 &testActionControlUpdateTemplateVersion,
			}

			actionControlTemplate, response, err := service.GetActionControlTemplateVersion(getActionControlTemplateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionControlTemplate).ToNot(BeNil())

			Expect(actionControlTemplate.Version).To(Equal(core.StringPtr("2")))
			Expect(actionControlTemplate.AccountID).To(Equal(&testAccountID))
			Expect(actionControlTemplate.State).To(Equal(core.StringPtr("active")))
		})
	})

	Describe(`CommitActionControlTemplate - Commit an action control template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CommitActionControlTemplate(commitActionControlTemplateOptions *CommitActionControlTemplateOptions)`, func() {
			commitActionControlTemplateOptions := &iampolicymanagementv1.CommitActionControlTemplateOptions{
				ActionControlTemplateID: &testActionControlTemplateID,
				Version:                 &testActionControlTemplateVersion,
			}

			response, err := service.CommitActionControlTemplate(commitActionControlTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`GetActionControlTemplate - Retrieve latest action control template version by template ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetActionControlTemplate(getActionControlTemplateOptions *GetActionControlTemplateOptions)`, func() {
			getActionControlTemplateOptions := &iampolicymanagementv1.GetActionControlTemplateOptions{
				ActionControlTemplateID: &testActionControlTemplateID,
			}

			actionControlTemplate, response, err := service.GetActionControlTemplate(getActionControlTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionControlTemplate).ToNot(BeNil())

			Expect(actionControlTemplate.Name).To(Equal(&exampleActionControlTemplateName))
			Expect(actionControlTemplate.AccountID).To(Equal(&testAccountID))
			Expect(actionControlTemplate.State).To(Equal(core.StringPtr("active")))
		})
	})

	Describe(`ListActionControlTemplateVersions - Retrieve action control template versions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListActionControlTemplateVersions(listActionControlTemplateVersionsOptions *ListActionControlTemplateVersionsOptions)`, func() {
			listActionControlTemplateVersionsOptions := &iampolicymanagementv1.ListActionControlTemplateVersionsOptions{
				ActionControlTemplateID: &testActionControlTemplateID,
			}

			actionControlTemplateVersionsCollection, response, err := service.ListActionControlTemplateVersions(listActionControlTemplateVersionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionControlTemplateVersionsCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateActionControlAssignments - Create action control assignments by action template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateActionControlTemplateAssignment(createActionControlTemplateAssignmentOptions *CreateActionControlTemplateAssignmentOptions)`, func() {
			template := iampolicymanagementv1.ActionControlAssignmentTemplate{
				ID:      &testActionControlTemplateID,
				Version: &testActionControlTemplateVersion,
			}
			templates := []iampolicymanagementv1.ActionControlAssignmentTemplate{
				template,
			}

			target := &iampolicymanagementv1.AssignmentTargetDetails{
				Type: &testTargetType,
				ID:   &testTargetAccountID,
			}

			createPolicyTemplateVersionOptions := &iampolicymanagementv1.CreateActionControlTemplateAssignmentOptions{
				Target:    target,
				Templates: templates,
			}

			actionControlAssignment, response, err := service.CreateActionControlTemplateAssignment(createPolicyTemplateVersionOptions)
			var assignmentDetails = actionControlAssignment.Assignments[0]
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(*assignmentDetails.Resources[0].ActionControl.ResourceCreated.ID).ToNot(BeNil())
			Expect(*assignmentDetails.Resources[0].Target.ID).ToNot(BeNil())
			Expect(*assignmentDetails.Resources[0].Target.ID).To(Equal(testTargetAccountID))
			Expect(*assignmentDetails.Resources[0].Target.Type).ToNot(BeNil())
			Expect(*assignmentDetails.Resources[0].Target.Type).To(Equal(testTargetType))
			testActionControlAssignmentETag = response.GetHeaders().Get(etagHeader)
			testActionControlAssignmentId = *assignmentDetails.ID
		})
	})

	Describe(`UpdateActionControlAssignment - update an action control assignment by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		Expect(testActionControlAssignmentETag).To(Not(BeNil()))
		It(`UpdateActionControlAssignment(updateActionControlAssignmentOptions *UpdateActionControlAssignmentOptions))`, func() {
			updatePolicyAssignmentOptions := &iampolicymanagementv1.UpdateActionControlAssignmentOptions{
				AssignmentID:    core.StringPtr(testActionControlAssignmentId),
				TemplateVersion: core.StringPtr(testActionControlUpdateTemplateVersion),
				IfMatch:         core.StringPtr(testActionControlAssignmentETag),
			}

			actionControlAssignment, response, err := service.UpdateActionControlAssignment(updatePolicyAssignmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(*actionControlAssignment.Resources[0].ActionControl.ResourceCreated.ID).ToNot(BeNil())
			Expect(*actionControlAssignment.ID).To(Equal(testActionControlAssignmentId))
		})
	})

	Describe(`ListActionControlAssignments - Get action control template assignments by attributes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListActionControlAssignments(listActionControlAssignmentsOptions *ListActionControlAssignmentsOptions)`, func() {
			listActionControlAssignmentsOptions := &iampolicymanagementv1.ListActionControlAssignmentsOptions{
				AccountID:      core.StringPtr(testAccountID),
				AcceptLanguage: core.StringPtr("default"),
			}

			templateAssignmentCollection, response, err := service.ListActionControlAssignments(listActionControlAssignmentsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateAssignmentCollection).ToNot(BeNil())
			var assignmentDetails = templateAssignmentCollection.Assignments[0]

			Expect(assignmentDetails).ToNot(BeNil())
			Expect(assignmentDetails.Template.ID).ToNot(BeNil())
			Expect(assignmentDetails.Target.Type).ToNot(BeNil())
			Expect(assignmentDetails.Template.Version).ToNot(BeNil())
			Expect(assignmentDetails.Target.ID).ToNot(BeNil())
			Expect(assignmentDetails.Status).ToNot(BeNil())
			Expect(assignmentDetails.AccountID).ToNot(BeNil())
			Expect(assignmentDetails.Resources).ToNot(BeNil())
			Expect(assignmentDetails.CreatedAt).ToNot(BeNil())
			Expect(assignmentDetails.CreatedByID).ToNot(BeNil())
			Expect(assignmentDetails.LastModifiedAt).ToNot(BeNil())
			Expect(assignmentDetails.LastModifiedByID).ToNot(BeNil())
			Expect(assignmentDetails.Href).ToNot(BeNil())
		})
	})

	Describe(`GetActionControlAssignment - Retrieve an action control assignment by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetActionControlAssignment(getActionControlAssignmentOptions *GetActionControlAssignmentOptions)`, func() {
			getActionControlAssignmentOptions := &iampolicymanagementv1.GetActionControlAssignmentOptions{
				AssignmentID: core.StringPtr(testActionControlAssignmentId),
			}

			assignmentDetails, response, err := service.GetActionControlAssignment(getActionControlAssignmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(assignmentDetails).ToNot(BeNil())
			Expect(assignmentDetails.Template.ID).ToNot(BeNil())
			Expect(assignmentDetails.Target.Type).ToNot(BeNil())
			Expect(assignmentDetails.Target.ID).ToNot(BeNil())
			Expect(assignmentDetails.Status).ToNot(BeNil())
			Expect(assignmentDetails.AccountID).ToNot(BeNil())
			Expect(assignmentDetails.CreatedAt).ToNot(BeNil())
			Expect(assignmentDetails.CreatedByID).ToNot(BeNil())
			Expect(assignmentDetails.LastModifiedAt).ToNot(BeNil())
			Expect(assignmentDetails.LastModifiedByID).ToNot(BeNil())
			Expect(assignmentDetails.Href).ToNot(BeNil())
		})
	})

	Describe(`DeleteActionControlAssignment - Delete an action control assignment by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteActionControlAssignment(deleteActionControlAssignmentOptions *DeleteActionControlAssignmentOptions)`, func() {
			deleteActionControlAssignmentOptions := &iampolicymanagementv1.DeleteActionControlAssignmentOptions{
				AssignmentID: core.StringPtr(testActionControlAssignmentId),
			}

			response, err := service.DeleteActionControlAssignment(deleteActionControlAssignmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteTestActionControlTemplate - Delete an action control template by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTestActionControlTemplate(deleteActionControlTemplateOptions *DeleteActionControlTemplateOptions)`, func() {
			deleteActionControlTemplateOptions := &iampolicymanagementv1.DeleteActionControlTemplateOptions{
				ActionControlTemplateID: &testActionControlTemplateID,
			}

			response, err := service.DeleteActionControlTemplate(deleteActionControlTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	// clean up all test groups
	AfterSuite(func() {
		if !configLoaded {
			return
		}

		fmt.Fprintf(GinkgoWriter, "Cleaning up test groups...\n")

		// list all policies in the account
		policyOptions := service.NewListPoliciesOptions(testAccountID)
		policyOptions.SetIamID(testUserId)
		policyResult, policyDetailedResponse, err := service.ListPolicies(policyOptions)
		Expect(err).To(BeNil())
		Expect(policyDetailedResponse.StatusCode).To(Equal(200))

		for _, policy := range policyResult.Policies {

			// delete the test policy (or any test policy older than 5 minutes)
			createdAt, err := time.Parse(time.RFC3339, policy.CreatedAt.String())
			if err != nil {
				fmt.Fprintf(GinkgoWriter, "time.Parse error occurred: %v\n", err)
				fmt.Fprintf(GinkgoWriter, "Cleanup of policy (%v) failed\n", *policy.ID)
				continue
			}
			fiveMinutesAgo := time.Now().Add(-(time.Duration(5) * time.Minute))
			if strings.Contains(*policy.Href, "v2/policies") {
				if *policy.ID == testV2PolicyId || createdAt.Before(fiveMinutesAgo) {
					options := service.NewDeleteV2PolicyOptions(*policy.ID)
					detailedResponse, err := service.DeleteV2Policy(options)
					Expect(err).To(BeNil())
					Expect(detailedResponse.StatusCode).To(Equal(204))
				}
			} else {
				if *policy.ID == testPolicyId || createdAt.Before(fiveMinutesAgo) {
					options := service.NewDeletePolicyOptions(*policy.ID)
					detailedResponse, err := service.DeletePolicy(options)
					Expect(err).To(BeNil())
					Expect(detailedResponse.StatusCode).To(Equal(204))
				}
			}
		}

		// List all custom roles in the account
		roleOptions := service.NewListRolesOptions()
		roleOptions.SetAccountID(testAccountID)
		roleResult, roleDetailedResponse, err := service.ListRoles(roleOptions)
		Expect(err).To(BeNil())
		Expect(roleDetailedResponse.StatusCode).To(Equal(200))

		for _, role := range roleResult.CustomRoles {

			// delete the role (or any test role older than 5 minutes)
			createdAt, err := time.Parse(time.RFC3339, role.CreatedAt.String())
			if err != nil {
				fmt.Fprintf(GinkgoWriter, "time.Parse error occurred: %v\n", err)
				fmt.Fprintf(GinkgoWriter, "Cleanup of role (%v) failed\n", *role.ID)
				continue
			}
			fiveMinutesAgo := time.Now().Add(-(time.Duration(5) * time.Minute))

			if *role.ID == testCustomRoleId || createdAt.Before(fiveMinutesAgo) {
				options := service.NewDeleteRoleOptions(*role.ID)
				detailedResponse, err := service.DeleteRole(options)
				Expect(err).To(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(204))
			}
		}

		fmt.Fprintf(GinkgoWriter, "Cleanup finished!\n")
	})
})
