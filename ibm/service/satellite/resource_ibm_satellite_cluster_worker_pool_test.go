// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSatelliteClusterWorkerPool_Basic(t *testing.T) {
	var instance string
	clusterName := fmt.Sprintf("tf-satellitecluster-%d", acctest.RandIntRange(10, 100))
	locationName := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	operatingSystem := "REDHAT_7_64"
	workerPoolName := fmt.Sprintf("tf-wp-%d", acctest.RandIntRange(10, 100))
	resource_prefix := "tf-satellite"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSatelliteClusterWorkerPoolDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckSatelliteClusterWorkerPoolCreate(clusterName, locationName, operatingSystem, workerPoolName, resource_prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteClusterWorkerPoolExists("ibm_satellite_cluster_worker_pool.create_wp", instance),
					resource.TestCheckResourceAttr("ibm_satellite_cluster.create_cluster", "name", clusterName),
					resource.TestCheckResourceAttr("ibm_satellite_cluster_worker_pool.create_wp", "name", workerPoolName),
					resource.TestCheckResourceAttr("ibm_satellite_cluster_worker_pool.create_wp", "operating_system", operatingSystem),
				),
			},
		},
	})
}

func TestAccSatelliteClusterWorkerPool_Entitlement(t *testing.T) {
	var instance string
	clusterName := fmt.Sprintf("tf-satellitecluster-%d", acctest.RandIntRange(10, 100))
	locationName := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	operatingSystem := "REDHAT_8_64"
	workerPoolName := fmt.Sprintf("tf-wp-%d", acctest.RandIntRange(10, 100))
	resource_prefix := "tf-satellite"
	publicKey := acc.SatelliteSSHPubKey

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckSatelliteSSH(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSatelliteClusterWorkerPoolDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckSatelliteClusterWorkerPoolCreateEntitlement(clusterName, locationName, operatingSystem, workerPoolName, resource_prefix, publicKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteClusterWorkerPoolExists("ibm_satellite_cluster_worker_pool.create_wp", instance),
					resource.TestCheckResourceAttr("ibm_satellite_cluster.create_cluster", "name", clusterName),
					resource.TestCheckResourceAttr("ibm_satellite_cluster_worker_pool.create_wp", "name", workerPoolName),
					resource.TestCheckResourceAttr("ibm_satellite_cluster_worker_pool.create_wp", "operating_system", operatingSystem),
					resource.TestCheckResourceAttr("data.ibm_satellite_cluster_worker_pool.read_created_wp", "openshift_license_source", "cloud_pak"),
				),
			},
		},
	})
}

func TestAccSatelliteClusterWorkerPool_Import(t *testing.T) {
	var instance string
	clusterName := fmt.Sprintf("tf-satellitecluster-%d", acctest.RandIntRange(10, 100))
	locationName := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	workerPoolName := fmt.Sprintf("tf-wp-%d", acctest.RandIntRange(10, 100))
	operatingSystem := "REDHAT_7_64"
	resource_prefix := "tf-satellite"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSatelliteClusterWorkerPoolDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckSatelliteClusterWorkerPoolCreate(clusterName, locationName, operatingSystem, workerPoolName, resource_prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteClusterWorkerPoolExists("ibm_satellite_cluster_worker_pool.create_wp", instance),
					resource.TestCheckResourceAttr("ibm_satellite_cluster_worker_pool.create_wp", "name", workerPoolName),
				),
			},
			{
				ResourceName:      "ibm_satellite_cluster_worker_pool.create_wp",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSatelliteClusterWorkerPoolExists(n string, instance string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		clusterID := parts[0]
		workerPoolID := parts[1]

		satClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatelliteClientSession()
		if err != nil {
			return err
		}

		getSatWorkerPoolOptions := &kubernetesserviceapiv1.GetWorkerPoolOptions{
			Cluster:    &clusterID,
			Workerpool: &workerPoolID,
		}

		wp, resp, err := satClient.GetWorkerPool(getSatWorkerPoolOptions)
		if err != nil {
			if resp != nil && resp.StatusCode == 404 {
				return nil
			}
			return fmt.Errorf("[ERROR] Error retrieving worker pool : %s\n Response code is: %+v", err, resp)
		}

		instance = *wp.ID

		return nil
	}
}

func testAccCheckSatelliteClusterWorkerPoolDestroy(s *terraform.State) error {
	satClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatelliteClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_satellite_cluster_worker_pool" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		clusterID := parts[0]
		workerPoolID := parts[1]

		getSatWorkerPoolOptions := &kubernetesserviceapiv1.GetWorkerPoolOptions{
			Cluster:    &clusterID,
			Workerpool: &workerPoolID,
		}

		_, _, err = satClient.GetWorkerPool(getSatWorkerPoolOptions)
		if err == nil {
			return fmt.Errorf("Worker pool still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckSatelliteClusterWorkerPoolCreate(clusterName, locationName, operatingSystem, workerPoolName, resource_prefix string) string {
	return fmt.Sprintf(`

	variable "location_zones" {
		description = "Allocate your hosts across these three zones"
		type        = list(string)
		default     = ["us-south-1", "us-south-2", "us-south-3"]
	}

	resource "ibm_satellite_location" "location" {
		location      = "%s"
		managed_from  = "wdc04"
		zones		  = var.location_zones
	}

	data "ibm_is_image" "rhel7" {
		name = "ibm-redhat-7-9-minimal-amd64-7"
	}

	data "ibm_satellite_attach_host_script" "script" {
		location          = ibm_satellite_location.location.id
		labels            = ["env:prod"]
		host_provider     = "ibm"
	}

	data "ibm_resource_group" "resource_group" {
		is_default = true
	}
	  
	resource "ibm_is_vpc" "satellite_vpc" {
		name = "%s-vpc-1"
	}
	  
	resource "ibm_is_subnet" "satellite_subnet" {
		count                    = 3

		name                     = "%s-subnet-${count.index}"
		vpc                      = ibm_is_vpc.satellite_vpc.id
		total_ipv4_address_count = 256
		zone                     = "us-south-${count.index + 1}"
	}
	  
	resource "ibm_is_ssh_key" "satellite_ssh" {	  
		name        = "%s-ibm-ssh"
		public_key  = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
	}
	  
	resource "ibm_is_instance" "satellite_instance" {
		count          = 3

		name           = "%s-instance-${count.index}"
		vpc            = ibm_is_vpc.satellite_vpc.id
		zone           = "us-south-${count.index + 1}"
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
		host_provider = "ibm"
	}

	resource "ibm_satellite_cluster" "create_cluster" {
		name                   = "%s"  
		location               = ibm_satellite_location.location.id
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
		worker_count       = 1   
		host_labels        = ["env:dev"]
		operating_system   = "%s"
		dynamic "zones" {
			for_each = var.location_zones
			content {
				id	= zones.value
			}
		}
		worker_pool_labels = {
			"test"  = "test-pool1" 
			"test1" = "test-pool2"
		}
	}

`, locationName, resource_prefix, resource_prefix, resource_prefix, resource_prefix, resource_prefix, clusterName, workerPoolName, operatingSystem)
}

func testAccCheckSatelliteClusterWorkerPoolCreateEntitlement(clusterName, locationName, operatingSystem, workerPoolName, resource_prefix, publicKey string) string {
	return fmt.Sprintf(`

	variable "location_zones" {
		description = "Allocate your hosts across these three zones"
		type        = list(string)
		default     = ["us-south-1", "us-south-2", "us-south-3"]
	}

	resource "ibm_satellite_location" "location" {
		location      = "%s"
		managed_from  = "dal10"
		zones		  = var.location_zones
		coreos_enabled = true
	}

	data "ibm_is_image" "rhel8" {
		name = "ibm-redhat-8-8-minimal-amd64-2"
	}

	data "ibm_satellite_attach_host_script" "script" {
		location          = ibm_satellite_location.location.id
		labels            = ["env:prod"]
		host_provider     = "ibm"
	}

	data "ibm_resource_group" "resource_group" {
		is_default = true
	}
	  
	resource "ibm_is_vpc" "satellite_vpc" {
		name = "%s-vpc-1"
	}
	  
	resource "ibm_is_subnet" "satellite_subnet" {
		count                    = 3

		name                     = "%s-subnet-${count.index}"
		vpc                      = ibm_is_vpc.satellite_vpc.id
		total_ipv4_address_count = 256
		zone                     = "us-south-${count.index + 1}"
	}
	  
	resource "ibm_is_ssh_key" "satellite_ssh" {	  
		name        = "%s-ibm-ssh"
		public_key  = "%s"
	}
	  
	resource "ibm_is_instance" "satellite_instance" {
		count          = 3

		name           = "%s-instance-${count.index}"
		vpc            = ibm_is_vpc.satellite_vpc.id
		zone           = "us-south-${count.index + 1}"
		image          = data.ibm_is_image.rhel8.id
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
		host_provider = "ibm"
	}

	resource "ibm_satellite_cluster" "create_cluster" {
		name                   = "%s"  
		location               = ibm_satellite_location.location.id
		enable_config_admin    = true
		kube_version           = "4.13_openshift"
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
		worker_count       = 1   
		host_labels        = ["env:dev"]
		operating_system   = "%s"
		entitlement        = "cloud_pak"
		dynamic "zones" {
			for_each = var.location_zones
			content {
				id	= zones.value
			}
		}
		worker_pool_labels = {
			"test"  = "test-pool1" 
			"test1" = "test-pool2"
		}
	}

	data "ibm_satellite_cluster_worker_pool" "read_created_wp" {
		name    = ibm_satellite_cluster_worker_pool.create_wp.name
		cluster = ibm_satellite_cluster.create_cluster.id
	}

`, locationName, resource_prefix, resource_prefix, resource_prefix, publicKey, resource_prefix, resource_prefix, clusterName, workerPoolName, operatingSystem)
}
