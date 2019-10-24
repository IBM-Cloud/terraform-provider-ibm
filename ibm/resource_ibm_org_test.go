package ibm

import (
	"fmt"
	"testing"

	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMOrg_Basic(t *testing.T) {
	var conf mccpv2.OrganizationFields
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMOrgDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMOrgCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMOrgExists("ibm_org.testacc_org", &conf),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "name", name),
				),
			},
			resource.TestStep{
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
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_org.testacc_org"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMOrgDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMOrgCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMOrgExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMOrg_with_roles(t *testing.T) {
	var conf mccpv2.OrganizationFields
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMOrgDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
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
			resource.TestStep{
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
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMOrgDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMOrgWithTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMOrgExists("ibm_org.testacc_org", &conf),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "name", name),
					resource.TestCheckResourceAttr("ibm_org.testacc_org", "tags.#", "1"),
				),
			},
			resource.TestStep{
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
		cfClient, err := testAccProvider.Meta().(ClientSession).MccpAPI()
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
	cfClient, err := testAccProvider.Meta().(ClientSession).MccpAPI()
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
				return fmt.Errorf("Error waiting for Organization (%s) to be destroyed: %s", rs.Primary.ID, err)
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
			name = "%s"
			auditors = ["%s"]
			managers = ["%s"]
			billing_managers = ["%s"]
			users = ["%s"]		
}`, name, ibmid1, ibmid1, ibmid1, ibmid1)
}

func testAccCheckIBMOrgUpdateWithRoles(updatedName string) string {
	return fmt.Sprintf(`
resource "ibm_org" "testacc_org" {
			name = "%s"
			auditors = ["%s"]
			managers = ["%s", "%s"]
			billing_managers = ["%s"]
			users = ["%s"]
}`, updatedName, ibmid2, ibmid2, ibmid1, ibmid2, ibmid2)
}

func testAccCheckIBMOrgWithTags(name string) string {
	return fmt.Sprintf(`
resource "ibm_org" "testacc_org" {
	name = "%s"
	tags = ["one"]
}`, name)
}

func testAccCheckIBMOrgWithUpdatedTags(name string) string {
	return fmt.Sprintf(`
resource "ibm_org" "testacc_org" {
	name = "%s"
	tags = ["one", "two"]
}`, name)
}
