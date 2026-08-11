// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package restapi_test

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const resourceGroupsURL = "https://resource-controller.cloud.ibm.com/v2/resource_groups"

// TestAccIBMRestApiRequestBasic creates, updates, reads back and destroys a
// resource group through the generic REST resource. Resource groups are used
// because they are cheap, account scoped and available to every test account.
func TestAccIBMRestApiRequestBasic(t *testing.T) {
	name := fmt.Sprintf("tf-restapi-%d", acctest.RandIntRange(10, 10000))
	nameUpdate := fmt.Sprintf("tf-restapi-%d", acctest.RandIntRange(10, 10000))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMRestApiRequestDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRestApiRequestConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRestApiRequestExists("ibm_restapi_request.restapi_request_instance"),
					resource.TestCheckResourceAttr("ibm_restapi_request.restapi_request_instance", "url", resourceGroupsURL),
					resource.TestCheckResourceAttr("ibm_restapi_request.restapi_request_instance", "create_method", "POST"),
					resource.TestCheckResourceAttr("ibm_restapi_request.restapi_request_instance", "response_status_code", "200"),
					resource.TestCheckResourceAttrSet("ibm_restapi_request.restapi_request_instance", "object_id"),
					resource.TestCheckResourceAttrSet("ibm_restapi_request.restapi_request_instance", "object_url"),
					resource.TestCheckResourceAttrSet("ibm_restapi_request.restapi_request_instance", "create_response_body"),
					resource.TestCheckResourceAttrSet("ibm_restapi_request.restapi_request_instance", "response_body"),
				),
			},
			{
				// Only the body changes, so the object must be updated in
				// place rather than replaced.
				Config: testAccCheckIBMRestApiRequestConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMRestApiRequestExists("ibm_restapi_request.restapi_request_instance"),
					resource.TestMatchResourceAttr("ibm_restapi_request.restapi_request_instance", "response_body", regexpForName(nameUpdate)),
				),
			},
		},
	})
}

// TestAccIBMRestApiRequestImport covers the two part import identifier, which
// is required whenever id_attribute is set and the create URL is therefore a
// collection URL.
func TestAccIBMRestApiRequestImport(t *testing.T) {
	name := fmt.Sprintf("tf-restapi-%d", acctest.RandIntRange(10, 10000))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMRestApiRequestDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRestApiRequestConfigBasic(name),
			},
			{
				ResourceName:      "ibm_restapi_request.restapi_request_instance",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccIBMRestApiRequestImportID("ibm_restapi_request.restapi_request_instance"),
				// These are inputs that the API never echoes back, so they
				// cannot be recovered on import.
				ImportStateVerifyIgnore: []string{
					"request_body", "create_method", "read_method", "update_method",
					"destroy_method", "id_attribute", "use_iam_auth", "read_on_create",
					"ignore_read_errors", "timeout_seconds", "create_response_body",
				},
			},
		},
	})
}

// TestAccIBMRestApiRequestAcceptStatusCodes checks that a non 2xx status code
// listed in accept_status_codes does not fail the run. The request targets a
// resource group name that already exists, which the API rejects with 409.
func TestAccIBMRestApiRequestAcceptStatusCodes(t *testing.T) {
	name := fmt.Sprintf("tf-restapi-%d", acctest.RandIntRange(10, 10000))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRestApiRequestConfigAcceptStatus(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_restapi_request.restapi_request_conflict", "response_status_code", "409"),
				),
			},
		},
	})
}

// regexpForName matches a resource group name inside a raw JSON response body.
func regexpForName(name string) *regexp.Regexp {
	return regexp.MustCompile(regexp.QuoteMeta(name))
}

func testAccCheckIBMRestApiRequestConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_restapi_request" "restapi_request_instance" {
			url          = "%s"
			id_attribute = "id"
			request_body = jsonencode({
				name       = "%s"
				account_id = "%s"
			})
			update_request_body = jsonencode({
				name = "%s"
			})
			update_method = "PATCH"
		}
	`, resourceGroupsURL, name, acc.IAMAccountId, name)
}

func testAccCheckIBMRestApiRequestConfigAcceptStatus(name string) string {
	return fmt.Sprintf(`
		resource "ibm_restapi_request" "restapi_request_first" {
			url          = "%[1]s"
			id_attribute = "id"
			request_body = jsonencode({
				name       = "%[2]s"
				account_id = "%[3]s"
			})
		}

		resource "ibm_restapi_request" "restapi_request_conflict" {
			depends_on          = [ibm_restapi_request.restapi_request_first]
			url                 = "%[1]s"
			read_on_create      = false
			ignore_read_errors  = true
			accept_status_codes = [409]
			request_body = jsonencode({
				name       = "%[2]s"
				account_id = "%[3]s"
			})
		}
	`, resourceGroupsURL, name, acc.IAMAccountId)
}

func testAccIBMRestApiRequestImportID(n string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return "", fmt.Errorf("not found: %s", n)
		}
		return fmt.Sprintf("%s,%s", rs.Primary.Attributes["url"], rs.Primary.Attributes["object_url"]), nil
	}
}

// getRestApiObject issues an authenticated GET against the object URL held in
// state and returns the status code.
func getRestApiObject(objectURL string) (int, error) {
	authenticator, err := acc.TestAccProvider.Meta().(conns.ClientSession).Authenticator()
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest(http.MethodGet, objectURL, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Accept", "application/json")
	if err := authenticator.Authenticate(req); err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

func testAccCheckIBMRestApiRequestExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}
		objectURL := rs.Primary.Attributes["object_url"]
		if objectURL == "" {
			return fmt.Errorf("no object_url is set for %s", n)
		}
		statusCode, err := getRestApiObject(objectURL)
		if err != nil {
			return err
		}
		if statusCode != http.StatusOK {
			return fmt.Errorf("expected %s to still exist, GET returned %d", objectURL, statusCode)
		}
		return nil
	}
}

func testAccCheckIBMRestApiRequestDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_restapi_request" {
			continue
		}
		objectURL := rs.Primary.Attributes["object_url"]
		if objectURL == "" {
			continue
		}
		statusCode, err := getRestApiObject(objectURL)
		if err != nil {
			return err
		}
		if statusCode != http.StatusNotFound && statusCode != http.StatusGone {
			return fmt.Errorf("expected %s to be destroyed, GET returned %d", objectURL, statusCode)
		}
	}
	return nil
}
