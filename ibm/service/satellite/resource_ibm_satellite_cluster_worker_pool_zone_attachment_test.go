// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	"github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIbmSatelliteClusterWorkerPoolZoneAttachmentBasic(t *testing.T) {
	var conf kubernetesserviceapiv1.GetWorkerPoolResponse
	clusterName := fmt.Sprintf("tf-satellitecluster-%d", acctest.RandIntRange(10, 100))
	locationName := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	zone := fmt.Sprintf("tf-zone-%d", acctest.RandIntRange(10, 100))
	resource_prefix := fmt.Sprintf("tf-satellite-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSatelliteClusterWorkerPoolZoneAttachmentConfigBasic(locationName, clusterName, resource_prefix, zone),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSatelliteClusterWorkerPoolZoneAttachmentExists("ibm_satellite_cluster_worker_pool_zone_attachment.satellite_cluster_worker_pool_zone_attachment", conf),
				),
			},
		},
	})
}

func TestAccIbmSatelliteClusterWorkerPoolZoneAttachmentAllArgs(t *testing.T) {
	var conf kubernetesserviceapiv1.GetWorkerPoolResponse
	clusterName := fmt.Sprintf("tf-satellitecluster-%d", acctest.RandIntRange(10, 100))
	locationName := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	zone := fmt.Sprintf("tf-zone-%d", acctest.RandIntRange(10, 100))
	resource_prefix := fmt.Sprintf("tf-satellite-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSatelliteClusterWorkerPoolZoneAttachmentConfig(locationName, clusterName, resource_prefix, zone),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSatelliteClusterWorkerPoolZoneAttachmentExists("ibm_satellite_cluster_worker_pool_zone_attachment.satellite_cluster_worker_pool_zone_attachment", conf),
					resource.TestCheckResourceAttr("ibm_satellite_cluster_worker_pool_zone_attachment.satellite_cluster_worker_pool_zone_attachment", "worker_pool", "default"),
					resource.TestCheckResourceAttr("ibm_satellite_cluster_worker_pool_zone_attachment.satellite_cluster_worker_pool_zone_attachment", "zone", zone),
				),
			},
			{
				ResourceName:      "ibm_satellite_cluster_worker_pool_zone_attachment.satellite_cluster_worker_pool_zone_attachment",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSatelliteClusterWorkerPoolZoneAttachmentConfigBasic(location, cluster, resource_prefix, zone string) string {
	return fmt.Sprintf(`

	provider "ibm" {
		region = "us-south"
	}

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
		name = "ibm-redhat-7-9-minimal-amd64-4"
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
		kube_version           = "4.6_openshift"
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

	resource "ibm_satellite_cluster_worker_pool_zone_attachment" "satellite_cluster_worker_pool_zone_attachment" {
		cluster = ibm_satellite_cluster.create_cluster.id
		worker_pool = "default"
		zone = "%s"
	}
	`, location, resource_prefix, resource_prefix, resource_prefix, resource_prefix, resource_prefix, cluster, zone)
}

func testAccCheckIbmSatelliteClusterWorkerPoolZoneAttachmentConfig(location, cluster, resource_prefix, zone string) string {
	return fmt.Sprintf(`

	provider "ibm" {
		region = "us-south"
	}

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
		name = "ibm-redhat-7-9-minimal-amd64-4"
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
		kube_version           = "4.6_openshift"
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

	resource "ibm_satellite_cluster_worker_pool_zone_attachment" "satellite_cluster_worker_pool_zone_attachment" {
		cluster = ibm_satellite_cluster.create_cluster.id
		worker_pool = "default"
		zone = "%s"
	}
	`, location, resource_prefix, resource_prefix, resource_prefix, resource_prefix, resource_prefix, cluster, zone)
}

func testAccCheckIbmSatelliteClusterWorkerPoolZoneAttachmentExists(n string, obj kubernetesserviceapiv1.GetWorkerPoolResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		satClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatelliteClientSession()
		if err != nil {
			return err
		}

		getWorkerPoolOptions := &kubernetesserviceapiv1.GetWorkerPoolOptions{}
		getWorkerPoolOptions.SetCluster(parts[0])
		getWorkerPoolOptions.SetWorkerpool(parts[1])

		getWorkerPoolResponse, _, err := satClient.GetWorkerPool(getWorkerPoolOptions)
		if err != nil {
			return err
		}

		if getWorkerPoolResponse != nil && getWorkerPoolResponse.Zones != nil {
			for _, zoneId := range getWorkerPoolResponse.Zones {
				if parts[2] == *zoneId.ID {
					obj = *getWorkerPoolResponse
				}
			}
		}

		return nil
	}
}
