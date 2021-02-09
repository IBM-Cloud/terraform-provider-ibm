/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMIAMRoleDataSourcebasic(t *testing.T) {
	serviceName := "kms"
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMRoleConfig(name, serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_roles.test", "service", serviceName),
				),
			},
		},
	})
}

func testAccCheckIBMIAMRoleConfig(name, serviceName string) string {
	return fmt.Sprintf(`


resource "ibm_iam_access_group" "accgrp" {
	name = "%s"
}

data "ibm_iam_roles" "test" {
	service = "%s"
  }

resource "ibm_iam_access_group_policy" "policy" {
	access_group_id = ibm_iam_access_group.accgrp.id
	roles           = [data.ibm_iam_roles.test.roles.4.name,"Viewer"]
	tags            = ["tag1"]
	resources {
	  service = "kms"
	}
}

`, name, serviceName)
}
