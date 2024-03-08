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

func TestAccIBMIAMServicePolicy_Basic(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "description", "IAM Service Policy Creation for test scenario"),
				),
			},
			{
				Config: testAccCheckIBMIAMServicePolicyUpdateRole(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "description", "IAM Service Policy Update for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Service(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyService(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMIAMServicePolicyUpdateServiceAndRegion(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.region", "us-south"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_ServiceType(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyServiceType(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service_type", "service"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_ResourceInstance(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyResourceInstance(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Resource_Group(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyResourceGroup(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service", "containers-kubernetes"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Resource_Type(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyResourceType(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_import(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_service_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
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

func TestAccIBMIAMServicePolicy_account_management(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_service_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyAccountManagement(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "account_management", "true"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicyWithCustomRole(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	crName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyWithCustomRole(name, crName, displayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Resource_Attributes(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyResourceAttributes(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resource_attributes.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMIAMServicePolicyResourceAttributesUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resource_attributes.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Resource_Tags(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyResourceTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resource_tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "description", "IAM Service Policy Creation for test scenario"),
				),
			},
			{
				Config: testAccCheckIBMIAMServicePolicyUpdateResourceTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resource_tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "description", "IAM Service Policy Update for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Transaction_Id(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyResourceTransactionId(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resource_attributes.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "transaction_id", "terrformServicePolicy")),
			},
			{
				Config: testAccCheckIBMIAMServicePolicyResourceTransactionIdUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resource_attributes.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "transaction_id", "terrformServicePolicyUpdate"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Time_Based_Conditions_Weekly_Custom(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyWeeklyCustomHours(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "pattern", "time-based-conditions:weekly:custom-hours"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "rule_conditions.#", "3"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "rule_conditions.2.key", "{{environment.attributes.day_of_week}}"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "rule_conditions.2.value.#", "5"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "description", "IAM Service Profile Policy Custom Hours Creation for test scenario"),
				),
			},
			{
				Config: testAccCheckIBMIAMServicePolicyUpdateConditions(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "pattern", "time-based-conditions:weekly:custom-hours"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "rule_conditions.2.value.#", "4"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "description", "IAM Service Profile Policy Custom Hours Update for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Time_Based_Conditions_Weekly_All_Day(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyWeeklyAllDay(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "pattern", "time-based-conditions:weekly:all-day"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "rule_conditions.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "rule_conditions.0.key", "{{environment.attributes.day_of_week}}"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "rule_conditions.0.value.#", "5"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "description", "IAM Service Profile Policy All Day Weekly Time-Based Conditions Creation for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Time_Based_Conditions_Once(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyTimeBasedOnce(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "pattern", "time-based-conditions:once"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "rule_conditions.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "rule_conditions.0.key", "{{environment.attributes.current_date_time}}"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "description", "IAM Service Profile Policy Once Time-Based Conditions Creation for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Update_To_Time_Based_Conditions(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyResourceAttributes(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resource_attributes.#", "2"),
				),
			},
			{
				Config:      testAccCheckIBMIAMServicePolicyUpdateConditions(name),
				ExpectError: regexp.MustCompile("Error: Cannot use rule_conditions, rule_operator, or pattern when updating v1/policy. Delete existing v1/policy and create using rule_conditions and pattern."),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_ServiceGroupID(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyWithServiceGroupId(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resource_attributes.0.value", "IAM"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMIAMServiceUpdatePolicyWithServiceGroupId(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Attribute_Based_Condition(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyAttributeBasedCondition(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "pattern", "attribute-based-condition:resource:literal-and-wildcard"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "description", "IAM Service Policy Attribute Based Condition Creation for test scenario"),
				),
			},
			{
				Config: testAccCheckIBMIAMServicePolicyUpdateAttributeBasedCondition(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "pattern", "attribute-based-condition:resource:literal-and-wildcard"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "description", "IAM Service Policy Attribute Based Condition Update for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Resource_Attributes_Without_Wildcard(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyResourceAttributesWithoutWildcard(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resource_attributes.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMServicePolicyDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_service_policy" {
			continue
		}
		policyID := rs.Primary.ID
		parts, err := flex.IdParts(policyID)
		if err != nil {
			return err
		}
		servicePolicyID := parts[1]

		getPolicyOptions := rsContClient.NewGetV2PolicyOptions(
			servicePolicyID,
		)

		// Try to find the key
		destroyedPolicy, response, err := rsContClient.GetV2Policy(getPolicyOptions)

		if err == nil && *destroyedPolicy.State != "deleted" {
			return fmt.Errorf("Service policy still exists: %s\n", rs.Primary.ID)
		} else if response.StatusCode != 404 && destroyedPolicy.State != nil && *destroyedPolicy.State != "deleted" {
			return fmt.Errorf("[ERROR] Error waiting for Service policy (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMServicePolicyExists(n string, obj iampolicymanagementv1.V2PolicyTemplateMetaData) resource.TestCheckFunc {

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
		servicePolicyID := parts[1]

		getPolicyOptions := rsContClient.NewGetV2PolicyOptions(
			servicePolicyID,
		)

		// Try to find the key
		policy, _, err := rsContClient.GetV2Policy(getPolicyOptions)
		if err != nil {
			return err
		}
		obj = *policy
		return nil
	}
}

func testAccCheckIBMIAMServicePolicyBasic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
			tags           = ["tag1"]
			description    = "IAM Service Policy Creation for test scenario"
	  	}

	`, name)
}

func testAccCheckIBMIAMServicePolicyUpdateRole(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer", "Manager"]
			tags           = ["tag1", "tag2"]
			description    = "IAM Service Policy Update for test scenario"
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyService(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
	  
			resources {
		 	 service = "cloudantnosqldb"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyServiceType(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
	  
			resources {
				service_type = "service"
				region = "us-south"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyUpdateServiceAndRegion(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer", "Manager"]
	  
			resources {
		  		service = "cloudantnosqldb"
		  		region  = "us-south"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyResourceInstance(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_resource_instance" "instance" {
			name     = "%s"
			service  = "kms"
			plan     = "tiered-pricing"
			location = "us-south"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Manager", "Viewer", "Administrator"]
	  
			resources {
		 		 service              = "kms"
		  		resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
			}
	  	}
		  

	`, name, name)
}

func testAccCheckIBMIAMServicePolicyResourceGroup(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	data "ibm_resource_group" "group" {
			is_default=true
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
	  
			resources {
		 		service           = "containers-kubernetes"
		  		resource_group_id = data.ibm_resource_group.group.id
			}
	  	}
		  

	`, name)
}

func testAccCheckIBMIAMServicePolicyResourceType(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	data "ibm_resource_group" "group" {
			is_default=true
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Administrator"]
	  
			resources {
		  		resource_type = "resource-group"
		  		resource      = data.ibm_resource_group.group.id
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyImport(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
	  	}

	`, name)
}

func testAccCheckIBMIAMServicePolicyAccountManagement(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id     = ibm_iam_service_id.serviceID.id
			roles              = ["Viewer"]
			account_management = true
	  	}

	`, name)
}

func testAccCheckIBMIAMServicePolicyWithCustomRole(name, crName, displayName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
		}
		resource "ibm_iam_custom_role" "customrole" {
			name         = "%s"
			display_name = "%s"
			description  = "role for test scenario1"
			service = "kms"
			actions      = ["kms.secrets.rotate"]
		}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = [ibm_iam_custom_role.customrole.display_name,"Viewer"]
			tags           = ["tag1"]
			resources {
				service           = "kms"
		   }
	  	}

	`, name, crName, displayName)
}

func testAccCheckIBMIAMServicePolicyResourceAttributes(name string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_service_id" "serviceID" {
		name = "%s"
	  }
  
	  resource "ibm_iam_service_policy" "policy" {
		iam_service_id     = ibm_iam_service_id.serviceID.id
		roles              = ["Viewer"]
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
	`, name)
}
func testAccCheckIBMIAMServicePolicyResourceAttributesUpdate(name string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_service_id" "serviceID" {
		name = "%s"
	  }
  
	  resource "ibm_iam_service_policy" "policy" {
		iam_service_id     = ibm_iam_service_id.serviceID.id
		roles              = ["Viewer"]
		resource_attributes {
			name     = "resource"
			value    = "test*"
		}
		resource_attributes {
			name     = "serviceName"
			value    = "messagehub"
		}
	  }
	`, name)
}

func testAccCheckIBMIAMServicePolicyResourceTags(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
			
			resource_tags {
				name  = "one"
				value = "Terraform"
			}
			description    = "IAM Service Policy Creation for test scenario"
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyUpdateResourceTags(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
			
			resource_tags {
				name  = "one"
				value = "Terraform"
			}
			resource_tags {
				name  = "two"
				value = "TerraformUpdate"
			}
			description    = "IAM Service Policy Update for test scenario"
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyResourceTransactionId(name string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_service_id" "serviceID" {
		name = "%s"
	  }
  
	  resource "ibm_iam_service_policy" "policy" {
		iam_service_id     = ibm_iam_service_id.serviceID.id
		roles              = ["Viewer"]
		transaction_id = "terrformServicePolicy"
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
	`, name)
}

func testAccCheckIBMIAMServicePolicyResourceTransactionIdUpdate(name string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_service_id" "serviceID" {
		name = "%s"
	  }
  
	  resource "ibm_iam_service_policy" "policy" {
		iam_service_id     = ibm_iam_service_id.serviceID.id
		roles              = ["Viewer"]
		transaction_id = "terrformServicePolicyUpdate"
		resource_attributes {
			name     = "resource"
			value    = "test*"
		}
		resource_attributes {
			name     = "serviceName"
			value    = "messagehub"
		}
	  }
	`, name)
}

func testAccCheckIBMIAMServicePolicyWeeklyCustomHours(name string) string {
	return fmt.Sprintf(`
	  resource "ibm_iam_service_id" "serviceID"  {
		  name = "%s"
		  }

		  resource "ibm_iam_service_policy" "policy" {
			iam_service_id     = ibm_iam_service_id.serviceID.id
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
			description = "IAM Service Profile Policy Custom Hours Creation for test scenario"
		}
	`, name)
}

func testAccCheckIBMIAMServicePolicyUpdateConditions(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_service_id" "serviceID"  {
			name = "%s"
			}

			resource "ibm_iam_service_policy" "policy" {
			iam_service_id     = ibm_iam_service_id.serviceID.id
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
			description = "IAM Service Profile Policy Custom Hours Update for test scenario"
		}
	`, name)
}

func testAccCheckIBMIAMServicePolicyWeeklyAllDay(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_service_id" "serviceID"  {
			name = "%s"
			}

			resource "ibm_iam_service_policy" "policy" {
			iam_service_id     = ibm_iam_service_id.serviceID.id
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
			description = "IAM Service Profile Policy All Day Weekly Time-Based Conditions Creation for test scenario"
		}
	`, name)
}

func testAccCheckIBMIAMServicePolicyTimeBasedOnce(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_service_id" "serviceID"  {
			name = "%s"
			}

			resource "ibm_iam_service_policy" "policy" {
			iam_service_id     = ibm_iam_service_id.serviceID.id
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
			description = "IAM Service Profile Policy Once Time-Based Conditions Creation for test scenario"
		}
	`, name)
}

func testAccCheckIBMIAMServicePolicyWithServiceGroupId(name string) string {
	return fmt.Sprintf(`
			resource "ibm_iam_service_id" "serviceID"  {
				name = "%s"
			}
			resource "ibm_iam_service_policy" "policy" {
				iam_service_id     = ibm_iam_service_id.serviceID.id
				roles           = ["Service ID creator"]
    		resource_attributes {
         		name     = "service_group_id"
         		operator = "stringEquals"
         		value    = "IAM"
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
			description = "IAM Service Profile Policy with service_group_id"
		}
	`, name)
}

func testAccCheckIBMIAMServiceUpdatePolicyWithServiceGroupId(name string) string {
	return fmt.Sprintf(`
			resource "ibm_iam_service_id" "serviceID"  {
				name = "%s"
			}
			resource "ibm_iam_service_policy" "policy" {
				iam_service_id     = ibm_iam_service_id.serviceID.id
				roles           = ["Service ID creator", "User API key creator"]
    		resource_attributes {
         		name     = "service_group_id"
         		operator = "stringEquals"
         		value    = "IAM"
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
			description = "IAM Service Profile Policy with service_group_id"
		}
	`, name)
}

func testAccCheckIBMIAMServicePolicyAttributeBasedCondition(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_service_id" "serviceID"  {
			name = "%s"
		}

		resource "ibm_iam_service_policy" "policy" {
			iam_service_id     = ibm_iam_service_id.serviceID.id
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
			description = "IAM Service Policy Attribute Based Condition Creation for test scenario"
		}
	`, name)
}

func testAccCheckIBMIAMServicePolicyUpdateAttributeBasedCondition(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_service_id" "serviceID"  {
			name = "%s"
		}

		resource "ibm_iam_service_policy" "policy" {
			iam_service_id     = ibm_iam_service_id.serviceID.id
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
			description = "IAM Service Policy Attribute Based Condition Update for test scenario"
		}
	`, name)
}

func testAccCheckIBMIAMServicePolicyResourceAttributesWithoutWildcard(name string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_service_id" "serviceID" {
		name = "%s"
	  }
  
	  resource "ibm_iam_service_policy" "policy" {
		iam_service_id     = ibm_iam_service_id.serviceID.id
		roles              = ["Viewer"]
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
	`, name)
}
