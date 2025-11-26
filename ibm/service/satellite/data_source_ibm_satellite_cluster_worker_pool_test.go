// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMSatelliteClusterWorkerPoolDataSourceBasic(t *testing.T) {
	clusterName := fmt.Sprintf("tf-satellitecluster-%d", acctest.RandIntRange(10, 100))
	locationName := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	workerPoolName := fmt.Sprintf("tf-cluster-wp-%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"
	zones := []string{"us-east-1", "us-east-2", "us-east-3"}
	resource_group := "default"
	region := "us-east"
	resource_prefix := "tf-satellite"
	host_provider := "ibm"
	operatingSystem := "REDHAT_7_64"
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSatelliteClusterorkerPoolDataSourceConfig(workerPoolName, clusterName, locationName, managed_from, operatingSystem, resource_group, resource_prefix, region, publicKey, host_provider, zones),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_satellite_cluster_worker_pool.read_wp", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_cluster_worker_pool.read_wp", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_cluster_worker_pool.read_wp", "cluster"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_cluster_worker_pool.read_wp", "operating_system"),
				),
			},
		},
	})
}

func testAccCheckIBMSatelliteClusterorkerPoolDataSourceConfig(workerPoolName, clusterName, locationName, managed_from, operatingSystem, resource_group, resource_prefix, region, publicKey, host_provider string, zones []string) string {
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

	data "ibm_is_image" "rhel7" {
		name = "ibm-redhat-7-9-minimal-amd64-7"
	}
	  
	resource "ibm_is_instance" "satellite_instance" {
		count          = 3

		name           = "%s-instance-${count.index}"
		vpc            = ibm_is_vpc.satellite_vpc.id
		zone           = "%s-${count.index + 1}"
		image          = data.ibm_is_image.rhel7.id
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

	resource "ibm_satellite_cluster" "create_cluster" {
		name                   = "%s"  
		location               = ibm_satellite_host.assign_host.0.location
		enable_config_admin    = true
		kube_version           = "4.9_openshift"
		wait_for_worker_update = true
		dynamic "zones" {
			for_each = var.location_zones
			content {
				id	= zones.value
			}
		}
		default_worker_pool_labels = {
			"test"  = "test-pool1" 
			"test1" = "test-pool2"
		}
	}

	resource "ibm_satellite_cluster_worker_pool" "create_wp" {
		name               = "%s"  
		cluster            = ibm_satellite_cluster.create_cluster.id
		operating_system   = "%s"
		worker_count       = 1   
		dynamic "zones" {
			for_each = var.location_zones
			content {
				id	= zones.value
			}
		}
		host_labels       	= ["env:dev"]
		worker_pool_labels = {
			"test"  = "test-pool1" 
			"test1" = "test-pool2"
		}
	}

	data "ibm_satellite_cluster_worker_pool" "read_wp" {
		name    = ibm_satellite_cluster_worker_pool.create_wp.name
		cluster = ibm_satellite_cluster.create_cluster.id
	}	  

`, locationName, managed_from, resource_group, resource_prefix, resource_prefix, region, resource_prefix, publicKey, resource_prefix, region, resource_prefix, host_provider, clusterName, workerPoolName, operatingSystem)
}
