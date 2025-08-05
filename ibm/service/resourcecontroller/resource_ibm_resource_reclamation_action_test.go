// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcecontroller_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMResourceReclamationAction_basic(t *testing.T) {
	// Use random reclamation IDs or create preconditioned reclamation resources for full test
	reclamationID := "put-valid-reclamation-id-here" // Adjust or create programmatically

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMResourceReclamationActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceReclamationActionConfig(reclamationID, "reclaim", "Terraform test reclaim action"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceReclamationActionExists("ibm_resource_reclamation_action.test"),
					resource.TestCheckResourceAttr("ibm_resource_reclamation_action.test", "id", reclamationID),
					resource.TestCheckResourceAttr("ibm_resource_reclamation_action.test", "state", "pending_reclamation"), // Example expected state; adjust
				),
			},
			{
				Config: testAccCheckIBMResourceReclamationActionConfig(reclamationID, "restore", "Terraform test restore action"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceReclamationActionExists("ibm_resource_reclamation_action.test"),
					resource.TestCheckResourceAttr("ibm_resource_reclamation_action.test", "id", reclamationID),
					resource.TestCheckResourceAttrSet("ibm_resource_reclamation_action.test", "state"),
				),
			},
		},
	})
}

func testAccCheckIBMResourceReclamationActionExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsConClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
		if err != nil {
			return err
		}
		reclamationID := rs.Primary.ID
		opts := rc.ListReclamationsOptions{}

		reclamationsList, resp, err := rsConClient.ListReclamations(&opts)
		if err != nil {
			return fmt.Errorf("Error listing reclamations: %s Response: %s", err, resp)
		}

		found := false
		for _, rec := range reclamationsList.Resources {
			if rec.ID != nil && *rec.ID == reclamationID {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("Reclamation not found: %s", reclamationID)
		}

		return nil
	}
}

func testAccCheckIBMResourceReclamationActionDestroy(s *terraform.State) error {
	rsConClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}

	// Verify all reclamations in state are removed or not found
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_resource_reclamation_action" {
			continue
		}

		reclamationID := rs.Primary.ID
		opts := rc.ListReclamationsOptions{}

		reclamationsList, resp, err := rsConClient.ListReclamations(&opts)
		if err != nil {
			if strings.Contains(err.Error(), "Not found") {
				continue
			}
			return fmt.Errorf("Error listing reclamations during destroy check: %s Response: %s", err, resp)
		}

		for _, rec := range reclamationsList.Resources {
			if rec.ID != nil && *rec.ID == reclamationID {
				return fmt.Errorf("Reclamation still exists after destroy: %s", reclamationID)
			}
		}
	}

	return nil
}

func testAccCheckIBMResourceReclamationActionConfig(id string, action string, comment string) string {
	return fmt.Sprintf(`
resource "ibm_resource_reclamation_action" "test" {
  id         = "%s"
  action     = "%s"
  comment    = "%s"
}
`, id, action, comment)
}
