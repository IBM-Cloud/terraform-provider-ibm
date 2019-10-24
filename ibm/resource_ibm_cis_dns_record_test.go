package ibm

import (
	"fmt"
	"testing"

	//v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMCisDNSRecord_Basic(t *testing.T) {
	//t.Parallel()
	var record string
	testName := "tf-acctest-basic"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", testName)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCis(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisDnsRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigCisDS_Basic(testName, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", testName+"."+cisDomainStatic),
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
				Config: testAccCheckIBMCisDNSRecordConfigCisDS_Basic("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "proxied", "false"), // default value
					resource.TestCheckResourceAttr(name, "name", "test."+cisDomainStatic),
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
	var record string
	testName := "tf-acctest-case-insensitive"
	resourceName := fmt.Sprintf("ibm_cis_dns_record.%s", testName)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisDnsRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigCaseSensitive(testName, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						resourceName, "name", testName+"."+cisDomainStatic),
				),
			},
			{
				Config: testAccCheckIBMCisDNSRecordConfigCaseSensitive("tf-acctest-CASE-INSENSITIVE", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists("ibm_cis_dns_record.tf-acctest-CASE-INSENSITIVE", &record),
					resource.TestCheckResourceAttr(
						"ibm_cis_dns_record.tf-acctest-CASE-INSENSITIVE", "name", "tf-acctest-case-insensitive."+cisDomainStatic),
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisDnsRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigApex(testName, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists(resourceName, &record),
					resource.TestCheckResourceAttr(
						// @ is replaced by domain name by CIS
						resourceName, "name", cisDomainStatic),
					resource.TestCheckResourceAttr(
						resourceName, "content", "192.168.0.10"),
				),
			},
		},
	})
}

func TestAccIBMCisDNSRecord_CreateAfterManualDestroy(t *testing.T) {
	t.Parallel()
	testName := "test_acc"
	var afterCreate, afterRecreate string
	name := "ibm_cis_dns_record.test_acc"

	afterCreate = "hello"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCis(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisDnsRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigCisDS_Basic(testName, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists(name, &afterCreate),
					testAccIBMCisManuallyDeleteDnsRecord(&afterCreate),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckIBMCisDNSRecordConfigCisDS_Basic(testName, cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists(name, &afterRecreate),
					testAccCheckIBMCisDnsRecordRecreated(&afterCreate, &afterRecreate),
				),
			},
		},
	})
}

func TestAccIBMCisDNSRecord_CreateAfterManualCisRIDestroy(t *testing.T) {
	t.Parallel()
	testName := "test_acc"
	var afterCreate, afterRecreate string
	name := "ibm_cis_dns_record.test_acc"

	afterCreate = "hello"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCis(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCisDnsRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisDNSRecordConfigCisRI_Basic(testName, cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists(name, &afterCreate),
					testAccIBMCisManuallyDeleteDnsRecord(&afterCreate),
					func(state *terraform.State) error {
						cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
						if err != nil {
							return err
						}
						for _, r := range state.RootModule().Resources {
							if r.Type == "ibm_cis_domain" {
								zoneId, cisId, _ := convertTftoCisTwoVar(r.Primary.ID)
								_ = cisClient.Zones().DeleteZone(cisId, zoneId)
								cisPtr := &cisId
								_ = testAccCisInstanceManuallyDelete(cisPtr)
							}

						}
						return nil
					},
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckIBMCisDNSRecordConfigCisRI_Basic(testName, cisDomainTest),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCisDnsRecordExists(name, &afterRecreate),
					testAccCheckIBMCisDnsRecordRecreated(&afterCreate, &afterRecreate),
				),
			},
		},
	})
}

func testAccIBMCisManuallyDeleteDnsRecord(tfRecordId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		if err != nil {
			return err
		}
		tfRecord := *tfRecordId
		recordId, zoneId, cisId, _ := convertTfToCisThreeVar(tfRecord)
		err = cisClient.Dns().DeleteDns(cisId, zoneId, recordId)
		if err != nil {
			return fmt.Errorf("Error deleting IBMCISDNS Record: %s", err)
		}
		return nil
	}
}

func testAccCheckIBMCisDnsRecordRecreated(beforeId, afterId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *beforeId == *afterId {
			return fmt.Errorf("Expected change of Record Ids, but both were %v", beforeId)
		}
		return nil
	}
}

func testAccCheckIBMCisDnsRecordDestroy(s *terraform.State) error {
	cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cis_record" {
			continue
		}

		recordId, zoneId, cisId, _ := convertTfToCisThreeVar(rs.Primary.ID)
		err = cisClient.Dns().DeleteDns(cisId, zoneId, recordId)
		if err == nil {
			return fmt.Errorf("Record still exists")
		}
	}

	return nil
}

func testAccCheckIBMCisDnsRecordExists(n string, tfRecordId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		tfRecord := *tfRecordId
		cisClient, err := testAccProvider.Meta().(ClientSession).CisAPI()
		recordId, zoneId, cisId, _ := convertTfToCisThreeVar(rs.Primary.ID)
		foundRecordPtr, err := cisClient.Dns().GetDns(rs.Primary.Attributes["cis_id"], zoneId, recordId)
		if err != nil {
			return err
		}

		foundRecord := *foundRecordPtr
		if foundRecord.Id != recordId {
			return fmt.Errorf("Record not found")
		}

		tfRecord = convertCisToTfThreeVar(foundRecord.Id, zoneId, cisId)
		*tfRecordId = tfRecord
		return nil
	}
}

func testAccCheckIBMCisDNSRecordConfigCisDS_Basic(resourceId string, cisDomain string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${data.ibm_cis.%[2]s.id}"
    domain_id = "${data.ibm_cis_domain.%[1]s.id}"

    name = "%[1]s"
    content = "192.168.0.10"
    type = "A"
}`, resourceId, cisInstance)
}

func testAccCheckIBMCisDNSRecordConfigCisRI_Basic(resourceId string, cisDomain string) string {
	return testAccCheckCisDomainConfigCisRI_basic(resourceId, cisDomain) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${ibm_cis.%[2]s.id}"
    domain_id = "${ibm_cis_domain.%[1]s.id}"

    name = "%[1]s"
    content = "192.168.0.10"
    type = "A"
}`, resourceId, "testacc_ds_cis")
}

func testAccCheckIBMCisDNSRecordConfigCaseSensitive(resourceId string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${data.ibm_cis.%[2]s.id}"
    domain_id = "${data.ibm_cis_domain.%[1]s.id}"

    name = "%[1]s"
    content = "192.168.0.10"
    type = "A"
}`, resourceId, cisInstance)
}

func testAccCheckIBMCisDNSRecordConfigApex(resourceId string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${data.ibm_cis.%[2]s.id}"
    domain_id = "${data.ibm_cis_domain.%[1]s.id}"
    name = "@"
    content = "192.168.0.10"
    type = "A"
}`, resourceId, cisInstance)
}

func testAccCheckIBMCisDNSRecordConfigLOC(resourceId string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${data.ibm_cis.%[2]s.id}"
    domain_id = "${data.ibm_cis_domain.%[1]s.id}"
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
}`, resourceId, cisInstance)
}

func testAccCheckIBMCisDNSRecordConfigSRV(resourceId string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${data.ibm_cis.%[2]s.id}"
    domain_id = "${data.ibm_cis_domain.%[1]s.id}"
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
}`, resourceId, cisInstance)
}

func testAccCheckIBMCisDNSRecordConfigProxied(resourceId string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceId, cisDomainStatic) + fmt.Sprintf(`
resource "ibm_cis_dns_record" "%[1]s" {
    cis_id = "${data.ibm_cis.%[2]s.id}"
    domain_id = "${data.ibm_cis_domain.%[1]s.id}"

    name = "%[1]s"
    content = "%[1]s"
    type = "CNAME"
    proxied = true
}`, resourceId, cisInstance)
}
