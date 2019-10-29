package ibm

import (
	"fmt"
	"testing"

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"

	"strings"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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
	instanceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
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
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "source_resource_type", "image"),
					resource.TestCheckResourceAttr("ibm_iam_authorization_policy.policy", "target_service_name", "cloud-object-storage"),
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
		name       = "%s"
		service    = "cloud-object-storage"
		plan       = "lite"
		location   = "global"
	  }
	  
	  resource "ibm_resource_instance" "instance2" {
		  name       = "%s"
		  service    = "kms"
		  plan       = "tiered-pricing"
		  location   = "us-south"
		}
	resource "ibm_iam_authorization_policy" "policy" {
		source_service_name = "cloud-object-storage"
		source_resource_instance_id = "${ibm_resource_instance.instance1.id}"
		target_service_name = "kms"
		target_resource_instance_id = "${ibm_resource_instance.instance2.id}"
		roles               = ["Reader"]
	  }
	  
	`, instanceName, instanceName)
}

func testAccCheckIBMIAMAuthorizationPolicyResourceType() string {
	return fmt.Sprintf(`
		  
	resource "ibm_iam_authorization_policy" "policy" {
		source_service_name = "is"
		source_resource_type = "image"
		target_service_name = "cloud-object-storage"
		roles               = ["Reader"]
	  }
	`)
}
