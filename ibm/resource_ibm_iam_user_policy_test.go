// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package ibm

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMIAMUserPolicy_Basic(t *testing.T) {
	var conf iampolicymanagementv1.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyUpdateRole(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Service(t *testing.T) {
	var conf iampolicymanagementv1.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyService(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyUpdateServiceAndRegion(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.service", "cloudantnosqldb"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.region", "us-south"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_ResourceInstance(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyResourceInstance(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Resource_Group(t *testing.T) {
	var conf iampolicymanagementv1.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyResourceGroup(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.service", "containers-kubernetes"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Resource_Type(t *testing.T) {
	var conf iampolicymanagementv1.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyResourceType(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_import(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_user_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"resources", "resource_attributes"},
			},
		},
	})
}
func TestAccIBMIAMUserPolicy_With_Resource_Attributes(t *testing.T) {
	var conf iampolicymanagementv1.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMUserPolicyResourceAttributes(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resource_attributes.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMIAMUserPolicyResourceAttributesUpdate(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resource_attributes.#", "2"),
				),
			},
		},
	})
}
func TestAccIBMIAMUserPolicy_account_management(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_user_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyAccountManagement(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "account_management", "true"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_Invalid_User(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMIAMUserPolicyInvalidUser(),
				ExpectError: regexp.MustCompile(`User test@in.ibm.com is not found`),
			},
		},
	})
}

func TestAccIBMIAMUserPolicyWithCustomRole(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	crName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyWithCustomRole(crName, displayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMUserPolicyDestroy(s *terraform.State) error {
	rsContClient, err := testAccProvider.Meta().(ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_user_policy" {
			continue
		}
		policyID := rs.Primary.ID
		parts, err := idParts(policyID)
		if err != nil {
			return err
		}

		userPolicyID := parts[1]

		getPolicyOptions := rsContClient.NewGetPolicyOptions(
			userPolicyID,
		)

		// Try to find the key
		destroyedPolicy, response, err := rsContClient.GetPolicy(getPolicyOptions)

		if err == nil && *destroyedPolicy.State != "deleted" {
			return fmt.Errorf("User policy still exists: %s\n", rs.Primary.ID)
		} else if response.StatusCode != 404 && *destroyedPolicy.State != "deleted" {
			return fmt.Errorf("Error waiting for user policy (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMUserPolicyExists(n string, obj iampolicymanagementv1.Policy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := testAccProvider.Meta().(ClientSession).IAMPolicyManagementV1API()
		if err != nil {
			return err
		}

		policyID := rs.Primary.ID
		parts, err := idParts(policyID)
		if err != nil {
			return err
		}
		userPolicyID := parts[1]

		getPolicyOptions := rsContClient.NewGetPolicyOptions(
			userPolicyID,
		)

		policy, _, err := rsContClient.GetPolicy(getPolicyOptions)
		if err != nil {
			return err
		}
		obj = *policy
		return nil
	}
}

func testAccCheckIBMIAMUserPolicyBasic() string {
	return fmt.Sprintf(`

		  
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer"]
			tags   = ["tag1"]
	  	}

	`, IAMUser)
}

func testAccCheckIBMIAMUserPolicyUpdateRole() string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer", "Manager"]
			tags   = ["tag1", "tag2"]
	  	}
	`, IAMUser)
}

func testAccCheckIBMIAMUserPolicyService() string {
	return fmt.Sprintf(`

		
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer"]
	  
			resources {
		 		 service = "cloudantnosqldb"
			}
	  	}

	`, IAMUser)
}

func testAccCheckIBMIAMUserPolicyUpdateServiceAndRegion() string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_user_policy" "policy" {
			ibm_id 		 = "%s"
			roles        = ["Viewer", "Manager"]
		  
			resources {
			  service = "cloudantnosqldb"
			  region  = "us-south"
			}
		  }
	`, IAMUser)
}

func testAccCheckIBMIAMUserPolicyResourceInstance(name string) string {
	return fmt.Sprintf(`

		resource "ibm_resource_instance" "instance" {
			name     = "%s"
			service  = "kms"
			plan     = "tiered-pricing"
			location = "us-south"
	  	}
	  
	  	resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Manager", "Viewer", "Administrator"]
	  
			resources {
		  		service              = "kms"
		  		resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
			}
	  	}
		  

	`, name, IAMUser)
}

func testAccCheckIBMIAMUserPolicyResourceGroup() string {
	return fmt.Sprintf(`

		  
		data "ibm_resource_group" "group" {
			is_default=true
	  	}
	  
	  	resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer"]
	  
			resources {
		 	 service           = "containers-kubernetes"
		  	 resource_group_id = data.ibm_resource_group.group.id
			}
	  	}
		  

	`, IAMUser)
}

func testAccCheckIBMIAMUserPolicyResourceType() string {
	return fmt.Sprintf(`

		  
		data "ibm_resource_group" "group" {
			is_default=true
		  }
		  
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Administrator"]
		  
			resources {
			  resource_type = "resource-group"
			  resource      = data.ibm_resource_group.group.id
			}
		  }
	`, IAMUser)
}

// TODO: do we need this test? It follows pattern of other policies, but has conflict with existing policy
func testAccCheckIBMIAMUserPolicyImport(name string) string {
	return fmt.Sprintf(`

	
		  resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles        = ["Viewer"]
		  }

	`, IAMUser)
}

func testAccCheckIBMIAMUserPolicyInvalidUser() string {
	return fmt.Sprintf(`

		  
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "test@in.ibm.com"
			roles  = ["Viewer"]
	  	}

	`)
}

func testAccCheckIBMIAMUserPolicyAccountManagement(name string) string {
	return fmt.Sprintf(`
	
		  resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = ["Viewer"]
			account_management = true
		  }

	`, IAMUser)
}

func testAccCheckIBMIAMUserPolicyWithCustomRole(crName, displayName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_custom_role" "customrole" {
			name         = "%s"
			display_name = "%s"
			description  = "role for test scenario1"
			service = "kms"
			actions      = ["kms.secrets.rotate"]
		}
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles  = [ibm_iam_custom_role.customrole.display_name,"Viewer"]
			resources {
				service = "kms"
			  }
	  	}

	`, crName, displayName, IAMUser)
}

func testAccCheckIBMIAMUserPolicyResourceAttributes() string {
	return fmt.Sprintf(`
  
	  resource "ibm_iam_user_policy" "policy" {
		ibm_id = "%s"
		roles  = ["Viewer"]
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
	  
`, IAMUser)
}
func testAccCheckIBMIAMUserPolicyResourceAttributesUpdate() string {
	return fmt.Sprintf(`
	resource "ibm_iam_user_policy" "policy" {
		ibm_id = "%s"
		roles  = ["Viewer"]
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
	`, IAMUser)
}
