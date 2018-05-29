package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMIAMServicePolicyDataSource_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicyDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_service_policy.testacc_ds_service_policy", "policies.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMServicePolicyDataSource_Multiple_Policies(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServicePolicyDataSourceMultiplePolicies(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_service_policy.testacc_ds_service_policy", "policies.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMServicePolicyDataSourceConfig(name string) string {
	return fmt.Sprintf(`

resource "ibm_iam_service_id" "serviceID" {
  name       	= "%s"
  description	= "Service ID for test"
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

data "ibm_iam_service_policy" "testacc_ds_service_policy" {
	iam_service_id = "${ibm_iam_service_policy.policy.iam_service_id}"
}`, name, name)

}

func testAccCheckIBMIAMServicePolicyDataSourceMultiplePolicies(name string) string {
	return fmt.Sprintf(`

resource "ibm_iam_service_id" "serviceID" {
  name       	= "%s"
  description	= "Service ID for test"
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

  data "ibm_resource_group" "group" {
	name = "default"
  }
  
resource "ibm_iam_service_policy" "policy1" {
	iam_service_id = "${ibm_iam_service_id.serviceID.id}"
	roles        = ["Viewer"]
  
	resources = [{
	  service           = "containers-kubernetes"
	  resource_group_id = "${data.ibm_resource_group.group.id}"
	}]
  }


data "ibm_iam_service_policy" "testacc_ds_service_policy" {
	iam_service_id = "${ibm_iam_service_policy.policy.iam_service_id}"
}`, name, name)

}
