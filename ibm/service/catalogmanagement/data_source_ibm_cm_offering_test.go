// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCmOfferingDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "catalog_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_id"),
				),
			},
		},
	})
}

func TestAccIBMCmOfferingDataSourceAllArgs(t *testing.T) {
	offeringURL := fmt.Sprintf("tf_url_%d", acctest.RandIntRange(10, 100))
	offeringCRN := fmt.Sprintf("tf_crn_%d", acctest.RandIntRange(10, 100))
	offeringLabel := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	offeringName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	offeringOfferingIconURL := fmt.Sprintf("tf_offering_icon_url_%d", acctest.RandIntRange(10, 100))
	offeringOfferingDocsURL := fmt.Sprintf("tf_offering_docs_url_%d", acctest.RandIntRange(10, 100))
	offeringOfferingSupportURL := fmt.Sprintf("tf_offering_support_url_%d", acctest.RandIntRange(10, 100))
	offeringShortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	offeringLongDescription := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	offeringPcManaged := "true"
	offeringPublishApproved := "true"
	offeringShareWithAll := "true"
	offeringShareWithIBM := "true"
	offeringShareEnabled := "false"
	offeringPermitRequestIBMPublicPublish := "false"
	offeringIBMPublishApproved := "true"
	offeringPublicPublishApproved := "false"
	offeringPublicOriginalCRN := fmt.Sprintf("tf_public_original_crn_%d", acctest.RandIntRange(10, 100))
	offeringPublishPublicCRN := fmt.Sprintf("tf_publish_public_crn_%d", acctest.RandIntRange(10, 100))
	offeringPortalApprovalRecord := fmt.Sprintf("tf_portal_approval_record_%d", acctest.RandIntRange(10, 100))
	offeringPortalUIURL := fmt.Sprintf("tf_portal_ui_url_%d", acctest.RandIntRange(10, 100))
	offeringCatalogName := fmt.Sprintf("tf_catalog_name_%d", acctest.RandIntRange(10, 100))
	offeringDisclaimer := fmt.Sprintf("tf_disclaimer_%d", acctest.RandIntRange(10, 100))
	offeringHidden := "true"
	offeringProvider := fmt.Sprintf("tf_provider_%d", acctest.RandIntRange(10, 100))
	offeringProductKind := fmt.Sprintf("tf_product_kind_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingDataSourceConfig(offeringURL, offeringCRN, offeringLabel, offeringName, offeringOfferingIconURL, offeringOfferingDocsURL, offeringOfferingSupportURL, offeringShortDescription, offeringLongDescription, offeringPcManaged, offeringPublishApproved, offeringShareWithAll, offeringShareWithIBM, offeringShareEnabled, offeringPermitRequestIBMPublicPublish, offeringIBMPublishApproved, offeringPublicPublishApproved, offeringPublicOriginalCRN, offeringPublishPublicCRN, offeringPortalApprovalRecord, offeringPortalUIURL, offeringCatalogName, offeringDisclaimer, offeringHidden, offeringProvider, offeringProductKind),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "catalog_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "label_i18n.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_icon_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_docs_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_support_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "keywords.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "rating.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "short_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "short_description_i18n.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "long_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "long_description_i18n.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "features.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "features.0.title"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "features.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.0.format_kind"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.0.install_kind"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.0.target_kind"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.0.created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "kinds.0.updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "pc_managed"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "publish_approved"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "share_with_all"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "share_with_ibm"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "share_enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "permit_request_ibm_public_publish"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "ibm_publish_approved"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "public_publish_approved"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "public_original_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "publish_public_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "portal_approval_record"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "portal_ui_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "catalog_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "metadata.%"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "disclaimer"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "hidden"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "provider"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "provider_info.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "repo_info.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "image_pull_keys.#"),
					resource.TestCheckResourceAttr("data.ibm_cm_offering.cm_offering", "image_pull_keys.0.name", offeringName),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "image_pull_keys.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "image_pull_keys.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "support.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "media.#"),
					resource.TestCheckResourceAttr("data.ibm_cm_offering.cm_offering", "media.0.url", offeringURL),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "media.0.api_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "media.0.caption"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "media.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "media.0.thumbnail_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "deprecate_pending.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "product_kind"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "badges.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "badges.0.id"),
					resource.TestCheckResourceAttr("data.ibm_cm_offering.cm_offering", "badges.0.label", offeringLabel),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "badges.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "badges.0.icon"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "badges.0.authority"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "badges.0.tag"),
				),
			},
		},
	})
}

func testAccCheckIBMCmOfferingDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
		}

		resource "ibm_cm_offering" "cm_offering" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
		}

		data "ibm_cm_offering" "cm_offering" {
			catalog_identifier = ibm_cm_offering.cm_offering.catalog_identifier
			offering_id = ibm_cm_offering.cm_offering.offering_id
		}
	`)
}

func testAccCheckIBMCmOfferingDataSourceConfig(offeringURL string, offeringCRN string, offeringLabel string, offeringName string, offeringOfferingIconURL string, offeringOfferingDocsURL string, offeringOfferingSupportURL string, offeringShortDescription string, offeringLongDescription string, offeringPcManaged string, offeringPublishApproved string, offeringShareWithAll string, offeringShareWithIBM string, offeringShareEnabled string, offeringPermitRequestIBMPublicPublish string, offeringIBMPublishApproved string, offeringPublicPublishApproved string, offeringPublicOriginalCRN string, offeringPublishPublicCRN string, offeringPortalApprovalRecord string, offeringPortalUIURL string, offeringCatalogName string, offeringDisclaimer string, offeringHidden string, offeringProvider string, offeringProductKind string) string {
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

		data "ibm_cm_offering" "cm_offering" {
			catalog_identifier = ibm_cm_offering.cm_offering.catalog_identifier
			offering_id = ibm_cm_offering.cm_offering.offering_id
		}
	`, offeringURL, offeringCRN, offeringLabel, offeringName, offeringOfferingIconURL, offeringOfferingDocsURL, offeringOfferingSupportURL, offeringShortDescription, offeringLongDescription, offeringPcManaged, offeringPublishApproved, offeringShareWithAll, offeringShareWithIBM, offeringShareEnabled, offeringPermitRequestIBMPublicPublish, offeringIBMPublishApproved, offeringPublicPublishApproved, offeringPublicOriginalCRN, offeringPublishPublicCRN, offeringPortalApprovalRecord, offeringPortalUIURL, offeringCatalogName, offeringDisclaimer, offeringHidden, offeringProvider, offeringProductKind)
}
