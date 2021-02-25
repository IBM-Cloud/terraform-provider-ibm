/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISFlowLogDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				//Create test case
				Config: testAccCheckIBMISFlowLogConfig(vpcname, name, flowlogname, sshname, publicKey, subnetname, serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFlowLogExists("ibm_is_flow_log.test_flow_log", instance),
					resource.TestCheckResourceAttr("ibm_is_flow_log.test_flow_log", "name", flowlogname),
				),
			},
			//update
			resource.TestStep{
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
		name = "default"
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
	  
	  `, vpcname, subnetname, ISZoneName, ISCIDR, sshname, publicKey, name, isImage, instanceProfileName, ISZoneName, serviceName, bucketName, bucketRegion, bucketClass, flowlogname, isActive)

}
func testAccCheckIBMISFlowLogDestroy(s *terraform.State) error {
	sess, err := vpcClient(testAccProvider.Meta())
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
		sess, err := vpcClient(testAccProvider.Meta())
		if err != nil {
			return err
		}
		getOptions := &vpcv1.GetFlowLogCollectorOptions{
			ID: &rs.Primary.ID,
		}
		instance1, response, err := sess.GetFlowLogCollector(getOptions)
		if err != nil {
			return fmt.Errorf("Error Getting Flow log: %s\n%s", err, response)
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISFlowLogDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISFlowLogConfig(vpcname, name, flowlogname, sshname, publicKey, subnetname, serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISFlowLogExists("ibm_is_flow_log.test_flow_log", instance),
					resource.TestCheckResourceAttr("ibm_is_flow_log.test_flow_log", "name", flowlogname),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_flow_log.test_flow_log",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
