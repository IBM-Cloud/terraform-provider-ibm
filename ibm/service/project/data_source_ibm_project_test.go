// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/project"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/project-go-sdk/projectv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmProjectDataSourceBasic(t *testing.T) {
	projectLocation := "us-south"
	projectResourceGroup := fmt.Sprintf("Default")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectDataSourceConfigBasic(projectLocation, projectResourceGroup),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "cumulative_needs_attention_view.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "configs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "environments.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project_instance", "definition.#"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectDataSourceConfigBasic(projectLocation string, projectResourceGroup string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			location = "%s"
			resource_group = "%s"
			definition {
                name = "acme-microservice"
                description = "acme-microservice description"
                destroy_on_delete = true
                monitoring_enabled = true
                auto_deploy = false
                auto_deploy_mode = "manual_approval"
            }
		}

		data "ibm_project" "project_instance" {
			project_id = ibm_project.project_instance.id
		}
	`, projectLocation, projectResourceGroup)
}

func TestDataSourceIbmProjectCumulativeNeedsAttentionToMap(t *testing.T) {
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

	result, err := project.DataSourceIbmProjectCumulativeNeedsAttentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectProjectConfigSummaryToMap(t *testing.T) {
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

	result, err := project.DataSourceIbmProjectProjectConfigSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectProjectConfigVersionSummaryToMap(t *testing.T) {
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

	result, err := project.DataSourceIbmProjectProjectConfigVersionSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectProjectConfigVersionDefinitionSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["environment_id"] = "testString"
		model["locator_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectConfigVersionDefinitionSummary)
	model.EnvironmentID = core.StringPtr("testString")
	model.LocatorID = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectProjectConfigVersionDefinitionSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectProjectConfigSummaryDefinitionToMap(t *testing.T) {
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

	result, err := project.DataSourceIbmProjectProjectConfigSummaryDefinitionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectProjectReferenceToMap(t *testing.T) {
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

	result, err := project.DataSourceIbmProjectProjectReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectProjectDefinitionReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectDefinitionReference)
	model.Name = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectProjectDefinitionReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectProjectEnvironmentSummaryToMap(t *testing.T) {
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

	result, err := project.DataSourceIbmProjectProjectEnvironmentSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectProjectEnvironmentSummaryDefinitionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["description"] = "testString"
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectEnvironmentSummaryDefinition)
	model.Description = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")

	result, err := project.DataSourceIbmProjectProjectEnvironmentSummaryDefinitionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectProjectDefinitionToMap(t *testing.T) {
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

	result, err := project.DataSourceIbmProjectProjectDefinitionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectProjectDefinitionStoreToMap(t *testing.T) {
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

	result, err := project.DataSourceIbmProjectProjectDefinitionStoreToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmProjectProjectTerraformEngineSettingsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["type"] = "terraform-enterprise"

		assert.Equal(t, result, model)
	}

	model := new(projectv1.ProjectTerraformEngineSettings)
	model.ID = core.StringPtr("testString")
	model.Type = core.StringPtr("terraform-enterprise")

	result, err := project.DataSourceIbmProjectProjectTerraformEngineSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
