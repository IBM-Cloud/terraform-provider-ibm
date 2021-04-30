// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.ibm.com/ibmcloud/kubernetesservice-go-sdk/kubernetesserviceapiv1"
)

func TestAccFunctionSatelliteHost_Basic(t *testing.T) {
	name := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"
	zones := []string{"us-east-1", "us-east-2", "us-east-3"}
	resource_group := "default"
	region := "us-east"
	resource_prefix := "tf-satellite"
	host_provider := "ibm"
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSatelliteHostDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckSatelliteHostCreate(name, managed_from, resource_group, resource_prefix, region, publicKey, host_provider, zones),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteHostExists("ibm_satellite_host.assign_host.0"),
					resource.TestCheckResourceAttr("ibm_satellite_host.assign_host.0", "host_provider", host_provider),
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

		satClient, err := testAccProvider.Meta().(ClientSession).SatelliteClientSession()
		if err != nil {
			return err
		}
		parts, err := idParts(rs.Primary.ID)
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
			return fmt.Errorf("Error retrieving satellite hosts: %s\n Response code is: %+v", err, resp)
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
	satClient, err := testAccProvider.Meta().(ClientSession).SatelliteClientSession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_satellite_host" {
			continue
		}

		parts, err := idParts(rs.Primary.ID)
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

func testAccCheckSatelliteHostCreate(name, managed_from, resource_group, resource_prefix, region, publicKey, host_provider string, zones []string) string {
	return fmt.Sprintf(`

	provider "ibm" {
		region = "us-east"
	}

	variable "location_zones" {
		description = "Allocate your hosts across these three zones"
		type        = list(string)
		default     = ["us-east-1", "us-east-2", "us-east-3"]
	}

	resource "ibm_satellite_location" "location" {
		location      = "%s"
		managed_from  = "%s"
		zones		  = var.location_zones
	}

	data "ibm_satellite_attach_host_script" "script" {
		location          = ibm_satellite_location.location.id
		labels            = ["env:prod"]
		host_provider     = "ibm"
	}

	data "ibm_resource_group" "resource_group" {
		name = "%s"
	}
	  
	resource "ibm_is_vpc" "satellite_vpc" {
		name = "%s-vpc-1"
	}
	  
	resource "ibm_is_subnet" "satellite_subnet" {
		count                    = 3

		name                     = "%s-subnet-${count.index}"
		vpc                      = ibm_is_vpc.satellite_vpc.id
		total_ipv4_address_count = 256
		zone                     = "%s-${count.index + 1}"
	  }
	  
	  resource "ibm_is_ssh_key" "satellite_ssh" {	  
		name        = "%s-ibm-ssh"
		public_key  = "%s"
	  }
	  
	  resource "ibm_is_instance" "satellite_instance" {
		count          = 3

		name           = "%s-instance-${count.index}"
		vpc            = ibm_is_vpc.satellite_vpc.id
		zone           = "%s-${count.index + 1}"
		image          = "r014-931515d2-fcc3-11e9-896d-3baa2797200f"
		profile        = "mx2-8x64"
		keys           = [ibm_is_ssh_key.satellite_ssh.id]
		resource_group = data.ibm_resource_group.resource_group.id
		user_data      = data.ibm_satellite_attach_host_script.script.host_script
		
		primary_network_interface {
		  subnet = ibm_is_subnet.satellite_subnet[count.index].id
		}
	  }
	  
	  resource "ibm_is_floating_ip" "satellite_ip" {
		count  = 3

		name   = "%s-fip-${count.index}"
		target = ibm_is_instance.satellite_instance[count.index].primary_network_interface[0].id
	  }
	  
	  resource "ibm_satellite_host" "assign_host" {
		count  = 3
	  
		location      = ibm_satellite_location.location.id
		host_id       = element(ibm_is_instance.satellite_instance[*].name, count.index)
		labels        = ["env:prod"]
		zone          = element(var.location_zones, count.index)
		host_provider = "%s"
	  }

`, name, managed_from, resource_group, resource_prefix, resource_prefix, region, resource_prefix, publicKey, resource_prefix, region, resource_prefix, host_provider)
}
