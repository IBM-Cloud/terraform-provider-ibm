// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasSourceRegistrationBasic(t *testing.T) {
	var conf backuprecoveryv1.SourceRegistrationReponseParams
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	environment := "kPhysical"
	xIbmTenantIDUpdate := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	environmentUpdate := "kSQL"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasSourceRegistrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasSourceRegistrationConfigBasic(xIbmTenantID, environment),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasSourceRegistrationExists("ibm_baas_source_registration.baas_source_registration_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "environment", environment),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmBaasSourceRegistrationConfigBasic(xIbmTenantIDUpdate, environmentUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "x_ibm_tenant_id", xIbmTenantIDUpdate),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "environment", environmentUpdate),
				),
			},
		},
	})
}

func TestAccIbmBaasSourceRegistrationAllArgs(t *testing.T) {
	var conf backuprecoveryv1.SourceRegistrationReponseParams
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	environment := "kPhysical"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	connectionID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	connectorGroupID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	dataSourceConnectionID := fmt.Sprintf("tf_data_source_connection_id_%d", acctest.RandIntRange(10, 100))
	xIbmTenantIDUpdate := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	environmentUpdate := "kSQL"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	connectionIDUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	connectorGroupIDUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	dataSourceConnectionIDUpdate := fmt.Sprintf("tf_data_source_connection_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasSourceRegistrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasSourceRegistrationConfig(xIbmTenantID, environment, name, connectionID, connectorGroupID, dataSourceConnectionID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasSourceRegistrationExists("ibm_baas_source_registration.baas_source_registration_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "environment", environment),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "connection_id", connectionID),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "connector_group_id", connectorGroupID),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "data_source_connection_id", dataSourceConnectionID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmBaasSourceRegistrationConfig(xIbmTenantIDUpdate, environmentUpdate, nameUpdate, connectionIDUpdate, connectorGroupIDUpdate, dataSourceConnectionIDUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "x_ibm_tenant_id", xIbmTenantIDUpdate),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "environment", environmentUpdate),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "connection_id", connectionIDUpdate),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "connector_group_id", connectorGroupIDUpdate),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "data_source_connection_id", dataSourceConnectionIDUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_source_registration.baas_source_registration",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasSourceRegistrationConfigBasic(xIbmTenantID string, environment string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_source_registration" "baas_source_registration_instance" {
			x_ibm_tenant_id = "%s"
			environment = "%s"
		}
	`, xIbmTenantID, environment)
}

func testAccCheckIbmBaasSourceRegistrationConfig(xIbmTenantID string, environment string, name string, connectionID string, connectorGroupID string, dataSourceConnectionID string) string {
	return fmt.Sprintf(`

		resource "ibm_baas_source_registration" "baas_source_registration_instance" {
			x_ibm_tenant_id = "%s"
			environment = "%s"
			name = "%s"
			connection_id = %s
			connections {
				connection_id = 1
				entity_id = 1
				connector_group_id = 1
				data_source_connection_id = "data_source_connection_id"
			}
			connector_group_id = %s
			data_source_connection_id = "%s"
			advanced_configs {
				key = "key"
				value = "value"
			}
			physical_params {
				endpoint = "endpoint"
				force_register = true
				host_type = "kLinux"
				physical_type = "kGroup"
				applications = [ "kSQL" ]
			}
		}
	`, xIbmTenantID, environment, name, connectionID, connectorGroupID, dataSourceConnectionID)
}

func testAccCheckIbmBaasSourceRegistrationExists(n string, obj backuprecoveryv1.SourceRegistrationReponseParams) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getProtectionSourceRegistrationOptions := &backuprecoveryv1.GetProtectionSourceRegistrationOptions{}

		num, _ := strconv.Atoi(rs.Primary.ID)
		getProtectionSourceRegistrationOptions.SetID(int64(num))

		sourceRegistrationReponseParams, _, err := backupRecoveryClient.GetProtectionSourceRegistration(getProtectionSourceRegistrationOptions)
		if err != nil {
			return err
		}

		obj = *sourceRegistrationReponseParams
		return nil
	}
}

func testAccCheckIbmBaasSourceRegistrationDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_source_registration" {
			continue
		}

		getProtectionSourceRegistrationOptions := &backuprecoveryv1.GetProtectionSourceRegistrationOptions{}

		num, _ := strconv.Atoi(rs.Primary.ID)

		getProtectionSourceRegistrationOptions.SetID(int64(num))

		// Try to find the key
		_, response, err := backupRecoveryClient.GetProtectionSourceRegistration(getProtectionSourceRegistrationOptions)

		if err == nil {
			return fmt.Errorf("baas_source_registration still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for baas_source_registration (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
