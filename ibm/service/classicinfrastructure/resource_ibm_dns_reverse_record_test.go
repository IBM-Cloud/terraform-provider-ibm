// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMDNSReverseRecord_Basic(t *testing.T) {
	var dns_domain_record datatypes.Dns_Domain_ResourceRecord

	host1 := acctest.RandString(10) + "ibm.com"
	host2 := acctest.RandString(10) + "ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDNSDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDNSReverseRecordConfigBasic(host1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSReverseRecordExists("ibm_dns_reverse_record.testreverserecord", &dns_domain_record),
					resource.TestCheckResourceAttr("ibm_dns_reverse_record.testreverserecord", "ipaddress", "1.2.3.4"),
					resource.TestCheckResourceAttr("ibm_dns_reverse_record.testreverserecord", "data", host1),
					resource.TestCheckResourceAttr("ibm_dns_reverse_record.testreverserecord", "ttl", "900"),
				),
			},
			{
				Config: testAccCheckIBMDNSReverseRecordConfigBasic(host2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSReverseRecordExists("ibm_dns_reverse_record.testreverserecord", &dns_domain_record),
					resource.TestCheckResourceAttr("ibm_dns_reverse_record.testreverserecord", "host", host2),
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
			return fmt.Errorf("[ERROR] Not  found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		dns_id, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetDnsDomainResourceRecordService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())
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
			ipaddress="1.2.3.4"
			hostname="%s"
			ttl=900
		}`, hostname)
}
