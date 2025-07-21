// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMTrustedProfilePolicyDataSource_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_trusted_profile_policy.policy", "policies.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicyDataSource_Multiple_Policies(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyDataSourceMultiplePolicies(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_trusted_profile_policy.policy", "policies.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicyDataSourceServiceSpecificAttributesConfig(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyDataSourceServiceSpecificAttributesConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_trusted_profile_policy.policy", "policies.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicyDataSource_Time_Based_Conditions_Weekly(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyDataSourceTimeBasedWeekly(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_trusted_profile_policy.policy", "policies.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicyDataSource_Time_Based_Conditions_Custom(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyDataSourceTimeBasedCustom(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_trusted_profile_policy.policy", "policies.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMTrustedProfilePolicyDataSource_ServiceGroupID(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMTrustedProfilePolicyDataSourceServiceGroupID(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_trusted_profile_policy.policy", "policies.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMTrustedProfilePolicyDataSourceConfig(name string) string {
	return fmt.Sprintf(`

resource "ibm_iam_trusted_profile" "profileID" {
  name        = "%s"
  description = "Profile ID for test"
}

resource "ibm_resource_instance" "instance" {
  name     = "%s"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

resource "ibm_iam_trusted_profile_policy" "policy" {
  iam_id = ibm_iam_trusted_profile.profileID.iam_id
  roles          = ["Manager", "Viewer", "Administrator"]

  resources {
    service              = "kms"
    resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
  }
}

data "ibm_iam_trusted_profile_policy" "policy" {
  iam_id = ibm_iam_trusted_profile_policy.policy.iam_id
}`, name, name)

}

func testAccCheckIBMIAMTrustedProfilePolicyDataSourceMultiplePolicies(name string) string {
	return fmt.Sprintf(`

resource "ibm_iam_trusted_profile" "profileID" {
  name        = "%s"
  description = "Profile ID for test"
}

resource "ibm_resource_instance" "instance" {
  name     = "%s"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

resource "ibm_iam_trusted_profile_policy" "policy" {
  iam_id = ibm_iam_trusted_profile.profileID.iam_id
  roles          = ["Manager", "Viewer", "Administrator"]

  resources {
    service              = "kms"
    resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
  }
}

data "ibm_resource_group" "group" {
  is_default=true
}

resource "ibm_iam_trusted_profile_policy" "policy1" {
  iam_id = ibm_iam_trusted_profile.profileID.iam_id
  roles          = ["Viewer"]

  resources {
    service           = "containers-kubernetes"
    resource_group_id = data.ibm_resource_group.group.id
  }
}

data "ibm_iam_trusted_profile_policy" "policy" {
  iam_id = ibm_iam_trusted_profile_policy.policy.iam_id
  sort = "created_at"
}`, name, name)

}

func testAccCheckIBMIAMTrustedProfilePolicyDataSourceServiceSpecificAttributesConfig(name string) string {
	return fmt.Sprintf(`
resource "ibm_iam_trusted_profile" "profileID" {
  name        = "%s"
  description = "Profile ID for test"
}

resource "ibm_iam_trusted_profile_policy" "policy" {
  iam_id = ibm_iam_trusted_profile.profileID.iam_id
  roles          = ["Manager", "Viewer", "Administrator"]

  resource_attributes {
		name     = "serviceName"
		value    = "containers-kubernetes"
	}
	resource_attributes {
		name     = "namespace"
		value    = "test"
	}
}

data "ibm_iam_trusted_profile_policy" "policy" {
  iam_id = ibm_iam_trusted_profile_policy.policy.iam_id
}`, name)

}

func testAccCheckIBMIAMTrustedProfilePolicyDataSourceTimeBasedWeekly(name string) string {
	return fmt.Sprintf(`


	resource "ibm_iam_trusted_profile" "profileID" {
		name        = "%s"
		description = "Profile ID for test"
	}

	resource "ibm_iam_trusted_profile_policy" "policy" {
		iam_id = ibm_iam_trusted_profile.profileID.iam_id
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

	data "ibm_iam_trusted_profile_policy" "policy" {
		iam_id = ibm_iam_trusted_profile_policy.policy.iam_id
	}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyDataSourceTimeBasedCustom(name string) string {
	return fmt.Sprintf(`


	resource "ibm_iam_trusted_profile" "profileID" {
		name        = "%s"
		description = "Profile ID for test"
	}

	resource "ibm_iam_trusted_profile_policy" "policy" {
		iam_id = ibm_iam_trusted_profile.profileID.iam_id
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

	data "ibm_iam_trusted_profile_policy" "policy" {
		iam_id = ibm_iam_trusted_profile_policy.policy.iam_id
	}
	`, name)
}

func testAccCheckIBMIAMTrustedProfilePolicyDataSourceServiceGroupID(name string) string {
	return fmt.Sprintf(`


	resource "ibm_iam_trusted_profile" "profileID" {
		name        = "%s"
		description = "Profile ID for test"
	}

	resource "ibm_iam_trusted_profile_policy" "policy" {
		iam_id = ibm_iam_trusted_profile.profileID.iam_id
		roles  = ["Viewer"]
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

	data "ibm_iam_trusted_profile_policy" "policy" {
		iam_id = ibm_iam_trusted_profile_policy.policy.iam_id
	}
	`, name)
}
