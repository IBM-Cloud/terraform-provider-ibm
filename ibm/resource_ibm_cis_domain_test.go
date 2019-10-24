package ibm

import (
	"fmt"
	"testing"

	// v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMCisDomain_basic(t *testing.T) {
	//rnd := acctest.RandString(10)
	name := "ibm_cis_domain." + "test_acc"
	testDomain := cisDomainTest

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this test it must have already been deleted
		// correctly during the resource destroy phase of test. The destroy of resource_ibm_cis used in testAccCheckCisPoolConfigBasic
		// will fail if this resource is not correctly deleted.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisDomainConfigCisRI_basic("test_acc", testDomain),
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
	var zoneOne, zoneTwo string
	name := "ibm_cis_domain." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisDomainConfigCisRI_basic("test", cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisDomainExists(name, &zoneOne),
					testAccCisDomainManuallyDelete(&zoneOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisDomainConfigCisRI_basic("test", cisDomainTest),
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
	var zoneOne, zoneTwo string
	name := "ibm_cis_domain." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCisDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisDomainConfigCisRI_basic("test", cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisDomainExists(name, &zoneOne),
					testAccCisDomainManuallyDelete(&zoneOne),
					testAccCisInstanceManuallyDelete(&zoneOne),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCisDomainConfigCisRI_basic("test", cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCisDomainExists(name, &zoneTwo),
					// No check for change in ID as CIS retains the same domainid across create/delete for a domain
				),
			},
		},
	})
}

func TestAccIBMCisDomain_import(t *testing.T) {
	name := "ibm_cis_domain.test_acc"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckCisDomainConfigCisRI_basic("test_acc", cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "domain", cisDomainTest),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
				),
			},
			resource.TestStep{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes"},
			},
		},
	})
}

func testAccCisDomainManuallyDelete(tfZoneId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		if err != nil {
			return err
		}
		tfZone := *tfZoneId
		zoneId, cisId, _ := convertTftoCisTwoVar(tfZone)
		err = cisClient.Zones().DeleteZone(cisId, zoneId)
		if err != nil {
			return fmt.Errorf("Error deleting IBMCISZone Record: %s", err)
		}
		return nil
	}
}

func testAccCheckCisDomainDestroy(s *terraform.State) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_domain" {
			continue
		}
		zoneId, cisId, _ := convertTftoCisTwoVar(rs.Primary.ID)
		_, err = cisClient.Zones().GetZone(cisId, zoneId)
		if err == nil {
			return fmt.Errorf("Domain still exists")
		}
	}

	return nil
}

func testAccCheckCisDomainExists(n string, tfZoneId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Domain ID is set")
		}

		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		zoneId, cisId, _ := convertTftoCisTwoVar(rs.Primary.ID)
		foundZone, err := cisClient.Zones().GetZone(cisId, zoneId)
		if err != nil {
			return err
		}

		*tfZoneId = convertCisToTfTwoVar(foundZone.Id, cisId)
		return nil
	}
}

func testAccCheckCisDomainConfigCisDS_basic(resourceName string, domain string) string {
	// Cis instance data source
	return testAccCheckCisInstanceDataSourceConfig_basic(cisResourceGroup, cisInstance) + fmt.Sprintf(`
				resource "ibm_cis_domain" "%[1]s" {
					cis_id = "${data.ibm_cis.testacc_ds_cis.id}"
                    domain = "%[2]s"
				}`, resourceName, domain)
}

func testAccCheckCisDomainConfigCisRI_basic(resourceName string, domain string) string {
	// Cis dynamically created resource instance
	return testAccCheckIBMCisInstance_basic(cisResourceGroup, "testacc_ds_cis") + fmt.Sprintf(`
				resource "ibm_cis_domain" "%[1]s" {
					cis_id = "${ibm_cis.testacc_ds_cis.id}"
                    domain = "%[2]s"
				}`, resourceName, domain)
}

func testAccCheckCisDomainDataSourceConfig_basic(resourceName string, domain string) string {
	return testAccCheckCisInstanceDataSourceConfig_basic(cisResourceGroup, cisInstance) + fmt.Sprintf(`
				data "ibm_cis_domain" "%[1]s" {
					cis_id = "${data.ibm_cis.testacc_ds_cis.id}"
                    domain = "%[2]s"
				}`, resourceName, domain)
}

func testAccCheckCisInstanceDataSourceConfig_basic(cisResourceGroup string, cisInstance string) string {
	// defaultResourceGroup from env vars
	//cisInstance from env vars
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  name = "%[1]s"
}

data "ibm_cis" "testacc_ds_cis" {
  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
  name = "%[2]s"
}`, cisResourceGroup, cisInstance)

}
