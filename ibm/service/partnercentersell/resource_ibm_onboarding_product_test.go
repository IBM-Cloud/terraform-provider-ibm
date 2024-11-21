// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package partnercentersell_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/partnercentersell"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnercentersellv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmOnboardingProductBasic(t *testing.T) {
	var conf partnercentersellv1.OnboardingProduct
	typeVar := "software"
	typeVarUpdate := "professional_service"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingProductDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingProductConfigBasic(typeVar),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingProductExists("ibm_onboarding_product.onboarding_product_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_product.onboarding_product_instance", "type", typeVar),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingProductConfigBasic(typeVarUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_product.onboarding_product_instance", "type", typeVarUpdate),
				),
			},
		},
	})
}

func TestAccIbmOnboardingProductAllArgs(t *testing.T) {
	var conf partnercentersellv1.OnboardingProduct
	typeVar := "software"
	eccnNumber := fmt.Sprintf("tf_eccn_number_%d", acctest.RandIntRange(10, 100))
	eroClass := fmt.Sprintf("tf_ero_class_%d", acctest.RandIntRange(10, 100))
	taxAssessment := fmt.Sprintf("tf_tax_assessment_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "professional_service"
	eccnNumberUpdate := fmt.Sprintf("tf_eccn_number_%d", acctest.RandIntRange(10, 100))
	eroClassUpdate := fmt.Sprintf("tf_ero_class_%d", acctest.RandIntRange(10, 100))
	taxAssessmentUpdate := fmt.Sprintf("tf_tax_assessment_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingProductDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingProductConfig(typeVar, eccnNumber, eroClass, taxAssessment),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingProductExists("ibm_onboarding_product.onboarding_product_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_product.onboarding_product_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_onboarding_product.onboarding_product_instance", "eccn_number", eccnNumber),
					resource.TestCheckResourceAttr("ibm_onboarding_product.onboarding_product_instance", "ero_class", eroClass),
					resource.TestCheckResourceAttr("ibm_onboarding_product.onboarding_product_instance", "tax_assessment", taxAssessment),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingProductConfig(typeVarUpdate, eccnNumberUpdate, eroClassUpdate, taxAssessmentUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_product.onboarding_product_instance", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_product.onboarding_product_instance", "eccn_number", eccnNumberUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_product.onboarding_product_instance", "ero_class", eroClassUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_product.onboarding_product_instance", "tax_assessment", taxAssessmentUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_onboarding_product.onboarding_product",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmOnboardingProductConfigBasic(typeVar string) string {
	return fmt.Sprintf(`
		resource "ibm_onboarding_product" "onboarding_product_instance" {
			type = "%s"
			primary_contact {
				name = "name"
				email = "email"
			}
		}
	`, typeVar)
}

func testAccCheckIbmOnboardingProductConfig(typeVar string, eccnNumber string, eroClass string, taxAssessment string) string {
	return fmt.Sprintf(`

		resource "ibm_onboarding_product" "onboarding_product_instance" {
			type = "%s"
			primary_contact {
				name = "name"
				email = "email"
			}
			eccn_number = "%s"
			ero_class = "%s"
			unspsc = "FIXME"
			tax_assessment = "%s"
			support {
				escalation_contacts {
					name = "name"
					email = "email"
					role = "role"
				}
			}
		}
	`, typeVar, eccnNumber, eroClass, taxAssessment)
}

func testAccCheckIbmOnboardingProductExists(n string, obj partnercentersellv1.OnboardingProduct) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
		if err != nil {
			return err
		}

		getOnboardingProductOptions := &partnercentersellv1.GetOnboardingProductOptions{}

		getOnboardingProductOptions.SetProductID(rs.Primary.ID)

		onboardingProduct, _, err := partnerCenterSellClient.GetOnboardingProduct(getOnboardingProductOptions)
		if err != nil {
			return err
		}

		obj = *onboardingProduct
		return nil
	}
}

func testAccCheckIbmOnboardingProductDestroy(s *terraform.State) error {
	partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_onboarding_product" {
			continue
		}

		getOnboardingProductOptions := &partnercentersellv1.GetOnboardingProductOptions{}

		getOnboardingProductOptions.SetProductID(rs.Primary.ID)

		// Try to find the key
		_, response, err := partnerCenterSellClient.GetOnboardingProduct(getOnboardingProductOptions)

		if err == nil {
			return fmt.Errorf("onboarding_product still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for onboarding_product (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmOnboardingProductPrimaryContactToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["email"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.PrimaryContact)
	model.Name = core.StringPtr("testString")
	model.Email = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingProductPrimaryContactToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingProductOnboardingProductSupportToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		onboardingProductSupportEscalationContactItemsModel := make(map[string]interface{})
		onboardingProductSupportEscalationContactItemsModel["name"] = "testString"
		onboardingProductSupportEscalationContactItemsModel["email"] = "testString"
		onboardingProductSupportEscalationContactItemsModel["role"] = "testString"

		model := make(map[string]interface{})
		model["escalation_contacts"] = []map[string]interface{}{onboardingProductSupportEscalationContactItemsModel}

		assert.Equal(t, result, model)
	}

	onboardingProductSupportEscalationContactItemsModel := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
	onboardingProductSupportEscalationContactItemsModel.Name = core.StringPtr("testString")
	onboardingProductSupportEscalationContactItemsModel.Email = core.StringPtr("testString")
	onboardingProductSupportEscalationContactItemsModel.Role = core.StringPtr("testString")

	model := new(partnercentersellv1.OnboardingProductSupport)
	model.EscalationContacts = []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel}

	result, err := partnercentersell.ResourceIbmOnboardingProductOnboardingProductSupportToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingProductOnboardingProductSupportEscalationContactItemsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["email"] = "testString"
		model["role"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
	model.Name = core.StringPtr("testString")
	model.Email = core.StringPtr("testString")
	model.Role = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingProductOnboardingProductSupportEscalationContactItemsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingProductMapToPrimaryContact(t *testing.T) {
	checkResult := func(result *partnercentersellv1.PrimaryContact) {
		model := new(partnercentersellv1.PrimaryContact)
		model.Name = core.StringPtr("testString")
		model.Email = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["email"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingProductMapToPrimaryContact(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingProductMapToOnboardingProductSupport(t *testing.T) {
	checkResult := func(result *partnercentersellv1.OnboardingProductSupport) {
		onboardingProductSupportEscalationContactItemsModel := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
		onboardingProductSupportEscalationContactItemsModel.Name = core.StringPtr("testString")
		onboardingProductSupportEscalationContactItemsModel.Email = core.StringPtr("testString")
		onboardingProductSupportEscalationContactItemsModel.Role = core.StringPtr("testString")

		model := new(partnercentersellv1.OnboardingProductSupport)
		model.EscalationContacts = []partnercentersellv1.OnboardingProductSupportEscalationContactItems{*onboardingProductSupportEscalationContactItemsModel}

		assert.Equal(t, result, model)
	}

	onboardingProductSupportEscalationContactItemsModel := make(map[string]interface{})
	onboardingProductSupportEscalationContactItemsModel["name"] = "testString"
	onboardingProductSupportEscalationContactItemsModel["email"] = "testString"
	onboardingProductSupportEscalationContactItemsModel["role"] = "testString"

	model := make(map[string]interface{})
	model["escalation_contacts"] = []interface{}{onboardingProductSupportEscalationContactItemsModel}

	result, err := partnercentersell.ResourceIbmOnboardingProductMapToOnboardingProductSupport(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingProductMapToOnboardingProductSupportEscalationContactItems(t *testing.T) {
	checkResult := func(result *partnercentersellv1.OnboardingProductSupportEscalationContactItems) {
		model := new(partnercentersellv1.OnboardingProductSupportEscalationContactItems)
		model.Name = core.StringPtr("testString")
		model.Email = core.StringPtr("testString")
		model.Role = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["email"] = "testString"
	model["role"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingProductMapToOnboardingProductSupportEscalationContactItems(model)
	assert.Nil(t, err)
	checkResult(result)
}
