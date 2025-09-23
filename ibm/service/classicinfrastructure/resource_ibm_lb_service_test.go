// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccBluemixIBMLbService_Basic(t *testing.T) {
	hostname := acctest.RandString(16)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckBluemixIBMLbServiceConfig_basic(hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb_service.test_service", "port", "80"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service.test_service", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service.test_service", "weight", "1"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service.test_service", "health_check_type", "DNS"),
					resource.TestCheckResourceAttrSet(
						"ibm_lb_service.test_service", "service_group_id"),
				),
			},
			{
				ResourceName:      "ibm_lb_service.test_service",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBluemixIBMLbServiceWithTag(t *testing.T) {
	hostname := acctest.RandString(16)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckBluemixIBMLbServiceWithTag(hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb_service.test_service", "port", "80"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service.test_service", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service.test_service", "weight", "1"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service.test_service", "health_check_type", "DNS"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service.test_service", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckBluemixIBMLbServiceWithUpdatedTag(hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb_service.test_service", "port", "80"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service.test_service", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service.test_service", "weight", "1"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service.test_service", "health_check_type", "DNS"),
					resource.TestCheckResourceAttr(
						"ibm_lb_service.test_service", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckBluemixIBMLbServiceConfig_basic(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "test_server_1" {
    hostname = "%s"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_9_64"
    datacenter = "dal06"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25, 10, 20]
    user_metadata = "{\"value\":\"newvalue\"}"
    dedicated_acct_host_only = true
    local_disk = false
}

resource "ibm_lb" "testacc_foobar_lb" {
    connections = 250
    datacenter    = "dal06"
    ha_enabled  = false
    dedicated = false
}

resource "ibm_lb_service_group" "test_service_group" {
    port = 82
    routing_method = "CONSISTENT_HASH_IP"
    routing_type = "HTTP"
    load_balancer_id = "${ibm_lb.testacc_foobar_lb.id}"
    allocation = 100
}

resource "ibm_lb_service" "test_service" {
    port = 80
    enabled = true
    service_group_id = "${ibm_lb_service_group.test_service_group.service_group_id}"
    weight = 1
    health_check_type = "DNS"
    ip_address_id = "${ibm_compute_vm_instance.test_server_1.ip_address_id}"
}`, hostname)
}

func testAccCheckBluemixIBMLbServiceWithTag(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "test_server_1" {
    hostname = "%s"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_9_64"
    datacenter = "dal06"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25, 10, 20]
    user_metadata = "{\"value\":\"newvalue\"}"
    dedicated_acct_host_only = true
    local_disk = false
}

resource "ibm_lb" "testacc_foobar_lb" {
    connections = 250
    datacenter    = "dal06"
    ha_enabled  = false
    dedicated = false
    tags = ["one", "two"]
}

resource "ibm_lb_service_group" "test_service_group" {
    port = 82
    routing_method = "CONSISTENT_HASH_IP"
    routing_type = "HTTP"
    load_balancer_id = "${ibm_lb.testacc_foobar_lb.id}"
    allocation = 100
}

resource "ibm_lb_service" "test_service" {
    port = 80
    enabled = true
    service_group_id = "${ibm_lb_service_group.test_service_group.service_group_id}"
    weight = 1
    health_check_type = "DNS"
    ip_address_id = "${ibm_compute_vm_instance.test_server_1.ip_address_id}"
    tags = ["one", "two"]
}`, hostname)
}

func testAccCheckBluemixIBMLbServiceWithUpdatedTag(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_vm_instance" "test_server_1" {
    hostname = "%s"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_9_64"
    datacenter = "dal06"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25, 10, 20]
    user_metadata = "{\"value\":\"newvalue\"}"
    dedicated_acct_host_only = true
    local_disk = false
}

resource "ibm_lb" "testacc_foobar_lb" {
    connections = 250
    datacenter    = "dal06"
    ha_enabled  = false
    dedicated = false
    tags = ["one", "two", "three"]
}

resource "ibm_lb_service_group" "test_service_group" {
    port = 82
    routing_method = "CONSISTENT_HASH_IP"
    routing_type = "HTTP"
    load_balancer_id = "${ibm_lb.testacc_foobar_lb.id}"
    allocation = 100
}

resource "ibm_lb_service" "test_service" {
    port = 80
    enabled = true
    service_group_id = "${ibm_lb_service_group.test_service_group.service_group_id}"
    weight = 1
    health_check_type = "DNS"
    ip_address_id = "${ibm_compute_vm_instance.test_server_1.ip_address_id}"
    tags = ["one", "two", "three"]
}`, hostname)
}
