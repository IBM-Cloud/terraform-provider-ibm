---
layout: "ibm"
page_title: "IBM : ibm_onboarding_catalog_product"
description: |-
  Manages onboarding_catalog_product.
subcategory: "Partner Center Sell"
---

# ibm_onboarding_catalog_product

**Note - Intended for internal use only. This resource is strictly experimental and subject to change without notice.**

Create, update, and delete onboarding_catalog_products with this resource.

## Example Usage

```hcl
resource "ibm_onboarding_catalog_product" "onboarding_catalog_product_instance" {
  active = true
  disabled = false
  images {
		image = "image"
  }
  kind = "service"
  metadata {
		rc_compatible = true
		ui {
			strings {
				en {
					bullets {
						description = "description"
						description_i18n = { "key" = "inner" }
						title = "title"
						title_i18n = { "key" = "inner" }
					}
					media {
						caption = "caption"
						caption_i18n = { "key" = "inner" }
						thumbnail = "thumbnail"
						type = "image"
						url = "url"
					}
					embeddable_dashboard = "embeddable_dashboard"
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
		}
		service {
			rc_provisionable = true
			iam_compatible = true
			bindable = true
			plan_updateable = true
			service_key_supported = true
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
  name = "1p-service-08-06"
  object_provider {
		name = "name"
		email = "email"
  }
  overview_ui {
		en {
			display_name = "display_name"
			description = "description"
			long_description = "long_description"
		}
  }
  product_id = ibm_onboarding_product.onboarding_product_instance.id
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `active` - (Required, Boolean) Whether the service is active.
* `disabled` - (Required, Boolean) Determines the global visibility for the catalog entry, and its children. If it is not enabled, all plans are disabled.
* `env` - (Optional, String) The environment to fetch this object from.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z]+$/`.
* `images` - (Optional, List) Images from the global catalog entry that help illustrate the service.
Nested schema for **images**:
	* `image` - (Optional, String) The URL for your product logo.
	  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
* `kind` - (Required, String) The kind of the global catalog object.
  * Constraints: Allowable values are: `service`, `platform_service`, `composite`.
* `metadata` - (Optional, List) The global catalog service metadata object.
Nested schema for **metadata**:
	* `other` - (Optional, List) The additional metadata of the service in global catalog.
	Nested schema for **other**:
		* `composite` - (Optional, List) Optional metadata of the service defining it as a composite.
		Nested schema for **composite**:
			* `children` - (Optional, List)
			  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
			Nested schema for **children**:
				* `kind` - (Optional, String) The type of the composite child.
				  * Constraints: Allowable values are: `service`, `platform_service`.
				* `name` - (Optional, String) The name of the composite child.
				  * Constraints: The maximum length is `100` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-z0-9\\-.]+$/`.
			* `composite_kind` - (Optional, String) The type of the composite service.
			  * Constraints: Allowable values are: `service`, `platform_service`.
			* `composite_tag` - (Optional, String) The tag used for the composite parent and its children.
			  * Constraints: The maximum length is `100` characters. The minimum length is `2` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `pc` - (Optional, List) The metadata of the service owned and managed by Partner Center - Sell.
		Nested schema for **pc**:
			* `support` - (Optional, List) The support metadata of the service.
			Nested schema for **support**:
				* `languages` - (Optional, List) The languages in which support is available.
				  * Constraints: The maximum length is `200` items. The minimum length is `0` items.
				* `locations` - (Optional, List) The countries in which your support is available. Provide a list of country codes.
				  * Constraints: The maximum length is `200` items. The minimum length is `0` items.
				* `process` - (Optional, String) The description of your support process.
				  * Constraints: The maximum length is `1500` characters. The minimum length is `0` characters.
				* `process_i18n` - (Optional, Map) The description of your support process.
				* `status_url` - (Optional, String) The URL where the status of your service is available.
				* `support_details` - (Optional, List) The support options for the service.
				  * Constraints: The maximum length is `6` items. The minimum length is `0` items.
				Nested schema for **support_details**:
					* `availability` - (Optional, List) The time period during which support is available for the service.
					Nested schema for **availability**:
						* `always_available` - (Optional, Boolean) Whether the support for the service is always available.
						* `times` - (Optional, List) The support hours available for the service.
						  * Constraints: The maximum length is `7` items. The minimum length is `0` items.
						Nested schema for **times**:
							* `day` - (Optional, Float) The number of days in a week when support is available for the service.
							* `end_time` - (Optional, String) The time in the day when support ends for the service.
							* `start_time` - (Optional, String) The time in the day when support starts for the service.
						* `timezone` - (Optional, String) The timezones in which support is available. Only relevant if `always_available` is set to false.
					* `contact` - (Optional, String) The contact information for this support channel.
					* `response_wait_time` - (Optional, List) The time interval of providing support in units and values.
					Nested schema for **response_wait_time**:
						* `type` - (Optional, String) The unit of the time.
						* `value` - (Optional, Float) The number of time units.
					* `type` - (Optional, String) The type of support for this support channel.
					  * Constraints: Allowable values are: `support_site`, `email`, `chat`, `slack`, `phone`, `other`.
				* `support_escalation` - (Optional, List) The details of the support escalation process.
				Nested schema for **support_escalation**:
					* `contact` - (Optional, String) The support contact information of the escalation team.
					* `escalation_wait_time` - (Optional, List) The time interval of providing support in units and values.
					Nested schema for **escalation_wait_time**:
						* `type` - (Optional, String) The unit of the time.
						* `value` - (Optional, Float) The number of time units.
					* `response_wait_time` - (Optional, List) The time interval of providing support in units and values.
					Nested schema for **response_wait_time**:
						* `type` - (Optional, String) The unit of the time.
						* `value` - (Optional, Float) The number of time units.
				* `support_type` - (Optional, String) The type of support provided.
				  * Constraints: Allowable values are: `community`, `third_party`, `ibm`, `ibm_cloud`.
				* `url` - (Optional, String) The support site URL where the support for your service is available.
	* `rc_compatible` - (Optional, Boolean) Whether the object is compatible with the resource controller service.
	* `service` - (Optional, List) The global catalog metadata of the service.
	Nested schema for **service**:
		* `bindable` - (Optional, Boolean) Deprecated. Controls the Connections tab on the Resource Details page.
		* `iam_compatible` - (Optional, Boolean) Whether the service is compatible with the IAM service.
		* `plan_updateable` - (Optional, Boolean) Indicates plan update support and controls the Plan tab on the Resource Details page.
		* `rc_provisionable` - (Optional, Boolean) Whether the service is provisionable by the resource controller service.
		* `service_key_supported` - (Optional, Boolean) Indicates service credentials support and controls the Service Credential tab on Resource Details page.
	* `ui` - (Optional, List) The UI metadata of this service.
	Nested schema for **ui**:
		* `hidden` - (Optional, Boolean) Whether the object is hidden from the consumption catalog.
		* `side_by_side_index` - (Optional, Float) When the objects are listed side-by-side, this value controls the ordering.
		* `strings` - (Optional, List) The data strings.
		Nested schema for **strings**:
			* `en` - (Optional, List) Translated content of additional information about the service.
			Nested schema for **en**:
				* `bullets` - (Optional, List) The list of features that highlights your product's attributes and benefits for users.
				  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
				Nested schema for **bullets**:
					* `description` - (Optional, String) The description about the features of the product.
					  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
					* `description_i18n` - (Optional, Map) The description about the features of the product in translation.
					* `title` - (Optional, String) The descriptive title for the feature.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
					* `title_i18n` - (Optional, Map) The descriptive title for the feature in translation.
				* `embeddable_dashboard` - (Optional, String) On a service kind record this controls if your service has a custom dashboard or Resource Detail page.
				  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
				* `media` - (Optional, List) The list of supporting media for this product.
				  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
				Nested schema for **media**:
					* `caption` - (Required, String) Provide a descriptive caption that indicates what the media illustrates. This caption is displayed in the catalog.
					  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
					* `caption_i18n` - (Optional, Map) The brief explanation for your images and videos in translation.
					* `thumbnail` - (Optional, String) The reduced-size version of your images and videos.
					  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
					* `type` - (Required, String) The type of the media.
					  * Constraints: Allowable values are: `image`, `youtube`, `video_mp_4`, `video_webm`.
					* `url` - (Required, String) The URL that links to the media that shows off the product.
					  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
		* `urls` - (Optional, List) Metadata with URLs related to a service.
		Nested schema for **urls**:
			* `apidocs_url` - (Optional, String) The URL for your product's API documentation.
			  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
			* `catalog_details_url` - (Optional, String) Controls the Provisioning page URL, if set the assumption is that this URL is the provisioning URL for your service.
			  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
			* `custom_create_page_url` - (Optional, String) Controls the Provisioning page URL, if set the assumption is that this URL is the provisioning URL for your service.
			  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
			* `dashboard` - (Optional, String) Controls if your service has a custom dashboard or Resource Detail page.
			  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
			* `doc_url` - (Optional, String) The URL for your product's documentation.
			  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
			* `instructions_url` - (Optional, String) Controls the Getting Started tab on the Resource Details page. Setting it the content is loaded from the specified URL.
			  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
			* `terms_url` - (Optional, String) The URL for your product's end user license agreement.
			  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
* `name` - (Required, String) The programmatic name of this product.
  * Constraints: The value must match regular expression `/^[a-z0-9\\-.]+$/`.
* `object_provider` - (Required, List) The provider or owner of the product.
Nested schema for **object_provider**:
	* `email` - (Optional, String) The email address of the provider.
	* `name` - (Optional, String) The name of the provider.
* `overview_ui` - (Optional, List) The object that contains the service details from the Overview page in global catalog.
Nested schema for **overview_ui**:
	* `en` - (Optional, List) Translated details about the service, for example, display name, short description, and long description.
	Nested schema for **en**:
		* `description` - (Optional, String) The short description of the product that is displayed in your catalog entry.
		* `display_name` - (Optional, String) The display name of the product.
		* `long_description` - (Optional, String) The detailed description of your product that is displayed at the beginning of your product page in the catalog. Markdown markup language is supported.
* `product_id` - (Required, Forces new resource, String) The unique ID of the product.
  * Constraints: The maximum length is `71` characters. The minimum length is `71` characters. The value must match regular expression `/^[a-zA-Z0-9]{32}:o:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `tags` - (Required, List) A list of tags that carry information about your product. These tags can be used to find your product in the IBM Cloud catalog.
  * Constraints: The list items must match regular expression `/^[a-z0-9\\-._]+$/`. The maximum length is `100` items. The minimum length is `0` items.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the onboarding_catalog_product.
* `catalog_product_id` - (String) The ID of a global catalog object.
* `url` - (String) The global catalog URL of your product.


## Import

You can import the `ibm_onboarding_catalog_product` resource by using `id`.
The `id` property can be formed from `product_id`, and `catalog_product_id` in the following format:

<pre>
&lt;product_id&gt;/&lt;catalog_product_id&gt;
</pre>
* `product_id`: A string. The unique ID of the product.
* `catalog_product_id`: A string. The ID of a global catalog object.

# Syntax
<pre>
$ terraform import ibm_onboarding_catalog_product.onboarding_catalog_product &lt;product_id&gt;/&lt;catalog_product_id&gt;
</pre>
