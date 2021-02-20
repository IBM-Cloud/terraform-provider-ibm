/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

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
