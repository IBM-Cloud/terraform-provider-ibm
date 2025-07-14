//go:build examples

/**
 * (C) Copyright IBM Corp. 2023.
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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the IAM Policy Management service.
//
// The following configuration properties are assumed to be defined:
//
// IAM_POLICY_MANAGEMENT_URL=<service url>
// IAM_POLICY_MANAGEMENT_AUTH_TYPE=iam
// IAM_POLICY_MANAGEMENT_AUTH_URL=<IAM token service URL - omit this if using the production environment>
// IAM_POLICY_MANAGEMENT_APIKEY=<YOUR_APIKEY>
// IAM_POLICY_MANAGEMENT_TEST_ACCOUNT_ID=<YOUR_ACCOUNT_ID>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of config file>
//
// Location of our config file.

var _ = Describe(`IamPolicyManagementV1 Examples Tests`, func() {
	const externalConfigFile = "../iam_policy_management.env"

	var (
		// TODO: Align
		iamPolicyManagementService *iampolicymanagementv1.IamPolicyManagementV1
		config                     map[string]string
		configLoaded               bool = false

		exampleUserID                           = "IBMid-user1"
		exampleServiceName                      = "iam-groups"
		exampleAccountID                        string
		examplePolicyID                         string
		examplePolicyETag                       string
		exampleCustomRoleID                     string
		exampleCustomRoleETag                   string
		examplePolicyTemplateName               = "PolicySampleTemplateTest"
		examplePolicyTemplateID                 string
		examplePolicyTemplateETag               string
		examplePolicyTemplateBaseVersion        string
		examplePolicyTemplateVersion            string
		testPolicyAssignmentId                  string
		exampleAssignmentPolicyID               string
		exampleTargetAccountID                  string = ""
		examplePolicyAssignmentETag             string = ""
		exampleAccountSettingsETag              string
		exampleETagHeader                       string = "ETag"
		exampleActionControlTemplateID          string
		exampleActionControlTemplateBaseVersion string
		exampleActionControlTemplateETag        string = ""
		exampleActionControlTemplateVersion     string
		exampleActionControlTemplateName               = "ActionControlTemplateGoSDKTest"
		exampleActionControlAssignmentETag      string = ""
		exampleActionControlAssignmentId        string
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
			config, err = core.GetServiceProperties(iampolicymanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			exampleAccountID = config["TEST_ACCOUNT_ID"]
			exampleTargetAccountID = config["TEST_TARGET_ACCOUNT_ID"]

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

			iamPolicyManagementServiceOptions := &iampolicymanagementv1.IamPolicyManagementV1Options{}

			iamPolicyManagementService, err = iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(iamPolicyManagementServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(iamPolicyManagementService).ToNot(BeNil())
		})
	})

	Describe(`IamPolicyManagementV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePolicy request example`, func() {
			fmt.Println("\nCreatePolicy() result:")
			// begin-create_policy

			subjectAttribute := &iampolicymanagementv1.SubjectAttribute{
				Name:  core.StringPtr("iam_id"),
				Value: &exampleUserID,
			}
			policySubjects := &iampolicymanagementv1.PolicySubject{
				Attributes: []iampolicymanagementv1.SubjectAttribute{*subjectAttribute},
			}
			policyRoles := &iampolicymanagementv1.PolicyRole{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::role:Viewer"),
			}
			accountIDResourceAttribute := &iampolicymanagementv1.ResourceAttribute{
				Name:     core.StringPtr("accountId"),
				Value:    core.StringPtr(exampleAccountID),
				Operator: core.StringPtr("stringEquals"),
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.ResourceAttribute{
				Name:     core.StringPtr("serviceType"),
				Value:    core.StringPtr("service"),
				Operator: core.StringPtr("stringEquals"),
			}
			policyResourceTag := &iampolicymanagementv1.ResourceTag{
				Name:     core.StringPtr("project"),
				Value:    core.StringPtr("prototype"),
				Operator: core.StringPtr("stringEquals"),
			}
			policyResources := &iampolicymanagementv1.PolicyResource{
				Attributes: []iampolicymanagementv1.ResourceAttribute{
					*accountIDResourceAttribute, *serviceNameResourceAttribute},
				Tags: []iampolicymanagementv1.ResourceTag{*policyResourceTag},
			}

			options := iamPolicyManagementService.NewCreatePolicyOptions(
				"access",
				[]iampolicymanagementv1.PolicySubject{*policySubjects},
				[]iampolicymanagementv1.PolicyRole{*policyRoles},
				[]iampolicymanagementv1.PolicyResource{*policyResources},
			)

			policy, response, err := iamPolicyManagementService.CreatePolicy(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			examplePolicyID = *policy.ID
			fmt.Println(string(b))

			// end-create_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policy).ToNot(BeNil())
		})
		It(`GetPolicy request example`, func() {
			fmt.Println("\nGetPolicy() result:")
			// begin-get_policy

			options := iamPolicyManagementService.NewGetPolicyOptions(
				examplePolicyID,
			)

			policy, response, err := iamPolicyManagementService.GetPolicy(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			examplePolicyETag = response.GetHeaders().Get("ETag")
			fmt.Println(string(b))

			// end-get_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
		})
		It(`ReplacePolicy request example`, func() {
			fmt.Println("\nReplacePolicy() result:")
			// begin-replace_policy

			subjectAttribute := &iampolicymanagementv1.SubjectAttribute{
				Name:  core.StringPtr("iam_id"),
				Value: &exampleUserID,
			}
			policySubjects := &iampolicymanagementv1.PolicySubject{
				Attributes: []iampolicymanagementv1.SubjectAttribute{*subjectAttribute},
			}
			accountIDResourceAttribute := &iampolicymanagementv1.ResourceAttribute{
				Name:     core.StringPtr("accountId"),
				Value:    core.StringPtr(exampleAccountID),
				Operator: core.StringPtr("stringEquals"),
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.ResourceAttribute{
				Name:     core.StringPtr("serviceType"),
				Value:    core.StringPtr("service"),
				Operator: core.StringPtr("stringEquals"),
			}
			policyResourceTag := &iampolicymanagementv1.ResourceTag{
				Name:     core.StringPtr("project"),
				Value:    core.StringPtr("prototype"),
				Operator: core.StringPtr("stringEquals"),
			}
			policyResources := &iampolicymanagementv1.PolicyResource{
				Attributes: []iampolicymanagementv1.ResourceAttribute{
					*accountIDResourceAttribute, *serviceNameResourceAttribute},
				Tags: []iampolicymanagementv1.ResourceTag{*policyResourceTag},
			}
			updatedPolicyRoles := &iampolicymanagementv1.PolicyRole{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::role:Editor"),
			}

			options := iamPolicyManagementService.NewReplacePolicyOptions(
				examplePolicyID,
				examplePolicyETag,
				"access",
				[]iampolicymanagementv1.PolicySubject{*policySubjects},
				[]iampolicymanagementv1.PolicyRole{*updatedPolicyRoles},
				[]iampolicymanagementv1.PolicyResource{*policyResources},
			)

			policy, response, err := iamPolicyManagementService.ReplacePolicy(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			examplePolicyETag = response.GetHeaders().Get("ETag")
			fmt.Println(string(b))

			// end-replace_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
		})
		It(`UpdatePolicy request example`, func() {
			fmt.Println("\nUpdatePolicyState() result:")
			// begin-update_policy_state

			options := iamPolicyManagementService.NewUpdatePolicyStateOptions(
				examplePolicyID,
				examplePolicyETag,
			)

			options.SetState("active")

			policy, response, err := iamPolicyManagementService.UpdatePolicyState(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-update_policy_state

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())

		})
		It(`ListPolicies request example`, func() {
			fmt.Println("\nListPolicies() result:")
			// begin-list_policies

			options := iamPolicyManagementService.NewListPoliciesOptions(
				exampleAccountID,
			)
			options.SetIamID(exampleUserID)
			options.SetFormat("include_last_permit")

			policyList, response, err := iamPolicyManagementService.ListPolicies(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyList, "", "  ")
			fmt.Println(string(b))

			// end-list_policies

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyList).ToNot(BeNil())

		})
		It(`DeletePolicy request example`, func() {
			// begin-delete_policy

			options := iamPolicyManagementService.NewDeletePolicyOptions(
				examplePolicyID,
			)

			response, err := iamPolicyManagementService.DeletePolicy(options)
			if err != nil {
				panic(err)
			}

			// end-delete_policy
			fmt.Printf("\nDeletePolicy() response status code: %d\n", response.StatusCode)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`CreateV2Policy request example`, func() {
			fmt.Println("\nCreateV2Policy() result:")
			// begin-create_v2_policy

			subjectAttribute := &iampolicymanagementv1.V2PolicySubjectAttribute{
				Key:      core.StringPtr("iam_id"),
				Operator: core.StringPtr("stringEquals"),
				Value:    &exampleUserID,
			}
			policySubject := &iampolicymanagementv1.V2PolicySubject{
				Attributes: []iampolicymanagementv1.V2PolicySubjectAttribute{*subjectAttribute},
			}
			policyRole := &iampolicymanagementv1.Roles{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::role:Viewer"),
			}
			v2PolicyGrant := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*policyRole},
			}
			v2PolicyControl := &iampolicymanagementv1.Control{
				Grant: v2PolicyGrant,
			}
			accountIDResourceAttribute := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("accountId"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr(exampleAccountID),
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceType"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("service"),
			}
			policyResourceTag := &iampolicymanagementv1.V2PolicyResourceTag{
				Key:      core.StringPtr("project"),
				Value:    core.StringPtr("prototype"),
				Operator: core.StringPtr("stringEquals"),
			}
			policyResource := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{
					*accountIDResourceAttribute, *serviceNameResourceAttribute},
				Tags: []iampolicymanagementv1.V2PolicyResourceTag{*policyResourceTag},
			}
			weeklyConditionAttribute := &iampolicymanagementv1.NestedCondition{
				Key:      core.StringPtr("{{environment.attributes.day_of_week}}"),
				Operator: core.StringPtr("dayOfWeekAnyOf"),
				Value:    []string{"1+00:00", "2+00:00", "3+00:00", "4+00:00", "5+00:00"},
			}
			startConditionAttribute := &iampolicymanagementv1.NestedCondition{
				Key:      core.StringPtr("{{environment.attributes.current_time}}"),
				Operator: core.StringPtr("timeGreaterThanOrEquals"),
				Value:    core.StringPtr("09:00:00+00:00"),
			}
			endConditionAttribute := &iampolicymanagementv1.NestedCondition{
				Key:      core.StringPtr("{{environment.attributes.current_time}}"),
				Operator: core.StringPtr("timeLessThanOrEquals"),
				Value:    core.StringPtr("17:00:00+00:00"),
			}
			policyRule := &iampolicymanagementv1.V2PolicyRule{
				Operator: core.StringPtr("and"),
				Conditions: []iampolicymanagementv1.NestedConditionIntf{
					weeklyConditionAttribute, startConditionAttribute, endConditionAttribute},
			}

			options := iamPolicyManagementService.NewCreateV2PolicyOptions(
				v2PolicyControl,
				"access",
			)
			options.SetSubject(policySubject)
			options.SetResource(policyResource)
			options.SetRule(policyRule)
			options.SetPattern(*core.StringPtr("time-based-conditions:weekly:custom-hours"))

			policy, response, err := iamPolicyManagementService.CreateV2Policy(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			examplePolicyID = *policy.ID
			fmt.Println(string(b))

			// end-create_v2_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policy).ToNot(BeNil())
		})
		It(`GetV2Policy request example`, func() {
			fmt.Println("\nGetV2Policy() result:")
			// begin-get_v2_policy

			options := iamPolicyManagementService.NewGetV2PolicyOptions(
				examplePolicyID,
			)

			policy, response, err := iamPolicyManagementService.GetV2Policy(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			examplePolicyETag = response.GetHeaders().Get("ETag")
			fmt.Println(string(b))

			// end-get_v2_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
		})
		It(`ReplaceV2Policy request example`, func() {
			fmt.Println("\nReplaceV2Policy() result:")
			// begin-replace_v2_policy

			subjectAttribute := &iampolicymanagementv1.V2PolicySubjectAttribute{
				Key:      core.StringPtr("iam_id"),
				Operator: core.StringPtr("stringEquals"),
				Value:    &exampleUserID,
			}
			policySubject := &iampolicymanagementv1.V2PolicySubject{
				Attributes: []iampolicymanagementv1.V2PolicySubjectAttribute{*subjectAttribute},
			}
			updatedPolicyRole := &iampolicymanagementv1.Roles{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::role:Editor"),
			}
			v2PolicyGrant := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*updatedPolicyRole},
			}
			v2PolicyControl := &iampolicymanagementv1.Control{
				Grant: v2PolicyGrant,
			}
			accountIDResourceAttribute := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("accountId"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr(exampleAccountID),
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceType"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("service"),
			}
			policyResource := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{
					*accountIDResourceAttribute, *serviceNameResourceAttribute},
			}

			options := iamPolicyManagementService.NewReplaceV2PolicyOptions(
				examplePolicyID,
				examplePolicyETag,
				v2PolicyControl,
				"access",
			)
			weeklyConditionAttribute := &iampolicymanagementv1.NestedCondition{
				Key:      core.StringPtr("{{environment.attributes.day_of_week}}"),
				Operator: core.StringPtr("dayOfWeekAnyOf"),
				Value:    []string{"1+00:00", "2+00:00", "3+00:00", "4+00:00"},
			}
			startConditionAttribute := &iampolicymanagementv1.NestedCondition{
				Key:      core.StringPtr("{{environment.attributes.current_time}}"),
				Operator: core.StringPtr("timeGreaterThanOrEquals"),
				Value:    core.StringPtr("09:00:00+00:00"),
			}
			endConditionAttribute := &iampolicymanagementv1.NestedCondition{
				Key:      core.StringPtr("{{environment.attributes.current_time}}"),
				Operator: core.StringPtr("timeLessThanOrEquals"),
				Value:    core.StringPtr("17:00:00+00:00"),
			}
			policyRule := &iampolicymanagementv1.V2PolicyRule{
				Operator: core.StringPtr("and"),
				Conditions: []iampolicymanagementv1.NestedConditionIntf{
					weeklyConditionAttribute, startConditionAttribute, endConditionAttribute},
			}
			options.SetRule(policyRule)
			options.SetPattern(*core.StringPtr("time-based-conditions:weekly:custom-hours"))
			options.SetSubject(policySubject)
			options.SetResource(policyResource)

			policy, response, err := iamPolicyManagementService.ReplaceV2Policy(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-replace_v2_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
		})
		It(`ListV2Policies request example`, func() {
			fmt.Println("\nListV2Policies() result:")
			// begin-list_v2_policies

			options := iamPolicyManagementService.NewListV2PoliciesOptions(
				exampleAccountID,
			)
			options.SetIamID(exampleUserID)
			options.SetFormat("include_last_permit")
			options.SetSort("-id")

			policyList, response, err := iamPolicyManagementService.ListV2Policies(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyList, "", "  ")
			fmt.Println(string(b))

			// end-list_v2_policies

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyList).ToNot(BeNil())

		})
		It(`DeleteV2Policy request example`, func() {
			// begin-delete_v2_policy

			options := iamPolicyManagementService.NewDeleteV2PolicyOptions(
				examplePolicyID,
			)

			response, err := iamPolicyManagementService.DeleteV2Policy(options)
			if err != nil {
				panic(err)
			}

			// end-delete_delete_v2_policypolicy
			fmt.Printf("\nDeleteV2Policy() response status code: %d\n", response.StatusCode)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`CreateRole request example`, func() {
			fmt.Println("\nCreateRole() result:")
			// begin-create_role

			options := iamPolicyManagementService.NewCreateRoleOptions(
				"IAM Groups read access",
				[]string{"iam-groups.groups.read"},
				"ExampleRoleIAMGroups",
				exampleAccountID,
				exampleServiceName,
			)

			customRole, response, err := iamPolicyManagementService.CreateRole(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(customRole, "", "  ")
			exampleCustomRoleID = *customRole.ID
			fmt.Println(string(b))

			// end-create_role

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(customRole).ToNot(BeNil())
		})
		It(`GetRole request example`, func() {
			fmt.Println("\nGetRole() result:")
			// begin-get_role

			options := iamPolicyManagementService.NewGetRoleOptions(
				exampleCustomRoleID,
			)

			customRole, response, err := iamPolicyManagementService.GetRole(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(customRole, "", "  ")
			exampleCustomRoleETag = response.Headers.Get("ETag")
			fmt.Println(string(b))

			// end-get_role

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customRole).ToNot(BeNil())
		})
		It(`ReplaceRole request example`, func() {
			fmt.Println("\nReplaceRole() result:")
			// begin-replace_role

			updatedRoleActions := []string{"iam-groups.groups.read", "iam-groups.groups.list"}

			options := iamPolicyManagementService.NewReplaceRoleOptions(
				exampleCustomRoleID,
				exampleCustomRoleETag,
				"ExampleRoleIAMGroups",
				updatedRoleActions,
			)

			customRole, response, err := iamPolicyManagementService.ReplaceRole(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(customRole, "", "  ")
			fmt.Println(string(b))

			// end-replace_role

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(customRole).ToNot(BeNil())

		})
		It(`ListRoles request example`, func() {
			fmt.Println("\nListRoles() result:")
			// begin-list_roles

			options := iamPolicyManagementService.NewListRolesOptions()
			options.SetAccountID(exampleAccountID)

			roleList, response, err := iamPolicyManagementService.ListRoles(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(roleList, "", "  ")
			fmt.Println(string(b))

			// end-list_roles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(roleList).ToNot(BeNil())

		})
		It(`DeleteRole request example`, func() {
			// begin-delete_role

			options := iamPolicyManagementService.NewDeleteRoleOptions(
				exampleCustomRoleID,
			)

			response, err := iamPolicyManagementService.DeleteRole(options)
			if err != nil {
				panic(err)
			}

			// end-delete_role
			fmt.Printf("\nDeleteRole() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ListPolicyTemplates request example`, func() {
			fmt.Println("\nListPolicyTemplates() result:")
			// begin-list_policy_templates

			listPolicyTemplatesOptions := iamPolicyManagementService.NewListPolicyTemplatesOptions(
				exampleAccountID,
			)

			policyTemplateCollection, response, err := iamPolicyManagementService.ListPolicyTemplates(listPolicyTemplatesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplateCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_policy_templates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplateCollection).ToNot(BeNil())
		})
		It(`CreatePolicyS2STemplate request example`, func() {
			fmt.Println("\nCreatePolicyTemplate() result:")
			// begin-create_policy_template

			policyRole := &iampolicymanagementv1.Roles{
				RoleID: core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Writer"),
			}
			v2PolicyGrant := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{*policyRole},
			}
			v2PolicyControl := &iampolicymanagementv1.Control{
				Grant: v2PolicyGrant,
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("cloud-object-storage"),
			}

			policyResource := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{
					*serviceNameResourceAttribute},
			}
			v2PolicySubjectAttributeModel := &iampolicymanagementv1.V2PolicySubjectAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("compliance"),
			}

			policySubject := &iampolicymanagementv1.V2PolicySubject{
				Attributes: []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel},
			}

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("authorization"),
				Description: core.StringPtr("Test Template"),
				Resource:    policyResource,
				Control:     v2PolicyControl,
				Subject:     policySubject,
			}

			createPolicyTemplateOptions := iamPolicyManagementService.NewCreatePolicyTemplateOptions(
				examplePolicyTemplateName,
				exampleAccountID,
				templatePolicyModel,
			)

			policyTemplate, response, err := iamPolicyManagementService.CreatePolicyTemplate(createPolicyTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplate, "", "  ")
			examplePolicyTemplateID = *policyTemplate.ID
			examplePolicyTemplateBaseVersion = *policyTemplate.Version
			examplePolicyTemplateETag = response.GetHeaders().Get("ETag")
			fmt.Println(string(b))

			// end-create_policy_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyTemplate).ToNot(BeNil())
		})

		It(`GetPolicyTemplate request example`, func() {
			fmt.Println("\nGetPolicyTemplate() result:")
			// begin-get_policy_template

			getPolicyTemplateOptions := iamPolicyManagementService.NewGetPolicyTemplateOptions(
				examplePolicyTemplateID,
			)

			policyTemplate, response, err := iamPolicyManagementService.GetPolicyTemplate(getPolicyTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplate, "", "  ")
			fmt.Println(string(b))

			// end-get_policy_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplate.AccountID).ToNot(BeNil())
			Expect(policyTemplate.Version).ToNot(BeNil())
			Expect(policyTemplate.Name).ToNot(BeNil())
			Expect(policyTemplate.Policy).ToNot(BeNil())
		})

		It(`CreatePolicyS2STemplateVersion request example`, func() {
			fmt.Println("\nCreatePolicyTemplateVersion() result:")
			// begin-create_policy_template_version
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

			createPolicyTemplateVersionOptions := &iampolicymanagementv1.CreatePolicyTemplateVersionOptions{
				Policy:           templatePolicyModel,
				PolicyTemplateID: core.StringPtr(examplePolicyTemplateID),
				Description:      core.StringPtr("Test PolicySampleTemplate"),
				Committed:        core.BoolPtr(true),
			}

			policyTemplate, response, err := iamPolicyManagementService.CreatePolicyTemplateVersion(createPolicyTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplate, "", "  ")
			examplePolicyTemplateVersion = *policyTemplate.Version
			fmt.Println(string(b))

			// end-create_policy_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyTemplate).ToNot(BeNil())
		})

		It(`ListPolicyTemplateVersions request example`, func() {
			fmt.Println("\nListPolicyTemplateVersions() result:")
			// begin-list_policy_template_versions

			listPolicyTemplateVersionsOptions := iamPolicyManagementService.NewListPolicyTemplateVersionsOptions(
				examplePolicyTemplateID,
			)

			policyTemplateVersionsCollection, response, err := iamPolicyManagementService.ListPolicyTemplateVersions(listPolicyTemplateVersionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplateVersionsCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_policy_template_versions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplateVersionsCollection).ToNot(BeNil())
		})

		It(`ReplacePolicyS2STemplate request example`, func() {
			fmt.Println("\nReplacePolicyTemplate() result:")
			// begin-replace_policy_template
			v2PolicyGrant := &iampolicymanagementv1.Grant{
				Roles: []iampolicymanagementv1.Roles{
					{core.StringPtr("crn:v1:bluemix:public:iam::::serviceRole:Reader")},
				},
			}

			v2PolicyControl := &iampolicymanagementv1.Control{
				Grant: v2PolicyGrant,
			}
			serviceNameResourceAttribute := &iampolicymanagementv1.V2PolicyResourceAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("appid"),
			}
			policyResource := &iampolicymanagementv1.V2PolicyResource{
				Attributes: []iampolicymanagementv1.V2PolicyResourceAttribute{
					*serviceNameResourceAttribute},
			}

			v2PolicySubjectAttributeModel := &iampolicymanagementv1.V2PolicySubjectAttribute{
				Key:      core.StringPtr("serviceName"),
				Operator: core.StringPtr("stringEquals"),
				Value:    core.StringPtr("compliance"),
			}

			policySubject := &iampolicymanagementv1.V2PolicySubject{
				Attributes: []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel},
			}

			templatePolicyModel := &iampolicymanagementv1.TemplatePolicy{
				Type:        core.StringPtr("authorization"),
				Description: core.StringPtr("Test Template v2"),
				Resource:    policyResource,
				Control:     v2PolicyControl,
				Subject:     policySubject,
			}

			replacePolicyTemplateOptions := iamPolicyManagementService.NewReplacePolicyTemplateOptions(
				examplePolicyTemplateID,
				examplePolicyTemplateBaseVersion,
				examplePolicyTemplateETag,
				templatePolicyModel,
			)

			policyTemplate, response, err := iamPolicyManagementService.ReplacePolicyTemplate(replacePolicyTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplate, "", "  ")
			examplePolicyTemplateVersion = *policyTemplate.Version
			examplePolicyTemplateETag = response.GetHeaders().Get("ETag")
			fmt.Println(string(b))

			// end-replace_policy_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplate).ToNot(BeNil())
		})
		It(`GetPolicyTemplateVersion request example`, func() {
			fmt.Println("\nGetPolicyTemplateVersion() result:")
			// begin-get_policy_template_version

			getPolicyTemplateVersionOptions := iamPolicyManagementService.NewGetPolicyTemplateVersionOptions(
				examplePolicyTemplateID,
				examplePolicyTemplateVersion,
			)

			policyTemplate, response, err := iamPolicyManagementService.GetPolicyTemplateVersion(getPolicyTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplate, "", "  ")
			fmt.Println(string(b))

			// end-get_policy_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyTemplate).ToNot(BeNil())
		})

		It(`CommitPolicyTemplate request example`, func() {
			fmt.Println("\nCommitPolicyTemplate() result:")
			// begin-commit_policy_template

			commitPolicyTemplateOptions := iamPolicyManagementService.NewCommitPolicyTemplateOptions(
				examplePolicyTemplateID,
				examplePolicyTemplateBaseVersion,
			)

			response, err := iamPolicyManagementService.CommitPolicyTemplate(commitPolicyTemplateOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from CommitPolicyTemplate(): %d\n", response.StatusCode)
			}

			// end-commit_policy_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

		It(`CreatePolicyAssignments request example`, func() {
			fmt.Println("\nCreatePolicyTemplateAssignment() result:")
			// begin-create_policy_template_assignment
			template := iampolicymanagementv1.AssignmentTemplateDetails{
				ID:      &examplePolicyTemplateID,
				Version: &examplePolicyTemplateBaseVersion,
			}
			templates := []iampolicymanagementv1.AssignmentTemplateDetails{
				template,
			}

			target := &iampolicymanagementv1.AssignmentTargetDetails{
				Type: core.StringPtr("Account"),
				ID:   &exampleTargetAccountID,
			}

			createPolicyTemplateVersionOptions := &iampolicymanagementv1.CreatePolicyTemplateAssignmentOptions{
				Version:   core.StringPtr("1.0"),
				Target:    target,
				Templates: templates,
			}

			policyAssignment, response, err := iamPolicyManagementService.CreatePolicyTemplateAssignment(createPolicyTemplateVersionOptions)
			b, _ := json.MarshalIndent(policyAssignment, "", "  ")
			fmt.Println(string(b))

			var assignmentDetails = policyAssignment.Assignments[0]
			// end-create_policy_template_assignment
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			examplePolicyAssignmentETag = response.GetHeaders().Get(exampleETagHeader)
			testPolicyAssignmentId = *assignmentDetails.ID
		})

		It(`UpdatePolicyAssignment request example))`, func() {
			// begin-update_policy_assignment
			updatePolicyAssignmentOptions := iamPolicyManagementService.NewUpdatePolicyAssignmentOptions(
				testPolicyAssignmentId,
				"1.0",
				examplePolicyAssignmentETag,
				examplePolicyTemplateVersion,
			)

			policyAssignment, response, err := iamPolicyManagementService.UpdatePolicyAssignment(updatePolicyAssignmentOptions)
			b, _ := json.MarshalIndent(policyAssignment, "", "  ")
			fmt.Println(string(b))
			// end-update_policy_assignment
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			testPolicyAssignmentId = *policyAssignment.ID
		})

		It(`ListPolicyAssignments request example`, func() {
			fmt.Println("\nListPolicyAssignments() result:")
			// begin-list_policy_assignments

			listPolicyAssignmentsOptions := iamPolicyManagementService.NewListPolicyAssignmentsOptions(
				"1.0",
				exampleAccountID,
			)

			policyTemplateAssignmentCollection, response, err := iamPolicyManagementService.ListPolicyAssignments(listPolicyAssignmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyTemplateAssignmentCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_policy_assignments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
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

		It(`GetPolicyAssignment request example`, func() {
			fmt.Println("\nGetPolicyAssignment() result:")
			// begin-get_policy_assignment

			getPolicyAssignmentOptions := iamPolicyManagementService.NewGetPolicyAssignmentOptions(
				testPolicyAssignmentId,
				"1.0",
			)

			policyAssignmentRecord, response, err := iamPolicyManagementService.GetPolicyAssignment(getPolicyAssignmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policyAssignmentRecord, "", "  ")
			fmt.Println(string(b))

			// end-get_policy_assignment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policyAssignmentRecord).ToNot(BeNil())
			var assignmentDetails = policyAssignmentRecord.(*iampolicymanagementv1.PolicyTemplateAssignmentItems)
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
			exampleAssignmentPolicyID = *assignmentDetails.Resources[0].Policy.ResourceCreated.ID
		})

		It(`GetV2Policy to get Template meta data request example`, func() {
			fmt.Println("\nGetV2Policy() result:")
			// begin-get_v2_policy template metadata

			options := iamPolicyManagementService.NewGetV2PolicyOptions(
				exampleAssignmentPolicyID,
			)

			policy, response, err := iamPolicyManagementService.GetV2Policy(options)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(policy, "", "  ")
			fmt.Println(string(b))

			// end-get_v2_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
			Expect(policy.Template).ToNot(BeNil())
			Expect(policy.Template.ID).ToNot(BeNil())
			Expect(policy.Template.Version).ToNot(BeNil())
			Expect(policy.Template.AssignmentID).ToNot(BeNil())
		})

		It(`DeletePolicyAssignment request example)`, func() {
			// begin-delete_policy_assignment
			deletePolicyAssignmentOptions := iamPolicyManagementService.NewDeletePolicyAssignmentOptions(
				testPolicyAssignmentId,
			)

			response, err := iamPolicyManagementService.DeletePolicyAssignment(deletePolicyAssignmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

		It(`DeletePolicyTemplateVersion request example`, func() {
			// begin-delete_policy_template_version

			deletePolicyTemplateVersionOptions := iamPolicyManagementService.NewDeletePolicyTemplateVersionOptions(
				examplePolicyTemplateID,
				examplePolicyTemplateVersion,
			)

			response, err := iamPolicyManagementService.DeletePolicyTemplateVersion(deletePolicyTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeletePolicyTemplateVersion(): %d\n", response.StatusCode)
			}

			// end-delete_policy_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

		It(`DeletePolicyTemplate request example`, func() {
			// begin-delete_policy_template

			deletePolicyTemplateOptions := iamPolicyManagementService.NewDeletePolicyTemplateOptions(
				examplePolicyTemplateID,
			)

			response, err := iamPolicyManagementService.DeletePolicyTemplate(deletePolicyTemplateOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeletePolicyTemplate(): %d\n", response.StatusCode)
			}

			// end-delete_policy_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

		It(`GetSettings request example`, func() {
			fmt.Println("\nGetSettings() result:")
			// begin-get_settings

			getSettingsOptions := iamPolicyManagementService.NewGetSettingsOptions(
				exampleAccountID,
			)

			accountSettingsAccessManagement, response, err := iamPolicyManagementService.GetSettings(getSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettingsAccessManagement, "", "  ")
			exampleAccountSettingsETag = response.GetHeaders().Get("ETag")
			fmt.Println(string(b))

			// end-get_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettingsAccessManagement).ToNot(BeNil())
		})

		It(`UpdateSettings request example`, func() {
			fmt.Println("\nUpdateSettings() result:")
			// begin-update_settings

			updateSettingsOptions := iamPolicyManagementService.NewUpdateSettingsOptions(
				exampleAccountID,
				exampleAccountSettingsETag,
			)
			identityTypesBase := &iampolicymanagementv1.IdentityTypesBase{
				State:                   core.StringPtr("monitor"),
				ExternalAllowedAccounts: []string{},
			}

			identityTypes := &iampolicymanagementv1.IdentityTypesPatch{
				User:      identityTypesBase,
				ServiceID: identityTypesBase,
				Service:   identityTypesBase,
			}

			externalAccountIdentityInteraction := &iampolicymanagementv1.ExternalAccountIdentityInteractionPatch{
				IdentityTypes: identityTypes,
			}
			updateSettingsOptions.SetExternalAccountIdentityInteraction(externalAccountIdentityInteraction)

			accountSettingsAccessManagement, response, err := iamPolicyManagementService.UpdateSettings(updateSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettingsAccessManagement, "", "  ")
			fmt.Println(string(b))

			// end-update_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettingsAccessManagement).ToNot(BeNil())
		})

		It(`CreateActionControlTemplate request example`, func() {
			fmt.Println("\nCreateActionControlTemplate() result:")
			// begin-create_action_control_template

			templateActionControl := &iampolicymanagementv1.TemplateActionControl{
				ServiceName: core.StringPtr("am-test-service"),
				Description: core.StringPtr("am-test-service service actionControl"),
				Actions:     []string{"am-test-service.test.create"},
			}

			createActionControlTemplateOptions := &iampolicymanagementv1.CreateActionControlTemplateOptions{
				Name:           &exampleActionControlTemplateName,
				AccountID:      &exampleAccountID,
				ActionControl:  templateActionControl,
				Description:    core.StringPtr("Test ActionControl Template from GO SDK"),
				Committed:      core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("default"),
			}

			actionControlTemplate, response, err := iamPolicyManagementService.CreateActionControlTemplate(createActionControlTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(actionControlTemplate, "", "  ")
			exampleActionControlTemplateID = *actionControlTemplate.ID
			exampleActionControlTemplateBaseVersion = *actionControlTemplate.Version
			exampleActionControlTemplateETag = response.GetHeaders().Get("ETag")
			fmt.Println(string(b))

			// end-create_action_control_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(actionControlTemplate).ToNot(BeNil())
		})

		It(`GetActionControlTemplate request example`, func() {
			fmt.Println("\nGetActionControlTemplate() result:")
			// begin-get_action_control_template

			getActionControlTemplateOptions := iamPolicyManagementService.NewGetActionControlTemplateOptions(
				exampleActionControlTemplateID,
			)

			actionControlTemplate, response, err := iamPolicyManagementService.GetActionControlTemplate(getActionControlTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(actionControlTemplate, "", "  ")
			fmt.Println(string(b))

			// end-get_action_control_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionControlTemplate.AccountID).ToNot(BeNil())
			Expect(actionControlTemplate.Version).ToNot(BeNil())
			Expect(actionControlTemplate.Name).ToNot(BeNil())
			Expect(actionControlTemplate.ActionControl).ToNot(BeNil())
		})

		It(`CreateActionControlTemplateVersion request example`, func() {
			fmt.Println("\nCreateActionControlTemplateVersion() result:")
			// begin-create_action_control_template_version
			templateActionControl := &iampolicymanagementv1.TemplateActionControl{
				ServiceName: core.StringPtr("am-test-service"),
				Description: core.StringPtr("am-test-service service actionControl"),
				Actions:     []string{"am-test-service.test.delete"},
			}

			updateActionControlTemplateVersionOptions := &iampolicymanagementv1.CreateActionControlTemplateVersionOptions{
				ActionControl:           templateActionControl,
				Description:             core.StringPtr("Test of ActionControl Template version from GO SDK"),
				ActionControlTemplateID: &exampleActionControlTemplateID,
			}

			actionControlTemplate, response, err := iamPolicyManagementService.CreateActionControlTemplateVersion(updateActionControlTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(actionControlTemplate, "", "  ")
			exampleActionControlTemplateVersion = *actionControlTemplate.Version
			exampleActionControlTemplateETag = response.GetHeaders().Get("ETag")
			fmt.Println(string(b))

			// end-create_action_control_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(actionControlTemplate).ToNot(BeNil())
		})

		It(`ListActionControlTemplateVersions request example`, func() {
			fmt.Println("\nListActionControlTemplateVersions() result:")
			// begin-list_action_control_templates

			listActionControlTemplateVersionsOptions := iamPolicyManagementService.NewListActionControlTemplateVersionsOptions(
				exampleActionControlTemplateID,
			)

			actionControlTemplateVersionsCollection, response, err := iamPolicyManagementService.ListActionControlTemplateVersions(listActionControlTemplateVersionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(actionControlTemplateVersionsCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_action_control_templates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionControlTemplateVersionsCollection).ToNot(BeNil())
		})

		It(`ReplaceActionControlTemplate request example`, func() {
			fmt.Println("\nReplaceActionControlTemplate() result:")
			// begin-replace_action_control_template
			templateActionControl := &iampolicymanagementv1.TemplateActionControl{
				ServiceName: core.StringPtr("am-test-service"),
				Description: core.StringPtr("am-test-service service actionControl"),
				Actions:     []string{"am-test-service.test.delete", "am-test-service.test.create"},
			}

			replaceActionControlTemplateVersionOptions := &iampolicymanagementv1.ReplaceActionControlTemplateOptions{
				ActionControl:           templateActionControl,
				Description:             core.StringPtr("Test update of ActionControl Template from GO SDK"),
				Committed:               core.BoolPtr(true),
				IfMatch:                 core.StringPtr(exampleActionControlTemplateETag),
				Version:                 core.StringPtr(exampleActionControlTemplateVersion),
				ActionControlTemplateID: &exampleActionControlTemplateID,
			}

			actionControlTemplate, response, err := iamPolicyManagementService.ReplaceActionControlTemplate(replaceActionControlTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(actionControlTemplate, "", "  ")
			exampleActionControlTemplateVersion = *actionControlTemplate.Version
			exampleActionControlTemplateETag = response.GetHeaders().Get("ETag")
			fmt.Println(string(b))

			// end-replace_action_control_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionControlTemplate).ToNot(BeNil())
		})

		It(`GetActionControlTemplateVersion request example`, func() {
			fmt.Println("\nGetActionControlTemplateVersion() result:")
			// begin-get_action_control_template_version

			getActionControlTemplateVersionOptions := iamPolicyManagementService.NewGetActionControlTemplateVersionOptions(
				exampleActionControlTemplateID,
				exampleActionControlTemplateVersion,
			)

			actionControlTemplate, response, err := iamPolicyManagementService.GetActionControlTemplateVersion(getActionControlTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(actionControlTemplate, "", "  ")
			fmt.Println(string(b))

			// end-get_action_control_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionControlTemplate).ToNot(BeNil())
		})

		It(`CommitActionControlTemplate request example`, func() {
			fmt.Println("\nCommitActionControlTemplate() result:")
			// begin-commit_action_control_template

			commitActionControlTemplateOptions := iamPolicyManagementService.NewCommitActionControlTemplateOptions(
				exampleActionControlTemplateID,
				exampleActionControlTemplateBaseVersion,
			)

			response, err := iamPolicyManagementService.CommitActionControlTemplate(commitActionControlTemplateOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from CommitActionControlTemplate(): %d\n", response.StatusCode)
			}

			// end-commit_action_control_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

		It(`ListActionControlTemplates request example`, func() {
			fmt.Println("\nListActionControlTemplates() result:")
			// begin-list_action_Control_templates

			listActionControlTemplatesOptions := iamPolicyManagementService.NewListActionControlTemplatesOptions(
				exampleAccountID,
			)

			actionControlTemplateCollection, response, err := iamPolicyManagementService.ListActionControlTemplates(listActionControlTemplatesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(actionControlTemplateCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_action_Control_templates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionControlTemplateCollection).ToNot(BeNil())
		})

		It(`CreateActionControlAssignments request example`, func() {
			fmt.Println("\nCreateActionControlTemplateAssignment() result:")
			// begin-create_action_control_template_assignment
			template := iampolicymanagementv1.ActionControlAssignmentTemplate{
				ID:      &exampleActionControlTemplateID,
				Version: &exampleActionControlTemplateVersion,
			}
			templates := []iampolicymanagementv1.ActionControlAssignmentTemplate{
				template,
			}

			target := &iampolicymanagementv1.AssignmentTargetDetails{
				Type: core.StringPtr("Account"),
				ID:   &exampleTargetAccountID,
			}

			createPolicyTemplateVersionOptions := &iampolicymanagementv1.CreateActionControlTemplateAssignmentOptions{
				Target:    target,
				Templates: templates,
			}

			actionControlAssignment, response, err := iamPolicyManagementService.CreateActionControlTemplateAssignment(createPolicyTemplateVersionOptions)

			b, _ := json.MarshalIndent(actionControlAssignment, "", "  ")
			fmt.Println(string(b))

			var assignmentDetails = actionControlAssignment.Assignments[0]
			// end-create_action_control_template_assignment
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			exampleActionControlAssignmentETag = response.GetHeaders().Get(exampleETagHeader)
			exampleActionControlAssignmentId = *assignmentDetails.ID
		})

		It(`UpdateActionControlAssignment request example))`, func() {
			// begin-update_action_control_assignment
			updatePolicyAssignmentOptions := iamPolicyManagementService.NewUpdateActionControlAssignmentOptions(
				exampleActionControlAssignmentId,
				exampleActionControlAssignmentETag,
				exampleActionControlTemplateBaseVersion,
			)

			actionControlAssignment, response, err := iamPolicyManagementService.UpdateActionControlAssignment(updatePolicyAssignmentOptions)
			b, _ := json.MarshalIndent(actionControlAssignment, "", "  ")
			fmt.Println(string(b))
			// end-update_action_control_assignment
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			exampleActionControlAssignmentId = *actionControlAssignment.ID
		})

		It(`ListActionControlAssignments request example`, func() {
			fmt.Println("\nListActionControlAssignments() result:")
			// begin-list_action_control_assignments

			listActionControlAssignmentsOptions := iamPolicyManagementService.NewListActionControlAssignmentsOptions(
				exampleAccountID,
			)

			actionControlTemplateAssignmentCollection, response, err := iamPolicyManagementService.ListActionControlAssignments(listActionControlAssignmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(actionControlTemplateAssignmentCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_action_control_assignments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			var assignmentDetails = actionControlTemplateAssignmentCollection.Assignments[0]
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

		It(`GetActionControlAssignment request example`, func() {
			fmt.Println("\nGetActionControlAssignment() result:")
			// begin-get_action_control_assignment

			getActionControlAssignmentOptions := iamPolicyManagementService.NewGetActionControlAssignmentOptions(
				exampleActionControlAssignmentId,
			)

			assignmentDetails, response, err := iamPolicyManagementService.GetActionControlAssignment(getActionControlAssignmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(assignmentDetails, "", "  ")
			fmt.Println(string(b))

			// end-get_action_control_assignment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(assignmentDetails).ToNot(BeNil())
			Expect(assignmentDetails.Template.ID).ToNot(BeNil())
			Expect(assignmentDetails.Target.Type).ToNot(BeNil())
			Expect(assignmentDetails.Template.Version).ToNot(BeNil())
			Expect(assignmentDetails.Target.ID).ToNot(BeNil())
			Expect(assignmentDetails.Status).ToNot(BeNil())
			Expect(assignmentDetails.AccountID).ToNot(BeNil())
			Expect(assignmentDetails.CreatedAt).ToNot(BeNil())
			Expect(assignmentDetails.CreatedByID).ToNot(BeNil())
			Expect(assignmentDetails.LastModifiedAt).ToNot(BeNil())
			Expect(assignmentDetails.LastModifiedByID).ToNot(BeNil())
			Expect(assignmentDetails.Href).ToNot(BeNil())
		})

		It(`DeleteActionControlAssignment request example)`, func() {
			// begin-delete_action_control_assignment
			deleteActionControlAssignmentOptions := iamPolicyManagementService.NewDeleteActionControlAssignmentOptions(
				exampleActionControlAssignmentId,
			)

			response, err := iamPolicyManagementService.DeleteActionControlAssignment(deleteActionControlAssignmentOptions)
			// end-delete_action_control_assignment
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

		It(`DeleteActionControlTemplateVersion request example`, func() {
			// begin-delete_action_control_template_version

			deleteActionControlTemplateVersionOptions := iamPolicyManagementService.NewDeleteActionControlTemplateVersionOptions(
				exampleActionControlTemplateID,
				exampleActionControlTemplateVersion,
			)

			response, err := iamPolicyManagementService.DeleteActionControlTemplateVersion(deleteActionControlTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteActionControlTemplateVersion(): %d\n", response.StatusCode)
			}

			// end-delete_action_control_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

		It(`DeleteActionControlTemplate request example`, func() {
			// begin-delete_action_control_template

			deleteActionControlTemplateOptions := iamPolicyManagementService.NewDeleteActionControlTemplateOptions(
				exampleActionControlTemplateID,
			)

			response, err := iamPolicyManagementService.DeleteActionControlTemplate(deleteActionControlTemplateOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeletePolicyTemplate(): %d\n", response.StatusCode)
			}

			// end-delete_action_control_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
