// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/power"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
)

func TestAccIBMPIInstanceBasic(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceConfig(name, power.OK),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceConfig(name, instanceHealthStatus string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_key" "key" {
		pi_cloud_instance_id = "%[1]s"
		pi_key_name          = "%[2]s"
		pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
	  }
	  data "ibm_pi_image" "power_image" {
		pi_cloud_instance_id = "%[1]s"
		pi_image_name        = "%[3]s"
	  }
	  data "ibm_pi_network" "power_networks" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[4]s"
	  }
	  resource "ibm_pi_volume" "power_volume" {
		pi_cloud_instance_id = "%[1]s"
		pi_volume_name       = "%[2]s"
		pi_volume_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_volume_shareable  = true
		pi_volume_size       = 20
		pi_volume_type       = "%[6]s"
	  }
	  resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id  = "%[1]s"
		pi_health_status      = "%[5]s"
		pi_image_id           = data.ibm_pi_image.power_image.id
		pi_instance_name      = "%[2]s"
		pi_key_pair_name      = ibm_pi_key.key.name
		pi_memory             = "2"
		pi_proc_type          = "shared"
		pi_processors         = "0.25"
		pi_storage_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_storage_type       = "%[6]s"
		pi_sys_type           = "s922"
		pi_volume_ids         = [ibm_pi_volume.power_volume.volume_id]
		pi_network {
			network_id = data.ibm_pi_network.power_networks.id
		}
	  }
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, instanceHealthStatus, acc.PiStorageType)
}

func TestAccIBMPIInstanceStorageConnection(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceStorageConnectionConfig(name, power.OK),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(instanceRes, "pi_storage_connection", acc.Pi_storage_connection),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceStorageConnectionConfig(name, instanceHealthStatus string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_volume" "power_volume" {
		pi_cloud_instance_id  = "%[1]s"
		pi_volume_size        = 1
		pi_volume_name        = "%[2]s"
		pi_volume_type        = "tier3"
	  }
	resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id = "%[1]s"
		pi_memory            = "2"
		pi_processors        = "1"
		pi_instance_name     = "%[2]s"
		pi_proc_type         = "shared"
		pi_image_id          = "%[3]s"
		pi_sys_type          = "s922"
		pi_network {
		  network_id = "%[4]s"
		}
		pi_storage_connection = "%[5]s"
		pi_health_status      = "%[6]s"
		pi_volume_ids         = [ibm_pi_volume.power_volume.volume_id]
	  }
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, acc.Pi_storage_connection, instanceHealthStatus)
}
func testAccCheckIBMPIInstanceNetworkSecurityGroupConfig(name, instanceHealthStatus string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_volume" "power_volume" {
		pi_cloud_instance_id  		 = "%[1]s"
		pi_volume_size        		 = 1
		pi_volume_name       		 = "%[2]s"
		pi_volume_type          	 = "tier3"
	  }
	resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id 		 = "%[1]s"
		pi_memory            		 = "2"
		pi_processors        		 = "0.25"
		pi_instance_name     		 = "%[2]s"
		pi_proc_type         		 = "shared"
		pi_image_id          		 = "%[3]s"
		pi_sys_type          		 = "s922"
		pi_network {
		  network_id                 = "%[4]s"
		  network_security_group_ids = ["%[6]s"]
		}
		pi_health_status             = "%[5]s"
		pi_volume_ids                = [ibm_pi_volume.power_volume.volume_id]
	  }
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, instanceHealthStatus, acc.Pi_network_security_group_id)
}
func TestAccIBMPIInstanceNetworkSecurityGroup(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceNetworkSecurityGroupConfig(name, power.OK),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(instanceRes, "pi_network.0.network_security_group_ids.0", acc.Pi_network_security_group_id),
				),
			},
		},
	})
}

func TestAccIBMPIInstanceDeploymentTarget(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceDeplomentTargetConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceDeplomentTargetConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id = "%[1]s"
		pi_memory            = "2"
		pi_processors        = "1"
		pi_instance_name     = "%[2]s"
		pi_proc_type         = "shared"
		pi_image_id          = "%[3]s"
		pi_sys_type          = "s922"
		pi_network {
		  network_id = "%[4]s"
		}
		pi_deployment_target {
			id = "308"
			type = "host"
		}
	  }
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name)
}

func TestAccIBMPIInstanceDeploymentType(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceDeploymentTypeConfig(name, power.OK, "EPIC", "e980"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
				),
			},
		},
	})
}

func TestAccIBMPIInstanceDeploymentTypeNoStorage(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceDeploymentTypeConfig(name, power.OK, "VMNoStorage", "s922"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceDeploymentTypeConfig(name, instanceHealthStatus, epic, systype string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_key" "key" {
		pi_cloud_instance_id = "%[1]s"
		pi_key_name          = "%[2]s"
		pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
	  }
	  data "ibm_pi_image" "power_image" {
		pi_cloud_instance_id = "%[1]s"
		pi_image_name        = "%[3]s"
	  }
	  data "ibm_pi_network" "power_networks" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[4]s"
	  }
	  resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id  = "%[1]s"
		pi_deployment_type          = "%[6]s"
		pi_health_status      = "%[5]s"
		pi_image_id           = data.ibm_pi_image.power_image.id
		pi_instance_name      = "%[2]s"
		pi_key_pair_name      = ibm_pi_key.key.name
		pi_memory             = "2"
		pi_proc_type          = "dedicated"
		pi_processors         = "1"
		pi_storage_type 	  = "%[8]s"
		pi_sys_type           = "%[7]s"
		pi_network {
			network_id = data.ibm_pi_network.power_networks.id
		}
	  }
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, instanceHealthStatus, epic, systype, acc.PiStorageType)
}

func TestAccIBMPIInstanceIBMiLicense(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceIBMiLicense(name, power.OK, true, 2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(instanceRes, "status", strings.ToUpper(power.State_Active)),
					resource.TestCheckResourceAttr(instanceRes, "pi_ibmi_css", "true"),
					resource.TestCheckResourceAttr(instanceRes, "ibmi_rds", "true"),
					resource.TestCheckResourceAttr(instanceRes, "pi_ibmi_rds_users", "2"),
				),
			},
			{
				Config: testAccCheckIBMPIInstanceIBMiLicense(name, power.OK, false, 0),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					testAccCheckIBMPIInstanceStatus(instanceRes, strings.ToUpper(power.State_Active)),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(instanceRes, "pi_ibmi_css", "false"),
					resource.TestCheckResourceAttr(instanceRes, "ibmi_rds", "false"),
					resource.TestCheckResourceAttr(instanceRes, "pi_ibmi_rds_users", "0"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceIBMiLicense(name, instanceHealthStatus string, IBMiCSS bool, IBMiRDSUsers int) string {
	return fmt.Sprintf(`
		  data "ibm_pi_image" "power_image" {
			pi_cloud_instance_id = "%[1]s"
			pi_image_name        = "%[3]s"
		  }
		  data "ibm_pi_network" "power_networks" {
			pi_cloud_instance_id = "%[1]s"
			pi_network_name      = "%[4]s"
		  }
		  resource "ibm_pi_volume" "power_volume" {
			pi_cloud_instance_id = "%[1]s"
			pi_volume_name       = "%[2]s"
			pi_volume_size       = 1
			pi_volume_type        = "tier3"
		  }
		  resource "ibm_pi_instance" "power_instance" {
			pi_cloud_instance_id  = "%[1]s"
			pi_health_status      = "%[5]s"
			pi_ibmi_css 		  = %[6]t
			pi_ibmi_rds_users 	  = %[7]d
			pi_image_id           = data.ibm_pi_image.power_image.id
			pi_instance_name      = "%[2]s"
			pi_memory             = "2"
			pi_proc_type          = "shared"
			pi_processors         = "0.25"
			pi_storage_pool       = data.ibm_pi_image.power_image.storage_pool
			pi_sys_type           = "s922"
			pi_volume_ids         = [ibm_pi_volume.power_volume.volume_id]
			pi_network {
				network_id = data.ibm_pi_network.power_networks.id
			}
		  }`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, instanceHealthStatus, IBMiCSS, IBMiRDSUsers)
}

func TestAccIBMPIInstanceReplicant(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceReplicantConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_replicants", "3"),
					resource.TestCheckResourceAttr(instanceRes, "pi_replication_policy", power.Affinity),
					resource.TestCheckResourceAttr(instanceRes, "pi_replication_scheme", "suffix"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckIBMPIInstanceReplicantConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_volume" "power_volume" {
		pi_cloud_instance_id = "%[1]s"
		pi_volume_name       = "%[2]s"
		pi_volume_shareable  = true
		pi_volume_size       = 1
		pi_volume_type        = "tier3"
	  }
	resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id = "%[1]s"
		pi_image_id          = "%[3]s"
		pi_instance_name     = "%[2]s"
		pi_memory            = "2"
		pi_proc_type         = "shared"
		pi_processors        = "1"
		pi_replicants         = 3
		pi_replication_policy = "affinity"
		pi_replication_scheme = "suffix"
		pi_sys_type          = "s922"
		pi_volume_ids        = [resource.ibm_pi_volume.power_volume.volume_id]
		pi_network {
			network_id = "%[4]s"
		  }
	  }
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name)
}

func TestAccIBMPIInstanceNetwork(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	privateNetIP := "192.168.17.253"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMPIInstanceNetworkConfig(name, privateNetIP),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttrSet(instanceRes, "pi_network.0.network_id"),
					resource.TestCheckResourceAttrSet(instanceRes, "pi_network.0.mac_address"),
					resource.TestCheckResourceAttr(instanceRes, "pi_network.0.ip_address", privateNetIP),
				),
			},
		},
	})
}

func testAccIBMPIInstanceNetworkConfig(name, privateNetIP string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_key" "key" {
		pi_cloud_instance_id = "%[1]s"
		pi_key_name          = "%[2]s"
		pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAQEArb2aK0mekAdbYdY9rwcmeNSxqVCwez3WZTYEq+1Nwju0x5/vQFPSD2Kp9LpKBbxx3OVLN4VffgGUJznz9DAr7veLkWaf3iwEil6U4rdrhBo32TuDtoBwiczkZ9gn1uJzfIaCJAJdnO80Kv9k0smbQFq5CSb9H+F5VGyFue/iVd5/b30MLYFAz6Jg1GGWgw8yzA4Gq+nO7HtyuA2FnvXdNA3yK/NmrTiPCdJAtEPZkGu9LcelkQ8y90ArlKfjtfzGzYDE4WhOufFxyWxciUePh425J2eZvElnXSdGha+FCfYjQcvqpCVoBAG70U4fJBGjB+HL/GpCXLyiYXPrSnzC9w=="
	}
	resource "ibm_pi_network" "power_networks" {
		pi_cidr              = "192.168.17.0/24"
		pi_cloud_instance_id = "%[1]s"
		pi_dns               = ["127.0.0.1"]
		pi_gateway           = "192.168.17.2"
		pi_network_name      = "%[2]s"
		pi_network_type      = "vlan"
		pi_ipaddress_range {
			pi_ending_ip_address = "192.168.17.254"
			pi_starting_ip_address = "192.168.17.3"
		}
	}
	resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id  = "%[1]s"
		pi_image_id           = "%[4]s"
		pi_instance_name      = "%[2]s"
		pi_key_pair_name      = ibm_pi_key.key.name
		pi_memory             = "2"
		pi_proc_type          = "shared"
		pi_processors         = "0.25"
		pi_storage_type 	  = "tier3"
		pi_sys_type           = "s922"
		pi_network {
			network_id = resource.ibm_pi_network.power_networks.network_id
			ip_address = "%[3]s"
		}
	}
	`, acc.Pi_cloud_instance_id, name, privateNetIP, acc.Pi_image)
}

func TestAccIBMPIInstanceVTL(t *testing.T) {
	instanceRes := "ibm_pi_instance.vtl_instance"
	name := fmt.Sprintf("tf-pi-vtl-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMPIInstanceVTLConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(instanceRes, "pi_license_repository_capacity", "3"),
				),
			},
		},
	})
}

func testAccIBMPIInstanceVTLConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_key" "vtl_key" {
		pi_cloud_instance_id = "%[1]s"
		pi_key_name          = "%[2]s"
		pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAQEArb2aK0mekAdbYdY9rwcmeNSxqVCwez3WZTYEq+1Nwju0x5/vQFPSD2Kp9LpKBbxx3OVLN4VffgGUJznz9DAr7veLkWaf3iwEil6U4rdrhBo32TuDtoBwiczkZ9gn1uJzfIaCJAJdnO80Kv9k0smbQFq5CSb9H+F5VGyFue/iVd5/b30MLYFAz6Jg1GGWgw8yzA4Gq+nO7HtyuA2FnvXdNA3yK/NmrTiPCdJAtEPZkGu9LcelkQ8y90ArlKfjtfzGzYDE4WhOufFxyWxciUePh425J2eZvElnXSdGha+FCfYjQcvqpCVoBAG70U4fJBGjB+HL/GpCXLyiYXPrSnzC9w=="
	}

	resource "ibm_pi_network" "vtl_network" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[2]s"
		pi_network_type      = "pub-vlan"
	}

	resource "ibm_pi_instance" "vtl_instance" {
		pi_cloud_instance_id  = "%[1]s"
		pi_image_id           = "%[3]s"
		pi_instance_name      = "%[2]s"
		pi_key_pair_name      = ibm_pi_key.vtl_key.name
		pi_license_repository_capacity = "3"
		pi_memory             = "22"
		pi_proc_type          = "shared"
		pi_processors         = "2"
		pi_storage_type 	  = "tier1"
		pi_sys_type           = "s922"
		pi_network {
			network_id = ibm_pi_network.vtl_network.network_id
		}
	  }

	`, acc.Pi_cloud_instance_id, name, acc.Pi_image)
}

func TestAccIBMPISAPInstance(t *testing.T) {
	instanceRes := "ibm_pi_instance.sap"
	name := fmt.Sprintf("tf-pi-sap-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMPISAPInstanceConfig(name, "tinytest-1x4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(instanceRes, "pi_sap_profile_id", "tinytest-1x4"),
				),
			},
			{
				Config: testAccIBMPISAPInstanceConfig(name, "tinytest-1x8"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(instanceRes, "pi_sap_profile_id", "tinytest-1x8"),
				),
			},
		},
	})
}

func testAccIBMPISAPInstanceConfig(name, sapProfile string) string {
	return fmt.Sprintf(`

	resource "ibm_pi_network" "power_network" {
		pi_cidr              = "192.168.17.0/24"
		pi_cloud_instance_id = "%[1]s"
		pi_dns               = ["127.0.0.1"]
		pi_gateway           = "192.168.17.2"
		pi_network_name      = "%[2]s"
		pi_network_type      = "vlan"
		pi_ipaddress_range {
			pi_ending_ip_address = "192.168.17.254"
			pi_starting_ip_address = "192.168.17.3"
		}
	}
	resource "ibm_pi_instance" "sap" {
		pi_cloud_instance_id  	= "%[1]s"
		pi_health_status		= "OK"
		pi_image_id           	= "%[4]s"
		pi_instance_name      	= "%[2]s"
		pi_sap_profile_id       = "%[3]s"
		pi_storage_type			= "tier1"
		pi_network {
			network_id = resource.ibm_pi_network.power_network.network_id
		}
	}
	`, acc.Pi_cloud_instance_id, name, sapProfile, acc.Pi_sap_image)
}

func TestAccIBMPIInstanceMixedStorage(t *testing.T) {
	instanceRes := "ibm_pi_instance.instance"
	name := fmt.Sprintf("tf-pi-mixedstorage-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMPIInstanceMixedStorage(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(instanceRes, "pi_storage_pool_affinity", "false"),
				),
			},
		},
	})
}

func testAccIBMPIInstanceMixedStorage(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_key" "key" {
		pi_cloud_instance_id = "%[1]s"
		pi_key_name          = "%[2]s"
		pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
	}
	resource "ibm_pi_network" "power_network" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[2]s"
		pi_network_type      = "vlan"
		pi_cidr              = "192.168.17.0/24"
	}
	resource "ibm_pi_volume" "power_volume" {
		pi_cloud_instance_id = "%[1]s"
		pi_volume_name       = "%[2]s"
		pi_volume_shareable  = true
		pi_volume_size       = 20
		pi_volume_type       = "tier3"
	}
	resource "ibm_pi_instance" "instance" {
		pi_cloud_instance_id     = "%[1]s"
		pi_image_id              = "%[3]s"
		pi_instance_name         = "%[2]s"
		pi_key_pair_name         = ibm_pi_key.key.name
		pi_memory                = "2"
		pi_proc_type             = "shared"
		pi_processors            = "0.25"
		pi_storage_pool_affinity = false
		pi_storage_type          = "tier1"
		pi_sys_type              = "s922"
		pi_network {
			network_id = ibm_pi_network.power_network.network_id
		}
	}
	resource "ibm_pi_volume_attach" "power_attach_volume"{
		pi_cloud_instance_id = "%[1]s"
		pi_instance_id       = ibm_pi_instance.instance.instance_id
		pi_volume_id         = ibm_pi_volume.power_volume.volume_id
	}
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image)
}

func TestAccIBMPIInstanceUpdateActiveState(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIActiveInstanceConfigUpdate(name, power.OK, "0.25", "2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(instanceRes, "status", strings.ToUpper(power.State_Active)),
				),
			},
			{
				Config: testAccCheckIBMPIActiveInstanceConfigUpdate(name, power.OK, "0.5", "4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceStatus(instanceRes, strings.ToUpper(power.State_Active)),
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
				),
			},
		},
	})
}

func TestAccIBMPIInstanceUpdateStoppedState(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIStoppedInstanceConfigUpdate(name, power.OK, "0.25", "2", "stop"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
				),
			},
			{
				Config: testAccCheckIBMPIStoppedInstanceConfigUpdate(name, power.OK, "0.5", "4", "stop"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceStatus(instanceRes, strings.ToUpper(power.State_Shutoff)),
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
				),
			},
		},
	})
}

func testAccCheckIBMPIActiveInstanceConfigUpdate(name, instanceHealthStatus, proc, memory string) string {
	return fmt.Sprintf(`
	data "ibm_pi_image" "power_image" {
		pi_cloud_instance_id = "%[1]s"
		pi_image_name        = "%[3]s"
	}
	data "ibm_pi_network" "power_networks" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[4]s"
	}
	resource "ibm_pi_volume" "power_volume" {
		pi_cloud_instance_id = "%[1]s"
		pi_volume_name       = "%[2]s"
		pi_volume_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_volume_shareable  = true
		pi_volume_size       = 20
	}
	resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id = "%[1]s"
		pi_health_status     = "%[5]s"
		pi_image_id          = data.ibm_pi_image.power_image.id
		pi_instance_name     = "%[2]s"
		pi_memory            = "%[7]s"
		pi_pin_policy        = "none"
		pi_proc_type         = "shared"
		pi_processors        = "%[6]s"
		pi_storage_pool      = data.ibm_pi_image.power_image.storage_pool
		pi_sys_type          = "s922"
		pi_volume_ids        = [ibm_pi_volume.power_volume.volume_id]
		pi_network {
			network_id = data.ibm_pi_network.power_networks.id
		}
	}
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, instanceHealthStatus, proc, memory)
}

func testAccCheckIBMPIStoppedInstanceConfigUpdate(name, instanceHealthStatus, proc, memory, action string) string {
	return fmt.Sprintf(`
	data "ibm_pi_image" "power_image" {
		pi_cloud_instance_id = "%[1]s"
		pi_image_name        = "%[3]s"
	}
	data "ibm_pi_network" "power_networks" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[4]s"
	}
	resource "ibm_pi_volume" "power_volume" {
		pi_cloud_instance_id = "%[1]s"
		pi_volume_name       = "%[2]s"
		pi_volume_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_volume_shareable  = true
		pi_volume_size       = 20
	}
	resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id  = "%[1]s"
		pi_health_status      = "%[5]s"
		pi_image_id           = data.ibm_pi_image.power_image.id
		pi_instance_name      = "%[2]s"
		pi_memory             = "%[7]s"
		pi_pin_policy         = "none"
		pi_proc_type          = "shared"
		pi_processors         = "%[6]s"
		pi_storage_pool       = data.ibm_pi_image.power_image.storage_pool		
		pi_sys_type           = "s922"
		pi_volume_ids         = [ibm_pi_volume.power_volume.volume_id]
		pi_network {
			network_id = data.ibm_pi_network.power_networks.id
		}
	}
	resource "ibm_pi_instance_action" "power_instance_action" {
  		pi_action            = "%[8]s"
  		pi_cloud_instance_id = "%[1]s"
  		pi_health_status     = "%[5]s"
  		pi_instance_id       = ibm_pi_instance.power_instance.instance_id
	}
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, instanceHealthStatus, proc, memory, action)
}

func TestAccIBMPIInstanceVirtualSerialNumber(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	description := "VSN for TF test"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceVirtualSerialNumber(name, power.OK, "s1022", "P05", description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttrSet(instanceRes, "pi_virtual_serial_number.0.serial"),
					resource.TestCheckResourceAttr(instanceRes, "pi_virtual_serial_number.0.description", description),
					resource.TestCheckResourceAttr(instanceRes, "pi_virtual_serial_number.0.software_tier", "P05"),
				),
			},
			{
				Config: testAccCheckIBMPIInstanceVirtualSerialNumber(name, power.OK, "s1022", "P10", description+" updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttrSet(instanceRes, "pi_virtual_serial_number.0.serial"),
					resource.TestCheckResourceAttr(instanceRes, "pi_virtual_serial_number.0.description", description+" updated"),
					resource.TestCheckResourceAttr(instanceRes, "pi_virtual_serial_number.0.software_tier", "P10"),
				),
			},
			{
				Config: testAccCheckIBMPIInstanceVirtualSerialNumber(name, power.OK, "s1022", "P05", description+" updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttrSet(instanceRes, "pi_virtual_serial_number.0.serial"),
					resource.TestCheckResourceAttr(instanceRes, "pi_virtual_serial_number.0.description", description+" updated"),
					resource.TestCheckResourceAttr(instanceRes, "pi_virtual_serial_number.0.software_tier", "P05"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceVirtualSerialNumber(name, instanceHealthStatus, systype string, softwareTier string, description string) string {
	return fmt.Sprintf(`
	  data "ibm_pi_image" "power_image" {
		pi_cloud_instance_id = "%[1]s"
		pi_image_name        = "%[3]s"
	  }
	  data "ibm_pi_network" "power_networks" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[4]s"
	  }
	  resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id  = "%[1]s"
		pi_health_status      = "%[5]s"
		pi_image_id           = data.ibm_pi_image.power_image.id
		pi_instance_name      = "%[2]s"
		pi_memory             = "2"
		pi_proc_type          = "shared"
		pi_processors         = "0.25"
		pi_storage_type 	  = "%[7]s"
		pi_sys_type           = "%[6]s"
		pi_network {
			network_id = data.ibm_pi_network.power_networks.id
		}
		pi_virtual_serial_number {
			serial        = "auto-assign"
			description   = "%[9]s"
			software_tier = "%[8]s"
		}
	  }
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, instanceHealthStatus, systype, acc.PiStorageType, softwareTier, description)
}

func TestAccIBMPIInstanceDeploymentGRS(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	bootVolumeData := "data.ibm_pi_volume.power_boot_volume_data"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMPIInstanceGRSConfig(name, power.OK, "2", "0.25"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(bootVolumeData, "replication_enabled", "true"),
				),
			},
		},
	})
}

func testAccIBMPIInstanceGRSConfig(name string, instanceHealthStatus string, memory string, proc string) string {
	return fmt.Sprintf(`
	data "ibm_pi_image" "power_image" {
		pi_image_name        = "%[3]s"
		pi_cloud_instance_id = "%[1]s"
	}
	data "ibm_pi_network" "power_networks" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[4]s"
	}  
	data "ibm_pi_volume" "power_boot_volume_data" {
		pi_cloud_instance_id = "%[1]s"
		pi_volume_name       = data.ibm_pi_instance_volumes.power_instance_volumes_data.instance_volumes[0].name
	}
	data "ibm_pi_instance_volumes" "power_instance_volumes_data" {
		pi_cloud_instance_id = "%[1]s"
		pi_instance_name     = ibm_pi_instance.power_instance.pi_instance_name
	}
	resource "ibm_pi_instance" "power_instance" {
		pi_boot_volume_replication_enabled = true
		pi_memory            			   = "%[7]s"
		pi_processors        			   = "%[6]s"
		pi_instance_name                           = "%[2]s"
		pi_proc_type          			   = "shared"
		pi_image_id           			   = data.ibm_pi_image.power_image.id
		pi_sys_type          			   = "e980"
		pi_cloud_instance_id  			   = "%[1]s"
		pi_storage_pool       			   = data.ibm_pi_image.power_image.storage_pool		
		pi_pin_policy        			   = "none"
		pi_health_status      			   = "%[5]s"
		pi_network {
			network_id = data.ibm_pi_network.power_networks.id
		}
	}
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, instanceHealthStatus, proc, memory)
}

func TestAccIBMPIInstanceUserTags(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	userTagsString := `["env:dev", "test_tag"]`
	userTagsStringUpdated := `["env:dev", "test_tag", "test_tag2"]`
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceUserTagsConfig(name, power.OK, userTagsString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(instanceRes, "pi_user_tags.#", "2"),
					resource.TestCheckTypeSetElemAttr(instanceRes, "pi_user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(instanceRes, "pi_user_tags.*", "test_tag"),
				),
			},
			{
				Config: testAccCheckIBMPIInstanceUserTagsConfig(name, power.OK, userTagsStringUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(instanceRes, "pi_user_tags.#", "3"),
					resource.TestCheckTypeSetElemAttr(instanceRes, "pi_user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(instanceRes, "pi_user_tags.*", "test_tag"),
					resource.TestCheckTypeSetElemAttr(instanceRes, "pi_user_tags.*", "test_tag2"),
				),
			},
		},
	})
}

func testAccCheckIBMPIInstanceUserTagsConfig(name, instanceHealthStatus string, userTagsString string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_key" "key" {
		pi_cloud_instance_id = "%[1]s"
		pi_key_name          = "%[2]s"
		pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
	  }
	  data "ibm_pi_image" "power_image" {
		pi_cloud_instance_id = "%[1]s"
		pi_image_name        = "%[3]s"
	  }
	  data "ibm_pi_network" "power_networks" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[4]s"
	  }
	  resource "ibm_pi_volume" "power_volume" {
		pi_cloud_instance_id = "%[1]s"
		pi_volume_name       = "%[2]s"
		pi_volume_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_volume_shareable  = true
		pi_volume_size       = 20
		pi_volume_type       = "%[6]s"
	  }
	  resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id  = "%[1]s"
		pi_health_status      = "%[5]s"
		pi_image_id           = data.ibm_pi_image.power_image.id
		pi_instance_name      = "%[2]s"
		pi_key_pair_name      = ibm_pi_key.key.name
		pi_memory             = "2"
		pi_proc_type          = "shared"
		pi_processors         = "0.25"
		pi_storage_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_storage_type       = "%[6]s"
		pi_sys_type           = "s922"
		pi_volume_ids         = [ibm_pi_volume.power_volume.volume_id]
		pi_network {
			network_id = data.ibm_pi_network.power_networks.id
		}
		pi_user_tags          = %[7]s
	  }
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, instanceHealthStatus, acc.PiStorageType, userTagsString)
}
func TestAccIBMPIInstanceCompMode(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPInstanceCompMode(name, power.OK, "0.25", "2", "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(instanceRes, "pi_preferred_processor_compatibility_mode", "default"),
					resource.TestCheckResourceAttr(instanceRes, "status", strings.ToUpper(power.State_Active)),
				),
			},
			{
				Config: testAccCheckIBMPInstanceCompMode(name, power.OK, "0.25", "2", "POWER10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(instanceRes, "pi_preferred_processor_compatibility_mode", "POWER10"),
				),
			},
		},
	})
}
func testAccCheckIBMPInstanceCompMode(name, instanceHealthStatus, proc, memory, compMode string) string {
	return fmt.Sprintf(`
	data "ibm_pi_image" "power_image" {
		pi_cloud_instance_id = "%[1]s"
		pi_image_name        = "%[3]s"
	}
	data "ibm_pi_network" "power_networks" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[4]s"
	}
	resource "ibm_pi_volume" "power_volume" {
		pi_cloud_instance_id = "%[1]s"
		pi_volume_name       = "%[2]s"
		pi_volume_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_volume_shareable  = true
		pi_volume_size       = 20
	}
	resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id = "%[1]s"
		pi_health_status     = "%[5]s"
		pi_image_id          = data.ibm_pi_image.power_image.id
		pi_instance_name     = "%[2]s"
		pi_memory            = "%[7]s"
		pi_pin_policy        = "none"
		pi_proc_type         = "shared"
		pi_processors        = "%[6]s"
		pi_preferred_processor_compatibility_mode =  "%[8]s"
		pi_storage_pool      = data.ibm_pi_image.power_image.storage_pool
		pi_sys_type          = "s1022"
		pi_volume_ids        = [ibm_pi_volume.power_volume.volume_id]
		pi_network {
			network_id = data.ibm_pi_network.power_networks.id
		}
	}
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, instanceHealthStatus, proc, memory, compMode)
}
func TestAccIBMPIInstanceVPMEM(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pvm-vpmem-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceVPMEMConfig(name, power.OK),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttr(instanceRes, "vpmem_volumes.#", "2"),
				),
			},
		},
	})
}
func testAccCheckIBMPIInstanceVPMEMConfig(name, instanceHealthStatus string) string {
	return fmt.Sprintf(`
	  data "ibm_pi_image" "power_image" {
		pi_cloud_instance_id = "%[1]s"
		pi_image_name        = "%[3]s"
	  }
	  data "ibm_pi_network" "power_networks" {
		pi_cloud_instance_id = "%[1]s"
		pi_network_name      = "%[4]s"
	  }
	  resource "ibm_pi_volume" "power_volume" {
		pi_cloud_instance_id = "%[1]s"
		pi_volume_name       = "%[2]s-vol"
		pi_volume_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_volume_shareable  = true
		pi_volume_size       = 20
		pi_volume_type       = "%[6]s"
	  }
	  resource "ibm_pi_instance" "power_instance" {
		pi_cloud_instance_id  = "%[1]s"
		pi_health_status      = "%[5]s"
		pi_image_id           = data.ibm_pi_image.power_image.id
		pi_instance_name      = "%[2]s"
		pi_memory             = "2"
		pi_proc_type          = "shared"
		pi_processors         = "0.25"
		pi_storage_pool       = data.ibm_pi_image.power_image.storage_pool
		pi_storage_type       = "%[6]s"
		pi_sys_type           = "s1022"
		pi_volume_ids         = [ibm_pi_volume.power_volume.volume_id]
		pi_network {
			network_id = data.ibm_pi_network.power_networks.id
		}
		pi_vpmem_volumes {
			name = "%[2]s-1"
			size = 2
		}
		pi_vpmem_volumes {
			name = "%[2]s-2"
			size = 3
		}
	  }
	`, acc.Pi_cloud_instance_id, name, acc.Pi_image, acc.Pi_network_name, instanceHealthStatus, acc.PiStorageType)
}
func testAccCheckIBMPIInstanceDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_instance" {
			continue
		}

		idArr, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudInstanceID := idArr[0]
		for _, instanceID := range idArr[1:] {
			client := st.NewIBMPIInstanceClient(context.Background(), sess, cloudInstanceID)
			_, err = client.Get(instanceID)
			if err == nil {
				return fmt.Errorf("PI Instance still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckIBMPIInstanceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}

		idArr, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		cloudInstanceID := idArr[0]
		for _, instanceID := range idArr[1:] {
			client := st.NewIBMPIInstanceClient(context.Background(), sess, cloudInstanceID)
			_, err = client.Get(instanceID)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func testAccCheckIBMPIInstanceStatus(n, status string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}

		cloudInstanceID, instanceID, err := splitID(rs.Primary.ID)
		if err == nil {
			return err
		}
		client := st.NewIBMPIInstanceClient(context.Background(), sess, cloudInstanceID)

		instance, err := client.Get(instanceID)
		if err != nil {
			return err
		}

		for {
			if instance.Status != &status {
				time.Sleep(2 * time.Minute)
			} else {
				break
			}
		}

		return nil
	}
}
