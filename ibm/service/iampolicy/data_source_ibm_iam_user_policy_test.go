// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TODO: test fails locally using test env because it returns 3 policies (even existing test)
func TestAccIBMIAMUserPolicyDataSource_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_user_policy.testacc_ds_user_policy", "policies.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicyDataSource_Multiple_Policies(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyDataSourceMultiplePolicies(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_user_policy.testacc_ds_user_policy", "policies.#"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicyDataSource_Service_Specific_Attributes(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyDataSourceServiceSpecificAttributesConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_user_policy.testacc_ds_user_policy", "policies.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicyDataSource_Time_Based_Conditions_Weekly(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyDataSourceTimeBasedWeekly(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_user_policy.policy", "policies.#"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicyDataSource_Time_Based_Conditions_Custom(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyDataSourceTimeBasedCustom(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_user_policy.policy", "policies.#"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicyDataSource_ServiceGroupID(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyDataSourceServiceGroupID(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_user_policy.policy", "policies.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMUserPolicyDataSourceConfig(name string) string {
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

data "ibm_iam_user_policy" "testacc_ds_user_policy" {
  ibm_id = ibm_iam_user_policy.policy.ibm_id
}
`, name, acc.IAMUser)

}

func testAccCheckIBMIAMUserPolicyDataSourceMultiplePolicies(name string) string {
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

data "ibm_resource_group" "group" {
  is_default=true
}

resource "ibm_iam_user_policy" "policy1" {
  ibm_id = "%s"
  roles  = ["Viewer"]

  resources {
    service           = "containers-kubernetes"
    resource_group_id = data.ibm_resource_group.group.id
  }
}


data "ibm_iam_user_policy" "testacc_ds_user_policy" {
  ibm_id = ibm_iam_user_policy.policy.ibm_id
  sort = "-id"
}`, name, acc.IAMUser, acc.IAMUser)

}

func testAccCheckIBMIAMUserPolicyDataSourceServiceSpecificAttributesConfig() string {
	return fmt.Sprintf(`

resource "ibm_iam_user_policy" "policy" {
	ibm_id = "%s"
	roles  = ["Viewer"]
	resource_attributes {
		name     = "serviceName"
		value    = "containers-kubernetes"
	}
	resource_attributes {
		name     = "namespace"
		value    = "test"
	}
}

data "ibm_iam_user_policy" "testacc_ds_user_policy" {
	ibm_id = ibm_iam_user_policy.policy.ibm_id
}
`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyDataSourceTimeBasedWeekly() string {
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
		}

	data "ibm_iam_user_policy" "policy" {
		ibm_id = ibm_iam_user_policy.policy.ibm_id
	}
	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyDataSourceTimeBasedCustom() string {
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
		}

	data "ibm_iam_user_policy" "policy" {
		ibm_id = ibm_iam_user_policy.policy.ibm_id
	}
	`, acc.IAMUser)
}

func testAccCheckIBMIAMUserPolicyDataSourceServiceGroupID() string {
	return fmt.Sprintf(`

	resource "ibm_iam_user_policy" "policy" {
		ibm_id = "%s"
		roles  = ["Service ID creator", "User API key creator"]
		resources {
			service_group_id = "IAM"
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
		}

	data "ibm_iam_user_policy" "policy" {
		ibm_id = ibm_iam_user_policy.policy.ibm_id
	}
	`, acc.IAMUser)
}
