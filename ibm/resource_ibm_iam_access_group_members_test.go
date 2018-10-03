package ibm

import (
	"fmt"
	"testing"

	"strings"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMIAMAccessGroupMember_Basic(t *testing.T) {

	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	sname := fmt.Sprintf("terraform_%d", acctest.RandInt())
	sname1 := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupMemberDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupMember_basic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_members.accgroupmem", "members.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupMember_Add_ServiceID(name, sname),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_members.accgroupmem", "members.#", "2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupMember_Add_Another_ServiceID(name, sname, sname1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_members.accgroupmem", "members.#", "3"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupMember_Remove_User_And_SID(name, sname, sname1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_members.accgroupmem", "members.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupMember_import(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	sname := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_iam_access_group.accgroup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupMemberDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupMember_Import(name, sname),
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
	accClient, err := testAccProvider.Meta().(ClientSession).IAMUUMAPI()
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

func testAccCheckIBMIAMAccessGroupMember_basic(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_access_group" "accgroup" {
			name              = "%s"		
		}

		resource "ibm_iam_access_group_members" "accgroupmem" {
			access_group_id   = "${ibm_iam_access_group.accgroup.id}"		
			ibm_ids           = ["%s"]
		}
	`, name, IAMUser)
}

func testAccCheckIBMIAMAccessGroupMember_Add_ServiceID(name, sname string) string {
	return fmt.Sprintf(`

	resource "ibm_iam_access_group" "accgroup" {
		name              = "%s"
	}

	resource "ibm_iam_service_id" "serviceID" {
		name              = "%s"
	}

	resource "ibm_iam_access_group_members" "accgroupmem" {
		access_group_id   = "${ibm_iam_access_group.accgroup.id}"
		ibm_ids           = ["%s"]
		iam_service_ids   = ["${ibm_iam_service_id.serviceID.id}"]
	}
	`, name, sname, IAMUser)
}

func testAccCheckIBMIAMAccessGroupMember_Add_Another_ServiceID(name, sname, sname1 string) string {
	return fmt.Sprintf(`

	resource "ibm_iam_access_group" "accgroup" {
		name              = "%s"
	}

	resource "ibm_iam_service_id" "serviceID" {
		name              = "%s"
	}

	resource "ibm_iam_service_id" "serviceID2" {
		name              = "%s"
	}


	resource "ibm_iam_access_group_members" "accgroupmem" {
		access_group_id   = "${ibm_iam_access_group.accgroup.id}"
		ibm_ids           = ["%s"]
		iam_service_ids   = ["${ibm_iam_service_id.serviceID.id}","${ibm_iam_service_id.serviceID2.id}"]
	}
	`, name, sname, sname1, IAMUser)
}

func testAccCheckIBMIAMAccessGroupMember_Remove_User_And_SID(name, sname, sname1 string) string {
	return fmt.Sprintf(`

	resource "ibm_iam_access_group" "accgroup" {
		name              = "%s"
	}

	resource "ibm_iam_service_id" "serviceID" {
		name              = "%s"
	}

	resource "ibm_iam_service_id" "serviceID2" {
		name              = "%s"
	}

	resource "ibm_iam_access_group_members" "accgroupmem" {
		access_group_id   = "${ibm_iam_access_group.accgroup.id}"
		iam_service_ids   = ["${ibm_iam_service_id.serviceID.id}"]
	}
	`, name, sname, sname1)
}

func testAccCheckIBMIAMAccessGroupMember_Import(name, sname string) string {
	return fmt.Sprintf(`

	resource "ibm_iam_access_group" "accgroup" {
		name              = "%s"		
	}

	resource "ibm_iam_service_id" "serviceID" {
		name              = "%s"		
	}

	resource "ibm_iam_access_group_members" "accgroupmem" {
		access_group_id   = "${ibm_iam_access_group.accgroup.id}"	
		ibm_ids           = ["%s"]	
		iam_service_ids   = ["${ibm_iam_service_id.serviceID.id}"]
	}
	`, name, sname, IAMUser)
}
