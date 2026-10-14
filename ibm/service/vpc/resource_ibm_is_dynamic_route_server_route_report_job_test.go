// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/stretchr/testify/assert"
)

const isDynamicRouteServerRouteReportJobResource = "ibm_is_dynamic_route_server_route_report_job.is_dynamic_route_server_route_report_job_instance"

func TestAccIBMIsDynamicRouteServerRouteReportJobBasic(t *testing.T) {
	var conf vpcv1.DynamicRouteServerRouteReportJob
	suffix := acctest.RandIntRange(10, 100)
	vpcname := fmt.Sprintf("tf-vpc-drs-rrj-%d", suffix)
	subnetname := fmt.Sprintf("tf-subnet-drs-rrj-%d", suffix)
	vniname := fmt.Sprintf("tf-vni-drs-rrj-%d", suffix)
	drsname := fmt.Sprintf("tf-drs-rrj-%d", suffix)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerRouteReportJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerRouteReportJobConfigBasic(vpcname, subnetname, vniname, drsname),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerRouteReportJobExists(isDynamicRouteServerRouteReportJobResource, conf),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobResource, "dynamic_route_server_id"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobResource, "dynamic_route_server_route_report_job_id"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobResource, "name"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobResource, "created_at"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobResource, "href"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobResource, "format", "json"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobResource, "schema_version", "v1"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobResource, "resource_type", "dynamic_route_server_route_report_job"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobResource, "status", "succeeded"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobResource, "status_reasons.#"),
				),
			},
		},
	})
}

func TestAccIBMIsDynamicRouteServerRouteReportJobAllArgs(t *testing.T) {
	var conf vpcv1.DynamicRouteServerRouteReportJob
	suffix := acctest.RandIntRange(10, 100)
	vpcname := fmt.Sprintf("tf-vpc-drs-rrj-%d", suffix)
	subnetname := fmt.Sprintf("tf-subnet-drs-rrj-%d", suffix)
	vniname := fmt.Sprintf("tf-vni-drs-rrj-%d", suffix)
	drsname := fmt.Sprintf("tf-drs-rrj-%d", suffix)
	name := fmt.Sprintf("tf-rrj-%d", suffix)
	nameUpdate := fmt.Sprintf("tf-rrj-upd-%d", suffix)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsDynamicRouteServerRouteReportJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsDynamicRouteServerRouteReportJobConfig(vpcname, subnetname, vniname, drsname, name, "v1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerRouteReportJobExists(isDynamicRouteServerRouteReportJobResource, conf),
					resource.TestCheckResourceAttrPair(isDynamicRouteServerRouteReportJobResource, "dynamic_route_server_id", "ibm_is_dynamic_route_server.is_dynamic_route_server_instance", "id"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobResource, "name", name),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobResource, "schema_version", "v1"),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobResource, "status", "succeeded"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobResource, "completed_at"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobResource, "started_at"),
					resource.TestCheckResourceAttrSet(isDynamicRouteServerRouteReportJobResource, "content"),
				),
			},
			resource.TestStep{
				// name is the only patchable field
				Config: testAccCheckIBMIsDynamicRouteServerRouteReportJobConfig(vpcname, subnetname, vniname, drsname, nameUpdate, "v1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsDynamicRouteServerRouteReportJobExists(isDynamicRouteServerRouteReportJobResource, conf),
					resource.TestCheckResourceAttr(isDynamicRouteServerRouteReportJobResource, "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      isDynamicRouteServerRouteReportJobResource,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobDependencies(vpcname, subnetname, vniname, drsname string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}

		resource "ibm_is_subnet" "testacc_subnet" {
			name            = "%s"
			vpc             = ibm_is_vpc.testacc_vpc.id
			zone            = "%s"
			ipv4_cidr_block = "%s"
		}

		resource "ibm_is_virtual_network_interface" "testacc_vni" {
			name   = "%s"
			subnet = ibm_is_subnet.testacc_subnet.id
		}

		resource "ibm_is_dynamic_route_server" "is_dynamic_route_server_instance" {
			name      = "%s"
			local_asn = 64512
			members {
				name = "%s-member"
				virtual_network_interfaces {
					id = ibm_is_virtual_network_interface.testacc_vni.id
				}
			}
		}
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, vniname, drsname, drsname)
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobConfigBasic(vpcname, subnetname, vniname, drsname string) string {
	return testAccCheckIBMIsDynamicRouteServerRouteReportJobDependencies(vpcname, subnetname, vniname, drsname) + `
		resource "ibm_is_dynamic_route_server_route_report_job" "is_dynamic_route_server_route_report_job_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server.is_dynamic_route_server_instance.id
		}
	`
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobConfig(vpcname, subnetname, vniname, drsname, name, schemaVersion string) string {
	return testAccCheckIBMIsDynamicRouteServerRouteReportJobDependencies(vpcname, subnetname, vniname, drsname) + fmt.Sprintf(`
		resource "ibm_is_dynamic_route_server_route_report_job" "is_dynamic_route_server_route_report_job_instance" {
			dynamic_route_server_id = ibm_is_dynamic_route_server.is_dynamic_route_server_instance.id
			name                    = "%s"
			schema_version          = "%s"
		}
	`, name, schemaVersion)
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobExists(n string, obj vpcv1.DynamicRouteServerRouteReportJob) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getDynamicRouteServerRouteReportJobOptions := &vpcv1.GetDynamicRouteServerRouteReportJobOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDynamicRouteServerRouteReportJobOptions.SetDynamicRouteServerID(parts[0])
		getDynamicRouteServerRouteReportJobOptions.SetID(parts[1])

		dynamicRouteServerRouteReportJob, _, err := vpcClient.GetDynamicRouteServerRouteReportJob(getDynamicRouteServerRouteReportJobOptions)
		if err != nil {
			return err
		}

		obj = *dynamicRouteServerRouteReportJob
		return nil
	}
}

func testAccCheckIBMIsDynamicRouteServerRouteReportJobDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_dynamic_route_server_route_report_job" {
			continue
		}

		getDynamicRouteServerRouteReportJobOptions := &vpcv1.GetDynamicRouteServerRouteReportJobOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDynamicRouteServerRouteReportJobOptions.SetDynamicRouteServerID(parts[0])
		getDynamicRouteServerRouteReportJobOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetDynamicRouteServerRouteReportJob(getDynamicRouteServerRouteReportJobOptions)

		if err == nil {
			return fmt.Errorf("DynamicRouteServerRouteReportJob still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for DynamicRouteServerRouteReportJob (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportStatusReasonToMap(t *testing.T) {
	t.Run("nil model", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportStatusReasonToMap(nil)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{}, result)
	})
	t.Run("empty model", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportStatusReasonToMap(&vpcv1.DynamicRouteServerRouteReportStatusReason{})
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{}, result)
	})
	t.Run("without more_info", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerRouteReportStatusReason{
			Code:    core.StringPtr("degraded"),
			Message: core.StringPtr("Route information unavailable for one or more members"),
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportStatusReasonToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{
			"code":    "degraded",
			"message": "Route information unavailable for one or more members",
		}, result)
	})
	t.Run("with more_info", func(t *testing.T) {
		model := &vpcv1.DynamicRouteServerRouteReportStatusReason{
			Code:     core.StringPtr("cannot_generate_route_report"),
			Message:  core.StringPtr("The route report could not be generated"),
			MoreInfo: core.StringPtr("https://cloud.ibm.com/docs/vpc"),
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportStatusReasonToMap(model)
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{
			"code":      "cannot_generate_route_report",
			"message":   "The route report could not be generated",
			"more_info": "https://cloud.ibm.com/docs/vpc",
		}, result)
	})
}

func TestResourceIBMIsDynamicRouteServerRouteReportJobContentToJSON(t *testing.T) {
	t.Run("nil content", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobContentToJSON(nil)
		assert.Nil(t, err)
		assert.Equal(t, "", result)
	})
	t.Run("nil typed pointer", func(t *testing.T) {
		var content *vpcv1.DynamicRouteServerRouteReportContent
		result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobContentToJSON(content)
		assert.Nil(t, err)
		assert.Equal(t, "", result)
	})
	t.Run("empty content", func(t *testing.T) {
		result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobContentToJSON(&vpcv1.DynamicRouteServerRouteReportContent{})
		assert.Nil(t, err)
		assert.Equal(t, `{"schema_version":null}`, result)
	})
	t.Run("content with schema version and members", func(t *testing.T) {
		content := &vpcv1.DynamicRouteServerRouteReportContent{
			SchemaVersion: core.StringPtr("v1"),
			Members:       []vpcv1.DynamicRouteServerRouteReportContentMember{},
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobContentToJSON(content)
		assert.Nil(t, err)
		assert.True(t, strings.Contains(result, `"schema_version":"v1"`))
	})
	t.Run("v1 content subtype with vpc", func(t *testing.T) {
		content := &vpcv1.DynamicRouteServerRouteReportContentV1{
			SchemaVersion: core.StringPtr("v1"),
			VPC: &vpcv1.DynamicRouteServerRouteReportContentVPC{
				Routes: []vpcv1.DynamicRouteServerRouteReportContentVPCRouteIntf{},
			},
		}
		result, err := vpc.ResourceIBMIsDynamicRouteServerRouteReportJobContentToJSON(content)
		assert.Nil(t, err)
		assert.True(t, strings.HasPrefix(result, "{"))
	})
}
