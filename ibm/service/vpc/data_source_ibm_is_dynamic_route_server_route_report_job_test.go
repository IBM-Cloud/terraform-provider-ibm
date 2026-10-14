// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

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

const isDynamicRouteServerRouteReportJobDataSourceByID = "data.ibm_is_dynamic_route_server_route_report_job.by_id"
const isDynamicRouteServerRouteReportJobDataSourceByName = "data.ibm_is_dynamic_route_server_route_report_job.by_name"

func TestAccIBMIsDynamicRouteServerRouteReportJobDataSourceBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsDynamicRouteServerRouteReportJobDataSourceConfigBasic(vpcname, subnetname, vniname, drsname, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobDataSourceByID, "id"),
					resource.TestCheckResourceAttrPair(isDynamicRouteServerRouteReportJobDataSourceByID, "dynamic_route_server_id", isDynamicRouteServerRouteReportJobResource, "dynamic_route_server_id"),
					resource.TestCheckResourceAttrPair(isDynamicRouteServerRouteReportJobDataSourceByID, "dynamic_route_server_route_report_job_id", isDynamicRouteServerRouteReportJobResource, "dynamic_route_server_route_report_job_id"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobDataSourceByID, "name", name),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobDataSourceByID, "created_at"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobDataSourceByID, "format", "json"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobDataSourceByID, "href"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobDataSourceByID, "resource_type", "dynamic_route_server_route_report_job"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobDataSourceByID, "schema_version", "v1"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobDataSourceByID, "status", "succeeded"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobDataSourceByID, "status_reasons.#"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobDataSourceByID, "content"),
					// lookup by name resolves to the same job
					resource.TestCheckResourceAttrPair(isDynamicRouteServerRouteReportJobDataSourceByName, "dynamic_route_server_route_report_job_id", isDynamicRouteServerRouteReportJobResource, "dynamic_route_server_route_report_job_id"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobDataSourceByName, "name", name),
					resource.TestCheckResourceAttrPair(isDynamicRouteServerRouteReportJobDataSourceByName, "href", isDynamicRouteServerRouteReportJobDataSourceByID, "href"),
				),
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobDataSourceConfigBasic(vpcname, subnetname, vniname, drsname, name string) string {
	return testAccCheckIBMIsDynamicRouteServerRouteReportJobConfig(vpcname, subnetname, vniname, drsname, name, "v1") + `
		data "ibm_is_dynamic_route_server_route_report_job" "by_id" {
			dynamic_route_server_id                     = ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance.dynamic_route_server_id
			dynamic_route_server_route_report_job_id = ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance.dynamic_route_server_route_report_job_id
		}

		data "ibm_is_dynamic_route_server_route_report_job" "by_name" {
			dynamic_route_server_id = ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance.dynamic_route_server_id
			name                    = ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance.name
		}
	`
}

func TestDataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportStatusReasonToMap(t *testing.T) {
	t.Run("nil model", func(t *testing.T) {
		result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportStatusReasonToMap(nil)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{}, result)
	})
	t.Run("empty model", func(t *testing.T) {
		result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportStatusReasonToMap(&vpcv1.DynamicRouteServerRouteReportStatusReason{})
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{}, result)
	})
	t.Run("without more_info", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerRouteReportStatusReason{
			Code:    core.StringPtr("initializing"),
			Message: core.StringPtr("The route report is initializing"),
		}
		result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportStatusReasonToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{
			"code":    "initializing",
			"message": "The route report is initializing",
		}, result)
	})
	t.Run("with more_info", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerRouteReportStatusReason{
			Code:     core.StringPtr("maintenance"),
			Message:  core.StringPtr("The route report is unavailable due to planned maintenance"),
			MoreInfo: core.StringPtr("https://cloud.ibm.com/docs/vpc"),
		}
		result, err := vpc.DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportStatusReasonToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{
			"code":      "maintenance",
			"message":   "The route report is unavailable due to planned maintenance",
			"more_info": "https://cloud.ibm.com/docs/vpc",
		}, result)
	})
}
