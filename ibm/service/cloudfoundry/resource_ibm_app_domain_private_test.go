// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudfoundry_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/Mavrickk3/bluemix-go/api/mccp/mccpv2"
)

func TestAccIBMAppDomainPrivate_Basic(t *testing.T) {
	var conf mccpv2.PrivateDomainFields
	name := fmt.Sprintf("terraform%d.com", acctest.RandIntRange(10, 100))
	updateName := fmt.Sprintf("terraformnew%d.com", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppDomainPrivate_basic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppDomainPrivateExists("ibm_app_domain_private.domain", &conf),
					resource.TestCheckResourceAttr("ibm_app_domain_private.domain", "name", name),
				),
			},
			{
				Config: testAccCheckIBMAppDomainPrivate_updateName(updateName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_domain_private.domain", "name", updateName),
				),
			},
		},
	})
}

func TestAccIBMAppDomainPrivate_With_Tags(t *testing.T) {
	var conf mccpv2.PrivateDomainFields
	name := fmt.Sprintf("terraform%d.com", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppDomainPrivate_with_tags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAppDomainPrivateExists("ibm_app_domain_private.domain", &conf),
					resource.TestCheckResourceAttr("ibm_app_domain_private.domain", "name", name),
					resource.TestCheckResourceAttr("ibm_app_domain_private.domain", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMAppDomainPrivate_with_updated_tags(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_domain_private.domain", "name", name),
					resource.TestCheckResourceAttr("ibm_app_domain_private.domain", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMAppDomainPrivateExists(n string, obj *mccpv2.PrivateDomainFields) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cfClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MccpAPI()
		if err != nil {
			return err
		}
		privateDomainGUID := rs.Primary.ID

		prdomain, err := cfClient.PrivateDomains().Get(privateDomainGUID)
		if err != nil {
			return err
		}

		*obj = *prdomain
		return nil
	}
}

func testAccCheckIBMAppDomainPrivateDestroy(s *terraform.State) error {
	cfClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MccpAPI()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_app_domain_private" {
			continue
		}

		privateDomainGUID := rs.Primary.ID

		// Try to find the private domain
		_, err := cfClient.PrivateDomains().Get(privateDomainGUID)

		if err == nil {
			return fmt.Errorf("CF private domain still exists: %s", rs.Primary.ID)
		} else if !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("[ERROR] Error waiting for CF private domain (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMAppDomainPrivate_basic(name string) string {
	return fmt.Sprintf(`
		
	data "ibm_org" "orgdata" {
		org = "%s"
	  }
	  
	  resource "ibm_app_domain_private" "domain" {
		name     = "%s"
		org_guid = data.ibm_org.orgdata.id
	  }
	`, acc.CfOrganization, name)
}

func testAccCheckIBMAppDomainPrivate_updateName(updateName string) string {
	return fmt.Sprintf(`
		
	data "ibm_org" "orgdata" {
		org = "%s"
	  }
	  
	  resource "ibm_app_domain_private" "domain" {
		name     = "%s"
		org_guid = data.ibm_org.orgdata.id
	  }
	`, acc.CfOrganization, updateName)
}

func testAccCheckIBMAppDomainPrivate_with_tags(name string) string {
	return fmt.Sprintf(`
		
	data "ibm_org" "orgdata" {
		org = "%s"
	  }
	  
	  resource "ibm_app_domain_private" "domain" {
		name     = "%s"
		org_guid = data.ibm_org.orgdata.id
		tags     = ["one", "two"]
	  }
	  
	`, acc.CfOrganization, name)
}

func testAccCheckIBMAppDomainPrivate_with_updated_tags(name string) string {
	return fmt.Sprintf(`
		
	data "ibm_org" "orgdata" {
		org = "%s"
	  }
	  
	  resource "ibm_app_domain_private" "domain" {
		name     = "%s"
		org_guid = data.ibm_org.orgdata.id
		tags     = ["one", "two", "three"]
	  }
	`, acc.CfOrganization, name)
}
