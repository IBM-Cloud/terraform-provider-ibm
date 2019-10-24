package ibm

import (
	"fmt"
	"testing"

	"github.com/IBM-Cloud/bluemix-go/models"

	"strings"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMIAMAccessGroup_Basic(t *testing.T) {
	var conf models.AccessGroup
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updateName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroup_basic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupExists("ibm_iam_access_group.accgroup", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "tags.#", "2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroup_updateWithSameName(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupExists("ibm_iam_access_group.accgroup", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "description", "AccessGroup for test scenario1"),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgroup", "tags.#", "3"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroup_update(updateName),
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
	var conf models.AccessGroup
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_iam_access_group.accgroup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroup_tag(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupExists(resourceName, conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "AccessGroup for test scenario2"),
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

func testAccCheckIBMIAMAccessGroupDestroy(s *terraform.State) error {
	accClient, err := testAccProvider.Meta().(ClientSession).IAMUUMAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_access_group" {
			continue
		}

		agID := rs.Primary.ID

		// Try to find the key
		_, _, err := accClient.AccessGroup().Get(agID)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for access group (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMAccessGroupExists(n string, obj models.AccessGroup) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		accClient, err := testAccProvider.Meta().(ClientSession).IAMUUMAPI()
		if err != nil {
			return err
		}
		agID := rs.Primary.ID

		accgroup, _, err := accClient.AccessGroup().Get(agID)

		if err != nil {
			return err
		}

		obj = *accgroup
		return nil
	}
}

func testAccCheckIBMIAMAccessGroup_basic(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_access_group" "accgroup" {
			name              = "%s"		
			tags              = ["tag1","tag2"]
		}
	`, name)
}

func testAccCheckIBMIAMAccessGroup_updateWithSameName(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_access_group" "accgroup" {
			name              = "%s"
			description       = "AccessGroup for test scenario1"
			tags              = ["tag1","tag2","db"]
		}
	`, name)
}

func testAccCheckIBMIAMAccessGroup_update(updateName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgroup" {
			name              = "%s"		
			description       = "AccessGroup for test scenario2"
			tags              = ["tag1"]
		}
	`, updateName)
}

func testAccCheckIBMIAMAccessGroup_tag(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgroup" {
			name              = "%s"		
			description       = "AccessGroup for test scenario2"
		}
	`, name)
}
