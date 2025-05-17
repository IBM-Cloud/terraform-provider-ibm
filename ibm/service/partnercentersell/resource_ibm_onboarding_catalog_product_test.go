// Copyright IBM Corp. 2025 All Rights Reserved.
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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/partnercentersell"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnercentersellv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmOnboardingCatalogProductBasic(t *testing.T) {
	var conf partnercentersellv1.GlobalCatalogProduct
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	active := "true"
	disabled := "true"
	kind := "service"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	activeUpdate := "false"
	disabledUpdate := "false"
	kindUpdate := "composite"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingCatalogProductDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogProductConfigBasic(name, active, disabled, kind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingCatalogProductExists("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "active", active),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "disabled", disabled),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "kind", kind),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogProductConfigBasic(nameUpdate, activeUpdate, disabledUpdate, kindUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "active", activeUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "disabled", disabledUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "kind", kindUpdate),
				),
			},
		},
	})
}

func TestAccIbmOnboardingCatalogProductAllArgs(t *testing.T) {
	var conf partnercentersellv1.GlobalCatalogProduct
	env := fmt.Sprintf("tf_env_%d", acctest.RandIntRange(10, 100))
	objectID := fmt.Sprintf("tf_object_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	active := "true"
	disabled := "true"
	kind := "service"
	envUpdate := fmt.Sprintf("tf_env_%d", acctest.RandIntRange(10, 100))
	objectIDUpdate := fmt.Sprintf("tf_object_id_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	activeUpdate := "false"
	disabledUpdate := "false"
	kindUpdate := "composite"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingCatalogProductDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogProductConfig(env, objectID, name, active, disabled, kind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingCatalogProductExists("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "env", env),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "object_id", objectID),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "active", active),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "disabled", disabled),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "kind", kind),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingCatalogProductConfig(envUpdate, objectIDUpdate, nameUpdate, activeUpdate, disabledUpdate, kindUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "env", envUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "object_id", objectIDUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "active", activeUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "disabled", disabledUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_catalog_product.onboarding_catalog_product_instance", "kind", kindUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_onboarding_catalog_product.onboarding_catalog_product",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmOnboardingCatalogProductConfigBasic(name string, active string, disabled string, kind string) string {
	return fmt.Sprintf(`
		resource "ibm_onboarding_catalog_product" "onboarding_catalog_product_instance" {
			product_id = ibm_onboarding_product.onboarding_product_instance.id
			name = "%s"
			active = %s
			disabled = %s
			kind = "%s"
			tags = "FIXME"
			object_provider {
				name = "name"
				email = "email"
			}
		}
	`, name, active, disabled, kind)
}

func testAccCheckIbmOnboardingCatalogProductConfig(env string, objectID string, name string, active string, disabled string, kind string) string {
	return fmt.Sprintf(`

		resource "ibm_onboarding_catalog_product" "onboarding_catalog_product_instance" {
			product_id = ibm_onboarding_product.onboarding_product_instance.id
			env = "%s"
			object_id = "%s"
			name = "%s"
			active = %s
			disabled = %s
			kind = "%s"
			overview_ui {
				en {
					display_name = "display_name"
					description = "description"
					long_description = "long_description"
				}
			}
			tags = "FIXME"
			images {
				image = "image"
			}
			object_provider {
				name = "name"
				email = "email"
			}
			metadata {
				rc_compatible = true
				ui {
					strings {
						en {
							bullets {
								description = "description"
								title = "title"
							}
							media {
								caption = "caption"
								thumbnail = "thumbnail"
								type = "image"
								url = "url"
							}
							navigation_items {
								id = "id"
								url = "url"
								label = "label"
							}
						}
					}
					urls {
						doc_url = "doc_url"
						apidocs_url = "apidocs_url"
						terms_url = "terms_url"
						instructions_url = "instructions_url"
						catalog_details_url = "catalog_details_url"
						custom_create_page_url = "custom_create_page_url"
						dashboard = "dashboard"
					}
					hidden = true
					side_by_side_index = 1.0
					embeddable_dashboard = "embeddable_dashboard"
					accessible_during_provision = true
					primary_offering_id = "primary_offering_id"
				}
				service {
					rc_provisionable = true
					iam_compatible = true
					service_key_supported = true
					unique_api_key = true
					async_provisioning_supported = true
					async_unprovisioning_supported = true
					custom_create_page_hybrid_enabled = true
					parameters {
						displayname = "displayname"
						name = "name"
						type = "text"
						options {
							displayname = "displayname"
							value = "value"
							i18n {
								en {
									displayname = "displayname"
									description = "description"
								}
								de {
									displayname = "displayname"
									description = "description"
								}
								es {
									displayname = "displayname"
									description = "description"
								}
								fr {
									displayname = "displayname"
									description = "description"
								}
								it {
									displayname = "displayname"
									description = "description"
								}
								ja {
									displayname = "displayname"
									description = "description"
								}
								ko {
									displayname = "displayname"
									description = "description"
								}
								pt_br {
									displayname = "displayname"
									description = "description"
								}
								zh_tw {
									displayname = "displayname"
									description = "description"
								}
								zh_cn {
									displayname = "displayname"
									description = "description"
								}
							}
						}
						value = [ "value" ]
						layout = "layout"
						associations {
							plan {
								show_for = [ "show_for" ]
								options_refresh = true
							}
							parameters {
								name = "name"
								show_for = [ "show_for" ]
								options_refresh = true
							}
							location {
								show_for = [ "show_for" ]
							}
						}
						validation_url = "validation_url"
						options_url = "options_url"
						invalidmessage = "invalidmessage"
						description = "description"
						required = true
						pattern = "pattern"
						placeholder = "placeholder"
						readonly = true
						hidden = true
						i18n {
							en {
								displayname = "displayname"
								description = "description"
							}
							de {
								displayname = "displayname"
								description = "description"
							}
							es {
								displayname = "displayname"
								description = "description"
							}
							fr {
								displayname = "displayname"
								description = "description"
							}
							it {
								displayname = "displayname"
								description = "description"
							}
							ja {
								displayname = "displayname"
								description = "description"
							}
							ko {
								displayname = "displayname"
								description = "description"
							}
							pt_br {
								displayname = "displayname"
								description = "description"
							}
							zh_tw {
								displayname = "displayname"
								description = "description"
							}
							zh_cn {
								displayname = "displayname"
								description = "description"
							}
						}
					}
				}
				other {
					pc {
						support {
							url = "url"
							status_url = "status_url"
							locations = [ "locations" ]
							languages = [ "languages" ]
							process = "process"
							process_i18n = { "key" = "inner" }
							support_type = "community"
							support_escalation {
								contact = "contact"
								escalation_wait_time {
									value = 1.0
									type = "type"
								}
								response_wait_time {
									value = 1.0
									type = "type"
								}
							}
							support_details {
								type = "support_site"
								contact = "contact"
								response_wait_time {
									value = 1.0
									type = "type"
								}
								availability {
									times {
										day = 1.0
										start_time = "start_time"
										end_time = "end_time"
									}
									timezone = "timezone"
									always_available = true
								}
							}
						}
					}
					composite {
						composite_kind = "service"
						composite_tag = "composite_tag"
						children {
							kind = "service"
							name = "name"
						}
					}
				}
			}
		}
	`, env, objectID, name, active, disabled, kind)
}

func testAccCheckIbmOnboardingCatalogProductExists(n string, obj partnercentersellv1.GlobalCatalogProduct) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
		if err != nil {
			return err
		}

		getCatalogProductOptions := &partnercentersellv1.GetCatalogProductOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getCatalogProductOptions.SetProductID(parts[0])
		getCatalogProductOptions.SetCatalogProductID(parts[1])

		globalCatalogProduct, _, err := partnerCenterSellClient.GetCatalogProduct(getCatalogProductOptions)
		if err != nil {
			return err
		}

		obj = *globalCatalogProduct
		return nil
	}
}

func testAccCheckIbmOnboardingCatalogProductDestroy(s *terraform.State) error {
	partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_onboarding_catalog_product" {
			continue
		}

		getCatalogProductOptions := &partnercentersellv1.GetCatalogProductOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getCatalogProductOptions.SetProductID(parts[0])
		getCatalogProductOptions.SetCatalogProductID(parts[1])

		// Try to find the key
		_, response, err := partnerCenterSellClient.GetCatalogProduct(getCatalogProductOptions)

		if err == nil {
			return fmt.Errorf("onboarding_catalog_product still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for onboarding_catalog_product (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogOverviewUIToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogOverviewUiTranslatedContentModel := make(map[string]interface{})
		globalCatalogOverviewUiTranslatedContentModel["display_name"] = "testString"
		globalCatalogOverviewUiTranslatedContentModel["description"] = "testString"
		globalCatalogOverviewUiTranslatedContentModel["long_description"] = "testString"

		model := make(map[string]interface{})
		model["en"] = []map[string]interface{}{globalCatalogOverviewUiTranslatedContentModel}

		assert.Equal(t, result, model)
	}

	globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
	globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
	globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
	globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

	model := new(partnercentersellv1.GlobalCatalogOverviewUI)
	model.En = globalCatalogOverviewUiTranslatedContentModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogOverviewUIToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogOverviewUITranslatedContentToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["display_name"] = "testString"
		model["description"] = "testString"
		model["long_description"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
	model.DisplayName = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")
	model.LongDescription = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogOverviewUITranslatedContentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogProductImagesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["image"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogProductImages)
	model.Image = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogProductImagesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductCatalogProductProviderToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["email"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.CatalogProductProvider)
	model.Name = core.StringPtr("testString")
	model.Email = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductCatalogProductProviderToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		catalogHighlightItemModel := make(map[string]interface{})
		catalogHighlightItemModel["description"] = "testString"
		catalogHighlightItemModel["title"] = "testString"

		catalogProductMediaItemModel := make(map[string]interface{})
		catalogProductMediaItemModel["caption"] = "testString"
		catalogProductMediaItemModel["thumbnail"] = "testString"
		catalogProductMediaItemModel["type"] = "image"
		catalogProductMediaItemModel["url"] = "testString"

		globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
		globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

		globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsContentModel["bullets"] = []map[string]interface{}{catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel["media"] = []map[string]interface{}{catalogProductMediaItemModel}
		globalCatalogMetadataUiStringsContentModel["navigation_items"] = []map[string]interface{}{globalCatalogMetadataUiNavigationItemModel}

		globalCatalogMetadataUiStringsModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsModel["en"] = []map[string]interface{}{globalCatalogMetadataUiStringsContentModel}

		globalCatalogMetadataUiUrlsModel := make(map[string]interface{})
		globalCatalogMetadataUiUrlsModel["doc_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["apidocs_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["terms_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["instructions_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["catalog_details_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["custom_create_page_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["dashboard"] = "testString"

		globalCatalogProductMetadataUiModel := make(map[string]interface{})
		globalCatalogProductMetadataUiModel["strings"] = []map[string]interface{}{globalCatalogMetadataUiStringsModel}
		globalCatalogProductMetadataUiModel["urls"] = []map[string]interface{}{globalCatalogMetadataUiUrlsModel}
		globalCatalogProductMetadataUiModel["hidden"] = true
		globalCatalogProductMetadataUiModel["side_by_side_index"] = float64(72.5)
		globalCatalogProductMetadataUiModel["embeddable_dashboard"] = "testString"
		globalCatalogProductMetadataUiModel["accessible_during_provision"] = true
		globalCatalogProductMetadataUiModel["primary_offering_id"] = "testString"

		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

		globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

		globalCatalogMetadataServiceCustomParametersOptionsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersOptionsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersOptionsModel["value"] = "testString"
		globalCatalogMetadataServiceCustomParametersOptionsModel["i18n"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []string{"testString"}

		globalCatalogMetadataServiceCustomParametersAssociationsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsModel["plan"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel["parameters"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel["location"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

		globalCatalogMetadataServiceCustomParametersModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["name"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["type"] = "text"
		globalCatalogMetadataServiceCustomParametersModel["options"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
		globalCatalogMetadataServiceCustomParametersModel["value"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersModel["layout"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["associations"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsModel}
		globalCatalogMetadataServiceCustomParametersModel["validation_url"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["options_url"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["invalidmessage"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["description"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["required"] = true
		globalCatalogMetadataServiceCustomParametersModel["pattern"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["placeholder"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["readonly"] = true
		globalCatalogMetadataServiceCustomParametersModel["hidden"] = true
		globalCatalogMetadataServiceCustomParametersModel["i18n"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

		globalCatalogProductMetadataServiceModel := make(map[string]interface{})
		globalCatalogProductMetadataServiceModel["rc_provisionable"] = true
		globalCatalogProductMetadataServiceModel["iam_compatible"] = true
		globalCatalogProductMetadataServiceModel["service_key_supported"] = true
		globalCatalogProductMetadataServiceModel["unique_api_key"] = true
		globalCatalogProductMetadataServiceModel["async_provisioning_supported"] = true
		globalCatalogProductMetadataServiceModel["async_unprovisioning_supported"] = true
		globalCatalogProductMetadataServiceModel["custom_create_page_hybrid_enabled"] = true
		globalCatalogProductMetadataServiceModel["parameters"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersModel}

		supportTimeIntervalModel := make(map[string]interface{})
		supportTimeIntervalModel["value"] = float64(72.5)
		supportTimeIntervalModel["type"] = "testString"

		supportEscalationModel := make(map[string]interface{})
		supportEscalationModel["contact"] = "testString"
		supportEscalationModel["escalation_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportEscalationModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}

		supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
		supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
		supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
		supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

		supportDetailsItemAvailabilityModel := make(map[string]interface{})
		supportDetailsItemAvailabilityModel["times"] = []map[string]interface{}{supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel["timezone"] = "testString"
		supportDetailsItemAvailabilityModel["always_available"] = true

		supportDetailsItemModel := make(map[string]interface{})
		supportDetailsItemModel["type"] = "support_site"
		supportDetailsItemModel["contact"] = "testString"
		supportDetailsItemModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportDetailsItemModel["availability"] = []map[string]interface{}{supportDetailsItemAvailabilityModel}

		globalCatalogProductMetadataOtherPcSupportModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherPcSupportModel["url"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["status_url"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["locations"] = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel["languages"] = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel["process"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["process_i18n"] = map[string]interface{}{"key1": "testString"}
		globalCatalogProductMetadataOtherPcSupportModel["support_type"] = "community"
		globalCatalogProductMetadataOtherPcSupportModel["support_escalation"] = []map[string]interface{}{supportEscalationModel}
		globalCatalogProductMetadataOtherPcSupportModel["support_details"] = []map[string]interface{}{supportDetailsItemModel}

		globalCatalogProductMetadataOtherPcModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherPcModel["support"] = []map[string]interface{}{globalCatalogProductMetadataOtherPcSupportModel}

		globalCatalogProductMetadataOtherCompositeChildModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherCompositeChildModel["kind"] = "service"
		globalCatalogProductMetadataOtherCompositeChildModel["name"] = "testString"

		globalCatalogProductMetadataOtherCompositeModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherCompositeModel["composite_kind"] = "service"
		globalCatalogProductMetadataOtherCompositeModel["composite_tag"] = "testString"
		globalCatalogProductMetadataOtherCompositeModel["children"] = []map[string]interface{}{globalCatalogProductMetadataOtherCompositeChildModel}

		globalCatalogProductMetadataOtherModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherModel["pc"] = []map[string]interface{}{globalCatalogProductMetadataOtherPcModel}
		globalCatalogProductMetadataOtherModel["composite"] = []map[string]interface{}{globalCatalogProductMetadataOtherCompositeModel}

		model := make(map[string]interface{})
		model["rc_compatible"] = true
		model["ui"] = []map[string]interface{}{globalCatalogProductMetadataUiModel}
		model["service"] = []map[string]interface{}{globalCatalogProductMetadataServiceModel}
		model["other"] = []map[string]interface{}{globalCatalogProductMetadataOtherModel}

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
	catalogHighlightItemModel.Description = core.StringPtr("testString")
	catalogHighlightItemModel.Title = core.StringPtr("testString")

	catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
	catalogProductMediaItemModel.Caption = core.StringPtr("testString")
	catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
	catalogProductMediaItemModel.Type = core.StringPtr("image")
	catalogProductMediaItemModel.URL = core.StringPtr("testString")

	globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
	globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

	globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
	globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
	globalCatalogMetadataUiStringsContentModel.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

	globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
	globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

	globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
	globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

	globalCatalogProductMetadataUiModel := new(partnercentersellv1.GlobalCatalogProductMetadataUI)
	globalCatalogProductMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
	globalCatalogProductMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
	globalCatalogProductMetadataUiModel.Hidden = core.BoolPtr(true)
	globalCatalogProductMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))
	globalCatalogProductMetadataUiModel.EmbeddableDashboard = core.StringPtr("testString")
	globalCatalogProductMetadataUiModel.AccessibleDuringProvision = core.BoolPtr(true)
	globalCatalogProductMetadataUiModel.PrimaryOfferingID = core.StringPtr("testString")

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

	globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
	globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

	globalCatalogMetadataServiceCustomParametersOptionsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
	globalCatalogMetadataServiceCustomParametersOptionsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersOptionsModel.Value = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersOptionsModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

	globalCatalogMetadataServiceCustomParametersAssociationsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

	globalCatalogMetadataServiceCustomParametersModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
	globalCatalogMetadataServiceCustomParametersModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Name = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Type = core.StringPtr("text")
	globalCatalogMetadataServiceCustomParametersModel.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
	globalCatalogMetadataServiceCustomParametersModel.Value = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersModel.Layout = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Associations = globalCatalogMetadataServiceCustomParametersAssociationsModel
	globalCatalogMetadataServiceCustomParametersModel.ValidationURL = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.OptionsURL = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Invalidmessage = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Description = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Required = core.BoolPtr(true)
	globalCatalogMetadataServiceCustomParametersModel.Pattern = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Placeholder = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Readonly = core.BoolPtr(true)
	globalCatalogMetadataServiceCustomParametersModel.Hidden = core.BoolPtr(true)
	globalCatalogMetadataServiceCustomParametersModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

	globalCatalogProductMetadataServiceModel := new(partnercentersellv1.GlobalCatalogProductMetadataService)
	globalCatalogProductMetadataServiceModel.RcProvisionable = core.BoolPtr(true)
	globalCatalogProductMetadataServiceModel.IamCompatible = core.BoolPtr(true)
	globalCatalogProductMetadataServiceModel.ServiceKeySupported = core.BoolPtr(true)
	globalCatalogProductMetadataServiceModel.UniqueApiKey = core.BoolPtr(true)
	globalCatalogProductMetadataServiceModel.AsyncProvisioningSupported = core.BoolPtr(true)
	globalCatalogProductMetadataServiceModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
	globalCatalogProductMetadataServiceModel.CustomCreatePageHybridEnabled = core.BoolPtr(true)
	globalCatalogProductMetadataServiceModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{*globalCatalogMetadataServiceCustomParametersModel}

	supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
	supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
	supportTimeIntervalModel.Type = core.StringPtr("testString")

	supportEscalationModel := new(partnercentersellv1.SupportEscalation)
	supportEscalationModel.Contact = core.StringPtr("testString")
	supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
	supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

	supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
	supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
	supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
	supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

	supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
	supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
	supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

	supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
	supportDetailsItemModel.Type = core.StringPtr("support_site")
	supportDetailsItemModel.Contact = core.StringPtr("testString")
	supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
	supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

	globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
	globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
	globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
	globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
	globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

	globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
	globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

	globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
	globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
	globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

	globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
	globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
	globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
	globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

	globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
	globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel
	globalCatalogProductMetadataOtherModel.Composite = globalCatalogProductMetadataOtherCompositeModel

	model := new(partnercentersellv1.GlobalCatalogProductMetadata)
	model.RcCompatible = core.BoolPtr(true)
	model.Ui = globalCatalogProductMetadataUiModel
	model.Service = globalCatalogProductMetadataServiceModel
	model.Other = globalCatalogProductMetadataOtherModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataUIToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		catalogHighlightItemModel := make(map[string]interface{})
		catalogHighlightItemModel["description"] = "testString"
		catalogHighlightItemModel["title"] = "testString"

		catalogProductMediaItemModel := make(map[string]interface{})
		catalogProductMediaItemModel["caption"] = "testString"
		catalogProductMediaItemModel["thumbnail"] = "testString"
		catalogProductMediaItemModel["type"] = "image"
		catalogProductMediaItemModel["url"] = "testString"

		globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
		globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

		globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsContentModel["bullets"] = []map[string]interface{}{catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel["media"] = []map[string]interface{}{catalogProductMediaItemModel}
		globalCatalogMetadataUiStringsContentModel["navigation_items"] = []map[string]interface{}{globalCatalogMetadataUiNavigationItemModel}

		globalCatalogMetadataUiStringsModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsModel["en"] = []map[string]interface{}{globalCatalogMetadataUiStringsContentModel}

		globalCatalogMetadataUiUrlsModel := make(map[string]interface{})
		globalCatalogMetadataUiUrlsModel["doc_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["apidocs_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["terms_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["instructions_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["catalog_details_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["custom_create_page_url"] = "testString"
		globalCatalogMetadataUiUrlsModel["dashboard"] = "testString"

		model := make(map[string]interface{})
		model["strings"] = []map[string]interface{}{globalCatalogMetadataUiStringsModel}
		model["urls"] = []map[string]interface{}{globalCatalogMetadataUiUrlsModel}
		model["hidden"] = true
		model["side_by_side_index"] = float64(72.5)
		model["embeddable_dashboard"] = "testString"
		model["accessible_during_provision"] = true
		model["primary_offering_id"] = "testString"

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
	catalogHighlightItemModel.Description = core.StringPtr("testString")
	catalogHighlightItemModel.Title = core.StringPtr("testString")

	catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
	catalogProductMediaItemModel.Caption = core.StringPtr("testString")
	catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
	catalogProductMediaItemModel.Type = core.StringPtr("image")
	catalogProductMediaItemModel.URL = core.StringPtr("testString")

	globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
	globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

	globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
	globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
	globalCatalogMetadataUiStringsContentModel.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

	globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
	globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

	globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
	globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
	globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

	model := new(partnercentersellv1.GlobalCatalogProductMetadataUI)
	model.Strings = globalCatalogMetadataUiStringsModel
	model.Urls = globalCatalogMetadataUiUrlsModel
	model.Hidden = core.BoolPtr(true)
	model.SideBySideIndex = core.Float64Ptr(float64(72.5))
	model.EmbeddableDashboard = core.StringPtr("testString")
	model.AccessibleDuringProvision = core.BoolPtr(true)
	model.PrimaryOfferingID = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataUIToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUIStringsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		catalogHighlightItemModel := make(map[string]interface{})
		catalogHighlightItemModel["description"] = "testString"
		catalogHighlightItemModel["title"] = "testString"

		catalogProductMediaItemModel := make(map[string]interface{})
		catalogProductMediaItemModel["caption"] = "testString"
		catalogProductMediaItemModel["thumbnail"] = "testString"
		catalogProductMediaItemModel["type"] = "image"
		catalogProductMediaItemModel["url"] = "testString"

		globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
		globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

		globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
		globalCatalogMetadataUiStringsContentModel["bullets"] = []map[string]interface{}{catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel["media"] = []map[string]interface{}{catalogProductMediaItemModel}
		globalCatalogMetadataUiStringsContentModel["navigation_items"] = []map[string]interface{}{globalCatalogMetadataUiNavigationItemModel}

		model := make(map[string]interface{})
		model["en"] = []map[string]interface{}{globalCatalogMetadataUiStringsContentModel}

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
	catalogHighlightItemModel.Description = core.StringPtr("testString")
	catalogHighlightItemModel.Title = core.StringPtr("testString")

	catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
	catalogProductMediaItemModel.Caption = core.StringPtr("testString")
	catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
	catalogProductMediaItemModel.Type = core.StringPtr("image")
	catalogProductMediaItemModel.URL = core.StringPtr("testString")

	globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
	globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

	globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
	globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
	globalCatalogMetadataUiStringsContentModel.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

	model := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
	model.En = globalCatalogMetadataUiStringsContentModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUIStringsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUIStringsContentToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		catalogHighlightItemModel := make(map[string]interface{})
		catalogHighlightItemModel["description"] = "testString"
		catalogHighlightItemModel["title"] = "testString"

		catalogProductMediaItemModel := make(map[string]interface{})
		catalogProductMediaItemModel["caption"] = "testString"
		catalogProductMediaItemModel["thumbnail"] = "testString"
		catalogProductMediaItemModel["type"] = "image"
		catalogProductMediaItemModel["url"] = "testString"

		globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
		globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
		globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

		model := make(map[string]interface{})
		model["bullets"] = []map[string]interface{}{catalogHighlightItemModel}
		model["media"] = []map[string]interface{}{catalogProductMediaItemModel}
		model["navigation_items"] = []map[string]interface{}{globalCatalogMetadataUiNavigationItemModel}

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
	catalogHighlightItemModel.Description = core.StringPtr("testString")
	catalogHighlightItemModel.Title = core.StringPtr("testString")

	catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
	catalogProductMediaItemModel.Caption = core.StringPtr("testString")
	catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
	catalogProductMediaItemModel.Type = core.StringPtr("image")
	catalogProductMediaItemModel.URL = core.StringPtr("testString")

	globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
	globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
	globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

	model := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
	model.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
	model.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
	model.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUIStringsContentToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductCatalogHighlightItemToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["description"] = "testString"
		model["title"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.CatalogHighlightItem)
	model.Description = core.StringPtr("testString")
	model.Title = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductCatalogHighlightItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductCatalogProductMediaItemToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["caption"] = "testString"
		model["thumbnail"] = "testString"
		model["type"] = "image"
		model["url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.CatalogProductMediaItem)
	model.Caption = core.StringPtr("testString")
	model.Thumbnail = core.StringPtr("testString")
	model.Type = core.StringPtr("image")
	model.URL = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductCatalogProductMediaItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUINavigationItemToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["url"] = "testString"
		model["label"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
	model.ID = core.StringPtr("testString")
	model.URL = core.StringPtr("testString")
	model.Label = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUINavigationItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUIUrlsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["doc_url"] = "testString"
		model["apidocs_url"] = "testString"
		model["terms_url"] = "testString"
		model["instructions_url"] = "testString"
		model["catalog_details_url"] = "testString"
		model["custom_create_page_url"] = "testString"
		model["dashboard"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
	model.DocURL = core.StringPtr("testString")
	model.ApidocsURL = core.StringPtr("testString")
	model.TermsURL = core.StringPtr("testString")
	model.InstructionsURL = core.StringPtr("testString")
	model.CatalogDetailsURL = core.StringPtr("testString")
	model.CustomCreatePageURL = core.StringPtr("testString")
	model.Dashboard = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataUIUrlsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataServiceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

		globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

		globalCatalogMetadataServiceCustomParametersOptionsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersOptionsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersOptionsModel["value"] = "testString"
		globalCatalogMetadataServiceCustomParametersOptionsModel["i18n"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []string{"testString"}

		globalCatalogMetadataServiceCustomParametersAssociationsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsModel["plan"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel["parameters"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel["location"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

		globalCatalogMetadataServiceCustomParametersModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["name"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["type"] = "text"
		globalCatalogMetadataServiceCustomParametersModel["options"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
		globalCatalogMetadataServiceCustomParametersModel["value"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersModel["layout"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["associations"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsModel}
		globalCatalogMetadataServiceCustomParametersModel["validation_url"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["options_url"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["invalidmessage"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["description"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["required"] = true
		globalCatalogMetadataServiceCustomParametersModel["pattern"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["placeholder"] = "testString"
		globalCatalogMetadataServiceCustomParametersModel["readonly"] = true
		globalCatalogMetadataServiceCustomParametersModel["hidden"] = true
		globalCatalogMetadataServiceCustomParametersModel["i18n"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

		model := make(map[string]interface{})
		model["rc_provisionable"] = true
		model["iam_compatible"] = true
		model["bindable"] = true
		model["plan_updateable"] = true
		model["service_key_supported"] = true
		model["unique_api_key"] = true
		model["async_provisioning_supported"] = true
		model["async_unprovisioning_supported"] = true
		model["custom_create_page_hybrid_enabled"] = true
		model["parameters"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersModel}

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

	globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
	globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

	globalCatalogMetadataServiceCustomParametersOptionsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
	globalCatalogMetadataServiceCustomParametersOptionsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersOptionsModel.Value = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersOptionsModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

	globalCatalogMetadataServiceCustomParametersAssociationsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

	globalCatalogMetadataServiceCustomParametersModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
	globalCatalogMetadataServiceCustomParametersModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Name = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Type = core.StringPtr("text")
	globalCatalogMetadataServiceCustomParametersModel.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
	globalCatalogMetadataServiceCustomParametersModel.Value = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersModel.Layout = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Associations = globalCatalogMetadataServiceCustomParametersAssociationsModel
	globalCatalogMetadataServiceCustomParametersModel.ValidationURL = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.OptionsURL = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Invalidmessage = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Description = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Required = core.BoolPtr(true)
	globalCatalogMetadataServiceCustomParametersModel.Pattern = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Placeholder = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersModel.Readonly = core.BoolPtr(true)
	globalCatalogMetadataServiceCustomParametersModel.Hidden = core.BoolPtr(true)
	globalCatalogMetadataServiceCustomParametersModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

	model := new(partnercentersellv1.GlobalCatalogProductMetadataService)
	model.RcProvisionable = core.BoolPtr(true)
	model.IamCompatible = core.BoolPtr(true)
	model.Bindable = core.BoolPtr(true)
	model.PlanUpdateable = core.BoolPtr(true)
	model.ServiceKeySupported = core.BoolPtr(true)
	model.UniqueApiKey = core.BoolPtr(true)
	model.AsyncProvisioningSupported = core.BoolPtr(true)
	model.AsyncUnprovisioningSupported = core.BoolPtr(true)
	model.CustomCreatePageHybridEnabled = core.BoolPtr(true)
	model.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{*globalCatalogMetadataServiceCustomParametersModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataServiceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

		globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

		globalCatalogMetadataServiceCustomParametersOptionsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersOptionsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersOptionsModel["value"] = "testString"
		globalCatalogMetadataServiceCustomParametersOptionsModel["i18n"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []string{"testString"}

		globalCatalogMetadataServiceCustomParametersAssociationsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsModel["plan"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel["parameters"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel["location"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

		model := make(map[string]interface{})
		model["displayname"] = "testString"
		model["name"] = "testString"
		model["type"] = "text"
		model["options"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
		model["value"] = []string{"testString"}
		model["layout"] = "testString"
		model["associations"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsModel}
		model["validation_url"] = "testString"
		model["options_url"] = "testString"
		model["invalidmessage"] = "testString"
		model["description"] = "testString"
		model["required"] = true
		model["pattern"] = "testString"
		model["placeholder"] = "testString"
		model["readonly"] = true
		model["hidden"] = true
		model["i18n"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

	globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
	globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

	globalCatalogMetadataServiceCustomParametersOptionsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
	globalCatalogMetadataServiceCustomParametersOptionsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersOptionsModel.Value = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersOptionsModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

	globalCatalogMetadataServiceCustomParametersAssociationsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
	model.Displayname = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.Type = core.StringPtr("text")
	model.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
	model.Value = []string{"testString"}
	model.Layout = core.StringPtr("testString")
	model.Associations = globalCatalogMetadataServiceCustomParametersAssociationsModel
	model.ValidationURL = core.StringPtr("testString")
	model.OptionsURL = core.StringPtr("testString")
	model.Invalidmessage = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")
	model.Required = core.BoolPtr(true)
	model.Pattern = core.StringPtr("testString")
	model.Placeholder = core.StringPtr("testString")
	model.Readonly = core.BoolPtr(true)
	model.Hidden = core.BoolPtr(true)
	model.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersOptionsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

		globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

		model := make(map[string]interface{})
		model["displayname"] = "testString"
		model["value"] = "testString"
		model["i18n"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

	globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
	globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
	model.Displayname = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")
	model.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersOptionsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersI18nToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

		model := make(map[string]interface{})
		model["en"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["de"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["es"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["fr"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["it"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["ja"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["ko"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["pt_br"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["zh_tw"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
		model["zh_cn"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
	model.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
	model.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersI18nToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersI18nFieldsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["displayname"] = "testString"
		model["description"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
	model.Displayname = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersI18nFieldsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersAssociationsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []string{"testString"}

		model := make(map[string]interface{})
		model["plan"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
		model["parameters"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		model["location"] = []map[string]interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
	model.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
	model.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	model.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersAssociationsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersAssociationsPlanToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["show_for"] = []string{"testString"}
		model["options_refresh"] = true

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
	model.ShowFor = []string{"testString"}
	model.OptionsRefresh = core.BoolPtr(true)

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersAssociationsPlanToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersAssociationsParametersItemToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["show_for"] = []string{"testString"}
		model["options_refresh"] = true

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
	model.Name = core.StringPtr("testString")
	model.ShowFor = []string{"testString"}
	model.OptionsRefresh = core.BoolPtr(true)

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersAssociationsParametersItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersAssociationsLocationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["show_for"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
	model.ShowFor = []string{"testString"}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogMetadataServiceCustomParametersAssociationsLocationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportTimeIntervalModel := make(map[string]interface{})
		supportTimeIntervalModel["value"] = float64(72.5)
		supportTimeIntervalModel["type"] = "testString"

		supportEscalationModel := make(map[string]interface{})
		supportEscalationModel["contact"] = "testString"
		supportEscalationModel["escalation_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportEscalationModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}

		supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
		supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
		supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
		supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

		supportDetailsItemAvailabilityModel := make(map[string]interface{})
		supportDetailsItemAvailabilityModel["times"] = []map[string]interface{}{supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel["timezone"] = "testString"
		supportDetailsItemAvailabilityModel["always_available"] = true

		supportDetailsItemModel := make(map[string]interface{})
		supportDetailsItemModel["type"] = "support_site"
		supportDetailsItemModel["contact"] = "testString"
		supportDetailsItemModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportDetailsItemModel["availability"] = []map[string]interface{}{supportDetailsItemAvailabilityModel}

		globalCatalogProductMetadataOtherPcSupportModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherPcSupportModel["url"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["status_url"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["locations"] = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel["languages"] = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel["process"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["process_i18n"] = map[string]interface{}{"key1": "testString"}
		globalCatalogProductMetadataOtherPcSupportModel["support_type"] = "community"
		globalCatalogProductMetadataOtherPcSupportModel["support_escalation"] = []map[string]interface{}{supportEscalationModel}
		globalCatalogProductMetadataOtherPcSupportModel["support_details"] = []map[string]interface{}{supportDetailsItemModel}

		globalCatalogProductMetadataOtherPcModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherPcModel["support"] = []map[string]interface{}{globalCatalogProductMetadataOtherPcSupportModel}

		globalCatalogProductMetadataOtherCompositeChildModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherCompositeChildModel["kind"] = "service"
		globalCatalogProductMetadataOtherCompositeChildModel["name"] = "testString"

		globalCatalogProductMetadataOtherCompositeModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherCompositeModel["composite_kind"] = "service"
		globalCatalogProductMetadataOtherCompositeModel["composite_tag"] = "testString"
		globalCatalogProductMetadataOtherCompositeModel["children"] = []map[string]interface{}{globalCatalogProductMetadataOtherCompositeChildModel}

		model := make(map[string]interface{})
		model["pc"] = []map[string]interface{}{globalCatalogProductMetadataOtherPcModel}
		model["composite"] = []map[string]interface{}{globalCatalogProductMetadataOtherCompositeModel}

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
	supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
	supportTimeIntervalModel.Type = core.StringPtr("testString")

	supportEscalationModel := new(partnercentersellv1.SupportEscalation)
	supportEscalationModel.Contact = core.StringPtr("testString")
	supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
	supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

	supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
	supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
	supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
	supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

	supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
	supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
	supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

	supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
	supportDetailsItemModel.Type = core.StringPtr("support_site")
	supportDetailsItemModel.Contact = core.StringPtr("testString")
	supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
	supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

	globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
	globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
	globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
	globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
	globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

	globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
	globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

	globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
	globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
	globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

	globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
	globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
	globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
	globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

	model := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
	model.PC = globalCatalogProductMetadataOtherPcModel
	model.Composite = globalCatalogProductMetadataOtherCompositeModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherPCToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportTimeIntervalModel := make(map[string]interface{})
		supportTimeIntervalModel["value"] = float64(72.5)
		supportTimeIntervalModel["type"] = "testString"

		supportEscalationModel := make(map[string]interface{})
		supportEscalationModel["contact"] = "testString"
		supportEscalationModel["escalation_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportEscalationModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}

		supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
		supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
		supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
		supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

		supportDetailsItemAvailabilityModel := make(map[string]interface{})
		supportDetailsItemAvailabilityModel["times"] = []map[string]interface{}{supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel["timezone"] = "testString"
		supportDetailsItemAvailabilityModel["always_available"] = true

		supportDetailsItemModel := make(map[string]interface{})
		supportDetailsItemModel["type"] = "support_site"
		supportDetailsItemModel["contact"] = "testString"
		supportDetailsItemModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportDetailsItemModel["availability"] = []map[string]interface{}{supportDetailsItemAvailabilityModel}

		globalCatalogProductMetadataOtherPcSupportModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherPcSupportModel["url"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["status_url"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["locations"] = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel["languages"] = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel["process"] = "testString"
		globalCatalogProductMetadataOtherPcSupportModel["process_i18n"] = map[string]interface{}{"key1": "testString"}
		globalCatalogProductMetadataOtherPcSupportModel["support_type"] = "community"
		globalCatalogProductMetadataOtherPcSupportModel["support_escalation"] = []map[string]interface{}{supportEscalationModel}
		globalCatalogProductMetadataOtherPcSupportModel["support_details"] = []map[string]interface{}{supportDetailsItemModel}

		model := make(map[string]interface{})
		model["support"] = []map[string]interface{}{globalCatalogProductMetadataOtherPcSupportModel}

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
	supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
	supportTimeIntervalModel.Type = core.StringPtr("testString")

	supportEscalationModel := new(partnercentersellv1.SupportEscalation)
	supportEscalationModel.Contact = core.StringPtr("testString")
	supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
	supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

	supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
	supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
	supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
	supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

	supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
	supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
	supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

	supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
	supportDetailsItemModel.Type = core.StringPtr("support_site")
	supportDetailsItemModel.Contact = core.StringPtr("testString")
	supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
	supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

	globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
	globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
	globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
	globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
	globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
	globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

	model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
	model.Support = globalCatalogProductMetadataOtherPcSupportModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherPCToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherPCSupportToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportTimeIntervalModel := make(map[string]interface{})
		supportTimeIntervalModel["value"] = float64(72.5)
		supportTimeIntervalModel["type"] = "testString"

		supportEscalationModel := make(map[string]interface{})
		supportEscalationModel["contact"] = "testString"
		supportEscalationModel["escalation_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportEscalationModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}

		supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
		supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
		supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
		supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

		supportDetailsItemAvailabilityModel := make(map[string]interface{})
		supportDetailsItemAvailabilityModel["times"] = []map[string]interface{}{supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel["timezone"] = "testString"
		supportDetailsItemAvailabilityModel["always_available"] = true

		supportDetailsItemModel := make(map[string]interface{})
		supportDetailsItemModel["type"] = "support_site"
		supportDetailsItemModel["contact"] = "testString"
		supportDetailsItemModel["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		supportDetailsItemModel["availability"] = []map[string]interface{}{supportDetailsItemAvailabilityModel}

		model := make(map[string]interface{})
		model["url"] = "testString"
		model["status_url"] = "testString"
		model["locations"] = []string{"testString"}
		model["languages"] = []string{"testString"}
		model["process"] = "testString"
		model["process_i18n"] = map[string]interface{}{"key1": "testString"}
		model["support_type"] = "community"
		model["support_escalation"] = []map[string]interface{}{supportEscalationModel}
		model["support_details"] = []map[string]interface{}{supportDetailsItemModel}

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
	supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
	supportTimeIntervalModel.Type = core.StringPtr("testString")

	supportEscalationModel := new(partnercentersellv1.SupportEscalation)
	supportEscalationModel.Contact = core.StringPtr("testString")
	supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
	supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

	supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
	supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
	supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
	supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

	supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
	supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
	supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

	supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
	supportDetailsItemModel.Type = core.StringPtr("support_site")
	supportDetailsItemModel.Contact = core.StringPtr("testString")
	supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
	supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

	model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
	model.URL = core.StringPtr("testString")
	model.StatusURL = core.StringPtr("testString")
	model.Locations = []string{"testString"}
	model.Languages = []string{"testString"}
	model.Process = core.StringPtr("testString")
	model.ProcessI18n = map[string]string{"key1": "testString"}
	model.SupportType = core.StringPtr("community")
	model.SupportEscalation = supportEscalationModel
	model.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherPCSupportToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductSupportEscalationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportTimeIntervalModel := make(map[string]interface{})
		supportTimeIntervalModel["value"] = float64(72.5)
		supportTimeIntervalModel["type"] = "testString"

		model := make(map[string]interface{})
		model["contact"] = "testString"
		model["escalation_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		model["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
	supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
	supportTimeIntervalModel.Type = core.StringPtr("testString")

	model := new(partnercentersellv1.SupportEscalation)
	model.Contact = core.StringPtr("testString")
	model.EscalationWaitTime = supportTimeIntervalModel
	model.ResponseWaitTime = supportTimeIntervalModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductSupportEscalationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductSupportTimeIntervalToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["value"] = float64(72.5)
		model["type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.SupportTimeInterval)
	model.Value = core.Float64Ptr(float64(72.5))
	model.Type = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductSupportTimeIntervalToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductSupportDetailsItemToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportTimeIntervalModel := make(map[string]interface{})
		supportTimeIntervalModel["value"] = float64(72.5)
		supportTimeIntervalModel["type"] = "testString"

		supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
		supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
		supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
		supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

		supportDetailsItemAvailabilityModel := make(map[string]interface{})
		supportDetailsItemAvailabilityModel["times"] = []map[string]interface{}{supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel["timezone"] = "testString"
		supportDetailsItemAvailabilityModel["always_available"] = true

		model := make(map[string]interface{})
		model["type"] = "support_site"
		model["contact"] = "testString"
		model["response_wait_time"] = []map[string]interface{}{supportTimeIntervalModel}
		model["availability"] = []map[string]interface{}{supportDetailsItemAvailabilityModel}

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
	supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
	supportTimeIntervalModel.Type = core.StringPtr("testString")

	supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
	supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
	supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
	supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

	supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
	supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
	supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

	model := new(partnercentersellv1.SupportDetailsItem)
	model.Type = core.StringPtr("support_site")
	model.Contact = core.StringPtr("testString")
	model.ResponseWaitTime = supportTimeIntervalModel
	model.Availability = supportDetailsItemAvailabilityModel

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductSupportDetailsItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductSupportDetailsItemAvailabilityToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
		supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
		supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
		supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

		model := make(map[string]interface{})
		model["times"] = []map[string]interface{}{supportDetailsItemAvailabilityTimeModel}
		model["timezone"] = "testString"
		model["always_available"] = true

		assert.Equal(t, result, model)
	}

	supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
	supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
	supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
	supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

	model := new(partnercentersellv1.SupportDetailsItemAvailability)
	model.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
	model.Timezone = core.StringPtr("testString")
	model.AlwaysAvailable = core.BoolPtr(true)

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductSupportDetailsItemAvailabilityToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductSupportDetailsItemAvailabilityTimeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day"] = float64(72.5)
		model["start_time"] = "testString"
		model["end_time"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
	model.Day = core.Float64Ptr(float64(72.5))
	model.StartTime = core.StringPtr("testString")
	model.EndTime = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductSupportDetailsItemAvailabilityTimeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherCompositeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		globalCatalogProductMetadataOtherCompositeChildModel := make(map[string]interface{})
		globalCatalogProductMetadataOtherCompositeChildModel["kind"] = "service"
		globalCatalogProductMetadataOtherCompositeChildModel["name"] = "testString"

		model := make(map[string]interface{})
		model["composite_kind"] = "service"
		model["composite_tag"] = "testString"
		model["children"] = []map[string]interface{}{globalCatalogProductMetadataOtherCompositeChildModel}

		assert.Equal(t, result, model)
	}

	globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
	globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
	globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

	model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
	model.CompositeKind = core.StringPtr("service")
	model.CompositeTag = core.StringPtr("testString")
	model.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherCompositeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherCompositeChildToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["kind"] = "service"
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
	model.Kind = core.StringPtr("service")
	model.Name = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductGlobalCatalogProductMetadataOtherCompositeChildToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToCatalogProductProvider(t *testing.T) {
	checkResult := func(result *partnercentersellv1.CatalogProductProvider) {
		model := new(partnercentersellv1.CatalogProductProvider)
		model.Name = core.StringPtr("testString")
		model.Email = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["email"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToCatalogProductProvider(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogOverviewUI(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogOverviewUI) {
		globalCatalogOverviewUiTranslatedContentModel := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
		globalCatalogOverviewUiTranslatedContentModel.DisplayName = core.StringPtr("testString")
		globalCatalogOverviewUiTranslatedContentModel.Description = core.StringPtr("testString")
		globalCatalogOverviewUiTranslatedContentModel.LongDescription = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogOverviewUI)
		model.En = globalCatalogOverviewUiTranslatedContentModel

		assert.Equal(t, result, model)
	}

	globalCatalogOverviewUiTranslatedContentModel := make(map[string]interface{})
	globalCatalogOverviewUiTranslatedContentModel["display_name"] = "testString"
	globalCatalogOverviewUiTranslatedContentModel["description"] = "testString"
	globalCatalogOverviewUiTranslatedContentModel["long_description"] = "testString"

	model := make(map[string]interface{})
	model["en"] = []interface{}{globalCatalogOverviewUiTranslatedContentModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogOverviewUI(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogOverviewUITranslatedContent(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogOverviewUITranslatedContent) {
		model := new(partnercentersellv1.GlobalCatalogOverviewUITranslatedContent)
		model.DisplayName = core.StringPtr("testString")
		model.Description = core.StringPtr("testString")
		model.LongDescription = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["display_name"] = "testString"
	model["description"] = "testString"
	model["long_description"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogOverviewUITranslatedContent(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductImages(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogProductImages) {
		model := new(partnercentersellv1.GlobalCatalogProductImages)
		model.Image = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["image"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductImages(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataPrototypePatch(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogProductMetadataPrototypePatch) {
		catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
		catalogHighlightItemModel.Description = core.StringPtr("testString")
		catalogHighlightItemModel.Title = core.StringPtr("testString")

		catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
		catalogProductMediaItemModel.Caption = core.StringPtr("testString")
		catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
		catalogProductMediaItemModel.Type = core.StringPtr("image")
		catalogProductMediaItemModel.URL = core.StringPtr("testString")

		globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
		globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

		globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
		globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
		globalCatalogMetadataUiStringsContentModel.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

		globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
		globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

		globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
		globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

		globalCatalogProductMetadataUiModel := new(partnercentersellv1.GlobalCatalogProductMetadataUI)
		globalCatalogProductMetadataUiModel.Strings = globalCatalogMetadataUiStringsModel
		globalCatalogProductMetadataUiModel.Urls = globalCatalogMetadataUiUrlsModel
		globalCatalogProductMetadataUiModel.Hidden = core.BoolPtr(true)
		globalCatalogProductMetadataUiModel.SideBySideIndex = core.Float64Ptr(float64(72.5))
		globalCatalogProductMetadataUiModel.EmbeddableDashboard = core.StringPtr("testString")
		globalCatalogProductMetadataUiModel.AccessibleDuringProvision = core.BoolPtr(true)
		globalCatalogProductMetadataUiModel.PrimaryOfferingID = core.StringPtr("testString")

		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

		globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
		globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

		globalCatalogMetadataServiceCustomParametersOptionsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
		globalCatalogMetadataServiceCustomParametersOptionsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersOptionsModel.Value = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersOptionsModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

		globalCatalogMetadataServiceCustomParametersAssociationsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

		globalCatalogMetadataServiceCustomParametersModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
		globalCatalogMetadataServiceCustomParametersModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Name = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Type = core.StringPtr("text")
		globalCatalogMetadataServiceCustomParametersModel.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
		globalCatalogMetadataServiceCustomParametersModel.Value = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersModel.Layout = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Associations = globalCatalogMetadataServiceCustomParametersAssociationsModel
		globalCatalogMetadataServiceCustomParametersModel.ValidationURL = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.OptionsURL = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Invalidmessage = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Description = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Required = core.BoolPtr(true)
		globalCatalogMetadataServiceCustomParametersModel.Pattern = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Placeholder = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Readonly = core.BoolPtr(true)
		globalCatalogMetadataServiceCustomParametersModel.Hidden = core.BoolPtr(true)
		globalCatalogMetadataServiceCustomParametersModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

		globalCatalogProductMetadataServicePrototypePatchModel := new(partnercentersellv1.GlobalCatalogProductMetadataServicePrototypePatch)
		globalCatalogProductMetadataServicePrototypePatchModel.RcProvisionable = core.BoolPtr(true)
		globalCatalogProductMetadataServicePrototypePatchModel.IamCompatible = core.BoolPtr(true)
		globalCatalogProductMetadataServicePrototypePatchModel.ServiceKeySupported = core.BoolPtr(true)
		globalCatalogProductMetadataServicePrototypePatchModel.UniqueApiKey = core.BoolPtr(true)
		globalCatalogProductMetadataServicePrototypePatchModel.AsyncProvisioningSupported = core.BoolPtr(true)
		globalCatalogProductMetadataServicePrototypePatchModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
		globalCatalogProductMetadataServicePrototypePatchModel.CustomCreatePageHybridEnabled = core.BoolPtr(true)
		globalCatalogProductMetadataServicePrototypePatchModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{*globalCatalogMetadataServiceCustomParametersModel}

		supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
		supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
		supportTimeIntervalModel.Type = core.StringPtr("testString")

		supportEscalationModel := new(partnercentersellv1.SupportEscalation)
		supportEscalationModel.Contact = core.StringPtr("testString")
		supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
		supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

		supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
		supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
		supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
		supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

		supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
		supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
		supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

		supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
		supportDetailsItemModel.Type = core.StringPtr("support_site")
		supportDetailsItemModel.Contact = core.StringPtr("testString")
		supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
		supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

		globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
		globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
		globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
		globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
		globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

		globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
		globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

		globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
		globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
		globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

		globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
		globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
		globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
		globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

		globalCatalogProductMetadataOtherModel := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
		globalCatalogProductMetadataOtherModel.PC = globalCatalogProductMetadataOtherPcModel
		globalCatalogProductMetadataOtherModel.Composite = globalCatalogProductMetadataOtherCompositeModel

		model := new(partnercentersellv1.GlobalCatalogProductMetadataPrototypePatch)
		model.RcCompatible = core.BoolPtr(true)
		model.Ui = globalCatalogProductMetadataUiModel
		model.Service = globalCatalogProductMetadataServicePrototypePatchModel
		model.Other = globalCatalogProductMetadataOtherModel

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := make(map[string]interface{})
	catalogHighlightItemModel["description"] = "testString"
	catalogHighlightItemModel["title"] = "testString"

	catalogProductMediaItemModel := make(map[string]interface{})
	catalogProductMediaItemModel["caption"] = "testString"
	catalogProductMediaItemModel["thumbnail"] = "testString"
	catalogProductMediaItemModel["type"] = "image"
	catalogProductMediaItemModel["url"] = "testString"

	globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
	globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

	globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsContentModel["bullets"] = []interface{}{catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel["media"] = []interface{}{catalogProductMediaItemModel}
	globalCatalogMetadataUiStringsContentModel["navigation_items"] = []interface{}{globalCatalogMetadataUiNavigationItemModel}

	globalCatalogMetadataUiStringsModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsModel["en"] = []interface{}{globalCatalogMetadataUiStringsContentModel}

	globalCatalogMetadataUiUrlsModel := make(map[string]interface{})
	globalCatalogMetadataUiUrlsModel["doc_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["apidocs_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["terms_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["instructions_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["catalog_details_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["custom_create_page_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["dashboard"] = "testString"

	globalCatalogProductMetadataUiModel := make(map[string]interface{})
	globalCatalogProductMetadataUiModel["strings"] = []interface{}{globalCatalogMetadataUiStringsModel}
	globalCatalogProductMetadataUiModel["urls"] = []interface{}{globalCatalogMetadataUiUrlsModel}
	globalCatalogProductMetadataUiModel["hidden"] = true
	globalCatalogProductMetadataUiModel["side_by_side_index"] = float64(72.5)
	globalCatalogProductMetadataUiModel["embeddable_dashboard"] = "testString"
	globalCatalogProductMetadataUiModel["accessible_during_provision"] = true
	globalCatalogProductMetadataUiModel["primary_offering_id"] = "testString"

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

	globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

	globalCatalogMetadataServiceCustomParametersOptionsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersOptionsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersOptionsModel["value"] = "testString"
	globalCatalogMetadataServiceCustomParametersOptionsModel["i18n"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []interface{}{"testString"}

	globalCatalogMetadataServiceCustomParametersAssociationsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsModel["plan"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel["parameters"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel["location"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

	globalCatalogMetadataServiceCustomParametersModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["name"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["type"] = "text"
	globalCatalogMetadataServiceCustomParametersModel["options"] = []interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
	globalCatalogMetadataServiceCustomParametersModel["value"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersModel["layout"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["associations"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsModel}
	globalCatalogMetadataServiceCustomParametersModel["validation_url"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["options_url"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["invalidmessage"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["description"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["required"] = true
	globalCatalogMetadataServiceCustomParametersModel["pattern"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["placeholder"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["readonly"] = true
	globalCatalogMetadataServiceCustomParametersModel["hidden"] = true
	globalCatalogMetadataServiceCustomParametersModel["i18n"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

	globalCatalogProductMetadataServicePrototypePatchModel := make(map[string]interface{})
	globalCatalogProductMetadataServicePrototypePatchModel["rc_provisionable"] = true
	globalCatalogProductMetadataServicePrototypePatchModel["iam_compatible"] = true
	globalCatalogProductMetadataServicePrototypePatchModel["service_key_supported"] = true
	globalCatalogProductMetadataServicePrototypePatchModel["unique_api_key"] = true
	globalCatalogProductMetadataServicePrototypePatchModel["async_provisioning_supported"] = true
	globalCatalogProductMetadataServicePrototypePatchModel["async_unprovisioning_supported"] = true
	globalCatalogProductMetadataServicePrototypePatchModel["custom_create_page_hybrid_enabled"] = true
	globalCatalogProductMetadataServicePrototypePatchModel["parameters"] = []interface{}{globalCatalogMetadataServiceCustomParametersModel}

	supportTimeIntervalModel := make(map[string]interface{})
	supportTimeIntervalModel["value"] = float64(72.5)
	supportTimeIntervalModel["type"] = "testString"

	supportEscalationModel := make(map[string]interface{})
	supportEscalationModel["contact"] = "testString"
	supportEscalationModel["escalation_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportEscalationModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}

	supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
	supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
	supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
	supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

	supportDetailsItemAvailabilityModel := make(map[string]interface{})
	supportDetailsItemAvailabilityModel["times"] = []interface{}{supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel["timezone"] = "testString"
	supportDetailsItemAvailabilityModel["always_available"] = true

	supportDetailsItemModel := make(map[string]interface{})
	supportDetailsItemModel["type"] = "support_site"
	supportDetailsItemModel["contact"] = "testString"
	supportDetailsItemModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportDetailsItemModel["availability"] = []interface{}{supportDetailsItemAvailabilityModel}

	globalCatalogProductMetadataOtherPcSupportModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherPcSupportModel["url"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["status_url"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["locations"] = []interface{}{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel["languages"] = []interface{}{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel["process"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["process_i18n"] = map[string]interface{}{"key1": "testString"}
	globalCatalogProductMetadataOtherPcSupportModel["support_type"] = "community"
	globalCatalogProductMetadataOtherPcSupportModel["support_escalation"] = []interface{}{supportEscalationModel}
	globalCatalogProductMetadataOtherPcSupportModel["support_details"] = []interface{}{supportDetailsItemModel}

	globalCatalogProductMetadataOtherPcModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherPcModel["support"] = []interface{}{globalCatalogProductMetadataOtherPcSupportModel}

	globalCatalogProductMetadataOtherCompositeChildModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherCompositeChildModel["kind"] = "service"
	globalCatalogProductMetadataOtherCompositeChildModel["name"] = "testString"

	globalCatalogProductMetadataOtherCompositeModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherCompositeModel["composite_kind"] = "service"
	globalCatalogProductMetadataOtherCompositeModel["composite_tag"] = "testString"
	globalCatalogProductMetadataOtherCompositeModel["children"] = []interface{}{globalCatalogProductMetadataOtherCompositeChildModel}

	globalCatalogProductMetadataOtherModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherModel["pc"] = []interface{}{globalCatalogProductMetadataOtherPcModel}
	globalCatalogProductMetadataOtherModel["composite"] = []interface{}{globalCatalogProductMetadataOtherCompositeModel}

	model := make(map[string]interface{})
	model["rc_compatible"] = true
	model["ui"] = []interface{}{globalCatalogProductMetadataUiModel}
	model["service"] = []interface{}{globalCatalogProductMetadataServicePrototypePatchModel}
	model["other"] = []interface{}{globalCatalogProductMetadataOtherModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataPrototypePatch(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataUI(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogProductMetadataUI) {
		catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
		catalogHighlightItemModel.Description = core.StringPtr("testString")
		catalogHighlightItemModel.Title = core.StringPtr("testString")

		catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
		catalogProductMediaItemModel.Caption = core.StringPtr("testString")
		catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
		catalogProductMediaItemModel.Type = core.StringPtr("image")
		catalogProductMediaItemModel.URL = core.StringPtr("testString")

		globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
		globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

		globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
		globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
		globalCatalogMetadataUiStringsContentModel.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

		globalCatalogMetadataUiStringsModel := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
		globalCatalogMetadataUiStringsModel.En = globalCatalogMetadataUiStringsContentModel

		globalCatalogMetadataUiUrlsModel := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
		globalCatalogMetadataUiUrlsModel.DocURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.ApidocsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.TermsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.InstructionsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.CatalogDetailsURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.CustomCreatePageURL = core.StringPtr("testString")
		globalCatalogMetadataUiUrlsModel.Dashboard = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogProductMetadataUI)
		model.Strings = globalCatalogMetadataUiStringsModel
		model.Urls = globalCatalogMetadataUiUrlsModel
		model.Hidden = core.BoolPtr(true)
		model.SideBySideIndex = core.Float64Ptr(float64(72.5))
		model.EmbeddableDashboard = core.StringPtr("testString")
		model.AccessibleDuringProvision = core.BoolPtr(true)
		model.PrimaryOfferingID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := make(map[string]interface{})
	catalogHighlightItemModel["description"] = "testString"
	catalogHighlightItemModel["title"] = "testString"

	catalogProductMediaItemModel := make(map[string]interface{})
	catalogProductMediaItemModel["caption"] = "testString"
	catalogProductMediaItemModel["thumbnail"] = "testString"
	catalogProductMediaItemModel["type"] = "image"
	catalogProductMediaItemModel["url"] = "testString"

	globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
	globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

	globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsContentModel["bullets"] = []interface{}{catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel["media"] = []interface{}{catalogProductMediaItemModel}
	globalCatalogMetadataUiStringsContentModel["navigation_items"] = []interface{}{globalCatalogMetadataUiNavigationItemModel}

	globalCatalogMetadataUiStringsModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsModel["en"] = []interface{}{globalCatalogMetadataUiStringsContentModel}

	globalCatalogMetadataUiUrlsModel := make(map[string]interface{})
	globalCatalogMetadataUiUrlsModel["doc_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["apidocs_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["terms_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["instructions_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["catalog_details_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["custom_create_page_url"] = "testString"
	globalCatalogMetadataUiUrlsModel["dashboard"] = "testString"

	model := make(map[string]interface{})
	model["strings"] = []interface{}{globalCatalogMetadataUiStringsModel}
	model["urls"] = []interface{}{globalCatalogMetadataUiUrlsModel}
	model["hidden"] = true
	model["side_by_side_index"] = float64(72.5)
	model["embeddable_dashboard"] = "testString"
	model["accessible_during_provision"] = true
	model["primary_offering_id"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataUI(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUIStrings(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataUIStrings) {
		catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
		catalogHighlightItemModel.Description = core.StringPtr("testString")
		catalogHighlightItemModel.Title = core.StringPtr("testString")

		catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
		catalogProductMediaItemModel.Caption = core.StringPtr("testString")
		catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
		catalogProductMediaItemModel.Type = core.StringPtr("image")
		catalogProductMediaItemModel.URL = core.StringPtr("testString")

		globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
		globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

		globalCatalogMetadataUiStringsContentModel := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
		globalCatalogMetadataUiStringsContentModel.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
		globalCatalogMetadataUiStringsContentModel.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
		globalCatalogMetadataUiStringsContentModel.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

		model := new(partnercentersellv1.GlobalCatalogMetadataUIStrings)
		model.En = globalCatalogMetadataUiStringsContentModel

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := make(map[string]interface{})
	catalogHighlightItemModel["description"] = "testString"
	catalogHighlightItemModel["title"] = "testString"

	catalogProductMediaItemModel := make(map[string]interface{})
	catalogProductMediaItemModel["caption"] = "testString"
	catalogProductMediaItemModel["thumbnail"] = "testString"
	catalogProductMediaItemModel["type"] = "image"
	catalogProductMediaItemModel["url"] = "testString"

	globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
	globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

	globalCatalogMetadataUiStringsContentModel := make(map[string]interface{})
	globalCatalogMetadataUiStringsContentModel["bullets"] = []interface{}{catalogHighlightItemModel}
	globalCatalogMetadataUiStringsContentModel["media"] = []interface{}{catalogProductMediaItemModel}
	globalCatalogMetadataUiStringsContentModel["navigation_items"] = []interface{}{globalCatalogMetadataUiNavigationItemModel}

	model := make(map[string]interface{})
	model["en"] = []interface{}{globalCatalogMetadataUiStringsContentModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUIStrings(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUIStringsContent(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataUIStringsContent) {
		catalogHighlightItemModel := new(partnercentersellv1.CatalogHighlightItem)
		catalogHighlightItemModel.Description = core.StringPtr("testString")
		catalogHighlightItemModel.Title = core.StringPtr("testString")

		catalogProductMediaItemModel := new(partnercentersellv1.CatalogProductMediaItem)
		catalogProductMediaItemModel.Caption = core.StringPtr("testString")
		catalogProductMediaItemModel.Thumbnail = core.StringPtr("testString")
		catalogProductMediaItemModel.Type = core.StringPtr("image")
		catalogProductMediaItemModel.URL = core.StringPtr("testString")

		globalCatalogMetadataUiNavigationItemModel := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
		globalCatalogMetadataUiNavigationItemModel.ID = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.URL = core.StringPtr("testString")
		globalCatalogMetadataUiNavigationItemModel.Label = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogMetadataUIStringsContent)
		model.Bullets = []partnercentersellv1.CatalogHighlightItem{*catalogHighlightItemModel}
		model.Media = []partnercentersellv1.CatalogProductMediaItem{*catalogProductMediaItemModel}
		model.NavigationItems = []partnercentersellv1.GlobalCatalogMetadataUINavigationItem{*globalCatalogMetadataUiNavigationItemModel}

		assert.Equal(t, result, model)
	}

	catalogHighlightItemModel := make(map[string]interface{})
	catalogHighlightItemModel["description"] = "testString"
	catalogHighlightItemModel["title"] = "testString"

	catalogProductMediaItemModel := make(map[string]interface{})
	catalogProductMediaItemModel["caption"] = "testString"
	catalogProductMediaItemModel["thumbnail"] = "testString"
	catalogProductMediaItemModel["type"] = "image"
	catalogProductMediaItemModel["url"] = "testString"

	globalCatalogMetadataUiNavigationItemModel := make(map[string]interface{})
	globalCatalogMetadataUiNavigationItemModel["id"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["url"] = "testString"
	globalCatalogMetadataUiNavigationItemModel["label"] = "testString"

	model := make(map[string]interface{})
	model["bullets"] = []interface{}{catalogHighlightItemModel}
	model["media"] = []interface{}{catalogProductMediaItemModel}
	model["navigation_items"] = []interface{}{globalCatalogMetadataUiNavigationItemModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUIStringsContent(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToCatalogHighlightItem(t *testing.T) {
	checkResult := func(result *partnercentersellv1.CatalogHighlightItem) {
		model := new(partnercentersellv1.CatalogHighlightItem)
		model.Description = core.StringPtr("testString")
		model.Title = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["description"] = "testString"
	model["title"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToCatalogHighlightItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToCatalogProductMediaItem(t *testing.T) {
	checkResult := func(result *partnercentersellv1.CatalogProductMediaItem) {
		model := new(partnercentersellv1.CatalogProductMediaItem)
		model.Caption = core.StringPtr("testString")
		model.Thumbnail = core.StringPtr("testString")
		model.Type = core.StringPtr("image")
		model.URL = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["caption"] = "testString"
	model["thumbnail"] = "testString"
	model["type"] = "image"
	model["url"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToCatalogProductMediaItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUINavigationItem(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataUINavigationItem) {
		model := new(partnercentersellv1.GlobalCatalogMetadataUINavigationItem)
		model.ID = core.StringPtr("testString")
		model.URL = core.StringPtr("testString")
		model.Label = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "testString"
	model["url"] = "testString"
	model["label"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUINavigationItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUIUrls(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataUIUrls) {
		model := new(partnercentersellv1.GlobalCatalogMetadataUIUrls)
		model.DocURL = core.StringPtr("testString")
		model.ApidocsURL = core.StringPtr("testString")
		model.TermsURL = core.StringPtr("testString")
		model.InstructionsURL = core.StringPtr("testString")
		model.CatalogDetailsURL = core.StringPtr("testString")
		model.CustomCreatePageURL = core.StringPtr("testString")
		model.Dashboard = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["doc_url"] = "testString"
	model["apidocs_url"] = "testString"
	model["terms_url"] = "testString"
	model["instructions_url"] = "testString"
	model["catalog_details_url"] = "testString"
	model["custom_create_page_url"] = "testString"
	model["dashboard"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataUIUrls(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataServicePrototypePatch(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogProductMetadataServicePrototypePatch) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

		globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
		globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

		globalCatalogMetadataServiceCustomParametersOptionsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
		globalCatalogMetadataServiceCustomParametersOptionsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersOptionsModel.Value = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersOptionsModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

		globalCatalogMetadataServiceCustomParametersAssociationsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

		globalCatalogMetadataServiceCustomParametersModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
		globalCatalogMetadataServiceCustomParametersModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Name = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Type = core.StringPtr("text")
		globalCatalogMetadataServiceCustomParametersModel.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
		globalCatalogMetadataServiceCustomParametersModel.Value = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersModel.Layout = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Associations = globalCatalogMetadataServiceCustomParametersAssociationsModel
		globalCatalogMetadataServiceCustomParametersModel.ValidationURL = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.OptionsURL = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Invalidmessage = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Description = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Required = core.BoolPtr(true)
		globalCatalogMetadataServiceCustomParametersModel.Pattern = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Placeholder = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersModel.Readonly = core.BoolPtr(true)
		globalCatalogMetadataServiceCustomParametersModel.Hidden = core.BoolPtr(true)
		globalCatalogMetadataServiceCustomParametersModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

		model := new(partnercentersellv1.GlobalCatalogProductMetadataServicePrototypePatch)
		model.RcProvisionable = core.BoolPtr(true)
		model.IamCompatible = core.BoolPtr(true)
		model.ServiceKeySupported = core.BoolPtr(true)
		model.UniqueApiKey = core.BoolPtr(true)
		model.AsyncProvisioningSupported = core.BoolPtr(true)
		model.AsyncUnprovisioningSupported = core.BoolPtr(true)
		model.CustomCreatePageHybridEnabled = core.BoolPtr(true)
		model.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters{*globalCatalogMetadataServiceCustomParametersModel}

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

	globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

	globalCatalogMetadataServiceCustomParametersOptionsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersOptionsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersOptionsModel["value"] = "testString"
	globalCatalogMetadataServiceCustomParametersOptionsModel["i18n"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []interface{}{"testString"}

	globalCatalogMetadataServiceCustomParametersAssociationsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsModel["plan"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel["parameters"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel["location"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

	globalCatalogMetadataServiceCustomParametersModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["name"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["type"] = "text"
	globalCatalogMetadataServiceCustomParametersModel["options"] = []interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
	globalCatalogMetadataServiceCustomParametersModel["value"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersModel["layout"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["associations"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsModel}
	globalCatalogMetadataServiceCustomParametersModel["validation_url"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["options_url"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["invalidmessage"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["description"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["required"] = true
	globalCatalogMetadataServiceCustomParametersModel["pattern"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["placeholder"] = "testString"
	globalCatalogMetadataServiceCustomParametersModel["readonly"] = true
	globalCatalogMetadataServiceCustomParametersModel["hidden"] = true
	globalCatalogMetadataServiceCustomParametersModel["i18n"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

	model := make(map[string]interface{})
	model["rc_provisionable"] = true
	model["iam_compatible"] = true
	model["service_key_supported"] = true
	model["unique_api_key"] = true
	model["async_provisioning_supported"] = true
	model["async_unprovisioning_supported"] = true
	model["custom_create_page_hybrid_enabled"] = true
	model["parameters"] = []interface{}{globalCatalogMetadataServiceCustomParametersModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataServicePrototypePatch(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParameters(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

		globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
		globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

		globalCatalogMetadataServiceCustomParametersOptionsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
		globalCatalogMetadataServiceCustomParametersOptionsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersOptionsModel.Value = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersOptionsModel.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

		globalCatalogMetadataServiceCustomParametersAssociationsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		globalCatalogMetadataServiceCustomParametersAssociationsModel.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParameters)
		model.Displayname = core.StringPtr("testString")
		model.Name = core.StringPtr("testString")
		model.Type = core.StringPtr("text")
		model.Options = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions{*globalCatalogMetadataServiceCustomParametersOptionsModel}
		model.Value = []string{"testString"}
		model.Layout = core.StringPtr("testString")
		model.Associations = globalCatalogMetadataServiceCustomParametersAssociationsModel
		model.ValidationURL = core.StringPtr("testString")
		model.OptionsURL = core.StringPtr("testString")
		model.Invalidmessage = core.StringPtr("testString")
		model.Description = core.StringPtr("testString")
		model.Required = core.BoolPtr(true)
		model.Pattern = core.StringPtr("testString")
		model.Placeholder = core.StringPtr("testString")
		model.Readonly = core.BoolPtr(true)
		model.Hidden = core.BoolPtr(true)
		model.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

	globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

	globalCatalogMetadataServiceCustomParametersOptionsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersOptionsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersOptionsModel["value"] = "testString"
	globalCatalogMetadataServiceCustomParametersOptionsModel["i18n"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []interface{}{"testString"}

	globalCatalogMetadataServiceCustomParametersAssociationsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsModel["plan"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel["parameters"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	globalCatalogMetadataServiceCustomParametersAssociationsModel["location"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

	model := make(map[string]interface{})
	model["displayname"] = "testString"
	model["name"] = "testString"
	model["type"] = "text"
	model["options"] = []interface{}{globalCatalogMetadataServiceCustomParametersOptionsModel}
	model["value"] = []interface{}{"testString"}
	model["layout"] = "testString"
	model["associations"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsModel}
	model["validation_url"] = "testString"
	model["options_url"] = "testString"
	model["invalidmessage"] = "testString"
	model["description"] = "testString"
	model["required"] = true
	model["pattern"] = "testString"
	model["placeholder"] = "testString"
	model["readonly"] = true
	model["hidden"] = true
	model["i18n"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParametersOptions(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

		globalCatalogMetadataServiceCustomParametersI18nModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
		globalCatalogMetadataServiceCustomParametersI18nModel.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		globalCatalogMetadataServiceCustomParametersI18nModel.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersOptions)
		model.Displayname = core.StringPtr("testString")
		model.Value = core.StringPtr("testString")
		model.I18n = globalCatalogMetadataServiceCustomParametersI18nModel

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

	globalCatalogMetadataServiceCustomParametersI18nModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nModel["en"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["de"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["es"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["fr"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["it"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ja"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["ko"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["pt_br"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_tw"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	globalCatalogMetadataServiceCustomParametersI18nModel["zh_cn"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

	model := make(map[string]interface{})
	model["displayname"] = "testString"
	model["value"] = "testString"
	model["i18n"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParametersOptions(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParametersI18n(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n) {
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Displayname = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersI18nFieldsModel.Description = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18n)
		model.En = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.De = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.Es = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.Fr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.It = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.Ja = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.Ko = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.PtBr = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.ZhTw = globalCatalogMetadataServiceCustomParametersI18nFieldsModel
		model.ZhCn = globalCatalogMetadataServiceCustomParametersI18nFieldsModel

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersI18nFieldsModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["displayname"] = "testString"
	globalCatalogMetadataServiceCustomParametersI18nFieldsModel["description"] = "testString"

	model := make(map[string]interface{})
	model["en"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["de"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["es"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["fr"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["it"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["ja"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["ko"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["pt_br"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["zh_tw"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}
	model["zh_cn"] = []interface{}{globalCatalogMetadataServiceCustomParametersI18nFieldsModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParametersI18n(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParametersI18nFields(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields) {
		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersI18nFields)
		model.Displayname = core.StringPtr("testString")
		model.Description = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["displayname"] = "testString"
	model["description"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParametersI18nFields(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParametersAssociations(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations) {
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsPlanModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.Name = core.StringPtr("testString")
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.ShowFor = []string{"testString"}
		globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel.OptionsRefresh = core.BoolPtr(true)

		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
		globalCatalogMetadataServiceCustomParametersAssociationsLocationModel.ShowFor = []string{"testString"}

		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociations)
		model.Plan = globalCatalogMetadataServiceCustomParametersAssociationsPlanModel
		model.Parameters = []partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem{*globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
		model.Location = globalCatalogMetadataServiceCustomParametersAssociationsLocationModel

		assert.Equal(t, result, model)
	}

	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsPlanModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["name"] = "testString"
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["show_for"] = []interface{}{"testString"}
	globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel["options_refresh"] = true

	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel := make(map[string]interface{})
	globalCatalogMetadataServiceCustomParametersAssociationsLocationModel["show_for"] = []interface{}{"testString"}

	model := make(map[string]interface{})
	model["plan"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsPlanModel}
	model["parameters"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsParametersItemModel}
	model["location"] = []interface{}{globalCatalogMetadataServiceCustomParametersAssociationsLocationModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParametersAssociations(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParametersAssociationsPlan(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan) {
		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsPlan)
		model.ShowFor = []string{"testString"}
		model.OptionsRefresh = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["show_for"] = []interface{}{"testString"}
	model["options_refresh"] = true

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParametersAssociationsPlan(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem) {
		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem)
		model.Name = core.StringPtr("testString")
		model.ShowFor = []string{"testString"}
		model.OptionsRefresh = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["show_for"] = []interface{}{"testString"}
	model["options_refresh"] = true

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParametersAssociationsParametersItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParametersAssociationsLocation(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation) {
		model := new(partnercentersellv1.GlobalCatalogMetadataServiceCustomParametersAssociationsLocation)
		model.ShowFor = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["show_for"] = []interface{}{"testString"}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogMetadataServiceCustomParametersAssociationsLocation(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOther(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogProductMetadataOther) {
		supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
		supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
		supportTimeIntervalModel.Type = core.StringPtr("testString")

		supportEscalationModel := new(partnercentersellv1.SupportEscalation)
		supportEscalationModel.Contact = core.StringPtr("testString")
		supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
		supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

		supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
		supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
		supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
		supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

		supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
		supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
		supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

		supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
		supportDetailsItemModel.Type = core.StringPtr("support_site")
		supportDetailsItemModel.Contact = core.StringPtr("testString")
		supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
		supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

		globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
		globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
		globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
		globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
		globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

		globalCatalogProductMetadataOtherPcModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
		globalCatalogProductMetadataOtherPcModel.Support = globalCatalogProductMetadataOtherPcSupportModel

		globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
		globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
		globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

		globalCatalogProductMetadataOtherCompositeModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
		globalCatalogProductMetadataOtherCompositeModel.CompositeKind = core.StringPtr("service")
		globalCatalogProductMetadataOtherCompositeModel.CompositeTag = core.StringPtr("testString")
		globalCatalogProductMetadataOtherCompositeModel.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

		model := new(partnercentersellv1.GlobalCatalogProductMetadataOther)
		model.PC = globalCatalogProductMetadataOtherPcModel
		model.Composite = globalCatalogProductMetadataOtherCompositeModel

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := make(map[string]interface{})
	supportTimeIntervalModel["value"] = float64(72.5)
	supportTimeIntervalModel["type"] = "testString"

	supportEscalationModel := make(map[string]interface{})
	supportEscalationModel["contact"] = "testString"
	supportEscalationModel["escalation_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportEscalationModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}

	supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
	supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
	supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
	supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

	supportDetailsItemAvailabilityModel := make(map[string]interface{})
	supportDetailsItemAvailabilityModel["times"] = []interface{}{supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel["timezone"] = "testString"
	supportDetailsItemAvailabilityModel["always_available"] = true

	supportDetailsItemModel := make(map[string]interface{})
	supportDetailsItemModel["type"] = "support_site"
	supportDetailsItemModel["contact"] = "testString"
	supportDetailsItemModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportDetailsItemModel["availability"] = []interface{}{supportDetailsItemAvailabilityModel}

	globalCatalogProductMetadataOtherPcSupportModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherPcSupportModel["url"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["status_url"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["locations"] = []interface{}{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel["languages"] = []interface{}{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel["process"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["process_i18n"] = map[string]interface{}{"key1": "testString"}
	globalCatalogProductMetadataOtherPcSupportModel["support_type"] = "community"
	globalCatalogProductMetadataOtherPcSupportModel["support_escalation"] = []interface{}{supportEscalationModel}
	globalCatalogProductMetadataOtherPcSupportModel["support_details"] = []interface{}{supportDetailsItemModel}

	globalCatalogProductMetadataOtherPcModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherPcModel["support"] = []interface{}{globalCatalogProductMetadataOtherPcSupportModel}

	globalCatalogProductMetadataOtherCompositeChildModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherCompositeChildModel["kind"] = "service"
	globalCatalogProductMetadataOtherCompositeChildModel["name"] = "testString"

	globalCatalogProductMetadataOtherCompositeModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherCompositeModel["composite_kind"] = "service"
	globalCatalogProductMetadataOtherCompositeModel["composite_tag"] = "testString"
	globalCatalogProductMetadataOtherCompositeModel["children"] = []interface{}{globalCatalogProductMetadataOtherCompositeChildModel}

	model := make(map[string]interface{})
	model["pc"] = []interface{}{globalCatalogProductMetadataOtherPcModel}
	model["composite"] = []interface{}{globalCatalogProductMetadataOtherCompositeModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOther(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOtherPC(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogProductMetadataOtherPC) {
		supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
		supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
		supportTimeIntervalModel.Type = core.StringPtr("testString")

		supportEscalationModel := new(partnercentersellv1.SupportEscalation)
		supportEscalationModel.Contact = core.StringPtr("testString")
		supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
		supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

		supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
		supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
		supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
		supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

		supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
		supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
		supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

		supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
		supportDetailsItemModel.Type = core.StringPtr("support_site")
		supportDetailsItemModel.Contact = core.StringPtr("testString")
		supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
		supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

		globalCatalogProductMetadataOtherPcSupportModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
		globalCatalogProductMetadataOtherPcSupportModel.URL = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.StatusURL = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.Locations = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel.Languages = []string{"testString"}
		globalCatalogProductMetadataOtherPcSupportModel.Process = core.StringPtr("testString")
		globalCatalogProductMetadataOtherPcSupportModel.ProcessI18n = map[string]string{"key1": "testString"}
		globalCatalogProductMetadataOtherPcSupportModel.SupportType = core.StringPtr("community")
		globalCatalogProductMetadataOtherPcSupportModel.SupportEscalation = supportEscalationModel
		globalCatalogProductMetadataOtherPcSupportModel.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

		model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPC)
		model.Support = globalCatalogProductMetadataOtherPcSupportModel

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := make(map[string]interface{})
	supportTimeIntervalModel["value"] = float64(72.5)
	supportTimeIntervalModel["type"] = "testString"

	supportEscalationModel := make(map[string]interface{})
	supportEscalationModel["contact"] = "testString"
	supportEscalationModel["escalation_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportEscalationModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}

	supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
	supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
	supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
	supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

	supportDetailsItemAvailabilityModel := make(map[string]interface{})
	supportDetailsItemAvailabilityModel["times"] = []interface{}{supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel["timezone"] = "testString"
	supportDetailsItemAvailabilityModel["always_available"] = true

	supportDetailsItemModel := make(map[string]interface{})
	supportDetailsItemModel["type"] = "support_site"
	supportDetailsItemModel["contact"] = "testString"
	supportDetailsItemModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportDetailsItemModel["availability"] = []interface{}{supportDetailsItemAvailabilityModel}

	globalCatalogProductMetadataOtherPcSupportModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherPcSupportModel["url"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["status_url"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["locations"] = []interface{}{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel["languages"] = []interface{}{"testString"}
	globalCatalogProductMetadataOtherPcSupportModel["process"] = "testString"
	globalCatalogProductMetadataOtherPcSupportModel["process_i18n"] = map[string]interface{}{"key1": "testString"}
	globalCatalogProductMetadataOtherPcSupportModel["support_type"] = "community"
	globalCatalogProductMetadataOtherPcSupportModel["support_escalation"] = []interface{}{supportEscalationModel}
	globalCatalogProductMetadataOtherPcSupportModel["support_details"] = []interface{}{supportDetailsItemModel}

	model := make(map[string]interface{})
	model["support"] = []interface{}{globalCatalogProductMetadataOtherPcSupportModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOtherPC(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOtherPCSupport(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport) {
		supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
		supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
		supportTimeIntervalModel.Type = core.StringPtr("testString")

		supportEscalationModel := new(partnercentersellv1.SupportEscalation)
		supportEscalationModel.Contact = core.StringPtr("testString")
		supportEscalationModel.EscalationWaitTime = supportTimeIntervalModel
		supportEscalationModel.ResponseWaitTime = supportTimeIntervalModel

		supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
		supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
		supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
		supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

		supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
		supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
		supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

		supportDetailsItemModel := new(partnercentersellv1.SupportDetailsItem)
		supportDetailsItemModel.Type = core.StringPtr("support_site")
		supportDetailsItemModel.Contact = core.StringPtr("testString")
		supportDetailsItemModel.ResponseWaitTime = supportTimeIntervalModel
		supportDetailsItemModel.Availability = supportDetailsItemAvailabilityModel

		model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherPCSupport)
		model.URL = core.StringPtr("testString")
		model.StatusURL = core.StringPtr("testString")
		model.Locations = []string{"testString"}
		model.Languages = []string{"testString"}
		model.Process = core.StringPtr("testString")
		model.ProcessI18n = map[string]string{"key1": "testString"}
		model.SupportType = core.StringPtr("community")
		model.SupportEscalation = supportEscalationModel
		model.SupportDetails = []partnercentersellv1.SupportDetailsItem{*supportDetailsItemModel}

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := make(map[string]interface{})
	supportTimeIntervalModel["value"] = float64(72.5)
	supportTimeIntervalModel["type"] = "testString"

	supportEscalationModel := make(map[string]interface{})
	supportEscalationModel["contact"] = "testString"
	supportEscalationModel["escalation_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportEscalationModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}

	supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
	supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
	supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
	supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

	supportDetailsItemAvailabilityModel := make(map[string]interface{})
	supportDetailsItemAvailabilityModel["times"] = []interface{}{supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel["timezone"] = "testString"
	supportDetailsItemAvailabilityModel["always_available"] = true

	supportDetailsItemModel := make(map[string]interface{})
	supportDetailsItemModel["type"] = "support_site"
	supportDetailsItemModel["contact"] = "testString"
	supportDetailsItemModel["response_wait_time"] = []interface{}{supportTimeIntervalModel}
	supportDetailsItemModel["availability"] = []interface{}{supportDetailsItemAvailabilityModel}

	model := make(map[string]interface{})
	model["url"] = "testString"
	model["status_url"] = "testString"
	model["locations"] = []interface{}{"testString"}
	model["languages"] = []interface{}{"testString"}
	model["process"] = "testString"
	model["process_i18n"] = map[string]interface{}{"key1": "testString"}
	model["support_type"] = "community"
	model["support_escalation"] = []interface{}{supportEscalationModel}
	model["support_details"] = []interface{}{supportDetailsItemModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOtherPCSupport(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToSupportEscalation(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportEscalation) {
		supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
		supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
		supportTimeIntervalModel.Type = core.StringPtr("testString")

		model := new(partnercentersellv1.SupportEscalation)
		model.Contact = core.StringPtr("testString")
		model.EscalationWaitTime = supportTimeIntervalModel
		model.ResponseWaitTime = supportTimeIntervalModel

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := make(map[string]interface{})
	supportTimeIntervalModel["value"] = float64(72.5)
	supportTimeIntervalModel["type"] = "testString"

	model := make(map[string]interface{})
	model["contact"] = "testString"
	model["escalation_wait_time"] = []interface{}{supportTimeIntervalModel}
	model["response_wait_time"] = []interface{}{supportTimeIntervalModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToSupportEscalation(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToSupportTimeInterval(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportTimeInterval) {
		model := new(partnercentersellv1.SupportTimeInterval)
		model.Value = core.Float64Ptr(float64(72.5))
		model.Type = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["value"] = float64(72.5)
	model["type"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToSupportTimeInterval(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToSupportDetailsItem(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportDetailsItem) {
		supportTimeIntervalModel := new(partnercentersellv1.SupportTimeInterval)
		supportTimeIntervalModel.Value = core.Float64Ptr(float64(72.5))
		supportTimeIntervalModel.Type = core.StringPtr("testString")

		supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
		supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
		supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
		supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

		supportDetailsItemAvailabilityModel := new(partnercentersellv1.SupportDetailsItemAvailability)
		supportDetailsItemAvailabilityModel.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
		supportDetailsItemAvailabilityModel.Timezone = core.StringPtr("testString")
		supportDetailsItemAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)

		model := new(partnercentersellv1.SupportDetailsItem)
		model.Type = core.StringPtr("support_site")
		model.Contact = core.StringPtr("testString")
		model.ResponseWaitTime = supportTimeIntervalModel
		model.Availability = supportDetailsItemAvailabilityModel

		assert.Equal(t, result, model)
	}

	supportTimeIntervalModel := make(map[string]interface{})
	supportTimeIntervalModel["value"] = float64(72.5)
	supportTimeIntervalModel["type"] = "testString"

	supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
	supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
	supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
	supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

	supportDetailsItemAvailabilityModel := make(map[string]interface{})
	supportDetailsItemAvailabilityModel["times"] = []interface{}{supportDetailsItemAvailabilityTimeModel}
	supportDetailsItemAvailabilityModel["timezone"] = "testString"
	supportDetailsItemAvailabilityModel["always_available"] = true

	model := make(map[string]interface{})
	model["type"] = "support_site"
	model["contact"] = "testString"
	model["response_wait_time"] = []interface{}{supportTimeIntervalModel}
	model["availability"] = []interface{}{supportDetailsItemAvailabilityModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToSupportDetailsItem(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToSupportDetailsItemAvailability(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportDetailsItemAvailability) {
		supportDetailsItemAvailabilityTimeModel := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
		supportDetailsItemAvailabilityTimeModel.Day = core.Float64Ptr(float64(72.5))
		supportDetailsItemAvailabilityTimeModel.StartTime = core.StringPtr("testString")
		supportDetailsItemAvailabilityTimeModel.EndTime = core.StringPtr("testString")

		model := new(partnercentersellv1.SupportDetailsItemAvailability)
		model.Times = []partnercentersellv1.SupportDetailsItemAvailabilityTime{*supportDetailsItemAvailabilityTimeModel}
		model.Timezone = core.StringPtr("testString")
		model.AlwaysAvailable = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	supportDetailsItemAvailabilityTimeModel := make(map[string]interface{})
	supportDetailsItemAvailabilityTimeModel["day"] = float64(72.5)
	supportDetailsItemAvailabilityTimeModel["start_time"] = "testString"
	supportDetailsItemAvailabilityTimeModel["end_time"] = "testString"

	model := make(map[string]interface{})
	model["times"] = []interface{}{supportDetailsItemAvailabilityTimeModel}
	model["timezone"] = "testString"
	model["always_available"] = true

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToSupportDetailsItemAvailability(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToSupportDetailsItemAvailabilityTime(t *testing.T) {
	checkResult := func(result *partnercentersellv1.SupportDetailsItemAvailabilityTime) {
		model := new(partnercentersellv1.SupportDetailsItemAvailabilityTime)
		model.Day = core.Float64Ptr(float64(72.5))
		model.StartTime = core.StringPtr("testString")
		model.EndTime = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["day"] = float64(72.5)
	model["start_time"] = "testString"
	model["end_time"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToSupportDetailsItemAvailabilityTime(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOtherComposite(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogProductMetadataOtherComposite) {
		globalCatalogProductMetadataOtherCompositeChildModel := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
		globalCatalogProductMetadataOtherCompositeChildModel.Kind = core.StringPtr("service")
		globalCatalogProductMetadataOtherCompositeChildModel.Name = core.StringPtr("testString")

		model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherComposite)
		model.CompositeKind = core.StringPtr("service")
		model.CompositeTag = core.StringPtr("testString")
		model.Children = []partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild{*globalCatalogProductMetadataOtherCompositeChildModel}

		assert.Equal(t, result, model)
	}

	globalCatalogProductMetadataOtherCompositeChildModel := make(map[string]interface{})
	globalCatalogProductMetadataOtherCompositeChildModel["kind"] = "service"
	globalCatalogProductMetadataOtherCompositeChildModel["name"] = "testString"

	model := make(map[string]interface{})
	model["composite_kind"] = "service"
	model["composite_tag"] = "testString"
	model["children"] = []interface{}{globalCatalogProductMetadataOtherCompositeChildModel}

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOtherComposite(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOtherCompositeChild(t *testing.T) {
	checkResult := func(result *partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild) {
		model := new(partnercentersellv1.GlobalCatalogProductMetadataOtherCompositeChild)
		model.Kind = core.StringPtr("service")
		model.Name = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["kind"] = "service"
	model["name"] = "testString"

	result, err := partnercentersell.ResourceIbmOnboardingCatalogProductMapToGlobalCatalogProductMetadataOtherCompositeChild(model)
	assert.Nil(t, err)
	checkResult(result)
}
