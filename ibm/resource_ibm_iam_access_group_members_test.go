// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMIAMAccessGroupMember_Basic(t *testing.T) {

	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	sname := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	sname1 := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupMemberDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupMemberBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_members.accgroupmem", "members.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupMemberAddServiceID(name, sname),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_members.accgroupmem", "members.#", "2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupMemberAddAnotherServiceID(name, sname, sname1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_members.accgroupmem", "members.#", "3"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupMemberRemoveUserAndSID(name, sname, sname1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_members.accgroupmem", "members.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupMember_import(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	sname := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_access_group.accgroup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupMemberDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupMemberImport(name, sname),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"iam_ids",
					"iam_service_ids",
				},
			},
		},
	})
}

func testAccCheckIBMIAMAccessGroupMemberDestroy(s *terraform.State) error {
	accClient, err := testAccProvider.Meta().(ClientSession).IAMUUMAPIV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_access_group_members" {
			continue
		}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		grpID := parts[0]

		// Try to find the members
		_, err = accClient.AccessGroupMember().List(grpID)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for access group members (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMAccessGroupMemberBasic(name string) string {
	return fmt.Sprintf(`
		
	resource "ibm_iam_access_group" "accgroup" {
  		name = "%s"
	}

	resource "ibm_iam_access_group_members" "accgroupmem" {
  		access_group_id = ibm_iam_access_group.accgroup.id
  		ibm_ids         = ["%s"]
	}`, name, IAMUser)
}

func testAccCheckIBMIAMAccessGroupMemberAddServiceID(name, sname string) string {
	return fmt.Sprintf(`

	resource "ibm_iam_access_group" "accgroup" {
  		name = "%s"
	}

	resource "ibm_iam_service_id" "serviceID" {
  		name = "%s"
	}

	resource "ibm_iam_access_group_members" "accgroupmem" {
  		access_group_id = ibm_iam_access_group.accgroup.id
  		ibm_ids         = ["%s"]
  		iam_service_ids = [ibm_iam_service_id.serviceID.id]
	}`, name, sname, IAMUser)
}

func testAccCheckIBMIAMAccessGroupMemberAddAnotherServiceID(name, sname, sname1 string) string {
	return fmt.Sprintf(`

	resource "ibm_iam_access_group" "accgroup" {
		name = "%s"
	  }
	  
	  resource "ibm_iam_service_id" "serviceID" {
		name = "%s"
	  }
	  
	  resource "ibm_iam_service_id" "serviceID2" {
		name = "%s"
	  }
	  
	  resource "ibm_iam_access_group_members" "accgroupmem" {
		access_group_id = ibm_iam_access_group.accgroup.id
		ibm_ids         = ["%s"]
		iam_service_ids = [ibm_iam_service_id.serviceID.id, ibm_iam_service_id.serviceID2.id]
	  }`, name, sname, sname1, IAMUser)
}

func testAccCheckIBMIAMAccessGroupMemberRemoveUserAndSID(name, sname, sname1 string) string {
	return fmt.Sprintf(`

	resource "ibm_iam_access_group" "accgroup" {
  		name = "%s"
	}

	resource "ibm_iam_service_id" "serviceID" {
  		name = "%s"
	}

	resource "ibm_iam_service_id" "serviceID2" {
 		name = "%s"
	}

	resource "ibm_iam_access_group_members" "accgroupmem" {
  		access_group_id = ibm_iam_access_group.accgroup.id
  		iam_service_ids = [ibm_iam_service_id.serviceID.id]
	}`, name, sname, sname1)
}

func testAccCheckIBMIAMAccessGroupMemberImport(name, sname string) string {
	return fmt.Sprintf(`

	resource "ibm_iam_access_group" "accgroup" {
		name = "%s"
	}
	  
	resource "ibm_iam_service_id" "serviceID" {
		name = "%s"
	}
	  
	resource "ibm_iam_access_group_members" "accgroupmem" {
		access_group_id = ibm_iam_access_group.accgroup.id
		ibm_ids         = ["%s"]
		iam_service_ids = [ibm_iam_service_id.serviceID.id]
	}`, name, sname, IAMUser)
}
