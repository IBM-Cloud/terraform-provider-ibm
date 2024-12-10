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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMCisDNSRecord_Basic(t *testing.T) {
	//t.Parallel()
	var record string
	testName := "tf-acctest-basic"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", testName)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCis(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisDNSRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigCisDSBasic(testName, acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", testName+"."+acc.CisDomainStatic),
					resource.TestCheckResourceAttr(
						resourceName, "content", "192.168.0.10"),
					resource.TestCheckResourceAttr(
						resourceName, "data.%", "0"),
				),
			},
		},
	})
}

func TestAccIBMCisDNSRecord_PTR(t *testing.T) {
	//t.Parallel()
	var record string
	testName := "tf-acctest-ptr"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", testName)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCis(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisDNSRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigPTR(testName, acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", "192.168.0.10."+acc.CisDomainStatic),
					resource.TestCheckResourceAttr(
						resourceName, "content", testName+"."+acc.CisDomainStatic),
					resource.TestCheckResourceAttr(
						resourceName, "data.%", "0"),
				),
			},
		},
	})
}

func TestAccIBMCisDNSRecord_import(t *testing.T) {
	name := "ibm_cis_dns_record.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigCisDSBasic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
					resource.TestCheckResourceAttr(name, "name", "test."+acc.CisDomainStatic),
					resource.TestCheckResourceAttr(name, "content", "192.168.0.10"),
					resource.TestCheckResourceAttr(name, "data.%", "0"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes"},
			},
		},
	})
}

func TestAccIBMCisDNSRecord_CaseInsensitive(t *testing.T) {
	//t.Parallel()
	var record string
	testName := "tf-acctest-case-insensitive"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", testName)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisDNSRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigCaseSensitive(testName, acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", testName+"."+acc.CisDomainStatic),
				),
			},
			{
				Config: testAccCheckIBMCisDNSRecordConfigCaseSensitive("tf-acctest-CASE-INSENSITIVE", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists("ibm_cis_dns_record.tf-acctest-CASE-INSENSITIVE", &record),
					resource.TestCheckResourceAttr(
						"ibm_cis_dns_record.tf-acctest-CASE-INSENSITIVE", "name", "tf-acctest-case-insensitive."+acc.CisDomainStatic),
				),
			},
		},
	})
}

func TestAccIBMCisDNSRecord_Apex(t *testing.T) {
	//t.Parallel()
	var record string
	testName := "test"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", testName)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisDNSRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigApex(testName, acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						// @ is replaced by domain name by CIS
						resourceName, "name", acc.CisDomainStatic),
					resource.TestCheckResourceAttr(
						resourceName, "content", "192.168.0.10"),
				),
			},
		},
	})
}

func TestAccIBMCisDNSRecord_CreateAfterManualDestroy(t *testing.T) {
	// t.Skip()
	testName := "test_acc"
	var afterCreate, afterRecreate string
	name := "ibm_cis_dns_record.test_acc"

	afterCreate = "hello"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCis(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCisDNSRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigCisDSBasic(testName, acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists(name, &afterCreate),
					testAccIBMCisManuallyDeleteDNSRecord(&afterCreate),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckIBMCisDNSRecordConfigCisDSBasic(testName, acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists(name, &afterRecreate),
					testAccCheckIBMCisDNSRecordRecreated(&afterCreate, &afterRecreate),
				),
			},
		},
	})
}

func testAccIBMCisManuallyDeleteDNSRecord(tfRecordID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisDNSRecordClientSession()
		if err != nil {
			return err
		}
		tfRecord := *tfRecordID
		recordID, zoneID, cisID, _ := flex.ConvertTfToCisThreeVar(tfRecord)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		delOpt := cisClient.NewDeleteDnsRecordOptions(recordID)
		_, _, err = cisClient.DeleteDnsRecord(delOpt)
		if err != nil {
			return fmt.Errorf("[ERROR] Error deleting IBMCISDNS Record: %s", err)
		}
		return nil
	}
}

func testAccCheckIBMCisDNSRecordRecreated(beforeID, afterID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *beforeID == *afterID {
			return fmt.Errorf("Expected change of Record Ids, but both were %v", beforeID)
		}
		return nil
	}
}

func testAccCheckIBMCisDNSRecordDestroy(s *terraform.State) error {
	cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisDNSRecordClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_record" {
			continue
		}

		recordID, zoneID, cisID, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		delOpt := cisClient.NewDeleteDnsRecordOptions(recordID)
		_, _, err = cisClient.DeleteDnsRecord(delOpt)
		if err != nil {
			return fmt.Errorf("Record still exists")
		}
	}

	return nil
}

func testAccCheckIBMCisDNSRecordExists(n string, tfRecordID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[ERROR] Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Record ID is set")
		}

		// tfRecord := *tfRecordID
		cisClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CisDNSRecordClientSession()
		if err != nil {
			return err
		}
		recordID, zoneID, cisID, _ := flex.ConvertTfToCisThreeVar(rs.Primary.ID)
		cisClient.Crn = core.StringPtr(cisID)
		cisClient.ZoneIdentifier = core.StringPtr(zoneID)
		opt := cisClient.NewGetDnsRecordOptions(recordID)
		foundRecord, _, err := cisClient.GetDnsRecord(opt)
		if err != nil {
			return err
		}

		if *foundRecord.Result.ID != recordID {
			return fmt.Errorf("Record not found")
		}

		tfRecord := flex.ConvertCisToTfThreeVar(*foundRecord.Result.ID, zoneID, cisID)
		*tfRecordID = tfRecord
		return nil
	}
}

func testAccCheckIBMCisDNSRecordConfigCisDSBasic(resourceID string, cisDomain string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_dns_record" "%[1]s" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id

		name    = "%[1]s"
		content = "192.168.0.10"
		type    = "A"
	  }
	  `, resourceID)
}

func testAccCheckIBMCisDNSRecordConfigPTR(resourceID string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_dns_record" "%[1]s" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		name    = "192.168.0.10"
		content = "%[1]s.%[2]s"
		type    = "PTR"
	  }
	`, resourceID, acc.CisDomainStatic)
}

func testAccCheckIBMCisDNSRecordConfigCaseSensitive(resourceID string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_dns_record" "%[1]s" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		name    = "%[1]s"
		content = "192.168.0.10"
		type    = "A"
	  }
`, resourceID)
}

func testAccCheckIBMCisDNSRecordConfigApex(resourceID string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_dns_record" "%[1]s" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		name      = "@"
		content   = "192.168.0.10"
		type      = "A"
	  }
`, resourceID)
}

func testAccCheckIBMCisDNSRecordConfigLOC(resourceID string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_dns_record" "%[1]s" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		name      = "%[1]s"
		data = {
		  "lat_degrees"    = "37"
		  "lat_minutes"    = "46"
		  "lat_seconds"    = "46"
		  "lat_direction"  = "N"
		  "long_degrees"   = "122"
		  "long_minutes"   = "23"
		  "long_seconds"   = "35"
		  "long_direction" = "W"
		  "altitude"       = 0
		  "size"           = 100
		  "precision_horz" = 0
		  "precision_vert" = 0
		}
		type = "LOC"
	  }
`, resourceID)
}

func testAccCheckIBMCisDNSRecordConfigSRV(resourceID string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_dns_record" "%[1]s" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		name      = "%[1]s"
		data = {
		  "priority" = 5
		  "weight"   = 0
		  "port"     = 5222
		  "target"   = "talk.l.google.com"
		  "service"  = "_xmpp-client"
		  "proto"    = "_tcp"
		}
		type = "SRV"
	  }
`, resourceID)
}

func testAccCheckIBMCisDNSRecordConfigProxied(resourceID string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_dns_record" "%[1]s" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id

		name    = "%[1]s"
		content = "%[1]s"
		type    = "CNAME"
		proxied = true
	  }
`, resourceID)
}
