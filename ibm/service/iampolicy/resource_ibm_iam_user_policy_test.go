// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package iampolicy_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMIAMUserPolicy_Basic(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "description", "IAM User Policy Creation for test scenario"),
				),
			},
			{
				Config: testAccCheckIBMIAMUserPolicyUpdateRole(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "description", "IAM User Policy Update for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Service(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyService(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMIAMUserPolicyUpdateServiceAndRegion(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.region", "us-south"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_ServiceType(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyServiceType(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.service_type", "service"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_ResourceInstance(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyResourceInstance(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Resource_Group(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyResourceGroup(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.service", "containers-kubernetes"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Resource_Type(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyResourceType(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_import(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_user_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"resources", "resource_attributes", "transaction_id"},
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Resource_Attributes(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyResourceAttributes(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resource_attributes.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMIAMUserPolicyResourceAttributesUpdate(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resource_attributes.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Resource_Attributes_Without_Wildcard(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyResourceAttributesWithoutWildcard(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resource_attributes.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_account_management(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_user_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyAccountManagement(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "account_management", "true"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_Invalid_User(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMIAMUserPolicyInvalidUser(),
				ExpectError: regexp.MustCompile(`User test@in.ibm.com is not found`),
			},
		},
	})
}

func TestAccIBMIAMUserPolicyWithCustomRole(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	crName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyWithCustomRole(crName, displayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicyWithSpecificServiceRole(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyWithServiceSpecificRole(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Resource_Tags(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyResourceTags(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resource_tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "description", "IAM User Policy Creation for test scenario"),
				),
			},
			{
				Config: testAccCheckIBMIAMUserPolicyResourceTagsUpdate(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resource_tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "description", "IAM User Policy Update for test scenario"),
				),
			},
		},
	})

}

func TestAccIBMIAMUserPolicy_With_Transaction_Id(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyTransactionId(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "transaction_id", "terrformUserPolicy"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Time_Based_Conditions_Weekly_Custom(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyWeeklyCustomHours(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "pattern", "time-based-conditions:weekly:custom-hours"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "rule_conditions.#", "3"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "rule_conditions.2.key", "{{environment.attributes.day_of_week}}"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "rule_conditions.2.value.#", "5"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "description", "IAM User Policy Custom Hours Creation for test scenario"),
				),
			},
			{
				Config: testAccCheckIBMIAMUserPolicyUpdateConditions(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "pattern", "time-based-conditions:weekly:custom-hours"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "rule_conditions.2.value.#", "4"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "description", "IAM User Policy Custom Hours Update for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Time_Based_Conditions_Weekly_All_Day(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyWeeklyAllDay(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "pattern", "time-based-conditions:weekly:all-day"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "rule_conditions.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "rule_conditions.0.key", "{{environment.attributes.day_of_week}}"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "rule_conditions.0.value.#", "5"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "description", "IAM User Policy All Day Weekly Time-Based Conditions Creation for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Time_Based_Conditions_Once(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyTimeBasedOnce(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "pattern", "time-based-conditions:once"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "rule_conditions.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "rule_conditions.0.key", "{{environment.attributes.current_date_time}}"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "description", "IAM User Policy Once Time-Based Conditions Creation for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Time_Based_Conditions_With_Resource_Attributes(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyTimeBasedWithResourceAttributes(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "pattern", "time-based-conditions:once"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "rule_conditions.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "rule_conditions.0.key", "{{environment.attributes.current_date_time}}"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "description", "IAM User Policy Once Time-Based Conditions Creation for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Update_To_Time_Based_Conditions(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyResourceAttributes(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resource_attributes.#", "2"),
				),
			},
			{
				Config:      testAccCheckIBMIAMUserPolicyUpdateConditions(),
				ExpectError: regexp.MustCompile("Error: Cannot use rule_conditions, rule_operator, or pattern when updating v1/policy. Delete existing v1/policy and create using rule_conditions and pattern."),
			},
		},
	})
}

func TestAccIBMIAMUSerPolicy_With_ServiceGroupID(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyWithServiceGroupId(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resource_attributes.0.value", "IAM"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Attribute_Based_Condition(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyAttributeBasedCondition(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "pattern", "attribute-based-condition:resource:literal-and-wildcard"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "description", "IAM User Policy Attribute Based Condition Creation for test scenario"),
				),
			},
			{
				Config: testAccCheckIBMIAMUserPolicyUpdateAttributeBasedCondition(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "pattern", "attribute-based-condition:resource:literal-and-wildcard"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "description", "IAM User Policy Attribute Based Condition Update for test scenario"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMUserPolicyDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_user_policy" {
			continue
		}
		policyID := rs.Primary.ID
		parts, err := flex.IdParts(policyID)
		if err != nil {
			return err
		}

		userPolicyID := parts[1]

		getPolicyOptions := rsContClient.NewGetV2PolicyOptions(
			userPolicyID,
		)

		// Try to find the key
		destroyedPolicy, response, err := rsContClient.GetV2Policy(getPolicyOptions)

		if err == nil && *destroyedPolicy.State != "deleted" {
			return fmt.Errorf("User policy still exists: %s\n", rs.Primary.ID)
		} else if response.StatusCode != 404 && destroyedPolicy.State != nil && *destroyedPolicy.State != "deleted" {
			return fmt.Errorf("[ERROR] Error waiting for user policy (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMUserPolicyExists(n string, obj iampolicymanagementv1.V2PolicyTemplateMetaData) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
		if err != nil {
			return err
		}

		policyID := rs.Primary.ID
		parts, err := flex.IdParts(policyID)
		if err != nil {
			return err
		}
		userPolicyID := parts[1]

		getPolicyOptions := rsContClient.NewGetV2PolicyOptions(
			userPolicyID,
		)

		policy, _, err := rsContClient.GetV2Policy(getPolicyOptions)
		if err != nil {
			return err
		}
		obj = *policy
		return nil
	}
}

func testAccCheckIBMIAMUserPolicyBasic() string {
	return fmt.Sprintf(`

		  
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer"]
			tags   = ["tag1"]
			description = "IAM User Policy Creation for test scenario"
	  	}

	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyUpdateRole() string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer", "Manager"]
			tags   = ["tag1", "tag2"]
			description = "IAM User Policy Update for test scenario"
	  	}
	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyService() string {
	return fmt.Sprintf(`

		
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer"]
	  
			resources {
		 		 service = "cloudantnosqldb"
			}
	  	}

	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyWithServiceSpecificRole() string {
	return fmt.Sprintf(`

		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = [ "Monitor", "Reader", "Viewer"]
			resources {
		 		 service = "cloudantnosqldb"
			}
	  	}

	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyServiceType() string {
	return fmt.Sprintf(`

		
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer"]
	  
			resources {
				service_type = "service"
				region = "us-south"
			}
	  	}

	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyUpdateServiceAndRegion() string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_user_policy" "policy" {
			ibm_id 		 = "%s"
			roles        = ["Viewer", "Manager"]
		  
			resources {
			  service = "cloudantnosqldb"
			  region  = "us-south"
			}
		  }
	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyResourceInstance(name string) string {
	return fmt.Sprintf(`

		resource "ibm_resource_instance" "instance" {
			name     = "%s"
			service  = "kms"
			plan     = "tiered-pricing"
			location = "us-south"
	  	}
	  
	  	resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Manager", "Viewer", "Administrator"]
	  
			resources {
		  		service              = "kms"
		  		resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
			}
	  	}
		  

	`, name, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyResourceGroup() string {
	return fmt.Sprintf(`

		  
		data "ibm_resource_group" "group" {
			is_default=true
	  	}
	  
	  	resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer"]
	  
			resources {
		 	 service           = "containers-kubernetes"
		  	 resource_group_id = data.ibm_resource_group.group.id
			}
	  	}
		  

	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyResourceType() string {
	return fmt.Sprintf(`

		  
		data "ibm_resource_group" "group" {
			is_default=true
		  }
		  
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Administrator"]
		  
			resources {
			  resource_type = "resource-group"
			  resource      = data.ibm_resource_group.group.id
			}
		  }
	`, acc.IAMUser)
}

// TODO: do we need this test? It follows pattern of other policies, but has conflict with existing policy
func testAccCheckIBMIAMUserPolicyImport(name string) string {
	return fmt.Sprintf(`

	
		  resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles        = ["Viewer"]
		  }

	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyInvalidUser() string {
	return `

		  
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "test@in.ibm.com"
			roles  = ["Viewer"]
	  	}

	`
}

func testAccCheckIBMIAMUserPolicyAccountManagement(name string) string {
	return fmt.Sprintf(`
	
		  resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer"]
			account_management = true
		  }

	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyWithCustomRole(crName, displayName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_custom_role" "customrole" {
			name         = "%s"
			display_name = "%s"
			description  = "role for test scenario1"
			service = "kms"
			actions      = ["kms.secrets.rotate"]
		}
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = [ibm_iam_custom_role.customrole.display_name,"Viewer"]
			resources {
				service = "kms"
			  }
	  	}

	`, crName, displayName, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyResourceAttributes() string {
	return fmt.Sprintf(`
  
	  resource "ibm_iam_user_policy" "policy" {
		ibm_id = "%s"
		roles  = ["Viewer"]
		resource_attributes {
			name     = "resource"
			value    = "test*"
			operator = "stringMatch"
		}
		resource_attributes {
			name     = "serviceName"
			value    = "messagehub"
		}
	  }
	  
`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyResourceAttributesWithoutWildcard() string {
	return fmt.Sprintf(`
  
	  resource "ibm_iam_user_policy" "policy" {
		ibm_id = "%s"
		roles  = ["Viewer"]
		resource_attributes {
			name     = "resource"
			value    = "test"
			operator = "stringMatch"
		}
		resource_attributes {
			name     = "serviceName"
			value    = "messagehub"
		}
	  }
	  
`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyResourceAttributesUpdate() string {
	return fmt.Sprintf(`
	resource "ibm_iam_user_policy" "policy" {
		ibm_id = "%s"
		roles  = ["Viewer"]
		resource_attributes {
			name     = "resource"
			value    = "test*"
			operator = "stringMatch"
		}
		resource_attributes {
			name     = "serviceName"
			value    = "messagehub"
		}
	  }
	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyResourceTags() string {
	return fmt.Sprintf(`
  
	resource "ibm_iam_user_policy" "policy" {
		ibm_id = "%s"
		roles  = ["Viewer"]
		description = "IAM User Policy Creation for test scenario"
		resources {
			service_type = "service"
		}
		
		resource_tags {
			name = "test"
			value = "terraform"
		}
	}
	  
`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyResourceTagsUpdate() string {
	return fmt.Sprintf(`
	
	resource "ibm_iam_user_policy" "policy" {
		ibm_id = "%s"
		roles  = ["Viewer", "Manager"]
		description = "IAM User Policy Update for test scenario"
		resources {
			service_type = "service"
		}
		
		resource_tags {
			name = "test"
			value = "terraform"
		}
		resource_tags {
			name = "two"
			value = "terrformupdate"
		}
	}
	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyTransactionId() string {
	return fmt.Sprintf(`

		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer"]
			transaction_id = "terrformUserPolicy"
			resources {
		 		 service = "cloudantnosqldb"
			}
	  	}

	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyWeeklyCustomHours() string {
	return fmt.Sprintf(`

		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer"]
			resources {
				 service = "kms"
			}
			rule_conditions {
				key = "{{environment.attributes.day_of_week}}"
				operator = "dayOfWeekAnyOf"
				value = ["1+00:00","2+00:00","3+00:00","4+00:00", "5+00:00"]
			}
			rule_conditions {
				key = "{{environment.attributes.current_time}}"
				operator = "timeGreaterThanOrEquals"
				value = ["09:00:00+00:00"]
			}
			rule_conditions {
				key = "{{environment.attributes.current_time}}"
				operator = "timeLessThanOrEquals"
				value = ["17:00:00+00:00"]
			}
			rule_operator = "and"
		  pattern = "time-based-conditions:weekly:custom-hours"
			description = "IAM User Policy Custom Hours Creation for test scenario"
		}
	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyUpdateConditions() string {
	return fmt.Sprintf(`

		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer", "Manager"]
			resources {
				 service = "kms"
			}
			rule_conditions {
				key = "{{environment.attributes.day_of_week}}"
				operator = "dayOfWeekAnyOf"
				value = ["1+00:00","2+00:00","3+00:00","4+00:00"]
			}
			rule_conditions {
				key = "{{environment.attributes.current_time}}"
				operator = "timeGreaterThanOrEquals"
				value = ["09:00:00+00:00"]
			}
			rule_conditions {
				key = "{{environment.attributes.current_time}}"
				operator = "timeLessThanOrEquals"
				value = ["17:00:00+00:00"]
			}
			rule_operator = "and"
		  pattern = "time-based-conditions:weekly:custom-hours"
			description = "IAM User Policy Custom Hours Update for test scenario"
		}
	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyWeeklyAllDay() string {
	return fmt.Sprintf(`

		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer"]
			resources {
				 service = "kms"
			}
			rule_conditions {
				key = "{{environment.attributes.day_of_week}}"
				operator = "dayOfWeekAnyOf"
				value = ["1+00:00","2+00:00","3+00:00","4+00:00", "5+00:00"]
			}

		  pattern = "time-based-conditions:weekly:all-day"
			description = "IAM User Policy All Day Weekly Time-Based Conditions Creation for test scenario"
		}
	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyTimeBasedOnce() string {
	return fmt.Sprintf(`

		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer"]
			resources {
				 service = "kms"
			}
			rule_conditions {
				key = "{{environment.attributes.current_date_time}}"
				operator = "dateTimeGreaterThanOrEquals"
				value = ["2022-10-01T12:00:00+00:00"]
			}
			rule_conditions {
				key = "{{environment.attributes.current_date_time}}"
				operator = "dateTimeLessThanOrEquals"
				value = ["2022-10-31T12:00:00+00:00"]
			}
			rule_operator = "and"
		  pattern = "time-based-conditions:once"
			description = "IAM User Policy Once Time-Based Conditions Creation for test scenario"
		}
	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyTimeBasedWithResourceAttributes() string {
	return fmt.Sprintf(`

		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer"]
			resource_attributes {
				name     = "resource"
				value    = "test*"
				operator = "stringMatch"
			}
			resource_attributes {
				name     = "serviceName"
				value    = "messagehub"
			}
			rule_conditions {
				key = "{{environment.attributes.current_date_time}}"
				operator = "dateTimeGreaterThanOrEquals"
				value = ["2022-10-01T12:00:00+00:00"]
			}
			rule_conditions {
				key = "{{environment.attributes.current_date_time}}"
				operator = "dateTimeLessThanOrEquals"
				value = ["2022-10-31T12:00:00+00:00"]
			}
			rule_operator = "and"
		  pattern = "time-based-conditions:once"
			description = "IAM User Policy Once Time-Based Conditions Creation for test scenario"
		}
	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyWithServiceGroupId(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"

			roles           = ["Service ID creator"]
    		resource_attributes {
         		name     = "service_group_id"
         		operator = "stringEquals"
         		value    = "IAM"
			}
		}
	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyAttributeBasedCondition() string {
	return fmt.Sprintf(`

		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Writer"]
			resource_attributes {
				value = "cloud-object-storage"
				operator = "stringEquals"
				name = "serviceName"
			}
			resource_attributes {
				value = "cos-instance"
				operator = "stringEquals"
				name = "serviceInstance"
			}
			resource_attributes {
				value = "bucket"
				operator = "stringEquals"
				name = "resourceType"
			}
			resource_attributes {
				value = "fgac-tf-test"
				operator = "stringEquals"
				name = "resource"
			}
			rule_conditions {
				operator = "and"
				conditions {
					key = "{{resource.attributes.prefix}}"
					operator = "stringMatch"
					value = ["folder1/subfolder1/*"]
				}
				conditions {
					key = "{{resource.attributes.delimiter}}"
					operator = "stringEqualsAnyOf"
					value = ["/",""]
				}
			}
			rule_conditions {
				key = "{{resource.attributes.path}}"
				operator = "stringMatch"
				value = ["folder1/subfolder1/*"]
			}
			rule_conditions {
				operator = "and"
				conditions {
					key = "{{resource.attributes.delimiter}}"
					operator = "stringExists"
					value = ["false"]
				}
				conditions {
					key = "{{resource.attributes.prefix}}"
					operator = "stringExists"
					value = ["false"]
				}
			}
			rule_operator = "or"
		  pattern = "attribute-based-condition:resource:literal-and-wildcard"
			description = "IAM User Policy Attribute Based Condition Creation for test scenario"
		}
	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyUpdateAttributeBasedCondition() string {
	return fmt.Sprintf(`
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Reader", "Writer"]
			resource_attributes {
				value = "cloud-object-storage"
				operator = "stringEquals"
				name = "serviceName"
			}
			resource_attributes {
				value = "cos-instance"
				operator = "stringEquals"
				name = "serviceInstance"
			}
			resource_attributes {
				value = "bucket"
				operator = "stringEquals"
				name = "resourceType"
			}
			resource_attributes {
				value = "fgac-tf-test"
				operator = "stringEquals"
				name = "resource"
			}
			rule_conditions {
				operator = "and"
				conditions {
					key = "{{resource.attributes.prefix}}"
					operator = "stringMatch"
					value = ["folder1/subfolder1/*"]
				}
				conditions {
					key = "{{resource.attributes.delimiter}}"
					operator = "stringEqualsAnyOf"
					value = ["/",""]
				}
			}
			rule_conditions {
				key = "{{resource.attributes.path}}"
				operator = "stringMatch"
				value = ["folder1/subfolder1/*"]
			}
			rule_conditions {
				operator = "and"
				conditions {
					key = "{{resource.attributes.delimiter}}"
					operator = "stringExists"
					value = ["false"]
				}
				conditions {
					key = "{{resource.attributes.prefix}}"
					operator = "stringExists"
					value = ["false"]
				}
			}
			rule_operator = "or"
		  pattern = "attribute-based-condition:resource:literal-and-wildcard"
			description = "IAM User Policy Attribute Based Condition Update for test scenario"
		}
	`, acc.IAMUser)
}
