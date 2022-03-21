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

func TestAccIBMIAMAccessGroupPolicy_Basic(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyUpdateRole(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_With_Service(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyService(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resources.0.service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyUpdateServiceAndRegion(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_With_ServiceType(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyServiceType(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resources.0.service_type", "service"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_With_ResourceInstance(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyResourceInstance(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_With_Resource_Group(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyResourceGroup(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resources.0.service", "containers-kubernetes"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_With_Resource_Type(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyResourceType(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_import(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_access_group_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyImport(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
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

func TestAccIBMIAMAccessGroupPolicy_account_management(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_access_group_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyAccountManagement(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "account_management", "true"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_With_Attributese(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyAttributes(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resources.0.service", "is"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_With_Resource_Attributes(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyResourceAttributes(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resource_attributes.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyResourceAttributesUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resource_attributes.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_With_Service_Specific_Roles(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyServiceSpecificRoles(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resource_attributes.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_WithCustomRole(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	crName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupPolicyWithCustomRole(name, crName, displayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupPolicy_With_Resource_Tags(t *testing.T) {
	var conf iampolicymanagementv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupPolicyResourceTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupPolicyExists("ibm_iam_access_group_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resource_tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupPolicyUpdateResourceTags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_access_group.accgrp", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "resource_tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMAccessGroupPolicyDestroy(s *terraform.State) error {
	iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_access_group_policy" {
			continue
		}
		policyID := rs.Primary.ID
		parts, err := flex.IdParts(policyID)
		if err != nil {
			return err
		}

		accessGroupPolicyID := parts[1]

		getPolicyOptions := iamPolicyManagementClient.NewGetPolicyOptions(
			accessGroupPolicyID,
		)

		destroyedPolicy, response, err := iamPolicyManagementClient.GetPolicy(getPolicyOptions)

		if err == nil && *destroyedPolicy.State != "deleted" {
			return fmt.Errorf("Access group policy still exists: %s\n", rs.Primary.ID)
		} else if response.StatusCode != 404 && destroyedPolicy.State != nil && *destroyedPolicy.State != "deleted" {
			return fmt.Errorf("[ERROR] Error waiting for access group policy (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMAccessGroupPolicyExists(n string, obj iampolicymanagementv1.Policy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
		if err != nil {
			return err
		}

		policyID := rs.Primary.ID

		parts, err := flex.IdParts(policyID)
		if err != nil {
			return err
		}

		accessGroupPolicyID := parts[1]

		getPolicyOptions := iamPolicyManagementClient.NewGetPolicyOptions(
			accessGroupPolicyID,
		)

		policy, _, err := iamPolicyManagementClient.GetPolicy(getPolicyOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error retrieving Policy %s err: %s", accessGroupPolicyID, err)
		}
		obj = *policy
		return nil
	}
}

func testAccCheckIBMIAMAccessGroupPolicyBasic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
  			name = "%s"
		}

		resource "ibm_iam_access_group_policy" "policy" {
  			access_group_id = ibm_iam_access_group.accgrp.id
  			roles           = ["Viewer"]
			tags            = ["tag1"]
		}

	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyUpdateRole(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Viewer", "Administrator"]
			tags            = ["tag1", "tag2"]
	  	}
	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyService(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
  		}

		resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles        = ["Viewer"]

			resources {
		  	service = "cloud-object-storage"
			}
		  }
		  
	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyServiceType(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
  		}

		resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles        = ["Viewer"]

			resources {
				service_type = "service"
				region = "us-south"
			}
		  }
		  
	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyUpdateServiceAndRegion(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Viewer", "Manager"]
	  
			resources {
		 	 service = "kms"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyResourceInstance(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	resource "ibm_resource_instance" "instance" {
			name     = "%s"
			service  = "kms"
			plan     = "tiered-pricing"
			location = "us-south"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Manager", "Viewer", "Administrator"]
	  
			resources {
		 	 service              = "kms"
		 	 resource_instance_id = element(split(":", ibm_resource_instance.instance.id), 7)
			}
	  	}
		  

	`, name, name)
}

func testAccCheckIBMIAMAccessGroupPolicyResourceGroup(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	data "ibm_resource_group" "group" {
			is_default=true
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Viewer"]
	  
			resources {
		 	 service           = "containers-kubernetes"
		 	 resource_group_id = data.ibm_resource_group.group.id
			}
	  	}
		  

	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyResourceType(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	data "ibm_resource_group" "group" {
			is_default=true
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Administrator"]
	  
			resources {
		  		resource_type = "resource-group"
		  		resource      = data.ibm_resource_group.group.id
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyImport(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	 	 }
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Viewer"]
	  	}

	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyAccountManagement(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id    = ibm_iam_access_group.accgrp.id
			roles              = ["Administrator"]
			account_management = true
	  	}

	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyAttributes(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Viewer"]
	  
			resources {
		  	service = "is"
		  	attributes = {
				"vpcId" = "*"
		  	}
			}
	  	}

	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyWithCustomRole(name, crName, displayName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
  			name = "%s"
		}

		resource "ibm_iam_custom_role" "customrole" {
			name         = "%s"
			display_name = "%s"
			description  = "role for test scenario1"
			service = "kms"
			actions      = ["kms.secrets.rotate"]
		}
		resource "ibm_iam_access_group_policy" "policy" {
  			access_group_id = ibm_iam_access_group.accgrp.id
  			roles           = [ibm_iam_custom_role.customrole.display_name,"Viewer"]
			  tags            = ["tag1"]
			  resources {
				service = "kms"
			  }
		}

	`, name, crName, displayName)
}
func testAccCheckIBMIAMAccessGroupPolicyResourceAttributes(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Viewer"]
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
func testAccCheckIBMIAMAccessGroupPolicyResourceAttributesUpdate(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Viewer"]
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

func testAccCheckIBMIAMAccessGroupPolicyServiceSpecificRoles(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_access_group" "accgrp" {
			name = "%s"
	  	}
	  
	  	resource "ibm_iam_access_group_policy" "policy" {
			access_group_id = ibm_iam_access_group.accgrp.id
			roles           = ["Satellite Link Source and Endpoint Controller"]
			resource_attributes {
				name     = "resource"
				value    = "test*"
			}
			resource_attributes {
				name     = "serviceName"
				value    = "satellite"
			}
	  	}
	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyResourceTags(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
  			name = "%s"
		}

		resource "ibm_iam_access_group_policy" "policy" {
  			access_group_id = ibm_iam_access_group.accgrp.id
  			roles           = ["Viewer"]
			
			resource_tags {
				name = "one"
				value = "terrformupdate"
			}
		}
	`, name)
}

func testAccCheckIBMIAMAccessGroupPolicyUpdateResourceTags(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group" "accgrp" {
  			name = "%s"
		}

		resource "ibm_iam_access_group_policy" "policy" {
  			access_group_id = ibm_iam_access_group.accgrp.id
  			roles           = ["Viewer"]
			
			resource_tags {
				name = "one"
				value = "terrformupdate"
			}

			resource_tags   {
				name = "two"
				value = "terrformupdate"
            }
		}
	`, name)
}
