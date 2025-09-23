// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMIAMServiceID_Basic(t *testing.T) {
	var conf string
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	updateName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServiceIDDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServiceIDBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServiceIDExists("ibm_iam_service_id.serviceID", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMIAMServiceIDUpdateWithSameName(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServiceIDExists("ibm_iam_service_id.serviceID", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "description", "ServiceID for test scenario1"),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "tags.#", "3"),
				),
			},
			{
				Config: testAccCheckIBMIAMServiceIDUpdate(updateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", updateName),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "description", "ServiceID for test scenario2"),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "tags.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMServiceID_import(t *testing.T) {
	var conf string
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_service_id.serviceID"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServiceIDDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServiceIDTag(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServiceIDExists(resourceName, conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "ServiceID for test scenario2"),
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

func testAccCheckIBMIAMServiceIDDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_service_id" {
			continue
		}

		serviceIDUUID := rs.Primary.ID
		getServiceIDOptions := iamidentityv1.GetServiceIDOptions{
			ID: &serviceIDUUID,
		}
		// Try to find the key
		_, resp, err := rsContClient.GetServiceID(&getServiceIDOptions)
		if err == nil {
			return fmt.Errorf("ServiceID still exists: %s %s", rs.Primary.ID, resp)
		} else if resp.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error waiting for serviceID (%s) to be destroyed: %s %s", rs.Primary.ID, err, resp)
		}
	}

	return nil
}

func testAccCheckIBMIAMServiceIDExists(n string, obj string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}
		serviceIDUUID := rs.Primary.ID
		getServiceIDOptions := iamidentityv1.GetServiceIDOptions{
			ID: &serviceIDUUID,
		}
		serviceID, resp, err := rsContClient.GetServiceID(&getServiceIDOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error retrieving serviceID: %s %s", err, resp)
		}

		obj = *serviceID.ID
		return nil
	}
}

func testAccCheckIBMIAMServiceIDBasic(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
			tags = ["tag1", "tag2"]
	  	}
	`, name)
}

func testAccCheckIBMIAMServiceIDUpdateWithSameName(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_service_id" "serviceID" {
			name        = "%s"
			description = "ServiceID for test scenario1"
			tags        = ["tag1", "tag2", "db"]
	  	}
	`, name)
}

func testAccCheckIBMIAMServiceIDUpdate(updateName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name              = "%s"		
			description       = "ServiceID for test scenario2"
			tags              = ["tag1"]
		}
	`, updateName)
}

func testAccCheckIBMIAMServiceIDTag(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name              = "%s"		
			description       = "ServiceID for test scenario2"
		}
	`, name)
}
