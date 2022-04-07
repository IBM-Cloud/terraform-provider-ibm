// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/scc-go-sdk/v3/posturemanagementv2"
)

func TestAccIBMSccPostureCredentialsBasic(t *testing.T) {
	var conf posturemanagementv2.Credential
	enabled := "true"
	typeVar := "ibm_cloud"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	purpose := "discovery_collection"
	enabledUpdate := "true"
	typeVarUpdate := "ibm_cloud"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	purposeUpdate := "discovery_fact_collection_remediation"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSccPostureCredentialsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSccPostureCredentialsConfigBasic(enabled, typeVar, name, description, purpose),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccPostureCredentialsExists("ibm_scc_posture_credential.credentials", conf),
					resource.TestCheckResourceAttr("ibm_scc_posture_credential.credentials", "enabled", enabled),
					resource.TestCheckResourceAttr("ibm_scc_posture_credential.credentials", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_scc_posture_credential.credentials", "name", name),
					resource.TestCheckResourceAttr("ibm_scc_posture_credential.credentials", "description", description),
					resource.TestCheckResourceAttr("ibm_scc_posture_credential.credentials", "purpose", purpose),
				),
			},
			{
				Config: testAccCheckIBMSccPostureCredentialsConfigBasic(enabledUpdate, typeVarUpdate, nameUpdate, descriptionUpdate, purposeUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccPostureCredentialsExists("ibm_scc_posture_credential.credentials", conf),
					resource.TestCheckResourceAttr("ibm_scc_posture_credential.credentials", "enabled", enabledUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_credential.credentials", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_credential.credentials", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_credential.credentials", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_credential.credentials", "purpose", purposeUpdate),
				),
			},
			{
				ResourceName: "ibm_scc_posture_credential.credentials",
				ImportState:  true,
				//ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSccPostureCredentialsConfigBasic(enabled string, typeVar string, name string, description string, purpose string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_posture_credential" "credentials" {
			enabled = %s
			type = "%s"
			name = "%s"
			description = "%s"
			display_fields {
				ibm_api_key = "sample_api_key"
				
			}
			group {
				id = "1"
				passphrase = "passphrase"
			}
			purpose = "%s"
		}
	`, enabled, typeVar, name, description, purpose)
}

func testAccCheckIBMSccPostureCredentialsExists(n string, obj posturemanagementv2.Credential) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		postureManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PostureManagementV2()
		if err != nil {
			return err
		}

		listCredentialsOptions := &posturemanagementv2.ListCredentialsOptions{}

		userDetails, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}

		listCredentialsOptions.SetAccountID(userDetails.UserAccount)

		newCredential, _, err := postureManagementClient.ListCredentials(listCredentialsOptions)
		if err != nil {
			return err
		}
		fmt.Println(rs)
		obj = (newCredential.Credentials[0])
		return nil
	}
}

func testAccCheckIBMSccPostureCredentialsDestroy(s *terraform.State) error {
	postureManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PostureManagementV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_posture_credential" {
			continue
		}

		listCredentialsOptions := &posturemanagementv2.ListCredentialsOptions{}

		userDetails, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}
		listCredentialsOptions.SetAccountID(userDetails.UserAccount)

		// Try to find the key
		_, response, err := postureManagementClient.ListCredentials(listCredentialsOptions)

		if err == nil {
			return nil
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for credentials (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
