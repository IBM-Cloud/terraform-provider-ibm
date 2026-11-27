// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanagerinstancemanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func TestAccIbmSmHsmPathBuildBasic(t *testing.T) {
	var conf secretsmanagerinstancemanagementv2.Path
	instanceID := hsmInstanceID(t, "SM_HSM_PPS_GATEWAY_CRN", "SM_HSM_SERVER_CA_PEM_PATH")
	caName := hsmRandName()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmHsmPathBuildDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: hsmDestinationConfig(instanceID) + hsmPathBlock("tf_parent_path", "") + fmt.Sprintf(`
					resource "ibm_sm_hsm_server_ca" "ca" {
						instance_id = ibm_sm_hsm_path.test.instance_id
						destination_id = ibm_sm_hsm_path.test.destination_id
						path_id = ibm_sm_hsm_path.test.path_id
						name = "%s"
						certificate = file("%s")
						status = "active"
					}

					resource "ibm_sm_hsm_path_build" "test" {
						instance_id = ibm_sm_hsm_server_ca.ca.instance_id
						destination_id = ibm_sm_hsm_server_ca.ca.destination_id
						path_id = ibm_sm_hsm_server_ca.ca.path_id
					}
				`, caName, acc.SecretsManagerHsmServerCaPemPath),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmHsmPathBuildExists("ibm_sm_hsm_path_build.test", conf),
					resource.TestCheckResourceAttrPair("ibm_sm_hsm_path_build.test", "instance_id", "ibm_sm_hsm_path.test", "instance_id"),
					resource.TestCheckResourceAttrPair("ibm_sm_hsm_path_build.test", "destination_id", "ibm_sm_hsm_path.test", "destination_id"),
					resource.TestCheckResourceAttrPair("ibm_sm_hsm_path_build.test", "path_id", "ibm_sm_hsm_path.test", "path_id"),
					resource.TestCheckResourceAttrSet("ibm_sm_hsm_path_build.test", "provisioning_state"),
				),
			},
		},
	})
}

func testAccCheckIbmSmHsmPathBuildExists(n string, obj secretsmanagerinstancemanagementv2.Path) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		secretsManagerInstanceManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerInstanceManagementV2()
		if err != nil {
			return err
		}

		getInstanceHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getInstanceHsmDestinationPathOptions.SetInstanceID(parts[0])
		getInstanceHsmDestinationPathOptions.SetDestinationID(parts[1])
		getInstanceHsmDestinationPathOptions.SetPathID(parts[2])

		vpeProvisioningResponse, _, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptions)
		if err != nil {
			return err
		}

		obj = *vpeProvisioningResponse
		return nil
	}
}

func testAccCheckIbmSmHsmPathBuildDestroy(s *terraform.State) error {
	secretsManagerInstanceManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_hsm_path_build" {
			continue
		}

		getInstanceHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getInstanceHsmDestinationPathOptions.SetInstanceID(parts[0])
		getInstanceHsmDestinationPathOptions.SetDestinationID(parts[1])
		getInstanceHsmDestinationPathOptions.SetPathID(parts[2])

		// Try to find the key
		_, response, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptions)

		if err == nil {
			return fmt.Errorf("sm_hsm_path_build still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("error checking for sm_hsm_path_build (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
