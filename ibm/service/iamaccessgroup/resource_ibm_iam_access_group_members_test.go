// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamaccessgroup_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMIAMAccessGroupMember_Basic(t *testing.T) {

	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	sname := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	sname1 := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	pname := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	pname1 := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupMemberBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_members.accgroupmem", "members.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessGroupMemberAddServiceID(name, sname),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_members.accgroupmem", "members.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessGroupMemberAddAnotherServiceID(name, sname, sname1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_members.accgroupmem", "members.#", "3"),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessGroupMemberAddProfileID(name, sname, sname1, pname),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_members.accgroupmem", "members.#", "4"),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessGroupMemberAddAnotherProfileID(name, sname, sname1, pname, pname1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_members.accgroupmem", "members.#", "5"),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessGroupMemberRemoveUserSIDAndProfileID(name, sname, sname1, pname, pname1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_members.accgroupmem", "members.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupMember_import(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	sname := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	pname := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_access_group.accgroup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupMemberImport(name, sname, pname),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"iam_ids",
					"iam_service_ids",
					"iam_profile_ids",
				},
			},
		},
	})
}

func testAccCheckIBMIAMAccessGroupMemberDestroy(s *terraform.State) error {
	accClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMAccessGroupsV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_access_group_members" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		grpID := parts[0]

		// Try to find the members
		listAccessGroupMembersOptions := &iamaccessgroupsv2.ListAccessGroupMembersOptions{
			AccessGroupID: &grpID,
		}
		_, detailResponse, err := accClient.ListAccessGroupMembers(listAccessGroupMembersOptions)

		if err != nil && detailResponse.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error waiting for access group members (%s) to be destroyed: %s", rs.Primary.ID, err)
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
	}`, name, acc.IAMUser)
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
	}`, name, sname, acc.IAMUser)
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
	}`, name, sname, sname1, acc.IAMUser)
}

func testAccCheckIBMIAMAccessGroupMemberAddProfileID(name, sname, sname1, pname string) string {
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

	resource "ibm_iam_trusted_profile" "profileID" {
  		name = "%s"
	}

	resource "ibm_iam_access_group_members" "accgroupmem" {
  		access_group_id = ibm_iam_access_group.accgroup.id
  		ibm_ids         = ["%s"]
		iam_service_ids = [ibm_iam_service_id.serviceID.id, ibm_iam_service_id.serviceID2.id]
  		iam_profile_ids = [ibm_iam_trusted_profile.profileID.id]
	}`, name, sname, sname1, pname, acc.IAMUser)
}

func testAccCheckIBMIAMAccessGroupMemberAddAnotherProfileID(name, sname, sname1, pname, pname1 string) string {
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

	resource "ibm_iam_trusted_profile" "profileID" {
  		name = "%s"
	}

	resource "ibm_iam_trusted_profile" "profileID2" {
		name = "%s"
    }

	resource "ibm_iam_access_group_members" "accgroupmem" {
  		access_group_id = ibm_iam_access_group.accgroup.id
  		ibm_ids         = ["%s"]
		iam_service_ids = [ibm_iam_service_id.serviceID.id, ibm_iam_service_id.serviceID2.id]
  		iam_profile_ids = [ibm_iam_trusted_profile.profileID.id, ibm_iam_trusted_profile.profileID2.id]
	}`, name, sname, sname1, pname, pname1, acc.IAMUser)
}

func testAccCheckIBMIAMAccessGroupMemberRemoveUserSIDAndProfileID(name, sname, sname1, pname, pname1 string) string {
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

	resource "ibm_iam_trusted_profile" "profileID" {
		name = "%s"
    }

	resource "ibm_iam_trusted_profile" "profileID2" {
		name = "%s"
    }
	
	resource "ibm_iam_access_group_members" "accgroupmem" {
  		access_group_id = ibm_iam_access_group.accgroup.id
  		iam_service_ids = [ibm_iam_service_id.serviceID.id]
		iam_profile_ids = [ibm_iam_trusted_profile.profileID.id]
	}`, name, sname, sname1, pname, pname1)
}

func testAccCheckIBMIAMAccessGroupMemberImport(name, sname, pname string) string {
	return fmt.Sprintf(`

	resource "ibm_iam_access_group" "accgroup" {
		name = "%s"
	}
	  
	resource "ibm_iam_service_id" "serviceID" {
		name = "%s"
	}

	resource "ibm_iam_trusted_profile" "profileID" {
		name = "%s"
	}
	  
	resource "ibm_iam_access_group_members" "accgroupmem" {
		access_group_id = ibm_iam_access_group.accgroup.id
		ibm_ids         = ["%s"]
		iam_service_ids = [ibm_iam_service_id.serviceID.id]
		iam_profile_ids = [ibm_iam_trusted_profile.profileID.id]
	}`, name, sname, pname, acc.IAMUser)
}
