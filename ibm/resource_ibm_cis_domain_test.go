// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
	"testing"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMCisDomain_basic(t *testing.T) {
	name := "ibm_cis_domain." + "cis_domain"
	testDomain := uuid.New().String() + cisDomainTest

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this test it must have already been deleted
		// correctly during the resource destroy phase of test. The destroy of resource_ibm_cis used in testAccCheckCisPoolConfigBasic
		// will fail if this resource is not correctly deleted.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisDomainConfigCisRIbasic("test_acc", testDomain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "domain", testDomain),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMCisDomain_CreateAfterManualDestroy(t *testing.T) {
	// Manual destroy of Domain resource
	//t.Parallel()
	t.Skip()
	var zoneOne, zoneTwo string
	name := "ibm_cis_domain." + "cis_domain"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisDomainConfigCisRIbasic("test", cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisDomainExists(name, &zoneOne),
					testAccCisDomainManuallyDelete(&zoneOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisDomainConfigCisRIbasic("test", cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisDomainExists(name, &zoneTwo),
					// No check for change in ID as CIS retains the same domainid across create/delete for a domain
				),
			},
		},
	})
}

func TestAccIBMCisDomain_CreateAfterManualCisRIDestroy(t *testing.T) {
	// Manual destroy of Domain resource & CIS Resource Instance
	//t.Parallel()
	t.Skip()
	var zoneOne, zoneTwo string
	name := "ibm_cis_domain." + "cis_domain"
	testDomain := uuid.New().String() + cisDomainTest
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisDomainConfigCisRIbasic("test", testDomain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisDomainExists(name, &zoneOne),
					testAccCisDomainManuallyDelete(&zoneOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisDomainConfigCisRIbasic("test", testDomain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisDomainExists(name, &zoneTwo),
					// No check for change in ID as CIS retains the same domainid across create/delete for a domain
				),
			},
		},
	})
}

func TestAccIBMCisDomain_import(t *testing.T) {
	name := "ibm_cis_domain.cis_domain"
	testDomain := uuid.New().String() + cisDomainTest
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisDomainConfigCisRIbasic("test_acc", testDomain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "status", "pending"),
					resource.TestCheckResourceAttr(name, "domain", testDomain),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
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

func testAccCisDomainManuallyDelete(tfZoneID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		tfZone := *tfZoneID

		cisClient, err := testAccProvider.Meta().(ClientSession).CisZonesV1ClientSession()
		if err != nil {
			return err
		}

		zoneID, crn, err := convertTftoCisTwoVar(tfZone)
		if err != nil {
			return err
		}
		cisClient.Crn = core.StringPtr(crn)
		delOpt := cisClient.NewDeleteZoneOptions(zoneID)
		_, resp, err := cisClient.DeleteZone(delOpt)
		if err != nil {
			return fmt.Errorf("[ERR] Error deleting zone %v", resp)
		}
		return nil
	}
}

func testAccCheckCisDomainDestroy(s *terraform.State) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisZonesV1ClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_domain" {
			continue
		}
		log.Println("check domain destroy : ", rs.Primary.ID)
		zoneID, crn, err := convertTftoCisTwoVar(rs.Primary.ID)
		if err != nil {
			return err
		}
		cisClient.Crn = core.StringPtr(crn)
		opt := cisClient.NewGetZoneOptions(zoneID)
		_, _, err = cisClient.GetZone(opt)
		if err == nil {
			return fmt.Errorf("Domain still exists when destroying")
		}
	}

	return nil
}

func testAccCheckCisDomainExists(n string, tfZoneID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Domain ID is set")
		}

		cisClient, err := testAccProvider.Meta().(ClientSession).CisZonesV1ClientSession()
		if err != nil {
			return err
		}
		zoneID, crn, err := convertTftoCisTwoVar(rs.Primary.ID)
		if err != nil {
			return err
		}
		cisClient.Crn = core.StringPtr(crn)
		opt := cisClient.NewGetZoneOptions(zoneID)
		foundZone, resp, err := cisClient.GetZone(opt)
		if err != nil {
			return fmt.Errorf("Domain does not exists: %v", resp)
		}
		*tfZoneID = convertCisToTfTwoVar(*foundZone.Result.ID, crn)
		return nil
	}
}

// func testAccCheckCisDomainConfigCisDS_basic(resourceName string, domain string) string {
// 	// Cis instance data source
// 	return testAccCheckCisInstanceDataSourceConfig_basic(cisResourceGroup, cisInstance) + fmt.Sprintf(`
// 	resource "ibm_cis_domain" "%[1]s" {
// 		cis_id = data.ibm_cis.testacc_ds_cis.id
// 		domain = "%[2]s"
// 	  }
// 	`, resourceName, domain)
// }

func testAccCheckCisDomainConfigCisRIbasic(resourceName string, domain string) string {
	// Cis dynamically created resource instance
	return testAccCheckIBMCisDataSourceConfig(cisInstance) + fmt.Sprintf(`
	resource "ibm_cis_domain" "cis_domain" {
		cis_id = data.ibm_cis.cis.id
		domain = "%[1]s"
	  }
	`, domain)
}

// func testAccCheckCisDomainDataSourceConfig_basic(resourceName string, domain string) string {
// 	return testAccCheckCisInstanceDataSourceConfig_basic(cisResourceGroup, cisInstance) + fmt.Sprintf(`
// 	data "ibm_cis_domain" "%[1]s" {
// 		cis_id = data.ibm_cis.testacc_ds_cis.id
// 		domain = "%[2]s"
// 	  }
// 	`, resourceName, domain)
// }

// func testAccCheckCisInstanceDataSourceConfig_basic(cisResourceGroup string, cisInstance string) string {
// 	// defaultResourceGroup from env vars
// 	//cisInstance from env vars
// 	return fmt.Sprintf(`
// 	data "ibm_resource_group" "test_acc" {
// 		name = "%[1]s"
// 	  }

// 	  data "ibm_cis" "testacc_ds_cis" {
// 		resource_group_id = data.ibm_resource_group.test_acc.id
// 		name              = "%[2]s"
// 	  }
// 	`, cisResourceGroup, cisInstance)

// }
