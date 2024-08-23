---
layout: "ibm"
page_title: "IBM : ibm_product_global_catalog_entry"
description: |-
  Manages product_global_catalog_entry.
subcategory: "Partner Center Sell"
---

# ibm_product_global_catalog_entry

Create, update, and delete product_global_catalog_entrys with this resource.

## Example Usage

```hcl
resource "ibm_product_global_catalog_entry" "product_global_catalog_entry_instance" {
  images {
		feature_image = "feature_image"
		image = "image"
		medium_image = "medium_image"
		small_image = "small_image"
  }
  metadata {
		other {
			composite {
				children {
					kind = "kind"
					name = "name"
				}
				composite_kind = "composite_kind"
				composite_tag = "composite_tag"
			}
			iam {
				key = "key"
				value = "value"
				supported_attributes {
					key = "key"
					options {
						operators = [ "operators" ]
						hidden = true
					}
					option_datasource {
						type = "type"
						values {
							value = "value"
							strings {
								en {
									display_name = "display_name"
								}
							}
						}
						gst_search_details {
							label_property_name = "label_property_name"
							query = "query"
							value_property_name = "value_property_name"
						}
					}
					type = "type"
					strings {
						en = { "key" = "anything as a string" }
					}
					description = { "key" = "anything as a string" }
					display_name = { "key" = "anything as a string" }
					ui {
						input_type = "input_type"
						input_details {
							type = "type"
							values {
								display_name {
									default = "default"
									en = "en"
									de = "de"
									fr = "fr"
									es = "es"
									it = "it"
									ja = "ja"
									ko = "ko"
									pt_br = "pt_br"
									zh_cn = "zh_cn"
									zh_tw = "zh_tw"
								}
								value = "value"
							}
						}
					}
				}
			}
			ui {
				hidden = true
			}
			swagger_urls {
				audience = "audience"
				children {
					audience = "audience"
					disable_download = true
					aliases = [ "aliases" ]
					category = [ "category" ]
					extensions = [ "extensions" ]
					file = "file"
					id = "id"
					subcollection = "subcollection"
					i18n {
						en = { "key" = "anything as a string" }
					}
					sdk = true
					links {
						cli = "cli"
						docs = "docs"
						terraform = "terraform"
					}
					versions {
						file = "file"
						id = "id"
						sdk = true
						i18n {
							en = { "key" = "anything as a string" }
						}
						version_label = "version_label"
					}
				}
				disable_download = true
				aliases = [ "aliases" ]
				category = [ "category" ]
				extensions = [ "extensions" ]
				file = "file"
				id = "id"
				subcollection = "subcollection"
				i18n {
					en = { "key" = "anything as a string" }
				}
				sdk = true
				links {
					cli = "cli"
					docs = "docs"
					terraform = "terraform"
				}
				versions {
					file = "file"
					id = "id"
					sdk = true
					i18n {
						en = { "key" = "anything as a string" }
					}
					version_label = "version_label"
				}
			}
		}
		rc_compatible = true
		original_name = "original_name"
		service {
			async_provisioning_supported = true
			user_defined_service = "user_defined_service"
			async_unprovisioning_supported = true
			bindable = true
			custom_create_page_hybrid_enabled = true
			iam_compatible = true
			parameters = [ "parameters" ]
			plan_updateable = true
			rc_provisionable = true
			service_check_enabled = true
			service_key_supported = true
			state = "state"
			test_check_interval = 1.0
			type = "type"
			extension = "extension"
			unique_api_key = true
		}
		ui {
			accessible_during_provision = true
			resource_hidden = true
			embeddable_dashboard = "embeddable_dashboard"
			hidden = true
			end_of_service_time = "end_of_service_time"
			strings {
				en {
					bullets {
						description = "description"
						title = "title"
					}
				}
			}
			urls {
				apidocs_url = "apidocs_url"
				terms_url = "terms_url"
				doc_url = "doc_url"
				instructions_url = "instructions_url"
				catalog_details_url = "catalog_details_url"
			}
		}
		callbacks {
			us_south = "us_south"
			eu_gb = "eu_gb"
			global = "global"
			regionless = "regionless"
			api_endpoint {
				global = "global"
				regionless = "regionless"
				us_south = "us_south"
				eu_gb = "eu_gb"
			}
		}
		deployment {
			location = "location"
			location_url = "location_url"
			service_crn = "service_crn"
			target_crn = "target_crn"
			broker {
				guid = "guid"
				name = "name"
			}
		}
		plan {
			allow_internal_users = true
			async_provisioning_supported = true
			async_unprovisioning_supported = true
			bindable = true
			reservable = true
			service_check_enabled = true
			test_check_interval = 1.0
			cf_guid = { "key" = "anything as a string" }
		}
		sla {
			dr {
				dr = true
			}
		}
		compliance = [ "compliance" ]
		pricing {
			origin = "origin"
			type = "type"
			url = "url"
		}
  }
  object_provider {
		email = "email"
		name = "name"
		contact = "contact"
		support_email = "support_email"
		phone = "phone"
  }
  overview_ui {
		de {
			description = "description"
			display_name = "display_name"
			long_description = "long_description"
		}
		en {
			description = "description"
			display_name = "display_name"
			long_description = "long_description"
		}
		es {
			description = "description"
			display_name = "display_name"
			long_description = "long_description"
		}
		fr {
			description = "description"
			display_name = "display_name"
			long_description = "long_description"
		}
		it {
			description = "description"
			display_name = "display_name"
			long_description = "long_description"
		}
		ja {
			description = "description"
			display_name = "display_name"
			long_description = "long_description"
		}
		ko {
			description = "description"
			display_name = "display_name"
			long_description = "long_description"
		}
		pt_br {
			description = "description"
			display_name = "display_name"
			long_description = "long_description"
		}
		zh_cn {
			description = "description"
			display_name = "display_name"
			long_description = "long_description"
		}
		zh_tw {
			description = "description"
			display_name = "display_name"
			long_description = "long_description"
		}
  }
  visibility {
		owner = "a/9152f5dd39ef492f8ece725c96baaa11"
		restrictions = "public"
		extendable = true
		include {
			accounts = { "key" = "anything as a string" }
		}
		exclude {
			accounts = { "key" = "anything as a string" }
		}
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `active` - (Optional, Boolean) The status of the gc entry.
* `disabled` - (Optional, Boolean) Is the current entry enabled or not.
* `geo_tags` - (Optional, List) List of geo tags.
  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `images` - (Optional, List) Links to Global Catalog images.
Nested schema for **images**:
	* `feature_image` - (Optional, String) Link to image.
	  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
	* `image` - (Optional, String) Link to image.
	  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
	* `medium_image` - (Optional, String) Link to image.
	  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
	* `small_image` - (Optional, String) Link to image.
	  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
* `kind` - (Optional, String) Type of entry.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\s]*$/`.
* `metadata` - (Optional, List) Meta data of the Global catalog.
Nested schema for **metadata**:
	* `callbacks` - (Optional, List)
	Nested schema for **callbacks**:
		* `api_endpoint` - (Optional, List)
		Nested schema for **api_endpoint**:
			* `eu_gb` - (Optional, String)
			* `global` - (Optional, String)
			* `regionless` - (Optional, String)
			* `us_south` - (Optional, String)
		* `eu_gb` - (Optional, String)
		* `global` - (Optional, String)
		* `regionless` - (Optional, String)
		* `us_south` - (Optional, String)
	* `compliance` - (Optional, List)
	* `deployment` - (Optional, List)
	Nested schema for **deployment**:
		* `broker` - (Optional, List)
		Nested schema for **broker**:
			* `guid` - (Optional, String)
			* `name` - (Optional, String)
		* `location` - (Optional, String)
		* `location_url` - (Optional, String)
		* `service_crn` - (Optional, String)
		* `target_crn` - (Optional, String)
	* `original_name` - (Optional, String)
	* `other` - (Optional, List) Metadata of Global Catalog entry.
	Nested schema for **other**:
		* `composite` - (Optional, List) Additonal composite service metadata.
		Nested schema for **composite**:
			* `children` - (Optional, List)
			Nested schema for **children**:
				* `kind` - (Optional, String)
				* `name` - (Optional, String)
			* `composite_kind` - (Optional, String)
			* `composite_tag` - (Optional, String)
		* `iam` - (Optional, List) Whatever IAM extenstion.
		Nested schema for **iam**:
			* `key` - (Optional, String)
			* `supported_attributes` - (Optional, List)
			  * Constraints: The minimum length is `0` items.
			Nested schema for **supported_attributes**:
				* `description` - (Optional, Map)
				* `display_name` - (Optional, Map)
				* `key` - (Optional, String)
				* `option_datasource` - (Optional, List)
				Nested schema for **option_datasource**:
					* `gst_search_details` - (Optional, List)
					Nested schema for **gst_search_details**:
						* `label_property_name` - (Optional, String)
						* `query` - (Optional, String)
						* `value_property_name` - (Optional, String)
					* `type` - (Optional, String)
					* `values` - (Optional, List)
					  * Constraints: The minimum length is `0` items.
					Nested schema for **values**:
						* `strings` - (Optional, List)
						Nested schema for **strings**:
							* `en` - (Optional, List)
							Nested schema for **en**:
								* `display_name` - (Optional, String)
						* `value` - (Optional, String)
				* `options` - (Optional, List)
				Nested schema for **options**:
					* `hidden` - (Optional, Boolean)
					* `operators` - (Optional, List)
					  * Constraints: The minimum length is `1` item.
				* `strings` - (Optional, List)
				Nested schema for **strings**:
					* `en` - (Optional, Map)
				* `type` - (Optional, String)
				* `ui` - (Optional, List)
				Nested schema for **ui**:
					* `input_details` - (Optional, List)
					Nested schema for **input_details**:
						* `type` - (Optional, String)
						* `values` - (Optional, List)
						Nested schema for **values**:
							* `display_name` - (Optional, List)
							Nested schema for **display_name**:
								* `de` - (Optional, String)
								* `default` - (Optional, String)
								* `en` - (Optional, String)
								* `es` - (Optional, String)
								* `fr` - (Optional, String)
								* `it` - (Optional, String)
								* `ja` - (Optional, String)
								* `ko` - (Optional, String)
								* `pt_br` - (Optional, String)
								* `zh_cn` - (Optional, String)
								* `zh_tw` - (Optional, String)
							* `value` - (Optional, String)
					* `input_type` - (Optional, String)
			* `value` - (Optional, String)
		* `swagger_urls` - (Optional, List)
		  * Constraints: The minimum length is `0` items.
		Nested schema for **swagger_urls**:
			* `aliases` - (Optional, List)
			  * Constraints: The minimum length is `0` items.
			* `audience` - (Optional, String)
			* `category` - (Optional, List)
			  * Constraints: The minimum length is `0` items.
			* `children` - (Optional, List)
			  * Constraints: The minimum length is `0` items.
			Nested schema for **children**:
				* `aliases` - (Optional, List)
				  * Constraints: The minimum length is `0` items.
				* `audience` - (Optional, String)
				* `category` - (Optional, List)
				  * Constraints: The minimum length is `0` items.
				* `disable_download` - (Optional, Boolean)
				* `extensions` - (Optional, List)
				  * Constraints: The minimum length is `0` items.
				* `file` - (Optional, String)
				* `i18n` - (Optional, List)
				Nested schema for **i18n**:
					* `en` - (Optional, Map)
				* `id` - (Optional, String)
				* `links` - (Optional, List)
				Nested schema for **links**:
					* `cli` - (Optional, String)
					* `docs` - (Optional, String)
					* `terraform` - (Optional, String)
				* `sdk` - (Optional, Boolean)
				* `subcollection` - (Optional, String)
				* `versions` - (Optional, List)
				  * Constraints: The minimum length is `0` items.
				Nested schema for **versions**:
					* `file` - (Optional, String)
					* `i18n` - (Optional, List)
					Nested schema for **i18n**:
						* `en` - (Optional, Map)
					* `id` - (Optional, String)
					* `sdk` - (Optional, Boolean)
					* `version_label` - (Optional, String)
			* `disable_download` - (Optional, Boolean)
			* `extensions` - (Optional, List)
			  * Constraints: The minimum length is `0` items.
			* `file` - (Optional, String)
			* `i18n` - (Optional, List)
			Nested schema for **i18n**:
				* `en` - (Optional, Map)
			* `id` - (Optional, String)
			* `links` - (Optional, List)
			Nested schema for **links**:
				* `cli` - (Optional, String)
				* `docs` - (Optional, String)
				* `terraform` - (Optional, String)
			* `sdk` - (Optional, Boolean)
			* `subcollection` - (Optional, String)
			* `versions` - (Optional, List)
			  * Constraints: The minimum length is `0` items.
			Nested schema for **versions**:
				* `file` - (Optional, String)
				* `i18n` - (Optional, List)
				Nested schema for **i18n**:
					* `en` - (Optional, Map)
				* `id` - (Optional, String)
				* `sdk` - (Optional, Boolean)
				* `version_label` - (Optional, String)
		* `ui` - (Optional, List)
		Nested schema for **ui**:
			* `hidden` - (Optional, Boolean)
	* `plan` - (Optional, List)
	Nested schema for **plan**:
		* `allow_internal_users` - (Optional, Boolean)
		* `async_provisioning_supported` - (Optional, Boolean)
		* `async_unprovisioning_supported` - (Optional, Boolean)
		* `bindable` - (Optional, Boolean)
		* `cf_guid` - (Optional, Map)
		* `reservable` - (Optional, Boolean)
		* `service_check_enabled` - (Optional, Boolean)
		* `test_check_interval` - (Optional, Float)
	* `pricing` - (Optional, List) Pricing related metadata.
	Nested schema for **pricing**:
		* `origin` - (Optional, String)
		* `type` - (Optional, String)
		* `url` - (Optional, String)
	* `rc_compatible` - (Optional, Boolean) Is compatible with rc.
	* `service` - (Optional, List) Metadata of the service.
	Nested schema for **service**:
		* `async_provisioning_supported` - (Optional, Boolean) Async provisioning supported or not.
		* `async_unprovisioning_supported` - (Optional, Boolean) Async unprovisioning Supported or not.
		* `bindable` - (Optional, Boolean) Bindable or not.
		* `custom_create_page_hybrid_enabled` - (Optional, Boolean) Custom create page hybrid enabled or not.
		* `extension` - (Optional, String)
		* `iam_compatible` - (Optional, Boolean) Iam compatible or not.
		* `parameters` - (Optional, List) Fake item for empty array.
		  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
		* `plan_updateable` - (Optional, Boolean) Plan updatable or not.
		* `rc_provisionable` - (Optional, Boolean) Rc provisionable or not.
		* `service_check_enabled` - (Optional, Boolean) Service check enabled or not.
		* `service_key_supported` - (Optional, Boolean) Service key supported or not.
		* `state` - (Optional, String) State of metadata.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `test_check_interval` - (Optional, Float) Test check interval.
		* `type` - (Optional, String) Type of metadata.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `unique_api_key` - (Optional, Boolean) HAs unique api key or not.
		* `user_defined_service` - (Optional, String)
	* `sla` - (Optional, List)
	Nested schema for **sla**:
		* `dr` - (Optional, List)
		Nested schema for **dr**:
			* `dr` - (Optional, Boolean)
	* `ui` - (Optional, List) UI Metadata.
	Nested schema for **ui**:
		* `accessible_during_provision` - (Optional, Boolean)
		* `embeddable_dashboard` - (Optional, String)
		* `end_of_service_time` - (Optional, String)
		* `hidden` - (Optional, Boolean)
		* `resource_hidden` - (Optional, Boolean)
		* `strings` - (Optional, List)
		Nested schema for **strings**:
			* `en` - (Optional, List)
			Nested schema for **en**:
				* `bullets` - (Optional, List)
				Nested schema for **bullets**:
					* `description` - (Optional, String)
					* `title` - (Optional, String)
		* `urls` - (Optional, List)
		Nested schema for **urls**:
			* `apidocs_url` - (Optional, String)
			* `catalog_details_url` - (Optional, String)
			* `doc_url` - (Optional, String)
			* `instructions_url` - (Optional, String)
			* `terms_url` - (Optional, String)
* `name` - (Optional, String) The cloud resource name in the catalog.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `object_id` - (Optional, String) The unique identifier of Catalog entry.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\s]*$/`.
* `object_provider` - (Optional, List) Provider for the entry.
Nested schema for **object_provider**:
	* `contact` - (Optional, String) Provider's contact name.
	  * Constraints: The maximum length is `512` characters. The minimum length is `5` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `email` - (Optional, String) Provider's email address for this catalog entry.
	  * Constraints: The maximum length is `320` characters. The minimum length is `5` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `name` - (Optional, String) Provider's name, for example, IBM.
	  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `phone` - (Optional, String) Provider's support email.
	  * Constraints: The maximum length is `512` characters. The minimum length is `5` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `support_email` - (Optional, String) Provider's support email.
	  * Constraints: The maximum length is `512` characters. The minimum length is `5` characters. The value must match regular expression `/^\\S+@\\S+$/`.
* `overview_ui` - (Optional, List) Ui overview language information.
Nested schema for **overview_ui**:
	* `de` - (Optional, List) German UI overview languages keys.
	Nested schema for **de**:
		* `description` - (Optional, String) Description in German.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `display_name` - (Optional, String) Display name in German.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `long_description` - (Optional, String) Long description in German.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
	* `en` - (Optional, List) English UI overview languages keys.
	Nested schema for **en**:
		* `description` - (Optional, String) Description in English.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `display_name` - (Optional, String) Display name in English.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `long_description` - (Optional, String) Long description in English.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `es` - (Optional, List) Spanish UI overview languages keys.
	Nested schema for **es**:
		* `description` - (Optional, String) Description in Spanish.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wáéíóúñ]+\\b/`.
		* `display_name` - (Optional, String) Display name in Spanish.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wáéíóúñ]+\\b/`.
		* `long_description` - (Optional, String) Long description in Spanish.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wáéíóúñ]+\\b/`.
	* `fr` - (Optional, List) French UI overview languages keys.
	Nested schema for **fr**:
		* `description` - (Optional, String) Description in French.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
		* `display_name` - (Optional, String) Display name in French.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
		* `long_description` - (Optional, String) Long description in French.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
	* `it` - (Optional, List) Italian UI overview languages keys.
	Nested schema for **it**:
		* `description` - (Optional, String) Description in Italian.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `display_name` - (Optional, String) Display name in Italian.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
		* `long_description` - (Optional, String) Long description in Italian.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
	* `ja` - (Optional, List) Japanese UI overview languages keys.
	Nested schema for **ja**:
		* `description` - (Optional, String) Description in Japanese.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `display_name` - (Optional, String) Display name in Japanese.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `long_description` - (Optional, String) Long description in Japanese.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
	* `ko` - (Optional, List) Korean UI overview languages keys.
	Nested schema for **ko**:
		* `description` - (Optional, String) Description in Korean.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `display_name` - (Optional, String) Display name in Korean.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `long_description` - (Optional, String) Long description in Korean.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `pt_br` - (Optional, List) Portuguese UI overview languages keys.
	Nested schema for **pt_br**:
		* `description` - (Optional, String) Description in Portuguese.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `display_name` - (Optional, String) Display name in Portuguese.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `long_description` - (Optional, String) Long description in Portuguese.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
	* `zh_cn` - (Optional, List) Chinese simplified UI overview languages keys.
	Nested schema for **zh_cn**:
		* `description` - (Optional, String) Description in Chinese simplified.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `display_name` - (Optional, String) Display name in Chinese simplified.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `long_description` - (Optional, String) Long description in Chinese simplified.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `zh_tw` - (Optional, List) Chinese traditional UI overview languages keys.
	Nested schema for **zh_tw**:
		* `description` - (Optional, String) Description in Chinese traditional.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `display_name` - (Optional, String) Display name in Chinese traditional.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `long_description` - (Optional, String) Long description in Chinese traditional.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
* `parent_id` - (Optional, String) 
* `pricing_tags` - (Optional, List) List of pricing tags.
  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `1` item.
* `tags` - (Optional, List) List of pricing tags.
  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `1` item.
* `visibility` - (Optional, List) Visibility option for the entry.
Nested schema for **visibility**:
	* `exclude` - (Optional, List)
	Nested schema for **exclude**:
		* `accounts` - (Optional, Map)
	* `extendable` - (Optional, Boolean)
	* `include` - (Optional, List)
	Nested schema for **include**:
		* `accounts` - (Optional, Map)
	* `owner` - (Required, String) The owner of this object prefixed with a/.
	* `restrictions` - (Optional, String) Visibility restriction for the entry.
	  * Constraints: Allowable values are: `public`, `ibm_only`, `private`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the product_global_catalog_entry.
* `catalog_crn` - (String) The cloud resource name in the catalog.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `children_url` - (String) Url to the children of the entry.
  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
* `created` - (String) Creation Date of the entry.
* `parent_url` - (String) 
* `updated` - (String) Time of last update.
* `url` - (String) Link to the catalog entry.
  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.


## Import

You can import the `ibm_product_global_catalog_entry` resource by using `id`. The unique identifier of Catalog entry.

# Syntax
<pre>
$ terraform import ibm_product_global_catalog_entry.product_global_catalog_entry &lt;id&gt;
</pre>
