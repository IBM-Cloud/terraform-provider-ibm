package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISInstanceGroup_dataBasic(t *testing.T) {
	name := fmt.Sprintf("testinstancegroup%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC0eZ4uNSH6rtxNM7MagrBtxwlASw0iKcxDdXq9eNu93xDpsvxdn6xE/JESlIHhf9/45oLw9AKpu/MZYwqQ0O+uedwgtLvorv++fyXI36cls4xmUCuNnEhoK1aXh26N+V+lxejqF3DJhMKHYprQCnyl/8RWkWIFc2Jo60ACZ98MY4rRHgBP/0t1tqmb0I4IdBaYLctVIdv16gYJ5zqGYKeJBMG7XtgkrtOeacVoArrmjHY6n2cNgE5jLt9n9MyyGLjzq1agIpwwsWJGfhzqo2I98UhGWtUlD2UeNHbuJZmbXeyibuoV7RhDZg9LafkOXOTojjZc9rUrd8BChoHhbW3X anil@Anils-MacBook-Pro.local
	`)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceGroupDConfig(name, publicKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance_group.instance_group_data", "name", name),
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance_group.instance_group_data", "status", "healthy"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceGroupDConfig(name, publicKey string) string {
	return fmt.Sprintf(`
	provider "ibm" {
		generation = 2
	}
	
	resource "ibm_is_vpc" "vpc2" {
	  name = "vpc2test"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "subnet2"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "ssh1"
	  public_key = "%s"
	}
	
	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "testtemplate"
	   image   = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
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

	data "ibm_is_instance_group" "instance_group_data" {
		name = ibm_is_instance_group.instance_group.name
	}
	`, publicKey, name)

}
