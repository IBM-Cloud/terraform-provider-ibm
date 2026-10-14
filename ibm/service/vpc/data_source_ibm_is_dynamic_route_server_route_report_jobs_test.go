// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

const isDynamicRouteServerRouteReportJobsDataSource = "data.ibm_is_dynamic_route_server_route_report_jobs.is_dynamic_route_server_route_report_jobs_instance"

func TestAccIBMIsDynamicRouteServerRouteReportJobsDataSourceBasic(t *testing.T) {
	suffix := acctest.RandIntRange(10, 100)
	vpcname := fmt.Sprintf("tf-vpc-drs-rrj-%d", suffix)
	subnetname := fmt.Sprintf("tf-subnet-drs-rrj-%d", suffix)
	vniname := fmt.Sprintf("tf-vni-drs-rrj-%d", suffix)
	drsname := fmt.Sprintf("tf-drs-rrj-%d", suffix)
	name := fmt.Sprintf("tf-rrj-%d", suffix)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerRouteReportJobsDataSourceConfigBasic(vpcname, subnetname, vniname, drsname, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobsDataSource, "id"),
					resource.TestCheckResourceAttrPair(isDynamicRouteServerRouteReportJobsDataSource, "dynamic_route_server_id", isDynamicRouteServerRouteReportJobResource, "dynamic_route_server_id"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobsDataSource, "sort", "name"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobsDataSource, "route_report_jobs.#", "1"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobsDataSource, "route_report_jobs.0.name", name),
					resource.TestCheckResourceAttrPair(isDynamicRouteServerRouteReportJobsDataSource, "route_report_jobs.0.id", isDynamicRouteServerRouteReportJobResource, "dynamic_route_server_route_report_job_id"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobsDataSource, "route_report_jobs.0.created_at"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobsDataSource, "route_report_jobs.0.format", "json"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobsDataSource, "route_report_jobs.0.href"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobsDataSource, "route_report_jobs.0.resource_type", "dynamic_route_server_route_report_job"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobsDataSource, "route_report_jobs.0.schema_version", "v1"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobsDataSource, "route_report_jobs.0.status", "succeeded"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobsDataSource, "route_report_jobs.0.status_reasons.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobsDataSourceConfigBasic(vpcname, subnetname, vniname, drsname, name string) string {
	return testAccCheckIBMIsDynamicRouteServerRouteReportJobConfig(vpcname, subnetname, vniname, drsname, name, "v1") + `
		data "ibm_is_dynamic_route_server_route_report_jobs" "is_dynamic_route_server_route_report_jobs_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance.dynamic_route_server_id
			sort                    = "name"
		}
	`
}

func TestDataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobToMap(t *testing.T) {
	createdAt := strfmt.DateTime(time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC))
	startedAt := strfmt.DateTime(time.Date(2026, 1, 2, 3, 5, 0, 0, time.UTC))
	completedAt := strfmt.DateTime(time.Date(2026, 1, 2, 3, 6, 0, 0, time.UTC))

	t.Run("nil model", func(t *testing.T) {
		result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobToMap(nil)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{}, result)
	})
	t.Run("empty model", func(t *testing.T) {
		result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobToMap(&vpcv1.DynamicRouteServerRouteReportJob{})
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"status_reasons": []map[string]interface{}{}}, result)
	})
	t.Run("queued job without started_at, completed_at and with nil status_reasons", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerRouteReportJob{
			CreatedAt:     &createdAt,
			Format:        core.StringPtr("json"),
			Href:          core.StringPtr("https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-aaaa/route_report_jobs/r006-bbbb"),
			ID:            core.StringPtr("r006-bbbb"),
			Name:          core.StringPtr("my-route-report-job"),
			ResourceType:  core.StringPtr("dynamic_route_server_route_report_job"),
			SchemaVersion: core.StringPtr("v1"),
			Status:        core.StringPtr("queued"),
		}
		result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{
			"created_at":     "2026-01-02T03:04:05.000Z",
			"format":         "json",
			"href":           "https://us-south.iaas.cloud.ibm.com/v1/dynamic_route_servers/r006-aaaa/route_report_jobs/r006-bbbb",
			"id":             "r006-bbbb",
			"name":           "my-route-report-job",
			"resource_type":  "dynamic_route_server_route_report_job",
			"schema_version": "v1",
			"status":         "queued",
			"status_reasons": []map[string]interface{}{},
		}, result)
		_, hasStarted := result["started_at"]
		_, hasCompleted := result["completed_at"]
		assert.False(t, hasStarted)
		assert.False(t, hasCompleted)
	})
	t.Run("empty status_reasons list behaves like nil", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerRouteReportJob{
			ID:            core.StringPtr("r006-bbbb"),
			StatusReasons: []vpcv1.DynamicRouteServerRouteReportStatusReason{},
		}
		result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, []map[string]interface{}{}, result["status_reasons"])
	})
	t.Run("succeeded job with single status reason", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerRouteReportJob{
			CompletedAt: &completedAt,
			CreatedAt:   &createdAt,
			ID:          core.StringPtr("r006-bbbb"),
			StartedAt:   &startedAt,
			Status:      core.StringPtr("succeeded"),
			StatusReasons: []vpcv1.DynamicRouteServerRouteReportStatusReason{
				{Code: core.StringPtr("degraded"), Message: core.StringPtr("degraded message")},
			},
		}
		result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, "2026-01-02T03:05:00.000Z", result["started_at"])
		assert.Equal(t, "2026-01-02T03:06:00.000Z", result["completed_at"])
		assert.Equal(t, []map[string]interface{}{
			{"code": "degraded", "message": "degraded message"},
		}, result["status_reasons"])
	})
	t.Run("failed job with multiple status reasons", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerRouteReportJob{
			ID:     core.StringPtr("r006-bbbb"),
			Status: core.StringPtr("failed"),
			StatusReasons: []vpcv1.DynamicRouteServerRouteReportStatusReason{
				{Code: core.StringPtr("cannot_generate_route_report"), Message: core.StringPtr("m1"), MoreInfo: core.StringPtr("https://cloud.ibm.com/docs/vpc")},
				{Code: core.StringPtr("stopped_by_user"), Message: core.StringPtr("m2")},
			},
		}
		result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, []map[string]interface{}{
			{"code": "cannot_generate_route_report", "message": "m1", "more_info": "https://cloud.ibm.com/docs/vpc"},
			{"code": "stopped_by_user", "message": "m2"},
		}, result["status_reasons"])
	})
}

func TestDataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportStatusReasonToMap(t *testing.T) {
	t.Run("nil model", func(t *testing.T) {
		result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportStatusReasonToMap(nil)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{}, result)
	})
	t.Run("empty model", func(t *testing.T) {
		result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportStatusReasonToMap(&vpcv1.DynamicRouteServerRouteReportStatusReason{})
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{}, result)
	})
	t.Run("without more_info", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerRouteReportStatusReason{
			Code:    core.StringPtr("route_report_unavailable"),
			Message: core.StringPtr("The route report is not available at this time"),
		}
		result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportStatusReasonToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{
			"code":    "route_report_unavailable",
			"message": "The route report is not available at this time",
		}, result)
	})
	t.Run("with more_info", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerRouteReportStatusReason{
			Code:     core.StringPtr("stopped_by_provider"),
			Message:  core.StringPtr("The route report has been stopped by IBM"),
			MoreInfo: core.StringPtr("https://cloud.ibm.com/docs/vpc"),
		}
		result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportStatusReasonToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{
			"code":      "stopped_by_provider",
			"message":   "The route report has been stopped by IBM",
			"more_info": "https://cloud.ibm.com/docs/vpc",
		}, result)
	})
}
