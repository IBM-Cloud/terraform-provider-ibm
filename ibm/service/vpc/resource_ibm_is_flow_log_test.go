// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISFlowLog_basic(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("flowlog-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("resource-instance-%d", acctest.RandIntRange(10, 100))
	flowlogname := fmt.Sprintf("flowlog-instance-%d", acctest.RandIntRange(10, 100))
	newflowlogname := fmt.Sprintf("newflowlog-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("flowlog-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))

	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISFlowLogDestroy,
		Steps: []resource.TestStep{

			{
				//Create test case
				Config: testAccCheckIBMISFlowLogConfig(vpcname, name, flowlogname, sshname, publicKey, subnetname, serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFlowLogExists("ibm_is_flow_log.test_flow_log", instance),
					resource.TestCheckResourceAttr("ibm_is_flow_log.test_flow_log", "name", flowlogname),
				),
			},
			//update
			{
				Config: testAccCheckIBMISFlowLogConfig(vpcname, name, newflowlogname, sshname, publicKey, subnetname, serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFlowLogExists("ibm_is_flow_log.test_flow_log", instance),
					resource.TestCheckResourceAttr("ibm_is_flow_log.test_flow_log", "name", newflowlogname),
					resource.TestCheckResourceAttr("ibm_is_flow_log.test_flow_log", "active", "false"),
				),
			},
		},
	},
	)
}
func TestAccIBMISFlowLog_vni(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("flowlog-vpc-%d", acctest.RandIntRange(10, 100))
	flowlogname := fmt.Sprintf("flowlog-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("flowlog-subnet-%d", acctest.RandIntRange(10, 100))

	vniname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISFlowLogDestroy,
		Steps: []resource.TestStep{

			{
				//Create test case
				Config: testAccCheckIBMISFlowLogVniConfig(vpcname, vniname, flowlogname, subnetname, bucketName, bucketRegionType, bucketRegion, bucketClass, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFlowLogExists("ibm_is_flow_log.test_flow_log", instance),
					resource.TestCheckResourceAttr("ibm_is_flow_log.test_flow_log", "name", flowlogname),
					resource.TestCheckResourceAttr("data.ibm_is_flow_log.is_flow_log_name", "target.0.resource_type", "virtual_network_interface"),
				),
			},
		},
	},
	)
}

func testAccCheckIBMISFlowLogConfig(vpcname, name, flowlogname, sshname, publicKey, subnetname, serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass string, isActive bool) string {
	return fmt.Sprintf(`	  	
	
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	}

	resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc     = ibm_is_vpc.testacc_vpc.id
		keys    = [ibm_is_ssh_key.testacc_sshkey.id]
		zone    = "%s"
	  }

	data "ibm_resource_group" "cos_group" {
		is_default=true
	}
	  
	resource "ibm_resource_instance" "instance2" {
		name              = "%s"
		resource_group_id = data.ibm_resource_group.cos_group.id
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
	  }

	  resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		region_location      = "%s"
		storage_class        = "%s"
	}	  
	// Authorisation policy is required between vpc Flowlogs and Object Storage
	
	resource "ibm_is_flow_log" "test_flow_log" {
		name    = "%s"
		target = ibm_is_instance.testacc_instance.id
		storage_bucket = ibm_cos_bucket.bucket2.bucket_name
		active = %v
	  } 
	  
	  `, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, serviceName, bucketName, bucketRegion, bucketClass, flowlogname, isActive)

}
func testAccCheckIBMISFlowLogVniConfig(vpcname, vniname, flowlogname, subnetname, bucketName, bucketRegionType, bucketRegion, bucketClass string, isActive bool) string {
	return fmt.Sprintf(`	  	
	
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		total_ipv4_address_count = 16
	}

	resource "ibm_is_virtual_network_interface" "testacc_vni"{
		name = "%s"
		subnet = ibm_is_subnet.testacc_subnet.id
	}

	resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		region_location      = "%s"
		storage_class        = "%s"
	}	  
	// Authorisation policy is required between vpc Flowlogs and Object Storage
	
	resource "ibm_is_flow_log" "test_flow_log" {
		name    = "%s"
		target = ibm_is_virtual_network_interface.testacc_vni.id
		storage_bucket = ibm_cos_bucket.bucket2.bucket_name
		active = %v
	}
	data "ibm_is_flow_log" "is_flow_log_name" {
		name = ibm_is_flow_log.test_flow_log.name
	}
	  
	  `, vpcname, subnetname, acc.ISZoneName, vniname, bucketName, acc.ISResourceCrn, bucketRegion, bucketClass, flowlogname, isActive)

}
func vpcClient(meta interface{}) (*vpcv1.VpcV1, error) {
	sess, err := meta.(conns.ClientSession).VpcV1API()
	return sess, err
}
func testAccCheckIBMISFlowLogDestroy(s *terraform.State) error {
	sess, err := vpcClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_flow_log" {
			continue
		}
		log.Printf("Destroy called ...%s", rs.Primary.ID)
		getOptions := &vpcv1.GetFlowLogCollectorOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetFlowLogCollector(getOptions)
		if err == nil {
			return fmt.Errorf("flow log still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISFlowLogExists(n string, instance string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		sess, err := vpcClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}
		getOptions := &vpcv1.GetFlowLogCollectorOptions{
			ID: &rs.Primary.ID,
		}
		instance1, response, err := sess.GetFlowLogCollector(getOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Getting Flow log: %s\n%s", err, response)
		}
		instance = *instance1.ID
		return nil
	}
}

func TestAccIBMISFlowLogImport(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("flowlog-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("resource-instance-%d", acctest.RandIntRange(10, 100))
	flowlogname := fmt.Sprintf("flowlog-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("flowlog-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))

	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISFlowLogDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISFlowLogConfig(vpcname, name, flowlogname, sshname, publicKey, subnetname, serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFlowLogExists("ibm_is_flow_log.test_flow_log", instance),
					resource.TestCheckResourceAttr("ibm_is_flow_log.test_flow_log", "name", flowlogname),
				),
			},
			{
				ResourceName:      "ibm_is_flow_log.test_flow_log",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
