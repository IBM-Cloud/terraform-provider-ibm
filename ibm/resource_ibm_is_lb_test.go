package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISLB_basic(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfupdate%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISLBDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISLBConfig(vpcname, subnetname, ISZoneName, ISCIDR, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_LB", "hostname"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMISLBConfig(vpcname, subnetname, ISZoneName, ISCIDR, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name1),
				),
			},
		},
	})
}

func TestAccIBMISLB_basic_network(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflb-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-subnet-name-%d", acctest.RandIntRange(10, 100))
	nlbName := fmt.Sprintf("tfnlbcreate%d", acctest.RandIntRange(10, 100))
	nlbName1 := fmt.Sprintf("tfnlbupdate%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISLBDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISLBNetworkConfig(vpcname, subnetname, ISZoneName, ISCIDR, nlbName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_NLB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_NLB", "name", nlbName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_NLB", "hostname"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMISLBNetworkConfig(vpcname, subnetname, ISZoneName, ISCIDR, nlbName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_NLB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_NLB", "name", nlbName1),
				),
			},
		},
	})
}

func TestAccIBMISLB_basic_private(t *testing.T) {
	var lb string
	vpcname := fmt.Sprintf("tflbt-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflb-create-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfupdate%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISLBDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISLBConfigPrivate(vpcname, subnetname, ISZoneName, ISCIDR, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMISLBConfigPrivate(vpcname, subnetname, ISZoneName, ISCIDR, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name1),
				),
			},
		},
	})
}

func testAccCheckIBMISLBDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_lb" {
				continue
			}

			getlboptions := &vpcclassicv1.GetLoadBalancerOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetLoadBalancer(getlboptions)
			if err == nil {
				return fmt.Errorf("LB still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_lb" {
				continue
			}

			getlboptions := &vpcv1.GetLoadBalancerOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetLoadBalancer(getlboptions)
			if err == nil {
				return fmt.Errorf("LB still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckIBMISLBExists(n, lb string) resource.TestCheckFunc {
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
			getlboptions := &vpcclassicv1.GetLoadBalancerOptions{
				ID: &rs.Primary.ID,
			}
			foundLB, _, err := sess.GetLoadBalancer(getlboptions)
			if err != nil {
				return err
			}
			lb = *foundLB.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getlboptions := &vpcv1.GetLoadBalancerOptions{
				ID: &rs.Primary.ID,
			}
			foundLB, _, err := sess.GetLoadBalancer(getlboptions)
			if err != nil {
				return err
			}
			lb = *foundLB.ID
		}
		return nil
	}
}

func testAccCheckIBMISLBConfig(vpcname, subnetname, zone, cidr, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
}`, vpcname, subnetname, zone, cidr, name)

}

func testAccCheckIBMISLBNetworkConfig(vpcname, subnetname, zone, cidr, nlbName string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_NLB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
		profile = "network-fixed"
    }`, vpcname, subnetname, zone, cidr, nlbName)

}

func testAccCheckIBMISLBConfigPrivate(vpcname, subnetname, zone, cidr, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = [ibm_is_subnet.testacc_subnet.id]
		type = "private"
}`, vpcname, subnetname, zone, cidr, name)

}
