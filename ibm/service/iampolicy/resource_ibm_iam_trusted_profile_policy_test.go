// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMIAMTrustedProfilePolicyBasic(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "description", "IAM Trusted Profile Policy Creation for test scenario"),
				),
			},
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyUpdateRole(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "description", "IAM Trusted Profile Policy Update for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_Service(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyService(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "resources.0.service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyUpdateServiceAndRegion(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "resources.0.service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "resources.0.region", "us-south"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_ServiceType(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyServiceType(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "resources.0.service_type", "service"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_ResourceInstance(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyResourceInstance(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_Resource_Group(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyResourceGroup(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "resources.0.service", "containers-kubernetes"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_Resource_Type(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyResourceType(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_import(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_trusted_profile_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
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

func TestAccIBMIAMTrustedProfilePolicy_account_management(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_trusted_profile_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyAccountManagement(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "account_management", "true"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicyWithCustomRole(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	crName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyWithCustomRole(name, crName, displayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_Resource_Attributes(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyResourceAttributes(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "resource_attributes.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyResourceAttributesUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "resource_attributes.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_Resource_Attributes_Without_Wildcard(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyResourceAttributesWithoutWildcard(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "resource_attributes.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_Resource_Tags(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyResourceTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "resource_tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyUpdateResourceTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "resource_tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_Transaction_Id(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyTransactionId(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "resources.0.service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "transaction_id", "terrformTrustedPolicy"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_Time_Based_Conditions_Weekly_Custom(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyWeeklyCustomHours(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "pattern", "time-based-conditions:weekly:custom-hours"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "rule_conditions.#", "3"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "rule_conditions.2.key", "{{environment.attributes.day_of_week}}"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "rule_conditions.2.value.#", "5"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "description", "IAM Trusted Profile Policy Custom Hours Creation for test scenario"),
				),
			},
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyUpdateConditions(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "pattern", "time-based-conditions:weekly:custom-hours"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "rule_conditions.2.value.#", "4"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "description", "IAM Trusted Profile Policy Custom Hours Update for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_Time_Based_Conditions_Weekly_All_Day(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyWeeklyAllDay(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "pattern", "time-based-conditions:weekly:all-day"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "rule_conditions.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "rule_conditions.0.key", "{{environment.attributes.day_of_week}}"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "rule_conditions.0.value.#", "5"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "description", "IAM Trusted Profile Policy All Day Weekly Time-Based Conditions Creation for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_Time_Based_Conditions_Once(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyTimeBasedOnce(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "pattern", "time-based-conditions:once"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "rule_conditions.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "rule_conditions.0.key", "{{environment.attributes.current_date_time}}"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "description", "IAM Trusted Profile Policy Once Time-Based Conditions Creation for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_Update_To_Time_Based_Conditions(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyResourceAttributes(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "resource_attributes.#", "2"),
				),
			},
			{
				Config:      testAccCheckIBMIAMTrustedProfilePolicyUpdateConditions(name),
				ExpectError: regexp.MustCompile("Error: Cannot use rule_conditions, rule_operator, or pattern when updating v1/policy. Delete existing v1/policy and create using rule_conditions and pattern."),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_ServiceGroupID(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyWithServiceGroupId(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "resource_attributes.0.value", "IAM"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyUpdateWithServiceGroupId(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicy_With_Attribute_Based_Condition(t *testing.T) {
	var conf iampolicymanagementv1.V2PolicyTemplateMetaData
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMTrustedProfilePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyAttributeBasedCondition(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMTrustedProfilePolicyExists("ibm_iam_trusted_profile_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "pattern", "attribute-based-condition:resource:literal-and-wildcard"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "description", "IAM Trusted Profile Policy Attribute Based Condition Creation for test scenario"),
				),
			},
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyUpdateAttributeBasedCondition(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile.profileID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "pattern", "attribute-based-condition:resource:literal-and-wildcard"),
					resource.TestCheckResourceAttr("ibm_iam_trusted_profile_policy.policy", "description", "IAM Trusted Profile Policy Attribute Based Condition Update for test scenario"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMTrustedProfilePolicyDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_trusted_profile_policy" {
			continue
		}
		policyID := rs.Primary.ID
		parts, err := flex.IdParts(policyID)
		if err != nil {
			return err
		}
		profilePolicyID := parts[1]

		getPolicyOptions := rsContClient.NewGetV2PolicyOptions(
			profilePolicyID,
		)

		// Try to find the key
		destroyedPolicy, response, err := rsContClient.GetV2Policy(getPolicyOptions)

		if err == nil && *destroyedPolicy.State != "deleted" {
			return fmt.Errorf("Trusted Profile policy still exists: %s\n", rs.Primary.ID)
		} else if response.StatusCode != 404 && destroyedPolicy.State != nil && *destroyedPolicy.State != "deleted" {
			return fmt.Errorf("[ERROR] Error waiting for trusted profile policy (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMTrustedProfilePolicyExists(n string, obj iampolicymanagementv1.V2PolicyTemplateMetaData) resource.TestCheckFunc {

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
		profilePolicyID := parts[1]

		getPolicyOptions := rsContClient.NewGetV2PolicyOptions(
			profilePolicyID,
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

func testAccCheckIBMIAMTrustedProfilePolicyBasic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
			roles          = ["Viewer"]
			tags           = ["tag1"]
			description    = "IAM Trusted Profile Policy Creation for test scenario"
	  	}

	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyUpdateRole(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
			roles          = ["Viewer", "Manager"]
			tags           = ["tag1", "tag2"]
			description    = "IAM Trusted Profile Policy Update for test scenario"
	  	}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyService(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
			roles          = ["Viewer"]
	  
			resources {
		 	 service = "cloudantnosqldb"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyServiceType(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
			roles          = ["Viewer"]
	  
			resources {
				service_type = "service"
				region = "us-south"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyUpdateServiceAndRegion(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
			roles          = ["Viewer", "Manager"]
	  
			resources {
		  		service = "cloudantnosqldb"
		  		region  = "us-south"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyResourceInstance(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_resource_instance" "instance" {
			name     = "%s"
			service  = "kms"
			plan     = "tiered-pricing"
			location = "us-south"
	  	}
	  
	  	resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
			roles          = ["Manager", "Viewer", "Administrator"]
	  
			resources {
		 		 service              = "kms"
		  		resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
			}
	  	}
		  

	`, name, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyResourceGroup(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
	  	}
	  
	  	data "ibm_resource_group" "group" {
			is_default=true
	  	}
	  
	  	resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
			roles          = ["Viewer"]
	  
			resources {
		 		service           = "containers-kubernetes"
		  		resource_group_id = data.ibm_resource_group.group.id
			}
	  	}
		  

	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyResourceType(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
	  	}
	  
	  	data "ibm_resource_group" "group" {
			is_default=true
	  	}
	  
	  	resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
			roles          = ["Administrator"]
	  
			resources {
		  		resource_type = "resource-group"
		  		resource      = data.ibm_resource_group.group.id
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyImport(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
			roles          = ["Viewer"]
	  	}

	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyAccountManagement(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id     = ibm_iam_trusted_profile.profileID.id
			roles              = ["Viewer"]
			account_management = true
	  	}

	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyWithCustomRole(name, crName, displayName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
		}
		resource "ibm_iam_custom_role" "customrole" {
			name         = "%s"
			display_name = "%s"
			description  = "role for test scenario1"
			service = "kms"
			actions      = ["kms.secrets.rotate"]
		}
	  
	  	resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
			roles          = [ibm_iam_custom_role.customrole.display_name,"Viewer"]
			tags           = ["tag1"]
			resources {
				service           = "kms"
		   }
	  	}

	`, name, crName, displayName)
}

func testAccCheckIBMIAMTrustedProfilePolicyResourceAttributes(name string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_trusted_profile" "profileID" {
		name = "%s"
	  }
  
	  resource "ibm_iam_trusted_profile_policy" "policy" {
		profile_id     = ibm_iam_trusted_profile.profileID.id
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

func testAccCheckIBMIAMTrustedProfilePolicyResourceAttributesWithoutWildcard(name string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_trusted_profile" "profileID" {
		name = "%s"
	  }
  
	  resource "ibm_iam_trusted_profile_policy" "policy" {
		profile_id     = ibm_iam_trusted_profile.profileID.id
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

func testAccCheckIBMIAMTrustedProfilePolicyResourceAttributesUpdate(name string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_trusted_profile" "profileID" {
		name = "%s"
	  }
  
	  resource "ibm_iam_trusted_profile_policy" "policy" {
		profile_id     = ibm_iam_trusted_profile.profileID.id
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

func testAccCheckIBMIAMTrustedProfilePolicyResourceTags(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
			roles          = ["Viewer"]
			resource_tags {
				name = "one"
				value = "Terraform"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyTransactionId(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
			roles          = ["Viewer"]
			transaction_id = "terrformTrustedPolicy"
	  
			resources {
		 	 service = "cloudantnosqldb"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyUpdateResourceTags(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
			roles          = ["Viewer"]
			resource_tags {
				name = "one"
				value = "Terraform"
			}
			resource_tags {
				name = "two"
				value = "TerraformUpdate"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyWeeklyCustomHours(name string) string {
	return fmt.Sprintf(`
	  resource "ibm_iam_trusted_profile" "profileID" {
		  name = "%s"
		  }

		  resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
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
			description = "IAM Trusted Profile Policy Custom Hours Creation for test scenario"
		}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyUpdateConditions(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
			}

			resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
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
			description = "IAM Trusted Profile Policy Custom Hours Update for test scenario"
		}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyWeeklyAllDay(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
			}

			resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
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
			description = "IAM Trusted Profile Policy All Day Weekly Time-Based Conditions Creation for test scenario"
		}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyTimeBasedOnce(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
			}

			resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
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
			description = "IAM Trusted Profile Policy Once Time-Based Conditions Creation for test scenario"
		}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyWithServiceGroupId(name string) string {
	return fmt.Sprintf(`
			resource "ibm_iam_trusted_profile" "profileID" {
				name = "%s"
			}
			resource "ibm_iam_trusted_profile_policy" "policy" {
				profile_id = ibm_iam_trusted_profile.profileID.id
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

func testAccCheckIBMIAMTrustedProfilePolicyUpdateWithServiceGroupId(name string) string {
	return fmt.Sprintf(`
			resource "ibm_iam_trusted_profile" "profileID" {
				name = "%s"
			}
			resource "ibm_iam_trusted_profile_policy" "policy" {
				profile_id = ibm_iam_trusted_profile.profileID.id
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

func testAccCheckIBMIAMTrustedProfilePolicyAttributeBasedCondition(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
		}

		resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
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
			description = "IAM Trusted Profile Policy Attribute Based Condition Creation for test scenario"
		}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyUpdateAttributeBasedCondition(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_trusted_profile" "profileID" {
			name = "%s"
		}

		resource "ibm_iam_trusted_profile_policy" "policy" {
			profile_id = ibm_iam_trusted_profile.profileID.id
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
			description = "IAM Trusted Profile Policy Attribute Based Condition Update for test scenario"
		}
	`, name)
}
