// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
)

func TestAccIBMDNSDomain_Basic(t *testing.T) {
	var dns_domain datatypes.Dns_Domain

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDNSDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(config, domainName1, target1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.acceptance_test_dns_domain-1", &dns_domain),
					testAccCheckIBMDNSDomainAttributes(&dns_domain),
					saveIBMDNSDomainId(&dns_domain, &firstDnsId),
					resource.TestCheckResourceAttr(
						"ibm_dns_domain.acceptance_test_dns_domain-1", "name", domainName1),
					resource.TestCheckResourceAttr(
						"ibm_dns_domain.acceptance_test_dns_domain-1", "target", target1),
				),
				Destroy: false,
			},
			{
				Config: fmt.Sprintf(config, domainName2, target1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.acceptance_test_dns_domain-1", &dns_domain),
					testAccCheckIBMDNSDomainAttributes(&dns_domain),
					resource.TestCheckResourceAttr(
						"ibm_dns_domain.acceptance_test_dns_domain-1", "name", domainName2),
					resource.TestCheckResourceAttr(
						"ibm_dns_domain.acceptance_test_dns_domain-1", "target", target1),
					testAccCheckIBMDNSDomainChanged(&dns_domain),
				),
				Destroy: false,
			},
			{
				Config: fmt.Sprintf(config, domainName2, target2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.acceptance_test_dns_domain-1", &dns_domain),
					testAccCheckIBMDNSDomainAttributes(&dns_domain),
					resource.TestCheckResourceAttr(
						"ibm_dns_domain.acceptance_test_dns_domain-1", "name", domainName2),
					resource.TestCheckResourceAttr(
						"ibm_dns_domain.acceptance_test_dns_domain-1", "target", target2),
				),
				Destroy: false,
			},
		},
	})
}

func TestAccIBMDNSDomainWithTag(t *testing.T) {
	var dns_domain datatypes.Dns_Domain

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDNSDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(configWithTag, domainName1, target1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.acceptance_test_dns_domain-1", &dns_domain),
					testAccCheckIBMDNSDomainAttributes(&dns_domain),
					saveIBMDNSDomainId(&dns_domain, &firstDnsId),
					resource.TestCheckResourceAttr(
						"ibm_dns_domain.acceptance_test_dns_domain-1", "name", domainName1),
					resource.TestCheckResourceAttr(
						"ibm_dns_domain.acceptance_test_dns_domain-1", "target", target1),
					resource.TestCheckResourceAttr(
						"ibm_dns_domain.acceptance_test_dns_domain-1", "tags.#", "2"),
				),
				Destroy: false,
			},
			{
				Config: fmt.Sprintf(configWithUpdatedTag, domainName1, target1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainExists("ibm_dns_domain.acceptance_test_dns_domain-1", &dns_domain),
					testAccCheckIBMDNSDomainAttributes(&dns_domain),
					resource.TestCheckResourceAttr(
						"ibm_dns_domain.acceptance_test_dns_domain-1", "name", domainName1),
					resource.TestCheckResourceAttr(
						"ibm_dns_domain.acceptance_test_dns_domain-1", "target", target1),
					resource.TestCheckResourceAttr(
						"ibm_dns_domain.acceptance_test_dns_domain-1", "tags.#", "3"),
				),
				Destroy: false,
			},
		},
	})
}

func testAccCheckIBMDNSDomainDestroy(s *terraform.State) error {
	service := services.GetDnsDomainService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dns_domain" {
			continue
		}

		dnsId, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the domain
		_, err := service.Id(dnsId).GetObject()

		if err == nil {
			return fmt.Errorf("Dns Domain with id %d still exists", dnsId)
		}
	}

	return nil
}

func testAccCheckIBMDNSDomainAttributes(dns *datatypes.Dns_Domain) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if name := sl.Get(dns.Name); name == "" {
			return errors.New("Empty dns domain name")
		}

		// find a record with host @; that will have the current target.
		foundTarget := false
		for _, record := range dns.ResourceRecords {
			if *record.Type == "a" && *record.Host == "@" {
				foundTarget = true
				break
			}
		}

		if !foundTarget {
			return fmt.Errorf("Target record not found for dns domain %s (%d)", sl.Get(dns.Name), sl.Get(dns.Id))
		}

		if id := sl.Get(dns.Id); id == 0 {
			return fmt.Errorf("Bad dns domain id: %d", id)
		}

		return nil
	}
}

func saveIBMDNSDomainId(dns *datatypes.Dns_Domain, id_holder *int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		*id_holder = *dns.Id

		return nil
	}
}

func testAccCheckIBMDNSDomainChanged(dns *datatypes.Dns_Domain) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		service := services.GetDnsDomainService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())

		_, err := service.Id(firstDnsId).Mask(
			"id,name,updateDate,resourceRecords",
		).GetObject()
		if err == nil {
			return fmt.Errorf("Dns domain with id %d still exists", firstDnsId)
		}

		return nil
	}
}

func testAccCheckIBMDNSDomainExists(n string, dns_domain *datatypes.Dns_Domain) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("[ERROR] Not  found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		dns_id, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetDnsDomainService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())
		found_domain, err := service.Id(dns_id).Mask(
			"id,name,updateDate,resourceRecords",
		).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*found_domain.Id)) != rs.Primary.ID {
			return errors.New("Record not found")
		}

		*dns_domain = found_domain

		return nil
	}
}

var config = `
resource "ibm_dns_domain" "acceptance_test_dns_domain-1" {
	name = "%s"
	target = "%s"
}
`

var configWithTag = `
resource "ibm_dns_domain" "acceptance_test_dns_domain-1" {
	name = "%s"
	target = "%s"
	tags = ["one", "two"]
}
`
var configWithUpdatedTag = `
resource "ibm_dns_domain" "acceptance_test_dns_domain-1" {
	name = "%s"
	target = "%s"
	tags = ["one", "two", "three"]
}
`

var domainName1 = fmt.Sprintf("tfuatdomain%s.com", acctest.RandString(10))
var domainName2 = fmt.Sprintf("tfuatdomain%s.com", acctest.RandString(10))
var target1 = "172.16.0.100"
var target2 = "172.16.0.101"
var firstDnsId = 0
