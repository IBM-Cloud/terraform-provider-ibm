// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMCisCertificateOrder_Basic(t *testing.T) {
	var monitor string
	name := "ibm_cis_certificate_order.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this point it must have already been deleted from CIS.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCertificateOrderConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisCertificateOrderExists(name, &monitor),
					resource.TestCheckResourceAttr(name, "hosts.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMCisCertificateOrder_import(t *testing.T) {
	name := "ibm_cis_certificate_order.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckCisCertificateOrderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCertificateOrderConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "hosts.#", "1"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMCisCertificateOrder_CreateAfterManualDestroy(t *testing.T) {
	var certOne, certTwo string
	name := "ibm_cis_certificate_order.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckCisCertificateOrderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCertificateOrderConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisCertificateOrderExists(name, &certOne),
					testAccCisCertificateOrderManuallyDelete(&certOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisCertificateOrderConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisCertificateOrderExists(name, &certTwo),
					func(state *terraform.State) error {
						if certOne == certTwo {
							return fmt.Errorf("certificate id is unchanged even after we thought we deleted it ( %s )",
								certTwo)
						}
						return nil
					},
				),
			},
		},
	})
}

func testAccCisCertificateOrderManuallyDelete(tfCertID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisSSLClientSession()
		if err != nil {
			return err
		}
		tfCert := *tfCertID
		certID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(tfCert)
		cisClient.Crn = core.StringPtr(crn)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewDeleteCertificateOptions(certID)
		_, err = cisClient.DeleteCertificate(opt)
		if err != nil {
			return fmt.Errorf("[ERROR] Error deleting certificate: %s", err)
		}
		return nil
	}
}

func testAccCheckCisCertificateOrderDestroy(s *terraform.State) error {
	cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisSSLClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_certificate_order" {
			continue
		}
		certID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(crn)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewGetCustomCertificateOptions(certID)
		_, _, err = cisClient.GetCustomCertificate(opt)
		if err == nil {
			return fmt.Errorf("Certificate still exists")
		}
	}

	return nil
}

func testAccCheckCisCertificateOrderExists(n string, tfCertID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[ERROR] Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Certificate ID is set")
		}

		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisSSLClientSession()
		if err != nil {
			return err
		}
		certID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(crn)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewGetCustomCertificateOptions(certID)
		result, _, err := cisClient.GetCustomCertificate(opt)
		if err != nil {
			return fmt.Errorf("Certificate still exists")
		}
		*tfCertID = flex.ConvertCisToTfThreeVar(*result.Result.ID, zoneID, crn)
		return nil
	}
}

func testAccCheckCisCertificateOrderConfigBasic() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_certificate_order" "test" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id
		hosts     = ["%[1]s"]
	  }
	`, acc.CisDomainStatic)
}
