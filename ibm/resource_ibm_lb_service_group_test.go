package ibm

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMLbServiceGroup_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbServiceGroupConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "port", "82"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "routing_method", "CONSISTENT_HASH_IP"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "routing_type", "HTTP"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "allocation", "50"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group2", "port", "83"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group2", "routing_method", "CONSISTENT_HASH_IP"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group2", "routing_type", "TCP"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group2", "allocation", "50"),
					resource.TestCheckResourceAttrSet(
						"ibm_lb_service_group.test_service_group2", "service_group_id"),
					resource.TestCheckResourceAttrSet(
						"ibm_lb_service_group.test_service_group2", "load_balancer_id"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMLbServiceGroupConfig_WithUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "port", "80"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "routing_method", "CONSISTENT_HASH_IP"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "routing_type", "TCP"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "allocation", "30"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group2", "port", "81"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group2", "routing_method", "CONSISTENT_HASH_IP"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group2", "routing_type", "HTTP"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group2", "allocation", "30"),
					resource.TestCheckResourceAttrSet(
						"ibm_lb_service_group.test_service_group2", "service_group_id"),
					resource.TestCheckResourceAttrSet(
						"ibm_lb_service_group.test_service_group2", "load_balancer_id"),
				),
			},

			resource.TestStep{
				ResourceName:      "ibm_lb_service_group.test_service_group1",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"virtual_server_id",
				},
			},
		},
	})
}

func TestAccIBMLbServiceGroupWithTag(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLbServiceGroupWithTag,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "port", "82"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "routing_method", "CONSISTENT_HASH_IP"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "routing_type", "HTTP"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "allocation", "50"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMLbServiceGroupWithUpdatedTag,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "port", "82"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "routing_method", "CONSISTENT_HASH_IP"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "routing_type", "HTTP"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "allocation", "50"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service_group.test_service_group1", "tags.#", "3"),
				),
			},
		},
	})
}

const testAccCheckIBMLbServiceGroupConfig_basic = `
resource "ibm_lb" "testacc_foobar_lb" {
    connections = 250
    datacenter    = "tor01"
    ha_enabled  = false
}

resource "ibm_lb_service_group" "test_service_group1" {
    port = 82
    routing_method = "CONSISTENT_HASH_IP"
    routing_type = "HTTP"
    load_balancer_id = "${ibm_lb.testacc_foobar_lb.id}"
    allocation = 50
}

resource "ibm_lb_service_group" "test_service_group2" {
    port = 83
    routing_method = "CONSISTENT_HASH_IP"
    routing_type = "TCP"
    load_balancer_id = "${ibm_lb.testacc_foobar_lb.id}"
    allocation = 50
}
`

const testAccCheckIBMLbServiceGroupConfig_WithUpdate = `
resource "ibm_lb" "testacc_foobar_lb" {
    connections = 250
    datacenter    = "tor01"
    ha_enabled  = false
}

resource "ibm_lb_service_group" "test_service_group1" {
    port = 80
    routing_method = "CONSISTENT_HASH_IP"
    routing_type = "TCP"
    load_balancer_id = "${ibm_lb.testacc_foobar_lb.id}"
    allocation = 30
}

resource "ibm_lb_service_group" "test_service_group2" {
    port = 81
    routing_method = "CONSISTENT_HASH_IP"
    routing_type = "HTTP"
    load_balancer_id = "${ibm_lb.testacc_foobar_lb.id}"
    allocation = 30
}
`

const testAccCheckIBMLbServiceGroupWithTag = `
resource "ibm_lb" "testacc_foobar_lb" {
    connections = 250
    datacenter    = "tor01"
    ha_enabled  = false
}

resource "ibm_lb_service_group" "test_service_group1" {
    port = 82
    routing_method = "CONSISTENT_HASH_IP"
    routing_type = "HTTP"
    load_balancer_id = "${ibm_lb.testacc_foobar_lb.id}"
	allocation = 50
	tags = ["one", "two"]
}
`

const testAccCheckIBMLbServiceGroupWithUpdatedTag = `
resource "ibm_lb" "testacc_foobar_lb" {
    connections = 250
    datacenter    = "tor01"
    ha_enabled  = false
}

resource "ibm_lb_service_group" "test_service_group1" {
    port = 82
    routing_method = "CONSISTENT_HASH_IP"
    routing_type = "HTTP"
    load_balancer_id = "${ibm_lb.testacc_foobar_lb.id}"
	allocation = 50
	tags = ["one", "two", "three"]
}
`
