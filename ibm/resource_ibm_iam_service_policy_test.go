package ibm

import (
	"fmt"
	"testing"

	"github.com/IBM-Cloud/bluemix-go/api/iam/iamv1"
	"github.com/IBM-Cloud/bluemix-go/models"

	"strings"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccIBMIAMServicePolicy_Basic(t *testing.T) {
	var conf models.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicy_basic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicy_updateRole(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "tags.#", "2"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_Service(t *testing.T) {
	var conf models.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicy_service(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists("ibm_iam_service_policy.policy", conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service", "cloud-object-storage"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicy_updateServiceAndRegion(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.service", "kms"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "resources.0.region", "us-south"),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicy_With_ResourceInstance(t *testing.T) {
	var conf models.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicy_resource_instance(name),
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
	var conf models.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicy_resource_group(name),
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
	var conf models.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicy_resource_type(name),
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
	var conf models.Policy
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())
	resourceName := "ibm_iam_service_policy.policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMIAMServicePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicy_import(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMServicePolicyExists(resourceName, conf),
					resource.TestCheckResourceAttr("ibm_iam_service_id.serviceID", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_service_policy.policy", "roles.#", "1"),
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

func testAccCheckIBMIAMServicePolicyDestroy(s *terraform.State) error {
	rsContClient, err := testAccProvider.Meta().(ClientSession).IAMAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_service_policy" {
			continue
		}
		policyID := rs.Primary.ID
		parts, err := idParts(policyID)
		if err != nil {
			return err
		}
		serviceIDUUID := parts[0]
		servicePolicyID := parts[1]

		bmxSess, err := testAccProvider.Meta().(ClientSession).BluemixSession()
		if err != nil {
			return err
		}

		mccpAPI, err := testAccProvider.Meta().(ClientSession).MccpAPI()
		if err != nil {
			return err
		}
		region, err := mccpAPI.Regions().FindRegionByName(bmxSess.Config.Region)
		if err != nil {
			return err
		}

		userDetails, err := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

		boundTo := GenerateBoundToCRN(*region, userDetails.userAccount)

		serviceID, err := rsContClient.ServiceIds().Get(serviceIDUUID)
		if err != nil {
			return nil
		}

		// Try to find the key
		err = rsContClient.ServicePolicies().Delete(iamv1.ServicePolicyIdentifier{
			Scope:    boundTo.ScopeSegment(),
			IAMID:    serviceID.IAMID,
			PolicyID: servicePolicyID,
		})

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for service policy (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMServicePolicyExists(n string, obj models.Policy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := testAccProvider.Meta().(ClientSession).IAMAPI()
		if err != nil {
			return err
		}

		policyID := rs.Primary.ID

		parts, err := idParts(policyID)
		if err != nil {
			return err
		}
		serviceIDUUID := parts[0]
		servicePolicyID := parts[1]
		bmxSess, err := testAccProvider.Meta().(ClientSession).BluemixSession()
		if err != nil {
			return err
		}

		mccpAPI, err := testAccProvider.Meta().(ClientSession).MccpAPI()
		if err != nil {
			return err
		}
		region, err := mccpAPI.Regions().FindRegionByName(bmxSess.Config.Region)
		if err != nil {
			return err
		}

		userDetails, err := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

		boundTo := GenerateBoundToCRN(*region, userDetails.userAccount)

		serviceID, err := rsContClient.ServiceIds().Get(serviceIDUUID)
		if err != nil {
			return err
		}

		// Try to find the key
		policy, err := rsContClient.ServicePolicies().Get(boundTo.ScopeSegment(), serviceID.IAMID, servicePolicyID)
		obj = policy
		return nil
	}
}

func testAccCheckIBMIAMServicePolicy_basic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
		  }
		  
		  resource "ibm_iam_service_policy" "policy" {
			iam_service_id = "${ibm_iam_service_id.serviceID.id}"
			roles        = ["Viewer"]
			tags         = ["tag1"]
		  }

	`, name)
}

func testAccCheckIBMIAMServicePolicy_updateRole(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
		  }
		  
		  resource "ibm_iam_service_policy" "policy" {
			iam_service_id = "${ibm_iam_service_id.serviceID.id}"
			roles        = ["Viewer","Manager"]
			tags         = ["tag1", "tag2"]
		  }
	`, name)
}

func testAccCheckIBMIAMServicePolicy_service(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
		  }
		  
		resource "ibm_iam_service_policy" "policy" {
			iam_service_id = "${ibm_iam_service_id.serviceID.id}"
			roles        = ["Viewer"]
		  
			resources = [{
			  service = "cloud-object-storage"
			}]
		  }

	`, name)
}

func testAccCheckIBMIAMServicePolicy_updateServiceAndRegion(name string) string {
	return fmt.Sprintf(`
		
		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
		  }
		  
		resource "ibm_iam_service_policy" "policy" {
			iam_service_id = "${ibm_iam_service_id.serviceID.id}"
			roles        = ["Viewer", "Manager"]
		  
			resources = [{
			  service = "kms"
			  region  = "us-south"
			}]
		  }
	`, name)
}

func testAccCheckIBMIAMServicePolicy_resource_instance(name string) string {
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
			iam_service_id = "${ibm_iam_service_id.serviceID.id}"
			roles        = ["Manager", "Viewer", "Administrator"]
		  
			resources = [{
			  service              = "kms"
			  region               = "us-south"
			  resource_instance_id = "${element(split(":",ibm_resource_instance.instance.id),7)}"
			}]
		  }
		  

	`, name, name)
}

func testAccCheckIBMIAMServicePolicy_resource_group(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
		  }
		  
		data "ibm_resource_group" "group" {
			name = "default"
		  }
		  
		resource "ibm_iam_service_policy" "policy" {
			iam_service_id = "${ibm_iam_service_id.serviceID.id}"
			roles        = ["Viewer"]
		  
			resources = [{
			  service           = "containers-kubernetes"
			  resource_group_id = "${data.ibm_resource_group.group.id}"
			}]
		  }
		  

	`, name)
}

func testAccCheckIBMIAMServicePolicy_resource_type(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
		  }
		  
		data "ibm_resource_group" "group" {
			name = "default"
		  }
		  
		resource "ibm_iam_service_policy" "policy" {
			iam_service_id = "${ibm_iam_service_id.serviceID.id}"
			roles        = ["Administrator"]
		  
			resources = [{
			  resource_type = "resource-group"
			  resource      = "${data.ibm_resource_group.group.id}"
			}]
		  }
	`, name)
}

func testAccCheckIBMIAMServicePolicy_import(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_service_id" "serviceID" {
			name = "%s"
		  }
		  
		  resource "ibm_iam_service_policy" "policy" {
			iam_service_id = "${ibm_iam_service_id.serviceID.id}"
			roles        = ["Viewer"]
		  }

	`, name)
}
