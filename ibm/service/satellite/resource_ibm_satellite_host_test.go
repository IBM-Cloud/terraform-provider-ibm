// Copyright IBM Corp. 2017, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFunctionSatelliteHost_Basic(t *testing.T) {
	name := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	resource_prefix := "tf-satellite"
	rhel_image_name := "ibm-redhat-8-8-minimal-amd64-3"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSatelliteHostDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckSatelliteHostCreate(name, resource_prefix, rhel_image_name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteHostExists("ibm_satellite_host.assign_host.0"),
					resource.TestCheckResourceAttr("ibm_satellite_host.assign_host.0", "host_provider", "ibm"),
				),
			},
		},
	})
}

func testAccCheckSatelliteHostExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		satClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatelliteClientSession()
		if err != nil {
			return err
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		ID := parts[0]
		hostID := parts[1]
		getSatOptions := &kubernetesserviceapiv1.GetSatelliteHostsOptions{
			Controller: &ID,
		}

		hostList, resp, err := satClient.GetSatelliteHosts(getSatOptions)
		if err != nil {
			if resp != nil && resp.StatusCode == 404 {
				return nil
			}
			return fmt.Errorf("[ERROR] Error retrieving satellite hosts: %s\n Response code is: %+v", err, resp)
		}

		isExist := false
		for _, h := range hostList {
			if hostID == *h.Name {
				isExist = true
			}
		}

		if isExist == false {
			return fmt.Errorf("Record not found")
		}

		return nil
	}
}

func testAccCheckSatelliteHostDestroy(s *terraform.State) error {
	satClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatelliteClientSession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_satellite_host" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		ID := parts[0]
		hostID := parts[1]

		removeSatHostOptions := &kubernetesserviceapiv1.RemoveSatelliteHostOptions{}
		removeSatHostOptions.Controller = &ID
		removeSatHostOptions.HostID = &hostID
		_, err = satClient.RemoveSatelliteHost(removeSatHostOptions)
		if err == nil {
			return fmt.Errorf("Satellite host still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckSatelliteHostCreate(name, resource_prefix, rhel_image_name string) string {
	return fmt.Sprintf(`
	variable "location_zones" {
		description = "Allocate your hosts across these three zones"
		type        = list(string)
		default     = ["location-zone-1", "location-zone-2", "location-zone-3"]
	  }
	  
	  resource "ibm_satellite_location" "location" {
		location     = "%[1]s"
		managed_from = "dal"
		zones        = var.location_zones
	  }
	  
	  data "ibm_satellite_attach_host_script" "script" {
		location      = ibm_satellite_location.location.id
		host_provider = "ibm"
	  }
	  
	  data "ibm_resource_group" "resource_group" {
		is_default = true
	  }
	  
	  resource "ibm_is_vpc" "satellite_vpc" {
		name                        = "%[2]s-vpc-1"
		resource_group              = data.ibm_resource_group.resource_group.id
		default_security_group_name = "%[2]s-default-sg"
		default_network_acl_name    = "%[2]s-default-acl"
		default_routing_table_name  = "%[2]s-default-rt"
	  }
	  
	  data "ibm_is_security_group" "default_group" {
		name = ibm_is_vpc.satellite_vpc.default_security_group_name
	  }
	  
	  resource "ibm_is_security_group_rule" "ssh_rule" {
		group     = data.ibm_is_security_group.default_group.id
		direction = "inbound"
		remote    = "0.0.0.0/0"
		tcp {
		  port_min = 22
		  port_max = 22
		}
	  }
	  
	  resource "ibm_is_security_group_rule" "ping_rule" {
		group     = data.ibm_is_security_group.default_group.id
		direction = "inbound"
		remote    = "0.0.0.0/0"
		icmp {
		  type = 8
		  code = 0
		}
	  }
	  
	  resource "ibm_is_subnet" "satellite_subnet" {
		count = 3
	  
		name                     = "%[2]s-subnet-${count.index}"
		vpc                      = ibm_is_vpc.satellite_vpc.id
		total_ipv4_address_count = 256
		zone                     = "us-south-${count.index + 1}"
	  }
	  
	  data "ibm_is_image" "rhel8" {
		name = "%[3]s"
	  }
	  
	  resource "ibm_is_instance" "satellite_instance" {
		count = 3
	  
		name           = "%[2]s-instance-${count.index}"
		vpc            = ibm_is_vpc.satellite_vpc.id
		zone           = "us-south-${count.index + 1}"
		image          = data.ibm_is_image.rhel8.id
		profile        = "mx2-8x64"
		keys           = []
		resource_group = data.ibm_resource_group.resource_group.id
		user_data      = data.ibm_satellite_attach_host_script.script.host_script
	  
		primary_network_interface {
		  name   = "eth0"
		  subnet = ibm_is_subnet.satellite_subnet[count.index].id
		}
	  }
	  
	  resource "ibm_is_floating_ip" "satellite_ip" {
		count = 3
	  
		name   = "%[2]s-fip-${count.index}"
		target = ibm_is_instance.satellite_instance[count.index].primary_network_interface[0].id
	  }
	  
	  resource "ibm_satellite_host" "assign_host" {
		count = 3
	  
		location      = ibm_satellite_location.location.id
		host_id       = element(ibm_is_instance.satellite_instance[*].name, count.index)
		zone          = element(var.location_zones, count.index)
		host_provider = "ibm"
	  }
`, name, resource_prefix, rhel_image_name)
}
