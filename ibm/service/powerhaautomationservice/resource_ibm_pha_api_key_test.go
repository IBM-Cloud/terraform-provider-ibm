// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package powerhaautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPhaAPIKeyBasic(t *testing.T) {
	var conf powerhaautomationservicev1.APIKeyResponse
	phaInstanceID := fmt.Sprintf("tf_pha_instance_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPhaAPIKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaAPIKeyConfigBasic(phaInstanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPhaAPIKeyExists("ibm_pha_api_key.pha_api_key_instance", conf),
					resource.TestCheckResourceAttr("ibm_pha_api_key.pha_api_key_instance", "pha_instance_id", phaInstanceID),
				),
			},
		},
	})
}

func TestAccIBMPhaAPIKeyAllArgs(t *testing.T) {
	var conf powerhaautomationservicev1.APIKeyResponse
	phaInstanceID := fmt.Sprintf("tf_pha_instance_id_%d", acctest.RandIntRange(10, 100))
	acceptLanguage := fmt.Sprintf("tf_accept_language_%d", acctest.RandIntRange(10, 100))
	ifNoneMatch := fmt.Sprintf("tf_if_none_match_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPhaAPIKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaAPIKeyConfig(phaInstanceID, acceptLanguage, ifNoneMatch),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPhaAPIKeyExists("ibm_pha_api_key.pha_api_key_instance", conf),
					resource.TestCheckResourceAttr("ibm_pha_api_key.pha_api_key_instance", "pha_instance_id", phaInstanceID),
					resource.TestCheckResourceAttr("ibm_pha_api_key.pha_api_key_instance", "accept_language", acceptLanguage),
					resource.TestCheckResourceAttr("ibm_pha_api_key.pha_api_key_instance", "if_none_match", ifNoneMatch),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_pha_api_key.pha_api_key_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMPhaAPIKeyConfigBasic(phaInstanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_api_key" "pha_api_key_instance" {
			pha_instance_id = "%s"
		}
	`, phaInstanceID)
}

func testAccCheckIBMPhaAPIKeyConfig(phaInstanceID string, acceptLanguage string, ifNoneMatch string) string {
	return fmt.Sprintf(`

		resource "ibm_pha_api_key" "pha_api_key_instance" {
			pha_instance_id = "%s"
			accept_language = "%s"
			if_none_match = "%s"
		}
	`, phaInstanceID, acceptLanguage, ifNoneMatch)
}

func testAccCheckIBMPhaAPIKeyExists(n string, obj powerhaautomationservicev1.APIKeyResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		powerhaAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PowerhaAutomationServiceV1()
		if err != nil {
			return err
		}

		getAPIKeyOptions := &powerhaautomationservicev1.GetAPIKeyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getAPIKeyOptions.SetPhaInstanceID(parts[0])
		getAPIKeyOptions.SetPhaInstanceID(parts[1])

		apiKeyResponse, _, err := powerhaAutomationServiceClient.GetAPIKey(getAPIKeyOptions)
		if err != nil {
			return err
		}

		obj = *apiKeyResponse
		return nil
	}
}

func testAccCheckIBMPhaAPIKeyDestroy(s *terraform.State) error {
	powerhaAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pha_api_key" {
			continue
		}

		getAPIKeyOptions := &powerhaautomationservicev1.GetAPIKeyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getAPIKeyOptions.SetPhaInstanceID(parts[0])
		getAPIKeyOptions.SetPhaInstanceID(parts[1])

		// Try to find the key
		_, response, err := powerhaAutomationServiceClient.GetAPIKey(getAPIKeyOptions)

		if err == nil {
			return fmt.Errorf("pha_api_key still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for pha_api_key (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
