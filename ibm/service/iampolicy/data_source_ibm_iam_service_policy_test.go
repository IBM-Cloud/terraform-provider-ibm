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

func TestAccIBMIAMServicePolicyDataSource_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_service_policy.testacc_ds_service_policy", "policies.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicyDataSource_Multiple_Policies(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyDataSourceMultiplePolicies(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_service_policy.testacc_ds_service_policy", "policies.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicyDataSourceServiceSpecificAttributesConfig(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyDataSourceServiceSpecificAttributesConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_service_policy.testacc_ds_service_policy", "policies.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMServicePolicyDataSourceConfig(name string) string {
	return fmt.Sprintf(`

resource "ibm_iam_service_id" "serviceID" {
  name        = "%s"
  description = "Service ID for test"
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

data "ibm_iam_service_policy" "testacc_ds_service_policy" {
  iam_service_id = ibm_iam_service_policy.policy.iam_service_id
}`, name, name)

}

func testAccCheckIBMIAMServicePolicyDataSourceMultiplePolicies(name string) string {
	return fmt.Sprintf(`

resource "ibm_iam_service_id" "serviceID" {
  name        = "%s"
  description = "Service ID for test"
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

data "ibm_resource_group" "group" {
  is_default=true
}

resource "ibm_iam_service_policy" "policy1" {
  iam_service_id = ibm_iam_service_id.serviceID.id
  roles          = ["Viewer"]

  resources {
    service           = "containers-kubernetes"
    resource_group_id = data.ibm_resource_group.group.id
  }
}


data "ibm_iam_service_policy" "testacc_ds_service_policy" {
  iam_service_id = ibm_iam_service_policy.policy.iam_service_id
  sort = "id"
}`, name, name)

}

func testAccCheckIBMIAMServicePolicyDataSourceServiceSpecificAttributesConfig(name string) string {
	return fmt.Sprintf(`

resource "ibm_iam_service_id" "serviceID" {
  name        = "%s"
  description = "Service ID for test"
}

resource "ibm_iam_service_policy" "policy" {
  iam_service_id = ibm_iam_service_id.serviceID.id
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

data "ibm_iam_service_policy" "testacc_ds_service_policy" {
  iam_service_id = ibm_iam_service_policy.policy.iam_service_id
}`, name)

}
