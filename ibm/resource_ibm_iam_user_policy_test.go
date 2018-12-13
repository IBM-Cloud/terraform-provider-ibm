package ibm

import (
	"fmt"
	"regexp"
	"testing"

	"strings"

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMIAMUserPolicy_Basic(t *testing.T) {
	var conf iampapv1.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicy_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicy_updateRole(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_Service(t *testing.T) {
	var conf iampapv1.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicy_service(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicy_updateServiceAndRegion(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "resources.0.region", "us-south"),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_With_ResourceInstance(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicy_resource_instance(name),
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
	var conf iampapv1.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicy_resource_group(),
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
	var conf iampapv1.Policy

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicy_resource_type(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists("ibm_iam_user_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicy_import(t *testing.T) {
	var conf iampapv1.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_iam_user_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMUserPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicy_import(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMUserPolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_user_policy.policy", "roles.#", "1"),
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

func TestAccIBMIAMUserPolicy_Invalid_User(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMIAMUserPolicy_invalid_user(),
				ExpectError: regexp.MustCompile(`User test@in.ibm.com does not exist within the given account.`),
			},
		},
	})
}

func testAccCheckIBMIAMUserPolicyDestroy(s *terraform.State) error {
	rsContClient, err := testAccProvider.Meta().(ClientSession).IAMPAPAPI()
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

		// Try to find the key
		_, err = rsContClient.V1Policy().Get(userPolicyID)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for user policy (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMUserPolicyExists(n string, obj iampapv1.Policy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := testAccProvider.Meta().(ClientSession).IAMPAPAPI()
		if err != nil {
			return err
		}

		policyID := rs.Primary.ID
		parts, err := idParts(policyID)
		if err != nil {
			return err
		}
		userPolicyID := parts[1]

		policy, err := rsContClient.V1Policy().Get(userPolicyID)
		if err != nil {
			return err
		}

		obj = policy
		return nil
	}
}

func testAccCheckIBMIAMUserPolicy_basic() string {
	return fmt.Sprintf(`

		  
		  resource "ibm_iam_user_policy" "policy" {
			ibm_id 		 = "%s"
			roles        = ["Viewer"]
			tags         = ["tag1"]
		  }

	`, IAMUser)
}

func testAccCheckIBMIAMUserPolicy_updateRole() string {
	return fmt.Sprintf(`
		
		  resource "ibm_iam_user_policy" "policy" {
			ibm_id 		 = "%s"
			roles        = ["Viewer","Manager"]
			tags         = ["tag1", "tag2"]
		  }
	`, IAMUser)
}

func testAccCheckIBMIAMUserPolicy_service() string {
	return fmt.Sprintf(`

		
		resource "ibm_iam_user_policy" "policy" {
			ibm_id 		 = "%s"
			roles        = ["Viewer"]
		  
			resources = [{
			  service = "cloud-object-storage"
			}]
		  }

	`, IAMUser)
}

func testAccCheckIBMIAMUserPolicy_updateServiceAndRegion() string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_user_policy" "policy" {
			ibm_id 		 = "%s"
			roles        = ["Viewer", "Manager"]
		  
			resources = [{
			  service = "kms"
			  region  = "us-south"
			}]
		  }
	`, IAMUser)
}

func testAccCheckIBMIAMUserPolicy_resource_instance(name string) string {
	return fmt.Sprintf(`

		resource "ibm_resource_instance" "instance" {
			name     = "%s"
			service  = "kms"
			plan     = "tiered-pricing"
			location = "us-south"
		  }
		  
		resource "ibm_iam_user_policy" "policy" {
			ibm_id 		 = "%s"
			roles        = ["Manager", "Viewer", "Administrator"]
		  
			resources = [{
			  service              = "kms"
			  region               = "us-south"
			  resource_instance_id = "${element(split(":",ibm_resource_instance.instance.id),7)}"
			}]
		  }
		  

	`, name, IAMUser)
}

func testAccCheckIBMIAMUserPolicy_resource_group() string {
	return fmt.Sprintf(`

		  
		data "ibm_resource_group" "group" {
			name = "default"
		  }
		  
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles        = ["Viewer"]
		  
			resources = [{
			  service           = "containers-kubernetes"
			  resource_group_id = "${data.ibm_resource_group.group.id}"
			}]
		  }
		  

	`, IAMUser)
}

func testAccCheckIBMIAMUserPolicy_resource_type() string {
	return fmt.Sprintf(`

		  
		data "ibm_resource_group" "group" {
			name = "default"
		  }
		  
		resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles        = ["Administrator"]
		  
			resources = [{
			  resource_type = "resource-group"
			  resource      = "${data.ibm_resource_group.group.id}"
			}]
		  }
	`, IAMUser)
}

func testAccCheckIBMIAMUserPolicy_import(name string) string {
	return fmt.Sprintf(`

	
		  resource "ibm_iam_user_policy" "policy" {
			ibm_id = "%s"
			roles        = ["Viewer"]
		  }

	`, IAMUser)
}

func testAccCheckIBMIAMUserPolicy_invalid_user() string {
	return fmt.Sprintf(`

		  
		  resource "ibm_iam_user_policy" "policy" {
			ibm_id 		 = "test@in.ibm.com"
			roles        = ["Viewer"]
		  }

	`)
}
