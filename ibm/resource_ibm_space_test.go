package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
)

func TestAccIBMSpace_Basic(t *testing.T) {
	var conf mccpv2.SpaceFields
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSpaceDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckIBMSpaceCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSpaceExists("ibm_space.space", &conf),
					resource.TestCheckResourceAttr("ibm_space.space", "org", cfOrganization),
					resource.TestCheckResourceAttr("ibm_space.space", "name", name),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMSpaceUpdate(updatedName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_space.space", "org", cfOrganization),
					resource.TestCheckResourceAttr("ibm_space.space", "name", updatedName),
				),
			},
		},
	})
}

func TestAccIBMSpace_Basic_Import(t *testing.T) {
	var conf mccpv2.SpaceFields
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_space.space"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSpaceDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckIBMSpaceCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSpaceExists(resourceName, &conf),
					resource.TestCheckResourceAttr(resourceName, "org", cfOrganization),
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

func TestAccIBMSpace_with_roles(t *testing.T) {
	var conf mccpv2.SpaceFields
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSpaceDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckIBMSpaceCreateWithRoles(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSpaceExists("ibm_space.space", &conf),
					resource.TestCheckResourceAttr("ibm_space.space", "org", cfOrganization),
					resource.TestCheckResourceAttr("ibm_space.space", "name", name),
					resource.TestCheckResourceAttr("ibm_space.space", "auditors.#", "1"),
					resource.TestCheckResourceAttr("ibm_space.space", "managers.#", "1"),
					resource.TestCheckResourceAttr("ibm_space.space", "developers.#", "1"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMSpaceUpdateWithRoles(updatedName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_space.space", "org", cfOrganization),
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
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSpaceDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckIBMSpaceWithTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSpaceExists("ibm_space.space", &conf),
					resource.TestCheckResourceAttr("ibm_space.space", "org", cfOrganization),
					resource.TestCheckResourceAttr("ibm_space.space", "name", name),
					resource.TestCheckResourceAttr("ibm_space.space", "tags.#", "1"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMSpaceWithUpdatedTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_space.space", "org", cfOrganization),
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

		cfClient, err := testAccProvider.Meta().(ClientSession).MccpAPI()
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
	cfClient, err := testAccProvider.Meta().(ClientSession).MccpAPI()
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
				return fmt.Errorf("Error waiting for Space (%s) to be destroyed: %s", rs.Primary.ID, err)
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
}`, cfOrganization, name)

}

func testAccCheckIBMSpaceUpdate(updatedName string) string {
	return fmt.Sprintf(`
	
resource "ibm_space" "space" {
    org = "%s"
	name = "%s"
}`, cfOrganization, updatedName)

}

func testAccCheckIBMSpaceCreateWithRoles(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_space" "space" {
    org = "%s"
	name = "%s"
	auditors = ["%s"]
	managers = ["%s"]
	developers = ["%s"]
}`, cfOrganization, name, ibmid1, ibmid1, ibmid1)

}

func testAccCheckIBMSpaceUpdateWithRoles(updatedName string) string {
	return fmt.Sprintf(`
resource "ibm_space" "space" {
    org = "%s"
	name = "%s"
	auditors = ["%s"]
	managers = ["%s", "%s"]
	developers = ["%s"]
}`, cfOrganization, updatedName, ibmid2, ibmid2, ibmid1, ibmid2)

}

func testAccCheckIBMSpaceWithTags(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_space" "space" {
    org = "%s"
	name = "%s"
	tags = ["one"]
}`, cfOrganization, name)

}
func testAccCheckIBMSpaceWithUpdatedTags(name string) string {
	return fmt.Sprintf(`
	
resource "ibm_space" "space" {
    org = "%s"
	name = "%s"
	tags = ["one", "two"]
}`, cfOrganization, name)

}
