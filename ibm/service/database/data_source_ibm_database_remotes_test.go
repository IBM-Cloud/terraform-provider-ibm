// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMDatabaseRemotesDataSourceBasic(t *testing.T) {
	var leader string

	testName := fmt.Sprintf("tf-Pgress-%s", acctest.RandString(16))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMDatabaseRemotesDataSourceConfigBasic(testName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExampleLeaderExists(testName, &leader),
					resource.TestCheckResourceAttrSet("data.ibm_database_remotes.database_remotes", "deployment_id"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseDataSourceConfig4(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
	}
	data "ibm_database" "%[1]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = ibm_database.db.name
	}
	resource "ibm_database" "db" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[1]s"
		service           = "databases-for-postgresql"
		plan              = "standard"
		location          = "%[2]s"
		tags              = ["one:two"]
	}
				`, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseRemotesDataSourceConfigBasic(name string) string {
	return testAccCheckIBMDatabaseDataSourceConfig4(name) + `
		data "ibm_database_remotes" "database_remotes" {
			deployment_id = ibm_database.db.id
		}
	`
}

func testAccCheckExampleLeaderExists(resourceName string, widget *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Widget ID is not set")
		}

		// retrieve the client from the test provider
		rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
		if err != nil {
			return err
		}

		instanceID := rs.Primary.ID

		rsInst := rc.GetResourceInstanceOptions{
			ID: &instanceID,
		}

		instance, response, err := rsContClient.GetResourceInstance(&rsInst)
		if err != nil {
			if strings.Contains(err.Error(), "Object not found") ||
				strings.Contains(err.Error(), "status code: 404") {
				*widget = ""
				return nil
			}
			return fmt.Errorf("[ERROR] Error retrieving resource instance: %s %s", err, response)
		}

		if strings.Contains(*instance.State, "removed") {
			*widget = ""
			return nil
		}

		*widget = instanceID
		return nil
	}
}
