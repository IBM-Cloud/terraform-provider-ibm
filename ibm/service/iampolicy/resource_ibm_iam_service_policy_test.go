// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMIAMServicePolicy_Basic(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "description", "IAM Service Policy Creation for test scenario"),
				),
			},
			{
				Config: testAccCheckIBMIAMServicePolicyUpdateRole(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "description", "IAM Service Policy Update for test scenario"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Service(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyService(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMIAMServicePolicyUpdateServiceAndRegion(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.region", "us-south"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_ServiceType(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyServiceType(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service_type", "service"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_ResourceInstance(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyResourceInstance(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Resource_Group(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyResourceGroup(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service", "containers-kubernetes"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Resource_Type(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyResourceType(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_import(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_service_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"resources", "resource_attributes"},
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_account_management(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_service_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyAccountManagement(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "account_management", "true"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicyWithCustomRole(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	crName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyWithCustomRole(name, crName, displayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Resource_Attributes(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyResourceAttributes(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resource_attributes.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMIAMServicePolicyResourceAttributesUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resource_attributes.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Resource_Tags(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMServicePolicyResourceTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resource_tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "description", "IAM Service Policy Creation for test scenario"),
				),
			},
			{
				Config: testAccCheckIBMIAMServicePolicyUpdateResourceTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resource_tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "description", "IAM Service Policy Update for test scenario"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMServicePolicyDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_service_policy" {
			continue
		}
		policyID := rs.Primary.ID
		parts, err := flex.IdParts(policyID)
		if err != nil {
			return err
		}
		servicePolicyID := parts[1]

		getPolicyOptions := rsContClient.NewGetPolicyOptions(
			servicePolicyID,
		)

		// Try to find the key
		destroyedPolicy, response, err := rsContClient.GetPolicy(getPolicyOptions)

		if err == nil && *destroyedPolicy.State != "deleted" {
			return fmt.Errorf("User policy still exists: %s\n", rs.Primary.ID)
		} else if response.StatusCode != 404 && destroyedPolicy.State != nil && *destroyedPolicy.State != "deleted" {
			return fmt.Errorf("[ERROR] Error waiting for user policy (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMServicePolicyExists(n string, obj iampolicymanagementv1.Policy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
		if err != nil {
			return err
		}

		policyID := rs.Primary.ID

		parts, err := flex.IdParts(policyID)
		if err != nil {
			return err
		}
		servicePolicyID := parts[1]

		getPolicyOptions := rsContClient.NewGetPolicyOptions(
			servicePolicyID,
		)

		// Try to find the key
		policy, _, err := rsContClient.GetPolicy(getPolicyOptions)
		if err != nil {
			return err
		}
		obj = *policy
		return nil
	}
}

func testAccCheckIBMIAMServicePolicyBasic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
			tags           = ["tag1"]
			description    = "IAM Service Policy Creation for test scenario"
	  	}

	`, name)
}

func testAccCheckIBMIAMServicePolicyUpdateRole(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer", "Manager"]
			tags           = ["tag1", "tag2"]
			description    = "IAM Service Policy Update for test scenario"
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyService(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
	  
			resources {
		 	 service = "cloudantnosqldb"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyServiceType(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
	  
			resources {
				service_type = "service"
				region = "us-south"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyUpdateServiceAndRegion(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer", "Manager"]
	  
			resources {
		  		service = "cloudantnosqldb"
		  		region  = "us-south"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyResourceInstance(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_resource_instance" "instance" {
			name     = "%s"
			service  = "kms"
			plan     = "tiered-pricing"
			location = "us-south"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Manager", "Viewer", "Administrator"]
	  
			resources {
		 		 service              = "kms"
		  		resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
			}
	  	}
		  

	`, name, name)
}

func testAccCheckIBMIAMServicePolicyResourceGroup(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	data "ibm_resource_group" "group" {
			is_default=true
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
	  
			resources {
		 		service           = "containers-kubernetes"
		  		resource_group_id = data.ibm_resource_group.group.id
			}
	  	}
		  

	`, name)
}

func testAccCheckIBMIAMServicePolicyResourceType(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	data "ibm_resource_group" "group" {
			is_default=true
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Administrator"]
	  
			resources {
		  		resource_type = "resource-group"
		  		resource      = data.ibm_resource_group.group.id
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMServicePolicyImport(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
	  	}

	`, name)
}

func testAccCheckIBMIAMServicePolicyAccountManagement(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id     = ibm_iam_service_id.serviceID.id
			roles              = ["Viewer"]
			account_management = true
	  	}

	`, name)
}

func testAccCheckIBMIAMServicePolicyWithCustomRole(name, crName, displayName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
		}
		resource "ibm_iam_custom_role" "customrole" {
			name         = "%s"
			display_name = "%s"
			description  = "role for test scenario1"
			service = "kms"
			actions      = ["kms.secrets.rotate"]
		}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = [ibm_iam_custom_role.customrole.display_name,"Viewer"]
			tags           = ["tag1"]
			resources {
				service           = "kms"
		   }
	  	}

	`, name, crName, displayName)
}

func testAccCheckIBMIAMServicePolicyResourceAttributes(name string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_service_id" "serviceID" {
		name = "%s"
	  }
  
	  resource "ibm_iam_service_policy" "policy" {
		iam_service_id     = ibm_iam_service_id.serviceID.id
		roles              = ["Viewer"]
		resource_attributes {
			name     = "resource"
			value    = "test*"
			operator = "stringMatch"
		}
		resource_attributes {
			name     = "serviceName"
			value    = "messagehub"
		}
	  }
	`, name)
}
func testAccCheckIBMIAMServicePolicyResourceAttributesUpdate(name string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_service_id" "serviceID" {
		name = "%s"
	  }
  
	  resource "ibm_iam_service_policy" "policy" {
		iam_service_id     = ibm_iam_service_id.serviceID.id
		roles              = ["Viewer"]
		resource_attributes {
			name     = "resource"
			value    = "test*"
		}
		resource_attributes {
			name     = "serviceName"
			value    = "messagehub"
		}
	  }
	`, name)
}

func testAccCheckIBMIAMServicePolicyResourceTags(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
			
			resource_tags {
				name  = "one"
				value = "Terraform"
			}

			description    = "IAM Service Policy Creation for test scenario"
	  	}

	`, name)
}

func testAccCheckIBMIAMServicePolicyUpdateResourceTags(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_service_policy" "policy" {
			iam_service_id = ibm_iam_service_id.serviceID.id
			roles          = ["Viewer"]
			
			resource_tags {
				name  = "one"
				value = "Terraform"
			}
			resource_tags {
				name  = "two"
				value = "TerraformUpdate"
			}
			description    = "IAM Service Policy Update for test scenario"
	  	}

	`, name)
}
