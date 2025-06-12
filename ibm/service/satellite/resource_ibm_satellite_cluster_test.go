// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccSatelliteCluster_Basic(t *testing.T) {
	var instance string
	clusterName := fmt.Sprintf("tf-satellitecluster-%d", acctest.RandIntRange(10, 100))
	locationName := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"
	operatingSystem := "REDHAT_7_64"
	zones := []string{"us-south-1", "us-south-2", "us-south-3"}
	resource_group := "default"
	region := "us-south"
	resource_prefix := "tf-satellite"
	host_provider := "ibm"
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSatelliteClusterDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckSatelliteClusterCreate(clusterName, locationName, managed_from, operatingSystem, resource_group, resource_prefix, region, publicKey, host_provider, zones),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteClusterExists("ibm_satellite_cluster.create_cluster", instance),
					resource.TestCheckResourceAttr("ibm_satellite_cluster.create_cluster", "name", clusterName),
				),
			},
		},
	})
}

func TestAccSatelliteCluster_SingleNode(t *testing.T) {
	var instance string
	clusterName := fmt.Sprintf("tf-singlenode-%d", acctest.RandIntRange(10, 100))
	locationName := fmt.Sprintf("tf-satellitelocation-singlenode-%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"
	operatingSystem := "REDHAT_7_64"
	zones := []string{"us-south-1", "us-south-2", "us-south-3"}
	resource_group := "Default"
	region := "us-south"
	resource_prefix := "tf-satellite"
	host_provider := "ibm"
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSatelliteClusterDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckSatelliteSingleNodeClusterCreate(clusterName, locationName, managed_from, operatingSystem, resource_group, resource_prefix, region, publicKey, host_provider, zones),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteClusterExists("ibm_satellite_cluster.create_cluster", instance),
					resource.TestCheckResourceAttr("ibm_satellite_cluster.create_cluster", "name", clusterName),
				),
			},
		},
	})
}

func TestAccSatelliteCluster_Entitlement(t *testing.T) {
	var instance string
	clusterName := fmt.Sprintf("tf-satellitecluster-%d", acctest.RandIntRange(10, 100))
	locationName := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	managed_from := "dal10"
	operatingSystem := "REDHAT_8_64"
	zones := []string{"us-south-1", "us-south-2", "us-south-3"}
	resource_group := "default"
	region := "us-south"
	resource_prefix := "tf-satellite"
	host_provider := "ibm"
	publicKey := acc.SatelliteSSHPubKey

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckSatelliteSSH(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSatelliteClusterDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckSatelliteEntitlementClusterCreate(clusterName, locationName, managed_from, operatingSystem, resource_group, resource_prefix, region, publicKey, host_provider, zones),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteClusterExists("ibm_satellite_cluster.create_cluster", instance),
					resource.TestCheckResourceAttr("data.ibm_satellite_cluster_worker_pool.read_default_wp", "openshift_license_source", "cloud_pak"),
				),
			},
		},
	})
}

func TestAccSatelliteCluster_Import(t *testing.T) {
	var instance string
	clusterName := fmt.Sprintf("tf-satellitecluster-%d", acctest.RandIntRange(10, 100))
	locationName := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"
	operatingSystem := "REDHAT_7_64"
	zones := []string{"us-south-1", "us-south-2", "us-south-3"}
	resource_group := "default"
	region := "us-south"
	resource_prefix := "tf-satellite"
	host_provider := "ibm"
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckSatelliteClusterDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckSatelliteClusterCreate(clusterName, locationName, managed_from, operatingSystem, resource_group, resource_prefix, region, publicKey, host_provider, zones),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteClusterExists("ibm_satellite_cluster.create_cluster", instance),
					resource.TestCheckResourceAttr("ibm_satellite_cluster.create_cluster", "name", clusterName),
				),
			},
			{
				ResourceName:      "ibm_satellite_cluster.create_cluster",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"enable_config_admin", "wait_for_worker_update", "location"},
			},
		},
	})
}

func testAccCheckSatelliteClusterExists(n string, instance string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		ID := rs.Primary.ID
		satClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatelliteClientSession()
		if err != nil {
			return err
		}

		getSatClusterOptions := &kubernetesserviceapiv1.GetClusterOptions{
			Cluster: &ID,
		}

		cluster, resp, err := satClient.GetCluster(getSatClusterOptions)
		if err != nil {
			if resp != nil && resp.StatusCode == 404 {
				return nil
			}
			return fmt.Errorf("[ERROR] Error retrieving satellite cluster: %s\n Response code is: %+v", err, resp)
		}

		instance = *cluster.ID

		return nil
	}
}

func testAccCheckSatelliteClusterDestroy(s *terraform.State) error {
	satClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatelliteClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_satellite_cluster" {
			continue
		}

		ID := rs.Primary.ID
		getSatClusterOptions := &kubernetesserviceapiv1.GetClusterOptions{
			Cluster: &ID,
		}

		_, _, err = satClient.GetCluster(getSatClusterOptions)
		if err == nil {
			return fmt.Errorf("Satellite Cluster still exists: %s", rs.Primary.ID)
		}

	}
	return nil
}

func testAccCheckSatelliteClusterCreate(clusterName, locationName, managed_from, operatingSystem, resource_group, resource_prefix, region, publicKey, host_provider string, zones []string) string {
	return fmt.Sprintf(`

	variable "location_zones" {
		description = "Allocate your hosts across these three zones"
		type        = list(string)
		default     = ["us-south-1", "us-south-2", "us-south-3"]
	}

	data "ibm_is_image" "rhel7" {
		name = "ibm-redhat-7-9-minimal-amd64-7"
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
		operating_system       = "%s"
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

`, locationName, managed_from, resource_group, resource_prefix, resource_prefix, region, resource_prefix, publicKey, resource_prefix, region, resource_prefix, host_provider, clusterName, operatingSystem)
}

func testAccCheckSatelliteSingleNodeClusterCreate(clusterName, locationName, managed_from, operatingSystem, resource_group, resource_prefix, region, publicKey, host_provider string, zones []string) string {
	return fmt.Sprintf(`

	variable "location_zones" {
		description = "Allocate your hosts across these three zones"
		type        = list(string)
		default     = ["us-south-1", "us-south-2", "us-south-3"]
	}

	data "ibm_is_image" "rhel8" {
		name = "ibm-redhat-8-6-minimal-amd64-3"
	}

	resource "ibm_satellite_location" "location" {
		location       = "%s"
		managed_from   = "%s"
		coreos_enabled = true
		zones		   = var.location_zones
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
		host_provider = "%s"
	  }

	  resource "ibm_satellite_cluster" "create_cluster" {
		name                   = "%s"  
		location               = ibm_satellite_host.assign_host.0.location
		enable_config_admin    = true
		kube_version           = "4.11_openshift"
		infrastructure_topology = "single-replica"
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

`, locationName, managed_from, resource_group, resource_prefix, resource_prefix, region, resource_prefix, publicKey, resource_prefix, region, resource_prefix, host_provider, clusterName)
}

func testAccCheckSatelliteEntitlementClusterCreate(clusterName, locationName, managed_from, operatingSystem, resource_group, resource_prefix, region, publicKey, host_provider string, zones []string) string {
	return fmt.Sprintf(`

	variable "location_zones" {
		description = "Allocate your hosts across these three zones"
		type        = list(string)
		default     = ["us-south-1", "us-south-2", "us-south-3"]
	}

	data "ibm_is_image" "rhel8" {
		name = "ibm-redhat-8-8-minimal-amd64-2"
	}

	resource "ibm_satellite_location" "location" {
		location      = "%s"
		managed_from  = "%s"
		zones		  = var.location_zones
		coreos_enabled = true
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
		host_provider = "%s"
	  }

	  resource "ibm_satellite_cluster" "create_cluster" {
		name                   = "%s"  
		location               = ibm_satellite_host.assign_host.0.location
		enable_config_admin    = true
		kube_version           = "4.13_openshift"
		operating_system       = "%s"
		entitlement            = "cloud_pak"
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

	data "ibm_satellite_cluster_worker_pool" "read_default_wp" {
		name    = "default"
		cluster = ibm_satellite_cluster.create_cluster.id
	}
`, locationName, managed_from, resource_group, resource_prefix, resource_prefix, region, resource_prefix, publicKey, resource_prefix, region, resource_prefix, host_provider, clusterName, operatingSystem)
}
