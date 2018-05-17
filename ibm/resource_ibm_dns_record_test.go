package ibm

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMDNSRecord_Basic(t *testing.T) {
	var dns_domain datatypes.Dns_Domain
	var dns_domain_record datatypes.Dns_Domain_ResourceRecord

	domainName := fmt.Sprintf("tfuatdomainr%s.ibm.com", acctest.RandString(10))
	host1 := acctest.RandString(10) + "ibm.com"
	host2 := acctest.RandString(10) + "ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDNSDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDNSRecordConfigBasic(domainName, host1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.test_dns_domain_records", &dns_domain),
					testAccCheckIBMDNSRecordExists("ibm_dns_record.recordA", &dns_domain_record),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "data", "127.0.0.1"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "expire", "900"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "minimum_ttl", "90"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "mx_priority", "1"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "refresh", "1"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "host", host1),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "responsible_person", "user@softlayer.com"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "ttl", "900"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "retry", "1"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "type", "a"),
				),
			},
			{
				Config: testAccCheckIBMDNSRecordConfigBasic(domainName, host2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.test_dns_domain_records", &dns_domain),
					testAccCheckIBMDNSRecordExists("ibm_dns_record.recordA", &dns_domain_record),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "host", host2),
				),
			},
		},
	})
}

func TestAccIBMDNSRecord_Types(t *testing.T) {
	var dns_domain datatypes.Dns_Domain
	var dns_domain_record datatypes.Dns_Domain_ResourceRecord

	domainName := acctest.RandString(10) + "dnstest.ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDNSDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckIBMDNSRecordConfig_all_types, domainName, "_tcp"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.test_dns_domain_record_types", &dns_domain),
					testAccCheckIBMDNSRecordExists("ibm_dns_record.recordA", &dns_domain_record),
					testAccCheckIBMDNSRecordExists("ibm_dns_record.recordAAAA", &dns_domain_record),
					testAccCheckIBMDNSRecordExists("ibm_dns_record.recordCNAME", &dns_domain_record),
					testAccCheckIBMDNSRecordExists("ibm_dns_record.recordMX", &dns_domain_record),
					testAccCheckIBMDNSRecordExists("ibm_dns_record.recordSPF", &dns_domain_record),
					testAccCheckIBMDNSRecordExists("ibm_dns_record.recordTXT", &dns_domain_record),
					testAccCheckIBMDNSRecordExists("ibm_dns_record.recordSRV", &dns_domain_record),
				),
			},

			{
				Config: fmt.Sprintf(testAccCheckIBMDNSRecordConfig_all_types, domainName, "_udp"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.test_dns_domain_record_types", &dns_domain),
					resource.TestCheckResourceAttr("ibm_dns_record.recordSRV", "protocol", "_udp"),
				),
			},
		},
	})
}

func TestAccIBMDNSRecordWithTag(t *testing.T) {
	var dns_domain datatypes.Dns_Domain
	var dns_domain_record datatypes.Dns_Domain_ResourceRecord

	domainName := fmt.Sprintf("tfuatdomainr%s.ibm.com", acctest.RandString(9))
	host1 := acctest.RandString(10) + "ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDNSDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDNSRecordWithTag(domainName, host1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.test_dns_domain_records", &dns_domain),
					testAccCheckIBMDNSRecordExists("ibm_dns_record.recordA", &dns_domain_record),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "data", "127.0.0.1"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "expire", "900"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "minimum_ttl", "90"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "mx_priority", "1"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "refresh", "1"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "host", host1),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "responsible_person", "user@softlayer.com"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "ttl", "900"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "retry", "1"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "type", "a"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "host", host1),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMDNSRecordWithUpdatedTag(domainName, host1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.test_dns_domain_records", &dns_domain),
					testAccCheckIBMDNSRecordExists("ibm_dns_record.recordA", &dns_domain_record),
					resource.TestCheckResourceAttr("ibm_dns_record.recordA", "tags.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMDNSRecord_MX_PRIORITY(t *testing.T) {
	var dns_domain datatypes.Dns_Domain
	var dns_domain_record datatypes.Dns_Domain_ResourceRecord

	domainName := fmt.Sprintf("tfuatdomainr%s.ibm.com", acctest.RandString(10))
	host1 := acctest.RandString(10) + "ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDNSDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDNSRecordMX(domainName, host1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.test_dns_domain_records", &dns_domain),
					testAccCheckIBMDNSRecordExists("ibm_dns_record.recordMX", &dns_domain_record),
					resource.TestCheckResourceAttr("ibm_dns_record.recordMX", "data", "email.example.com"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordMX", "expire", "0"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordMX", "minimum_ttl", "0"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordMX", "mx_priority", "0"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordMX", "refresh", "0"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordMX", "host", host1),
					resource.TestCheckResourceAttr("ibm_dns_record.recordMX", "responsible_person", "user@softlayer.com"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordMX", "ttl", "900"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordMX", "type", "mx"),
				),
			},
		},
	})
}

func TestAccIBMDNSRecord_SRV_PRIORITY_ZERO(t *testing.T) {
	var dns_domain datatypes.Dns_Domain
	var dns_domain_record datatypes.Dns_Domain_ResourceRecord

	domainName := fmt.Sprintf("tfuatdomainr%s.ibm.com", acctest.RandString(10))
	protocol := "_udp"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDNSDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDNSRecordSRVWithPriorityZero(domainName, protocol),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.test_dns_domain_records", &dns_domain),
					resource.TestCheckResourceAttr("ibm_dns_record.recordSRV", "protocol", "_udp"),
					testAccCheckIBMDNSRecordExists("ibm_dns_record.recordSRV", &dns_domain_record),
					resource.TestCheckResourceAttr("ibm_dns_record.recordSRV", "priority", "0"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordSRV", "weight", "0"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordSRV", "host", "hosta-srv.com"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordSRV", "responsible_person", "user@softlayer.com"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordSRV", "ttl", "900"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordSRV", "type", "srv"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordSRV", "port", "8080"),
					resource.TestCheckResourceAttr("ibm_dns_record.recordSRV", "service", "_mail"),
				),
			},
		},
	})
}

func testAccCheckIBMDNSRecordExists(n string, dns_domain_record *datatypes.Dns_Domain_ResourceRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		dns_id, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetDnsDomainResourceRecordService(testAccProvider.Meta().(ClientSession).SoftLayerSession())
		found_domain_record, err := service.Id(dns_id).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*found_domain_record.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record %d not found", dns_id)
		}

		*dns_domain_record = found_domain_record

		return nil
	}
}

func testAccCheckIBMDNSRecordConfigBasic(domainName, hostname string) string {
	return fmt.Sprintf(`
resource "ibm_dns_domain" "test_dns_domain_records" {
	name = "%s"
	target = "172.16.0.100"
}

resource "ibm_dns_record" "recordA" {
    data = "127.0.0.1"
    domain_id = "${ibm_dns_domain.test_dns_domain_records.id}"
    expire = 900
    minimum_ttl = 90
    mx_priority = 1
    refresh = 1
    host = "%s"
    responsible_person = "user@softlayer.com"
    ttl = 900
    retry = 1
    type = "a"
}`, domainName, hostname)
}

var testAccCheckIBMDNSRecordConfig_all_types = `
resource "ibm_dns_domain" "test_dns_domain_record_types" {
	name = "%s"
	target = "172.16.12.100"
}

resource "ibm_dns_record" "recordA" {
    data = "127.0.0.1"
    domain_id = "${ibm_dns_domain.test_dns_domain_record_types.id}"
    host = "hosta.com"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "a"
}

resource "ibm_dns_record" "recordAAAA" {
    data = "fe80::202:b3ff:fe1e:8329"
    domain_id = "${ibm_dns_domain.test_dns_domain_record_types.id}"
    host = "hosta-2.com"
    responsible_person = "user2changed@softlayer.com"
    ttl = 1000
    type = "aaaa"
}

resource "ibm_dns_record" "recordCNAME" {
    data = "testsssaaaass.com."
    domain_id = "${ibm_dns_domain.test_dns_domain_record_types.id}"
    host = "hosta-cname.com"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "cname"
}

resource "ibm_dns_record" "recordMX" {
    data = "email.example.com."
    domain_id = "${ibm_dns_domain.test_dns_domain_record_types.id}"
    host = "hosta-mx.com"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "mx"
}

resource "ibm_dns_record" "recordSPF" {
    data = "v=spf1 mx:mail.example.org ~all"
    domain_id = "${ibm_dns_domain.test_dns_domain_record_types.id}"
    host = "hosta-spf"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "spf"
}

resource "ibm_dns_record" "recordTXT" {
    data = "127.0.0.1"
    domain_id = "${ibm_dns_domain.test_dns_domain_record_types.id}"
    host = "hosta-txt.com"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "txt"
}

resource "ibm_dns_record" "recordSRV" {
    data = "ns1.example.org"
    domain_id = "${ibm_dns_domain.test_dns_domain_record_types.id}"
    host = "hosta-srv.com"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "srv"
	port = 8080
	priority = 3
	protocol = "%s"
	weight = 3
	service = "_mail"
}
`

func testAccCheckIBMDNSRecordWithTag(domainName, hostname string) string {
	return fmt.Sprintf(`
resource "ibm_dns_domain" "test_dns_domain_records" {
	name = "%s"
	target = "172.16.0.100"
}

resource "ibm_dns_record" "recordA" {
    data = "127.0.0.1"
    domain_id = "${ibm_dns_domain.test_dns_domain_records.id}"
    expire = 900
    minimum_ttl = 90
    mx_priority = 1
    refresh = 1
    host = "%s"
    responsible_person = "user@softlayer.com"
    ttl = 900
    retry = 1
	type = "a"
	tags = ["one", "two"]
}`, domainName, hostname)
}

func testAccCheckIBMDNSRecordWithUpdatedTag(domainName, hostname string) string {
	return fmt.Sprintf(`
resource "ibm_dns_domain" "test_dns_domain_records" {
	name = "%s"
	target = "172.16.0.102"
}

resource "ibm_dns_record" "recordA" {
    data = "127.0.0.1"
    domain_id = "${ibm_dns_domain.test_dns_domain_records.id}"
    expire = 900
    minimum_ttl = 90
    mx_priority = 1
    refresh = 1
    host = "%s"
    responsible_person = "user@softlayer.com"
    ttl = 900
    retry = 1
	type = "a"
	tags = ["one", "two", "three"]
}`, domainName, hostname)
}

func testAccCheckIBMDNSRecordMX(domainName, hostname string) string {
	return fmt.Sprintf(`
resource "ibm_dns_domain" "test_dns_domain_records" {
	name = "%s"
	target = "172.16.0.100"
}

resource "ibm_dns_record" "recordMX" {
    data = "email.example.com"
    domain_id = "${ibm_dns_domain.test_dns_domain_records.id}"
    host = "%s"
    responsible_person = "user@softlayer.com"
    ttl = 900
	type = "mx"
	mx_priority = 0
}`, domainName, hostname)
}

func testAccCheckIBMDNSRecordSRVWithPriorityZero(domainName, protocol string) string {
	return fmt.Sprintf(`
resource "ibm_dns_domain" "test_dns_domain_records" {
	name = "%s"
	target = "172.16.0.100"
}

resource "ibm_dns_record" "recordSRV" {
    data = "ns1.example.org"
    domain_id = "${ibm_dns_domain.test_dns_domain_records.id}"
    host = "hosta-srv.com"
    responsible_person = "user@softlayer.com"
    ttl = 900
    type = "srv"
	port = 8080
	priority = 0
	protocol = "%s"
	weight = 0
	service = "_mail"
}`, domainName, protocol)
}
