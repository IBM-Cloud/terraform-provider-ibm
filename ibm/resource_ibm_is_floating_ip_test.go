package ibm

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

func TestAccIBMISFloatingIP_basic(t *testing.T) {
	var ip string
	vpcname := fmt.Sprintf("tfip-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfip-%d", acctest.RandIntRange(10, 100))
	instancename := fmt.Sprintf("tfip-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tfip-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISFloatingIPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISFloatingIPConfig(vpcname, subnetname, sshname, publicKey, instancename, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFloatingIPExists("ibm_is_floating_ip.testacc_floatingip", ip),
					resource.TestCheckResourceAttr(
						"ibm_is_floating_ip.testacc_floatingip", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_floating_ip.testacc_floatingip", "zone", ISZoneName),
				),
			},
		},
	})
}

func TestAccIBMISFloatingIP_NoTarget(t *testing.T) {
	var ip string
	name := fmt.Sprintf("tfip-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISFloatingIPDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISFloatingIPNoTargetConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFloatingIPExists("ibm_is_floating_ip.testacc_floatingip", ip),
					resource.TestCheckResourceAttr(
						"ibm_is_floating_ip.testacc_floatingip", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_floating_ip.testacc_floatingip", "zone", ISZoneName),
				),
			},
		},
	})
}

func testAccCheckIBMISFloatingIPDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_floating_ip" {
				continue
			}

			getfipoptions := &vpcclassicv1.GetFloatingIpOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetFloatingIp(getfipoptions)
			if err == nil {
				return fmt.Errorf("Floating IP still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_floating_ip" {
				continue
			}

			getfipoptions := &vpcv1.GetFloatingIpOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetFloatingIp(getfipoptions)
			if err == nil {
				return fmt.Errorf("Floating IP still exists: %s", rs.Primary.ID)
			}
		}
	}
	return nil
}

func testAccCheckIBMISFloatingIPExists(n, ip string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
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
			getfipoptions := &vpcclassicv1.GetFloatingIpOptions{
				ID: &rs.Primary.ID,
			}
			foundip, _, err := sess.GetFloatingIp(getfipoptions)
			if err != nil {
				return err
			}
			ip = *foundip.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getfipoptions := &vpcv1.GetFloatingIpOptions{
				ID: &rs.Primary.ID,
			}
			foundip, _, err := sess.GetFloatingIp(getfipoptions)
			if err != nil {
				return err
			}
			ip = *foundip.ID
		}
		return nil
	}
}

func testAccCheckIBMISFloatingIPConfig(vpcname, subnetname, sshname, publicKey, instancename, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  port_speed = "100"
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	  }
	  
	  resource "ibm_is_floating_ip" "testacc_floatingip" {
		name   = "%s"
		target = ibm_is_instance.testacc_instance.primary_network_interface[0].id
	  }
`, vpcname, subnetname, ISZoneName, ISCIDR, sshname, publicKey, instancename, isImage, instanceProfileName, ISZoneName, name)
}

func testAccCheckIBMISFloatingIPNoTargetConfig(name string) string {
	return fmt.Sprintf(`
	  resource "ibm_is_floating_ip" "testacc_floatingip" {
		name   = "%s"
		zone   = "%s"
	  }
`, name, ISZoneName)
}
