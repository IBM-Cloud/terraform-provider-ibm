// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.0-d6dec9d7-20241008-212902
 */

package db2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/db2"
	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmDb2SaasConnectionInfoDataSourceBasic(t *testing.T) {
	db2DeploymentId := "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A"
	db2XDeploymentId := "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDb2SaasConnectionInfoDataSourceConfigBasic(db2DeploymentId, db2XDeploymentId),
				Check: resource.ComposeTestCheckFunc(
					//resource.TestCheckResourceAttrSet("data.ibm_db2_saas_connection_info.db2_saas_connection_info_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_connection_info.db2_saas_connection_info_instance", "deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_connection_info.db2_saas_connection_info_instance", "x_deployment_id"),
				),
			},
		},
	})
}

func testAccCheckIbmDb2SaasConnectionInfoDataSourceConfigBasic(db2DeploymentId, db2XDeploymentId string) string {
	return fmt.Sprintf(`
		data "ibm_db2_saas_connection_info" "db2_saas_connection_info_instance" {
			deployment_id = "[%1s]"
            x_deployment_id = "[%2s]"
            
		}
	`, db2DeploymentId, db2XDeploymentId)
}

func TestDataSourceIbmDb2SaasConnectionInfoSuccessConnectionInfoPublicToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hostname"] = "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast-private.bt1ibm.dev.db2.ibmappdomain.cloud"
		model["database_name"] = "bluedb"
		model["ssl_port"] = "30450"
		model["ssl"] = true
		model["database_version"] = "11.5.0"

		assert.Equal(t, result, model)
	}

	model := new(db2saasv1.SuccessConnectionInfoPublic)
	model.Hostname = core.StringPtr("84792aeb-2a9c-4dee-bfad-2e529f16945d-useast-private.bt1ibm.dev.db2.ibmappdomain.cloud")
	model.DatabaseName = core.StringPtr("bluedb")
	model.SslPort = core.StringPtr("30450")
	model.Ssl = core.BoolPtr(true)
	model.DatabaseVersion = core.StringPtr("11.5.0")

	result, err := db2.DataSourceIbmDb2SaasConnectionInfoSuccessConnectionInfoPublicToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmDb2SaasConnectionInfoSuccessConnectionInfoPrivateToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hostname"] = "84792aeb-2a9c-4dee-bfad-2e529f16945d-useast.bt1ibm.dev.db2.ibmappdomain.cloud"
		model["database_name"] = "bluedb"
		model["ssl_port"] = "30450"
		model["ssl"] = true
		model["database_version"] = "11.5.0"
		model["private_service_name"] = "us-south-private.db2oc.test.saas.ibm.com:32764"
		model["cloud_service_offering"] = "dashdb-for-transactions"
		model["vpe_service_crn"] = "crn:v1:staging:public:dashdb-for-transactions:us-south:::endpoint:feea41a1-ff88-4541-8865-0698ccb7c5dc-us-south-private.bt1ibm.dev.db2.ibmappdomain.cloud"
		model["db_vpc_endpoint_service"] = "feea41a1-ff88-4541-8865-0698ccb7c5dc-ussouth-private.bt1ibm.dev.db2.ibmappdomain.cloud:32679"

		assert.Equal(t, result, model)
	}

	model := new(db2saasv1.SuccessConnectionInfoPrivate)
	model.Hostname = core.StringPtr("84792aeb-2a9c-4dee-bfad-2e529f16945d-useast.bt1ibm.dev.db2.ibmappdomain.cloud")
	model.DatabaseName = core.StringPtr("bluedb")
	model.SslPort = core.StringPtr("30450")
	model.Ssl = core.BoolPtr(true)
	model.DatabaseVersion = core.StringPtr("11.5.0")
	model.PrivateServiceName = core.StringPtr("us-south-private.db2oc.test.saas.ibm.com:32764")
	model.CloudServiceOffering = core.StringPtr("dashdb-for-transactions")
	model.VpeServiceCrn = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:::endpoint:feea41a1-ff88-4541-8865-0698ccb7c5dc-us-south-private.bt1ibm.dev.db2.ibmappdomain.cloud")
	model.DbVpcEndpointService = core.StringPtr("feea41a1-ff88-4541-8865-0698ccb7c5dc-ussouth-private.bt1ibm.dev.db2.ibmappdomain.cloud:32679")

	result, err := db2.DataSourceIbmDb2SaasConnectionInfoSuccessConnectionInfoPrivateToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
