package ibm

import (
	"fmt"
	"testing"

	"strings"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
)

func TestAccIBMAppRoute_Basic(t *testing.T) {
	var conf mccpv2.RouteFields
	host := fmt.Sprintf("terraform_%d", acctest.RandInt())
	updateHost := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAppRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAppRoute_basic(host),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppRouteExists("ibm_app_route.route", &conf),
					resource.TestCheckResourceAttr("ibm_app_route.route", "host", host),
					resource.TestCheckResourceAttr("ibm_app_route.route", "path", "/app"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMAppRoute_updatePath(host),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppRouteExists("ibm_app_route.route", &conf),
					resource.TestCheckResourceAttr("ibm_app_route.route", "host", host),
					resource.TestCheckResourceAttr("ibm_app_route.route", "path", "/app1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMAppRoute_updateHost(updateHost),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_route.route", "host", updateHost),
					resource.TestCheckResourceAttr("ibm_app_route.route", "path", ""),
				),
			},
		},
	})
}

func TestAccIBMAppRoute_With_Tags(t *testing.T) {
	var conf mccpv2.RouteFields
	host := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAppRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAppRoute_with_tags(host),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppRouteExists("ibm_app_route.route", &conf),
					resource.TestCheckResourceAttr("ibm_app_route.route", "tags.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMAppRoute_with_updated_tags(host),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppRouteExists("ibm_app_route.route", &conf),
					resource.TestCheckResourceAttr("ibm_app_route.route", "tags.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMAppRouteDestroy(s *terraform.State) error {
	cfClient, err := testAccProvider.Meta().(ClientSession).MccpAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_app_route" {
			continue
		}

		routeGuid := rs.Primary.ID

		// Try to find the key
		_, err := cfClient.Routes().Get(routeGuid)

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for CF route (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMAppRouteExists(n string, obj *mccpv2.RouteFields) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cfClient, err := testAccProvider.Meta().(ClientSession).MccpAPI()
		if err != nil {
			return err
		}
		routeGuid := rs.Primary.ID

		route, err := cfClient.Routes().Get(routeGuid)
		if err != nil {
			return err
		}

		*obj = *route
		return nil
	}
}

func testAccCheckIBMAppRoute_basic(host string) string {
	return fmt.Sprintf(`
	
		data "ibm_space" "spacedata" {
			org    = "%s"
			space  = "%s"
		}
		
		data "ibm_app_domain_shared" "domain" {
			name        = "mybluemix.net"
		}
		
		resource "ibm_app_route" "route" {
			domain_guid       = "${data.ibm_app_domain_shared.domain.id}"
			space_guid        = "${data.ibm_space.spacedata.id}"
			host              = "%s"
			path              = "/app"
		}
	`, cfOrganization, cfSpace, host)
}

func testAccCheckIBMAppRoute_updatePath(host string) string {
	return fmt.Sprintf(`
	
		data "ibm_space" "spacedata" {
			org    = "%s"
			space  = "%s"
		}
		
		data "ibm_app_domain_shared" "domain" {
			name        = "mybluemix.net"
		}
		
		resource "ibm_app_route" "route" {
			domain_guid       = "${data.ibm_app_domain_shared.domain.id}"
			space_guid        = "${data.ibm_space.spacedata.id}"
			host              = "%s"
			path              = "/app1"
		}
	`, cfOrganization, cfSpace, host)
}

func testAccCheckIBMAppRoute_updateHost(updateHost string) string {
	return fmt.Sprintf(`
		
		data "ibm_space" "spacedata" {
			org    = "%s"
			space  = "%s"
		}
		
		data "ibm_app_domain_shared" "domain" {
			name        = "mybluemix.net"
		}
		
		resource "ibm_app_route" "route" {
			domain_guid       = "${data.ibm_app_domain_shared.domain.id}"
			space_guid        = "${data.ibm_space.spacedata.id}"
			host              = "%s"
		}
	`, cfOrganization, cfSpace, updateHost)
}

func testAccCheckIBMAppRoute_with_tags(host string) string {
	return fmt.Sprintf(`
	
		data "ibm_space" "spacedata" {
			org    = "%s"
			space  = "%s"
		}
		
		data "ibm_app_domain_shared" "domain" {
			name        = "mybluemix.net"
		}
		
		resource "ibm_app_route" "route" {
			domain_guid       = "${data.ibm_app_domain_shared.domain.id}"
			space_guid        = "${data.ibm_space.spacedata.id}"
			host              = "%s"
			path              = "/app"
			tags              = ["one"]
		}
	`, cfOrganization, cfSpace, host)
}

func testAccCheckIBMAppRoute_with_updated_tags(host string) string {
	return fmt.Sprintf(`
	
		data "ibm_space" "spacedata" {
			org    = "%s"
			space  = "%s"
		}
		
		data "ibm_app_domain_shared" "domain" {
			name        = "mybluemix.net"
		}
		
		resource "ibm_app_route" "route" {
			domain_guid       = "${data.ibm_app_domain_shared.domain.id}"
			space_guid        = "${data.ibm_space.spacedata.id}"
			host              = "%s"
			path              = "/app"
			tags              = ["one", "two"]
		}
	`, cfOrganization, cfSpace, host)
}
