// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISInstanceTemplate_basic(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("tf-testvpc%d", randInt)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateConfig(vpcName, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceTemplate_WithVolumeAttachment(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("tf-testvpc%d", randInt)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	volAttachName := fmt.Sprintf("tf-testvolattach%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateWithVolume(vpcName, subnetName, sshKeyName, publicKey, templateName, volAttachName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceTemplateDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance_template" {
			continue
		}

		getInstanceTemplateOptions := vpcv1.GetInstanceTemplateOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetInstanceTemplate(&getInstanceTemplateOptions)

		if err == nil {
			return fmt.Errorf("instance template still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISInstanceTemplateConfig(vpcName, subnetName, sshKeyName, publicKey, templateName string) string {
	return fmt.Sprintf(`
	provider "ibm" {
		generation = 2
	}
	
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

	data "ibm_is_images" "is_images" {
	}

	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = data.ibm_is_images.is_images.images.0.id
	   profile = "bx2-8x32"
	
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
	   }
	
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	
	`, vpcName, subnetName, sshKeyName, publicKey, templateName)

}

func testAccCheckIBMISInstanceTemplateWithVolume(vpcName, subnetName, sshKeyName, publicKey, templateName, volAttachName string) string {
	return fmt.Sprintf(`
	provider "ibm" {
		generation = 2
	}
	
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

	data "ibm_is_images" "is_images" {
	}

	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = data.ibm_is_images.is_images.images.0.id
	   profile = "bx2-8x32"
	
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
	   }
	   volume_attachments {
        delete_volume_on_instance_delete = true
        name                             = "%s"
			new_volume {
				iops = 9000
				profile = "general-purpose"
				capacity = 100
			}   
    	}
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	
	`, vpcName, subnetName, sshKeyName, publicKey, templateName, volAttachName)

}
