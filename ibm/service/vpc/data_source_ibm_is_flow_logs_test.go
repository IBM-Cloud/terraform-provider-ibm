// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISFlowLogsDataSource_basic(t *testing.T) {
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
	resName := "data.ibm_is_flow_logs.display_flow_logs"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISFlowLogDestroy,
		Steps: []resource.TestStep{

			{
				//Create test case
				Config: testAccCheckIBMISFlowLogsDataSourceConfig(vpcname, name, flowlogname, sshname, publicKey, subnetname, serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFlowLogExists("ibm_is_flow_log.test_flow_log", instance),
					resource.TestCheckResourceAttr("ibm_is_flow_log.test_flow_log", "name", flowlogname),
				),
			},
			//update
			{
				Config: testAccCheckIBMISFlowLogsDataSourceConfig(vpcname, name, flowlogname, sshname, publicKey, subnetname, serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFlowLogExists("ibm_is_flow_log.test_flow_log", instance),
					resource.TestCheckResourceAttr("ibm_is_flow_log.test_flow_log", "active", "false"),
					resource.TestCheckResourceAttrSet(resName, "flow_log_collectors.0.id"),
					resource.TestCheckResourceAttrSet(resName, "flow_log_collectors.0.crn"),
					resource.TestCheckResourceAttrSet(resName, "flow_log_collectors.0.href"),
					resource.TestCheckResourceAttrSet(resName, "flow_log_collectors.0.name"),
				),
			},
		},
	},
	)
}

func testAccCheckIBMISFlowLogsDataSourceConfig(vpcname, name, flowlogname, sshname, publicKey, subnetname, serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass string, isActive bool) string {
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
        
         resource "ibm_iam_authorization_policy" "cos_policy" {
		source_service_name  = "is"
		source_resource_type = "flow-log-collector"
		target_service_name  = "cloud-object-storage"
		roles                = ["Writer"]
	  }

	resource "ibm_is_flow_log" "test_flow_log" {
		depends_on = [ibm_iam_authorization_policy.cos_policy]
                name    = "%s"
		target = ibm_is_instance.testacc_instance.id
		storage_bucket = ibm_cos_bucket.bucket2.bucket_name
		active = %v
	  } 

	  data "ibm_is_flow_logs" "display_flow_logs" {
	}
	  
	  `, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, serviceName, bucketName, bucketRegion, bucketClass, flowlogname, isActive)

}
