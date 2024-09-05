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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
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

func TestResourceIbmBaasSourceRegistrationConnectionConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationConnectionConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationKeyValuePairToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.KeyValuePair)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationKeyValuePairToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationPhysicalSourceRegistrationParamsToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationPhysicalSourceRegistrationParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationObjectToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationObjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationSharepointObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["site_web_url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SharepointObjectParams)
	model.SiteWebURL = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationSharepointObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationObjectSummaryToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationObjectSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationObjectTypeVCenterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["is_cloud_env"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeVCenterParams)
	model.IsCloudEnv = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationObjectTypeVCenterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationObjectTypeWindowsClusterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cluster_source_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	model.ClusterSourceType = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationObjectTypeWindowsClusterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationObjectProtectionStatsSummaryToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationObjectProtectionStatsSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationPermissionInfoToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationPermissionInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationUserToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationUserToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationGroupToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationTenantToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationTenantToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationExternalVendorTenantMetadataToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationExternalVendorTenantMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationIbmTenantMetadataParamsToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationIbmTenantMetadataParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationExternalVendorCustomPropertiesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ExternalVendorCustomProperties)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationExternalVendorCustomPropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationTenantNetworkToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationTenantNetworkToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationObjectMssqlParamsToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationObjectMssqlParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationAAGInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["object_id"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AAGInfo)
	model.Name = core.StringPtr("testString")
	model.ObjectID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationAAGInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationHostInformationToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationHostInformationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationObjectPhysicalParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enable_system_backup"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectPhysicalParams)
	model.EnableSystemBackup = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationObjectPhysicalParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationEntityExternalMetadataToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationEntityExternalMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationMaintenanceModeConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationMaintenanceModeConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationTimeRangeUsecsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["end_time_usecs"] = int(26)
		model["start_time_usecs"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.TimeRangeUsecs)
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.StartTimeUsecs = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationTimeRangeUsecsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationScheduleToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationTimeWindowToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationTimeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hour"] = int(38)
		model["minute"] = int(38)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.Time)
	model.Hour = core.Int64Ptr(int64(38))
	model.Minute = core.Int64Ptr(int64(38))

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationTimeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationWorkflowInterventionSpecToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["intervention"] = "NoIntervention"
		model["workflow_type"] = "BackupRun"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.WorkflowInterventionSpec)
	model.Intervention = core.StringPtr("NoIntervention")
	model.WorkflowType = core.StringPtr("BackupRun")

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationWorkflowInterventionSpecToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationMapToConnectionConfig(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.ConnectionConfig) {
		model := new(backuprecoveryv1.ConnectionConfig)
		model.ConnectionID = core.Int64Ptr(int64(26))
		model.EntityID = core.Int64Ptr(int64(26))
		model.ConnectorGroupID = core.Int64Ptr(int64(26))
		model.DataSourceConnectionID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["connection_id"] = int(26)
	model["entity_id"] = int(26)
	model["connector_group_id"] = int(26)
	model["data_source_connection_id"] = "testString"

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationMapToConnectionConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationMapToKeyValuePair(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.KeyValuePair) {
		model := new(backuprecoveryv1.KeyValuePair)
		model.Key = core.StringPtr("testString")
		model.Value = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["key"] = "testString"
	model["value"] = "testString"

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationMapToKeyValuePair(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasSourceRegistrationMapToPhysicalSourceRegistrationParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.PhysicalSourceRegistrationParams) {
		model := new(backuprecoveryv1.PhysicalSourceRegistrationParams)
		model.Endpoint = core.StringPtr("testString")
		model.ForceRegister = core.BoolPtr(true)
		model.HostType = core.StringPtr("kLinux")
		model.PhysicalType = core.StringPtr("kGroup")
		model.Applications = []string{"kSQL"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["endpoint"] = "testString"
	model["force_register"] = true
	model["host_type"] = "kLinux"
	model["physical_type"] = "kGroup"
	model["applications"] = []interface{}{"kSQL"}

	result, err := backuprecovery.ResourceIbmBaasSourceRegistrationMapToPhysicalSourceRegistrationParams(model)
	assert.Nil(t, err)
	checkResult(result)
}
