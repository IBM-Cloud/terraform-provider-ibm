package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMIAMUserPolicyDataSource_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_user_policy.testacc_ds_user_policy", "policies.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMUserPolicyDataSource_Multiple_Policies(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserPolicyDataSourceMultiplePolicies(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_user_policy.testacc_ds_user_policy", "policies.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMUserPolicyDataSourceConfig(name string) string {
	return fmt.Sprintf(`

resource "ibm_resource_instance" "instance" {
	name     = "%s"
	service  = "kms"
	plan     = "tiered-pricing"
	location = "us-south"
}
  
resource "ibm_iam_user_policy" "policy" {
	ibm_id = "%s"
	roles        = ["Manager", "Viewer", "Administrator"]
  
	resources = [{
	  service              = "kms"
	  region               = "us-south"
	  resource_instance_id = "${element(split(":",ibm_resource_instance.instance.id),7)}"
	}]
	}
	
	data "ibm_iam_user_policy" "testacc_ds_user_policy" {
		ibm_id = "${ibm_iam_user_policy.policy.ibm_id}"
	}
`, name, IAMUser)

}

func testAccCheckIBMIAMUserPolicyDataSourceMultiplePolicies(name string) string {
	return fmt.Sprintf(`

resource "ibm_resource_instance" "instance" {
	name     = "%s"
	service  = "kms"
	plan     = "tiered-pricing"
	location = "us-south"
}

resource "ibm_iam_user_policy" "policy" {
	ibm_id = "%s"
	roles        = ["Manager", "Viewer", "Administrator"]
  
	resources = [{
	  service              = "kms"
	  region               = "us-south"
	  resource_instance_id = "${element(split(":",ibm_resource_instance.instance.id),7)}"
	}]
  }

  data "ibm_resource_group" "group" {
	name = "default"
  }
  
resource "ibm_iam_user_policy" "policy1" {
	ibm_id = "%s"
	roles        = ["Viewer"]
  
	resources = [{
	  service           = "containers-kubernetes"
	  resource_group_id = "${data.ibm_resource_group.group.id}"
	}]
  }


data "ibm_iam_user_policy" "testacc_ds_user_policy" {
	ibm_id = "${ibm_iam_user_policy.policy.ibm_id}"
}`, name, IAMUser, IAMUser)

}
