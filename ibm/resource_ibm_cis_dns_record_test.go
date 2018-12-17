package ibm

import (
	"fmt"
	"testing"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMCisDNSRecord_Basic(t *testing.T) {
	//t.Parallel()
	var record v1.DnsRecord
	testName := "tf-acctest-basic"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", testName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this point it must have already been deleted from CIS.
		// If the DNS record failed to delete, the destroy of resource_ibm_cis used in this test suite will have been failed by the Resource Manager
		// and test execution aborted prior to this test.
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigBasic(testName, cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", testName+"."+cis_domain),
					resource.TestCheckResourceAttr(
						resourceName, "content", "192.168.0.10"),
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
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCisDNSRecordConfigBasic("test", cis_domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
					resource.TestCheckResourceAttr(name, "name", "test."+cis_domain),
					resource.TestCheckResourceAttr(name, "content", "192.168.0.10"),
					resource.TestCheckResourceAttr(name, "data.%", "0"),
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

func TestAccIBMCisDNSRecord_CaseInsensitive(t *testing.T) {
	//t.Parallel()
	var record v1.DnsRecord
	testName := "tf-acctest-case-insensitive"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", "test")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigCaseSensitive(testName, cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", testName+"."+cis_domain),
				),
			},
			{
				Config:   testAccCheckIBMCisDNSRecordConfigCaseSensitive("tf-acctest-CASE-INSENSITIVE", cis_domain),
				PlanOnly: true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", "tf-acctest-case-insensitive"+"."+cis_domain),
				),
			},
		},
	})
}

func TestAccIBMCisDNSRecord_Apex(t *testing.T) {
	//t.Parallel()
	var record v1.DnsRecord
	testName := "test"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", testName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// Remove check destroy as this occurs after the CIS instance is deleted and fails with an auth error
		//CheckDestroy: testAccCheckIBMCisDNSRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigApex(testName, cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						// @ is replaced by domain name by CIS
						resourceName, "name", cis_domain),
					resource.TestCheckResourceAttr(
						resourceName, "content", "192.168.0.10"),
				),
			},
		},
	})
}

func testAccCheckIBMCisDNSRecordDestroy(s *terraform.State) error {
	cisClient, _ := testAccProvider.Meta().(ClientSession).CisAPI()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_dns_record" {
			continue
		}

		_, err := cisClient.Dns().GetDns(rs.Primary.Attributes["cis_id"], rs.Primary.Attributes["domain_id"], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Record still exists")
		}
	}

	return nil
}

func testAccCheckIBMCisDNSRecordExists(n string, record *v1.DnsRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		recordId, zoneId, _, _ := convertTfToCisThreeVar(rs.Primary.ID)
		foundRecord, err := cisClient.Dns().GetDns(rs.Primary.Attributes["cis_id"], zoneId, recordId)
		if err != nil {
			return err
		}

		if foundRecord.Id != recordId {
			return fmt.Errorf("Record not found")
		}

		record = foundRecord

		return nil
	}
}

func testAccCheckIBMCisDNSRecordConfigBasic(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.%[1]s.id}"

    name = "%[1]s"
    content = "192.168.0.10"
    type = "A"
}`, resourceId)
}

func testAccCheckIBMCisDNSRecordConfigCaseSensitive(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic("test", cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "test" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.test.id}"

    name = "%[1]s"
    content = "192.168.0.10"
    type = "A"
}`, resourceId)
}

func testAccCheckIBMCisDNSRecordConfigApex(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.%[1]s.id}"
    name = "@"
    content = "192.168.0.10"
    type = "A"
}`, resourceId)
}

func testAccCheckIBMCisDNSRecordConfigLOC(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.%[1]s.id}"
    name = "%[1]s"
    data {
      "lat_degrees" =  "37"
      "lat_minutes" = "46"
      "lat_seconds" = "46"
      "lat_direction" = "N"
      "long_degrees" = "122"
      "long_minutes" = "23"
      "long_seconds" = "35"
      "long_direction" = "W"
      "altitude" = 0
      "size" = 100
      "precision_horz" = 0
      "precision_vert" = 0
    }
    type = "LOC"
}`, resourceId)
}

func testAccCheckIBMCisDNSRecordConfigSRV(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "foobar" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.%[1]s.id}"
    name = "%[1]s"
    data {
      "priority" = 5
      "weight" = 0
      "port" = 5222
      "target" = "talk.l.google.com"
      "service" = "_xmpp-client"
      "proto" = "_tcp"
    }
    type = "SRV"
}`, resourceId)
}

func testAccCheckIBMCisDNSRecordConfigProxied(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "foobar" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.%[1]s.id}"

    name = "%[1]s"
    content = "%[1]s"
    type = "CNAME"
    proxied = true
}`, resourceId)
}
