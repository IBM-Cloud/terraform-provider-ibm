// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func TestAccIBMIAMAuthorizationPolicy_Basic(t *testing.T) {
	var conf iampolicymanagementv1.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAuthorizationPolicyBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "description", "Authorization Policy for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicy_Resource_Instance(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	instanceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_authorization_policy.policy"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAuthorizationPolicyResourceInstance(instanceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", "kms"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// TODO: Invalid authorizatoin header
func TestAccIBMIAMAuthorizationPolicy_Resource_Group(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	sResourceGroup := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	tResourceGroup := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_authorization_policy.policy"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAuthorizationPolicyResourceGroup(sResourceGroup, tResourceGroup),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", "kms"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicy_ResourceType(t *testing.T) {
	var conf iampolicymanagementv1.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			{
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
func TestAccIBMIAMAuthorizationPolicyDelegatorRole(t *testing.T) {
	var conf iampolicymanagementv1.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAuthorizationPolicyDelegatorRole(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", "databases-for-redis"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", "kms"),
				),
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicy_ResourceAttributes(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	sServiceInstance := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	tServiceInstance := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAuthorizationPolicyResourceAttributes(sServiceInstance, tServiceInstance, acc.Tg_cross_network_account_id, acc.Tg_cross_network_account_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttrSet("ibm_iam_authorization_policy.policy", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMAuthorizationPolicyDestroy(s *terraform.State) error {
	iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_authorization_policy" {
			continue
		}

		authPolicyID := rs.Primary.ID

		getPolicyOptions := iamPolicyManagementClient.NewGetPolicyOptions(
			authPolicyID,
		)
		destroyedPolicy, response, err := iamPolicyManagementClient.GetPolicy(getPolicyOptions)

		if err == nil && *destroyedPolicy.State != "deleted" {
			return fmt.Errorf("Authorization policy still exists: %s\n", rs.Primary.ID)
		} else if response.StatusCode != 404 && destroyedPolicy.State != nil && *destroyedPolicy.State != "deleted" {
			return fmt.Errorf("[ERROR] Error waiting for authorization policy (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMAuthorizationPolicyExists(n string, obj iampolicymanagementv1.Policy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
		if err != nil {
			return err
		}

		authPolicyID := rs.Primary.ID

		getPolicyOptions := iamPolicyManagementClient.NewGetPolicyOptions(
			authPolicyID,
		)

		policy, resp, err := iamPolicyManagementClient.GetPolicy(getPolicyOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Getting Policy %s, %s", err, resp)
		}
		obj = *policy
		return nil
	}
}

func testAccCheckIBMIAMAuthorizationPolicyBasic() string {
	return `
	resource "ibm_iam_authorization_policy" "policy" {
		source_service_name = "cloud-object-storage"
		target_service_name = "kms"
		roles               = ["Reader"]
		description = "Authorization Policy for test scenario"
	  }
	`
}

func testAccCheckIBMIAMAuthorizationPolicyResourceInstance(instanceName string) string {
	return fmt.Sprintf(`
		  
	resource "ibm_resource_instance" "instance1" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "lite"
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
	return `
	resource "ibm_iam_authorization_policy" "policy" {
		source_service_name  = "is"
		source_resource_type = "load-balancer"
		target_service_name  = "cloudcerts"
		roles                = ["Reader"]
	  }
	`
}
func testAccCheckIBMIAMAuthorizationPolicyDelegatorRole() string {
	return `
	resource "ibm_iam_authorization_policy" "policy" {
		source_service_name         = "databases-for-redis"
		target_service_name         = "kms"
		roles                       = ["Reader", "Authorization Delegator"]
	  }
	`
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

func testAccCheckIBMIAMAuthorizationPolicyResourceAttributes(sServiceInstance, tServiceInstance, sAccountID, tAccountID string) string {

	return fmt.Sprintf(`
	
	resource "ibm_resource_instance" "cos" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "lite"
		location = "global"
	}
	
	resource "ibm_resource_instance" "kms" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	}

	resource "ibm_iam_authorization_policy" "policy" {
		roles                       = ["Reader"]

		subject_attributes {
			name   = "accountId"
			value = "%s"
		}
		subject_attributes {
			name   = "serviceInstance"
			value = ibm_resource_instance.cos.id
		}
		subject_attributes {
			name   = "serviceName"
			value = "cloud-object-storage"
		}

		resource_attributes {
			name   = "serviceName"
			value = "kms"
		}
		resource_attributes {
			name   = "accountId"
			value = "%s"
		}
		resource_attributes {
			name   = "serviceInstance"
			value = ibm_resource_instance.kms.id
		}
	}
	`, sServiceInstance, tServiceInstance, sAccountID, tAccountID)
}
