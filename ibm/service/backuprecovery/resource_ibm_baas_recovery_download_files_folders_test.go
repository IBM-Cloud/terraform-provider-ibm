// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasRecoveryDownloadFilesFoldersBasic(t *testing.T) {
	var conf backuprecoveryv1.Recovery
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasRecoveryDownloadFilesFoldersDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasRecoveryDownloadFilesFoldersConfigBasic(xIbmTenantID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasRecoveryDownloadFilesFoldersExists("ibm_baas_recovery_download_files_folders.baas_recovery_download_files_folders_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_recovery_download_files_folders.baas_recovery_download_files_folders_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_recovery_download_files_folders.baas_recovery_download_files_folders_instance", "name", name),
				),
			},
		},
	})
}

func TestAccIbmBaasRecoveryDownloadFilesFoldersAllArgs(t *testing.T) {
	var conf backuprecoveryv1.Recovery
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	parentRecoveryID := fmt.Sprintf("tf_parent_recovery_id_%d", acctest.RandIntRange(10, 100))
	glacierRetrievalType := "kStandard"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasRecoveryDownloadFilesFoldersDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasRecoveryDownloadFilesFoldersConfig(xIbmTenantID, name, parentRecoveryID, glacierRetrievalType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasRecoveryDownloadFilesFoldersExists("ibm_baas_recovery_download_files_folders.baas_recovery_download_files_folders_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_recovery_download_files_folders.baas_recovery_download_files_folders_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_recovery_download_files_folders.baas_recovery_download_files_folders_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_baas_recovery_download_files_folders.baas_recovery_download_files_folders_instance", "parent_recovery_id", parentRecoveryID),
					resource.TestCheckResourceAttr("ibm_baas_recovery_download_files_folders.baas_recovery_download_files_folders_instance", "glacier_retrieval_type", glacierRetrievalType),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_recovery_download_files_folders.baas_recovery_download_files_folders",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasRecoveryDownloadFilesFoldersConfigBasic(xIbmTenantID string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_recovery_download_files_folders" "baas_recovery_download_files_folders_instance" {
			x_ibm_tenant_id = "%s"
			name = "%s"
			object {
				snapshot_id = "snapshot_id"
				point_in_time_usecs = 1
				protection_group_id = "protection_group_id"
				protection_group_name = "protection_group_name"
				object_info {
					id = 1
					name = "name"
					source_id = 1
					source_name = "source_name"
					environment = "kPhysical"
					object_hash = "object_hash"
					object_type = "kCluster"
					logical_size_bytes = 1
					uuid = "uuid"
					global_id = "global_id"
					protection_type = "kAgent"
					sharepoint_site_summary {
						site_web_url = "site_web_url"
					}
					os_type = "kLinux"
					child_objects {
						id = 1
						name = "name"
						source_id = 1
						source_name = "source_name"
						environment = "kPhysical"
						object_hash = "object_hash"
						object_type = "kCluster"
						logical_size_bytes = 1
						uuid = "uuid"
						global_id = "global_id"
						protection_type = "kAgent"
						sharepoint_site_summary {
							site_web_url = "site_web_url"
						}
						os_type = "kLinux"
						v_center_summary {
							is_cloud_env = true
						}
						windows_cluster_summary {
							cluster_source_type = "cluster_source_type"
						}
					}
					v_center_summary {
						is_cloud_env = true
					}
					windows_cluster_summary {
						cluster_source_type = "cluster_source_type"
					}
				}
				archival_target_info {
					target_id = 1
					archival_task_id = "archival_task_id"
					target_name = "target_name"
					target_type = "Tape"
					usage_type = "Archival"
					ownership_context = "Local"
					tier_settings {
						aws_tiering {
							tiers {
								move_after_unit = "Days"
								move_after = 1
								tier_type = "kAmazonS3Standard"
							}
						}
						azure_tiering {
							tiers {
								move_after_unit = "Days"
								move_after = 1
								tier_type = "kAzureTierHot"
							}
						}
						cloud_platform = "AWS"
						google_tiering {
							tiers {
								move_after_unit = "Days"
								move_after = 1
								tier_type = "kGoogleStandard"
							}
						}
						oracle_tiering {
							tiers {
								move_after_unit = "Days"
								move_after = 1
								tier_type = "kOracleTierStandard"
							}
						}
						current_tier_type = "kAmazonS3Standard"
					}
				}
				recover_from_standby = true
			}
			files_and_folders {
				absolute_path = "absolute_path"
				is_directory = true
			}
		}
	`, xIbmTenantID, name)
}

func testAccCheckIbmBaasRecoveryDownloadFilesFoldersConfig(xIbmTenantID string, name string, parentRecoveryID string, glacierRetrievalType string) string {
	return fmt.Sprintf(`

		resource "ibm_baas_recovery_download_files_folders" "baas_recovery_download_files_folders_instance" {
			x_ibm_tenant_id = "%s"
			documents {
				is_directory = true
				item_id = "item_id"
			}
			name = "%s"
			object {
				snapshot_id = "snapshot_id"
				point_in_time_usecs = 1
				protection_group_id = "protection_group_id"
				protection_group_name = "protection_group_name"
				object_info {
					id = 1
					name = "name"
					source_id = 1
					source_name = "source_name"
					environment = "kPhysical"
					object_hash = "object_hash"
					object_type = "kCluster"
					logical_size_bytes = 1
					uuid = "uuid"
					global_id = "global_id"
					protection_type = "kAgent"
					sharepoint_site_summary {
						site_web_url = "site_web_url"
					}
					os_type = "kLinux"
					child_objects {
						id = 1
						name = "name"
						source_id = 1
						source_name = "source_name"
						environment = "kPhysical"
						object_hash = "object_hash"
						object_type = "kCluster"
						logical_size_bytes = 1
						uuid = "uuid"
						global_id = "global_id"
						protection_type = "kAgent"
						sharepoint_site_summary {
							site_web_url = "site_web_url"
						}
						os_type = "kLinux"
						v_center_summary {
							is_cloud_env = true
						}
						windows_cluster_summary {
							cluster_source_type = "cluster_source_type"
						}
					}
					v_center_summary {
						is_cloud_env = true
					}
					windows_cluster_summary {
						cluster_source_type = "cluster_source_type"
					}
				}
				archival_target_info {
					target_id = 1
					archival_task_id = "archival_task_id"
					target_name = "target_name"
					target_type = "Tape"
					usage_type = "Archival"
					ownership_context = "Local"
					tier_settings {
						aws_tiering {
							tiers {
								move_after_unit = "Days"
								move_after = 1
								tier_type = "kAmazonS3Standard"
							}
						}
						azure_tiering {
							tiers {
								move_after_unit = "Days"
								move_after = 1
								tier_type = "kAzureTierHot"
							}
						}
						cloud_platform = "AWS"
						google_tiering {
							tiers {
								move_after_unit = "Days"
								move_after = 1
								tier_type = "kGoogleStandard"
							}
						}
						oracle_tiering {
							tiers {
								move_after_unit = "Days"
								move_after = 1
								tier_type = "kOracleTierStandard"
							}
						}
						current_tier_type = "kAmazonS3Standard"
					}
				}
				recover_from_standby = true
			}
			parent_recovery_id = "%s"
			files_and_folders {
				absolute_path = "absolute_path"
				is_directory = true
			}
			glacier_retrieval_type = "%s"
		}
	`, xIbmTenantID, name, parentRecoveryID, glacierRetrievalType)
}

func testAccCheckIbmBaasRecoveryDownloadFilesFoldersExists(n string, obj backuprecoveryv1.Recovery) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getRecoveryByIdOptions := &backuprecoveryv1.GetRecoveryByIdOptions{}

		getRecoveryByIdOptions.SetID(rs.Primary.ID)

		downloadFilesAndFoldersRequestParams, _, err := backupRecoveryClient.GetRecoveryByID(getRecoveryByIdOptions)
		if err != nil {
			return err
		}

		obj = *downloadFilesAndFoldersRequestParams
		return nil
	}
}

func testAccCheckIbmBaasRecoveryDownloadFilesFoldersDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_recovery_download_files_folders" {
			continue
		}

		getRecoveryByIdOptions := &backuprecoveryv1.GetRecoveryByIdOptions{}

		getRecoveryByIdOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := backupRecoveryClient.GetRecoveryByID(getRecoveryByIdOptions)

		if err == nil {
			return fmt.Errorf("Download Files And Folders Recovery Params. still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Download Files And Folders Recovery Params. (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
