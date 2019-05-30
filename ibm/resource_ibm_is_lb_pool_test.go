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

func TestAccIBMISLBPool_basic(t *testing.T) {
	var lb *models.Pool
	vpcname := fmt.Sprintf("terraformLBuat_vpc_%d", acctest.RandInt())
	subnetname := fmt.Sprintf("terraformLBuat_create_step_name_%d", acctest.RandInt())
	name := fmt.Sprintf("tf_create_step_name_%d", acctest.RandInt())
	poolName := fmt.Sprintf("tf_lbpoolc_%d", acctest.RandInt())
	poolName1 := fmt.Sprintf("tf_lbpoolu_%d", acctest.RandInt())
	alg1 := "round_robin"
	protocol1 := "http"
	delay1 := "45"
	retries1 := "5"
	timeout1 := "15"
	healthType1 := "http"

	alg2 := "least_connections"
	protocol2 := "tcp"
	delay2 := "60"
	retries2 := "3"
	timeout2 := "30"
	healthType2 := "tcp"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISLBPoolDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISLBPoolConfig(vpcname, subnetname, ISZoneName, ISCIDR, name, poolName, alg1, protocol1, delay1, retries1, timeout1, healthType1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBPoolExists("ibm_is_lb_pool.testacc_lb_pool", &lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb.testacc_LB", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "name", poolName),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "algorithm", alg1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "protocol", protocol1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_delay", delay1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_retries", retries1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_timeout", timeout1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_type", healthType1),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMISLBPoolConfig(vpcname, subnetname, ISZoneName, ISCIDR, name, poolName1, alg2, protocol2, delay2, retries2, timeout2, healthType2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISLBPoolExists("ibm_is_lb_pool.testacc_lb_pool", &lb),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "name", poolName1),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "algorithm", alg2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "protocol", protocol2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_delay", delay2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_retries", retries2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_timeout", timeout2),
					resource.TestCheckResourceAttr(
						"ibm_is_lb_pool.testacc_lb_pool", "health_type", healthType2),
				),
			},
		},
	})
}

func testAccCheckIBMISLBPoolDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	LBC := lbaas.NewLoadBalancerClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_lb_pool" {
			continue
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		lbID := parts[0]
		lbPoolID := parts[1]
		_, err = LBC.GetPool(lbID, lbPoolID)

		if err == nil {
			return fmt.Errorf("LB Pool still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISLBPoolExists(n string, lbPool **models.Pool) resource.TestCheckFunc {
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
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		lbID := parts[0]
		lbPoolID := parts[1]
		foundLBPool, err := client.GetPool(lbID, lbPoolID)

		if err != nil {
			return err
		}

		*lbPool = foundLBPool
		return nil
	}
}

func testAccCheckIBMISLBPoolConfig(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType string) string {
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
		algorithm = "%s"
		protocol = "%s"
		health_delay= %s
		health_retries = %s
		health_timeout = %s
		health_type = "%s"
}`, vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType)

}
