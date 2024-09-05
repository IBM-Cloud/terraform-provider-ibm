// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasSearchObjectsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasSearchObjectsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_search_objects.baas_search_objects_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_search_objects.baas_search_objects_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasSearchObjectsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_baas_search_objects" "baas_search_objects_instance" {
			X-IBM-Tenant-Id = "X-IBM-Tenant-Id"
			requestInitiatorType = "UIUser"
			searchString = "searchString"
			environments = [ "kPhysical" ]
			protectionTypes = [ "kAgent" ]
			protectionGroupIds = [ "protectionGroupIds" ]
			objectIds = [ 1 ]
			osTypes = [ "kLinux" ]
			sourceIds = [ 1 ]
			sourceUuids = [ "sourceUuids" ]
			isProtected = true
			isDeleted = true
			lastRunStatusList = [ "Accepted" ]
			clusterIdentifiers = [ "clusterIdentifiers" ]
			includeDeletedObjects = true
			paginationCookie = "paginationCookie"
			object_count = 1
			mustHaveTagIds = [ "mustHaveTagIds" ]
			mightHaveTagIds = [ "mightHaveTagIds" ]
			mustHaveSnapshotTagIds = [ "mustHaveSnapshotTagIds" ]
			mightHaveSnapshotTagIds = [ "mightHaveSnapshotTagIds" ]
			tagSearchName = "tagSearchName"
			tagNames = [ "tagNames" ]
			tagTypes = [ "System" ]
			tagCategories = [ "Security" ]
			tagSubCategories = [ "Classification" ]
			includeHeliosTagInfoForObjects = true
			externalFilters = [ "externalFilters" ]
		}
	`)
}

func TestDataSourceIbmBaasSearchObjectsSearchObjectToMap(t *testing.T) {
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

		objectProtectionStatsSummaryModel := make(map[string]interface{})
		objectProtectionStatsSummaryModel["environment"] = "kPhysical"
		objectProtectionStatsSummaryModel["protected_count"] = int(26)
		objectProtectionStatsSummaryModel["unprotected_count"] = int(26)
		objectProtectionStatsSummaryModel["deleted_protected_count"] = int(26)
		objectProtectionStatsSummaryModel["protected_size_bytes"] = int(26)
		objectProtectionStatsSummaryModel["unprotected_size_bytes"] = int(26)

		userModel := make(map[string]interface{})
		userModel["name"] = "testString"
		userModel["sid"] = "testString"
		userModel["domain"] = "testString"

		groupModel := make(map[string]interface{})
		groupModel["name"] = "testString"
		groupModel["sid"] = "testString"
		groupModel["domain"] = "testString"

		externalVendorCustomPropertiesModel := make(map[string]interface{})
		externalVendorCustomPropertiesModel["key"] = "testString"
		externalVendorCustomPropertiesModel["value"] = "testString"

		ibmTenantMetadataParamsModel := make(map[string]interface{})
		ibmTenantMetadataParamsModel["account_id"] = "testString"
		ibmTenantMetadataParamsModel["crn"] = "testString"
		ibmTenantMetadataParamsModel["custom_properties"] = []map[string]interface{}{externalVendorCustomPropertiesModel}
		ibmTenantMetadataParamsModel["liveness_mode"] = "Active"
		ibmTenantMetadataParamsModel["ownership_mode"] = "Primary"
		ibmTenantMetadataParamsModel["resource_group_id"] = "testString"

		externalVendorTenantMetadataModel := make(map[string]interface{})
		externalVendorTenantMetadataModel["ibm_tenant_metadata_params"] = []map[string]interface{}{ibmTenantMetadataParamsModel}
		externalVendorTenantMetadataModel["type"] = "IBM"

		tenantNetworkModel := make(map[string]interface{})
		tenantNetworkModel["connector_enabled"] = true
		tenantNetworkModel["cluster_hostname"] = "testString"
		tenantNetworkModel["cluster_ips"] = []string{"testString"}

		tenantModel := make(map[string]interface{})
		tenantModel["description"] = "testString"
		tenantModel["external_vendor_metadata"] = []map[string]interface{}{externalVendorTenantMetadataModel}
		tenantModel["id"] = "testString"
		tenantModel["is_managed_on_helios"] = true
		tenantModel["name"] = "testString"
		tenantModel["network"] = []map[string]interface{}{tenantNetworkModel}
		tenantModel["status"] = "Active"

		permissionInfoModel := make(map[string]interface{})
		permissionInfoModel["object_id"] = int(26)
		permissionInfoModel["users"] = []map[string]interface{}{userModel}
		permissionInfoModel["groups"] = []map[string]interface{}{groupModel}
		permissionInfoModel["tenant"] = []map[string]interface{}{tenantModel}

		aagInfoModel := make(map[string]interface{})
		aagInfoModel["name"] = "testString"
		aagInfoModel["object_id"] = int(26)

		hostInformationModel := make(map[string]interface{})
		hostInformationModel["id"] = "testString"
		hostInformationModel["name"] = "testString"
		hostInformationModel["environment"] = "kPhysical"

		searchObjectMssqlParamsModel := make(map[string]interface{})
		searchObjectMssqlParamsModel["aag_info"] = []map[string]interface{}{aagInfoModel}
		searchObjectMssqlParamsModel["host_info"] = []map[string]interface{}{hostInformationModel}
		searchObjectMssqlParamsModel["is_encrypted"] = true

		searchObjectPhysicalParamsModel := make(map[string]interface{})
		searchObjectPhysicalParamsModel["enable_system_backup"] = true

		tagInfoModel := make(map[string]interface{})
		tagInfoModel["tag_id"] = "testString"

		snapshotTagInfoModel := make(map[string]interface{})
		snapshotTagInfoModel["tag_id"] = "testString"
		snapshotTagInfoModel["run_ids"] = []string{"testString"}

		heliosTagInfoModel := make(map[string]interface{})
		heliosTagInfoModel["category"] = "Security"
		heliosTagInfoModel["name"] = "testString"
		heliosTagInfoModel["sub_category"] = "Classification"
		heliosTagInfoModel["third_party_name"] = "testString"
		heliosTagInfoModel["type"] = "System"
		heliosTagInfoModel["ui_color"] = "testString"
		heliosTagInfoModel["updated_time_usecs"] = int(26)
		heliosTagInfoModel["uuid"] = "testString"

		searchObjectSourceInfoMssqlParamsModel := make(map[string]interface{})
		searchObjectSourceInfoMssqlParamsModel["aag_info"] = []map[string]interface{}{aagInfoModel}
		searchObjectSourceInfoMssqlParamsModel["host_info"] = []map[string]interface{}{hostInformationModel}
		searchObjectSourceInfoMssqlParamsModel["is_encrypted"] = true

		searchObjectSourceInfoPhysicalParamsModel := make(map[string]interface{})
		searchObjectSourceInfoPhysicalParamsModel["enable_system_backup"] = true

		searchObjectSourceInfoModel := make(map[string]interface{})
		searchObjectSourceInfoModel["id"] = int(26)
		searchObjectSourceInfoModel["name"] = "testString"
		searchObjectSourceInfoModel["source_id"] = int(26)
		searchObjectSourceInfoModel["source_name"] = "testString"
		searchObjectSourceInfoModel["environment"] = "kPhysical"
		searchObjectSourceInfoModel["object_hash"] = "testString"
		searchObjectSourceInfoModel["object_type"] = "kCluster"
		searchObjectSourceInfoModel["logical_size_bytes"] = int(26)
		searchObjectSourceInfoModel["uuid"] = "testString"
		searchObjectSourceInfoModel["global_id"] = "testString"
		searchObjectSourceInfoModel["protection_type"] = "kAgent"
		searchObjectSourceInfoModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		searchObjectSourceInfoModel["os_type"] = "kLinux"
		searchObjectSourceInfoModel["child_objects"] = []map[string]interface{}{objectSummaryModel}
		searchObjectSourceInfoModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		searchObjectSourceInfoModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}
		searchObjectSourceInfoModel["protection_stats"] = []map[string]interface{}{objectProtectionStatsSummaryModel}
		searchObjectSourceInfoModel["permissions"] = []map[string]interface{}{permissionInfoModel}
		searchObjectSourceInfoModel["mssql_params"] = []map[string]interface{}{searchObjectSourceInfoMssqlParamsModel}
		searchObjectSourceInfoModel["physical_params"] = []map[string]interface{}{searchObjectSourceInfoPhysicalParamsModel}

		objectProtectionGroupSummaryModel := make(map[string]interface{})
		objectProtectionGroupSummaryModel["name"] = "testString"
		objectProtectionGroupSummaryModel["id"] = "testString"
		objectProtectionGroupSummaryModel["protection_env_type"] = "kAgent"
		objectProtectionGroupSummaryModel["policy_name"] = "testString"
		objectProtectionGroupSummaryModel["policy_id"] = "testString"
		objectProtectionGroupSummaryModel["last_backup_run_status"] = "Accepted"
		objectProtectionGroupSummaryModel["last_archival_run_status"] = "Accepted"
		objectProtectionGroupSummaryModel["last_replication_run_status"] = "Accepted"
		objectProtectionGroupSummaryModel["last_run_sla_violated"] = true

		protectionSummaryModel := make(map[string]interface{})
		protectionSummaryModel["policy_name"] = "testString"
		protectionSummaryModel["policy_id"] = "testString"
		protectionSummaryModel["last_backup_run_status"] = "Accepted"
		protectionSummaryModel["last_archival_run_status"] = "Accepted"
		protectionSummaryModel["last_replication_run_status"] = "Accepted"
		protectionSummaryModel["last_run_sla_violated"] = true

		objectProtectionInfoModel := make(map[string]interface{})
		objectProtectionInfoModel["object_id"] = int(26)
		objectProtectionInfoModel["source_id"] = int(26)
		objectProtectionInfoModel["view_id"] = int(26)
		objectProtectionInfoModel["region_id"] = "testString"
		objectProtectionInfoModel["cluster_id"] = int(26)
		objectProtectionInfoModel["cluster_incarnation_id"] = int(26)
		objectProtectionInfoModel["tenant_ids"] = []string{"testString"}
		objectProtectionInfoModel["is_deleted"] = true
		objectProtectionInfoModel["protection_groups"] = []map[string]interface{}{objectProtectionGroupSummaryModel}
		objectProtectionInfoModel["object_backup_configuration"] = []map[string]interface{}{protectionSummaryModel}
		objectProtectionInfoModel["last_run_status"] = "Accepted"

		secondaryIdModel := make(map[string]interface{})
		secondaryIdModel["name"] = "testString"
		secondaryIdModel["value"] = "testString"

		taggedSnapshotInfoModel := make(map[string]interface{})
		taggedSnapshotInfoModel["cluster_id"] = int(26)
		taggedSnapshotInfoModel["cluster_incarnation_id"] = int(26)
		taggedSnapshotInfoModel["job_id"] = int(26)
		taggedSnapshotInfoModel["object_uuid"] = "testString"
		taggedSnapshotInfoModel["run_start_time_usecs"] = int(26)
		taggedSnapshotInfoModel["tags"] = []map[string]interface{}{heliosTagInfoModel}

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
		model["protection_stats"] = []map[string]interface{}{objectProtectionStatsSummaryModel}
		model["permissions"] = []map[string]interface{}{permissionInfoModel}
		model["mssql_params"] = []map[string]interface{}{searchObjectMssqlParamsModel}
		model["physical_params"] = []map[string]interface{}{searchObjectPhysicalParamsModel}
		model["tags"] = []map[string]interface{}{tagInfoModel}
		model["snapshot_tags"] = []map[string]interface{}{snapshotTagInfoModel}
		model["helios_tags"] = []map[string]interface{}{heliosTagInfoModel}
		model["source_info"] = []map[string]interface{}{searchObjectSourceInfoModel}
		model["object_protection_infos"] = []map[string]interface{}{objectProtectionInfoModel}
		model["secondary_ids"] = []map[string]interface{}{secondaryIdModel}
		model["tagged_snapshots"] = []map[string]interface{}{taggedSnapshotInfoModel}

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

	objectProtectionStatsSummaryModel := new(backuprecoveryv1.ObjectProtectionStatsSummary)
	objectProtectionStatsSummaryModel.Environment = core.StringPtr("kPhysical")
	objectProtectionStatsSummaryModel.ProtectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.UnprotectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.DeletedProtectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.ProtectedSizeBytes = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.UnprotectedSizeBytes = core.Int64Ptr(int64(26))

	userModel := new(backuprecoveryv1.User)
	userModel.Name = core.StringPtr("testString")
	userModel.Sid = core.StringPtr("testString")
	userModel.Domain = core.StringPtr("testString")

	groupModel := new(backuprecoveryv1.Group)
	groupModel.Name = core.StringPtr("testString")
	groupModel.Sid = core.StringPtr("testString")
	groupModel.Domain = core.StringPtr("testString")

	externalVendorCustomPropertiesModel := new(backuprecoveryv1.ExternalVendorCustomProperties)
	externalVendorCustomPropertiesModel.Key = core.StringPtr("testString")
	externalVendorCustomPropertiesModel.Value = core.StringPtr("testString")

	ibmTenantMetadataParamsModel := new(backuprecoveryv1.IbmTenantMetadataParams)
	ibmTenantMetadataParamsModel.AccountID = core.StringPtr("testString")
	ibmTenantMetadataParamsModel.Crn = core.StringPtr("testString")
	ibmTenantMetadataParamsModel.CustomProperties = []backuprecoveryv1.ExternalVendorCustomProperties{*externalVendorCustomPropertiesModel}
	ibmTenantMetadataParamsModel.LivenessMode = core.StringPtr("Active")
	ibmTenantMetadataParamsModel.OwnershipMode = core.StringPtr("Primary")
	ibmTenantMetadataParamsModel.ResourceGroupID = core.StringPtr("testString")

	externalVendorTenantMetadataModel := new(backuprecoveryv1.ExternalVendorTenantMetadata)
	externalVendorTenantMetadataModel.IbmTenantMetadataParams = ibmTenantMetadataParamsModel
	externalVendorTenantMetadataModel.Type = core.StringPtr("IBM")

	tenantNetworkModel := new(backuprecoveryv1.TenantNetwork)
	tenantNetworkModel.ConnectorEnabled = core.BoolPtr(true)
	tenantNetworkModel.ClusterHostname = core.StringPtr("testString")
	tenantNetworkModel.ClusterIps = []string{"testString"}

	tenantModel := new(backuprecoveryv1.Tenant)
	tenantModel.Description = core.StringPtr("testString")
	tenantModel.ExternalVendorMetadata = externalVendorTenantMetadataModel
	tenantModel.ID = core.StringPtr("testString")
	tenantModel.IsManagedOnHelios = core.BoolPtr(true)
	tenantModel.Name = core.StringPtr("testString")
	tenantModel.Network = tenantNetworkModel
	tenantModel.Status = core.StringPtr("Active")

	permissionInfoModel := new(backuprecoveryv1.PermissionInfo)
	permissionInfoModel.ObjectID = core.Int64Ptr(int64(26))
	permissionInfoModel.Users = []backuprecoveryv1.User{*userModel}
	permissionInfoModel.Groups = []backuprecoveryv1.Group{*groupModel}
	permissionInfoModel.Tenant = tenantModel

	aagInfoModel := new(backuprecoveryv1.AAGInfo)
	aagInfoModel.Name = core.StringPtr("testString")
	aagInfoModel.ObjectID = core.Int64Ptr(int64(26))

	hostInformationModel := new(backuprecoveryv1.HostInformation)
	hostInformationModel.ID = core.StringPtr("testString")
	hostInformationModel.Name = core.StringPtr("testString")
	hostInformationModel.Environment = core.StringPtr("kPhysical")

	searchObjectMssqlParamsModel := new(backuprecoveryv1.SearchObjectMssqlParams)
	searchObjectMssqlParamsModel.AagInfo = aagInfoModel
	searchObjectMssqlParamsModel.HostInfo = hostInformationModel
	searchObjectMssqlParamsModel.IsEncrypted = core.BoolPtr(true)

	searchObjectPhysicalParamsModel := new(backuprecoveryv1.SearchObjectPhysicalParams)
	searchObjectPhysicalParamsModel.EnableSystemBackup = core.BoolPtr(true)

	tagInfoModel := new(backuprecoveryv1.TagInfo)
	tagInfoModel.TagID = core.StringPtr("testString")

	snapshotTagInfoModel := new(backuprecoveryv1.SnapshotTagInfo)
	snapshotTagInfoModel.TagID = core.StringPtr("testString")
	snapshotTagInfoModel.RunIds = []string{"testString"}

	heliosTagInfoModel := new(backuprecoveryv1.HeliosTagInfo)
	heliosTagInfoModel.Category = core.StringPtr("Security")
	heliosTagInfoModel.Name = core.StringPtr("testString")
	heliosTagInfoModel.SubCategory = core.StringPtr("Classification")
	heliosTagInfoModel.ThirdPartyName = core.StringPtr("testString")
	heliosTagInfoModel.Type = core.StringPtr("System")
	heliosTagInfoModel.UiColor = core.StringPtr("testString")
	heliosTagInfoModel.UpdatedTimeUsecs = core.Int64Ptr(int64(26))
	heliosTagInfoModel.UUID = core.StringPtr("testString")

	searchObjectSourceInfoMssqlParamsModel := new(backuprecoveryv1.SearchObjectSourceInfoMssqlParams)
	searchObjectSourceInfoMssqlParamsModel.AagInfo = aagInfoModel
	searchObjectSourceInfoMssqlParamsModel.HostInfo = hostInformationModel
	searchObjectSourceInfoMssqlParamsModel.IsEncrypted = core.BoolPtr(true)

	searchObjectSourceInfoPhysicalParamsModel := new(backuprecoveryv1.SearchObjectSourceInfoPhysicalParams)
	searchObjectSourceInfoPhysicalParamsModel.EnableSystemBackup = core.BoolPtr(true)

	searchObjectSourceInfoModel := new(backuprecoveryv1.SearchObjectSourceInfo)
	searchObjectSourceInfoModel.ID = core.Int64Ptr(int64(26))
	searchObjectSourceInfoModel.Name = core.StringPtr("testString")
	searchObjectSourceInfoModel.SourceID = core.Int64Ptr(int64(26))
	searchObjectSourceInfoModel.SourceName = core.StringPtr("testString")
	searchObjectSourceInfoModel.Environment = core.StringPtr("kPhysical")
	searchObjectSourceInfoModel.ObjectHash = core.StringPtr("testString")
	searchObjectSourceInfoModel.ObjectType = core.StringPtr("kCluster")
	searchObjectSourceInfoModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	searchObjectSourceInfoModel.UUID = core.StringPtr("testString")
	searchObjectSourceInfoModel.GlobalID = core.StringPtr("testString")
	searchObjectSourceInfoModel.ProtectionType = core.StringPtr("kAgent")
	searchObjectSourceInfoModel.SharepointSiteSummary = sharepointObjectParamsModel
	searchObjectSourceInfoModel.OsType = core.StringPtr("kLinux")
	searchObjectSourceInfoModel.ChildObjects = []backuprecoveryv1.ObjectSummary{*objectSummaryModel}
	searchObjectSourceInfoModel.VCenterSummary = objectTypeVCenterParamsModel
	searchObjectSourceInfoModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel
	searchObjectSourceInfoModel.ProtectionStats = []backuprecoveryv1.ObjectProtectionStatsSummary{*objectProtectionStatsSummaryModel}
	searchObjectSourceInfoModel.Permissions = permissionInfoModel
	searchObjectSourceInfoModel.MssqlParams = searchObjectSourceInfoMssqlParamsModel
	searchObjectSourceInfoModel.PhysicalParams = searchObjectSourceInfoPhysicalParamsModel

	objectProtectionGroupSummaryModel := new(backuprecoveryv1.ObjectProtectionGroupSummary)
	objectProtectionGroupSummaryModel.Name = core.StringPtr("testString")
	objectProtectionGroupSummaryModel.ID = core.StringPtr("testString")
	objectProtectionGroupSummaryModel.ProtectionEnvType = core.StringPtr("kAgent")
	objectProtectionGroupSummaryModel.PolicyName = core.StringPtr("testString")
	objectProtectionGroupSummaryModel.PolicyID = core.StringPtr("testString")
	objectProtectionGroupSummaryModel.LastBackupRunStatus = core.StringPtr("Accepted")
	objectProtectionGroupSummaryModel.LastArchivalRunStatus = core.StringPtr("Accepted")
	objectProtectionGroupSummaryModel.LastReplicationRunStatus = core.StringPtr("Accepted")
	objectProtectionGroupSummaryModel.LastRunSlaViolated = core.BoolPtr(true)

	protectionSummaryModel := new(backuprecoveryv1.ProtectionSummary)
	protectionSummaryModel.PolicyName = core.StringPtr("testString")
	protectionSummaryModel.PolicyID = core.StringPtr("testString")
	protectionSummaryModel.LastBackupRunStatus = core.StringPtr("Accepted")
	protectionSummaryModel.LastArchivalRunStatus = core.StringPtr("Accepted")
	protectionSummaryModel.LastReplicationRunStatus = core.StringPtr("Accepted")
	protectionSummaryModel.LastRunSlaViolated = core.BoolPtr(true)

	objectProtectionInfoModel := new(backuprecoveryv1.ObjectProtectionInfo)
	objectProtectionInfoModel.ObjectID = core.Int64Ptr(int64(26))
	objectProtectionInfoModel.SourceID = core.Int64Ptr(int64(26))
	objectProtectionInfoModel.ViewID = core.Int64Ptr(int64(26))
	objectProtectionInfoModel.RegionID = core.StringPtr("testString")
	objectProtectionInfoModel.ClusterID = core.Int64Ptr(int64(26))
	objectProtectionInfoModel.ClusterIncarnationID = core.Int64Ptr(int64(26))
	objectProtectionInfoModel.TenantIds = []string{"testString"}
	objectProtectionInfoModel.IsDeleted = core.BoolPtr(true)
	objectProtectionInfoModel.ProtectionGroups = []backuprecoveryv1.ObjectProtectionGroupSummary{*objectProtectionGroupSummaryModel}
	objectProtectionInfoModel.ObjectBackupConfiguration = []backuprecoveryv1.ProtectionSummary{*protectionSummaryModel}
	objectProtectionInfoModel.LastRunStatus = core.StringPtr("Accepted")

	secondaryIdModel := new(backuprecoveryv1.SecondaryID)
	secondaryIdModel.Name = core.StringPtr("testString")
	secondaryIdModel.Value = core.StringPtr("testString")

	taggedSnapshotInfoModel := new(backuprecoveryv1.TaggedSnapshotInfo)
	taggedSnapshotInfoModel.ClusterID = core.Int64Ptr(int64(26))
	taggedSnapshotInfoModel.ClusterIncarnationID = core.Int64Ptr(int64(26))
	taggedSnapshotInfoModel.JobID = core.Int64Ptr(int64(26))
	taggedSnapshotInfoModel.ObjectUUID = core.StringPtr("testString")
	taggedSnapshotInfoModel.RunStartTimeUsecs = core.Int64Ptr(int64(26))
	taggedSnapshotInfoModel.Tags = []backuprecoveryv1.HeliosTagInfo{*heliosTagInfoModel}

	model := new(backuprecoveryv1.SearchObject)
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
	model.ProtectionStats = []backuprecoveryv1.ObjectProtectionStatsSummary{*objectProtectionStatsSummaryModel}
	model.Permissions = permissionInfoModel
	model.MssqlParams = searchObjectMssqlParamsModel
	model.PhysicalParams = searchObjectPhysicalParamsModel
	model.Tags = []backuprecoveryv1.TagInfo{*tagInfoModel}
	model.SnapshotTags = []backuprecoveryv1.SnapshotTagInfo{*snapshotTagInfoModel}
	model.HeliosTags = []backuprecoveryv1.HeliosTagInfo{*heliosTagInfoModel}
	model.SourceInfo = searchObjectSourceInfoModel
	model.ObjectProtectionInfos = []backuprecoveryv1.ObjectProtectionInfo{*objectProtectionInfoModel}
	model.SecondaryIds = []backuprecoveryv1.SecondaryID{*secondaryIdModel}
	model.TaggedSnapshots = []backuprecoveryv1.TaggedSnapshotInfo{*taggedSnapshotInfoModel}

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsSearchObjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsSharepointObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["site_web_url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SharepointObjectParams)
	model.SiteWebURL = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsSharepointObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsObjectSummaryToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsObjectSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsObjectTypeVCenterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["is_cloud_env"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeVCenterParams)
	model.IsCloudEnv = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsObjectTypeVCenterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsObjectTypeWindowsClusterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cluster_source_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	model.ClusterSourceType = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsObjectTypeWindowsClusterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsObjectProtectionStatsSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["environment"] = "kPhysical"
		model["protected_count"] = int(26)
		model["unprotected_count"] = int(26)
		model["deleted_protected_count"] = int(26)
		model["protected_size_bytes"] = int(26)
		model["unprotected_size_bytes"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectProtectionStatsSummary)
	model.Environment = core.StringPtr("kPhysical")
	model.ProtectedCount = core.Int64Ptr(int64(26))
	model.UnprotectedCount = core.Int64Ptr(int64(26))
	model.DeletedProtectedCount = core.Int64Ptr(int64(26))
	model.ProtectedSizeBytes = core.Int64Ptr(int64(26))
	model.UnprotectedSizeBytes = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsObjectProtectionStatsSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsPermissionInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		userModel := make(map[string]interface{})
		userModel["name"] = "testString"
		userModel["sid"] = "testString"
		userModel["domain"] = "testString"

		groupModel := make(map[string]interface{})
		groupModel["name"] = "testString"
		groupModel["sid"] = "testString"
		groupModel["domain"] = "testString"

		externalVendorCustomPropertiesModel := make(map[string]interface{})
		externalVendorCustomPropertiesModel["key"] = "testString"
		externalVendorCustomPropertiesModel["value"] = "testString"

		ibmTenantMetadataParamsModel := make(map[string]interface{})
		ibmTenantMetadataParamsModel["account_id"] = "testString"
		ibmTenantMetadataParamsModel["crn"] = "testString"
		ibmTenantMetadataParamsModel["custom_properties"] = []map[string]interface{}{externalVendorCustomPropertiesModel}
		ibmTenantMetadataParamsModel["liveness_mode"] = "Active"
		ibmTenantMetadataParamsModel["ownership_mode"] = "Primary"
		ibmTenantMetadataParamsModel["resource_group_id"] = "testString"

		externalVendorTenantMetadataModel := make(map[string]interface{})
		externalVendorTenantMetadataModel["ibm_tenant_metadata_params"] = []map[string]interface{}{ibmTenantMetadataParamsModel}
		externalVendorTenantMetadataModel["type"] = "IBM"

		tenantNetworkModel := make(map[string]interface{})
		tenantNetworkModel["connector_enabled"] = true
		tenantNetworkModel["cluster_hostname"] = "testString"
		tenantNetworkModel["cluster_ips"] = []string{"testString"}

		tenantModel := make(map[string]interface{})
		tenantModel["description"] = "testString"
		tenantModel["external_vendor_metadata"] = []map[string]interface{}{externalVendorTenantMetadataModel}
		tenantModel["id"] = "testString"
		tenantModel["is_managed_on_helios"] = true
		tenantModel["name"] = "testString"
		tenantModel["network"] = []map[string]interface{}{tenantNetworkModel}
		tenantModel["status"] = "Active"

		model := make(map[string]interface{})
		model["object_id"] = int(26)
		model["users"] = []map[string]interface{}{userModel}
		model["groups"] = []map[string]interface{}{groupModel}
		model["tenant"] = []map[string]interface{}{tenantModel}

		assert.Equal(t, result, model)
	}

	userModel := new(backuprecoveryv1.User)
	userModel.Name = core.StringPtr("testString")
	userModel.Sid = core.StringPtr("testString")
	userModel.Domain = core.StringPtr("testString")

	groupModel := new(backuprecoveryv1.Group)
	groupModel.Name = core.StringPtr("testString")
	groupModel.Sid = core.StringPtr("testString")
	groupModel.Domain = core.StringPtr("testString")

	externalVendorCustomPropertiesModel := new(backuprecoveryv1.ExternalVendorCustomProperties)
	externalVendorCustomPropertiesModel.Key = core.StringPtr("testString")
	externalVendorCustomPropertiesModel.Value = core.StringPtr("testString")

	ibmTenantMetadataParamsModel := new(backuprecoveryv1.IbmTenantMetadataParams)
	ibmTenantMetadataParamsModel.AccountID = core.StringPtr("testString")
	ibmTenantMetadataParamsModel.Crn = core.StringPtr("testString")
	ibmTenantMetadataParamsModel.CustomProperties = []backuprecoveryv1.ExternalVendorCustomProperties{*externalVendorCustomPropertiesModel}
	ibmTenantMetadataParamsModel.LivenessMode = core.StringPtr("Active")
	ibmTenantMetadataParamsModel.OwnershipMode = core.StringPtr("Primary")
	ibmTenantMetadataParamsModel.ResourceGroupID = core.StringPtr("testString")

	externalVendorTenantMetadataModel := new(backuprecoveryv1.ExternalVendorTenantMetadata)
	externalVendorTenantMetadataModel.IbmTenantMetadataParams = ibmTenantMetadataParamsModel
	externalVendorTenantMetadataModel.Type = core.StringPtr("IBM")

	tenantNetworkModel := new(backuprecoveryv1.TenantNetwork)
	tenantNetworkModel.ConnectorEnabled = core.BoolPtr(true)
	tenantNetworkModel.ClusterHostname = core.StringPtr("testString")
	tenantNetworkModel.ClusterIps = []string{"testString"}

	tenantModel := new(backuprecoveryv1.Tenant)
	tenantModel.Description = core.StringPtr("testString")
	tenantModel.ExternalVendorMetadata = externalVendorTenantMetadataModel
	tenantModel.ID = core.StringPtr("testString")
	tenantModel.IsManagedOnHelios = core.BoolPtr(true)
	tenantModel.Name = core.StringPtr("testString")
	tenantModel.Network = tenantNetworkModel
	tenantModel.Status = core.StringPtr("Active")

	model := new(backuprecoveryv1.PermissionInfo)
	model.ObjectID = core.Int64Ptr(int64(26))
	model.Users = []backuprecoveryv1.User{*userModel}
	model.Groups = []backuprecoveryv1.Group{*groupModel}
	model.Tenant = tenantModel

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsPermissionInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsUserToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["sid"] = "testString"
		model["domain"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.User)
	model.Name = core.StringPtr("testString")
	model.Sid = core.StringPtr("testString")
	model.Domain = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsUserToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsGroupToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["sid"] = "testString"
		model["domain"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.Group)
	model.Name = core.StringPtr("testString")
	model.Sid = core.StringPtr("testString")
	model.Domain = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsTenantToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		externalVendorCustomPropertiesModel := make(map[string]interface{})
		externalVendorCustomPropertiesModel["key"] = "testString"
		externalVendorCustomPropertiesModel["value"] = "testString"

		ibmTenantMetadataParamsModel := make(map[string]interface{})
		ibmTenantMetadataParamsModel["account_id"] = "testString"
		ibmTenantMetadataParamsModel["crn"] = "testString"
		ibmTenantMetadataParamsModel["custom_properties"] = []map[string]interface{}{externalVendorCustomPropertiesModel}
		ibmTenantMetadataParamsModel["liveness_mode"] = "Active"
		ibmTenantMetadataParamsModel["ownership_mode"] = "Primary"
		ibmTenantMetadataParamsModel["resource_group_id"] = "testString"

		externalVendorTenantMetadataModel := make(map[string]interface{})
		externalVendorTenantMetadataModel["ibm_tenant_metadata_params"] = []map[string]interface{}{ibmTenantMetadataParamsModel}
		externalVendorTenantMetadataModel["type"] = "IBM"

		tenantNetworkModel := make(map[string]interface{})
		tenantNetworkModel["connector_enabled"] = true
		tenantNetworkModel["cluster_hostname"] = "testString"
		tenantNetworkModel["cluster_ips"] = []string{"testString"}

		model := make(map[string]interface{})
		model["created_at_time_msecs"] = int(26)
		model["deleted_at_time_msecs"] = int(26)
		model["description"] = "testString"
		model["external_vendor_metadata"] = []map[string]interface{}{externalVendorTenantMetadataModel}
		model["id"] = "testString"
		model["is_managed_on_helios"] = true
		model["last_updated_at_time_msecs"] = int(26)
		model["name"] = "testString"
		model["network"] = []map[string]interface{}{tenantNetworkModel}
		model["status"] = "Active"

		assert.Equal(t, result, model)
	}

	externalVendorCustomPropertiesModel := new(backuprecoveryv1.ExternalVendorCustomProperties)
	externalVendorCustomPropertiesModel.Key = core.StringPtr("testString")
	externalVendorCustomPropertiesModel.Value = core.StringPtr("testString")

	ibmTenantMetadataParamsModel := new(backuprecoveryv1.IbmTenantMetadataParams)
	ibmTenantMetadataParamsModel.AccountID = core.StringPtr("testString")
	ibmTenantMetadataParamsModel.Crn = core.StringPtr("testString")
	ibmTenantMetadataParamsModel.CustomProperties = []backuprecoveryv1.ExternalVendorCustomProperties{*externalVendorCustomPropertiesModel}
	ibmTenantMetadataParamsModel.LivenessMode = core.StringPtr("Active")
	ibmTenantMetadataParamsModel.OwnershipMode = core.StringPtr("Primary")
	ibmTenantMetadataParamsModel.ResourceGroupID = core.StringPtr("testString")

	externalVendorTenantMetadataModel := new(backuprecoveryv1.ExternalVendorTenantMetadata)
	externalVendorTenantMetadataModel.IbmTenantMetadataParams = ibmTenantMetadataParamsModel
	externalVendorTenantMetadataModel.Type = core.StringPtr("IBM")

	tenantNetworkModel := new(backuprecoveryv1.TenantNetwork)
	tenantNetworkModel.ConnectorEnabled = core.BoolPtr(true)
	tenantNetworkModel.ClusterHostname = core.StringPtr("testString")
	tenantNetworkModel.ClusterIps = []string{"testString"}

	model := new(backuprecoveryv1.Tenant)
	model.CreatedAtTimeMsecs = core.Int64Ptr(int64(26))
	model.DeletedAtTimeMsecs = core.Int64Ptr(int64(26))
	model.Description = core.StringPtr("testString")
	model.ExternalVendorMetadata = externalVendorTenantMetadataModel
	model.ID = core.StringPtr("testString")
	model.IsManagedOnHelios = core.BoolPtr(true)
	model.LastUpdatedAtTimeMsecs = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.Network = tenantNetworkModel
	model.Status = core.StringPtr("Active")

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsTenantToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsExternalVendorTenantMetadataToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		externalVendorCustomPropertiesModel := make(map[string]interface{})
		externalVendorCustomPropertiesModel["key"] = "testString"
		externalVendorCustomPropertiesModel["value"] = "testString"

		ibmTenantMetadataParamsModel := make(map[string]interface{})
		ibmTenantMetadataParamsModel["account_id"] = "testString"
		ibmTenantMetadataParamsModel["crn"] = "testString"
		ibmTenantMetadataParamsModel["custom_properties"] = []map[string]interface{}{externalVendorCustomPropertiesModel}
		ibmTenantMetadataParamsModel["liveness_mode"] = "Active"
		ibmTenantMetadataParamsModel["ownership_mode"] = "Primary"
		ibmTenantMetadataParamsModel["resource_group_id"] = "testString"

		model := make(map[string]interface{})
		model["ibm_tenant_metadata_params"] = []map[string]interface{}{ibmTenantMetadataParamsModel}
		model["type"] = "IBM"

		assert.Equal(t, result, model)
	}

	externalVendorCustomPropertiesModel := new(backuprecoveryv1.ExternalVendorCustomProperties)
	externalVendorCustomPropertiesModel.Key = core.StringPtr("testString")
	externalVendorCustomPropertiesModel.Value = core.StringPtr("testString")

	ibmTenantMetadataParamsModel := new(backuprecoveryv1.IbmTenantMetadataParams)
	ibmTenantMetadataParamsModel.AccountID = core.StringPtr("testString")
	ibmTenantMetadataParamsModel.Crn = core.StringPtr("testString")
	ibmTenantMetadataParamsModel.CustomProperties = []backuprecoveryv1.ExternalVendorCustomProperties{*externalVendorCustomPropertiesModel}
	ibmTenantMetadataParamsModel.LivenessMode = core.StringPtr("Active")
	ibmTenantMetadataParamsModel.OwnershipMode = core.StringPtr("Primary")
	ibmTenantMetadataParamsModel.ResourceGroupID = core.StringPtr("testString")

	model := new(backuprecoveryv1.ExternalVendorTenantMetadata)
	model.IbmTenantMetadataParams = ibmTenantMetadataParamsModel
	model.Type = core.StringPtr("IBM")

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsExternalVendorTenantMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsIbmTenantMetadataParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		externalVendorCustomPropertiesModel := make(map[string]interface{})
		externalVendorCustomPropertiesModel["key"] = "testString"
		externalVendorCustomPropertiesModel["value"] = "testString"

		model := make(map[string]interface{})
		model["account_id"] = "testString"
		model["crn"] = "testString"
		model["custom_properties"] = []map[string]interface{}{externalVendorCustomPropertiesModel}
		model["liveness_mode"] = "Active"
		model["ownership_mode"] = "Primary"
		model["resource_group_id"] = "testString"

		assert.Equal(t, result, model)
	}

	externalVendorCustomPropertiesModel := new(backuprecoveryv1.ExternalVendorCustomProperties)
	externalVendorCustomPropertiesModel.Key = core.StringPtr("testString")
	externalVendorCustomPropertiesModel.Value = core.StringPtr("testString")

	model := new(backuprecoveryv1.IbmTenantMetadataParams)
	model.AccountID = core.StringPtr("testString")
	model.Crn = core.StringPtr("testString")
	model.CustomProperties = []backuprecoveryv1.ExternalVendorCustomProperties{*externalVendorCustomPropertiesModel}
	model.LivenessMode = core.StringPtr("Active")
	model.OwnershipMode = core.StringPtr("Primary")
	model.ResourceGroupID = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsIbmTenantMetadataParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsExternalVendorCustomPropertiesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ExternalVendorCustomProperties)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsExternalVendorCustomPropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsTenantNetworkToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["connector_enabled"] = true
		model["cluster_hostname"] = "testString"
		model["cluster_ips"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.TenantNetwork)
	model.ConnectorEnabled = core.BoolPtr(true)
	model.ClusterHostname = core.StringPtr("testString")
	model.ClusterIps = []string{"testString"}

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsTenantNetworkToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsSearchObjectMssqlParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		aagInfoModel := make(map[string]interface{})
		aagInfoModel["name"] = "testString"
		aagInfoModel["object_id"] = int(26)

		hostInformationModel := make(map[string]interface{})
		hostInformationModel["id"] = "testString"
		hostInformationModel["name"] = "testString"
		hostInformationModel["environment"] = "kPhysical"

		model := make(map[string]interface{})
		model["aag_info"] = []map[string]interface{}{aagInfoModel}
		model["host_info"] = []map[string]interface{}{hostInformationModel}
		model["is_encrypted"] = true

		assert.Equal(t, result, model)
	}

	aagInfoModel := new(backuprecoveryv1.AAGInfo)
	aagInfoModel.Name = core.StringPtr("testString")
	aagInfoModel.ObjectID = core.Int64Ptr(int64(26))

	hostInformationModel := new(backuprecoveryv1.HostInformation)
	hostInformationModel.ID = core.StringPtr("testString")
	hostInformationModel.Name = core.StringPtr("testString")
	hostInformationModel.Environment = core.StringPtr("kPhysical")

	model := new(backuprecoveryv1.SearchObjectMssqlParams)
	model.AagInfo = aagInfoModel
	model.HostInfo = hostInformationModel
	model.IsEncrypted = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsSearchObjectMssqlParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsAAGInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["object_id"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AAGInfo)
	model.Name = core.StringPtr("testString")
	model.ObjectID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsAAGInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsHostInformationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["name"] = "testString"
		model["environment"] = "kPhysical"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.HostInformation)
	model.ID = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.Environment = core.StringPtr("kPhysical")

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsHostInformationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsSearchObjectPhysicalParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enable_system_backup"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SearchObjectPhysicalParams)
	model.EnableSystemBackup = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsSearchObjectPhysicalParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsTagInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["tag_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.TagInfo)
	model.TagID = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsTagInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsSnapshotTagInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["tag_id"] = "testString"
		model["run_ids"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SnapshotTagInfo)
	model.TagID = core.StringPtr("testString")
	model.RunIds = []string{"testString"}

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsSnapshotTagInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsHeliosTagInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["category"] = "Security"
		model["name"] = "testString"
		model["sub_category"] = "Classification"
		model["third_party_name"] = "testString"
		model["type"] = "System"
		model["ui_color"] = "testString"
		model["updated_time_usecs"] = int(26)
		model["uuid"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.HeliosTagInfo)
	model.Category = core.StringPtr("Security")
	model.Name = core.StringPtr("testString")
	model.SubCategory = core.StringPtr("Classification")
	model.ThirdPartyName = core.StringPtr("testString")
	model.Type = core.StringPtr("System")
	model.UiColor = core.StringPtr("testString")
	model.UpdatedTimeUsecs = core.Int64Ptr(int64(26))
	model.UUID = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsHeliosTagInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsSearchObjectSourceInfoToMap(t *testing.T) {
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

		objectProtectionStatsSummaryModel := make(map[string]interface{})
		objectProtectionStatsSummaryModel["environment"] = "kPhysical"
		objectProtectionStatsSummaryModel["protected_count"] = int(26)
		objectProtectionStatsSummaryModel["unprotected_count"] = int(26)
		objectProtectionStatsSummaryModel["deleted_protected_count"] = int(26)
		objectProtectionStatsSummaryModel["protected_size_bytes"] = int(26)
		objectProtectionStatsSummaryModel["unprotected_size_bytes"] = int(26)

		userModel := make(map[string]interface{})
		userModel["name"] = "testString"
		userModel["sid"] = "testString"
		userModel["domain"] = "testString"

		groupModel := make(map[string]interface{})
		groupModel["name"] = "testString"
		groupModel["sid"] = "testString"
		groupModel["domain"] = "testString"

		externalVendorCustomPropertiesModel := make(map[string]interface{})
		externalVendorCustomPropertiesModel["key"] = "testString"
		externalVendorCustomPropertiesModel["value"] = "testString"

		ibmTenantMetadataParamsModel := make(map[string]interface{})
		ibmTenantMetadataParamsModel["account_id"] = "testString"
		ibmTenantMetadataParamsModel["crn"] = "testString"
		ibmTenantMetadataParamsModel["custom_properties"] = []map[string]interface{}{externalVendorCustomPropertiesModel}
		ibmTenantMetadataParamsModel["liveness_mode"] = "Active"
		ibmTenantMetadataParamsModel["ownership_mode"] = "Primary"
		ibmTenantMetadataParamsModel["resource_group_id"] = "testString"

		externalVendorTenantMetadataModel := make(map[string]interface{})
		externalVendorTenantMetadataModel["ibm_tenant_metadata_params"] = []map[string]interface{}{ibmTenantMetadataParamsModel}
		externalVendorTenantMetadataModel["type"] = "IBM"

		tenantNetworkModel := make(map[string]interface{})
		tenantNetworkModel["connector_enabled"] = true
		tenantNetworkModel["cluster_hostname"] = "testString"
		tenantNetworkModel["cluster_ips"] = []string{"testString"}

		tenantModel := make(map[string]interface{})
		tenantModel["description"] = "testString"
		tenantModel["external_vendor_metadata"] = []map[string]interface{}{externalVendorTenantMetadataModel}
		tenantModel["id"] = "testString"
		tenantModel["is_managed_on_helios"] = true
		tenantModel["name"] = "testString"
		tenantModel["network"] = []map[string]interface{}{tenantNetworkModel}
		tenantModel["status"] = "Active"

		permissionInfoModel := make(map[string]interface{})
		permissionInfoModel["object_id"] = int(26)
		permissionInfoModel["users"] = []map[string]interface{}{userModel}
		permissionInfoModel["groups"] = []map[string]interface{}{groupModel}
		permissionInfoModel["tenant"] = []map[string]interface{}{tenantModel}

		aagInfoModel := make(map[string]interface{})
		aagInfoModel["name"] = "testString"
		aagInfoModel["object_id"] = int(26)

		hostInformationModel := make(map[string]interface{})
		hostInformationModel["id"] = "testString"
		hostInformationModel["name"] = "testString"
		hostInformationModel["environment"] = "kPhysical"

		searchObjectSourceInfoMssqlParamsModel := make(map[string]interface{})
		searchObjectSourceInfoMssqlParamsModel["aag_info"] = []map[string]interface{}{aagInfoModel}
		searchObjectSourceInfoMssqlParamsModel["host_info"] = []map[string]interface{}{hostInformationModel}
		searchObjectSourceInfoMssqlParamsModel["is_encrypted"] = true

		searchObjectSourceInfoPhysicalParamsModel := make(map[string]interface{})
		searchObjectSourceInfoPhysicalParamsModel["enable_system_backup"] = true

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
		model["protection_stats"] = []map[string]interface{}{objectProtectionStatsSummaryModel}
		model["permissions"] = []map[string]interface{}{permissionInfoModel}
		model["mssql_params"] = []map[string]interface{}{searchObjectSourceInfoMssqlParamsModel}
		model["physical_params"] = []map[string]interface{}{searchObjectSourceInfoPhysicalParamsModel}

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

	objectProtectionStatsSummaryModel := new(backuprecoveryv1.ObjectProtectionStatsSummary)
	objectProtectionStatsSummaryModel.Environment = core.StringPtr("kPhysical")
	objectProtectionStatsSummaryModel.ProtectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.UnprotectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.DeletedProtectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.ProtectedSizeBytes = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.UnprotectedSizeBytes = core.Int64Ptr(int64(26))

	userModel := new(backuprecoveryv1.User)
	userModel.Name = core.StringPtr("testString")
	userModel.Sid = core.StringPtr("testString")
	userModel.Domain = core.StringPtr("testString")

	groupModel := new(backuprecoveryv1.Group)
	groupModel.Name = core.StringPtr("testString")
	groupModel.Sid = core.StringPtr("testString")
	groupModel.Domain = core.StringPtr("testString")

	externalVendorCustomPropertiesModel := new(backuprecoveryv1.ExternalVendorCustomProperties)
	externalVendorCustomPropertiesModel.Key = core.StringPtr("testString")
	externalVendorCustomPropertiesModel.Value = core.StringPtr("testString")

	ibmTenantMetadataParamsModel := new(backuprecoveryv1.IbmTenantMetadataParams)
	ibmTenantMetadataParamsModel.AccountID = core.StringPtr("testString")
	ibmTenantMetadataParamsModel.Crn = core.StringPtr("testString")
	ibmTenantMetadataParamsModel.CustomProperties = []backuprecoveryv1.ExternalVendorCustomProperties{*externalVendorCustomPropertiesModel}
	ibmTenantMetadataParamsModel.LivenessMode = core.StringPtr("Active")
	ibmTenantMetadataParamsModel.OwnershipMode = core.StringPtr("Primary")
	ibmTenantMetadataParamsModel.ResourceGroupID = core.StringPtr("testString")

	externalVendorTenantMetadataModel := new(backuprecoveryv1.ExternalVendorTenantMetadata)
	externalVendorTenantMetadataModel.IbmTenantMetadataParams = ibmTenantMetadataParamsModel
	externalVendorTenantMetadataModel.Type = core.StringPtr("IBM")

	tenantNetworkModel := new(backuprecoveryv1.TenantNetwork)
	tenantNetworkModel.ConnectorEnabled = core.BoolPtr(true)
	tenantNetworkModel.ClusterHostname = core.StringPtr("testString")
	tenantNetworkModel.ClusterIps = []string{"testString"}

	tenantModel := new(backuprecoveryv1.Tenant)
	tenantModel.Description = core.StringPtr("testString")
	tenantModel.ExternalVendorMetadata = externalVendorTenantMetadataModel
	tenantModel.ID = core.StringPtr("testString")
	tenantModel.IsManagedOnHelios = core.BoolPtr(true)
	tenantModel.Name = core.StringPtr("testString")
	tenantModel.Network = tenantNetworkModel
	tenantModel.Status = core.StringPtr("Active")

	permissionInfoModel := new(backuprecoveryv1.PermissionInfo)
	permissionInfoModel.ObjectID = core.Int64Ptr(int64(26))
	permissionInfoModel.Users = []backuprecoveryv1.User{*userModel}
	permissionInfoModel.Groups = []backuprecoveryv1.Group{*groupModel}
	permissionInfoModel.Tenant = tenantModel

	aagInfoModel := new(backuprecoveryv1.AAGInfo)
	aagInfoModel.Name = core.StringPtr("testString")
	aagInfoModel.ObjectID = core.Int64Ptr(int64(26))

	hostInformationModel := new(backuprecoveryv1.HostInformation)
	hostInformationModel.ID = core.StringPtr("testString")
	hostInformationModel.Name = core.StringPtr("testString")
	hostInformationModel.Environment = core.StringPtr("kPhysical")

	searchObjectSourceInfoMssqlParamsModel := new(backuprecoveryv1.SearchObjectSourceInfoMssqlParams)
	searchObjectSourceInfoMssqlParamsModel.AagInfo = aagInfoModel
	searchObjectSourceInfoMssqlParamsModel.HostInfo = hostInformationModel
	searchObjectSourceInfoMssqlParamsModel.IsEncrypted = core.BoolPtr(true)

	searchObjectSourceInfoPhysicalParamsModel := new(backuprecoveryv1.SearchObjectSourceInfoPhysicalParams)
	searchObjectSourceInfoPhysicalParamsModel.EnableSystemBackup = core.BoolPtr(true)

	model := new(backuprecoveryv1.SearchObjectSourceInfo)
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
	model.ProtectionStats = []backuprecoveryv1.ObjectProtectionStatsSummary{*objectProtectionStatsSummaryModel}
	model.Permissions = permissionInfoModel
	model.MssqlParams = searchObjectSourceInfoMssqlParamsModel
	model.PhysicalParams = searchObjectSourceInfoPhysicalParamsModel

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsSearchObjectSourceInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsSearchObjectSourceInfoMssqlParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		aagInfoModel := make(map[string]interface{})
		aagInfoModel["name"] = "testString"
		aagInfoModel["object_id"] = int(26)

		hostInformationModel := make(map[string]interface{})
		hostInformationModel["id"] = "testString"
		hostInformationModel["name"] = "testString"
		hostInformationModel["environment"] = "kPhysical"

		model := make(map[string]interface{})
		model["aag_info"] = []map[string]interface{}{aagInfoModel}
		model["host_info"] = []map[string]interface{}{hostInformationModel}
		model["is_encrypted"] = true

		assert.Equal(t, result, model)
	}

	aagInfoModel := new(backuprecoveryv1.AAGInfo)
	aagInfoModel.Name = core.StringPtr("testString")
	aagInfoModel.ObjectID = core.Int64Ptr(int64(26))

	hostInformationModel := new(backuprecoveryv1.HostInformation)
	hostInformationModel.ID = core.StringPtr("testString")
	hostInformationModel.Name = core.StringPtr("testString")
	hostInformationModel.Environment = core.StringPtr("kPhysical")

	model := new(backuprecoveryv1.SearchObjectSourceInfoMssqlParams)
	model.AagInfo = aagInfoModel
	model.HostInfo = hostInformationModel
	model.IsEncrypted = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsSearchObjectSourceInfoMssqlParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsSearchObjectSourceInfoPhysicalParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enable_system_backup"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SearchObjectSourceInfoPhysicalParams)
	model.EnableSystemBackup = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsSearchObjectSourceInfoPhysicalParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsObjectProtectionInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		objectProtectionGroupSummaryModel := make(map[string]interface{})
		objectProtectionGroupSummaryModel["name"] = "testString"
		objectProtectionGroupSummaryModel["id"] = "testString"
		objectProtectionGroupSummaryModel["protection_env_type"] = "kAgent"
		objectProtectionGroupSummaryModel["policy_name"] = "testString"
		objectProtectionGroupSummaryModel["policy_id"] = "testString"
		objectProtectionGroupSummaryModel["last_backup_run_status"] = "Accepted"
		objectProtectionGroupSummaryModel["last_archival_run_status"] = "Accepted"
		objectProtectionGroupSummaryModel["last_replication_run_status"] = "Accepted"
		objectProtectionGroupSummaryModel["last_run_sla_violated"] = true

		protectionSummaryModel := make(map[string]interface{})
		protectionSummaryModel["policy_name"] = "testString"
		protectionSummaryModel["policy_id"] = "testString"
		protectionSummaryModel["last_backup_run_status"] = "Accepted"
		protectionSummaryModel["last_archival_run_status"] = "Accepted"
		protectionSummaryModel["last_replication_run_status"] = "Accepted"
		protectionSummaryModel["last_run_sla_violated"] = true

		model := make(map[string]interface{})
		model["object_id"] = int(26)
		model["source_id"] = int(26)
		model["view_id"] = int(26)
		model["region_id"] = "testString"
		model["cluster_id"] = int(26)
		model["cluster_incarnation_id"] = int(26)
		model["tenant_ids"] = []string{"testString"}
		model["is_deleted"] = true
		model["protection_groups"] = []map[string]interface{}{objectProtectionGroupSummaryModel}
		model["object_backup_configuration"] = []map[string]interface{}{protectionSummaryModel}
		model["last_run_status"] = "Accepted"

		assert.Equal(t, result, model)
	}

	objectProtectionGroupSummaryModel := new(backuprecoveryv1.ObjectProtectionGroupSummary)
	objectProtectionGroupSummaryModel.Name = core.StringPtr("testString")
	objectProtectionGroupSummaryModel.ID = core.StringPtr("testString")
	objectProtectionGroupSummaryModel.ProtectionEnvType = core.StringPtr("kAgent")
	objectProtectionGroupSummaryModel.PolicyName = core.StringPtr("testString")
	objectProtectionGroupSummaryModel.PolicyID = core.StringPtr("testString")
	objectProtectionGroupSummaryModel.LastBackupRunStatus = core.StringPtr("Accepted")
	objectProtectionGroupSummaryModel.LastArchivalRunStatus = core.StringPtr("Accepted")
	objectProtectionGroupSummaryModel.LastReplicationRunStatus = core.StringPtr("Accepted")
	objectProtectionGroupSummaryModel.LastRunSlaViolated = core.BoolPtr(true)

	protectionSummaryModel := new(backuprecoveryv1.ProtectionSummary)
	protectionSummaryModel.PolicyName = core.StringPtr("testString")
	protectionSummaryModel.PolicyID = core.StringPtr("testString")
	protectionSummaryModel.LastBackupRunStatus = core.StringPtr("Accepted")
	protectionSummaryModel.LastArchivalRunStatus = core.StringPtr("Accepted")
	protectionSummaryModel.LastReplicationRunStatus = core.StringPtr("Accepted")
	protectionSummaryModel.LastRunSlaViolated = core.BoolPtr(true)

	model := new(backuprecoveryv1.ObjectProtectionInfo)
	model.ObjectID = core.Int64Ptr(int64(26))
	model.SourceID = core.Int64Ptr(int64(26))
	model.ViewID = core.Int64Ptr(int64(26))
	model.RegionID = core.StringPtr("testString")
	model.ClusterID = core.Int64Ptr(int64(26))
	model.ClusterIncarnationID = core.Int64Ptr(int64(26))
	model.TenantIds = []string{"testString"}
	model.IsDeleted = core.BoolPtr(true)
	model.ProtectionGroups = []backuprecoveryv1.ObjectProtectionGroupSummary{*objectProtectionGroupSummaryModel}
	model.ObjectBackupConfiguration = []backuprecoveryv1.ProtectionSummary{*protectionSummaryModel}
	model.LastRunStatus = core.StringPtr("Accepted")

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsObjectProtectionInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsObjectProtectionGroupSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["id"] = "testString"
		model["protection_env_type"] = "kAgent"
		model["policy_name"] = "testString"
		model["policy_id"] = "testString"
		model["last_backup_run_status"] = "Accepted"
		model["last_archival_run_status"] = "Accepted"
		model["last_replication_run_status"] = "Accepted"
		model["last_run_sla_violated"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectProtectionGroupSummary)
	model.Name = core.StringPtr("testString")
	model.ID = core.StringPtr("testString")
	model.ProtectionEnvType = core.StringPtr("kAgent")
	model.PolicyName = core.StringPtr("testString")
	model.PolicyID = core.StringPtr("testString")
	model.LastBackupRunStatus = core.StringPtr("Accepted")
	model.LastArchivalRunStatus = core.StringPtr("Accepted")
	model.LastReplicationRunStatus = core.StringPtr("Accepted")
	model.LastRunSlaViolated = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsObjectProtectionGroupSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsProtectionSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["policy_name"] = "testString"
		model["policy_id"] = "testString"
		model["last_backup_run_status"] = "Accepted"
		model["last_archival_run_status"] = "Accepted"
		model["last_replication_run_status"] = "Accepted"
		model["last_run_sla_violated"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ProtectionSummary)
	model.PolicyName = core.StringPtr("testString")
	model.PolicyID = core.StringPtr("testString")
	model.LastBackupRunStatus = core.StringPtr("Accepted")
	model.LastArchivalRunStatus = core.StringPtr("Accepted")
	model.LastReplicationRunStatus = core.StringPtr("Accepted")
	model.LastRunSlaViolated = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsProtectionSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsSecondaryIDToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SecondaryID)
	model.Name = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsSecondaryIDToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSearchObjectsTaggedSnapshotInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		heliosTagInfoModel := make(map[string]interface{})
		heliosTagInfoModel["category"] = "Security"
		heliosTagInfoModel["name"] = "testString"
		heliosTagInfoModel["sub_category"] = "Classification"
		heliosTagInfoModel["third_party_name"] = "testString"
		heliosTagInfoModel["type"] = "System"
		heliosTagInfoModel["ui_color"] = "testString"
		heliosTagInfoModel["updated_time_usecs"] = int(26)
		heliosTagInfoModel["uuid"] = "testString"

		model := make(map[string]interface{})
		model["cluster_id"] = int(26)
		model["cluster_incarnation_id"] = int(26)
		model["job_id"] = int(26)
		model["object_uuid"] = "testString"
		model["run_start_time_usecs"] = int(26)
		model["tags"] = []map[string]interface{}{heliosTagInfoModel}

		assert.Equal(t, result, model)
	}

	heliosTagInfoModel := new(backuprecoveryv1.HeliosTagInfo)
	heliosTagInfoModel.Category = core.StringPtr("Security")
	heliosTagInfoModel.Name = core.StringPtr("testString")
	heliosTagInfoModel.SubCategory = core.StringPtr("Classification")
	heliosTagInfoModel.ThirdPartyName = core.StringPtr("testString")
	heliosTagInfoModel.Type = core.StringPtr("System")
	heliosTagInfoModel.UiColor = core.StringPtr("testString")
	heliosTagInfoModel.UpdatedTimeUsecs = core.Int64Ptr(int64(26))
	heliosTagInfoModel.UUID = core.StringPtr("testString")

	model := new(backuprecoveryv1.TaggedSnapshotInfo)
	model.ClusterID = core.Int64Ptr(int64(26))
	model.ClusterIncarnationID = core.Int64Ptr(int64(26))
	model.JobID = core.Int64Ptr(int64(26))
	model.ObjectUUID = core.StringPtr("testString")
	model.RunStartTimeUsecs = core.Int64Ptr(int64(26))
	model.Tags = []backuprecoveryv1.HeliosTagInfo{*heliosTagInfoModel}

	result, err := backuprecovery.DataSourceIbmBaasSearchObjectsTaggedSnapshotInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
