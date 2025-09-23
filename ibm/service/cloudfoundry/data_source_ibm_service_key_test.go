// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudfoundry_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMServiceKeyDataSource_basic(t *testing.T) {
	t.Skip()
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	serviceKey := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMServiceKeyDataSourceConfig(serviceName, serviceKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_service_key.testacc_ds_service_key", "name", serviceKey),
					resource.TestCheckResourceAttr("data.ibm_service_key.testacc_ds_service_key", "credentials.%", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMServiceKeyDataSourceConfig(serviceName, serviceKey string) string {
	return fmt.Sprintf(`
	data "ibm_space" "spacedata" {
		org   = "%s"
		space = "%s"
	  }
	  
	  resource "ibm_service_instance" "service" {
		name       = "%s"
		space_guid = data.ibm_space.spacedata.id
		service    = "speech_to_text"
		plan       = "lite"
		tags       = ["cluster-service", "cluster-bind"]
	  }
	  
	  resource "ibm_service_key" "servicekey" {
		name                  = "%s"
		service_instance_guid = ibm_service_instance.service.id
	  }
	  
	  data "ibm_service_key" "testacc_ds_service_key" {
		name                  = ibm_service_key.servicekey.name
		service_instance_name = ibm_service_instance.service.name
		space_guid            = data.ibm_space.spacedata.id
	  }
`, acc.CfOrganization, acc.CfSpace, serviceName, serviceKey)

}
