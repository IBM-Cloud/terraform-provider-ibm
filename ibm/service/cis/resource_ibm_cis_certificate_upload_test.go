// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMCisCertificateUpload_Basic(t *testing.T) {
	t.Skip()
	var cert string

	name := "ibm_cis_certificate_upload." + "test"
	certMgrInstanceName := fmt.Sprintf("testacc-cert-manager-%s",
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	domainName := fmt.Sprintf("%s.%s",
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum), acc.CisDomainStatic)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCis(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckCisCertificateUploadDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCertificateUploadConfigBasic(certMgrInstanceName, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisCertificateUploadExists(name, &cert),
					resource.TestCheckResourceAttr(name, "bundle_method", "ubiquitous"),
					resource.TestCheckResourceAttr(name, "priority", "20"),
				),
			},
			{
				Config: testAccCheckCisCertificateUploadConfigUpdate(certMgrInstanceName, domainName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisCertificateUploadExists(name, &cert),
					resource.TestCheckResourceAttr(name, "bundle_method", "ubiquitous"),
					resource.TestCheckResourceAttr(name, "priority", "1"),
				),
			},
		},
	})
}

func TestAccIBMCisCertificateUpload_import(t *testing.T) {
	t.Skip()
	name := "ibm_cis_certificate_upload.test"
	certMgrInstanceName := fmt.Sprintf("testacc-cert-manager-%s",
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	domainName := fmt.Sprintf("%s.%s",
		acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum), acc.CisDomainStatic)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCertificateUploadConfigBasic(certMgrInstanceName, domainName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "priority", "20"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"certificate",
					"private_key"},
			},
		},
	})
}

func testAccCheckCisCertificateUploadDestroy(s *terraform.State) error {
	cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisSSLClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_certificate_upload" {
			continue
		}
		certID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		if err != nil {
			return err
		}
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

func testAccCheckCisCertificateUploadExists(n string, tfCertID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[ERROR] Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No certificate ID is set")
		}
		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisSSLClientSession()
		if err != nil {
			return err
		}
		certID, zoneID, crn, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		if err != nil {
			return err
		}
		cisClient.Crn = core.StringPtr(crn)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)

		opt := cisClient.NewGetCustomCertificateOptions(certID)

		result, _, err := cisClient.GetCustomCertificate(opt)
		if err != nil {
			return fmt.Errorf("Certificate exists")
		}
		*tfCertID = flex.ConvertCisToTfThreeVar(*result.Result.ID, zoneID, crn)
		return nil
	}
}

func testAccCheckCisCertificateUploadConfigBasic(certMgrInstanceName, domainName string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_resource_instance" "cm" {
		name     = "%[1]s"
		location = "%[2]s"
		service  = "cloudcerts"
		plan     = "free"
		resource_group_id = data.ibm_resource_group.test_acc.id
	}
	resource "ibm_certificate_manager_order" "cert" {
		certificate_manager_instance_id = ibm_resource_instance.cm.id
		name                            = "cis-test-certificate-upload"
		domains                         = ["%[3]s"]
		dns_provider_instance_crn       = data.ibm_cis.cis.id
	}
	data "ibm_certificate_manager_certificate" "data_cert" {
		certificate_manager_instance_id = ibm_certificate_manager_order.cert.certificate_manager_instance_id
		name                            = ibm_certificate_manager_order.cert.name
	}
	resource "ibm_cis_certificate_upload" "test" {
		cis_id        = data.ibm_cis.cis.id
		domain_id     = data.ibm_cis_domain.cis_domain.id
		certificate   = data.ibm_certificate_manager_certificate.data_cert.certificate_details.0.data.content
		private_key   = data.ibm_certificate_manager_certificate.data_cert.certificate_details.0.data.priv_key
		bundle_method = "ubiquitous"
		priority      = 20
	  }
	`, certMgrInstanceName, acc.RegionName, domainName)
}

func testAccCheckCisCertificateUploadConfigUpdate(certMgrInstanceName, domainName string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_resource_instance" "cm" {
		name     = "%[1]s"
		location = "%[2]s"
		service  = "cloudcerts"
		plan     = "free"
		resource_group_id = data.ibm_resource_group.test_acc.id
	}
	resource "ibm_certificate_manager_order" "cert" {
		certificate_manager_instance_id = ibm_resource_instance.cm.id
		name                            = "cis-test-certificate-upload"
		domains                         = ["%[3]s"]
		dns_provider_instance_crn       = data.ibm_cis.cis.id
	}
	data "ibm_certificate_manager_certificate" "data_cert" {
		certificate_manager_instance_id = ibm_certificate_manager_order.cert.certificate_manager_instance_id
		name                            = ibm_certificate_manager_order.cert.name
	}
	resource "ibm_cis_certificate_upload" "test" {
		cis_id        = data.ibm_cis.cis.id
		domain_id     = data.ibm_cis_domain.cis_domain.id
		certificate   = data.ibm_certificate_manager_certificate.data_cert.certificate_details.0.data.content
		private_key   = data.ibm_certificate_manager_certificate.data_cert.certificate_details.0.data.priv_key
		bundle_method = "ubiquitous"
		priority      = 1
	  }
	`, certMgrInstanceName, acc.RegionName, domainName)
}
