// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudfoundry_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/Mavrickk3/bluemix-go/api/mccp/mccpv2"
	"github.com/Mavrickk3/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMOrg_Basic(t *testing.T) {
	var conf mccpv2.OrganizationFields
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMOrgDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMOrgCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMOrgExists("ibm_org.testacc_org", &conf),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "name", name),
				),
			},
			{
				Config: testAccCheckIBMOrgUpdate(updatedName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "name", updatedName),
				),
			},
		},
	})
}

func TestAccIBMOrg_Basic_Import(t *testing.T) {
	var conf mccpv2.OrganizationFields
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_org.testacc_org"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMOrgDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMOrgCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMOrgExists(resourceName, &conf),
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

func TestAccIBMOrg_with_roles(t *testing.T) {
	var conf mccpv2.OrganizationFields
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMOrgDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMOrgCreateWithRoles(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMOrgExists("ibm_org.testacc_org", &conf),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "name", name),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "auditors.#", "1"),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "managers.#", "1"),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "billing_managers.#", "1"),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "users.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMOrgUpdateWithRoles(updatedName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "auditors.#", "1"),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "managers.#", "2"),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "billing_managers.#", "1"),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "users.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMOrg_With_Tags(t *testing.T) {
	var conf mccpv2.OrganizationFields
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMOrgDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMOrgWithTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMOrgExists("ibm_org.testacc_org", &conf),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "name", name),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMOrgWithUpdatedTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "tags.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMOrgExists(n string, obj *mccpv2.OrganizationFields) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		cfClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MccpAPI()
		if err != nil {
			return err
		}
		orgGUID := rs.Primary.ID
		org, err := cfClient.Organizations().Get(orgGUID)
		if err != nil {
			return err
		}
		*obj = *org
		return nil
	}
}

func testAccCheckIBMOrgDestroy(s *terraform.State) error {
	cfClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_org" {
			continue
		}
		orgGUID := rs.Primary.ID
		_, err := cfClient.Organizations().Get(orgGUID)
		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("[ERROR] Error waiting for Organization (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckIBMOrgCreate(name string) string {
	return fmt.Sprintf(`
resource "ibm_org" "testacc_org" {
    name = "%s"
}`, name)
}

func testAccCheckIBMOrgUpdate(updatedName string) string {
	return fmt.Sprintf(`	
resource "ibm_org" "testacc_org" {
	name = "%s"
}`, updatedName)
}

func testAccCheckIBMOrgCreateWithRoles(name string) string {
	return fmt.Sprintf(`
	resource "ibm_org" "testacc_org" {
		name             = "%s"
		auditors         = ["%s"]
		managers         = ["%s"]
		billing_managers = ["%s"]
		users            = ["%s"]
	  }
	  
`, name, acc.Ibmid1, acc.Ibmid1, acc.Ibmid1, acc.Ibmid1)
}

func testAccCheckIBMOrgUpdateWithRoles(updatedName string) string {
	return fmt.Sprintf(`
	resource "ibm_org" "testacc_org" {
		name             = "%s"
		auditors         = ["%s"]
		managers         = ["%s", "%s"]
		billing_managers = ["%s"]
		users            = ["%s"]
	  }
	  
`, updatedName, acc.Ibmid2, acc.Ibmid2, acc.Ibmid1, acc.Ibmid2, acc.Ibmid2)
}

func testAccCheckIBMOrgWithTags(name string) string {
	return fmt.Sprintf(`
resource "ibm_org" "testacc_org" {
	name = "%s"
	tags = ["one"]
}
`, name)
}

func testAccCheckIBMOrgWithUpdatedTags(name string) string {
	return fmt.Sprintf(`
resource "ibm_org" "testacc_org" {
	name = "%s"
	tags = ["one", "two"]
}
`, name)
}
