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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmRecoveryDownloadFilesFoldersBasic(t *testing.T) {
	var conf backuprecoveryv1.Recovery
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmRecoveryDownloadFilesFoldersDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmRecoveryDownloadFilesFoldersConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmRecoveryDownloadFilesFoldersExists("ibm_recovery_download_files_folders.recovery_download_files_folders_instance", conf),
					resource.TestCheckResourceAttr("ibm_recovery_download_files_folders.recovery_download_files_folders_instance", "name", name),
				),
			},
		},
	})
}

func TestAccIbmRecoveryDownloadFilesFoldersAllArgs(t *testing.T) {
	var conf backuprecoveryv1.Recovery
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	parentRecoveryID := fmt.Sprintf("tf_parent_recovery_id_%d", acctest.RandIntRange(10, 100))
	glacierRetrievalType := "kStandard"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmRecoveryDownloadFilesFoldersDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmRecoveryDownloadFilesFoldersConfig(name, parentRecoveryID, glacierRetrievalType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmRecoveryDownloadFilesFoldersExists("ibm_recovery_download_files_folders.recovery_download_files_folders_instance", conf),
					resource.TestCheckResourceAttr("ibm_recovery_download_files_folders.recovery_download_files_folders_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_recovery_download_files_folders.recovery_download_files_folders_instance", "parent_recovery_id", parentRecoveryID),
					resource.TestCheckResourceAttr("ibm_recovery_download_files_folders.recovery_download_files_folders_instance", "glacier_retrieval_type", glacierRetrievalType),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_recovery_download_files_folders.recovery_download_files_folders",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmRecoveryDownloadFilesFoldersConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_recovery_download_files_folders" "recovery_download_files_folders_instance" {
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
	`, name)
}

func testAccCheckIbmRecoveryDownloadFilesFoldersConfig(name string, parentRecoveryID string, glacierRetrievalType string) string {
	return fmt.Sprintf(`

		resource "ibm_recovery_download_files_folders" "recovery_download_files_folders_instance" {
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
	`, name, parentRecoveryID, glacierRetrievalType)
}

func testAccCheckIbmRecoveryDownloadFilesFoldersExists(n string, obj backuprecoveryv1.Recovery) resource.TestCheckFunc {

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

func testAccCheckIbmRecoveryDownloadFilesFoldersDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_recovery_download_files_folders" {
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

func TestResourceIbmRecoveryDownloadFilesFoldersDocumentObjectToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["is_directory"] = true
		model["item_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.DocumentObject)
	model.IsDirectory = core.BoolPtr(true)
	model.ItemID = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersDocumentObjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sharepointObjectParamsModel := make(map[string]interface{})
		sharepointObjectParamsModel["site_web_url"] = "testString"

		objectTypeVCenterParamsModel := make(map[string]interface{})
		objectTypeVCenterParamsModel["is_cloud_env"] = true

		objectTypeWindowsClusterParamsModel := make(map[string]interface{})
		objectTypeWindowsClusterParamsModel["cluster_source_type"] = "testString"

		objectSummaryModel := make(map[string]interface{})
		objectSummaryModel["id"] = int(26)
		objectSummaryModel["name"] = "testString"
		objectSummaryModel["source_id"] = int(26)
		objectSummaryModel["source_name"] = "testString"
		objectSummaryModel["environment"] = "kPhysical"
		objectSummaryModel["object_hash"] = "testString"
		objectSummaryModel["object_type"] = "kCluster"
		objectSummaryModel["logical_size_bytes"] = int(26)
		objectSummaryModel["uuid"] = "testString"
		objectSummaryModel["global_id"] = "testString"
		objectSummaryModel["protection_type"] = "kAgent"
		objectSummaryModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		objectSummaryModel["os_type"] = "kLinux"
		objectSummaryModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		objectSummaryModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		commonRecoverObjectSnapshotParamsObjectInfoModel := make(map[string]interface{})
		commonRecoverObjectSnapshotParamsObjectInfoModel["id"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["name"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["source_id"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["source_name"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["environment"] = "kPhysical"
		commonRecoverObjectSnapshotParamsObjectInfoModel["object_hash"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["object_type"] = "kCluster"
		commonRecoverObjectSnapshotParamsObjectInfoModel["logical_size_bytes"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["uuid"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["global_id"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["protection_type"] = "kAgent"
		commonRecoverObjectSnapshotParamsObjectInfoModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["os_type"] = "kLinux"
		commonRecoverObjectSnapshotParamsObjectInfoModel["child_objects"] = []map[string]interface{}{objectSummaryModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		awsTierModel := make(map[string]interface{})
		awsTierModel["move_after_unit"] = "Days"
		awsTierModel["move_after"] = int(26)
		awsTierModel["tier_type"] = "kAmazonS3Standard"

		awsTiersModel := make(map[string]interface{})
		awsTiersModel["tiers"] = []map[string]interface{}{awsTierModel}

		azureTierModel := make(map[string]interface{})
		azureTierModel["move_after_unit"] = "Days"
		azureTierModel["move_after"] = int(26)
		azureTierModel["tier_type"] = "kAzureTierHot"

		azureTiersModel := make(map[string]interface{})
		azureTiersModel["tiers"] = []map[string]interface{}{azureTierModel}

		googleTierModel := make(map[string]interface{})
		googleTierModel["move_after_unit"] = "Days"
		googleTierModel["move_after"] = int(26)
		googleTierModel["tier_type"] = "kGoogleStandard"

		googleTiersModel := make(map[string]interface{})
		googleTiersModel["tiers"] = []map[string]interface{}{googleTierModel}

		oracleTierModel := make(map[string]interface{})
		oracleTierModel["move_after_unit"] = "Days"
		oracleTierModel["move_after"] = int(26)
		oracleTierModel["tier_type"] = "kOracleTierStandard"

		oracleTiersModel := make(map[string]interface{})
		oracleTiersModel["tiers"] = []map[string]interface{}{oracleTierModel}

		archivalTargetTierInfoModel := make(map[string]interface{})
		archivalTargetTierInfoModel["aws_tiering"] = []map[string]interface{}{awsTiersModel}
		archivalTargetTierInfoModel["azure_tiering"] = []map[string]interface{}{azureTiersModel}
		archivalTargetTierInfoModel["cloud_platform"] = "AWS"
		archivalTargetTierInfoModel["google_tiering"] = []map[string]interface{}{googleTiersModel}
		archivalTargetTierInfoModel["oracle_tiering"] = []map[string]interface{}{oracleTiersModel}
		archivalTargetTierInfoModel["current_tier_type"] = "kAmazonS3Standard"

		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel := make(map[string]interface{})
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_id"] = int(26)
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["archival_task_id"] = "testString"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_name"] = "testString"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_type"] = "Tape"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["usage_type"] = "Archival"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["ownership_context"] = "Local"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["tier_settings"] = []map[string]interface{}{archivalTargetTierInfoModel}

		model := make(map[string]interface{})
		model["snapshot_id"] = "testString"
		model["point_in_time_usecs"] = int(26)
		model["protection_group_id"] = "testString"
		model["protection_group_name"] = "testString"
		model["snapshot_creation_time_usecs"] = int(26)
		model["object_info"] = []map[string]interface{}{commonRecoverObjectSnapshotParamsObjectInfoModel}
		model["snapshot_target_type"] = "Local"
		model["storage_domain_id"] = int(26)
		model["archival_target_info"] = []map[string]interface{}{commonRecoverObjectSnapshotParamsArchivalTargetInfoModel}
		model["progress_task_id"] = "testString"
		model["recover_from_standby"] = true
		model["status"] = "Accepted"
		model["start_time_usecs"] = int(26)
		model["end_time_usecs"] = int(26)
		model["messages"] = []string{"testString"}
		model["bytes_restored"] = int(26)

		assert.Equal(t, result, model)
	}

	sharepointObjectParamsModel := new(backuprecoveryv1.SharepointObjectParams)
	sharepointObjectParamsModel.SiteWebURL = core.StringPtr("testString")

	objectTypeVCenterParamsModel := new(backuprecoveryv1.ObjectTypeVCenterParams)
	objectTypeVCenterParamsModel.IsCloudEnv = core.BoolPtr(true)

	objectTypeWindowsClusterParamsModel := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	objectTypeWindowsClusterParamsModel.ClusterSourceType = core.StringPtr("testString")

	objectSummaryModel := new(backuprecoveryv1.ObjectSummary)
	objectSummaryModel.ID = core.Int64Ptr(int64(26))
	objectSummaryModel.Name = core.StringPtr("testString")
	objectSummaryModel.SourceID = core.Int64Ptr(int64(26))
	objectSummaryModel.SourceName = core.StringPtr("testString")
	objectSummaryModel.Environment = core.StringPtr("kPhysical")
	objectSummaryModel.ObjectHash = core.StringPtr("testString")
	objectSummaryModel.ObjectType = core.StringPtr("kCluster")
	objectSummaryModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	objectSummaryModel.UUID = core.StringPtr("testString")
	objectSummaryModel.GlobalID = core.StringPtr("testString")
	objectSummaryModel.ProtectionType = core.StringPtr("kAgent")
	objectSummaryModel.SharepointSiteSummary = sharepointObjectParamsModel
	objectSummaryModel.OsType = core.StringPtr("kLinux")
	objectSummaryModel.VCenterSummary = objectTypeVCenterParamsModel
	objectSummaryModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	commonRecoverObjectSnapshotParamsObjectInfoModel := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo)
	commonRecoverObjectSnapshotParamsObjectInfoModel.ID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.Name = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.SourceID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.SourceName = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.Environment = core.StringPtr("kPhysical")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ObjectHash = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ObjectType = core.StringPtr("kCluster")
	commonRecoverObjectSnapshotParamsObjectInfoModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.UUID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.GlobalID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ProtectionType = core.StringPtr("kAgent")
	commonRecoverObjectSnapshotParamsObjectInfoModel.SharepointSiteSummary = sharepointObjectParamsModel
	commonRecoverObjectSnapshotParamsObjectInfoModel.OsType = core.StringPtr("kLinux")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ChildObjects = []backuprecoveryv1.ObjectSummary{*objectSummaryModel}
	commonRecoverObjectSnapshotParamsObjectInfoModel.VCenterSummary = objectTypeVCenterParamsModel
	commonRecoverObjectSnapshotParamsObjectInfoModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	awsTierModel := new(backuprecoveryv1.AWSTier)
	awsTierModel.MoveAfterUnit = core.StringPtr("Days")
	awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
	awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

	awsTiersModel := new(backuprecoveryv1.AWSTiers)
	awsTiersModel.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

	azureTierModel := new(backuprecoveryv1.AzureTier)
	azureTierModel.MoveAfterUnit = core.StringPtr("Days")
	azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
	azureTierModel.TierType = core.StringPtr("kAzureTierHot")

	azureTiersModel := new(backuprecoveryv1.AzureTiers)
	azureTiersModel.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

	googleTierModel := new(backuprecoveryv1.GoogleTier)
	googleTierModel.MoveAfterUnit = core.StringPtr("Days")
	googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	googleTierModel.TierType = core.StringPtr("kGoogleStandard")

	googleTiersModel := new(backuprecoveryv1.GoogleTiers)
	googleTiersModel.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

	oracleTierModel := new(backuprecoveryv1.OracleTier)
	oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
	oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

	oracleTiersModel := new(backuprecoveryv1.OracleTiers)
	oracleTiersModel.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

	archivalTargetTierInfoModel := new(backuprecoveryv1.ArchivalTargetTierInfo)
	archivalTargetTierInfoModel.AwsTiering = awsTiersModel
	archivalTargetTierInfoModel.AzureTiering = azureTiersModel
	archivalTargetTierInfoModel.CloudPlatform = core.StringPtr("AWS")
	archivalTargetTierInfoModel.GoogleTiering = googleTiersModel
	archivalTargetTierInfoModel.OracleTiering = oracleTiersModel
	archivalTargetTierInfoModel.CurrentTierType = core.StringPtr("kAmazonS3Standard")

	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo)
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.ArchivalTaskID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetName = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetType = core.StringPtr("Tape")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.UsageType = core.StringPtr("Archival")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.OwnershipContext = core.StringPtr("Local")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TierSettings = archivalTargetTierInfoModel

	model := new(backuprecoveryv1.CommonRecoverObjectSnapshotParams)
	model.SnapshotID = core.StringPtr("testString")
	model.PointInTimeUsecs = core.Int64Ptr(int64(26))
	model.ProtectionGroupID = core.StringPtr("testString")
	model.ProtectionGroupName = core.StringPtr("testString")
	model.SnapshotCreationTimeUsecs = core.Int64Ptr(int64(26))
	model.ObjectInfo = commonRecoverObjectSnapshotParamsObjectInfoModel
	model.SnapshotTargetType = core.StringPtr("Local")
	model.StorageDomainID = core.Int64Ptr(int64(26))
	model.ArchivalTargetInfo = commonRecoverObjectSnapshotParamsArchivalTargetInfoModel
	model.ProgressTaskID = core.StringPtr("testString")
	model.RecoverFromStandby = core.BoolPtr(true)
	model.Status = core.StringPtr("Accepted")
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.Messages = []string{"testString"}
	model.BytesRestored = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsObjectInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sharepointObjectParamsModel := make(map[string]interface{})
		sharepointObjectParamsModel["site_web_url"] = "testString"

		objectTypeVCenterParamsModel := make(map[string]interface{})
		objectTypeVCenterParamsModel["is_cloud_env"] = true

		objectTypeWindowsClusterParamsModel := make(map[string]interface{})
		objectTypeWindowsClusterParamsModel["cluster_source_type"] = "testString"

		objectSummaryModel := make(map[string]interface{})
		objectSummaryModel["id"] = int(26)
		objectSummaryModel["name"] = "testString"
		objectSummaryModel["source_id"] = int(26)
		objectSummaryModel["source_name"] = "testString"
		objectSummaryModel["environment"] = "kPhysical"
		objectSummaryModel["object_hash"] = "testString"
		objectSummaryModel["object_type"] = "kCluster"
		objectSummaryModel["logical_size_bytes"] = int(26)
		objectSummaryModel["uuid"] = "testString"
		objectSummaryModel["global_id"] = "testString"
		objectSummaryModel["protection_type"] = "kAgent"
		objectSummaryModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		objectSummaryModel["os_type"] = "kLinux"
		objectSummaryModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		objectSummaryModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["source_id"] = int(26)
		model["source_name"] = "testString"
		model["environment"] = "kPhysical"
		model["object_hash"] = "testString"
		model["object_type"] = "kCluster"
		model["logical_size_bytes"] = int(26)
		model["uuid"] = "testString"
		model["global_id"] = "testString"
		model["protection_type"] = "kAgent"
		model["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		model["os_type"] = "kLinux"
		model["child_objects"] = []map[string]interface{}{objectSummaryModel}
		model["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		model["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		assert.Equal(t, result, model)
	}

	sharepointObjectParamsModel := new(backuprecoveryv1.SharepointObjectParams)
	sharepointObjectParamsModel.SiteWebURL = core.StringPtr("testString")

	objectTypeVCenterParamsModel := new(backuprecoveryv1.ObjectTypeVCenterParams)
	objectTypeVCenterParamsModel.IsCloudEnv = core.BoolPtr(true)

	objectTypeWindowsClusterParamsModel := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	objectTypeWindowsClusterParamsModel.ClusterSourceType = core.StringPtr("testString")

	objectSummaryModel := new(backuprecoveryv1.ObjectSummary)
	objectSummaryModel.ID = core.Int64Ptr(int64(26))
	objectSummaryModel.Name = core.StringPtr("testString")
	objectSummaryModel.SourceID = core.Int64Ptr(int64(26))
	objectSummaryModel.SourceName = core.StringPtr("testString")
	objectSummaryModel.Environment = core.StringPtr("kPhysical")
	objectSummaryModel.ObjectHash = core.StringPtr("testString")
	objectSummaryModel.ObjectType = core.StringPtr("kCluster")
	objectSummaryModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	objectSummaryModel.UUID = core.StringPtr("testString")
	objectSummaryModel.GlobalID = core.StringPtr("testString")
	objectSummaryModel.ProtectionType = core.StringPtr("kAgent")
	objectSummaryModel.SharepointSiteSummary = sharepointObjectParamsModel
	objectSummaryModel.OsType = core.StringPtr("kLinux")
	objectSummaryModel.VCenterSummary = objectTypeVCenterParamsModel
	objectSummaryModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	model := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.SourceID = core.Int64Ptr(int64(26))
	model.SourceName = core.StringPtr("testString")
	model.Environment = core.StringPtr("kPhysical")
	model.ObjectHash = core.StringPtr("testString")
	model.ObjectType = core.StringPtr("kCluster")
	model.LogicalSizeBytes = core.Int64Ptr(int64(26))
	model.UUID = core.StringPtr("testString")
	model.GlobalID = core.StringPtr("testString")
	model.ProtectionType = core.StringPtr("kAgent")
	model.SharepointSiteSummary = sharepointObjectParamsModel
	model.OsType = core.StringPtr("kLinux")
	model.ChildObjects = []backuprecoveryv1.ObjectSummary{*objectSummaryModel}
	model.VCenterSummary = objectTypeVCenterParamsModel
	model.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsObjectInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersSharepointObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["site_web_url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SharepointObjectParams)
	model.SiteWebURL = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersSharepointObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersObjectSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sharepointObjectParamsModel := make(map[string]interface{})
		sharepointObjectParamsModel["site_web_url"] = "testString"

		objectTypeVCenterParamsModel := make(map[string]interface{})
		objectTypeVCenterParamsModel["is_cloud_env"] = true

		objectTypeWindowsClusterParamsModel := make(map[string]interface{})
		objectTypeWindowsClusterParamsModel["cluster_source_type"] = "testString"

		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["source_id"] = int(26)
		model["source_name"] = "testString"
		model["environment"] = "kPhysical"
		model["object_hash"] = "testString"
		model["object_type"] = "kCluster"
		model["logical_size_bytes"] = int(26)
		model["uuid"] = "testString"
		model["global_id"] = "testString"
		model["protection_type"] = "kAgent"
		model["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		model["os_type"] = "kLinux"
		model["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		model["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		assert.Equal(t, result, model)
	}

	sharepointObjectParamsModel := new(backuprecoveryv1.SharepointObjectParams)
	sharepointObjectParamsModel.SiteWebURL = core.StringPtr("testString")

	objectTypeVCenterParamsModel := new(backuprecoveryv1.ObjectTypeVCenterParams)
	objectTypeVCenterParamsModel.IsCloudEnv = core.BoolPtr(true)

	objectTypeWindowsClusterParamsModel := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	objectTypeWindowsClusterParamsModel.ClusterSourceType = core.StringPtr("testString")

	model := new(backuprecoveryv1.ObjectSummary)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.SourceID = core.Int64Ptr(int64(26))
	model.SourceName = core.StringPtr("testString")
	model.Environment = core.StringPtr("kPhysical")
	model.ObjectHash = core.StringPtr("testString")
	model.ObjectType = core.StringPtr("kCluster")
	model.LogicalSizeBytes = core.Int64Ptr(int64(26))
	model.UUID = core.StringPtr("testString")
	model.GlobalID = core.StringPtr("testString")
	model.ProtectionType = core.StringPtr("kAgent")
	model.SharepointSiteSummary = sharepointObjectParamsModel
	model.OsType = core.StringPtr("kLinux")
	model.VCenterSummary = objectTypeVCenterParamsModel
	model.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersObjectSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersObjectTypeVCenterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["is_cloud_env"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeVCenterParams)
	model.IsCloudEnv = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersObjectTypeVCenterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersObjectTypeWindowsClusterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cluster_source_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	model.ClusterSourceType = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersObjectTypeWindowsClusterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsArchivalTargetInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		awsTierModel := make(map[string]interface{})
		awsTierModel["move_after_unit"] = "Days"
		awsTierModel["move_after"] = int(26)
		awsTierModel["tier_type"] = "kAmazonS3Standard"

		awsTiersModel := make(map[string]interface{})
		awsTiersModel["tiers"] = []map[string]interface{}{awsTierModel}

		azureTierModel := make(map[string]interface{})
		azureTierModel["move_after_unit"] = "Days"
		azureTierModel["move_after"] = int(26)
		azureTierModel["tier_type"] = "kAzureTierHot"

		azureTiersModel := make(map[string]interface{})
		azureTiersModel["tiers"] = []map[string]interface{}{azureTierModel}

		googleTierModel := make(map[string]interface{})
		googleTierModel["move_after_unit"] = "Days"
		googleTierModel["move_after"] = int(26)
		googleTierModel["tier_type"] = "kGoogleStandard"

		googleTiersModel := make(map[string]interface{})
		googleTiersModel["tiers"] = []map[string]interface{}{googleTierModel}

		oracleTierModel := make(map[string]interface{})
		oracleTierModel["move_after_unit"] = "Days"
		oracleTierModel["move_after"] = int(26)
		oracleTierModel["tier_type"] = "kOracleTierStandard"

		oracleTiersModel := make(map[string]interface{})
		oracleTiersModel["tiers"] = []map[string]interface{}{oracleTierModel}

		archivalTargetTierInfoModel := make(map[string]interface{})
		archivalTargetTierInfoModel["aws_tiering"] = []map[string]interface{}{awsTiersModel}
		archivalTargetTierInfoModel["azure_tiering"] = []map[string]interface{}{azureTiersModel}
		archivalTargetTierInfoModel["cloud_platform"] = "AWS"
		archivalTargetTierInfoModel["google_tiering"] = []map[string]interface{}{googleTiersModel}
		archivalTargetTierInfoModel["oracle_tiering"] = []map[string]interface{}{oracleTiersModel}
		archivalTargetTierInfoModel["current_tier_type"] = "kAmazonS3Standard"

		model := make(map[string]interface{})
		model["target_id"] = int(26)
		model["archival_task_id"] = "testString"
		model["target_name"] = "testString"
		model["target_type"] = "Tape"
		model["usage_type"] = "Archival"
		model["ownership_context"] = "Local"
		model["tier_settings"] = []map[string]interface{}{archivalTargetTierInfoModel}

		assert.Equal(t, result, model)
	}

	awsTierModel := new(backuprecoveryv1.AWSTier)
	awsTierModel.MoveAfterUnit = core.StringPtr("Days")
	awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
	awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

	awsTiersModel := new(backuprecoveryv1.AWSTiers)
	awsTiersModel.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

	azureTierModel := new(backuprecoveryv1.AzureTier)
	azureTierModel.MoveAfterUnit = core.StringPtr("Days")
	azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
	azureTierModel.TierType = core.StringPtr("kAzureTierHot")

	azureTiersModel := new(backuprecoveryv1.AzureTiers)
	azureTiersModel.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

	googleTierModel := new(backuprecoveryv1.GoogleTier)
	googleTierModel.MoveAfterUnit = core.StringPtr("Days")
	googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	googleTierModel.TierType = core.StringPtr("kGoogleStandard")

	googleTiersModel := new(backuprecoveryv1.GoogleTiers)
	googleTiersModel.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

	oracleTierModel := new(backuprecoveryv1.OracleTier)
	oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
	oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

	oracleTiersModel := new(backuprecoveryv1.OracleTiers)
	oracleTiersModel.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

	archivalTargetTierInfoModel := new(backuprecoveryv1.ArchivalTargetTierInfo)
	archivalTargetTierInfoModel.AwsTiering = awsTiersModel
	archivalTargetTierInfoModel.AzureTiering = azureTiersModel
	archivalTargetTierInfoModel.CloudPlatform = core.StringPtr("AWS")
	archivalTargetTierInfoModel.GoogleTiering = googleTiersModel
	archivalTargetTierInfoModel.OracleTiering = oracleTiersModel
	archivalTargetTierInfoModel.CurrentTierType = core.StringPtr("kAmazonS3Standard")

	model := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo)
	model.TargetID = core.Int64Ptr(int64(26))
	model.ArchivalTaskID = core.StringPtr("testString")
	model.TargetName = core.StringPtr("testString")
	model.TargetType = core.StringPtr("Tape")
	model.UsageType = core.StringPtr("Archival")
	model.OwnershipContext = core.StringPtr("Local")
	model.TierSettings = archivalTargetTierInfoModel

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersCommonRecoverObjectSnapshotParamsArchivalTargetInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersArchivalTargetTierInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		awsTierModel := make(map[string]interface{})
		awsTierModel["move_after_unit"] = "Days"
		awsTierModel["move_after"] = int(26)
		awsTierModel["tier_type"] = "kAmazonS3Standard"

		awsTiersModel := make(map[string]interface{})
		awsTiersModel["tiers"] = []map[string]interface{}{awsTierModel}

		azureTierModel := make(map[string]interface{})
		azureTierModel["move_after_unit"] = "Days"
		azureTierModel["move_after"] = int(26)
		azureTierModel["tier_type"] = "kAzureTierHot"

		azureTiersModel := make(map[string]interface{})
		azureTiersModel["tiers"] = []map[string]interface{}{azureTierModel}

		googleTierModel := make(map[string]interface{})
		googleTierModel["move_after_unit"] = "Days"
		googleTierModel["move_after"] = int(26)
		googleTierModel["tier_type"] = "kGoogleStandard"

		googleTiersModel := make(map[string]interface{})
		googleTiersModel["tiers"] = []map[string]interface{}{googleTierModel}

		oracleTierModel := make(map[string]interface{})
		oracleTierModel["move_after_unit"] = "Days"
		oracleTierModel["move_after"] = int(26)
		oracleTierModel["tier_type"] = "kOracleTierStandard"

		oracleTiersModel := make(map[string]interface{})
		oracleTiersModel["tiers"] = []map[string]interface{}{oracleTierModel}

		model := make(map[string]interface{})
		model["aws_tiering"] = []map[string]interface{}{awsTiersModel}
		model["azure_tiering"] = []map[string]interface{}{azureTiersModel}
		model["cloud_platform"] = "AWS"
		model["google_tiering"] = []map[string]interface{}{googleTiersModel}
		model["oracle_tiering"] = []map[string]interface{}{oracleTiersModel}
		model["current_tier_type"] = "kAmazonS3Standard"

		assert.Equal(t, result, model)
	}

	awsTierModel := new(backuprecoveryv1.AWSTier)
	awsTierModel.MoveAfterUnit = core.StringPtr("Days")
	awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
	awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

	awsTiersModel := new(backuprecoveryv1.AWSTiers)
	awsTiersModel.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

	azureTierModel := new(backuprecoveryv1.AzureTier)
	azureTierModel.MoveAfterUnit = core.StringPtr("Days")
	azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
	azureTierModel.TierType = core.StringPtr("kAzureTierHot")

	azureTiersModel := new(backuprecoveryv1.AzureTiers)
	azureTiersModel.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

	googleTierModel := new(backuprecoveryv1.GoogleTier)
	googleTierModel.MoveAfterUnit = core.StringPtr("Days")
	googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	googleTierModel.TierType = core.StringPtr("kGoogleStandard")

	googleTiersModel := new(backuprecoveryv1.GoogleTiers)
	googleTiersModel.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

	oracleTierModel := new(backuprecoveryv1.OracleTier)
	oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
	oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

	oracleTiersModel := new(backuprecoveryv1.OracleTiers)
	oracleTiersModel.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

	model := new(backuprecoveryv1.ArchivalTargetTierInfo)
	model.AwsTiering = awsTiersModel
	model.AzureTiering = azureTiersModel
	model.CloudPlatform = core.StringPtr("AWS")
	model.GoogleTiering = googleTiersModel
	model.OracleTiering = oracleTiersModel
	model.CurrentTierType = core.StringPtr("kAmazonS3Standard")

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersArchivalTargetTierInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersAWSTiersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		awsTierModel := make(map[string]interface{})
		awsTierModel["move_after_unit"] = "Days"
		awsTierModel["move_after"] = int(26)
		awsTierModel["tier_type"] = "kAmazonS3Standard"

		model := make(map[string]interface{})
		model["tiers"] = []map[string]interface{}{awsTierModel}

		assert.Equal(t, result, model)
	}

	awsTierModel := new(backuprecoveryv1.AWSTier)
	awsTierModel.MoveAfterUnit = core.StringPtr("Days")
	awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
	awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

	model := new(backuprecoveryv1.AWSTiers)
	model.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersAWSTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersAWSTierToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["move_after_unit"] = "Days"
		model["move_after"] = int(26)
		model["tier_type"] = "kAmazonS3Standard"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AWSTier)
	model.MoveAfterUnit = core.StringPtr("Days")
	model.MoveAfter = core.Int64Ptr(int64(26))
	model.TierType = core.StringPtr("kAmazonS3Standard")

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersAWSTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersAzureTiersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		azureTierModel := make(map[string]interface{})
		azureTierModel["move_after_unit"] = "Days"
		azureTierModel["move_after"] = int(26)
		azureTierModel["tier_type"] = "kAzureTierHot"

		model := make(map[string]interface{})
		model["tiers"] = []map[string]interface{}{azureTierModel}

		assert.Equal(t, result, model)
	}

	azureTierModel := new(backuprecoveryv1.AzureTier)
	azureTierModel.MoveAfterUnit = core.StringPtr("Days")
	azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
	azureTierModel.TierType = core.StringPtr("kAzureTierHot")

	model := new(backuprecoveryv1.AzureTiers)
	model.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersAzureTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersAzureTierToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["move_after_unit"] = "Days"
		model["move_after"] = int(26)
		model["tier_type"] = "kAzureTierHot"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AzureTier)
	model.MoveAfterUnit = core.StringPtr("Days")
	model.MoveAfter = core.Int64Ptr(int64(26))
	model.TierType = core.StringPtr("kAzureTierHot")

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersAzureTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersGoogleTiersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		googleTierModel := make(map[string]interface{})
		googleTierModel["move_after_unit"] = "Days"
		googleTierModel["move_after"] = int(26)
		googleTierModel["tier_type"] = "kGoogleStandard"

		model := make(map[string]interface{})
		model["tiers"] = []map[string]interface{}{googleTierModel}

		assert.Equal(t, result, model)
	}

	googleTierModel := new(backuprecoveryv1.GoogleTier)
	googleTierModel.MoveAfterUnit = core.StringPtr("Days")
	googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	googleTierModel.TierType = core.StringPtr("kGoogleStandard")

	model := new(backuprecoveryv1.GoogleTiers)
	model.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersGoogleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersGoogleTierToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["move_after_unit"] = "Days"
		model["move_after"] = int(26)
		model["tier_type"] = "kGoogleStandard"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.GoogleTier)
	model.MoveAfterUnit = core.StringPtr("Days")
	model.MoveAfter = core.Int64Ptr(int64(26))
	model.TierType = core.StringPtr("kGoogleStandard")

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersGoogleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersOracleTiersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		oracleTierModel := make(map[string]interface{})
		oracleTierModel["move_after_unit"] = "Days"
		oracleTierModel["move_after"] = int(26)
		oracleTierModel["tier_type"] = "kOracleTierStandard"

		model := make(map[string]interface{})
		model["tiers"] = []map[string]interface{}{oracleTierModel}

		assert.Equal(t, result, model)
	}

	oracleTierModel := new(backuprecoveryv1.OracleTier)
	oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
	oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

	model := new(backuprecoveryv1.OracleTiers)
	model.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersOracleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersOracleTierToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["move_after_unit"] = "Days"
		model["move_after"] = int(26)
		model["tier_type"] = "kOracleTierStandard"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.OracleTier)
	model.MoveAfterUnit = core.StringPtr("Days")
	model.MoveAfter = core.Int64Ptr(int64(26))
	model.TierType = core.StringPtr("kOracleTierStandard")

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersOracleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersFilesAndFoldersObjectToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["absolute_path"] = "testString"
		model["is_directory"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.FilesAndFoldersObject)
	model.AbsolutePath = core.StringPtr("testString")
	model.IsDirectory = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersFilesAndFoldersObjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.CommonRecoverObjectSnapshotParams) {
		sharepointObjectParamsModel := new(backuprecoveryv1.SharepointObjectParams)
		sharepointObjectParamsModel.SiteWebURL = core.StringPtr("testString")

		objectTypeVCenterParamsModel := new(backuprecoveryv1.ObjectTypeVCenterParams)
		objectTypeVCenterParamsModel.IsCloudEnv = core.BoolPtr(true)

		objectTypeWindowsClusterParamsModel := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
		objectTypeWindowsClusterParamsModel.ClusterSourceType = core.StringPtr("testString")

		objectSummaryModel := new(backuprecoveryv1.ObjectSummary)
		objectSummaryModel.ID = core.Int64Ptr(int64(26))
		objectSummaryModel.Name = core.StringPtr("testString")
		objectSummaryModel.SourceID = core.Int64Ptr(int64(26))
		objectSummaryModel.SourceName = core.StringPtr("testString")
		objectSummaryModel.Environment = core.StringPtr("kPhysical")
		objectSummaryModel.ObjectHash = core.StringPtr("testString")
		objectSummaryModel.ObjectType = core.StringPtr("kCluster")
		objectSummaryModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
		objectSummaryModel.UUID = core.StringPtr("testString")
		objectSummaryModel.GlobalID = core.StringPtr("testString")
		objectSummaryModel.ProtectionType = core.StringPtr("kAgent")
		objectSummaryModel.SharepointSiteSummary = sharepointObjectParamsModel
		objectSummaryModel.OsType = core.StringPtr("kLinux")
		objectSummaryModel.VCenterSummary = objectTypeVCenterParamsModel
		objectSummaryModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

		commonRecoverObjectSnapshotParamsObjectInfoModel := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo)
		commonRecoverObjectSnapshotParamsObjectInfoModel.ID = core.Int64Ptr(int64(26))
		commonRecoverObjectSnapshotParamsObjectInfoModel.Name = core.StringPtr("testString")
		commonRecoverObjectSnapshotParamsObjectInfoModel.SourceID = core.Int64Ptr(int64(26))
		commonRecoverObjectSnapshotParamsObjectInfoModel.SourceName = core.StringPtr("testString")
		commonRecoverObjectSnapshotParamsObjectInfoModel.Environment = core.StringPtr("kPhysical")
		commonRecoverObjectSnapshotParamsObjectInfoModel.ObjectHash = core.StringPtr("testString")
		commonRecoverObjectSnapshotParamsObjectInfoModel.ObjectType = core.StringPtr("kCluster")
		commonRecoverObjectSnapshotParamsObjectInfoModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
		commonRecoverObjectSnapshotParamsObjectInfoModel.UUID = core.StringPtr("testString")
		commonRecoverObjectSnapshotParamsObjectInfoModel.GlobalID = core.StringPtr("testString")
		commonRecoverObjectSnapshotParamsObjectInfoModel.ProtectionType = core.StringPtr("kAgent")
		commonRecoverObjectSnapshotParamsObjectInfoModel.SharepointSiteSummary = sharepointObjectParamsModel
		commonRecoverObjectSnapshotParamsObjectInfoModel.OsType = core.StringPtr("kLinux")
		commonRecoverObjectSnapshotParamsObjectInfoModel.ChildObjects = []backuprecoveryv1.ObjectSummary{*objectSummaryModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel.VCenterSummary = objectTypeVCenterParamsModel
		commonRecoverObjectSnapshotParamsObjectInfoModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

		awsTierModel := new(backuprecoveryv1.AWSTier)
		awsTierModel.MoveAfterUnit = core.StringPtr("Days")
		awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
		awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

		awsTiersModel := new(backuprecoveryv1.AWSTiers)
		awsTiersModel.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

		azureTierModel := new(backuprecoveryv1.AzureTier)
		azureTierModel.MoveAfterUnit = core.StringPtr("Days")
		azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
		azureTierModel.TierType = core.StringPtr("kAzureTierHot")

		azureTiersModel := new(backuprecoveryv1.AzureTiers)
		azureTiersModel.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

		googleTierModel := new(backuprecoveryv1.GoogleTier)
		googleTierModel.MoveAfterUnit = core.StringPtr("Days")
		googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
		googleTierModel.TierType = core.StringPtr("kGoogleStandard")

		googleTiersModel := new(backuprecoveryv1.GoogleTiers)
		googleTiersModel.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

		oracleTierModel := new(backuprecoveryv1.OracleTier)
		oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
		oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
		oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

		oracleTiersModel := new(backuprecoveryv1.OracleTiers)
		oracleTiersModel.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

		archivalTargetTierInfoModel := new(backuprecoveryv1.ArchivalTargetTierInfo)
		archivalTargetTierInfoModel.AwsTiering = awsTiersModel
		archivalTargetTierInfoModel.AzureTiering = azureTiersModel
		archivalTargetTierInfoModel.CloudPlatform = core.StringPtr("AWS")
		archivalTargetTierInfoModel.GoogleTiering = googleTiersModel
		archivalTargetTierInfoModel.OracleTiering = oracleTiersModel
		archivalTargetTierInfoModel.CurrentTierType = core.StringPtr("kAmazonS3Standard")

		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo)
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetID = core.Int64Ptr(int64(26))
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.ArchivalTaskID = core.StringPtr("testString")
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetName = core.StringPtr("testString")
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetType = core.StringPtr("Tape")
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.UsageType = core.StringPtr("Archival")
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.OwnershipContext = core.StringPtr("Local")
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TierSettings = archivalTargetTierInfoModel

		model := new(backuprecoveryv1.CommonRecoverObjectSnapshotParams)
		model.SnapshotID = core.StringPtr("testString")
		model.PointInTimeUsecs = core.Int64Ptr(int64(26))
		model.ProtectionGroupID = core.StringPtr("testString")
		model.ProtectionGroupName = core.StringPtr("testString")
		model.SnapshotCreationTimeUsecs = core.Int64Ptr(int64(26))
		model.ObjectInfo = commonRecoverObjectSnapshotParamsObjectInfoModel
		model.SnapshotTargetType = core.StringPtr("Local")
		model.StorageDomainID = core.Int64Ptr(int64(26))
		model.ArchivalTargetInfo = commonRecoverObjectSnapshotParamsArchivalTargetInfoModel
		model.ProgressTaskID = core.StringPtr("testString")
		model.RecoverFromStandby = core.BoolPtr(true)
		model.Status = core.StringPtr("Accepted")
		model.StartTimeUsecs = core.Int64Ptr(int64(26))
		model.EndTimeUsecs = core.Int64Ptr(int64(26))
		model.Messages = []string{"testString"}
		model.BytesRestored = core.Int64Ptr(int64(26))

		assert.Equal(t, result, model)
	}

	sharepointObjectParamsModel := make(map[string]interface{})
	sharepointObjectParamsModel["site_web_url"] = "testString"

	objectTypeVCenterParamsModel := make(map[string]interface{})
	objectTypeVCenterParamsModel["is_cloud_env"] = true

	objectTypeWindowsClusterParamsModel := make(map[string]interface{})
	objectTypeWindowsClusterParamsModel["cluster_source_type"] = "testString"

	objectSummaryModel := make(map[string]interface{})
	objectSummaryModel["id"] = int(26)
	objectSummaryModel["name"] = "testString"
	objectSummaryModel["source_id"] = int(26)
	objectSummaryModel["source_name"] = "testString"
	objectSummaryModel["environment"] = "kPhysical"
	objectSummaryModel["object_hash"] = "testString"
	objectSummaryModel["object_type"] = "kCluster"
	objectSummaryModel["logical_size_bytes"] = int(26)
	objectSummaryModel["uuid"] = "testString"
	objectSummaryModel["global_id"] = "testString"
	objectSummaryModel["protection_type"] = "kAgent"
	objectSummaryModel["sharepoint_site_summary"] = []interface{}{sharepointObjectParamsModel}
	objectSummaryModel["os_type"] = "kLinux"
	objectSummaryModel["v_center_summary"] = []interface{}{objectTypeVCenterParamsModel}
	objectSummaryModel["windows_cluster_summary"] = []interface{}{objectTypeWindowsClusterParamsModel}

	commonRecoverObjectSnapshotParamsObjectInfoModel := make(map[string]interface{})
	commonRecoverObjectSnapshotParamsObjectInfoModel["id"] = int(26)
	commonRecoverObjectSnapshotParamsObjectInfoModel["name"] = "testString"
	commonRecoverObjectSnapshotParamsObjectInfoModel["source_id"] = int(26)
	commonRecoverObjectSnapshotParamsObjectInfoModel["source_name"] = "testString"
	commonRecoverObjectSnapshotParamsObjectInfoModel["environment"] = "kPhysical"
	commonRecoverObjectSnapshotParamsObjectInfoModel["object_hash"] = "testString"
	commonRecoverObjectSnapshotParamsObjectInfoModel["object_type"] = "kCluster"
	commonRecoverObjectSnapshotParamsObjectInfoModel["logical_size_bytes"] = int(26)
	commonRecoverObjectSnapshotParamsObjectInfoModel["uuid"] = "testString"
	commonRecoverObjectSnapshotParamsObjectInfoModel["global_id"] = "testString"
	commonRecoverObjectSnapshotParamsObjectInfoModel["protection_type"] = "kAgent"
	commonRecoverObjectSnapshotParamsObjectInfoModel["sharepoint_site_summary"] = []interface{}{sharepointObjectParamsModel}
	commonRecoverObjectSnapshotParamsObjectInfoModel["os_type"] = "kLinux"
	commonRecoverObjectSnapshotParamsObjectInfoModel["child_objects"] = []interface{}{objectSummaryModel}
	commonRecoverObjectSnapshotParamsObjectInfoModel["v_center_summary"] = []interface{}{objectTypeVCenterParamsModel}
	commonRecoverObjectSnapshotParamsObjectInfoModel["windows_cluster_summary"] = []interface{}{objectTypeWindowsClusterParamsModel}

	awsTierModel := make(map[string]interface{})
	awsTierModel["move_after_unit"] = "Days"
	awsTierModel["move_after"] = int(26)
	awsTierModel["tier_type"] = "kAmazonS3Standard"

	awsTiersModel := make(map[string]interface{})
	awsTiersModel["tiers"] = []interface{}{awsTierModel}

	azureTierModel := make(map[string]interface{})
	azureTierModel["move_after_unit"] = "Days"
	azureTierModel["move_after"] = int(26)
	azureTierModel["tier_type"] = "kAzureTierHot"

	azureTiersModel := make(map[string]interface{})
	azureTiersModel["tiers"] = []interface{}{azureTierModel}

	googleTierModel := make(map[string]interface{})
	googleTierModel["move_after_unit"] = "Days"
	googleTierModel["move_after"] = int(26)
	googleTierModel["tier_type"] = "kGoogleStandard"

	googleTiersModel := make(map[string]interface{})
	googleTiersModel["tiers"] = []interface{}{googleTierModel}

	oracleTierModel := make(map[string]interface{})
	oracleTierModel["move_after_unit"] = "Days"
	oracleTierModel["move_after"] = int(26)
	oracleTierModel["tier_type"] = "kOracleTierStandard"

	oracleTiersModel := make(map[string]interface{})
	oracleTiersModel["tiers"] = []interface{}{oracleTierModel}

	archivalTargetTierInfoModel := make(map[string]interface{})
	archivalTargetTierInfoModel["aws_tiering"] = []interface{}{awsTiersModel}
	archivalTargetTierInfoModel["azure_tiering"] = []interface{}{azureTiersModel}
	archivalTargetTierInfoModel["cloud_platform"] = "AWS"
	archivalTargetTierInfoModel["google_tiering"] = []interface{}{googleTiersModel}
	archivalTargetTierInfoModel["oracle_tiering"] = []interface{}{oracleTiersModel}
	archivalTargetTierInfoModel["current_tier_type"] = "kAmazonS3Standard"

	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel := make(map[string]interface{})
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_id"] = int(26)
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["archival_task_id"] = "testString"
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_name"] = "testString"
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_type"] = "Tape"
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["usage_type"] = "Archival"
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["ownership_context"] = "Local"
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["tier_settings"] = []interface{}{archivalTargetTierInfoModel}

	model := make(map[string]interface{})
	model["snapshot_id"] = "testString"
	model["point_in_time_usecs"] = int(26)
	model["protection_group_id"] = "testString"
	model["protection_group_name"] = "testString"
	model["snapshot_creation_time_usecs"] = int(26)
	model["object_info"] = []interface{}{commonRecoverObjectSnapshotParamsObjectInfoModel}
	model["snapshot_target_type"] = "Local"
	model["storage_domain_id"] = int(26)
	model["archival_target_info"] = []interface{}{commonRecoverObjectSnapshotParamsArchivalTargetInfoModel}
	model["progress_task_id"] = "testString"
	model["recover_from_standby"] = true
	model["status"] = "Accepted"
	model["start_time_usecs"] = int(26)
	model["end_time_usecs"] = int(26)
	model["messages"] = []interface{}{"testString"}
	model["bytes_restored"] = int(26)

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParamsObjectInfo(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo) {
		sharepointObjectParamsModel := new(backuprecoveryv1.SharepointObjectParams)
		sharepointObjectParamsModel.SiteWebURL = core.StringPtr("testString")

		objectTypeVCenterParamsModel := new(backuprecoveryv1.ObjectTypeVCenterParams)
		objectTypeVCenterParamsModel.IsCloudEnv = core.BoolPtr(true)

		objectTypeWindowsClusterParamsModel := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
		objectTypeWindowsClusterParamsModel.ClusterSourceType = core.StringPtr("testString")

		objectSummaryModel := new(backuprecoveryv1.ObjectSummary)
		objectSummaryModel.ID = core.Int64Ptr(int64(26))
		objectSummaryModel.Name = core.StringPtr("testString")
		objectSummaryModel.SourceID = core.Int64Ptr(int64(26))
		objectSummaryModel.SourceName = core.StringPtr("testString")
		objectSummaryModel.Environment = core.StringPtr("kPhysical")
		objectSummaryModel.ObjectHash = core.StringPtr("testString")
		objectSummaryModel.ObjectType = core.StringPtr("kCluster")
		objectSummaryModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
		objectSummaryModel.UUID = core.StringPtr("testString")
		objectSummaryModel.GlobalID = core.StringPtr("testString")
		objectSummaryModel.ProtectionType = core.StringPtr("kAgent")
		objectSummaryModel.SharepointSiteSummary = sharepointObjectParamsModel
		objectSummaryModel.OsType = core.StringPtr("kLinux")
		objectSummaryModel.VCenterSummary = objectTypeVCenterParamsModel
		objectSummaryModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

		model := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo)
		model.ID = core.Int64Ptr(int64(26))
		model.Name = core.StringPtr("testString")
		model.SourceID = core.Int64Ptr(int64(26))
		model.SourceName = core.StringPtr("testString")
		model.Environment = core.StringPtr("kPhysical")
		model.ObjectHash = core.StringPtr("testString")
		model.ObjectType = core.StringPtr("kCluster")
		model.LogicalSizeBytes = core.Int64Ptr(int64(26))
		model.UUID = core.StringPtr("testString")
		model.GlobalID = core.StringPtr("testString")
		model.ProtectionType = core.StringPtr("kAgent")
		model.SharepointSiteSummary = sharepointObjectParamsModel
		model.OsType = core.StringPtr("kLinux")
		model.ChildObjects = []backuprecoveryv1.ObjectSummary{*objectSummaryModel}
		model.VCenterSummary = objectTypeVCenterParamsModel
		model.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

		assert.Equal(t, result, model)
	}

	sharepointObjectParamsModel := make(map[string]interface{})
	sharepointObjectParamsModel["site_web_url"] = "testString"

	objectTypeVCenterParamsModel := make(map[string]interface{})
	objectTypeVCenterParamsModel["is_cloud_env"] = true

	objectTypeWindowsClusterParamsModel := make(map[string]interface{})
	objectTypeWindowsClusterParamsModel["cluster_source_type"] = "testString"

	objectSummaryModel := make(map[string]interface{})
	objectSummaryModel["id"] = int(26)
	objectSummaryModel["name"] = "testString"
	objectSummaryModel["source_id"] = int(26)
	objectSummaryModel["source_name"] = "testString"
	objectSummaryModel["environment"] = "kPhysical"
	objectSummaryModel["object_hash"] = "testString"
	objectSummaryModel["object_type"] = "kCluster"
	objectSummaryModel["logical_size_bytes"] = int(26)
	objectSummaryModel["uuid"] = "testString"
	objectSummaryModel["global_id"] = "testString"
	objectSummaryModel["protection_type"] = "kAgent"
	objectSummaryModel["sharepoint_site_summary"] = []interface{}{sharepointObjectParamsModel}
	objectSummaryModel["os_type"] = "kLinux"
	objectSummaryModel["v_center_summary"] = []interface{}{objectTypeVCenterParamsModel}
	objectSummaryModel["windows_cluster_summary"] = []interface{}{objectTypeWindowsClusterParamsModel}

	model := make(map[string]interface{})
	model["id"] = int(26)
	model["name"] = "testString"
	model["source_id"] = int(26)
	model["source_name"] = "testString"
	model["environment"] = "kPhysical"
	model["object_hash"] = "testString"
	model["object_type"] = "kCluster"
	model["logical_size_bytes"] = int(26)
	model["uuid"] = "testString"
	model["global_id"] = "testString"
	model["protection_type"] = "kAgent"
	model["sharepoint_site_summary"] = []interface{}{sharepointObjectParamsModel}
	model["os_type"] = "kLinux"
	model["child_objects"] = []interface{}{objectSummaryModel}
	model["v_center_summary"] = []interface{}{objectTypeVCenterParamsModel}
	model["windows_cluster_summary"] = []interface{}{objectTypeWindowsClusterParamsModel}

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParamsObjectInfo(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToSharepointObjectParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.SharepointObjectParams) {
		model := new(backuprecoveryv1.SharepointObjectParams)
		model.SiteWebURL = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["site_web_url"] = "testString"

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToSharepointObjectParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToObjectSummary(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.ObjectSummary) {
		sharepointObjectParamsModel := new(backuprecoveryv1.SharepointObjectParams)
		sharepointObjectParamsModel.SiteWebURL = core.StringPtr("testString")

		objectTypeVCenterParamsModel := new(backuprecoveryv1.ObjectTypeVCenterParams)
		objectTypeVCenterParamsModel.IsCloudEnv = core.BoolPtr(true)

		objectTypeWindowsClusterParamsModel := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
		objectTypeWindowsClusterParamsModel.ClusterSourceType = core.StringPtr("testString")

		model := new(backuprecoveryv1.ObjectSummary)
		model.ID = core.Int64Ptr(int64(26))
		model.Name = core.StringPtr("testString")
		model.SourceID = core.Int64Ptr(int64(26))
		model.SourceName = core.StringPtr("testString")
		model.Environment = core.StringPtr("kPhysical")
		model.ObjectHash = core.StringPtr("testString")
		model.ObjectType = core.StringPtr("kCluster")
		model.LogicalSizeBytes = core.Int64Ptr(int64(26))
		model.UUID = core.StringPtr("testString")
		model.GlobalID = core.StringPtr("testString")
		model.ProtectionType = core.StringPtr("kAgent")
		model.SharepointSiteSummary = sharepointObjectParamsModel
		model.OsType = core.StringPtr("kLinux")
		model.VCenterSummary = objectTypeVCenterParamsModel
		model.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

		assert.Equal(t, result, model)
	}

	sharepointObjectParamsModel := make(map[string]interface{})
	sharepointObjectParamsModel["site_web_url"] = "testString"

	objectTypeVCenterParamsModel := make(map[string]interface{})
	objectTypeVCenterParamsModel["is_cloud_env"] = true

	objectTypeWindowsClusterParamsModel := make(map[string]interface{})
	objectTypeWindowsClusterParamsModel["cluster_source_type"] = "testString"

	model := make(map[string]interface{})
	model["id"] = int(26)
	model["name"] = "testString"
	model["source_id"] = int(26)
	model["source_name"] = "testString"
	model["environment"] = "kPhysical"
	model["object_hash"] = "testString"
	model["object_type"] = "kCluster"
	model["logical_size_bytes"] = int(26)
	model["uuid"] = "testString"
	model["global_id"] = "testString"
	model["protection_type"] = "kAgent"
	model["sharepoint_site_summary"] = []interface{}{sharepointObjectParamsModel}
	model["os_type"] = "kLinux"
	model["v_center_summary"] = []interface{}{objectTypeVCenterParamsModel}
	model["windows_cluster_summary"] = []interface{}{objectTypeWindowsClusterParamsModel}

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToObjectSummary(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToObjectTypeVCenterParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.ObjectTypeVCenterParams) {
		model := new(backuprecoveryv1.ObjectTypeVCenterParams)
		model.IsCloudEnv = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["is_cloud_env"] = true

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToObjectTypeVCenterParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToObjectTypeWindowsClusterParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.ObjectTypeWindowsClusterParams) {
		model := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
		model.ClusterSourceType = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["cluster_source_type"] = "testString"

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToObjectTypeWindowsClusterParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParamsArchivalTargetInfo(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo) {
		awsTierModel := new(backuprecoveryv1.AWSTier)
		awsTierModel.MoveAfterUnit = core.StringPtr("Days")
		awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
		awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

		awsTiersModel := new(backuprecoveryv1.AWSTiers)
		awsTiersModel.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

		azureTierModel := new(backuprecoveryv1.AzureTier)
		azureTierModel.MoveAfterUnit = core.StringPtr("Days")
		azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
		azureTierModel.TierType = core.StringPtr("kAzureTierHot")

		azureTiersModel := new(backuprecoveryv1.AzureTiers)
		azureTiersModel.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

		googleTierModel := new(backuprecoveryv1.GoogleTier)
		googleTierModel.MoveAfterUnit = core.StringPtr("Days")
		googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
		googleTierModel.TierType = core.StringPtr("kGoogleStandard")

		googleTiersModel := new(backuprecoveryv1.GoogleTiers)
		googleTiersModel.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

		oracleTierModel := new(backuprecoveryv1.OracleTier)
		oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
		oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
		oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

		oracleTiersModel := new(backuprecoveryv1.OracleTiers)
		oracleTiersModel.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

		archivalTargetTierInfoModel := new(backuprecoveryv1.ArchivalTargetTierInfo)
		archivalTargetTierInfoModel.AwsTiering = awsTiersModel
		archivalTargetTierInfoModel.AzureTiering = azureTiersModel
		archivalTargetTierInfoModel.CloudPlatform = core.StringPtr("AWS")
		archivalTargetTierInfoModel.GoogleTiering = googleTiersModel
		archivalTargetTierInfoModel.OracleTiering = oracleTiersModel
		archivalTargetTierInfoModel.CurrentTierType = core.StringPtr("kAmazonS3Standard")

		model := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo)
		model.TargetID = core.Int64Ptr(int64(26))
		model.ArchivalTaskID = core.StringPtr("testString")
		model.TargetName = core.StringPtr("testString")
		model.TargetType = core.StringPtr("Tape")
		model.UsageType = core.StringPtr("Archival")
		model.OwnershipContext = core.StringPtr("Local")
		model.TierSettings = archivalTargetTierInfoModel

		assert.Equal(t, result, model)
	}

	awsTierModel := make(map[string]interface{})
	awsTierModel["move_after_unit"] = "Days"
	awsTierModel["move_after"] = int(26)
	awsTierModel["tier_type"] = "kAmazonS3Standard"

	awsTiersModel := make(map[string]interface{})
	awsTiersModel["tiers"] = []interface{}{awsTierModel}

	azureTierModel := make(map[string]interface{})
	azureTierModel["move_after_unit"] = "Days"
	azureTierModel["move_after"] = int(26)
	azureTierModel["tier_type"] = "kAzureTierHot"

	azureTiersModel := make(map[string]interface{})
	azureTiersModel["tiers"] = []interface{}{azureTierModel}

	googleTierModel := make(map[string]interface{})
	googleTierModel["move_after_unit"] = "Days"
	googleTierModel["move_after"] = int(26)
	googleTierModel["tier_type"] = "kGoogleStandard"

	googleTiersModel := make(map[string]interface{})
	googleTiersModel["tiers"] = []interface{}{googleTierModel}

	oracleTierModel := make(map[string]interface{})
	oracleTierModel["move_after_unit"] = "Days"
	oracleTierModel["move_after"] = int(26)
	oracleTierModel["tier_type"] = "kOracleTierStandard"

	oracleTiersModel := make(map[string]interface{})
	oracleTiersModel["tiers"] = []interface{}{oracleTierModel}

	archivalTargetTierInfoModel := make(map[string]interface{})
	archivalTargetTierInfoModel["aws_tiering"] = []interface{}{awsTiersModel}
	archivalTargetTierInfoModel["azure_tiering"] = []interface{}{azureTiersModel}
	archivalTargetTierInfoModel["cloud_platform"] = "AWS"
	archivalTargetTierInfoModel["google_tiering"] = []interface{}{googleTiersModel}
	archivalTargetTierInfoModel["oracle_tiering"] = []interface{}{oracleTiersModel}
	archivalTargetTierInfoModel["current_tier_type"] = "kAmazonS3Standard"

	model := make(map[string]interface{})
	model["target_id"] = int(26)
	model["archival_task_id"] = "testString"
	model["target_name"] = "testString"
	model["target_type"] = "Tape"
	model["usage_type"] = "Archival"
	model["ownership_context"] = "Local"
	model["tier_settings"] = []interface{}{archivalTargetTierInfoModel}

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToCommonRecoverObjectSnapshotParamsArchivalTargetInfo(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToArchivalTargetTierInfo(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.ArchivalTargetTierInfo) {
		awsTierModel := new(backuprecoveryv1.AWSTier)
		awsTierModel.MoveAfterUnit = core.StringPtr("Days")
		awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
		awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

		awsTiersModel := new(backuprecoveryv1.AWSTiers)
		awsTiersModel.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

		azureTierModel := new(backuprecoveryv1.AzureTier)
		azureTierModel.MoveAfterUnit = core.StringPtr("Days")
		azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
		azureTierModel.TierType = core.StringPtr("kAzureTierHot")

		azureTiersModel := new(backuprecoveryv1.AzureTiers)
		azureTiersModel.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

		googleTierModel := new(backuprecoveryv1.GoogleTier)
		googleTierModel.MoveAfterUnit = core.StringPtr("Days")
		googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
		googleTierModel.TierType = core.StringPtr("kGoogleStandard")

		googleTiersModel := new(backuprecoveryv1.GoogleTiers)
		googleTiersModel.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

		oracleTierModel := new(backuprecoveryv1.OracleTier)
		oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
		oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
		oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

		oracleTiersModel := new(backuprecoveryv1.OracleTiers)
		oracleTiersModel.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

		model := new(backuprecoveryv1.ArchivalTargetTierInfo)
		model.AwsTiering = awsTiersModel
		model.AzureTiering = azureTiersModel
		model.CloudPlatform = core.StringPtr("AWS")
		model.GoogleTiering = googleTiersModel
		model.OracleTiering = oracleTiersModel
		model.CurrentTierType = core.StringPtr("kAmazonS3Standard")

		assert.Equal(t, result, model)
	}

	awsTierModel := make(map[string]interface{})
	awsTierModel["move_after_unit"] = "Days"
	awsTierModel["move_after"] = int(26)
	awsTierModel["tier_type"] = "kAmazonS3Standard"

	awsTiersModel := make(map[string]interface{})
	awsTiersModel["tiers"] = []interface{}{awsTierModel}

	azureTierModel := make(map[string]interface{})
	azureTierModel["move_after_unit"] = "Days"
	azureTierModel["move_after"] = int(26)
	azureTierModel["tier_type"] = "kAzureTierHot"

	azureTiersModel := make(map[string]interface{})
	azureTiersModel["tiers"] = []interface{}{azureTierModel}

	googleTierModel := make(map[string]interface{})
	googleTierModel["move_after_unit"] = "Days"
	googleTierModel["move_after"] = int(26)
	googleTierModel["tier_type"] = "kGoogleStandard"

	googleTiersModel := make(map[string]interface{})
	googleTiersModel["tiers"] = []interface{}{googleTierModel}

	oracleTierModel := make(map[string]interface{})
	oracleTierModel["move_after_unit"] = "Days"
	oracleTierModel["move_after"] = int(26)
	oracleTierModel["tier_type"] = "kOracleTierStandard"

	oracleTiersModel := make(map[string]interface{})
	oracleTiersModel["tiers"] = []interface{}{oracleTierModel}

	model := make(map[string]interface{})
	model["aws_tiering"] = []interface{}{awsTiersModel}
	model["azure_tiering"] = []interface{}{azureTiersModel}
	model["cloud_platform"] = "AWS"
	model["google_tiering"] = []interface{}{googleTiersModel}
	model["oracle_tiering"] = []interface{}{oracleTiersModel}
	model["current_tier_type"] = "kAmazonS3Standard"

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToArchivalTargetTierInfo(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToAWSTiers(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.AWSTiers) {
		awsTierModel := new(backuprecoveryv1.AWSTier)
		awsTierModel.MoveAfterUnit = core.StringPtr("Days")
		awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
		awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

		model := new(backuprecoveryv1.AWSTiers)
		model.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

		assert.Equal(t, result, model)
	}

	awsTierModel := make(map[string]interface{})
	awsTierModel["move_after_unit"] = "Days"
	awsTierModel["move_after"] = int(26)
	awsTierModel["tier_type"] = "kAmazonS3Standard"

	model := make(map[string]interface{})
	model["tiers"] = []interface{}{awsTierModel}

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToAWSTiers(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToAWSTier(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.AWSTier) {
		model := new(backuprecoveryv1.AWSTier)
		model.MoveAfterUnit = core.StringPtr("Days")
		model.MoveAfter = core.Int64Ptr(int64(26))
		model.TierType = core.StringPtr("kAmazonS3Standard")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["move_after_unit"] = "Days"
	model["move_after"] = int(26)
	model["tier_type"] = "kAmazonS3Standard"

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToAWSTier(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToAzureTiers(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.AzureTiers) {
		azureTierModel := new(backuprecoveryv1.AzureTier)
		azureTierModel.MoveAfterUnit = core.StringPtr("Days")
		azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
		azureTierModel.TierType = core.StringPtr("kAzureTierHot")

		model := new(backuprecoveryv1.AzureTiers)
		model.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

		assert.Equal(t, result, model)
	}

	azureTierModel := make(map[string]interface{})
	azureTierModel["move_after_unit"] = "Days"
	azureTierModel["move_after"] = int(26)
	azureTierModel["tier_type"] = "kAzureTierHot"

	model := make(map[string]interface{})
	model["tiers"] = []interface{}{azureTierModel}

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToAzureTiers(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToAzureTier(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.AzureTier) {
		model := new(backuprecoveryv1.AzureTier)
		model.MoveAfterUnit = core.StringPtr("Days")
		model.MoveAfter = core.Int64Ptr(int64(26))
		model.TierType = core.StringPtr("kAzureTierHot")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["move_after_unit"] = "Days"
	model["move_after"] = int(26)
	model["tier_type"] = "kAzureTierHot"

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToAzureTier(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToGoogleTiers(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.GoogleTiers) {
		googleTierModel := new(backuprecoveryv1.GoogleTier)
		googleTierModel.MoveAfterUnit = core.StringPtr("Days")
		googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
		googleTierModel.TierType = core.StringPtr("kGoogleStandard")

		model := new(backuprecoveryv1.GoogleTiers)
		model.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

		assert.Equal(t, result, model)
	}

	googleTierModel := make(map[string]interface{})
	googleTierModel["move_after_unit"] = "Days"
	googleTierModel["move_after"] = int(26)
	googleTierModel["tier_type"] = "kGoogleStandard"

	model := make(map[string]interface{})
	model["tiers"] = []interface{}{googleTierModel}

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToGoogleTiers(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToGoogleTier(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.GoogleTier) {
		model := new(backuprecoveryv1.GoogleTier)
		model.MoveAfterUnit = core.StringPtr("Days")
		model.MoveAfter = core.Int64Ptr(int64(26))
		model.TierType = core.StringPtr("kGoogleStandard")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["move_after_unit"] = "Days"
	model["move_after"] = int(26)
	model["tier_type"] = "kGoogleStandard"

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToGoogleTier(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToOracleTiers(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.OracleTiers) {
		oracleTierModel := new(backuprecoveryv1.OracleTier)
		oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
		oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
		oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

		model := new(backuprecoveryv1.OracleTiers)
		model.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

		assert.Equal(t, result, model)
	}

	oracleTierModel := make(map[string]interface{})
	oracleTierModel["move_after_unit"] = "Days"
	oracleTierModel["move_after"] = int(26)
	oracleTierModel["tier_type"] = "kOracleTierStandard"

	model := make(map[string]interface{})
	model["tiers"] = []interface{}{oracleTierModel}

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToOracleTiers(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToOracleTier(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.OracleTier) {
		model := new(backuprecoveryv1.OracleTier)
		model.MoveAfterUnit = core.StringPtr("Days")
		model.MoveAfter = core.Int64Ptr(int64(26))
		model.TierType = core.StringPtr("kOracleTierStandard")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["move_after_unit"] = "Days"
	model["move_after"] = int(26)
	model["tier_type"] = "kOracleTierStandard"

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToOracleTier(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToFilesAndFoldersObject(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.FilesAndFoldersObject) {
		model := new(backuprecoveryv1.FilesAndFoldersObject)
		model.AbsolutePath = core.StringPtr("testString")
		model.IsDirectory = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["absolute_path"] = "testString"
	model["is_directory"] = true

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToFilesAndFoldersObject(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRecoveryDownloadFilesFoldersMapToDocumentObject(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.DocumentObject) {
		model := new(backuprecoveryv1.DocumentObject)
		model.IsDirectory = core.BoolPtr(true)
		model.ItemID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["is_directory"] = true
	model["item_id"] = "testString"

	result, err := backuprecovery.ResourceIbmRecoveryDownloadFilesFoldersMapToDocumentObject(model)
	assert.Nil(t, err)
	checkResult(result)
}
