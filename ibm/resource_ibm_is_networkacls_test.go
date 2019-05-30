package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

func TestNetworkACL(t *testing.T) {
	var nwACL *models.NetworkACL
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISNetworkACLConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.isExampleACL", &nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "name", "is-example-acl"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "rules.#", "2"),
				),
			},
		},
	})
}

func checkNetworkACLDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	nwaclC := network.NewNetworkAclClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_network_acl" {
			continue
		}

		_, err := nwaclC.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("network acl still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISNetworkACLExists(n string, nwACL **models.NetworkACL) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fmt.Println("siv ", s.RootModule().Resources)
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		nwaclC := network.NewNetworkAclClient(sess)
		foundNwACL, err := nwaclC.Get(rs.Primary.ID)

		if err != nil {
			return err
		}

		*nwACL = foundNwACL
		return nil
	}
}

func testAccCheckIBMISNetworkACLConfig() string {
	return fmt.Sprintf(`
		resource "ibm_is_network_acl" "isExampleACL" {
			name = "is-example-acl"
			rules=[
			{
				name = "egress"
				action = "allow"
				source = "0.0.0.0/0"
				destination = "0.0.0.0/0"
				direction = "egress"
				icmp=[
				{
					code = 1
					type = 1
				}]
				# Optionals : 
				# port_max = 
				# port_min = 
			},
			{
				name = "ingress"
				action = "allow"
				source = "0.0.0.0/0"
				destination = "0.0.0.0/0"
				direction = "ingress"
				icmp=[
				{
					code = 1
					type = 1
				}]
				# Optionals : 
				# port_max = 
				# port_min = 
			}
			]
		}
	`)
}
