// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/scc-go-sdk/posturemanagementv2"
)

func TestAccIBMSccPostureV2CredentialsBasic(t *testing.T) {
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccPostureV2CredentialsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureV2CredentialsConfigBasic(enabled, typeVar, name, description, purpose),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccPostureV2CredentialsExists("ibm_scc_posture_v2_credentials.credentials", conf),
					resource.TestCheckResourceAttr("ibm_scc_posture_v2_credentials.credentials", "enabled", enabled),
					resource.TestCheckResourceAttr("ibm_scc_posture_v2_credentials.credentials", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_scc_posture_v2_credentials.credentials", "name", name),
					resource.TestCheckResourceAttr("ibm_scc_posture_v2_credentials.credentials", "description", description),
					resource.TestCheckResourceAttr("ibm_scc_posture_v2_credentials.credentials", "purpose", purpose),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSccPostureV2CredentialsConfigBasic(enabledUpdate, typeVarUpdate, nameUpdate, descriptionUpdate, purposeUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccPostureV2CredentialsExists("ibm_scc_posture_v2_credentials.credentials", conf),
					resource.TestCheckResourceAttr("ibm_scc_posture_v2_credentials.credentials", "enabled", enabledUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_v2_credentials.credentials", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_v2_credentials.credentials", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_v2_credentials.credentials", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_v2_credentials.credentials", "purpose", purposeUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_posture_v2_credentials.credentials",
				ImportState:       true,
				//ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSccPostureV2CredentialsConfigBasic(enabled string, typeVar string, name string, description string, purpose string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_posture_v2_credentials" "credentials" {
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

func testAccCheckIBMSccPostureV2CredentialsExists(n string, obj posturemanagementv2.Credential) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		postureManagementClient, err := testAccProvider.Meta().(ClientSession).PostureManagementV2()
		if err != nil {
			return err
		}

		listCredentialsOptions := &posturemanagementv2.ListCredentialsOptions{}
		listCredentialsOptions.SetAccountID(os.Getenv("SCC_POSTURE_V2_ACCOUNT_ID"))


		newCredential, _, err := postureManagementClient.ListCredentials(listCredentialsOptions)
		if err != nil {
			return err
		}
		fmt.Println(rs)
		obj = (newCredential.Credentials[0])
		return nil
	}
}

func testAccCheckIBMSccPostureV2CredentialsDestroy(s *terraform.State) error {
	postureManagementClient, err := testAccProvider.Meta().(ClientSession).PostureManagementV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_posture_v2_credentials" {
			continue
		}

		listCredentialsOptions := &posturemanagementv2.ListCredentialsOptions{}
		listCredentialsOptions.SetAccountID(os.Getenv("SCC_POSTURE_V2_ACCOUNT_ID"))

		// Try to find the key
		_, response, err := postureManagementClient.ListCredentials(listCredentialsOptions)

		if err == nil {
			return nil
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for credentials (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
