// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMAccessGroupPolicyDataSource_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupPolicyDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_access_group_policy.testacc_ds_access_group_policy", "policies.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicyDataSource_Multiple_Policies(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupPolicyDataSourceMultiplePolicies(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_access_group_policy.testacc_ds_access_group_policy", "policies.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMAccessGroupPolicyDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_resource_instance" "instance" {
	name     = "%s"
	service  = "kms"
	plan     = "tiered-pricing"
	location = "us-south"
}

resource "ibm_iam_access_group" "accgrp" {
	name = "%s"
}

resource "ibm_iam_access_group_policy" "policy" {
	access_group_id = ibm_iam_access_group.accgrp.id
	roles  = ["Manager", "Viewer", "Administrator"]
	
	resources {
		service              = "kms"
		resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
	}
}

data "ibm_iam_access_group_policy" "testacc_ds_access_group_policy" {
	access_group_id = ibm_iam_access_group_policy.policy.access_group_id
}
`, name, name)

}

func testAccCheckIBMIAMAccessGroupPolicyDataSourceMultiplePolicies(name string) string {
	return fmt.Sprintf(`

resource "ibm_resource_instance" "instance" {
	name     = "%s"
	service  = "kms"
	plan     = "tiered-pricing"
	location = "us-south"
}

resource "ibm_iam_access_group" "accgrp" {
	name = "%s"
}

resource "ibm_iam_access_group_policy" "policy" {
	access_group_id = ibm_iam_access_group.accgrp.id
	roles  = ["Manager", "Viewer", "Administrator"]
	resources {
		service              = "kms"
		resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
	}
}

data "ibm_resource_group" "group" {
	is_default=true
}

resource "ibm_iam_access_group_policy" "policy1" {
	access_group_id = ibm_iam_access_group.accgrp.id
	roles  = ["Viewer"]

	resources {
		service           = "containers-kubernetes"
		resource_group_id = data.ibm_resource_group.group.id
	}
}


data "ibm_iam_access_group_policy" "testacc_ds_access_group_policy" {
	access_group_id = ibm_iam_access_group_policy.policy.access_group_id
	sort = "-id"
}`, name, name)

}
