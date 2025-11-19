// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/project"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/project-go-sdk/projectv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmProjectBasic(t *testing.T) {
	var conf projectv1.Project
	location := "us-south"
	resourceGroup := fmt.Sprintf("Default")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigBasic(location, resourceGroup),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectExists("ibm_project.project_instance", conf),
					resource.TestCheckResourceAttr("ibm_project.project_instance", "location", location),
					resource.TestCheckResourceAttr("ibm_project.project_instance", "resource_group", resourceGroup),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_project.project_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmProjectConfigBasic(location string, resourceGroup string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			location = "%s"
			resource_group = "%s"
			definition {
                name = "acme-microservice"
                description = "acme-microservice description"
                destroy_on_delete = true
                monitoring_enabled = true
                auto_deploy = true
                auto_deploy_mode = "auto_approval"
            }
		}
	`, location, resourceGroup)
}

func testAccCheckIbmProjectExists(n string, obj projectv1.Project) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
		if err != nil {
			return err
		}

		getProjectOptions := &projectv1.GetProjectOptions{}

		getProjectOptions.SetID(rs.Primary.ID)

		project, _, err := projectClient.GetProject(getProjectOptions)
		if err != nil {
			return err
		}

		obj = *project
		return nil
	}
}

func testAccCheckIbmProjectDestroy(s *terraform.State) error {
	projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_project" {
			continue
		}

		getProjectOptions := &projectv1.GetProjectOptions{}

		getProjectOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := projectClient.GetProject(getProjectOptions)

		if err == nil {
			return fmt.Errorf("project still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for project (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmProjectProjectConfigSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		projectConfigVersionDefinitionSummaryModel := make(map[string]interface{})
		projectConfigVersionDefinitionSummaryModel["environment_id"] = "testString"
		projectConfigVersionDefinitionSummaryModel["locator_id"] = "testString"

		projectConfigVersionSummaryModel := make(map[string]interface{})
		projectConfigVersionSummaryModel["definition"] = []map[string]interface{}{projectConfigVersionDefinitionSummaryModel}
		projectConfigVersionSummaryModel["container_state"] = "approved"
		projectConfigVersionSummaryModel["state"] = "approved"
		projectConfigVersionSummaryModel["version"] = int(0)
		projectConfigVersionSummaryModel["href"] = "testString"

		projectConfigSummaryDefinitionModel := make(map[string]interface{})
		projectConfigSummaryDefinitionModel["description"] = "testString"
		projectConfigSummaryDefinitionModel["name"] = "testString"
		projectConfigSummaryDefinitionModel["locator_id"] = "testString"

		projectDefinitionReferenceModel := make(map[string]interface{})
		projectDefinitionReferenceModel["name"] = "testString"

		projectReferenceModel := make(map[string]interface{})
		projectReferenceModel["id"] = "testString"
		projectReferenceModel["href"] = "testString"
		projectReferenceModel["definition"] = []map[string]interface{}{projectDefinitionReferenceModel}
		projectReferenceModel["crn"] = "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"

		model := make(map[string]interface{})
		model["approved_version"] = []map[string]interface{}{projectConfigVersionSummaryModel}
		model["deployed_version"] = []map[string]interface{}{projectConfigVersionSummaryModel}
		model["id"] = "testString"
		model["version"] = int(0)
		model["container_state"] = "approved"
		model["container_state_code"] = "awaiting_input"
		model["state"] = "approved"
		model["state_code"] = "awaiting_input"
		model["created_at"] = "2019-01-01T12:00:00.000Z"
		model["modified_at"] = "2019-01-01T12:00:00.000Z"
		model["href"] = "testString"
		model["definition"] = []map[string]interface{}{projectConfigSummaryDefinitionModel}
		model["project"] = []map[string]interface{}{projectReferenceModel}
		model["deployment_model"] = "project_deployed"

		assert.Equal(t, result, model)
	}

	projectConfigVersionDefinitionSummaryModel := new(projectv1.ProjectConfigVersionDefinitionSummary)
	projectConfigVersionDefinitionSummaryModel.EnvironmentID = core.StringPtr("testString")
	projectConfigVersionDefinitionSummaryModel.LocatorID = core.StringPtr("testString")

	projectConfigVersionSummaryModel := new(projectv1.ProjectConfigVersionSummary)
	projectConfigVersionSummaryModel.Definition = projectConfigVersionDefinitionSummaryModel
	projectConfigVersionSummaryModel.ContainerState = core.StringPtr("approved")
	projectConfigVersionSummaryModel.State = core.StringPtr("approved")
	projectConfigVersionSummaryModel.Version = core.Int64Ptr(int64(0))
	projectConfigVersionSummaryModel.Href = core.StringPtr("testString")

	projectConfigSummaryDefinitionModel := new(projectv1.ProjectConfigSummaryDefinition)
	projectConfigSummaryDefinitionModel.Description = core.StringPtr("testString")
	projectConfigSummaryDefinitionModel.Name = core.StringPtr("testString")
	projectConfigSummaryDefinitionModel.LocatorID = core.StringPtr("testString")

	projectDefinitionReferenceModel := new(projectv1.ProjectDefinitionReference)
	projectDefinitionReferenceModel.Name = core.StringPtr("testString")

	projectReferenceModel := new(projectv1.ProjectReference)
	projectReferenceModel.ID = core.StringPtr("testString")
	projectReferenceModel.Href = core.StringPtr("testString")
	projectReferenceModel.Definition = projectDefinitionReferenceModel
	projectReferenceModel.Crn = core.StringPtr("crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")

	model := new(projectv1.ProjectConfigSummary)
	model.ApprovedVersion = projectConfigVersionSummaryModel
	model.DeployedVersion = projectConfigVersionSummaryModel
	model.ID = core.StringPtr("testString")
	model.Version = core.Int64Ptr(int64(0))
	model.ContainerState = core.StringPtr("approved")
	model.ContainerStateCode = core.StringPtr("awaiting_input")
	model.State = core.StringPtr("approved")
	model.StateCode = core.StringPtr("awaiting_input")
	model.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.ModifiedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.Href = core.StringPtr("testString")
	model.Definition = projectConfigSummaryDefinitionModel
	model.Project = projectReferenceModel
	model.DeploymentModel = core.StringPtr("project_deployed")

	result, err := project.ResourceIbmProjectProjectConfigSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectProjectConfigVersionSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		projectConfigVersionDefinitionSummaryModel := make(map[string]interface{})
		projectConfigVersionDefinitionSummaryModel["environment_id"] = "testString"
		projectConfigVersionDefinitionSummaryModel["locator_id"] = "testString"

		model := make(map[string]interface{})
		model["definition"] = []map[string]interface{}{projectConfigVersionDefinitionSummaryModel}
		model["container_state"] = "approved"
		model["state"] = "approved"
		model["version"] = int(0)
		model["href"] = "testString"

		assert.Equal(t, result, model)
	}

	projectConfigVersionDefinitionSummaryModel := new(projectv1.ProjectConfigVersionDefinitionSummary)
	projectConfigVersionDefinitionSummaryModel.EnvironmentID = core.StringPtr("testString")
	projectConfigVersionDefinitionSummaryModel.LocatorID = core.StringPtr("testString")

	model := new(projectv1.ProjectConfigVersionSummary)
	model.Definition = projectConfigVersionDefinitionSummaryModel
	model.ContainerState = core.StringPtr("approved")
	model.State = core.StringPtr("approved")
	model.Version = core.Int64Ptr(int64(0))
	model.Href = core.StringPtr("testString")

	result, err := project.ResourceIbmProjectProjectConfigVersionSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectProjectConfigVersionDefinitionSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["environment_id"] = "testString"
		model["locator_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectConfigVersionDefinitionSummary)
	model.EnvironmentID = core.StringPtr("testString")
	model.LocatorID = core.StringPtr("testString")

	result, err := project.ResourceIbmProjectProjectConfigVersionDefinitionSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectProjectConfigSummaryDefinitionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["description"] = "testString"
		model["name"] = "testString"
		model["locator_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectConfigSummaryDefinition)
	model.Description = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.LocatorID = core.StringPtr("testString")

	result, err := project.ResourceIbmProjectProjectConfigSummaryDefinitionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectProjectReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		projectDefinitionReferenceModel := make(map[string]interface{})
		projectDefinitionReferenceModel["name"] = "testString"

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["href"] = "testString"
		model["definition"] = []map[string]interface{}{projectDefinitionReferenceModel}
		model["crn"] = "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"

		assert.Equal(t, result, model)
	}

	projectDefinitionReferenceModel := new(projectv1.ProjectDefinitionReference)
	projectDefinitionReferenceModel.Name = core.StringPtr("testString")

	model := new(projectv1.ProjectReference)
	model.ID = core.StringPtr("testString")
	model.Href = core.StringPtr("testString")
	model.Definition = projectDefinitionReferenceModel
	model.Crn = core.StringPtr("crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")

	result, err := project.ResourceIbmProjectProjectReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectProjectDefinitionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectDefinitionReference)
	model.Name = core.StringPtr("testString")

	result, err := project.ResourceIbmProjectProjectDefinitionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectProjectEnvironmentSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		projectDefinitionReferenceModel := make(map[string]interface{})
		projectDefinitionReferenceModel["name"] = "testString"

		projectReferenceModel := make(map[string]interface{})
		projectReferenceModel["id"] = "testString"
		projectReferenceModel["href"] = "testString"
		projectReferenceModel["definition"] = []map[string]interface{}{projectDefinitionReferenceModel}
		projectReferenceModel["crn"] = "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"

		projectEnvironmentSummaryDefinitionModel := make(map[string]interface{})
		projectEnvironmentSummaryDefinitionModel["description"] = "testString"
		projectEnvironmentSummaryDefinitionModel["name"] = "testString"

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["project"] = []map[string]interface{}{projectReferenceModel}
		model["created_at"] = "2019-01-01T12:00:00.000Z"
		model["href"] = "testString"
		model["definition"] = []map[string]interface{}{projectEnvironmentSummaryDefinitionModel}

		assert.Equal(t, result, model)
	}

	projectDefinitionReferenceModel := new(projectv1.ProjectDefinitionReference)
	projectDefinitionReferenceModel.Name = core.StringPtr("testString")

	projectReferenceModel := new(projectv1.ProjectReference)
	projectReferenceModel.ID = core.StringPtr("testString")
	projectReferenceModel.Href = core.StringPtr("testString")
	projectReferenceModel.Definition = projectDefinitionReferenceModel
	projectReferenceModel.Crn = core.StringPtr("crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")

	projectEnvironmentSummaryDefinitionModel := new(projectv1.ProjectEnvironmentSummaryDefinition)
	projectEnvironmentSummaryDefinitionModel.Description = core.StringPtr("testString")
	projectEnvironmentSummaryDefinitionModel.Name = core.StringPtr("testString")

	model := new(projectv1.ProjectEnvironmentSummary)
	model.ID = core.StringPtr("testString")
	model.Project = projectReferenceModel
	model.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.Href = core.StringPtr("testString")
	model.Definition = projectEnvironmentSummaryDefinitionModel

	result, err := project.ResourceIbmProjectProjectEnvironmentSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectProjectEnvironmentSummaryDefinitionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["description"] = "testString"
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectEnvironmentSummaryDefinition)
	model.Description = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")

	result, err := project.ResourceIbmProjectProjectEnvironmentSummaryDefinitionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectProjectDefinitionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		projectDefinitionStoreModel := make(map[string]interface{})
		projectDefinitionStoreModel["type"] = "gh"
		projectDefinitionStoreModel["url"] = "testString"
		projectDefinitionStoreModel["token"] = "testString"
		projectDefinitionStoreModel["config_directory"] = "testString"

		projectTerraformEngineSettingsModel := make(map[string]interface{})
		projectTerraformEngineSettingsModel["id"] = "testString"
		projectTerraformEngineSettingsModel["type"] = "terraform-enterprise"

		model := make(map[string]interface{})
		model["name"] = "testString"
		model["description"] = "testString"
		model["auto_deploy_mode"] = "manual_approval"
		model["monitoring_enabled"] = false
		model["destroy_on_delete"] = true
		model["store"] = []map[string]interface{}{projectDefinitionStoreModel}
		model["terraform_engine"] = []map[string]interface{}{projectTerraformEngineSettingsModel}
		model["auto_deploy"] = false

		assert.Equal(t, result, model)
	}

	projectDefinitionStoreModel := new(projectv1.ProjectDefinitionStore)
	projectDefinitionStoreModel.Type = core.StringPtr("gh")
	projectDefinitionStoreModel.URL = core.StringPtr("testString")
	projectDefinitionStoreModel.Token = core.StringPtr("testString")
	projectDefinitionStoreModel.ConfigDirectory = core.StringPtr("testString")

	projectTerraformEngineSettingsModel := new(projectv1.ProjectTerraformEngineSettings)
	projectTerraformEngineSettingsModel.ID = core.StringPtr("testString")
	projectTerraformEngineSettingsModel.Type = core.StringPtr("terraform-enterprise")

	model := new(projectv1.ProjectDefinition)
	model.Name = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")
	model.AutoDeployMode = core.StringPtr("manual_approval")
	model.MonitoringEnabled = core.BoolPtr(false)
	model.DestroyOnDelete = core.BoolPtr(true)
	model.Store = projectDefinitionStoreModel
	model.TerraformEngine = projectTerraformEngineSettingsModel
	model.AutoDeploy = core.BoolPtr(false)

	result, err := project.ResourceIbmProjectProjectDefinitionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectProjectDefinitionStoreToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["type"] = "gh"
		model["url"] = "testString"
		model["token"] = "testString"
		model["config_directory"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectDefinitionStore)
	model.Type = core.StringPtr("gh")
	model.URL = core.StringPtr("testString")
	model.Token = core.StringPtr("testString")
	model.ConfigDirectory = core.StringPtr("testString")

	result, err := project.ResourceIbmProjectProjectDefinitionStoreToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectProjectTerraformEngineSettingsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["type"] = "terraform-enterprise"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectTerraformEngineSettings)
	model.ID = core.StringPtr("testString")
	model.Type = core.StringPtr("terraform-enterprise")

	result, err := project.ResourceIbmProjectProjectTerraformEngineSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectCumulativeNeedsAttentionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["event"] = "testString"
		model["event_id"] = "testString"
		model["config_id"] = "testString"
		model["config_version"] = int(0)

		assert.Equal(t, result, model)
	}

	model := new(projectv1.CumulativeNeedsAttention)
	model.Event = core.StringPtr("testString")
	model.EventID = core.StringPtr("testString")
	model.ConfigID = core.StringPtr("testString")
	model.ConfigVersion = core.Int64Ptr(int64(0))

	result, err := project.ResourceIbmProjectCumulativeNeedsAttentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectMapToProjectPrototypeDefinition(t *testing.T) {
	checkResult := func(result *projectv1.ProjectPrototypeDefinition) {
		projectDefinitionStoreModel := new(projectv1.ProjectDefinitionStore)
		projectDefinitionStoreModel.Type = core.StringPtr("gh")
		projectDefinitionStoreModel.URL = core.StringPtr("testString")
		projectDefinitionStoreModel.Token = core.StringPtr("testString")
		projectDefinitionStoreModel.ConfigDirectory = core.StringPtr("testString")

		projectTerraformEngineSettingsModel := new(projectv1.ProjectTerraformEngineSettings)
		projectTerraformEngineSettingsModel.ID = core.StringPtr("testString")
		projectTerraformEngineSettingsModel.Type = core.StringPtr("terraform-enterprise")

		model := new(projectv1.ProjectPrototypeDefinition)
		model.Name = core.StringPtr("testString")
		model.Description = core.StringPtr("testString")
		model.AutoDeployMode = core.StringPtr("manual_approval")
		model.MonitoringEnabled = core.BoolPtr(false)
		model.DestroyOnDelete = core.BoolPtr(true)
		model.Store = projectDefinitionStoreModel
		model.TerraformEngine = projectTerraformEngineSettingsModel
		model.AutoDeploy = core.BoolPtr(false)

		assert.Equal(t, result, model)
	}

	projectDefinitionStoreModel := make(map[string]interface{})
	projectDefinitionStoreModel["type"] = "gh"
	projectDefinitionStoreModel["url"] = "testString"
	projectDefinitionStoreModel["token"] = "testString"
	projectDefinitionStoreModel["config_directory"] = "testString"

	projectTerraformEngineSettingsModel := make(map[string]interface{})
	projectTerraformEngineSettingsModel["id"] = "testString"
	projectTerraformEngineSettingsModel["type"] = "terraform-enterprise"

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["description"] = "testString"
	model["auto_deploy_mode"] = "manual_approval"
	model["monitoring_enabled"] = false
	model["destroy_on_delete"] = true
	model["store"] = []interface{}{projectDefinitionStoreModel}
	model["terraform_engine"] = []interface{}{projectTerraformEngineSettingsModel}
	model["auto_deploy"] = false

	result, err := project.ResourceIbmProjectMapToProjectPrototypeDefinition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectMapToProjectDefinitionStore(t *testing.T) {
	checkResult := func(result *projectv1.ProjectDefinitionStore) {
		model := new(projectv1.ProjectDefinitionStore)
		model.Type = core.StringPtr("gh")
		model.URL = core.StringPtr("testString")
		model.Token = core.StringPtr("testString")
		model.ConfigDirectory = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["type"] = "gh"
	model["url"] = "testString"
	model["token"] = "testString"
	model["config_directory"] = "testString"

	result, err := project.ResourceIbmProjectMapToProjectDefinitionStore(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectMapToProjectTerraformEngineSettings(t *testing.T) {
	checkResult := func(result *projectv1.ProjectTerraformEngineSettings) {
		model := new(projectv1.ProjectTerraformEngineSettings)
		model.ID = core.StringPtr("testString")
		model.Type = core.StringPtr("terraform-enterprise")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "testString"
	model["type"] = "terraform-enterprise"

	result, err := project.ResourceIbmProjectMapToProjectTerraformEngineSettings(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectMapToProjectComplianceProfile(t *testing.T) {
	checkResult := func(result projectv1.ProjectComplianceProfileIntf) {
		model := new(projectv1.ProjectComplianceProfile)
		model.ID = core.StringPtr("testString")
		model.InstanceID = core.StringPtr("testString")
		model.InstanceLocation = core.StringPtr("us-south")
		model.AttachmentID = core.StringPtr("someattachmentid")
		model.ProfileName = core.StringPtr("SCCProfilev1.0")
		model.WpPolicyID = core.StringPtr("testString")
		model.WpInstanceID = core.StringPtr("testString")
		model.WpInstanceName = core.StringPtr("testString")
		model.WpInstanceLocation = core.StringPtr("us-south")
		model.WpZoneID = core.StringPtr("testString")
		model.WpZoneName = core.StringPtr("testString")
		model.WpPolicyName = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "testString"
	model["instance_id"] = "testString"
	model["instance_location"] = "us-south"
	model["attachment_id"] = "someattachmentid"
	model["profile_name"] = "SCCProfilev1.0"
	model["wp_policy_id"] = "testString"
	model["wp_instance_id"] = "testString"
	model["wp_instance_name"] = "testString"
	model["wp_instance_location"] = "us-south"
	model["wp_zone_id"] = "testString"
	model["wp_zone_name"] = "testString"
	model["wp_policy_name"] = "testString"

	result, err := project.ResourceIbmProjectMapToProjectComplianceProfile(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectMapToProjectComplianceProfileNullableObject(t *testing.T) {
	checkResult := func(result *projectv1.ProjectComplianceProfileNullableObject) {
		model := new(projectv1.ProjectComplianceProfileNullableObject)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})

	result, err := project.ResourceIbmProjectMapToProjectComplianceProfileNullableObject(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectMapToProjectComplianceProfileV1(t *testing.T) {
	checkResult := func(result *projectv1.ProjectComplianceProfileV1) {
		model := new(projectv1.ProjectComplianceProfileV1)
		model.ID = core.StringPtr("testString")
		model.InstanceID = core.StringPtr("testString")
		model.InstanceLocation = core.StringPtr("us-south")
		model.AttachmentID = core.StringPtr("testString")
		model.ProfileName = core.StringPtr("testString")
		model.WpPolicyID = core.StringPtr("testString")
		model.WpInstanceID = core.StringPtr("testString")
		model.WpInstanceName = core.StringPtr("testString")
		model.WpInstanceLocation = core.StringPtr("us-south")
		model.WpZoneID = core.StringPtr("testString")
		model.WpZoneName = core.StringPtr("testString")
		model.WpPolicyName = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "testString"
	model["instance_id"] = "testString"
	model["instance_location"] = "us-south"
	model["attachment_id"] = "testString"
	model["profile_name"] = "testString"
	model["wp_policy_id"] = "testString"
	model["wp_instance_id"] = "testString"
	model["wp_instance_name"] = "testString"
	model["wp_instance_location"] = "us-south"
	model["wp_zone_id"] = "testString"
	model["wp_zone_name"] = "testString"
	model["wp_policy_name"] = "testString"

	result, err := project.ResourceIbmProjectMapToProjectComplianceProfileV1(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectMapToStackMember(t *testing.T) {
	checkResult := func(result *projectv1.StackMember) {
		model := new(projectv1.StackMember)
		model.Name = core.StringPtr("testString")
		model.ConfigID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["config_id"] = "testString"

	result, err := project.ResourceIbmProjectMapToStackMember(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectMapToProjectConfigUses(t *testing.T) {
	checkResult := func(result *projectv1.ProjectConfigUses) {
		model := new(projectv1.ProjectConfigUses)
		model.ConfigID = core.StringPtr("testString")
		model.ProjectID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["config_id"] = "testString"
	model["project_id"] = "testString"

	result, err := project.ResourceIbmProjectMapToProjectConfigUses(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectMapToProjectConfigAuth(t *testing.T) {
	checkResult := func(result *projectv1.ProjectConfigAuth) {
		model := new(projectv1.ProjectConfigAuth)
		model.TrustedProfileID = core.StringPtr("Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12")
		model.Method = core.StringPtr("trusted_profile")
		model.ApiKey = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["trusted_profile_id"] = "Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12"
	model["method"] = "trusted_profile"
	model["api_key"] = "testString"

	result, err := project.ResourceIbmProjectMapToProjectConfigAuth(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectMapToProjectConfigDefinitionPrototypeResourceConfigDefinitionPropertiesPrototype(t *testing.T) {
	checkResult := func(result *projectv1.ProjectConfigDefinitionPrototypeResourceConfigDefinitionPropertiesPrototype) {
		projectConfigAuthModel := new(projectv1.ProjectConfigAuth)
		projectConfigAuthModel.TrustedProfileID = core.StringPtr("Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12")
		projectConfigAuthModel.Method = core.StringPtr("trusted_profile")
		projectConfigAuthModel.ApiKey = core.StringPtr("testString")

		model := new(projectv1.ProjectConfigDefinitionPrototypeResourceConfigDefinitionPropertiesPrototype)
		model.ResourceCrns = []string{"crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"}
		model.Description = core.StringPtr("testString")
		model.Name = core.StringPtr("testString")
		model.Authorizations = projectConfigAuthModel
		model.Inputs = map[string]interface{}{"anyKey": "anyValue"}
		model.Settings = map[string]interface{}{"anyKey": "anyValue"}
		model.EnvironmentID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	projectConfigAuthModel := make(map[string]interface{})
	projectConfigAuthModel["trusted_profile_id"] = "Profile-9ac10c5c-195c-41ef-b465-68a6b6dg5f12"
	projectConfigAuthModel["method"] = "trusted_profile"
	projectConfigAuthModel["api_key"] = "testString"

	model := make(map[string]interface{})
	model["resource_crns"] = []interface{}{"crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"}
	model["description"] = "testString"
	model["name"] = "testString"
	model["authorizations"] = []interface{}{projectConfigAuthModel}
	model["inputs"] = map[string]interface{}{"anyKey": "anyValue"}
	model["settings"] = map[string]interface{}{"anyKey": "anyValue"}
	model["environment_id"] = "testString"

	result, err := project.ResourceIbmProjectMapToProjectConfigDefinitionPrototypeResourceConfigDefinitionPropertiesPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectMapToSchematicsWorkspace(t *testing.T) {
	checkResult := func(result *projectv1.SchematicsWorkspace) {
		model := new(projectv1.SchematicsWorkspace)
		model.WorkspaceCrn = core.StringPtr("crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["workspace_crn"] = "crn:v1:staging:public:project:us-south:a/4e1c48fcf8ac33c0a2441e4139f189ae:bf40ad13-b107-446a-8286-c6d576183bb1::"

	result, err := project.ResourceIbmProjectMapToSchematicsWorkspace(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmProjectMapToProjectDefinitionPatch(t *testing.T) {
	checkResult := func(result *projectv1.ProjectDefinitionPatch) {
		projectDefinitionStoreModel := new(projectv1.ProjectDefinitionStore)
		projectDefinitionStoreModel.Type = core.StringPtr("gh")
		projectDefinitionStoreModel.URL = core.StringPtr("testString")
		projectDefinitionStoreModel.Token = core.StringPtr("testString")
		projectDefinitionStoreModel.ConfigDirectory = core.StringPtr("testString")

		projectTerraformEngineSettingsModel := new(projectv1.ProjectTerraformEngineSettings)
		projectTerraformEngineSettingsModel.ID = core.StringPtr("testString")
		projectTerraformEngineSettingsModel.Type = core.StringPtr("terraform-enterprise")

		model := new(projectv1.ProjectDefinitionPatch)
		model.Name = core.StringPtr("testString")
		model.Description = core.StringPtr("testString")
		model.AutoDeployMode = core.StringPtr("auto_approval")
		model.MonitoringEnabled = core.BoolPtr(true)
		model.DestroyOnDelete = core.BoolPtr(true)
		model.Store = projectDefinitionStoreModel
		model.TerraformEngine = projectTerraformEngineSettingsModel
		model.AutoDeploy = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	projectDefinitionStoreModel := make(map[string]interface{})
	projectDefinitionStoreModel["type"] = "gh"
	projectDefinitionStoreModel["url"] = "testString"
	projectDefinitionStoreModel["token"] = "testString"
	projectDefinitionStoreModel["config_directory"] = "testString"

	projectTerraformEngineSettingsModel := make(map[string]interface{})
	projectTerraformEngineSettingsModel["id"] = "testString"
	projectTerraformEngineSettingsModel["type"] = "terraform-enterprise"

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["description"] = "testString"
	model["auto_deploy_mode"] = "auto_approval"
	model["monitoring_enabled"] = true
	model["destroy_on_delete"] = true
	model["store"] = []interface{}{projectDefinitionStoreModel}
	model["terraform_engine"] = []interface{}{projectTerraformEngineSettingsModel}
	model["auto_deploy"] = true

	result, err := project.ResourceIbmProjectMapToProjectDefinitionPatch(model)
	assert.Nil(t, err)
	checkResult(result)
}
