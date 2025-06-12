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
	var conf iampolicymanagementv1.PolicyTemplateMetaData

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

func TestAccIBMIAMAuthorizationPolicyUpdate_Basic(t *testing.T) {
	var conf iampolicymanagementv1.PolicyTemplateMetaData

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
			{
				Config: testAccCheckIBMIAMAuthorizationPolicyUpdate(acc.Tg_cross_network_account_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_account", acc.Tg_cross_network_account_id),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "description", "Authorization Policy for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicy_Resource_Instance(t *testing.T) {
	var conf iampolicymanagementv1.PolicyTemplateMetaData
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
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"transaction_id"},
			},
		},
	})
}

// TODO: Invalid authorizatoin header
func TestAccIBMIAMAuthorizationPolicy_Resource_Group(t *testing.T) {
	var conf iampolicymanagementv1.PolicyTemplateMetaData
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
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerifyIgnore: []string{"transaction_id"},
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicy_ResourceType(t *testing.T) {
	var conf iampolicymanagementv1.PolicyTemplateMetaData

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
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", "secrets-manager"),
				),
			},
		},
	})
}
func TestAccIBMIAMAuthorizationPolicyDelegatorRole(t *testing.T) {
	var conf iampolicymanagementv1.PolicyTemplateMetaData

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
	var conf iampolicymanagementv1.PolicyTemplateMetaData
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

func TestAccIBMIAMAuthorizationPolicy_SourceResourceGroupId(t *testing.T) {
	var conf iampolicymanagementv1.PolicyTemplateMetaData
	resourceName := "ibm_iam_authorization_policy.policy"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAuthorizationPolicySourceResourceGroupId(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttrSet("ibm_iam_authorization_policy.policy", "id"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", ""),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", "cloud-object-storage"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"transaction_id"},
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicy_SourceResourceGroupId_ResourceAttributes(t *testing.T) {
	var conf iampolicymanagementv1.PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAuthorizationPolicySourceResourceGroupIdResourceAttributes(acc.Tg_cross_network_account_id, acc.Tg_cross_network_account_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttrSet("ibm_iam_authorization_policy.policy", "id"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", ""),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", "cloud-object-storage"),
				),
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicy_SourceResourceGroupId_ResourceAttributes_WildCard(t *testing.T) {
	var conf iampolicymanagementv1.PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAuthorizationPolicySourceResourceGroupIdResourceAttributesWildCard(acc.Tg_cross_network_account_id, acc.Tg_cross_network_account_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttrSet("ibm_iam_authorization_policy.policy", "id"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", ""),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", "cloud-object-storage"),
				),
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicy_TargetResourceType(t *testing.T) {
	var conf iampolicymanagementv1.PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAuthorizationPolicyTargetResourceType(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", ""),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", "project"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_resource_type", "resource-group"),
				),
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicy_TargetResourceTypeAndResourceAttributes(t *testing.T) {
	var conf iampolicymanagementv1.PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAuthorizationPolicyResourceTypeAndResourceAttributes(acc.Tg_cross_network_account_id, acc.Tg_cross_network_account_id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", ""),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", "project"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_resource_type", "resource-group"),
				),
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicy_With_Transaction_id(t *testing.T) {
	var conf iampolicymanagementv1.PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAuthorizationPolicyTransactionId(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAuthorizationPolicyExists("ibm_iam_authorization_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_service_name", "databases-for-redis"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "transaction_id", "terrformAuthorizationPolicy"),
				),
			},
		},
	})
}

func TestAccIBMIAMAuthorizationPolicy_SourceResourceGroupIdWithStringExistsInSubjectAttributes(t *testing.T) {
	var conf iampolicymanagementv1.PolicyTemplateMetaData

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAuthorizationPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAuthorizationPolicySourceResourceGroupIdWithStringExistsInSubjectAttributes(acc.Tg_cross_network_account_id, acc.Tg_cross_network_account_id),
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

func testAccCheckIBMIAMAuthorizationPolicyExists(n string, obj iampolicymanagementv1.PolicyTemplateMetaData) resource.TestCheckFunc {

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

func testAccCheckIBMIAMAuthorizationPolicyUpdate(accountId string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_authorization_policy" "policy" {
		source_service_name = "cloud-object-storage"
		source_service_account = "%s"
		target_service_name = "kms"
		roles               = ["Reader"]
		description = "Authorization Policy for test scenario"
	  }
	`, accountId)
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
		source_resource_instance_id = ibm_resource_instance.instance1.guid
		target_service_name         = "kms"
		target_resource_instance_id = ibm_resource_instance.instance2.guid
		roles                       = ["Reader"]
	  }
	  
	`, instanceName, instanceName)
}

func testAccCheckIBMIAMAuthorizationPolicyResourceType() string {
	return `
	resource "ibm_iam_authorization_policy" "policy" {
		source_service_name  = "is"
		source_resource_type = "load-balancer"
		target_service_name  = "secrets-manager"
		roles                = ["SecretsReader"]
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
			value = ibm_resource_instance.cos.guid
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
			value = ibm_resource_instance.kms.guid
		}
	}
	`, sServiceInstance, tServiceInstance, sAccountID, tAccountID)
}

func testAccCheckIBMIAMAuthorizationPolicyTransactionId() string {
	return `
	resource "ibm_iam_authorization_policy" "policy" {
		source_service_name         = "databases-for-redis"
		target_service_name         = "kms"
		roles                       = ["Reader", "Authorization Delegator"]
		transaction_id 				= "terrformAuthorizationPolicy"
	  }
	`
}

func testAccCheckIBMIAMAuthorizationPolicySourceResourceGroupId() string {
	return fmt.Sprintf(`
	  resource "ibm_iam_authorization_policy" "policy" {
			source_resource_group_id    = "123-456-abc-def"
			target_service_name         = "cloud-object-storage"
			roles                       = ["Reader"]
	  }

	`)
}

func testAccCheckIBMIAMAuthorizationPolicySourceResourceGroupIdResourceAttributes(sAccountID, tAccountID string) string {

	return fmt.Sprintf(`

	resource "ibm_iam_authorization_policy" "policy" {
		roles    = ["Reader"]
		subject_attributes {
			name   = "accountId"
			value  = "%s"
		}
		subject_attributes {
			name   = "resourceGroupId"
			value  = "def-abc-456-123"
		}

		resource_attributes {
			name   = "serviceName"
			value  = "cloud-object-storage"
		}
		resource_attributes {
			name   = "accountId"
			value  = "%s"
		}
	}
	`, sAccountID, tAccountID)
}

func testAccCheckIBMIAMAuthorizationPolicySourceResourceGroupIdResourceAttributesWildCard(sAccountID, tAccountID string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_authorization_policy" "policy" {
		roles    = ["Reader"]
		subject_attributes {
			name   = "accountId"
			value  = "%s"
		}
		subject_attributes {
			name   = "resourceGroupId"
			value  = "*"
		}

		resource_attributes {
			name   = "serviceName"
			value  = "cloud-object-storage"
		}
		resource_attributes {
			name   = "accountId"
			value  = "%s"
		}
	}
	`, sAccountID, tAccountID)
}

func testAccCheckIBMIAMAuthorizationPolicyTargetResourceType() string {
	return `
	resource "ibm_iam_authorization_policy" "policy" {
		source_service_name = "project"
		target_resource_type  = "resource-group"
		roles                = ["Viewer"]
	  }
	`
}

func testAccCheckIBMIAMAuthorizationPolicyResourceTypeAndResourceAttributes(sAccountID, tAccountID string) string {

	return fmt.Sprintf(`

	resource "ibm_iam_authorization_policy" "policy" {
		roles    = ["Viewer"]
		subject_attributes {
			name   = "accountId"
			value  = "%s"
		}
		subject_attributes {
			name   = "serviceName"
			value  = "project"
		}

		resource_attributes {
			name   = "resourceType"
			value  = "resource-group"
		}
		resource_attributes {
			name   = "accountId"
			value  = "%s"
		}

	}
	`, sAccountID, tAccountID)
}

func testAccCheckIBMIAMAuthorizationPolicySourceResourceGroupIdWithStringExistsInSubjectAttributes(sAccountID, tAccountID string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_authorization_policy" "policy" {
		roles    = ["Reader"]
		subject_attributes {
			name   = "accountId"
			value  = "%s"
		}
		subject_attributes {
			name     = "resourceGroupId"
			operator = "stringExists"
			value    = "true"
		}

		resource_attributes {
			name   = "serviceName"
			value  = "cloud-object-storage"
		}
		resource_attributes {
			name   = "accountId"
			value  = "%s"
		}
	}
	`, sAccountID, tAccountID)
}
