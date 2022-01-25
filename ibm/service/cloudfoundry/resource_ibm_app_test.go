// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudfoundry_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMApp_Invalid_Application_Path(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppDestroy,
		Steps: []resource.TestStep{

			{
				Config:      testAccCheckIBMAppInvalidPath(name),
				ExpectError: regexp.MustCompile(`The given app path:  doesn't exist`),
			},
		},
	})
}

func TestAccIBMApp_Basic(t *testing.T) {
	var conf mccpv2.AppFields
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMAppCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppExists("ibm_app.app", &conf),
					resource.TestCheckResourceAttr("ibm_app.app", "name", name),
				),
			},
			{
				Config: testAccCheckIBMAppUpdate(updatedName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app.app", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_app.app", "instances", "2"),
					resource.TestCheckResourceAttr("ibm_app.app", "memory", "128"),
					resource.TestCheckResourceAttr("ibm_app.app", "disk_quota", "512"),
					resource.TestCheckResourceAttr("ibm_app.app", "environment_json.test", "test1"),
				),
			},
		},
	})
}

func TestAccIBMApp_with_routes(t *testing.T) {
	var conf mccpv2.AppFields
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	route1 := fmt.Sprintf("terraform-%d", acctest.RandIntRange(10, 100))
	route2 := fmt.Sprintf("terraform-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMAppBindRoute(name, route1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppExists("ibm_app.app", &conf),
					resource.TestCheckResourceAttr("ibm_app.app", "name", name),
					resource.TestCheckResourceAttr("ibm_app.app", "instances", "1"),
					resource.TestCheckResourceAttr("ibm_app.app", "memory", "128"),
					resource.TestCheckResourceAttr("ibm_app.app", "disk_quota", "512"),
					resource.TestCheckResourceAttr("ibm_app.app", "environment_json.test", "test1"),
					resource.TestCheckResourceAttr("ibm_app.app", "route_guid.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMAppAddMultipleRoute(name, route1, route2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppExists("ibm_app.app", &conf),
					resource.TestCheckResourceAttr("ibm_app.app", "name", name),
					resource.TestCheckResourceAttr("ibm_app.app", "instances", "1"),
					resource.TestCheckResourceAttr("ibm_app.app", "memory", "128"),
					resource.TestCheckResourceAttr("ibm_app.app", "disk_quota", "512"),
					resource.TestCheckResourceAttr("ibm_app.app", "environment_json.test", "test1"),
					resource.TestCheckResourceAttr("ibm_app.app", "route_guid.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMAppUnBindRoute(name, route1, route2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppExists("ibm_app.app", &conf),
					resource.TestCheckResourceAttr("ibm_app.app", "name", name),
					resource.TestCheckResourceAttr("ibm_app.app", "instances", "1"),
					resource.TestCheckResourceAttr("ibm_app.app", "memory", "128"),
					resource.TestCheckResourceAttr("ibm_app.app", "disk_quota", "512"),
					resource.TestCheckResourceAttr("ibm_app.app", "environment_json.test", "test1"),
					resource.TestCheckResourceAttr("ibm_app.app", "route_guid.#", "1"),
				),
			},
		},
	})

}

func TestAccIBMApp_with_service_instances(t *testing.T) {
	var conf mccpv2.AppFields
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	route := fmt.Sprintf("terraform-%d", acctest.RandIntRange(10, 100))
	serviceName1 := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	serviceName2 := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMAppBindService(name, route, serviceName1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppExists("ibm_app.app", &conf),
					resource.TestCheckResourceAttr("ibm_app.app", "name", name),
					resource.TestCheckResourceAttr("ibm_app.app", "instances", "1"),
					resource.TestCheckResourceAttr("ibm_app.app", "memory", "128"),
					resource.TestCheckResourceAttr("ibm_app.app", "disk_quota", "512"),
					resource.TestCheckResourceAttr("ibm_app.app", "environment_json.test", "test1"),
					resource.TestCheckResourceAttr("ibm_app.app", "route_guid.#", "1"),
					resource.TestCheckResourceAttr("ibm_app.app", "service_instance_guid.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMAppAddMultipleService(name, route, serviceName1, serviceName2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppExists("ibm_app.app", &conf),
					resource.TestCheckResourceAttr("ibm_app.app", "name", name),
					resource.TestCheckResourceAttr("ibm_app.app", "instances", "1"),
					resource.TestCheckResourceAttr("ibm_app.app", "memory", "128"),
					resource.TestCheckResourceAttr("ibm_app.app", "disk_quota", "512"),
					resource.TestCheckResourceAttr("ibm_app.app", "environment_json.test", "test1"),
					resource.TestCheckResourceAttr("ibm_app.app", "route_guid.#", "1"),
					resource.TestCheckResourceAttr("ibm_app.app", "service_instance_guid.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMAppUnBindService(name, route, serviceName1, serviceName2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppExists("ibm_app.app", &conf),
					resource.TestCheckResourceAttr("ibm_app.app", "name", name),
					resource.TestCheckResourceAttr("ibm_app.app", "instances", "1"),
					resource.TestCheckResourceAttr("ibm_app.app", "memory", "128"),
					resource.TestCheckResourceAttr("ibm_app.app", "disk_quota", "512"),
					resource.TestCheckResourceAttr("ibm_app.app", "environment_json.test", "test1"),
					resource.TestCheckResourceAttr("ibm_app.app", "environment_json.mockport", "443"),
					resource.TestCheckResourceAttr("ibm_app.app", "environment_json.floatval", "0.67"),
					resource.TestCheckResourceAttr("ibm_app.app", "route_guid.#", "1"),
					resource.TestCheckResourceAttr("ibm_app.app", "service_instance_guid.#", "1"),
				),
			},
		},
	})

}

func TestAccIBMApp_With_Tags(t *testing.T) {
	var conf mccpv2.AppFields
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMAppCreate_With_Tags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppExists("ibm_app.app", &conf),
					resource.TestCheckResourceAttr("ibm_app.app", "name", name),
					resource.TestCheckResourceAttr("ibm_app.app", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMAppCreate_With_Updated_Tags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app.app", "name", name),
					resource.TestCheckResourceAttr("ibm_app.app", "tags.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMApp_With_Health_Check(t *testing.T) {
	var conf mccpv2.AppFields
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMAppWithHealthCheck(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppExists("ibm_app.app", &conf),
					resource.TestCheckResourceAttr("ibm_app.app", "name", name),
					resource.TestCheckResourceAttr("ibm_app.app", "instances", "1"),
					resource.TestCheckResourceAttr("ibm_app.app", "memory", "128"),
					resource.TestCheckResourceAttr("ibm_app.app", "disk_quota", "512"),
					resource.TestCheckResourceAttr("ibm_app.app", "health_check_type", "port"),
					resource.TestCheckResourceAttr("ibm_app.app", "health_check_timeout", "120"),
				),
			},
			{
				Config: testAccCheckIBMAppWithHealthCheckUpdate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppExists("ibm_app.app", &conf),
					resource.TestCheckResourceAttr("ibm_app.app", "name", name),
					resource.TestCheckResourceAttr("ibm_app.app", "instances", "1"),
					resource.TestCheckResourceAttr("ibm_app.app", "memory", "128"),
					resource.TestCheckResourceAttr("ibm_app.app", "disk_quota", "512"),
					resource.TestCheckResourceAttr("ibm_app.app", "health_check_type", "port"),
					resource.TestCheckResourceAttr("ibm_app.app", "health_check_timeout", "180"),
				),
			},
		},
	})
}

func testAccCheckIBMAppDestroy(s *terraform.State) error {
	cfClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_app" {
			continue
		}
		appGUID := rs.Primary.ID

		_, err := cfClient.Apps().Get(appGUID)
		if err == nil {
			return fmt.Errorf("App still exists: %s", rs.Primary.ID)
		}
	}

	return nil

}

func testAccCheckIBMAppExists(n string, obj *mccpv2.AppFields) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cfClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MccpAPI()
		if err != nil {
			return err
		}
		appGUID := rs.Primary.ID

		app, err := cfClient.Apps().Get(appGUID)
		if err != nil {
			return err
		}

		*obj = *app
		return nil
	}
}

func testAccCheckIBMAppInvalidPath(name string) string {
	return fmt.Sprintf(`

	data "ibm_space" "space" {
		org   = "%s"
		space = "%s"
	  }
	  
	  resource "ibm_app" "app" {
		name              = "%s"
		space_guid        = data.ibm_space.space.id
		app_path          = ""
		wait_time_minutes = 90
		buildpack         = "sdk-for-nodejs"
	  }
	  
`, acc.CfOrganization, acc.CfSpace, name)

}

func testAccCheckIBMAppCreate(name string) string {
	return fmt.Sprintf(`

	data "ibm_space" "space" {
		org   = "%s"
		space = "%s"
	  }
	  
	  resource "ibm_app" "app" {
		name              = "%s"
		space_guid        = data.ibm_space.space.id
		app_path          = "../../test-fixtures/app1.zip"
		wait_time_minutes = 90
		buildpack         = "sdk-for-nodejs"
	  }
`, acc.CfOrganization, acc.CfSpace, name)

}

func testAccCheckIBMAppUpdate(name string) string {
	return fmt.Sprintf(`
	data "ibm_space" "space" {
		org   = "%s"
		space = "%s"
	  }
	  
	  resource "ibm_app" "app" {
		name              = "%s"
		space_guid        = data.ibm_space.space.id
		app_path          = "../../test-fixtures/app1.zip"
		wait_time_minutes = 20
		buildpack         = "sdk-for-nodejs"
		disk_quota        = 512
		memory            = 128
		instances         = 2
	  
		environment_json = {
		  "test" = "test1"
		}
	  }
`, acc.CfOrganization, acc.CfSpace, name)

}

func testAccCheckIBMAppBindRoute(name, route1 string) string {
	return fmt.Sprintf(`

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
	  
	  resource "ibm_app" "app" {
		name              = "%s"
		space_guid        = data.ibm_space.space.id
		app_path          = "../../test-fixtures/app1.zip"
		wait_time_minutes = 20
		buildpack         = "sdk-for-nodejs"
		instances         = 1
		route_guid        = [ibm_app_route.route.id]
		disk_quota        = 512
		memory            = 128
	  
		environment_json = {
		  "test" = "test1"
		}
	  }	  
`, acc.CfOrganization, acc.CfSpace, route1, name)

}

func testAccCheckIBMAppAddMultipleRoute(name, route1, route2 string) string {
	return fmt.Sprintf(`

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
	  
	  resource "ibm_app_route" "route1" {
		domain_guid = data.ibm_app_domain_shared.domain.id
		space_guid  = data.ibm_space.space.id
		host        = "%s"
	  }
	  
	  resource "ibm_app" "app" {
		name              = "%s"
		space_guid        = data.ibm_space.space.id
		app_path          = "../../test-fixtures/app1.zip"
		wait_time_minutes = 20
		buildpack         = "sdk-for-nodejs"
		instances         = 1
		route_guid        = [ibm_app_route.route.id, ibm_app_route.route1.id]
		disk_quota        = 512
		memory            = 128
		disk_quota        = 512
	  
		environment_json = {
		  "test" = "test1"
		}
	  }
`, acc.CfOrganization, acc.CfSpace, route1, route2, name)

}

func testAccCheckIBMAppUnBindRoute(name, route1, route2 string) string {
	return fmt.Sprintf(`

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
	  
	  resource "ibm_app_route" "route1" {
		domain_guid = data.ibm_app_domain_shared.domain.id
		space_guid  = data.ibm_space.space.id
		host        = "%s"
	  }
	  
	  resource "ibm_app" "app" {
		name              = "%s"
		space_guid        = data.ibm_space.space.id
		app_path          = "../../test-fixtures/app1.zip"
		wait_time_minutes = 20
		buildpack         = "sdk-for-nodejs"
		instances         = 1
		route_guid        = [ibm_app_route.route.id]
		disk_quota        = 512
		memory            = 128
		instances         = 1
		disk_quota        = 512
	  
		environment_json = {
		  "test" = "test1"
		}
	  }	  
`, acc.CfOrganization, acc.CfSpace, route1, route2, name)

}

func testAccCheckIBMAppBindService(name, route1, serviceName string) string {
	return fmt.Sprintf(`

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
		tags       = ["cluster-service", "cluster-bind"]
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
	  
		environment_json = {
		  "test" = "test1"
		}
	  }
`, acc.CfOrganization, acc.CfSpace, route1, serviceName, name)

}

func testAccCheckIBMAppAddMultipleService(name, route, serviceName1, serviceName2 string) string {
	return fmt.Sprintf(`

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
		tags       = ["cluster-service", "cluster-bind"]
	  }
	  
	  resource "ibm_service_instance" "service1" {
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
		service_instance_guid = [ibm_service_instance.service.id, ibm_service_instance.service1.id]
		disk_quota            = 512
		memory                = 128
		instances             = 1
		disk_quota            = 512
	  
		environment_json = {
		  "test" = "test1"
		}
	  }
`, acc.CfOrganization, acc.CfSpace, route, serviceName1, serviceName2, name)

}

func testAccCheckIBMAppUnBindService(name, route1, serviceName1, serviceName2 string) string {
	return fmt.Sprintf(`

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
		tags       = ["cluster-service", "cluster-bind"]
	  }
	  
	  resource "ibm_service_instance" "service1" {
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
		  "floatval" = 0.67
		}
	 }	  
`, acc.CfOrganization, acc.CfSpace, route1, serviceName1, serviceName2, name)

}

func testAccCheckIBMAppCreate_With_Tags(name string) string {
	return fmt.Sprintf(`

	data "ibm_space" "space" {
		org   = "%s"
		space = "%s"
	  }
	  
	  resource "ibm_app" "app" {
		name              = "%s"
		space_guid        = data.ibm_space.space.id
		app_path          = "../../test-fixtures/app1.zip"
		wait_time_minutes = 90
		buildpack         = "sdk-for-nodejs"
		tags              = ["one", "two"]
	  }
`, acc.CfOrganization, acc.CfSpace, name)

}

func testAccCheckIBMAppCreate_With_Updated_Tags(name string) string {
	return fmt.Sprintf(`

	data "ibm_space" "space" {
		org   = "%s"
		space = "%s"
	  }
	  
	  resource "ibm_app" "app" {
		name              = "%s"
		space_guid        = data.ibm_space.space.id
		app_path          = "../../test-fixtures/app1.zip"
		wait_time_minutes = 90
		buildpack         = "sdk-for-nodejs"
		tags              = ["one", "two", "three"]
	  }
`, acc.CfOrganization, acc.CfSpace, name)

}

func testAccCheckIBMAppWithHealthCheck(name string) string {
	return fmt.Sprintf(`
	data "ibm_space" "space" {
		org   = "%s"
		space = "%s"
	  }
	  
	  resource "ibm_app" "app" {
		name                 = "%s"
		space_guid           = data.ibm_space.space.id
		app_path             = "../../test-fixtures/app1.zip"
		wait_time_minutes    = 90
		health_check_timeout = 120
		instances            = 1
		disk_quota           = 512
		memory               = 128
	  }
`, acc.CfOrganization, acc.CfSpace, name)

}

func testAccCheckIBMAppWithHealthCheckUpdate(name string) string {
	return fmt.Sprintf(`
	data "ibm_space" "space" {
		org   = "%s"
		space = "%s"
	  }
	  
	  resource "ibm_app" "app" {
		name                 = "%s"
		space_guid           = data.ibm_space.space.id
		app_path             = "../../test-fixtures/app1.zip"
		wait_time_minutes    = 90
		health_check_timeout = 180
		instances            = 1
		disk_quota           = 512
		memory               = 128
	  }
`, acc.CfOrganization, acc.CfSpace, name)

}
