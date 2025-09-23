// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudfoundry_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/Mavrickk3/bluemix-go/api/mccp/mccpv2"
	"github.com/Mavrickk3/bluemix-go/bmxerror"
)

func TestAccIBMSpace_Basic(t *testing.T) {
	var conf mccpv2.SpaceFields
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSpaceDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMSpaceCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSpaceExists("ibm_space.space", &conf),
					resource.TestCheckResourceAttr("ibm_space.space", "org", acc.CfOrganization),
					resource.TestCheckResourceAttr("ibm_space.space", "name", name),
				),
			},

			{
				Config: testAccCheckIBMSpaceUpdate(updatedName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_space.space", "org", acc.CfOrganization),
					resource.TestCheckResourceAttr("ibm_space.space", "name", updatedName),
				),
			},
		},
	})
}

func TestAccIBMSpace_Basic_Import(t *testing.T) {
	var conf mccpv2.SpaceFields
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_space.space"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSpaceDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMSpaceCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSpaceExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "org", acc.CfOrganization),
					resource.TestCheckResourceAttr(resourceName, "name", name),
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

func TestAccIBMSpace_with_roles(t *testing.T) {
	var conf mccpv2.SpaceFields
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSpaceDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMSpaceCreateWithRoles(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSpaceExists("ibm_space.space", &conf),
					resource.TestCheckResourceAttr("ibm_space.space", "org", acc.CfOrganization),
					resource.TestCheckResourceAttr("ibm_space.space", "name", name),
					resource.TestCheckResourceAttr("ibm_space.space", "auditors.#", "1"),
					resource.TestCheckResourceAttr("ibm_space.space", "managers.#", "1"),
					resource.TestCheckResourceAttr("ibm_space.space", "developers.#", "1"),
				),
			},

			{
				Config: testAccCheckIBMSpaceUpdateWithRoles(updatedName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_space.space", "org", acc.CfOrganization),
					resource.TestCheckResourceAttr("ibm_space.space", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_space.space", "auditors.#", "1"),
					resource.TestCheckResourceAttr("ibm_space.space", "managers.#", "2"),
					resource.TestCheckResourceAttr("ibm_space.space", "developers.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMSpace_With_Tags(t *testing.T) {
	var conf mccpv2.SpaceFields
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSpaceDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMSpaceWithTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSpaceExists("ibm_space.space", &conf),
					resource.TestCheckResourceAttr("ibm_space.space", "org", acc.CfOrganization),
					resource.TestCheckResourceAttr("ibm_space.space", "name", name),
					resource.TestCheckResourceAttr("ibm_space.space", "tags.#", "1"),
				),
			},

			{
				Config: testAccCheckIBMSpaceWithUpdatedTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_space.space", "org", acc.CfOrganization),
					resource.TestCheckResourceAttr("ibm_space.space", "tags.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMSpaceExists(n string, obj *mccpv2.SpaceFields) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cfClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MccpAPI()
		if err != nil {
			return err
		}
		spaceGUID := rs.Primary.ID

		space, err := cfClient.Spaces().Get(spaceGUID)
		if err != nil {
			return err
		}

		*obj = *space
		return nil
	}
}

func testAccCheckIBMSpaceDestroy(s *terraform.State) error {
	cfClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MccpAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_space" {
			continue
		}

		spaceGUID := rs.Primary.ID
		_, err := cfClient.Spaces().Get(spaceGUID)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("[ERROR] Error waiting for Space (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckIBMSpaceCreate(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_space" "space" {
    org = "%s"
	name = "%s"
}`, acc.CfOrganization, name)

}

func testAccCheckIBMSpaceUpdate(updatedName string) string {
	return fmt.Sprintf(`
	
resource "ibm_space" "space" {
    org = "%s"
	name = "%s"
}`, acc.CfOrganization, updatedName)

}

func testAccCheckIBMSpaceCreateWithRoles(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_space" "space" {
    org = "%s"
	name = "%s"
	auditors = ["%s"]
	managers = ["%s"]
	developers = ["%s"]
}
`, acc.CfOrganization, name, acc.Ibmid1, acc.Ibmid1, acc.Ibmid1)

}

func testAccCheckIBMSpaceUpdateWithRoles(updatedName string) string {
	return fmt.Sprintf(`
resource "ibm_space" "space" {
    org = "%s"
	name = "%s"
	auditors = ["%s"]
	managers = ["%s", "%s"]
	developers = ["%s"]
}`, acc.CfOrganization, updatedName, acc.Ibmid2, acc.Ibmid2, acc.Ibmid1, acc.Ibmid2)

}

func testAccCheckIBMSpaceWithTags(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_space" "space" {
    org = "%s"
	name = "%s"
	tags = ["one"]
}`, acc.CfOrganization, name)

}
func testAccCheckIBMSpaceWithUpdatedTags(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_space" "space" {
    org = "%s"
	name = "%s"
	tags = ["one", "two"]
}`, acc.CfOrganization, name)

}
