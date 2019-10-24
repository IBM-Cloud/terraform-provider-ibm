package ibm

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
	"log"
	"strconv"
	"testing"
)

func TestAccIBMDNSDomainRegistration_Nameservers_Basic(t *testing.T) {
	var dns_domain_registration datatypes.Dns_Domain_Registration

	var config = `
resource "ibm_dns_domain_registration_nameservers" "acceptance_test_dns_domain-1" {
	name_servers = ["%[1]s", "%[2]s"]
	dns_registration_id = "${data.ibm_dns_domain_registration.wcpclouduk.id}"
}

data "ibm_dns_domain_registration" "wcpclouduk" {
    name = "%[3]s"
}

`

	var domainName1 = "wcpclouduk.com"
	var nameServer1 = "ns008.name.cloud.ibm.com"
	var nameServer2 = "ns017.name.cloud.ibm.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(config, nameServer1, nameServer2, domainName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDNSDomainRegistrationAttributes("ibm_dns_domain_registration_nameservers.acceptance_test_dns_domain-1",
						&dns_domain_registration, nameServer1, nameServer2),
					resource.TestCheckResourceAttr(
						"ibm_dns_domain_registration_nameservers.acceptance_test_dns_domain-1", "name_servers.#", "2"),
				),
				Destroy: false,
			},
		},
	})
}

func testAccCheckIBMDNSDomainRegistrationAttributes(n string, dns_reg *datatypes.Dns_Domain_Registration, ns1 string, ns2 string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		// Get name servers from DNS to verify they have been set correctly
		service := services.GetDnsDomainRegistrationService(testAccProvider.Meta().(ClientSession).SoftLayerSession())
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		dnsId, _ := strconv.Atoi(rs.Primary.ID)

		log.Printf("DNS ID of Registered name servers %v\n", dnsId)
		dns_domain_nameservers, err := service.Id(dnsId).
			Mask("nameservers.name").
			GetDomainNameservers()

		if err != nil {
			return fmt.Errorf("Error retrieving domain registration: %s", err)
		}

		if len(dns_domain_nameservers) == 0 {
			return fmt.Errorf("No domain found with id [%d]", dnsId)
		}

		log.Printf("list %v\n", dns_domain_nameservers)
		ns := make([]string, len(dns_domain_nameservers[0].Nameservers))
		for i, elem := range dns_domain_nameservers[0].Nameservers {
			ns[i] = *elem.Name
		}

		log.Printf("names %v\n", ns)

		ns1Found := false
		for _, elem := range ns {
			if elem == ns1 {
				ns1Found = true
			}
		}
		ns2Found := false
		for _, elem := range ns {
			if elem == ns2 {
				ns2Found = true
			}
		}

		if ns1Found != true || ns2Found != true {
			return fmt.Errorf("Error domain registration nameservers not set as required: %v", ns)
		}

		return nil
	}
}
