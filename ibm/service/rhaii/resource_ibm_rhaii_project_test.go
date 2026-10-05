// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMRhaiiProject_basic(t *testing.T) {
	resourceName := "ibm_rhaii_project.project"
	name := fmt.Sprintf("tf-rhaii-%s", acctest.RandString(8))
	updateName := fmt.Sprintf("tf-rhaii-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMRhaiiProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRhaiiProjectConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRhaiiProjectExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "service", "instructlab"),
					resource.TestCheckResourceAttr(resourceName, "plan", "instructlab-pricing-plan"),
					resource.TestCheckResourceAttr(resourceName, "location", "us-east"),
					resource.TestCheckResourceAttr(resourceName, "state", "active"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrPair(resourceName, "project_id", resourceName, "guid"),
					resource.TestMatchResourceAttr(resourceName, "endpoint", regexp.MustCompile(`^https://us-east\.rhai\.ibm\.com/v1/projects/[0-9a-f-]{36}$`)),
				),
			},
			{
				Config: testAccCheckIBMRhaiiProjectConfig(updateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updateName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parameters"},
			},
		},
	})
}

func testAccCheckIBMRhaiiProjectConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_rhaii_project" "project" {
		name = "%s"
	}
	`, name)
}

func testAccCheckIBMRhaiiProjectDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_rhaii_project" {
			continue
		}

		instanceID := rs.Primary.ID
		instance, resp, err := rsContClient.GetResourceInstance(&rc.GetResourceInstanceOptions{
			ID: &instanceID,
		})
		if err != nil {
			if resp != nil && resp.StatusCode == 404 {
				continue
			}
			return fmt.Errorf("[ERROR] Error checking if RHAII project (%s) has been destroyed: %s with resp code: %s", rs.Primary.ID, err, resp)
		}
		if instance.State != nil && !strings.Contains(*instance.State, "removed") && !strings.Contains(*instance.State, "pending_reclamation") {
			return fmt.Errorf("RHAII project still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMRhaiiProjectExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
		if err != nil {
			return err
		}
		instanceID := rs.Primary.ID
		instance, resp, err := rsContClient.GetResourceInstance(&rc.GetResourceInstanceOptions{
			ID: &instanceID,
		})
		if err != nil {
			return fmt.Errorf("[ERROR] Error retrieving RHAII project: %s with resp code: %s", err, resp)
		}
		if instance.ResourceID == nil || *instance.ResourceID != "instructlab" {
			return fmt.Errorf("Resource instance %s is not an RHAII project", instanceID)
		}
		return nil
	}
}
