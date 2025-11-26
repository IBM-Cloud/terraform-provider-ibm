// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMContainerIngressInstance_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMContainerIngressInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerIngressInstanceBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_instance.instance", "secret_group_id", acc.SecretGroupID),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_instance.instance", "instance_crn", acc.InstanceCRN),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_instance.instance", "is_default", "true"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_instance.instance", "status", "created"),
				),
			},
			{
				Config: testAccCheckIBMContainerIngressInstanceUpdate(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_instance.instance", "instance_crn", acc.InstanceCRN),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_instance.instance", "is_default", "false"),
					resource.TestCheckResourceAttr(
						"ibm_container_ingress_instance.instance", "status", "created"),
				),
			},
			{
				ResourceName:      "ibm_container_ingress_instance.instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMContainerIngressInstanceDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_container_ingress_instance" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		clusterID := parts[0]
		instanceName := parts[1]

		ingressClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcContainerAPI()
		if err != nil {
			return err
		}

		ingressAPI := ingressClient.Ingresses()
		resp, err := ingressAPI.GetIngressInstance(clusterID, instanceName)
		if err == nil && &resp != nil && resp.Status == "deleted" {
			return nil
		} else if err == nil || !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMContainerIngressInstanceBasic() string {
	return fmt.Sprintf(`
resource "ibm_container_ingress_instance" "instance" {
  instance_crn    = "%s"
  secret_group_id = "%s"
  is_default = "%t"
  cluster  = "%s"
}`, acc.InstanceCRN, acc.SecretGroupID, true, acc.ClusterName)
}

func testAccCheckIBMContainerIngressInstanceUpdate() string {
	return fmt.Sprintf(`
resource "ibm_container_ingress_instance" "instance" {
  instance_crn    = "%s"
  is_default = "%t"
  cluster  = "%s"
}`, acc.InstanceCRN, false, acc.ClusterName)
}
