package ibm

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISLBPoolMember_basic(t *testing.T) {
	var lb string

	vpcname := fmt.Sprintf("tflbpm-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflbpmc-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	poolName := fmt.Sprintf("tflbpoolc%d", acctest.RandIntRange(10, 100))
	port := "8080"
	port1 := "9000"
	address := "127.0.0.1"
	address1 := "192.168.0.1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISLBPoolMemberDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISLBPoolMemberConfig(vpcname, subnetname, ISZoneName, ISCIDR, name, poolName, port, address),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBPoolMemberExists("ibm_is_lb_pool_member.testacc_lb_mem", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool_member.testacc_lb_mem", "port", port),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool_member.testacc_lb_mem", "target_address", address),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMISLBPoolMemberConfig(vpcname, subnetname, ISZoneName, ISCIDR, name, poolName, port1, address1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBPoolMemberExists("ibm_is_lb_pool_member.testacc_lb_mem", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool_member.testacc_lb_mem", "port", port1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool_member.testacc_lb_mem", "target_address", address1),
				),
			},
		},
	})
}

func TestAccIBMISLBPoolMember_basic_network(t *testing.T) {
	var lb string

	vpcname := fmt.Sprintf("tflbpm-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflbpmc-name-%d", acctest.RandIntRange(10, 100))
	nlbPoolName := fmt.Sprintf("tfnlbpoolc%d", acctest.RandIntRange(10, 100))

	nlbName := fmt.Sprintf("tfnlbcreate%d", acctest.RandIntRange(10, 100))
	nlbName1 := fmt.Sprintf("tfnlbupdate%d", acctest.RandIntRange(10, 100))

	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	vsiName := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISLBPoolMemberDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISLBPoolMemberIDConfig(vpcname, subnetname, ISZoneName, ISCIDR, sshname, publicKey, vsiName, nlbName, nlbPoolName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBPoolMemberExists("ibm_is_lb_pool_member.testacc_nlb_mem", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool_member.testacc_nlb_mem", "weight", "20"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMISLBPoolMemberIDConfig(vpcname, subnetname, ISZoneName, ISCIDR, sshname, publicKey, vsiName, nlbName1, nlbPoolName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBPoolMemberExists("ibm_is_lb_pool_member.testacc_nlb_mem", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool_member.testacc_nlb_mem", "port", "8080"),
				),
			},
		},
	})
}

func testAccCheckIBMISLBPoolMemberDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_lb_pool_member" {
				continue
			}
			parts, err := idParts(rs.Primary.ID)
			if err != nil {
				return err
			}

			lbID := parts[0]
			lbPoolID := parts[1]
			lbPoolMemID := parts[2]
			getlbpmoptions := &vpcclassicv1.GetLoadBalancerPoolMemberOptions{
				LoadBalancerID: &lbID,
				PoolID:         &lbPoolID,
				ID:             &lbPoolMemID,
			}
			_, _, err1 := sess.GetLoadBalancerPoolMember(getlbpmoptions)

			if err1 == nil {
				return fmt.Errorf("LB Pool member still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_lb_pool_member" {
				continue
			}
			parts, err := idParts(rs.Primary.ID)
			if err != nil {
				return err
			}

			lbID := parts[0]
			lbPoolID := parts[1]
			lbPoolMemID := parts[2]
			getlbpmoptions := &vpcv1.GetLoadBalancerPoolMemberOptions{
				LoadBalancerID: &lbID,
				PoolID:         &lbPoolID,
				ID:             &lbPoolMemID,
			}
			_, _, err1 := sess.GetLoadBalancerPoolMember(getlbpmoptions)

			if err1 == nil {
				return fmt.Errorf("LB Pool member still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckIBMISLBPoolMemberExists(n, lbPoolMember string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		lbID := parts[0]
		lbPoolID := parts[1]
		lbPoolMemID := parts[2]
		userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()
		if userDetails.generation == 1 {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
			getlbpmoptions := &vpcclassicv1.GetLoadBalancerPoolMemberOptions{
				LoadBalancerID: &lbID,
				PoolID:         &lbPoolID,
				ID:             &lbPoolMemID,
			}
			foundLBPoolMember, _, err := sess.GetLoadBalancerPoolMember(getlbpmoptions)
			if err != nil {
				return err
			}
			lbPoolMember = *foundLBPoolMember.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getlbpmoptions := &vpcv1.GetLoadBalancerPoolMemberOptions{
				LoadBalancerID: &lbID,
				PoolID:         &lbPoolID,
				ID:             &lbPoolMemID,
			}
			foundLBPoolMember, _, err := sess.GetLoadBalancerPoolMember(getlbpmoptions)
			if err != nil {
				return err
			}
			lbPoolMember = *foundLBPoolMember.ID
		}
		return nil
	}
}

func testAccCheckIBMISLBPoolMemberConfig(vpcname, subnetname, zone, cidr, name, poolName, port, address string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
	}
	resource "ibm_is_lb_pool" "testacc_lb_pool" {
		name = "%s"
		lb = "${ibm_is_lb.testacc_LB.id}"
		algorithm = "round_robin"
		protocol = "http"
		health_delay= 45
		health_retries = 5
		health_timeout = 30
		health_type = "tcp"
	}
	resource "ibm_is_lb_pool_member" "testacc_lb_mem" {
		lb = "${ibm_is_lb.testacc_LB.id}"
		pool = "${element(split("/",ibm_is_lb_pool.testacc_lb_pool.id),1)}"
		port 	=	"%s"
		target_address = "%s"
}`, vpcname, subnetname, zone, cidr, name, poolName, port, address)
}

func testAccCheckIBMISLBPoolMemberIDConfig(vpcname, subnetname, zone, cidr, sshname, publicKey, vsiName, nlbName, nlbPoolName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
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
	resource "ibm_is_lb" "testacc_NLB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
		profile = "network-fixed"
	}
	resource "ibm_is_lb_pool" "testacc_nlb_pool" {
		name = "%s"
		lb = "${ibm_is_lb.testacc_NLB.id}"
		algorithm      = "weighted_round_robin"
        protocol       = "tcp"
        health_delay   = 60
        health_retries = 5
        health_timeout = 30
        health_type    = "tcp"
	}
	resource "ibm_is_lb_pool_member" "testacc_nlb_mem" {
		lb = "${ibm_is_lb.testacc_NLB.id}"
		pool = "${element(split("/",ibm_is_lb_pool.testacc_nlb_pool.id),1)}"
		port           = 8080
        weight = 20
		target_id = "${ibm_is_instance.testacc_instance.id}"
	}
`, vpcname, subnetname, zone, cidr, sshname, publicKey, vsiName, isImage, instanceProfileName, zone, nlbName, nlbPoolName)
}
