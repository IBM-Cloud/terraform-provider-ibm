// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package pag_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMResourceInstanceBasic(t *testing.T) {
	serviceName := fmt.Sprintf("tf-pag-%d", acctest.RandIntRange(10, 100))
	cos_instance_name := acc.PagCosInstanceName
	cos_bucket_name := acc.PagCosBucketName
	cos_bucket_region := acc.PagCosBucketRegion
	vpc_name := acc.PagVpcName
	service_plan := acc.PagServicePlan
	subnet_ins_1 := acc.PagVpcSubnetNameInstance_1
	subnet_ins_2 := acc.PagVpcSubnetNameInstance_2
	sg_ins_1 := acc.PagVpcSgInstance_1
	sg_ins_2 := acc.PagVpcSgInstance_2
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMResourceInstanceDestroy,
		Steps: []resource.TestStep{
			{

				Config: testAccCheckIBMResourceInstanceBasic(cos_instance_name, cos_bucket_name, cos_bucket_region, vpc_name, sg_ins_1, sg_ins_2, subnet_ins_1, subnet_ins_2, serviceName, service_plan),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMResourceInstanceExists("ibm_pag_instance.pag"),
					resource.TestCheckResourceAttr("ibm_pag_instance.pag", "name", serviceName),
					resource.TestCheckResourceAttr("ibm_pag_instance.pag", "service", "privileged-access-gateway"),
					resource.TestCheckResourceAttr("ibm_pag_instance.pag", "plan", service_plan),
					resource.TestCheckResourceAttr("ibm_pag_instance.pag", "location", "us-south"),
				),
			},
		},
	})
}

func testAccCheckIBMResourceInstanceDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pag_instance" {
			continue
		}

		instanceID := rs.Primary.ID
		resourceInstanceGet := rc.GetResourceInstanceOptions{
			ID: &instanceID,
		}

		// Try to find the key
		instance, resp, err := rsContClient.GetResourceInstance(&resourceInstanceGet)

		if err == nil {
			if *instance.State == "active" {
				return fmt.Errorf("Resource Instance still exists: %s", rs.Primary.ID)
			}
		} else {
			if !strings.Contains(err.Error(), "404") {
				return fmt.Errorf("[ERROR] Error checking if Resource Instance (%s) has been destroyed: %s with resp code: %s", rs.Primary.ID, err, resp)
			}
		}
	}

	return nil
}

func testAccCheckIBMResourceInstanceExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
		if err != nil {
			return err
		}
		instanceID := rs.Primary.ID
		resourceInstanceGet := rc.GetResourceInstanceOptions{
			ID: &instanceID,
		}

		_, resp, err := rsContClient.GetResourceInstance(&resourceInstanceGet)

		if err != nil {
			return fmt.Errorf("Get resource instance error: %s with resp code: %s", err, resp)
		}

		return nil
	}
}

func testAccCheckIBMResourceInstanceBasic(cos_instance_name string, cos_bucket_name string, cos_bucket_region string, vpc_name string, sg_ins_1 string, sg_ins_2 string, subnet_ins_1 string, subnet_ins_2 string, serviceName string, service_plan string) string {
	return fmt.Sprintf(`
	
    data "ibm_resource_group" "pag" {
		is_default = true
	  }
	  
	  data "ibm_resource_instance" "pag-cos" {
		name              = "%s"
		resource_group_id = data.ibm_resource_group.pag.id
		service           = "cloud-object-storage"
	  }
	  
	  data "ibm_cos_bucket" "pag-cos-bucket" {
		bucket_name          = "%s"
		resource_instance_id = data.ibm_resource_instance.pag-cos.id
		bucket_type          = "single_site_location"
		bucket_region        = "%s"
	  }
	  
	  
	  data "ibm_is_vpc" "pag" {
		name = "%s"
	  }

	  data "ibm_is_security_group" "pag_instance_1" {
		name     = "%s"
	  }
	  
	  data "ibm_is_security_group" "pag_instance_2" {
		name     = "%s"
	  }
	  
	  data "ibm_is_subnet" "pag_instance_1" {
		name = "%s"
	  }
	  
	  data "ibm_is_subnet" "pag_instance_2" {
		name = "%s"
	  }
	  	  
	  resource "ibm_pag_instance" "pag" {
		name              = "%s"
		resource_group_id = data.ibm_resource_group.pag.id
		service           = "privileged-access-gateway"
		plan              = "%s"
		location          = "us-south"
		parameters_json = jsonencode(
		  {
			"cosinstance" : data.ibm_resource_instance.pag-cos.crn,
			"cosbucket" : data.ibm_cos_bucket.pag-cos-bucket.bucket_name,
			"cosendpoint" : data.ibm_cos_bucket.pag-cos-bucket.s3_endpoint_direct
			"proxies" : [
			  {
				"name" : "proxy1",
				"securitygroups" : split(",", data.ibm_is_security_group.pag_instance_1.id),
				"subnet" : {
				  "crn" : data.ibm_is_subnet.pag_instance_1.crn,
				  "cidr" : data.ibm_is_subnet.pag_instance_1.ipv4_cidr_block
				}
			  },
			  {
				"name" : "proxy2",
				"securitygroups" : split(",", data.ibm_is_security_group.pag_instance_2.id),
				"subnet" : {
				  "crn" : data.ibm_is_subnet.pag_instance_2.crn,
				  "cidr" : data.ibm_is_subnet.pag_instance_2.ipv4_cidr_block
				}
			  }
			],
			"settings" : {
				"inactivity_timeout" : 15,
				"system_use_notification" : ""
			  },
			  "vpc_id" : data.ibm_is_vpc.pag.id
		  }
		)
		timeouts {
		  create = "30m"
		  update = "30m"
		  delete = "30m"
		}
	  }
	  
	  resource "ibm_iam_authorization_policy" "pag-cos-iam-policy" {
		source_service_name         = "privileged-access-gateway"
		source_resource_instance_id = ibm_pag_instance.pag.guid
		roles                       = ["Object Writer"]
		resource_attributes {
		  name     = "serviceName"
		  operator = "stringEquals"
		  value    = "cloud-object-storage"
		}
	  
		resource_attributes {
		  name     = "accountId"
		  operator = "stringEquals"
		  value    = data.ibm_resource_group.pag.account_id
		}
		resource_attributes {
		  name     = "serviceInstance"
		  operator = "stringEquals"
		  value    = data.ibm_resource_instance.pag-cos.guid
		}
	  
		resource_attributes {
		  name     = "resourceType"
		  operator = "stringEquals"
		  value    = "bucket"
		}
	  
		resource_attributes {
		  name     = "resource"
		  operator = "stringEquals"
		  value    = "cloud-object-storage-h2-cos-standard-bxo"
		}
	  
	  }

	`, cos_instance_name, cos_bucket_name, cos_bucket_region, vpc_name, sg_ins_1, sg_ins_2, subnet_ins_1, subnet_ins_2, serviceName, service_plan,
	)
}
