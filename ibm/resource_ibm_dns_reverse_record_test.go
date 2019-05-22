package ibm

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMDNSReverseRecord_Basic(t *testing.T) {
	var dns_domain datatypes.Dns_Domain
	var dns_domain_record datatypes.Dns_Domain_ResourceRecord

	host1 := acctest.RandString(10) + "ibm.com"
	host2 := acctest.RandString(10) + "ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDNSDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDNSReverseRecordConfigBasic(host1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.test_dns_domain_records", &dns_domain),
					testAccCheckIBMDNSRecordExists("ibm_dns_reverse_record.recordA", &dns_domain_record),
					resource.TestCheckResourceAttr("ibm_dns_reverse_record.recordA", "ipaddress", "10.132.90.123"),
					resource.TestCheckResourceAttr("ibm_dns_reverse_record.recordA", "data", host1),
					resource.TestCheckResourceAttr("ibm_dns_reverse_record.recordA", "ttl", "900"),
				),
			},
			{
				Config: testAccCheckIBMDNSReverseRecordConfigBasic(host2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.test_dns_domain_records", &dns_domain),
					testAccCheckIBMDNSRecordExists("ibm_dns_reverse_record.recordA", &dns_domain_record),
					resource.TestCheckResourceAttr("ibm_dns_reverse_record.recordA", "host", host2),
				),
			},
		},
	})
}

func testAccCheckIBMDNSReverseRecordExists(n string, dns_domain_record *datatypes.Dns_Domain_ResourceRecord) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		log.Printf("inside reverse record exist function")
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
func testAccCheckIBMDNSReverseRecordConfigBasic(hostname string) string {
	return fmt.Sprintf(`
		resource "ibm_dns_reverse_record" "testreverserecord" {
			ipaddress="158.175.87.35"
			hostname="%s"
			ttl=900
		}`, hostname)
}
