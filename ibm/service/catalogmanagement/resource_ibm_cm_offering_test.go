// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func TestAccIBMCmOfferingBasic(t *testing.T) {
	var conf catalogmanagementv1.Offering

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmOfferingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmOfferingExists("ibm_cm_offering.cm_offering", conf),
				),
			},
		},
	})
}

func TestAccIBMCmOfferingAllArgs(t *testing.T) {
	var conf catalogmanagementv1.Offering
	url := fmt.Sprintf("tf_url_%d", acctest.RandIntRange(10, 100))
	crn := fmt.Sprintf("tf_crn_%d", acctest.RandIntRange(10, 100))
	label := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	offeringIconURL := fmt.Sprintf("tf_offering_icon_url_%d", acctest.RandIntRange(10, 100))
	offeringDocsURL := fmt.Sprintf("tf_offering_docs_url_%d", acctest.RandIntRange(10, 100))
	offeringSupportURL := fmt.Sprintf("tf_offering_support_url_%d", acctest.RandIntRange(10, 100))
	shortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	longDescription := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	pcManaged := "true"
	publishApproved := "true"
	shareWithAll := "true"
	shareWithIBM := "true"
	shareEnabled := "false"
	permitRequestIBMPublicPublish := "false"
	ibmPublishApproved := "true"
	publicPublishApproved := "false"
	publicOriginalCRN := fmt.Sprintf("tf_public_original_crn_%d", acctest.RandIntRange(10, 100))
	publishPublicCRN := fmt.Sprintf("tf_publish_public_crn_%d", acctest.RandIntRange(10, 100))
	portalApprovalRecord := fmt.Sprintf("tf_portal_approval_record_%d", acctest.RandIntRange(10, 100))
	portalUIURL := fmt.Sprintf("tf_portal_ui_url_%d", acctest.RandIntRange(10, 100))
	catalogName := fmt.Sprintf("tf_catalog_name_%d", acctest.RandIntRange(10, 100))
	disclaimer := fmt.Sprintf("tf_disclaimer_%d", acctest.RandIntRange(10, 100))
	hidden := "true"
	provider := fmt.Sprintf("tf_provider_%d", acctest.RandIntRange(10, 100))
	productKind := fmt.Sprintf("tf_product_kind_%d", acctest.RandIntRange(10, 100))
	urlUpdate := fmt.Sprintf("tf_url_%d", acctest.RandIntRange(10, 100))
	crnUpdate := fmt.Sprintf("tf_crn_%d", acctest.RandIntRange(10, 100))
	labelUpdate := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	offeringIconURLUpdate := fmt.Sprintf("tf_offering_icon_url_%d", acctest.RandIntRange(10, 100))
	offeringDocsURLUpdate := fmt.Sprintf("tf_offering_docs_url_%d", acctest.RandIntRange(10, 100))
	offeringSupportURLUpdate := fmt.Sprintf("tf_offering_support_url_%d", acctest.RandIntRange(10, 100))
	shortDescriptionUpdate := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	longDescriptionUpdate := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	pcManagedUpdate := "false"
	publishApprovedUpdate := "false"
	shareWithAllUpdate := "false"
	shareWithIBMUpdate := "false"
	shareEnabledUpdate := "true"
	permitRequestIBMPublicPublishUpdate := "true"
	ibmPublishApprovedUpdate := "false"
	publicPublishApprovedUpdate := "true"
	publicOriginalCRNUpdate := fmt.Sprintf("tf_public_original_crn_%d", acctest.RandIntRange(10, 100))
	publishPublicCRNUpdate := fmt.Sprintf("tf_publish_public_crn_%d", acctest.RandIntRange(10, 100))
	portalApprovalRecordUpdate := fmt.Sprintf("tf_portal_approval_record_%d", acctest.RandIntRange(10, 100))
	portalUIURLUpdate := fmt.Sprintf("tf_portal_ui_url_%d", acctest.RandIntRange(10, 100))
	catalogNameUpdate := fmt.Sprintf("tf_catalog_name_%d", acctest.RandIntRange(10, 100))
	disclaimerUpdate := fmt.Sprintf("tf_disclaimer_%d", acctest.RandIntRange(10, 100))
	hiddenUpdate := "false"
	providerUpdate := fmt.Sprintf("tf_provider_%d", acctest.RandIntRange(10, 100))
	productKindUpdate := fmt.Sprintf("tf_product_kind_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmOfferingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingConfig(url, crn, label, name, offeringIconURL, offeringDocsURL, offeringSupportURL, shortDescription, longDescription, pcManaged, publishApproved, shareWithAll, shareWithIBM, shareEnabled, permitRequestIBMPublicPublish, ibmPublishApproved, publicPublishApproved, publicOriginalCRN, publishPublicCRN, portalApprovalRecord, portalUIURL, catalogName, disclaimer, hidden, provider, productKind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmOfferingExists("ibm_cm_offering.cm_offering", conf),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "url", url),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "crn", crn),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "label", label),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "name", name),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "offering_icon_url", offeringIconURL),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "offering_docs_url", offeringDocsURL),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "offering_support_url", offeringSupportURL),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "short_description", shortDescription),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "long_description", longDescription),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "pc_managed", pcManaged),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "publish_approved", publishApproved),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "share_with_all", shareWithAll),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "share_with_ibm", shareWithIBM),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "share_enabled", shareEnabled),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "permit_request_ibm_public_publish", permitRequestIBMPublicPublish),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "ibm_publish_approved", ibmPublishApproved),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "public_publish_approved", publicPublishApproved),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "public_original_crn", publicOriginalCRN),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "publish_public_crn", publishPublicCRN),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "portal_approval_record", portalApprovalRecord),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "portal_ui_url", portalUIURL),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "catalog_name", catalogName),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "disclaimer", disclaimer),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "hidden", hidden),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "provider", provider),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "product_kind", productKind),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingConfig(urlUpdate, crnUpdate, labelUpdate, nameUpdate, offeringIconURLUpdate, offeringDocsURLUpdate, offeringSupportURLUpdate, shortDescriptionUpdate, longDescriptionUpdate, pcManagedUpdate, publishApprovedUpdate, shareWithAllUpdate, shareWithIBMUpdate, shareEnabledUpdate, permitRequestIBMPublicPublishUpdate, ibmPublishApprovedUpdate, publicPublishApprovedUpdate, publicOriginalCRNUpdate, publishPublicCRNUpdate, portalApprovalRecordUpdate, portalUIURLUpdate, catalogNameUpdate, disclaimerUpdate, hiddenUpdate, providerUpdate, productKindUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "url", urlUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "crn", crnUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "label", labelUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "offering_icon_url", offeringIconURLUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "offering_docs_url", offeringDocsURLUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "offering_support_url", offeringSupportURLUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "short_description", shortDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "long_description", longDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "pc_managed", pcManagedUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "publish_approved", publishApprovedUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "share_with_all", shareWithAllUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "share_with_ibm", shareWithIBMUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "share_enabled", shareEnabledUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "permit_request_ibm_public_publish", permitRequestIBMPublicPublishUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "ibm_publish_approved", ibmPublishApprovedUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "public_publish_approved", publicPublishApprovedUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "public_original_crn", publicOriginalCRNUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "publish_public_crn", publishPublicCRNUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "portal_approval_record", portalApprovalRecordUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "portal_ui_url", portalUIURLUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "catalog_name", catalogNameUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "disclaimer", disclaimerUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "hidden", hiddenUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "provider", providerUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "product_kind", productKindUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cm_offering.cm_offering",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCmOfferingConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
		}

		resource "ibm_cm_offering" "cm_offering" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
		}
	`)
}

func testAccCheckIBMCmOfferingConfig(url string, crn string, label string, name string, offeringIconURL string, offeringDocsURL string, offeringSupportURL string, shortDescription string, longDescription string, pcManaged string, publishApproved string, shareWithAll string, shareWithIBM string, shareEnabled string, permitRequestIBMPublicPublish string, ibmPublishApproved string, publicPublishApproved string, publicOriginalCRN string, publishPublicCRN string, portalApprovalRecord string, portalUIURL string, catalogName string, disclaimer string, hidden string, provider string, productKind string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
		}

		resource "ibm_cm_offering" "cm_offering" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
			url = "%s"
			crn = "%s"
			label = "%s"
			label_i18n = "FIXME"
			name = "%s"
			offering_icon_url = "%s"
			offering_docs_url = "%s"
			offering_support_url = "%s"
			tags = "FIXME"
			keywords = "FIXME"
			rating {
				one_star_count = 1
				two_star_count = 1
				three_star_count = 1
				four_star_count = 1
			}
			created = "2004-10-28T04:39:00.000Z"
			updated = "2004-10-28T04:39:00.000Z"
			short_description = "%s"
			short_description_i18n = "FIXME"
			long_description = "%s"
			long_description_i18n = "FIXME"
			features {
				title = "title"
				title_i18n = { "key": "inner" }
				description = "description"
				description_i18n = { "key": "inner" }
			}
			kinds {
				id = "id"
				format_kind = "format_kind"
				install_kind = "install_kind"
				target_kind = "target_kind"
				metadata = { "key": null }
				tags = [ "tags" ]
				additional_features {
					title = "title"
					title_i18n = { "key": "inner" }
					description = "description"
					description_i18n = { "key": "inner" }
				}
				created = "2021-01-31T09:44:12Z"
				updated = "2021-01-31T09:44:12Z"
				versions {
					id = "id"
					rev = "rev"
					crn = "crn"
					version = "version"
					flavor {
						name = "name"
						label = "label"
						label_i18n = { "key": "inner" }
						index = 1
					}
					sha = "sha"
					created = "2021-01-31T09:44:12Z"
					updated = "2021-01-31T09:44:12Z"
					offering_id = "offering_id"
					catalog_id = "catalog_id"
					kind_id = "kind_id"
					tags = [ "tags" ]
					repo_url = "repo_url"
					source_url = "source_url"
					tgz_url = "tgz_url"
					configuration {
						key = "key"
						type = "type"
						display_name = "display_name"
						value_constraint = "value_constraint"
						description = "description"
						required = true
						options = [ null ]
						hidden = true
						custom_config {
							type = "type"
							grouping = "grouping"
							original_grouping = "original_grouping"
							grouping_index = 1
							config_constraints = { "key": null }
							associations {
								parameters {
									name = "name"
									options_refresh = true
								}
							}
						}
						type_metadata = "type_metadata"
					}
					outputs {
						key = "key"
						description = "description"
					}
					iam_permissions {
						service_name = "service_name"
						role_crns = [ "role_crns" ]
						resources {
							name = "name"
							description = "description"
							role_crns = [ "role_crns" ]
						}
					}
					metadata = { "key": null }
					validation {
						validated = "2021-01-31T09:44:12Z"
						requested = "2021-01-31T09:44:12Z"
						state = "state"
						last_operation = "last_operation"
						target = { "key": null }
						message = "message"
					}
					required_resources {
						type = "mem"
					}
					single_instance = true
					install {
						instructions = "instructions"
						instructions_i18n = { "key": "inner" }
						script = "script"
						script_permission = "script_permission"
						delete_script = "delete_script"
						scope = "scope"
					}
					pre_install {
						instructions = "instructions"
						instructions_i18n = { "key": "inner" }
						script = "script"
						script_permission = "script_permission"
						delete_script = "delete_script"
						scope = "scope"
					}
					entitlement {
						provider_name = "provider_name"
						provider_id = "provider_id"
						product_id = "product_id"
						part_numbers = [ "part_numbers" ]
						image_repo_name = "image_repo_name"
					}
					licenses {
						id = "id"
						name = "name"
						type = "type"
						url = "url"
						description = "description"
					}
					image_manifest_url = "image_manifest_url"
					deprecated = true
					package_version = "package_version"
					state {
						current = "current"
						current_entered = "2021-01-31T09:44:12Z"
						pending = "pending"
						pending_requested = "2021-01-31T09:44:12Z"
						previous = "previous"
					}
					version_locator = "version_locator"
					long_description = "long_description"
					long_description_i18n = { "key": "inner" }
					whitelisted_accounts = [ "whitelisted_accounts" ]
					image_pull_key_name = "image_pull_key_name"
					deprecate_pending {
						deprecate_date = "2021-01-31T09:44:12Z"
						deprecate_state = "deprecate_state"
						description = "description"
					}
					solution_info {
						architecture_diagrams {
							diagram {
								url = "url"
								api_url = "api_url"
								url_proxy {
									url = "url"
									sha = "sha"
								}
								caption = "caption"
								caption_i18n = { "key": "inner" }
								type = "type"
								thumbnail_url = "thumbnail_url"
							}
							description = "description"
							description_i18n = { "key": "inner" }
						}
						features {
							title = "title"
							title_i18n = { "key": "inner" }
							description = "description"
							description_i18n = { "key": "inner" }
						}
						cost_estimate {
							version = "version"
							currency = "currency"
							projects {
								name = "name"
								metadata = { "key": null }
								past_breakdown {
									total_hourly_cost = "total_hourly_cost"
									total_monthly_c_ost = "total_monthly_c_ost"
									resources {
										name = "name"
										metadata = { "key": null }
										hourly_cost = "hourly_cost"
										monthly_cost = "monthly_cost"
										cost_components {
											name = "name"
											unit = "unit"
											hourly_quantity = "hourly_quantity"
											monthly_quantity = "monthly_quantity"
											price = "price"
											hourly_cost = "hourly_cost"
											monthly_cost = "monthly_cost"
										}
									}
								}
								breakdown {
									total_hourly_cost = "total_hourly_cost"
									total_monthly_c_ost = "total_monthly_c_ost"
									resources {
										name = "name"
										metadata = { "key": null }
										hourly_cost = "hourly_cost"
										monthly_cost = "monthly_cost"
										cost_components {
											name = "name"
											unit = "unit"
											hourly_quantity = "hourly_quantity"
											monthly_quantity = "monthly_quantity"
											price = "price"
											hourly_cost = "hourly_cost"
											monthly_cost = "monthly_cost"
										}
									}
								}
								diff {
									total_hourly_cost = "total_hourly_cost"
									total_monthly_c_ost = "total_monthly_c_ost"
									resources {
										name = "name"
										metadata = { "key": null }
										hourly_cost = "hourly_cost"
										monthly_cost = "monthly_cost"
										cost_components {
											name = "name"
											unit = "unit"
											hourly_quantity = "hourly_quantity"
											monthly_quantity = "monthly_quantity"
											price = "price"
											hourly_cost = "hourly_cost"
											monthly_cost = "monthly_cost"
										}
									}
								}
								summary {
									total_detected_resources = 1
									total_supported_resources = 1
									total_unsupported_resources = 1
									total_usage_based_resources = 1
									total_no_price_resources = 1
									unsupported_resource_counts = { "key": 1 }
									no_price_resource_counts = { "key": 1 }
								}
							}
							summary {
								total_detected_resources = 1
								total_supported_resources = 1
								total_unsupported_resources = 1
								total_usage_based_resources = 1
								total_no_price_resources = 1
								unsupported_resource_counts = { "key": 1 }
								no_price_resource_counts = { "key": 1 }
							}
							total_hourly_cost = "total_hourly_cost"
							total_monthly_cost = "total_monthly_cost"
							past_total_hourly_cost = "past_total_hourly_cost"
							past_total_monthly_cost = "past_total_monthly_cost"
							diff_total_hourly_cost = "diff_total_hourly_cost"
							diff_total_monthly_cost = "diff_total_monthly_cost"
							time_generated = "2021-01-31T09:44:12Z"
						}
						dependencies {
							catalog_id = "catalog_id"
							id = "id"
							name = "name"
							version = "version"
							flavors = [ "flavors" ]
						}
					}
					is_consumable = true
				}
				plans {
					id = "id"
					label = "label"
					name = "name"
					short_description = "short_description"
					long_description = "long_description"
					metadata = { "key": null }
					tags = [ "tags" ]
					additional_features {
						title = "title"
						title_i18n = { "key": "inner" }
						description = "description"
						description_i18n = { "key": "inner" }
					}
					created = "2021-01-31T09:44:12Z"
					updated = "2021-01-31T09:44:12Z"
					deployments {
						id = "id"
						label = "label"
						name = "name"
						short_description = "short_description"
						long_description = "long_description"
						metadata = { "key": null }
						tags = [ "tags" ]
						created = "2021-01-31T09:44:12Z"
						updated = "2021-01-31T09:44:12Z"
					}
				}
			}
			pc_managed = %s
			publish_approved = %s
			share_with_all = %s
			share_with_ibm = %s
			share_enabled = %s
			permit_request_ibm_public_publish = %s
			ibm_publish_approved = %s
			public_publish_approved = %s
			public_original_crn = "%s"
			publish_public_crn = "%s"
			portal_approval_record = "%s"
			portal_ui_url = "%s"
			catalog_id = ibm_cm_catalog.cm_catalog.id
			catalog_name = "%s"
			metadata = "FIXME"
			disclaimer = "%s"
			hidden = %s
			provider = "%s"
			provider_info {
				id = "id"
				name = "name"
			}
			repo_info {
				token = "token"
				type = "type"
			}
			image_pull_keys {
				name = "name"
				value = "value"
				description = "description"
			}
			support {
				url = "url"
				process = "process"
				process_i18n = { "key": "inner" }
				locations = [ "locations" ]
				support_details {
					type = "type"
					contact = "contact"
					response_wait_time {
						value = 1
						type = "type"
					}
					availability {
						times {
							day = 1
							start_time = "start_time"
							end_time = "end_time"
						}
						timezone = "timezone"
						always_available = true
					}
				}
				support_escalation {
					escalation_wait_time {
						value = 1
						type = "type"
					}
					response_wait_time {
						value = 1
						type = "type"
					}
					contact = "contact"
				}
				support_type = "support_type"
			}
			media {
				url = "url"
				api_url = "api_url"
				url_proxy {
					url = "url"
					sha = "sha"
				}
				caption = "caption"
				caption_i18n = { "key": "inner" }
				type = "type"
				thumbnail_url = "thumbnail_url"
			}
			deprecate_pending {
				deprecate_date = "2021-01-31T09:44:12Z"
				deprecate_state = "deprecate_state"
				description = "description"
			}
			product_kind = "%s"
			badges {
				id = "id"
				label = "label"
				label_i18n = { "key": "inner" }
				description = "description"
				description_i18n = { "key": "inner" }
				icon = "icon"
				authority = "authority"
				tag = "tag"
				learn_more_links {
					first_party = "first_party"
					third_party = "third_party"
				}
				constraints {
					type = "type"
				}
			}
		}
	`, url, crn, label, name, offeringIconURL, offeringDocsURL, offeringSupportURL, shortDescription, longDescription, pcManaged, publishApproved, shareWithAll, shareWithIBM, shareEnabled, permitRequestIBMPublicPublish, ibmPublishApproved, publicPublishApproved, publicOriginalCRN, publishPublicCRN, portalApprovalRecord, portalUIURL, catalogName, disclaimer, hidden, provider, productKind)
}

func testAccCheckIBMCmOfferingExists(n string, obj catalogmanagementv1.Offering) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getOfferingOptions.SetCatalogIdentifier(parts[0])
		getOfferingOptions.SetOfferingID(parts[1])

		offering, _, err := catalogManagementClient.GetOffering(getOfferingOptions)
		if err != nil {
			return err
		}

		obj = *offering
		return nil
	}
}

func testAccCheckIBMCmOfferingDestroy(s *terraform.State) error {
	catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cm_offering" {
			continue
		}

		getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getOfferingOptions.SetCatalogIdentifier(parts[0])
		getOfferingOptions.SetOfferingID(parts[1])

		// Try to find the key
		_, response, err := catalogManagementClient.GetOffering(getOfferingOptions)

		if err == nil {
			return fmt.Errorf("cm_offering still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cm_offering (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
