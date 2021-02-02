// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
)

func TestAccIBMIAMAuthorizationPolicy_Basic(t *testing.T) {
	var conf iampapv1.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAuthorizationPolicyBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", "kms"),
				),
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicy_Resource_Instance(t *testing.T) {
	var conf iampapv1.Policy
	instanceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_authorization_policy.policy"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAuthorizationPolicyResourceInstance(instanceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", "kms"),
				),
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicy_Resource_Group(t *testing.T) {
	var conf iampapv1.Policy
	sResourceGroup := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	tResourceGroup := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_authorization_policy.policy"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAuthorizationPolicyResourceGroup(sResourceGroup, tResourceGroup),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", "kms"),
				),
			},
			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicy_ResourceType(t *testing.T) {
	var conf iampapv1.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAuthorizationPolicyResourceType(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", "is"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_resource_type", "load-balancer"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", "cloudcerts"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMAuthorizationPolicyDestroy(s *terraform.State) error {
	iampapClient, err := testAccProvider.Meta().(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_authorization_policy" {
			continue
		}

		authPolicyID := rs.Primary.ID

		err = iampapClient.V1Policy().Delete(authPolicyID)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for authorization policy (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMAuthorizationPolicyExists(n string, obj iampapv1.Policy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iampapClient, err := testAccProvider.Meta().(ClientSession).IAMPAPAPI()
		if err != nil {
			return err
		}

		authPolicyID := rs.Primary.ID

		policy, err := iampapClient.V1Policy().Get(authPolicyID)
		obj = policy
		return nil
	}
}

func testAccCheckIBMIAMAuthorizationPolicyBasic() string {
	return fmt.Sprintf(`
		  
	resource "ibm_iam_authorization_policy" "policy" {
		source_service_name = "cloud-object-storage"
		target_service_name = "kms"
		roles               = ["Reader"]
	  }
	`)
}

func testAccCheckIBMIAMAuthorizationPolicyResourceInstance(instanceName string) string {
	return fmt.Sprintf(`
		  
	resource "ibm_resource_instance" "instance1" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "standard"
		location = "global"
	  }
	  
	  resource "ibm_resource_instance" "instance2" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	  }
	  
	  resource "ibm_iam_authorization_policy" "policy" {
		source_service_name         = "cloud-object-storage"
		source_resource_instance_id = ibm_resource_instance.instance1.id
		target_service_name         = "kms"
		target_resource_instance_id = ibm_resource_instance.instance2.id
		roles                       = ["Reader"]
	  }
	  
	`, instanceName, instanceName)
}

func testAccCheckIBMIAMAuthorizationPolicyResourceType() string {
	return fmt.Sprintf(`
		  
	resource "ibm_iam_authorization_policy" "policy" {
		source_service_name  = "is"
		source_resource_type = "load-balancer"
		target_service_name  = "cloudcerts"
		roles                = ["Reader"]
	  }
	`)
}

func testAccCheckIBMIAMAuthorizationPolicyResourceGroup(sResourceGroup, tResourceGroup string) string {
	return fmt.Sprintf(`
		  
	resource "ibm_resource_group" "source_resource_group" {
		name     = "%s"
	  }
	  
	  resource "ibm_resource_group" "target_resource_group" {
		name     = "%s"
	  }
	  
	  resource "ibm_iam_authorization_policy" "policy" {
		source_service_name         = "cloud-object-storage"
		source_resource_group_id = ibm_resource_group.source_resource_group.id
		target_service_name         = "kms"
		target_resource_group_id = ibm_resource_group.target_resource_group.id
		roles                       = ["Reader"]
	  }
	  
	`, sResourceGroup, tResourceGroup)
}
