package power_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
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

func testAccCheckIBMPIWorkspaceConfig(name string) string {
	return fmt.Sprintf(`
	 resource "ibm_pi_workspace" "powervs_service_instance" {
		pi_name              = "%[1]s"
		pi_datacenter        = "dal"
		pi_resource_group_id = "%[2]s"
		pi_plan              = "public"
	  }
	`, name, acc.Pi_resource_group_id)
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
		client := st.NewIBMPIWorkspacesClient(context.Background(), sess, cloudInstanceID)
		workspace, resp, err := client.GetRC(cloudInstanceID)
		if err == nil {
			if *workspace.State == "active" {
				return fmt.Errorf("Resource Instance still exists: %s", rs.Primary.ID)
			}
		} else {
			if !strings.Contains(err.Error(), "404") {
				return fmt.Errorf("[ERROR] Error checking if Resource Instance (%s) has been destroyed: %s with resp code: %s", rs.Primary.ID, err, resp)
			}
		}
	}
	return nil
}

func testAccCheckIBMPIWorkspaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}

		cloudInstanceID := rs.Primary.ID
		client := st.NewIBMPIWorkspacesClient(context.Background(), sess, cloudInstanceID)
		_, _, err = client.GetRC(cloudInstanceID)
		if err != nil {
			return err
		}
		return nil
	}
}
