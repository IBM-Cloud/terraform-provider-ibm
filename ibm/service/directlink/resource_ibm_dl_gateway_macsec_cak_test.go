// Copyright IBM Corp. 2017, 2021, 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM/networking-go-sdk/directlinkv1"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMDLGatewayMacsecCak_basic(t *testing.T) {
	var cakID string
	cakName := fmt.Sprintf("DD%d", acctest.RandIntRange(00, 99))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDLGatewayMacsecCakDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMDLGatewayMacsecCakConfig(cakName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayMacsecCakExists("ibm_dl_gateway_macsec_cak.test_cak", &cakID),
					resource.TestCheckResourceAttr("ibm_dl_gateway_macsec_cak.test_cak", "session", "fallback"),
					resource.TestCheckResourceAttr("ibm_dl_gateway_macsec_cak.test_cak", "name", cakName),
				),
			},
		},
	})
}

func testAccIBMDLGatewayMacsecCakConfig(cakName string) string {
	return fmt.Sprintf(`
	resource "ibm_dl_gateway_macsec_cak" "test_cak" {
		gateway = "9c95f464-1ba9-471e-85b4-d2bf188cb273"
		key {
			crn = "crn:v1:staging:public:hs-crypto:us-south:a/3f455c4c574447adbc14bda52f80e62f:b2044455-b89e-4c57-96ae-3f17c092dd31:key:6f79b964-229c-45ab-b1d9-47e111cd03f6"
		}
		name = "%s"
		session = "fallback"
	}
	`, cakName)
}

func testAccCheckIBMDLGatewayMacsecCakDestroy(s *terraform.State) error {
	directLink, err := directlinkClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dl_gateway_macsec_cak" {
			continue
		}

		gatewayId := rs.Primary.ID
		cakId := rs.Primary.Attributes["cak_id"]

		delOptions := &directlinkv1.DeleteGatewayMacsecCakOptions{
			ID:    &gatewayId,
			CakID: &cakId,
		}

		_, err := directLink.DeleteGatewayMacsecCak(delOptions)

		if err == nil {
			return fmt.Errorf("Macsec CAK still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMDLGatewayMacsecCakExists(n string, cakID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		directLink, err := directlinkClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		gatewayId := rs.Primary.ID
		id := rs.Primary.Attributes["cak_id"]
		*cakID = id

		opts := &directlinkv1.GetGatewayMacsecCakOptions{
			ID:    &gatewayId,
			CakID: &id,
		}
		_, response, err := directLink.GetGatewayMacsecCak(opts)
		if err != nil {
			return fmt.Errorf("Error reading Macsec CAK: %v\nResponse: %v", err, response)
		}
		return nil
	}
}
