// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamaccessgroup_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMIAMAccessGroup_Basic(t *testing.T) {
	var conf iamaccessgroupsv2.Group
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	updateName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupExists("ibm_iam_access_group.accgroup", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessGroupUpdateWithSameName(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupExists("ibm_iam_access_group.accgroup", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "description", "AccessGroup for test scenario1"),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "tags.#", "3"),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessGroupUpdate(updateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", updateName),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "description", "AccessGroup for test scenario2"),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "tags.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroup_import(t *testing.T) {
	var conf iamaccessgroupsv2.Group
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_access_group.accgroup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupTag(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupExists(resourceName, conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "AccessGroup for test scenario2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIAMAccessGroupDestroy(s *terraform.State) error {
	accClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMAccessGroupsV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_access_group" {
			continue
		}

		agID := rs.Primary.ID

		// Try to find the key
		getAccessGroupOptions := &iamaccessgroupsv2.GetAccessGroupOptions{
			AccessGroupID: &agID,
		}
		_, detailResponse, err := accClient.GetAccessGroup(getAccessGroupOptions)
		if err == nil {
			return fmt.Errorf("Access group still exists: %s", rs.Primary.ID)
		} else if detailResponse.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error waiting for access group (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMAccessGroupExists(n string, obj iamaccessgroupsv2.Group) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		accClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMAccessGroupsV2()
		if err != nil {
			return err
		}
		agID := rs.Primary.ID
		getAccessGroupOptions := &iamaccessgroupsv2.GetAccessGroupOptions{
			AccessGroupID: &agID,
		}
		accgroup, _, err := accClient.GetAccessGroup(getAccessGroupOptions)

		if err != nil {
			return err
		}

		obj = *accgroup
		return nil
	}
}

func testAccCheckIBMIAMAccessGroupBasic(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_access_group" "accgroup" {
			name = "%s"
			tags = ["tag1", "tag2"]
	  	}
	`, name)
}

func testAccCheckIBMIAMAccessGroupUpdateWithSameName(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_access_group" "accgroup" {
			name        = "%s"
			description = "AccessGroup for test scenario1"
			tags        = ["tag1", "tag2", "db"]
	  	}
	`, name)
}

func testAccCheckIBMIAMAccessGroupUpdate(updateName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgroup" {
			name        = "%s"
			description = "AccessGroup for test scenario2"
			tags        = ["tag1"]
	 	}
	`, updateName)
}

func testAccCheckIBMIAMAccessGroupTag(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgroup" {
			name              = "%s"		
			description       = "AccessGroup for test scenario2"
		}
	`, name)
}
