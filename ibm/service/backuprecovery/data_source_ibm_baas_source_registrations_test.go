// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasSourceRegistrationsDataSourceBasic(t *testing.T) {
	sourceRegistrationReponseParamsXIBMTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	sourceRegistrationReponseParamsEnvironment := "kPhysical"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasSourceRegistrationsDataSourceConfigBasic(sourceRegistrationReponseParamsXIBMTenantID, sourceRegistrationReponseParamsEnvironment),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func TestAccIbmBaasSourceRegistrationsDataSourceAllArgs(t *testing.T) {
	sourceRegistrationReponseParamsXIBMTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	sourceRegistrationReponseParamsEnvironment := "kPhysical"
	sourceRegistrationReponseParamsName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	sourceRegistrationReponseParamsConnectionID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	sourceRegistrationReponseParamsConnectorGroupID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	sourceRegistrationReponseParamsDataSourceConnectionID := fmt.Sprintf("tf_data_source_connection_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasSourceRegistrationsDataSourceConfig(sourceRegistrationReponseParamsXIBMTenantID, sourceRegistrationReponseParamsEnvironment, sourceRegistrationReponseParamsName, sourceRegistrationReponseParamsConnectionID, sourceRegistrationReponseParamsConnectorGroupID, sourceRegistrationReponseParamsDataSourceConnectionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "x_ibm_tenant_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "ids"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "include_source_credentials"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "encryption_key"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "use_cached_data"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "include_external_metadata"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "ignore_tenant_migration_in_progress_check"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.#"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.source_id"),
					resource.TestCheckResourceAttr("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.environment", sourceRegistrationReponseParamsEnvironment),
					resource.TestCheckResourceAttr("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.name", sourceRegistrationReponseParamsName),
					resource.TestCheckResourceAttr("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.connection_id", sourceRegistrationReponseParamsConnectionID),
					resource.TestCheckResourceAttr("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.connector_group_id", sourceRegistrationReponseParamsConnectorGroupID),
					resource.TestCheckResourceAttr("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.data_source_connection_id", sourceRegistrationReponseParamsDataSourceConnectionID),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.authentication_status"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.registration_time_msecs"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.last_refreshed_time_msecs"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasSourceRegistrationsDataSourceConfigBasic(sourceRegistrationReponseParamsXIBMTenantID string, sourceRegistrationReponseParamsEnvironment string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_source_registration" "baas_source_registration_instance" {
			x_ibm_tenant_id = "%s"
			environment = "%s"
		}

		data "ibm_baas_source_registrations" "baas_source_registrations_instance" {
			x_ibm_tenant_id = ibm_baas_source_registration.baas_source_registration_instance.x_ibm_tenant_id
			ids = [ 1 ]
			include_source_credentials = true
			encryption_key = "encryption_key"
			use_cached_data = true
			include_external_metadata = true
			ignore_tenant_migration_in_progress_check = true
		}
	`, sourceRegistrationReponseParamsXIBMTenantID, sourceRegistrationReponseParamsEnvironment)
}

func testAccCheckIbmBaasSourceRegistrationsDataSourceConfig(sourceRegistrationReponseParamsXIBMTenantID string, sourceRegistrationReponseParamsEnvironment string, sourceRegistrationReponseParamsName string, sourceRegistrationReponseParamsConnectionID string, sourceRegistrationReponseParamsConnectorGroupID string, sourceRegistrationReponseParamsDataSourceConnectionID string) string {
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

		data "ibm_baas_source_registrations" "baas_source_registrations_instance" {
			x_ibm_tenant_id = ibm_baas_source_registration.baas_source_registration_instance.x_ibm_tenant_id
			ids = [ 1 ]
			include_source_credentials = true
			encryption_key = "encryption_key"
			use_cached_data = true
			include_external_metadata = true
			ignore_tenant_migration_in_progress_check = true
		}
	`, sourceRegistrationReponseParamsXIBMTenantID, sourceRegistrationReponseParamsEnvironment, sourceRegistrationReponseParamsName, sourceRegistrationReponseParamsConnectionID, sourceRegistrationReponseParamsConnectorGroupID, sourceRegistrationReponseParamsDataSourceConnectionID)
}

func TestDataSourceIbmBaasSourceRegistrationsSourceRegistrationReponseParamsToMap(t *testing.T) {
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

		objectMssqlParamsModel := make(map[string]interface{})
		objectMssqlParamsModel["aag_info"] = []map[string]interface{}{aagInfoModel}
		objectMssqlParamsModel["host_info"] = []map[string]interface{}{hostInformationModel}
		objectMssqlParamsModel["is_encrypted"] = true

		objectPhysicalParamsModel := make(map[string]interface{})
		objectPhysicalParamsModel["enable_system_backup"] = true

		objectModel := make(map[string]interface{})
		objectModel["id"] = int(26)
		objectModel["name"] = "testString"
		objectModel["source_id"] = int(26)
		objectModel["source_name"] = "testString"
		objectModel["environment"] = "kPhysical"
		objectModel["object_hash"] = "testString"
		objectModel["object_type"] = "kCluster"
		objectModel["logical_size_bytes"] = int(26)
		objectModel["uuid"] = "testString"
		objectModel["global_id"] = "testString"
		objectModel["protection_type"] = "kAgent"
		objectModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		objectModel["os_type"] = "kLinux"
		objectModel["child_objects"] = []map[string]interface{}{objectSummaryModel}
		objectModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		objectModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}
		objectModel["protection_stats"] = []map[string]interface{}{objectProtectionStatsSummaryModel}
		objectModel["permissions"] = []map[string]interface{}{permissionInfoModel}
		objectModel["mssql_params"] = []map[string]interface{}{objectMssqlParamsModel}
		objectModel["physical_params"] = []map[string]interface{}{objectPhysicalParamsModel}

		connectionConfigModel := make(map[string]interface{})
		connectionConfigModel["connection_id"] = int(26)
		connectionConfigModel["entity_id"] = int(26)
		connectionConfigModel["connector_group_id"] = int(26)
		connectionConfigModel["data_source_connection_id"] = "testString"

		keyValuePairModel := make(map[string]interface{})
		keyValuePairModel["key"] = "testString"
		keyValuePairModel["value"] = "testString"

		timeRangeUsecsModel := make(map[string]interface{})
		timeRangeUsecsModel["end_time_usecs"] = int(26)
		timeRangeUsecsModel["start_time_usecs"] = int(26)

		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		timeWindowModel := make(map[string]interface{})
		timeWindowModel["day_of_the_week"] = "Sunday"
		timeWindowModel["end_time"] = []map[string]interface{}{timeModel}
		timeWindowModel["start_time"] = []map[string]interface{}{timeModel}

		scheduleModel := make(map[string]interface{})
		scheduleModel["periodic_time_windows"] = []map[string]interface{}{timeWindowModel}
		scheduleModel["schedule_type"] = "PeriodicTimeWindows"
		scheduleModel["time_ranges"] = []map[string]interface{}{timeRangeUsecsModel}
		scheduleModel["timezone"] = "testString"

		workflowInterventionSpecModel := make(map[string]interface{})
		workflowInterventionSpecModel["intervention"] = "NoIntervention"
		workflowInterventionSpecModel["workflow_type"] = "BackupRun"

		maintenanceModeConfigModel := make(map[string]interface{})
		maintenanceModeConfigModel["activation_time_intervals"] = []map[string]interface{}{timeRangeUsecsModel}
		maintenanceModeConfigModel["maintenance_schedule"] = []map[string]interface{}{scheduleModel}
		maintenanceModeConfigModel["user_message"] = "testString"
		maintenanceModeConfigModel["workflow_intervention_spec_list"] = []map[string]interface{}{workflowInterventionSpecModel}

		entityExternalMetadataModel := make(map[string]interface{})
		entityExternalMetadataModel["maintenance_mode_config"] = []map[string]interface{}{maintenanceModeConfigModel}

		physicalSourceRegistrationParamsModel := make(map[string]interface{})
		physicalSourceRegistrationParamsModel["endpoint"] = "testString"
		physicalSourceRegistrationParamsModel["force_register"] = true
		physicalSourceRegistrationParamsModel["host_type"] = "kLinux"
		physicalSourceRegistrationParamsModel["physical_type"] = "kGroup"
		physicalSourceRegistrationParamsModel["applications"] = []string{"kSQL"}

		model := make(map[string]interface{})
		model["id"] = int(26)
		model["source_id"] = int(26)
		model["source_info"] = []map[string]interface{}{objectModel}
		model["environment"] = "kPhysical"
		model["name"] = "testString"
		model["connection_id"] = int(26)
		model["connections"] = []map[string]interface{}{connectionConfigModel}
		model["connector_group_id"] = int(26)
		model["data_source_connection_id"] = "testString"
		model["advanced_configs"] = []map[string]interface{}{keyValuePairModel}
		model["authentication_status"] = "Pending"
		model["registration_time_msecs"] = int(26)
		model["last_refreshed_time_msecs"] = int(26)
		model["external_metadata"] = []map[string]interface{}{entityExternalMetadataModel}
		model["physical_params"] = []map[string]interface{}{physicalSourceRegistrationParamsModel}

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

	objectMssqlParamsModel := new(backuprecoveryv1.ObjectMssqlParams)
	objectMssqlParamsModel.AagInfo = aagInfoModel
	objectMssqlParamsModel.HostInfo = hostInformationModel
	objectMssqlParamsModel.IsEncrypted = core.BoolPtr(true)

	objectPhysicalParamsModel := new(backuprecoveryv1.ObjectPhysicalParams)
	objectPhysicalParamsModel.EnableSystemBackup = core.BoolPtr(true)

	objectModel := new(backuprecoveryv1.Object)
	objectModel.ID = core.Int64Ptr(int64(26))
	objectModel.Name = core.StringPtr("testString")
	objectModel.SourceID = core.Int64Ptr(int64(26))
	objectModel.SourceName = core.StringPtr("testString")
	objectModel.Environment = core.StringPtr("kPhysical")
	objectModel.ObjectHash = core.StringPtr("testString")
	objectModel.ObjectType = core.StringPtr("kCluster")
	objectModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	objectModel.UUID = core.StringPtr("testString")
	objectModel.GlobalID = core.StringPtr("testString")
	objectModel.ProtectionType = core.StringPtr("kAgent")
	objectModel.SharepointSiteSummary = sharepointObjectParamsModel
	objectModel.OsType = core.StringPtr("kLinux")
	objectModel.ChildObjects = []backuprecoveryv1.ObjectSummary{*objectSummaryModel}
	objectModel.VCenterSummary = objectTypeVCenterParamsModel
	objectModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel
	objectModel.ProtectionStats = []backuprecoveryv1.ObjectProtectionStatsSummary{*objectProtectionStatsSummaryModel}
	objectModel.Permissions = permissionInfoModel
	objectModel.MssqlParams = objectMssqlParamsModel
	objectModel.PhysicalParams = objectPhysicalParamsModel

	connectionConfigModel := new(backuprecoveryv1.ConnectionConfig)
	connectionConfigModel.ConnectionID = core.Int64Ptr(int64(26))
	connectionConfigModel.EntityID = core.Int64Ptr(int64(26))
	connectionConfigModel.ConnectorGroupID = core.Int64Ptr(int64(26))
	connectionConfigModel.DataSourceConnectionID = core.StringPtr("testString")

	keyValuePairModel := new(backuprecoveryv1.KeyValuePair)
	keyValuePairModel.Key = core.StringPtr("testString")
	keyValuePairModel.Value = core.StringPtr("testString")

	timeRangeUsecsModel := new(backuprecoveryv1.TimeRangeUsecs)
	timeRangeUsecsModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	timeRangeUsecsModel.StartTimeUsecs = core.Int64Ptr(int64(26))

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	timeWindowModel := new(backuprecoveryv1.TimeWindow)
	timeWindowModel.DayOfTheWeek = core.StringPtr("Sunday")
	timeWindowModel.EndTime = timeModel
	timeWindowModel.StartTime = timeModel

	scheduleModel := new(backuprecoveryv1.Schedule)
	scheduleModel.PeriodicTimeWindows = []backuprecoveryv1.TimeWindow{*timeWindowModel}
	scheduleModel.ScheduleType = core.StringPtr("PeriodicTimeWindows")
	scheduleModel.TimeRanges = []backuprecoveryv1.TimeRangeUsecs{*timeRangeUsecsModel}
	scheduleModel.Timezone = core.StringPtr("testString")

	workflowInterventionSpecModel := new(backuprecoveryv1.WorkflowInterventionSpec)
	workflowInterventionSpecModel.Intervention = core.StringPtr("NoIntervention")
	workflowInterventionSpecModel.WorkflowType = core.StringPtr("BackupRun")

	maintenanceModeConfigModel := new(backuprecoveryv1.MaintenanceModeConfig)
	maintenanceModeConfigModel.ActivationTimeIntervals = []backuprecoveryv1.TimeRangeUsecs{*timeRangeUsecsModel}
	maintenanceModeConfigModel.MaintenanceSchedule = scheduleModel
	maintenanceModeConfigModel.UserMessage = core.StringPtr("testString")
	maintenanceModeConfigModel.WorkflowInterventionSpecList = []backuprecoveryv1.WorkflowInterventionSpec{*workflowInterventionSpecModel}

	entityExternalMetadataModel := new(backuprecoveryv1.EntityExternalMetadata)
	entityExternalMetadataModel.MaintenanceModeConfig = maintenanceModeConfigModel

	physicalSourceRegistrationParamsModel := new(backuprecoveryv1.PhysicalSourceRegistrationParams)
	physicalSourceRegistrationParamsModel.Endpoint = core.StringPtr("testString")
	physicalSourceRegistrationParamsModel.ForceRegister = core.BoolPtr(true)
	physicalSourceRegistrationParamsModel.HostType = core.StringPtr("kLinux")
	physicalSourceRegistrationParamsModel.PhysicalType = core.StringPtr("kGroup")
	physicalSourceRegistrationParamsModel.Applications = []string{"kSQL"}

	model := new(backuprecoveryv1.SourceRegistrationReponseParams)
	model.ID = core.Int64Ptr(int64(26))
	model.SourceID = core.Int64Ptr(int64(26))
	model.SourceInfo = objectModel
	model.Environment = core.StringPtr("kPhysical")
	model.Name = core.StringPtr("testString")
	model.ConnectionID = core.Int64Ptr(int64(26))
	model.Connections = []backuprecoveryv1.ConnectionConfig{*connectionConfigModel}
	model.ConnectorGroupID = core.Int64Ptr(int64(26))
	model.DataSourceConnectionID = core.StringPtr("testString")
	model.AdvancedConfigs = []backuprecoveryv1.KeyValuePair{*keyValuePairModel}
	model.AuthenticationStatus = core.StringPtr("Pending")
	model.RegistrationTimeMsecs = core.Int64Ptr(int64(26))
	model.LastRefreshedTimeMsecs = core.Int64Ptr(int64(26))
	model.ExternalMetadata = entityExternalMetadataModel
	model.PhysicalParams = physicalSourceRegistrationParamsModel

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsSourceRegistrationReponseParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsObjectToMap(t *testing.T) {
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

		objectMssqlParamsModel := make(map[string]interface{})
		objectMssqlParamsModel["aag_info"] = []map[string]interface{}{aagInfoModel}
		objectMssqlParamsModel["host_info"] = []map[string]interface{}{hostInformationModel}
		objectMssqlParamsModel["is_encrypted"] = true

		objectPhysicalParamsModel := make(map[string]interface{})
		objectPhysicalParamsModel["enable_system_backup"] = true

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
		model["mssql_params"] = []map[string]interface{}{objectMssqlParamsModel}
		model["physical_params"] = []map[string]interface{}{objectPhysicalParamsModel}

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

	objectMssqlParamsModel := new(backuprecoveryv1.ObjectMssqlParams)
	objectMssqlParamsModel.AagInfo = aagInfoModel
	objectMssqlParamsModel.HostInfo = hostInformationModel
	objectMssqlParamsModel.IsEncrypted = core.BoolPtr(true)

	objectPhysicalParamsModel := new(backuprecoveryv1.ObjectPhysicalParams)
	objectPhysicalParamsModel.EnableSystemBackup = core.BoolPtr(true)

	model := new(backuprecoveryv1.Object)
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
	model.MssqlParams = objectMssqlParamsModel
	model.PhysicalParams = objectPhysicalParamsModel

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsObjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsSharepointObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["site_web_url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SharepointObjectParams)
	model.SiteWebURL = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsSharepointObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsObjectSummaryToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsObjectSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsObjectTypeVCenterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["is_cloud_env"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeVCenterParams)
	model.IsCloudEnv = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsObjectTypeVCenterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsObjectTypeWindowsClusterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cluster_source_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	model.ClusterSourceType = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsObjectTypeWindowsClusterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsObjectProtectionStatsSummaryToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsObjectProtectionStatsSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsPermissionInfoToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsPermissionInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsUserToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsUserToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsGroupToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsTenantToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsTenantToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsExternalVendorTenantMetadataToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsExternalVendorTenantMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsIbmTenantMetadataParamsToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsIbmTenantMetadataParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsExternalVendorCustomPropertiesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ExternalVendorCustomProperties)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsExternalVendorCustomPropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsTenantNetworkToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsTenantNetworkToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsObjectMssqlParamsToMap(t *testing.T) {
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

	model := new(backuprecoveryv1.ObjectMssqlParams)
	model.AagInfo = aagInfoModel
	model.HostInfo = hostInformationModel
	model.IsEncrypted = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsObjectMssqlParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsAAGInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["object_id"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AAGInfo)
	model.Name = core.StringPtr("testString")
	model.ObjectID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsAAGInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsHostInformationToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsHostInformationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsObjectPhysicalParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enable_system_backup"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectPhysicalParams)
	model.EnableSystemBackup = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsObjectPhysicalParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsConnectionConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["connection_id"] = int(26)
		model["entity_id"] = int(26)
		model["connector_group_id"] = int(26)
		model["data_source_connection_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ConnectionConfig)
	model.ConnectionID = core.Int64Ptr(int64(26))
	model.EntityID = core.Int64Ptr(int64(26))
	model.ConnectorGroupID = core.Int64Ptr(int64(26))
	model.DataSourceConnectionID = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsConnectionConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsKeyValuePairToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.KeyValuePair)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsKeyValuePairToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsEntityExternalMetadataToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeRangeUsecsModel := make(map[string]interface{})
		timeRangeUsecsModel["end_time_usecs"] = int(26)
		timeRangeUsecsModel["start_time_usecs"] = int(26)

		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		timeWindowModel := make(map[string]interface{})
		timeWindowModel["day_of_the_week"] = "Sunday"
		timeWindowModel["end_time"] = []map[string]interface{}{timeModel}
		timeWindowModel["start_time"] = []map[string]interface{}{timeModel}

		scheduleModel := make(map[string]interface{})
		scheduleModel["periodic_time_windows"] = []map[string]interface{}{timeWindowModel}
		scheduleModel["schedule_type"] = "PeriodicTimeWindows"
		scheduleModel["time_ranges"] = []map[string]interface{}{timeRangeUsecsModel}
		scheduleModel["timezone"] = "testString"

		workflowInterventionSpecModel := make(map[string]interface{})
		workflowInterventionSpecModel["intervention"] = "NoIntervention"
		workflowInterventionSpecModel["workflow_type"] = "BackupRun"

		maintenanceModeConfigModel := make(map[string]interface{})
		maintenanceModeConfigModel["activation_time_intervals"] = []map[string]interface{}{timeRangeUsecsModel}
		maintenanceModeConfigModel["maintenance_schedule"] = []map[string]interface{}{scheduleModel}
		maintenanceModeConfigModel["user_message"] = "testString"
		maintenanceModeConfigModel["workflow_intervention_spec_list"] = []map[string]interface{}{workflowInterventionSpecModel}

		model := make(map[string]interface{})
		model["maintenance_mode_config"] = []map[string]interface{}{maintenanceModeConfigModel}

		assert.Equal(t, result, model)
	}

	timeRangeUsecsModel := new(backuprecoveryv1.TimeRangeUsecs)
	timeRangeUsecsModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	timeRangeUsecsModel.StartTimeUsecs = core.Int64Ptr(int64(26))

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	timeWindowModel := new(backuprecoveryv1.TimeWindow)
	timeWindowModel.DayOfTheWeek = core.StringPtr("Sunday")
	timeWindowModel.EndTime = timeModel
	timeWindowModel.StartTime = timeModel

	scheduleModel := new(backuprecoveryv1.Schedule)
	scheduleModel.PeriodicTimeWindows = []backuprecoveryv1.TimeWindow{*timeWindowModel}
	scheduleModel.ScheduleType = core.StringPtr("PeriodicTimeWindows")
	scheduleModel.TimeRanges = []backuprecoveryv1.TimeRangeUsecs{*timeRangeUsecsModel}
	scheduleModel.Timezone = core.StringPtr("testString")

	workflowInterventionSpecModel := new(backuprecoveryv1.WorkflowInterventionSpec)
	workflowInterventionSpecModel.Intervention = core.StringPtr("NoIntervention")
	workflowInterventionSpecModel.WorkflowType = core.StringPtr("BackupRun")

	maintenanceModeConfigModel := new(backuprecoveryv1.MaintenanceModeConfig)
	maintenanceModeConfigModel.ActivationTimeIntervals = []backuprecoveryv1.TimeRangeUsecs{*timeRangeUsecsModel}
	maintenanceModeConfigModel.MaintenanceSchedule = scheduleModel
	maintenanceModeConfigModel.UserMessage = core.StringPtr("testString")
	maintenanceModeConfigModel.WorkflowInterventionSpecList = []backuprecoveryv1.WorkflowInterventionSpec{*workflowInterventionSpecModel}

	model := new(backuprecoveryv1.EntityExternalMetadata)
	model.MaintenanceModeConfig = maintenanceModeConfigModel

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsEntityExternalMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsMaintenanceModeConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeRangeUsecsModel := make(map[string]interface{})
		timeRangeUsecsModel["end_time_usecs"] = int(26)
		timeRangeUsecsModel["start_time_usecs"] = int(26)

		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		timeWindowModel := make(map[string]interface{})
		timeWindowModel["day_of_the_week"] = "Sunday"
		timeWindowModel["end_time"] = []map[string]interface{}{timeModel}
		timeWindowModel["start_time"] = []map[string]interface{}{timeModel}

		scheduleModel := make(map[string]interface{})
		scheduleModel["periodic_time_windows"] = []map[string]interface{}{timeWindowModel}
		scheduleModel["schedule_type"] = "PeriodicTimeWindows"
		scheduleModel["time_ranges"] = []map[string]interface{}{timeRangeUsecsModel}
		scheduleModel["timezone"] = "testString"

		workflowInterventionSpecModel := make(map[string]interface{})
		workflowInterventionSpecModel["intervention"] = "NoIntervention"
		workflowInterventionSpecModel["workflow_type"] = "BackupRun"

		model := make(map[string]interface{})
		model["activation_time_intervals"] = []map[string]interface{}{timeRangeUsecsModel}
		model["maintenance_schedule"] = []map[string]interface{}{scheduleModel}
		model["user_message"] = "testString"
		model["workflow_intervention_spec_list"] = []map[string]interface{}{workflowInterventionSpecModel}

		assert.Equal(t, result, model)
	}

	timeRangeUsecsModel := new(backuprecoveryv1.TimeRangeUsecs)
	timeRangeUsecsModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	timeRangeUsecsModel.StartTimeUsecs = core.Int64Ptr(int64(26))

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	timeWindowModel := new(backuprecoveryv1.TimeWindow)
	timeWindowModel.DayOfTheWeek = core.StringPtr("Sunday")
	timeWindowModel.EndTime = timeModel
	timeWindowModel.StartTime = timeModel

	scheduleModel := new(backuprecoveryv1.Schedule)
	scheduleModel.PeriodicTimeWindows = []backuprecoveryv1.TimeWindow{*timeWindowModel}
	scheduleModel.ScheduleType = core.StringPtr("PeriodicTimeWindows")
	scheduleModel.TimeRanges = []backuprecoveryv1.TimeRangeUsecs{*timeRangeUsecsModel}
	scheduleModel.Timezone = core.StringPtr("testString")

	workflowInterventionSpecModel := new(backuprecoveryv1.WorkflowInterventionSpec)
	workflowInterventionSpecModel.Intervention = core.StringPtr("NoIntervention")
	workflowInterventionSpecModel.WorkflowType = core.StringPtr("BackupRun")

	model := new(backuprecoveryv1.MaintenanceModeConfig)
	model.ActivationTimeIntervals = []backuprecoveryv1.TimeRangeUsecs{*timeRangeUsecsModel}
	model.MaintenanceSchedule = scheduleModel
	model.UserMessage = core.StringPtr("testString")
	model.WorkflowInterventionSpecList = []backuprecoveryv1.WorkflowInterventionSpec{*workflowInterventionSpecModel}

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsMaintenanceModeConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsTimeRangeUsecsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["end_time_usecs"] = int(26)
		model["start_time_usecs"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.TimeRangeUsecs)
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.StartTimeUsecs = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsTimeRangeUsecsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		timeWindowModel := make(map[string]interface{})
		timeWindowModel["day_of_the_week"] = "Sunday"
		timeWindowModel["end_time"] = []map[string]interface{}{timeModel}
		timeWindowModel["start_time"] = []map[string]interface{}{timeModel}

		timeRangeUsecsModel := make(map[string]interface{})
		timeRangeUsecsModel["end_time_usecs"] = int(26)
		timeRangeUsecsModel["start_time_usecs"] = int(26)

		model := make(map[string]interface{})
		model["periodic_time_windows"] = []map[string]interface{}{timeWindowModel}
		model["schedule_type"] = "PeriodicTimeWindows"
		model["time_ranges"] = []map[string]interface{}{timeRangeUsecsModel}
		model["timezone"] = "testString"

		assert.Equal(t, result, model)
	}

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	timeWindowModel := new(backuprecoveryv1.TimeWindow)
	timeWindowModel.DayOfTheWeek = core.StringPtr("Sunday")
	timeWindowModel.EndTime = timeModel
	timeWindowModel.StartTime = timeModel

	timeRangeUsecsModel := new(backuprecoveryv1.TimeRangeUsecs)
	timeRangeUsecsModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	timeRangeUsecsModel.StartTimeUsecs = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.Schedule)
	model.PeriodicTimeWindows = []backuprecoveryv1.TimeWindow{*timeWindowModel}
	model.ScheduleType = core.StringPtr("PeriodicTimeWindows")
	model.TimeRanges = []backuprecoveryv1.TimeRangeUsecs{*timeRangeUsecsModel}
	model.Timezone = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		model := make(map[string]interface{})
		model["day_of_the_week"] = "Sunday"
		model["end_time"] = []map[string]interface{}{timeModel}
		model["start_time"] = []map[string]interface{}{timeModel}

		assert.Equal(t, result, model)
	}

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	model := new(backuprecoveryv1.TimeWindow)
	model.DayOfTheWeek = core.StringPtr("Sunday")
	model.EndTime = timeModel
	model.StartTime = timeModel

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsTimeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hour"] = int(38)
		model["minute"] = int(38)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.Time)
	model.Hour = core.Int64Ptr(int64(38))
	model.Minute = core.Int64Ptr(int64(38))

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsTimeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsWorkflowInterventionSpecToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["intervention"] = "NoIntervention"
		model["workflow_type"] = "BackupRun"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.WorkflowInterventionSpec)
	model.Intervention = core.StringPtr("NoIntervention")
	model.WorkflowType = core.StringPtr("BackupRun")

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsWorkflowInterventionSpecToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasSourceRegistrationsPhysicalSourceRegistrationParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["endpoint"] = "testString"
		model["force_register"] = true
		model["host_type"] = "kLinux"
		model["physical_type"] = "kGroup"
		model["applications"] = []string{"kSQL"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.PhysicalSourceRegistrationParams)
	model.Endpoint = core.StringPtr("testString")
	model.ForceRegister = core.BoolPtr(true)
	model.HostType = core.StringPtr("kLinux")
	model.PhysicalType = core.StringPtr("kGroup")
	model.Applications = []string{"kSQL"}

	result, err := backuprecovery.DataSourceIbmBaasSourceRegistrationsPhysicalSourceRegistrationParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
