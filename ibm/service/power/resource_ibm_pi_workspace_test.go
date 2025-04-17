package power_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/power"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPIWorkspaceBasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-workspace-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMPIWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIWorkspaceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIWorkspaceExists("ibm_pi_workspace.powervs_service_instance"),
					resource.TestCheckResourceAttrSet("ibm_pi_workspace.powervs_service_instance", "id"),
				),
			},
		},
	})
}

func TestAccIBMPIWorkspaceUserTags(t *testing.T) {
	name := fmt.Sprintf("tf-pi-workspace-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccIBMPIWorkspaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIWorkspaceUserTagConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIWorkspaceExists("ibm_pi_workspace.powervs_service_instance"),
					resource.TestCheckResourceAttrSet("ibm_pi_workspace.powervs_service_instance", "id"),
					resource.TestCheckResourceAttr("ibm_pi_workspace.powervs_service_instance", "pi_user_tags.#", "2"),
					resource.TestCheckTypeSetElemAttr("ibm_pi_workspace.powervs_service_instance", "pi_user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr("ibm_pi_workspace.powervs_service_instance", "pi_user_tags.*", "dataresidency:france"),
				),
			},
		},
	})
}

func testAccCheckIBMPIWorkspaceConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_workspace" "powervs_service_instance" {
			pi_name              = "%[1]s"
			pi_datacenter        = "dal12"
			pi_resource_group_id = "%[2]s"
		}`, name, acc.Pi_resource_group_id)
}

func testAccCheckIBMPIWorkspaceUserTagConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_workspace" "powervs_service_instance" {
			pi_name              = "%[1]s"
			pi_datacenter        = "dal12"
			pi_resource_group_id = "%[2]s"
			pi_user_tags = ["env:dev", "dataresidency:france"]
		}`, name, acc.Pi_resource_group_id)
}

func testAccIBMPIWorkspaceDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_workspace" {
			continue
		}
		cloudInstanceID := rs.Primary.ID
		client := instance.NewIBMPIWorkspacesClient(context.Background(), sess, cloudInstanceID)
		workspace, resp, err := client.GetRC(cloudInstanceID)
		if err == nil {
			if *workspace.State == power.State_Active {
				return flex.FmtErrorf("Resource Instance still exists: %s", rs.Primary.ID)
			}
		} else {
			if !strings.Contains(err.Error(), "404") {
				return flex.FmtErrorf("[ERROR] Error checking if Resource Instance (%s) has been destroyed: %s with resp code: %s", rs.Primary.ID, err, resp)
			}
		}
	}
	return nil
}

func testAccCheckIBMPIWorkspaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return flex.FmtErrorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}

		cloudInstanceID := rs.Primary.ID
		client := instance.NewIBMPIWorkspacesClient(context.Background(), sess, cloudInstanceID)
		_, _, err = client.GetRC(cloudInstanceID)
		if err != nil {
			return err
		}
		return nil
	}
}
