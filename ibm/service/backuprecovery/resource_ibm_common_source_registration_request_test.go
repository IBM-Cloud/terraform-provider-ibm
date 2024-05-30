// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

// func TestAccIbmCommonSourceRegistrationRequestBasic(t *testing.T) {
// 	var conf backuprecoveryv0.CommonSourceRegistrationRequestParams
// 	environment := "kPhysical"
// 	environmentUpdate := "kOracle"

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { acc.TestAccPreCheck(t) },
// 		Providers:    acc.TestAccProviders,
// 		CheckDestroy: testAccCheckIbmCommonSourceRegistrationRequestDestroy,
// 		Steps: []resource.TestStep{
// 			resource.TestStep{
// 				Config: testAccCheckIbmCommonSourceRegistrationRequestConfigBasic(environment),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					testAccCheckIbmCommonSourceRegistrationRequestExists("ibm_common_source_registration_request.common_source_registration_request_instance", conf),
// 					resource.TestCheckResourceAttr("ibm_common_source_registration_request.common_source_registration_request_instance", "environment", environment),
// 				),
// 			},
// 			resource.TestStep{
// 				Config: testAccCheckIbmCommonSourceRegistrationRequestConfigBasic(environmentUpdate),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr("ibm_common_source_registration_request.common_source_registration_request_instance", "environment", environmentUpdate),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccIbmCommonSourceRegistrationRequestAllArgs(t *testing.T) {
// 	var conf backuprecoveryv0.CommonSourceRegistrationRequestParams
// 	environment := "kPhysical"
// 	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
// 	isInternalEncrypted := "false"
// 	encryptionKey := fmt.Sprintf("tf_encryption_key_%d", acctest.RandIntRange(10, 100))
// 	connectionID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
// 	connectorGroupID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
// 	environmentUpdate := "kOracle"
// 	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
// 	isInternalEncryptedUpdate := "true"
// 	encryptionKeyUpdate := fmt.Sprintf("tf_encryption_key_%d", acctest.RandIntRange(10, 100))
// 	connectionIDUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
// 	connectorGroupIDUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { acc.TestAccPreCheck(t) },
// 		Providers:    acc.TestAccProviders,
// 		CheckDestroy: testAccCheckIbmCommonSourceRegistrationRequestDestroy,
// 		Steps: []resource.TestStep{
// 			resource.TestStep{
// 				Config: testAccCheckIbmCommonSourceRegistrationRequestConfig(environment, name, isInternalEncrypted, encryptionKey, connectionID, connectorGroupID),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					testAccCheckIbmCommonSourceRegistrationRequestExists("ibm_common_source_registration_request.common_source_registration_request_instance", conf),
// 					resource.TestCheckResourceAttr("ibm_common_source_registration_request.common_source_registration_request_instance", "environment", environment),
// 					resource.TestCheckResourceAttr("ibm_common_source_registration_request.common_source_registration_request_instance", "name", name),
// 					resource.TestCheckResourceAttr("ibm_common_source_registration_request.common_source_registration_request_instance", "is_internal_encrypted", isInternalEncrypted),
// 					resource.TestCheckResourceAttr("ibm_common_source_registration_request.common_source_registration_request_instance", "encryption_key", encryptionKey),
// 					resource.TestCheckResourceAttr("ibm_common_source_registration_request.common_source_registration_request_instance", "connection_id", connectionID),
// 					resource.TestCheckResourceAttr("ibm_common_source_registration_request.common_source_registration_request_instance", "connector_group_id", connectorGroupID),
// 				),
// 			},
// 			resource.TestStep{
// 				Config: testAccCheckIbmCommonSourceRegistrationRequestConfig(environmentUpdate, nameUpdate, isInternalEncryptedUpdate, encryptionKeyUpdate, connectionIDUpdate, connectorGroupIDUpdate),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestCheckResourceAttr("ibm_common_source_registration_request.common_source_registration_request_instance", "environment", environmentUpdate),
// 					resource.TestCheckResourceAttr("ibm_common_source_registration_request.common_source_registration_request_instance", "name", nameUpdate),
// 					resource.TestCheckResourceAttr("ibm_common_source_registration_request.common_source_registration_request_instance", "is_internal_encrypted", isInternalEncryptedUpdate),
// 					resource.TestCheckResourceAttr("ibm_common_source_registration_request.common_source_registration_request_instance", "encryption_key", encryptionKeyUpdate),
// 					resource.TestCheckResourceAttr("ibm_common_source_registration_request.common_source_registration_request_instance", "connection_id", connectionIDUpdate),
// 					resource.TestCheckResourceAttr("ibm_common_source_registration_request.common_source_registration_request_instance", "connector_group_id", connectorGroupIDUpdate),
// 				),
// 			},
// 			resource.TestStep{
// 				ResourceName:      "ibm_common_source_registration_request.common_source_registration_request",
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 			},
// 		},
// 	})
// }

// func testAccCheckIbmCommonSourceRegistrationRequestConfigBasic(environment string) string {
// 	return fmt.Sprintf(`
// 		resource "ibm_common_source_registration_request" "common_source_registration_request_instance" {
// 			environment = "%s"
// 		}
// 	`, environment)
// }

// func testAccCheckIbmCommonSourceRegistrationRequestConfig(environment string, name string, isInternalEncrypted string, encryptionKey string, connectionID string, connectorGroupID string) string {
// 	return fmt.Sprintf(`

// 		resource "ibm_common_source_registration_request" "common_source_registration_request_instance" {
// 			environment = "%s"
// 			name = "%s"
// 			is_internal_encrypted = %s
// 			encryption_key = "%s"
// 			connection_id = %s
// 			connections {
// 				connection_id = 1
// 				entity_id = 1
// 				connector_group_id = 1
// 			}
// 			connector_group_id = %s
// 			advanced_configs {
// 				key = "key"
// 				value = "value"
// 			}
// 			physical_params {
// 				endpoint = "endpoint"
// 				force_register = true
// 				host_type = "kLinux"
// 				physical_type = "kGroup"
// 				applications = [ "kSQL" ]
// 			}
// 			oracle_params {
// 				database_entity_info {
// 					container_database_info {
// 						database_id = "database_id"
// 						database_name = "database_name"
// 					}
// 					data_guard_info {
// 						role = "kPrimary"
// 						standby_type = "kPhysical"
// 					}
// 				}
// 				host_info {
// 					id = "id"
// 					name = "name"
// 					environment = "kPhysical"
// 				}
// 			}
// 		}
// 	`, environment, name, isInternalEncrypted, encryptionKey, connectionID, connectorGroupID)
// }

// func testAccCheckIbmCommonSourceRegistrationRequestExists(n string, obj backuprecoveryv0.CommonSourceRegistrationRequestParams) resource.TestCheckFunc {

// 	return func(s *terraform.State) error {
// 		rs, ok := s.RootModule().Resources[n]
// 		if !ok {
// 			return fmt.Errorf("Not found: %s", n)
// 		}

// 		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV0()
// 		if err != nil {
// 			return err
// 		}

// 		getSourceRegistrationsOptions := &backuprecoveryv0.GetSourceRegistrationsOptions{}

// 		getSourceRegistrationsOptions.SetID(rs.Primary.ID)

// 		commonSourceRegistrationRequestParams, _, err := backupRecoveryClient.GetSourceRegistrations(getSourceRegistrationsOptions)
// 		if err != nil {
// 			return err
// 		}

// 		obj = *commonSourceRegistrationRequestParams
// 		return nil
// 	}
// }

// func testAccCheckIbmCommonSourceRegistrationRequestDestroy(s *terraform.State) error {
// 	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV0()
// 	if err != nil {
// 		return err
// 	}
// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "ibm_common_source_registration_request" {
// 			continue
// 		}

// 		getSourceRegistrationsOptions := &backuprecoveryv0.GetSourceRegistrationsOptions{}

// 		getSourceRegistrationsOptions.SetID(rs.Primary.ID)

// 		// Try to find the key
// 		_, response, err := backupRecoveryClient.GetSourceRegistrations(getSourceRegistrationsOptions)

// 		if err == nil {
// 			return fmt.Errorf("common_source_registration_request still exists: %s", rs.Primary.ID)
// 		} else if response.StatusCode != 404 {
// 			return fmt.Errorf("Error checking for common_source_registration_request (%s) has been destroyed: %s", rs.Primary.ID, err)
// 		}
// 	}

// 	return nil
// }
