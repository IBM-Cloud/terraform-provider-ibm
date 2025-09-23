// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package hpcs_test

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"testing"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/resourcecontroller"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMHPCSInstanceBasic(t *testing.T) {
	var hpcsInstance string
	testName := fmt.Sprintf("tf-hpcs-%d", acctest.RandIntRange(10, 100))
	name := "ibm_hpcs.hpcs"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckHPCS(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMHPCSInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMHPCSInstanceBasic(testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMHPCSInstanceExists(name, hpcsInstance),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", "us-south"),
					resource.TestCheckResourceAttr(name, "admins.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMHPCSInstanceAdminUpdate(testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMHPCSInstanceExists(name, hpcsInstance),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", "us-south"),
					resource.TestCheckResourceAttr(name, "admins.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMHPCSInstanceAdminDelete(testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMHPCSInstanceExists(name, hpcsInstance),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", "us-south"),
					resource.TestCheckResourceAttr(name, "admins.#", "1"),
				),
			},
			{
				Config:      testAccCheckIBMHPCSInstanceUnitsUpdate(testName),
				ExpectError: regexp.MustCompile(`'units' attribute is immutable and can't be changed`),
			},
		},
	})
}

func testAccCheckIBMHPCSInstanceDestroy(s *terraform.State) error {
	rsConClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_hpcs" {
			continue
		}

		instanceID := rs.Primary.ID
		rsInst := rc.GetResourceInstanceOptions{
			ID: &instanceID,
		}
		instance, response, err := rsConClient.GetResourceInstance(&rsInst)
		if err == nil {
			if instance != nil && (strings.Contains(*instance.State, "removed") || strings.Contains(*instance.State, resourcecontroller.RsInstanceReclamation)) {
				log.Printf("[WARN]Returning nil because it's in removed or pending_reclamation state")
				return nil
			}
			return fmt.Errorf("Instance still exists: %s", rs.Primary.ID)
		} else if strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error checking if instance (%s) has been destroyed: %s %s", rs.Primary.ID, err, response)
		}
	}
	return nil
}

func testAccCheckIBMHPCSInstanceExists(n string, tfHPCSID string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsConClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
		if err != nil {
			return err
		}
		instanceID := rs.Primary.ID

		rsInst := rc.GetResourceInstanceOptions{
			ID: &instanceID,
		}
		instance, response, err := rsConClient.GetResourceInstance(&rsInst)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				tfHPCSID = ""
				return nil
			}
			return fmt.Errorf("[ERROR] Error retrieving resource instance: %s %s", err, response)
		}
		if strings.Contains(*instance.State, "removed") {
			tfHPCSID = ""
			return nil
		}

		tfHPCSID = instanceID
		return nil
	}
}

func testAccCheckIBMHPCSInstanceBasic(name string) string {
	return fmt.Sprintf(`
	resource ibm_hpcs hpcs {
		location             = "us-south"
		name                 = "%s"
		plan                 = "standard"
		units                = 2
		signature_threshold  = 1
		revocation_threshold = 1
		admins {
			name  = "ad1"
			key   = "%s"
			token = "%s"
		}
	}
	`, name, acc.HpcsAdmin1, acc.HpcsToken1)
}
func testAccCheckIBMHPCSInstanceAdminUpdate(name string) string {
	return fmt.Sprintf(`
	resource ibm_hpcs hpcs {
		location             = "us-south"
		name                 = "%s"
		plan                 = "standard"
		units                = 2
		signature_threshold  = 1
		revocation_threshold = 1
		admins {
			name  = "ad1"
			key   = "%s"
			token = "%s"
		}
		admins {
			name  = "ad2"
			key   = "%s"
			token = "%s"
		}
	}
	`, name, acc.HpcsAdmin1, acc.HpcsToken1, acc.HpcsAdmin2, acc.HpcsToken2)
}
func testAccCheckIBMHPCSInstanceAdminDelete(name string) string {
	return fmt.Sprintf(`
	resource ibm_hpcs hpcs {
		location             = "us-south"
		name                 = "%s"
		plan                 = "standard"
		units                = 2
		signature_threshold  = 1
		revocation_threshold = 1
		admins {
			name  = "ad1"
			key   = "%s"
			token = "%s"
		}
	}
	`, name, acc.HpcsAdmin1, acc.HpcsToken1)
}
func testAccCheckIBMHPCSInstanceUnitsUpdate(name string) string {
	return fmt.Sprintf(`
	resource ibm_hpcs hpcs {
		location             = "us-south"
		name                 = "%s"
		plan                 = "standard"
		units                = 3
		signature_threshold  = 1
		revocation_threshold = 1
		admins {
			name  = "ad1"
			key   = "%s"
			token = "%s"
		}
	}
	`, name, acc.HpcsAdmin1, acc.HpcsToken1)
}
