// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/scc-go-sdk/posturemanagementv2"
)

func TestAccIBMSccPostureScopesBasic(t *testing.T) {
	var conf posturemanagementv2.ScopeItem
	name := fmt.Sprintf("tf_name_%d", time.Now().UnixNano())
	description := fmt.Sprintf("tf_description_%d", time.Now().UnixNano())
	credentialType := "ibm"
	nameUpdate := fmt.Sprintf("tf_name_%d", time.Now().UnixNano())
	descriptionUpdate := fmt.Sprintf("tf_description_%d", time.Now().UnixNano())
	credentialTypeUpdate := "ibm"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccPostureScopesDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureScopesConfigBasic(name, description, scc_posture_credential_id_scope, credentialType, scc_posture_collector_id_scope),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccPostureScopesExists("ibm_scc_posture_scope.scopes", conf),
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "name", name),
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "description", description),
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "credential_id", scc_posture_credential_id_scope),
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "credential_type", credentialType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSccPostureScopesConfigBasic(nameUpdate, descriptionUpdate, scc_posture_credential_id_scope_update, credentialTypeUpdate, scc_posture_collector_id_scope_update),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "credential_id", scc_posture_credential_id_scope_update),
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "credential_type", credentialTypeUpdate),
				),
			},
		},
	})
}

func TestAccIBMScopesAllArgs(t *testing.T) {
	var conf posturemanagementv2.ScopeItem
	name := fmt.Sprintf("tf_name_%d", time.Now().UnixNano())
	description := fmt.Sprintf("tf_description_%d", time.Now().UnixNano())
	credentialType := "ibm"
	nameUpdate := fmt.Sprintf("tf_name_%d", time.Now().UnixNano())
	descriptionUpdate := fmt.Sprintf("tf_description_%d", time.Now().UnixNano())
	credentialTypeUpdate := "ibm"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSccPostureScopesDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureScopesConfig(name, description, scc_posture_credential_id_scope, credentialType, scc_posture_collector_id_scope),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccPostureScopesExists("ibm_scc_posture_scope.scopes", conf),
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "name", name),
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "description", description),
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "credential_id", scc_posture_credential_id_scope),
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "credential_type", credentialType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSccPostureScopesConfig(nameUpdate, descriptionUpdate, scc_posture_credential_id_scope_update, credentialTypeUpdate, scc_posture_collector_id_scope_update),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "credential_id", scc_posture_credential_id_scope_update),
					resource.TestCheckResourceAttr("ibm_scc_posture_scope.scopes", "credential_type", credentialTypeUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_posture_scope.scopes",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSccPostureScopesConfigBasic(name string, description string, credentialID string, credentialType string, collectorID []string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_posture_scope" "scopes" {
			name = "%s"
			description = "%s"
			credential_id = "%s"
			credential_type = "%s"
			collector_ids = %q
		}
	`, name, description, credentialID, credentialType, collectorID)
}

func testAccCheckIBMSccPostureScopesConfig(name string, description string, credentialID string, credentialType string, collectorID []string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_posture_scope" "scopes" {
			name = "%s"
			description = "%s"
			credential_id = "%s"
			credential_type = "%s"
			collector_ids = %q
		}
	`, name, description, credentialID, credentialType, collectorID)
}

func testAccCheckIBMSccPostureScopesExists(n string, obj posturemanagementv2.ScopeItem) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		postureManagementClient, err := testAccProvider.Meta().(ClientSession).PostureManagementV2()
		if err != nil {
			return err
		}

		listScopesOptions := &posturemanagementv2.ListScopesOptions{}
		listScopesOptions.SetAccountID(os.Getenv("SCC_POSTURE_ACCOUNT_ID"))

		newScope, _, err := postureManagementClient.ListScopes(listScopesOptions)
		if err != nil {
			return err
		}
		fmt.Println(rs)
		obj = (newScope.Scopes[0])
		return nil
	}
}

func testAccCheckIBMSccPostureScopesDestroy(s *terraform.State) error {
	postureManagementClient, err := testAccProvider.Meta().(ClientSession).PostureManagementV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_posture_scope" {
			continue
		}

		listScopesOptions := &posturemanagementv2.ListScopesOptions{}
		listScopesOptions.SetAccountID(os.Getenv("SCC_POSTURE_ACCOUNT_ID"))

		// Try to find the key
		_, response, err := postureManagementClient.ListScopes(listScopesOptions)

		if err == nil {
			return err
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scopes (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
