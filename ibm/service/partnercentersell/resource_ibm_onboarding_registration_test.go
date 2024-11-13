// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package partnercentersell_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/partnercentersell"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnercentersellv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmOnboardingRegistrationBasic(t *testing.T) {
	var conf partnercentersellv1.Registration
	accountID := acc.PcsRegistrationAccountId
	companyName := "Test_company"
	accountIDUpdate := acc.PcsRegistrationAccountId
	companyNameUpdate := "Test_company_up"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingRegistrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingRegistrationConfigBasic(accountID, companyName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingRegistrationExists("ibm_onboarding_registration.onboarding_registration_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_registration.onboarding_registration_instance", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_onboarding_registration.onboarding_registration_instance", "company_name", companyName),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingRegistrationConfigBasic(accountIDUpdate, companyNameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_registration.onboarding_registration_instance", "account_id", accountIDUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_registration.onboarding_registration_instance", "company_name", companyNameUpdate),
				),
			},
		},
	})
}

func TestAccIbmOnboardingRegistrationAllArgs(t *testing.T) {
	var conf partnercentersellv1.Registration
	accountID := acc.PcsRegistrationAccountId
	companyName := "Test_company"
	defaultPrivateCatalogID := "772b632e-fab4-4c41-b0b7-0a92fa40cf67"
	providerAccessGroup := "AccessGroupId-b08e7bb5-d480-4c26-b193-d57dd9311608"
	accountIDUpdate := acc.PcsRegistrationAccountId
	companyNameUpdate := "Test_company_up"
	defaultPrivateCatalogIDUpdate := "772b632e-fab4-4c41-b0b7-0a92fa40cf67"
	providerAccessGroupUpdate := "AccessGroupId-b08e7bb5-d480-4c26-b193-d57dd9311608"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingRegistrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingRegistrationConfig(accountID, companyName, defaultPrivateCatalogID, providerAccessGroup),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingRegistrationExists("ibm_onboarding_registration.onboarding_registration_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_registration.onboarding_registration_instance", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_onboarding_registration.onboarding_registration_instance", "company_name", companyName),
					resource.TestCheckResourceAttr("ibm_onboarding_registration.onboarding_registration_instance", "default_private_catalog_id", defaultPrivateCatalogID),
					resource.TestCheckResourceAttr("ibm_onboarding_registration.onboarding_registration_instance", "provider_access_group", providerAccessGroup),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingRegistrationConfig(accountIDUpdate, companyNameUpdate, defaultPrivateCatalogIDUpdate, providerAccessGroupUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_registration.onboarding_registration_instance", "account_id", accountIDUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_registration.onboarding_registration_instance", "company_name", companyNameUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_registration.onboarding_registration_instance", "default_private_catalog_id", defaultPrivateCatalogIDUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_registration.onboarding_registration_instance", "provider_access_group", providerAccessGroupUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_onboarding_registration.onboarding_registration_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmOnboardingRegistrationConfigBasic(accountID string, companyName string) string {
	return fmt.Sprintf(`
		resource "ibm_onboarding_registration" "onboarding_registration_instance" {
			account_id = "%s"
			company_name = "%s"
			primary_contact {
				name = "Petra"
				email = "petra@ibm.com"
			}
		}
	`, accountID, companyName)
}

func testAccCheckIbmOnboardingRegistrationConfig(accountID string, companyName string, defaultPrivateCatalogID string, providerAccessGroup string) string {
	return fmt.Sprintf(`

		resource "ibm_onboarding_registration" "onboarding_registration_instance" {
			account_id = "%s"
			company_name = "%s"
			primary_contact {
				name = "Petra"
				email = "petra@ibm.com"
			}
			default_private_catalog_id = "%s"
			provider_access_group = "%s"
		}
	`, accountID, companyName, defaultPrivateCatalogID, providerAccessGroup)
}

func testAccCheckIbmOnboardingRegistrationExists(n string, obj partnercentersellv1.Registration) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
		if err != nil {
			return err
		}

		getRegistrationOptions := &partnercentersellv1.GetRegistrationOptions{}

		getRegistrationOptions.SetRegistrationID(rs.Primary.ID)

		registration, _, err := partnerCenterSellClient.GetRegistration(getRegistrationOptions)
		if err != nil {
			return err
		}

		obj = *registration
		return nil
	}
}

func testAccCheckIbmOnboardingRegistrationDestroy(s *terraform.State) error {
	partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_onboarding_registration" {
			continue
		}

		getRegistrationOptions := &partnercentersellv1.GetRegistrationOptions{}

		getRegistrationOptions.SetRegistrationID(rs.Primary.ID)

		// Try to find the key
		_, response, err := partnerCenterSellClient.GetRegistration(getRegistrationOptions)

		if err == nil {
			return fmt.Errorf("onboarding_registration still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for onboarding_registration (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmOnboardingRegistrationPrimaryContactToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["email"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.PrimaryContact)
	model.Name = core.StringPtr("testString")
	model.Email = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingRegistrationPrimaryContactToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingRegistrationMapToPrimaryContact(t *testing.T) {
	checkResult := func(result *partnercentersellv1.PrimaryContact) {
		model := new(partnercentersellv1.PrimaryContact)
		model.Name = core.StringPtr("testString")
		model.Email = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["email"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingRegistrationMapToPrimaryContact(model)
	assert.Nil(t, err)
	checkResult(result)
}
