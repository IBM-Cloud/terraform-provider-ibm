package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

func TestNetworkACLGen1(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISNetworkACLConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.isExampleACL", nwACL),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "name", "is-example-acl"),
					resource.TestCheckResourceAttr(
						"ibm_is_network_acl.isExampleACL", "rules.#", "2"),
				),
			},
		},
	})
}

func TestNetworkACLGen2(t *testing.T) {
	var nwACL string
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkNetworkACLDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISNetworkACLConfig1(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISNetworkACLExists("ibm_is_network_acl.isExampleACL", nwACL),
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
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()
	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_network_acl" {
				continue
			}

			getnwacloptions := &vpcclassicv1.GetNetworkAclOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetNetworkAcl(getnwacloptions)
			if err == nil {
				return fmt.Errorf("network acl still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_network_acl" {
				continue
			}

			getnwacloptions := &vpcv1.GetNetworkAclOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetNetworkAcl(getnwacloptions)
			if err == nil {
				return fmt.Errorf("network acl still exists: %s", rs.Primary.ID)
			}
		}
	}
	return nil
}

func testAccCheckIBMISNetworkACLExists(n, nwACL string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fmt.Println("siv ", s.RootModule().Resources)
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

		if userDetails.generation == 1 {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
			getnwacloptions := &vpcclassicv1.GetNetworkAclOptions{
				ID: &rs.Primary.ID,
			}
			foundNwACL, _, err := sess.GetNetworkAcl(getnwacloptions)
			if err != nil {
				return err
			}
			nwACL = *foundNwACL.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getnwacloptions := &vpcv1.GetNetworkAclOptions{
				ID: &rs.Primary.ID,
			}
			foundNwACL, _, err := sess.GetNetworkAcl(getnwacloptions)
			if err != nil {
				return err
			}
			nwACL = *foundNwACL.ID
		}
		return nil
	}
}

func testAccCheckIBMISNetworkACLConfig() string {
	return fmt.Sprintf(`
	resource "ibm_is_network_acl" "isExampleACL" {
		name = "is-example-acl"
		rules {
		  name        = "outbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "outbound"
		  icmp {
			code = 1
			type = 1
		  }
		  # Optionals :
		  # port_max =
		  # port_min =
		}
		rules {
		  name        = "inbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  icmp {
			code = 1
			type = 1
		  }
		  # Optionals :
		  # port_max =
		  # port_min =
		}
	  }
	`)
}

func testAccCheckIBMISNetworkACLConfig1() string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "vpctest"
	  }

	resource "ibm_is_network_acl" "isExampleACL" {
		name = "is-example-acl"
		vpc  = ibm_is_vpc.testacc_vpc.id
		rules {
		  name        = "outbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "outbound"
		  icmp {
			code = 1
			type = 1
		  }
		  # Optionals :
		  # port_max =
		  # port_min =
		}
		rules {
		  name        = "inbound"
		  action      = "allow"
		  source      = "0.0.0.0/0"
		  destination = "0.0.0.0/0"
		  direction   = "inbound"
		  icmp {
			code = 1
			type = 1
		  }
		  # Optionals :
		  # port_max =
		  # port_min =
		}
	  }
	`)
}
