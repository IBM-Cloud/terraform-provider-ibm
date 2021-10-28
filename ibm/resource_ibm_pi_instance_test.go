// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
)

func testAccCheckIBMPIInstanceConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_key" "key" {
		pi_cloud_instance_id = "%[1]s"
		pi_key_name          = "%[2]s"
		pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
	  }
	  
	  resource "ibm_pi_volume" "power_volume" {
		pi_volume_size       = 20
		pi_volume_name       = "%[2]s"
		pi_volume_type       = "tier3"
		pi_volume_shareable  = true
		pi_cloud_instance_id = "%[1]s"
	  }
	  resource "ibm_pi_instance" "power_instance" {
		pi_memory             = "4"
		pi_processors         = "2"
		pi_instance_name      = "%[2]s"
		pi_proc_type          = "shared"
		pi_image_id           = "f4501cad-d0f4-4517-9eea-85402309d90d"
		pi_network_ids        = ["1bfad140-588f-45d3-aaaa-a3f0abbd441f"]
		pi_key_pair_name      = ibm_pi_key.key.key_id
		pi_sys_type           = "s922"
		pi_cloud_instance_id  = "%[1]s"
		pi_volume_ids         = [ibm_pi_volume.power_volume.volume_id]
	  }
	`, pi_cloud_instance_id, name)
}

func testAccIBMPIInstanceNetworkConfig(name, privateNetIP string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_key" "key" {
		pi_cloud_instance_id = "%[1]s"
		pi_key_name          = "%[2]s"
		pi_ssh_key           = "ssh-rsa AAAAB3NzaC1yc2EAAAABJQAAAQEArb2aK0mekAdbYdY9rwcmeNSxqVCwez3WZTYEq+1Nwju0x5/vQFPSD2Kp9LpKBbxx3OVLN4VffgGUJznz9DAr7veLkWaf3iwEil6U4rdrhBo32TuDtoBwiczkZ9gn1uJzfIaCJAJdnO80Kv9k0smbQFq5CSb9H+F5VGyFue/iVd5/b30MLYFAz6Jg1GGWgw8yzA4Gq+nO7HtyuA2FnvXdNA3yK/NmrTiPCdJAtEPZkGu9LcelkQ8y90ArlKfjtfzGzYDE4WhOufFxyWxciUePh425J2eZvElnXSdGha+FCfYjQcvqpCVoBAG70U4fJBGjB+HL/GpCXLyiYXPrSnzC9w=="
	}
	resource "ibm_pi_instance" "power_instance" {
		pi_memory             = "2"
		pi_processors         = "0.25"
		pi_instance_name      = "%[2]s"
		pi_proc_type          = "shared"
		pi_image_id           = "f4501cad-d0f4-4517-9eea-85402309d90d"
		pi_key_pair_name      = ibm_pi_key.key.key_id
		pi_sys_type           = "e980"
		pi_storage_type 	  = "tier3"
		pi_cloud_instance_id  = "%[1]s"
		pi_network {
			network_id = "tf-cloudconnection-23"
			ip_address = "%[3]s"
		}
	}
	`, pi_cloud_instance_id, name, privateNetIP)
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
		pi_memory             = "22"
		pi_processors         = "2"
		pi_instance_name      = "%[2]s"
		pi_license_repository_capacity = "3"
		pi_proc_type          = "shared"
		pi_image_id           = "ca4ea55f-b329-4cf5-bdce-d2f38cfc6da3"
		pi_network_ids        = [ibm_pi_network.vtl_network.network_id]
		pi_key_pair_name      = ibm_pi_key.vtl_key.key_id
		pi_sys_type           = "s922"
		pi_cloud_instance_id  = "%[1]s"
		pi_storage_type 	  = "tier1"
	  }
	
	`, pi_cloud_instance_id, name)
}

func testAccCheckIBMPIInstanceDestroy(s *terraform.State) error {

	sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_instance" {
			continue
		}
		parts, err := idParts(rs.Primary.ID)
		powerinstanceid := parts[0]
		networkC := st.NewIBMPIInstanceClient(sess, powerinstanceid)
		_, err = networkC.Get(parts[1], powerinstanceid, getTimeOut)
		if err == nil {
			return fmt.Errorf("PI Instance still exists: %s", rs.Primary.ID)
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

		sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		powerinstanceid := parts[0]
		client := st.NewIBMPIInstanceClient(sess, powerinstanceid)

		instance, err := client.Get(parts[1], powerinstanceid, getTimeOut)
		if err != nil {
			return err
		}
		parts[1] = *instance.PvmInstanceID
		return nil

	}
}

func TestAccIBMPIInstanceBasic(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIInstanceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
				),
			},
		},
	})
}

func TestAccIBMPIInstanceNetwork(t *testing.T) {
	instanceRes := "ibm_pi_instance.power_instance"
	name := fmt.Sprintf("tf-pi-instance-%d", acctest.RandIntRange(10, 100))
	privateNetIP := "192.112.111.220"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
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

func TestAccIBMPIInstanceVTL(t *testing.T) {
	instanceRes := "ibm_pi_instance.vtl_instance"
	name := fmt.Sprintf("tf-pi-vtl-instance-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPIInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMPIInstanceVTLConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIInstanceExists(instanceRes),
					resource.TestCheckResourceAttr(instanceRes, "pi_instance_name", name),
					resource.TestCheckResourceAttrSet(instanceRes, "license_repository_capacity"),
					resource.TestCheckResourceAttr(instanceRes, "license_repository_capacity", "3"),
				),
			},
		},
	})
}
