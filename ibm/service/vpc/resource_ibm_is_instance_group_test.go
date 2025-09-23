// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISInstanceGroup_basic(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceGroupConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group", "name", instanceGroupName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group", "instance_count", "2"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceGroup_basic_loadbalancer(t *testing.T) {
	// var lb string
	randInt := acctest.RandIntRange(10, 100)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)

	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	poolName := fmt.Sprintf("tflbpoolc%d", acctest.RandIntRange(10, 100))
	alg1 := "round_robin"
	protocol1 := "http"
	// proxyProtocol1 := "disabled"
	delay1 := "45"
	retries1 := "5"
	timeout1 := "15"
	healthType1 := "http"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupDestroy,
		Steps: []resource.TestStep{
			// {
			// 	Config: testAccCheckIBMISLBPoolConfig(vpcName, subnetName, acc.ISZoneName, acc.ISCIDR, name, poolName, alg1, protocol1, delay1, retries1, timeout1, healthType1),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccCheckIBMISLBPoolExists("ibm_is_lb_pool.testacc_lb_pool", lb),
			// 		resource.TestCheckResourceAttr(
			// 			"ibm_is_lb.testacc_LB", "name", name),
			// 		resource.TestCheckResourceAttr(
			// 			"ibm_is_lb_pool.testacc_lb_pool", "name", poolName),
			// 		resource.TestCheckResourceAttr(
			// 			"ibm_is_lb_pool.testacc_lb_pool", "algorithm", alg1),
			// 		resource.TestCheckResourceAttr(
			// 			"ibm_is_lb_pool.testacc_lb_pool", "protocol", protocol1),
			// 		resource.TestCheckResourceAttr(
			// 			"ibm_is_lb_pool.testacc_lb_pool", "proxy_protocol", proxyProtocol1),
			// 		resource.TestCheckResourceAttr(
			// 			"ibm_is_lb_pool.testacc_lb_pool", "health_delay", delay1),
			// 		resource.TestCheckResourceAttr(
			// 			"ibm_is_lb_pool.testacc_lb_pool", "health_retries", retries1),
			// 		resource.TestCheckResourceAttr(
			// 			"ibm_is_lb_pool.testacc_lb_pool", "health_timeout", timeout1),
			// 		resource.TestCheckResourceAttr(
			// 			"ibm_is_lb_pool.testacc_lb_pool", "health_type", healthType1),
			// 	),
			// },
			{
				Config: testAccCheckIBMISInstanceGrouplbConfig(vpcName, subnetName, acc.ISZoneName, acc.ISCIDR, name, poolName, alg1, protocol1, delay1, retries1, timeout1, healthType1, sshKeyName, publicKey, templateName, instanceGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group", "name", instanceGroupName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group", "instance_count", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceGroupDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance_group" {
			continue
		}

		getInstanceGroupOptions := vpcv1.GetInstanceGroupOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetInstanceGroup(&getInstanceGroupOptions)

		if err == nil {
			return fmt.Errorf("instance group still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISInstanceGrouplbConfig(vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType, sshKeyName, publicKey, templateName, instanceGroupName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_lb" "testacc_LB" {
		name = "%s"
		subnets = ["${ibm_is_subnet.testacc_subnet.id}"]
	}
	resource "ibm_is_lb_pool" "testacc_lb_pool" {
		name = "%s"
		lb = "${ibm_is_lb.testacc_LB.id}"
		algorithm = "%s"
		protocol = "%s"
		health_delay= %s
		health_retries = %s
		health_timeout = %s
		health_type = "%s"	
	}
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}
	
	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = "%s"
	   profile = "bx2-8x32"
	
	   primary_network_interface {
		 subnet = ibm_is_subnet.testacc_subnet.id
	   }
	
	   vpc       = ibm_is_vpc.testacc_vpc.id
	   zone      = "%s"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	resource "ibm_is_instance_group" "instance_group" {
		name =  "%s"
		instance_template = ibm_is_instance_template.instancetemplate1.id
		instance_count =  1
		subnets = [ibm_is_subnet.testacc_subnet.id]
		load_balancer = ibm_is_lb.testacc_LB.id
        load_balancer_pool = ibm_is_lb_pool.testacc_lb_pool.pool_id
        application_port = "2364"
	}
	`, vpcname, subnetname, zone, cidr, name, poolName, algorithm, protocol, delay, retries, timeout, healthType, sshKeyName, publicKey, templateName, acc.IsImage, zone, instanceGroupName)

}

func testAccCheckIBMISInstanceGroupConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName string) string {
	return fmt.Sprintf(`
	
	
	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "%s"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}
	
	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = "%s"
	   profile = "bx2-8x32"
	
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
	   }
	
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	resource "ibm_is_instance_group" "instance_group" {
		name =  "%s"
		instance_template = ibm_is_instance_template.instancetemplate1.id
		instance_count =  2
		subnets = [ibm_is_subnet.subnet2.id]
	}
	`, vpcName, subnetName, sshKeyName, publicKey, templateName, acc.IsImage, instanceGroupName)

}
