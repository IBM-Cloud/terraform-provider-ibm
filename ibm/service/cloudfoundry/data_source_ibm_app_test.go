// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudfoundry_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppDataSource_Basic(t *testing.T) {
	var conf mccpv2.AppFields
	appName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	routeHostName := fmt.Sprintf("terraform-route-host-%d", acctest.RandIntRange(10, 100))
	svcName := fmt.Sprintf("tfsvc-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMAppDataSourceBasic(routeHostName, svcName, appName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppExists("ibm_app.app", &conf),
					resource.TestCheckResourceAttr("ibm_app.app", "name", appName),
					resource.TestCheckResourceAttrSet("data.ibm_app.ds", "id"),
					resource.TestCheckResourceAttr("data.ibm_app.ds", "name", appName),
					resource.TestCheckResourceAttr("data.ibm_app.ds", "buildpack", "sdk-for-nodejs"),
					resource.TestCheckResourceAttr("data.ibm_app.ds", "environment_json.%", "2"),
					resource.TestCheckResourceAttr("data.ibm_app.ds", "environment_json.test", "test1"),
					resource.TestCheckResourceAttr("data.ibm_app.ds", "environment_json.mockport", "443"),
					resource.TestCheckResourceAttr("data.ibm_app.ds", "state", "STARTED"),
					resource.TestCheckResourceAttr("data.ibm_app.ds", "package_state", "STAGED"),
					resource.TestCheckResourceAttr("data.ibm_app.ds", "route_guid.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_app.ds", "service_instance_guid.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_app.ds", "memory", "128"),
					resource.TestCheckResourceAttr("data.ibm_app.ds", "instances", "1"),
					resource.TestCheckResourceAttr("data.ibm_app.ds", "disk_quota", "512"),
				),
			},
		},
	})
}

func testAccCheckIBMAppDataSourceBasic(routeHost, serviceInstanceName, appName string) (config string) {
	config = fmt.Sprintf(`
	data "ibm_space" "space" {
		org   = "%s"
		space = "%s"
	  }
	  
	  data "ibm_app_domain_shared" "domain" {
		name = "mybluemix.net"
	  }
	  
	  resource "ibm_app_route" "route" {
		domain_guid = data.ibm_app_domain_shared.domain.id
		space_guid  = data.ibm_space.space.id
		host        = "%s"
	  }
	  
	  resource "ibm_service_instance" "service" {
		name       = "%s"
		space_guid = data.ibm_space.space.id
		service    = "speech_to_text"
		plan       = "lite"
		tags       = ["cluster-service"]
	  }
	  
	  resource "ibm_app" "app" {
		name                  = "%s"
		space_guid            = data.ibm_space.space.id
		app_path              = "../../test-fixtures/app1.zip"
		wait_time_minutes     = 20
		buildpack             = "sdk-for-nodejs"
		instances             = 1
		route_guid            = [ibm_app_route.route.id]
		service_instance_guid = [ibm_service_instance.service.id]
		disk_quota            = 512
		memory                = 128
		instances             = 1
		disk_quota            = 512
	  
		environment_json = {
		  "test"     = "test1"
		  "mockport" = 443
		}
	  }
	  
	  data "ibm_app" "ds" {
		name       = ibm_app.app.name
		space_guid = data.ibm_space.space.id
	  }
	  
`, acc.CfOrganization, acc.CfSpace, routeHost, serviceInstanceName, appName)
	return
}
