package ibm

import (
	"fmt"
	"regexp"
	"testing"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMCISDNSRecord_Basic(t *testing.T) {
	t.Parallel()
	var record v1.DnsRecord
	testName := "tf-acctest-basic"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", testName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// Remove check destroy as this occurs after the CIS instance is deleted and fails with an auth error
		//CheckDestroy: testAccCheckIBMCISDNSRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCISDNSRecordConfigBasic(testName, cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCISDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", testName+"."+cis_domain),
					resource.TestCheckResourceAttr(
						resourceName, "content", "192.168.0.10"),
					resource.TestCheckResourceAttr(
						resourceName, "data.%", "0"),
					resource.TestMatchResourceAttr(
						resourceName, "domain_id", regexp.MustCompile("^[a-z0-9]{32}$")),
				),
			},
		},
	})
}

func TestAccIBMCISDNSRecord_CaseInsensitive(t *testing.T) {
	t.Parallel()
	var record v1.DnsRecord
	testName := "tf-acctest-case-insensitive"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", "test")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// Remove check destroy as this occurs after the CIS instance is deleted and fails with an auth error
		//CheckDestroy: testAccCheckIBMCISDNSRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCISDNSRecordConfigCaseSensitive(testName, cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCISDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", testName+"."+cis_domain),
				),
			},
			{
				Config:   testAccCheckIBMCISDNSRecordConfigCaseSensitive("tf-acctest-CASE-INSENSITIVE", cis_domain),
				PlanOnly: true,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCISDNSRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", "tf-acctest-case-insensitive"+"."+cis_domain),
				),
			},
		},
	})
}

func TestAccIBMCISDNSRecord_Apex(t *testing.T) {
	t.Parallel()
	var record v1.DnsRecord
	testName := "test"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", testName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// Remove check destroy as this occurs after the CIS instance is deleted and fails with an auth error
		//CheckDestroy: testAccCheckIBMCISDNSRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCISDNSRecordConfigApex(testName, cis_domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCISDNSRecordExists(resourceName, &record),
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

func testAccCheckIBMCISDNSRecordDestroy(s *terraform.State) error {
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

func testAccCheckIBMCISDNSRecordExists(n string, record *v1.DnsRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		foundRecord, err := cisClient.Dns().GetDns(rs.Primary.Attributes["cis_id"], rs.Primary.Attributes["domain_id"], rs.Primary.ID)
		if err != nil {
			return err
		}

		if foundRecord.Id != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		record = foundRecord

		return nil
	}
}

func testAccCheckIBMCISDNSRecordConfigBasic(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.%[1]s.id}"

    name = "%[1]s"
    content = "192.168.0.10"
    type = "A"
}`, resourceId, cis_domain)
}

func testAccCheckIBMCISDNSRecordConfigCaseSensitive(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic("test", cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "test" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.test.id}"

    name = "%[1]s"
    content = "192.168.0.10"
    type = "A"
}`, resourceId, cis_domain)
}

func testAccCheckIBMCISDNSRecordConfigApex(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.%[1]s.id}"
    name = "@"
    content = "192.168.0.10"
    type = "A"
}`, resourceId, cis_domain)
}

func testAccCheckIBMCISDNSRecordConfigLOC(resourceId string, cis_domain string) string {
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
}`, resourceId, cis_domain)
}

func testAccCheckIBMCISDNSRecordConfigSRV(resourceId string, cis_domain string) string {
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
}`, resourceId, cis_domain)
}

func testAccCheckIBMCISDNSRecordConfigProxied(resourceId string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic(resourceId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "foobar" {
    cis_id = "${ibm_cis.instance.id}"
    domain_id = "${ibm_cis_domain.%[1]s.id}"

    name = "%[1]s"
    content = "%[1]s"
    type = "CNAME"
    proxied = true
}`, resourceId, cis_domain)
}
