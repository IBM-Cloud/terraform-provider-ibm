// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.0-a902401e-20260427-192904
 */

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIsDynamicRouteServerRouteReportJobDataSourceBasic(t *testing.T) {
	dynamicRouteServerRouteReportJobDynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerRouteReportJobDataSourceConfigBasic(dynamicRouteServerRouteReportJobDynamicRouteServerID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "dynamic_route_server_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "is_dynamic_route_server_route_report_job_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "format"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "status_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "storage_bucket.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "storage_href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "storage_object.#"),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerRouteReportJobDataSourceAllArgs(t *testing.T) {
	dynamicRouteServerRouteReportJobDynamicRouteServerID := fmt.Sprintf("tf_dynamic_route_server_id_%d", acctest.RandIntRange(10, 100))
	dynamicRouteServerRouteReportJobFormat := "json"
	dynamicRouteServerRouteReportJobName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerRouteReportJobDataSourceConfig(dynamicRouteServerRouteReportJobDynamicRouteServerID, dynamicRouteServerRouteReportJobFormat, dynamicRouteServerRouteReportJobName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "dynamic_route_server_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "is_dynamic_route_server_route_report_job_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "completed_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "format"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "json_schema.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "started_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "status_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "status_reasons.0.code"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "status_reasons.0.message"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "status_reasons.0.more_info"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "storage_bucket.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "storage_href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance", "storage_object.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobDataSourceConfigBasic(dynamicRouteServerRouteReportJobDynamicRouteServerID string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_route_report_job" "is_dynamic_route_server_route_report_job_instance" {
			dynamic_route_server_id = "%s"
			storage_bucket {
				crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue"
				name = "bucket-27200-lwx4cfvcue"
			}
		}

		data "ibm_is_dynamic_route_server_route_report_job" "is_dynamic_route_server_route_report_job_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance.dynamic_route_server_id
			is_dynamic_route_server_route_report_job_id = ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance.is_dynamic_route_server_route_report_job_id
		}
	`, dynamicRouteServerRouteReportJobDynamicRouteServerID)
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobDataSourceConfig(dynamicRouteServerRouteReportJobDynamicRouteServerID string, dynamicRouteServerRouteReportJobFormat string, dynamicRouteServerRouteReportJobName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_route_report_job" "is_dynamic_route_server_route_report_job_instance" {
			dynamic_route_server_id = "%s"
			format = "%s"
			name = "%s"
			storage_bucket {
				crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue"
				name = "bucket-27200-lwx4cfvcue"
			}
		}

		data "ibm_is_dynamic_route_server_route_report_job" "is_dynamic_route_server_route_report_job_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance.dynamic_route_server_id
			is_dynamic_route_server_route_report_job_id = ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance.is_dynamic_route_server_route_report_job_id
		}
	`, dynamicRouteServerRouteReportJobDynamicRouteServerID, dynamicRouteServerRouteReportJobFormat, dynamicRouteServerRouteReportJobName)
}

func TestDataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportSchemaToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["href"] = "https://github.ibm.com/cloud-content/downloads/blob/publish/drs/schema/v1/drs-route-report.json"
		model["type"] = "json"
		model["version"] = "v1"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerRouteReportSchema)
	model.Href = core.StringPtr("https://github.ibm.com/cloud-content/downloads/blob/publish/drs/schema/v1/drs-route-report.json")
	model.Type = core.StringPtr("json")
	model.Version = core.StringPtr("v1")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportSchemaToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportJobStatusReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "cannot_access_storage_bucket"
		model["message"] = "A failure occurred"
		model["more_info"] = "https://cloud.ibm.com/docs/vpc?topic=vpc-object-storage-prereq"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.DynamicRouteServerRouteReportJobStatusReason)
	model.Code = core.StringPtr("cannot_access_storage_bucket")
	model.Message = core.StringPtr("A failure occurred")
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/docs/vpc?topic=vpc-object-storage-prereq")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportJobStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerCloudObjectStorageBucketReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		deletedModel := make(map[string]interface{})
		deletedModel["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		model := make(map[string]interface{})
		model["crn"] = "crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue"
		model["deleted"] = []map[string]interface{}{deletedModel}
		model["name"] = "bucket-27200-lwx4cfvcue"
		model["resource_type"] = "cos_bucket"

		assert.Equal(t, result, model)
	}

	deletedModel := new(vpcv1.Deleted)
	deletedModel.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	model := new(vpcv1.DynamicRouteServerCloudObjectStorageBucketReference)
	model.CRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue")
	model.Deleted = deletedModel
	model.Name = core.StringPtr("bucket-27200-lwx4cfvcue")
	model.ResourceType = core.StringPtr("cos_bucket")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerCloudObjectStorageBucketReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerRouteReportJobDeletedToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["more_info"] = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.Deleted)
	model.MoreInfo = core.StringPtr("https://cloud.ibm.com/apidocs/vpc#deleted-resources")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobDeletedToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIsDynamicRouteServerRouteReportJobCloudObjectStorageObjectReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "my-object"

		assert.Equal(t, result, model)
	}

	model := new(vpcv1.CloudObjectStorageObjectReference)
	model.Name = core.StringPtr("my-object")

	result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobCloudObjectStorageObjectReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
