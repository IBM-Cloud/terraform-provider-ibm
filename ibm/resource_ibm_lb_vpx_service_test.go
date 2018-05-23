package ibm

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMLbVpxService_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMLbVpxServiceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service1", "name", "test_load_balancer_service1"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service1", "destination_port", "89"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service1", "weight", "55"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service2", "name", "test_load_balancer_service2"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service2", "destination_port", "89"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service2", "weight", "55"),
				),
			},
		},
	})
}

func TestAccIBMLbVpxService_Basic105(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMLbVpxServiceConfig_basic105,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service3", "name", "testacc_service3"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service3", "destination_port", "89"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service4", "name", "testacc_service4"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service4", "destination_port", "89"),
				),
			},
		},
	})
}

func TestAccIBMLbVpxServiceWithTag(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMLbVpxServiceWithTag,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service1", "name", "test_load_balancer_service1"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service1", "destination_port", "89"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service1", "weight", "55"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service1", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMLbVpxServiceWithUpdatedTag,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service1", "name", "test_load_balancer_service1"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service1", "destination_port", "89"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service1", "weight", "55"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_service.testacc_service1", "tags.#", "3"),
				),
			},
		},
	})
}

var testAccCheckIBMLbVpxServiceConfig_basic = `

resource "ibm_compute_vm_instance" "vm1" {
    hostname = "vm1"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal09"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_compute_vm_instance" "vm2" {
    hostname = "vm2"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal09"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_lb_vpx" "testacc_vpx" {
    datacenter = "dal09"
    speed = 10
    version = "10.1"
    plan = "Standard"
    ip_count = 2
}

resource "ibm_lb_vpx_vip" "testacc_vip" {
    name = "test_load_balancer_vip"
    nad_controller_id = "${ibm_lb_vpx.testacc_vpx.id}"
    load_balancing_method = "lc"
    source_port = 80
    type = "HTTP"
    virtual_ip_address = "${ibm_lb_vpx.testacc_vpx.vip_pool[0]}"
}

resource "ibm_lb_vpx_service" "testacc_service1" {
  name = "test_load_balancer_service1"
  vip_id = "${ibm_lb_vpx_vip.testacc_vip.id}"
  destination_ip_address = "${ibm_compute_vm_instance.vm1.ipv4_address}"
  destination_port = 89
  weight = 55
  connection_limit = 5000
  health_check = "HTTP"
}

resource "ibm_lb_vpx_service" "testacc_service2" {
  name = "test_load_balancer_service2"
  vip_id = "${ibm_lb_vpx_vip.testacc_vip.id}"
  destination_ip_address = "${ibm_compute_vm_instance.vm2.ipv4_address}"
  destination_port = 89
  weight = 55
  connection_limit = 5000
  health_check = "HTTP"
}
`

var testAccCheckIBMLbVpxServiceConfig_basic105 = `

resource "ibm_compute_vm_instance" "vm3" {
    hostname = "vm3"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal09"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_compute_vm_instance" "vm4" {
    hostname = "vm4"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal09"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_lb_vpx" "testacc_vpx2" {
    datacenter = "dal09"
    speed = 10
    version = "10.5"
    plan = "Standard"
    ip_count = 2
}

resource "ibm_lb_vpx_vip" "testacc_vip2" {
    name = "testacc_vip2"
    nad_controller_id = "${ibm_lb_vpx.testacc_vpx2.id}"
    load_balancing_method = "lc"
    source_port = 80
    type = "HTTP"
    virtual_ip_address = "${ibm_lb_vpx.testacc_vpx2.vip_pool[0]}"
}

resource "ibm_lb_vpx_service" "testacc_service3" {
  name = "testacc_service3"
  vip_id = "${ibm_lb_vpx_vip.testacc_vip2.id}"
  destination_ip_address = "${ibm_compute_vm_instance.vm3.ipv4_address}"
  destination_port = 89
  weight = 55
  connection_limit = 5000
  health_check = "HTTP"
}

resource "ibm_lb_vpx_service" "testacc_service4" {
  name = "testacc_service4"
  vip_id = "${ibm_lb_vpx_vip.testacc_vip2.id}"
  destination_ip_address = "${ibm_compute_vm_instance.vm4.ipv4_address}"
  destination_port = 89
  weight = 55
  connection_limit = 5000
  health_check = "HTTP"
}
`
var testAccCheckIBMLbVpxServiceWithTag = `

resource "ibm_compute_vm_instance" "vm1" {
    hostname = "vm1"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal09"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_lb_vpx" "testacc_vpx" {
    datacenter = "dal09"
    speed = 10
    version = "10.1"
    plan = "Standard"
    ip_count = 2
}

resource "ibm_lb_vpx_vip" "testacc_vip" {
    name = "test_load_balancer_vip"
    nad_controller_id = "${ibm_lb_vpx.testacc_vpx.id}"
    load_balancing_method = "lc"
    source_port = 80
    type = "HTTP"
    virtual_ip_address = "${ibm_lb_vpx.testacc_vpx.vip_pool[0]}"
}

resource "ibm_lb_vpx_service" "testacc_service1" {
  name = "test_load_balancer_service1"
  vip_id = "${ibm_lb_vpx_vip.testacc_vip.id}"
  destination_ip_address = "${ibm_compute_vm_instance.vm1.ipv4_address}"
  destination_port = 89
  weight = 55
  connection_limit = 5000
  health_check = "HTTP"
  tags = ["one", "two"]
}
`
var testAccCheckIBMLbVpxServiceWithUpdatedTag = `

resource "ibm_compute_vm_instance" "vm1" {
    hostname = "vm1"
    domain = "terraformuat.ibm.com"
    os_reference_code = "DEBIAN_8_64"
    datacenter = "dal09"
    network_speed = 10
    hourly_billing = true
    private_network_only = false
    cores = 1
    memory = 1024
    disks = [25]
    local_disk = false
}

resource "ibm_lb_vpx" "testacc_vpx" {
    datacenter = "dal09"
    speed = 10
    version = "10.1"
    plan = "Standard"
    ip_count = 2
}

resource "ibm_lb_vpx_vip" "testacc_vip" {
    name = "test_load_balancer_vip"
    nad_controller_id = "${ibm_lb_vpx.testacc_vpx.id}"
    load_balancing_method = "lc"
    source_port = 80
    type = "HTTP"
    virtual_ip_address = "${ibm_lb_vpx.testacc_vpx.vip_pool[0]}"
}

resource "ibm_lb_vpx_service" "testacc_service1" {
  name = "test_load_balancer_service1"
  vip_id = "${ibm_lb_vpx_vip.testacc_vip.id}"
  destination_ip_address = "${ibm_compute_vm_instance.vm1.ipv4_address}"
  destination_port = 89
  weight = 55
  connection_limit = 5000
  health_check = "HTTP"
  tags = ["one", "two", "three"]
}
`

/*
var testAccCheckIBMLbVpxServiceConfig_basic105 = `

resource "ibm_lb_vpx_vip" "testacc_vip2" {
    name = "testacc_vip2"
    nad_controller_id = 21641
    load_balancing_method = "lc"
    source_port = 80
    type = "HTTP"
    virtual_ip_address = "169.54.227.96"
}

resource "ibm_lb_vpx_service" "testacc_service3" {
  name = "testacc_service3"
  vip_id = "${ibm_lb_vpx_vip.testacc_vip2.id}"
  destination_ip_address = "10.10.10.10"
  destination_port = 89
  weight = 55
  connection_limit = 5000
  health_check = "HTTP"
}

resource "ibm_lb_vpx_service" "testacc_service4" {
  name = "testacc_service4"
  vip_id = "${ibm_lb_vpx_vip.testacc_vip2.id}"
  destination_ip_address = "20.20.20.20"
  destination_port = 89
  weight = 55
  connection_limit = 5000
  health_check = "HTTP"
}
`
*/
