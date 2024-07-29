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

func TestAccIbmSourceRegistrationBasic(t *testing.T) {
	var conf backuprecoveryv1.SourceRegistrationReponseParams
	environment := "kPhysical"
	environmentUpdate := "kOracle"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSourceRegistrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSourceRegistrationConfigBasic(environment),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSourceRegistrationExists("ibm_source_registration.source_registration_instance", conf),
					resource.TestCheckResourceAttr("ibm_source_registration.source_registration_instance", "environment", environment),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSourceRegistrationConfigBasic(environmentUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_source_registration.source_registration_instance", "environment", environmentUpdate),
				),
			},
		},
	})
}

func TestAccIbmSourceRegistrationAllArgs(t *testing.T) {
	var conf backuprecoveryv1.SourceRegistrationReponseParams
	environment := "kPhysical"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	connectionID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	connectorGroupID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	environmentUpdate := "kOracle"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	connectionIDUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	connectorGroupIDUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSourceRegistrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSourceRegistrationConfig(environment, name, connectionID, connectorGroupID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSourceRegistrationExists("ibm_source_registration.source_registration_instance", conf),
					resource.TestCheckResourceAttr("ibm_source_registration.source_registration_instance", "environment", environment),
					resource.TestCheckResourceAttr("ibm_source_registration.source_registration_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_source_registration.source_registration_instance", "connection_id", connectionID),
					resource.TestCheckResourceAttr("ibm_source_registration.source_registration_instance", "connector_group_id", connectorGroupID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSourceRegistrationConfig(environmentUpdate, nameUpdate, connectionIDUpdate, connectorGroupIDUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_source_registration.source_registration_instance", "environment", environmentUpdate),
					resource.TestCheckResourceAttr("ibm_source_registration.source_registration_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_source_registration.source_registration_instance", "connection_id", connectionIDUpdate),
					resource.TestCheckResourceAttr("ibm_source_registration.source_registration_instance", "connector_group_id", connectorGroupIDUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_source_registration.source_registration",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSourceRegistrationConfigBasic(environment string) string {
	return fmt.Sprintf(`
		resource "ibm_source_registration" "source_registration_instance" {
			environment = "%s"
		}
	`, environment)
}

func testAccCheckIbmSourceRegistrationConfig(environment string, name string, connectionID string, connectorGroupID string) string {
	return fmt.Sprintf(`

		resource "ibm_source_registration" "source_registration_instance" {
			environment = "%s"
			name = "%s"
			connection_id = %s
			connections {
				connection_id = 1
				entity_id = 1
				connector_group_id = 1
			}
			connector_group_id = %s
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
			oracle_params {
				database_entity_info {
					container_database_info {
						database_id = "database_id"
						database_name = "database_name"
					}
					data_guard_info {
						role = "kPrimary"
						standby_type = "kPhysical"
					}
				}
				host_info {
					id = "id"
					name = "name"
					environment = "kPhysical"
				}
			}
		}
	`, environment, name, connectionID, connectorGroupID)
}

func testAccCheckIbmSourceRegistrationExists(n string, obj backuprecoveryv1.SourceRegistrationReponseParams) resource.TestCheckFunc {

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

func testAccCheckIbmSourceRegistrationDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_source_registration" {
			continue
		}

		getProtectionSourceRegistrationOptions := &backuprecoveryv1.GetProtectionSourceRegistrationOptions{}

		num, _ := strconv.Atoi(rs.Primary.ID)

		getProtectionSourceRegistrationOptions.SetID(int64(num))

		// Try to find the key
		_, response, err := backupRecoveryClient.GetProtectionSourceRegistration(getProtectionSourceRegistrationOptions)

		if err == nil {
			return fmt.Errorf("source_registration still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for source_registration (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
