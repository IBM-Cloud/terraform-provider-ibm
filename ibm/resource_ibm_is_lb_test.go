package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.ibm.com/Bluemix/riaas-go-client/clients/lbaas"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

func TestAccIBMISLB_basic(t *testing.T) {
	var lb *models.LoadBalancer
	vpcname := fmt.Sprintf("terraformLBuat_vpc_%d", acctest.RandInt())
	subnetname := fmt.Sprintf("terraformLBuat_create_step_name_%d", acctest.RandInt())
	name := fmt.Sprintf("tf_create_step_name_%d", acctest.RandInt())
	name1 := fmt.Sprintf("tf_update_step_name_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISLBDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISLBConfig(vpcname, subnetname, ISZoneName, ISCIDR, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", &lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttrSet(
						"ibm_is_lb.testacc_LB", "hostname"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMISLBConfig(vpcname, subnetname, ISZoneName, ISCIDR, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", &lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name1),
				),
			},
		},
	})
}

func TestAccIBMISLB_basic_private(t *testing.T) {
	var lb *models.LoadBalancer
	vpcname := fmt.Sprintf("terraformLBuat_vpc_%d", acctest.RandInt())
	subnetname := fmt.Sprintf("terraformLBuat_create_step_name_%d", acctest.RandInt())
	name := fmt.Sprintf("tf_create_step_name_%d", acctest.RandInt())
	name1 := fmt.Sprintf("tf_update_step_name_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISLBDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISLBConfigPrivate(vpcname, subnetname, ISZoneName, ISCIDR, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", &lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMISLBConfigPrivate(vpcname, subnetname, ISZoneName, ISCIDR, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBExists("ibm_is_lb.testacc_LB", &lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name1),
				),
			},
		},
	})
}

func testAccCheckIBMISLBDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	LBC := lbaas.NewLoadBalancerClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_lb" {
			continue
		}

		_, err := LBC.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("LB still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISLBExists(n string, lb **models.LoadBalancer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		client := lbaas.NewLoadBalancerClient(sess)
		foundLB, err := client.Get(rs.Primary.ID)

		if err != nil {
			return err
		}

		*lb = foundLB
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
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
}`, vpcname, subnetname, zone, cidr, name)

}

func testAccCheckIBMISLBConfigPrivate(vpcname, subnetname, zone, cidr, name string) string {
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
		type = "private"
}`, vpcname, subnetname, zone, cidr, name)

}
