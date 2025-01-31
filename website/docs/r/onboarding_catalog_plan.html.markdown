---
layout: "ibm"
page_title: "IBM : ibm_onboarding_catalog_plan"
description: |-
  Manages onboarding_catalog_plan.
subcategory: "Partner Center Sell"
---

# ibm_onboarding_catalog_plan

**Note - Intended for internal use only. This resource is strictly experimental and subject to change without notice.**

Create, update, and delete onboarding_catalog_plans with this resource.

## Example Usage

```hcl
resource "ibm_onboarding_catalog_plan" "onboarding_catalog_plan_instance" {
  active = true
  catalog_product_id = ibm_onboarding_catalog_product.onboarding_catalog_product_instance.onboarding_catalog_product_id
  disabled = false
  kind = "plan"
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
		pricing {
			type = "free"
			origin = "global_catalog"
		}
		plan {
			allow_internal_users = true
			bindable = true
		}
  }
  name = "free-plan2"
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
* `catalog_product_id` - (Required, Forces new resource, String) The unique ID of this global catalog product.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z\\-_\\d]+$/`.
* `disabled` - (Required, Boolean) Determines the global visibility for the catalog entry, and its children. If it is not enabled, all plans are disabled.
* `env` - (Optional, String) The environment to fetch this object from.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z]+$/`.
* `kind` - (Required, String) The kind of the global catalog object.
  * Constraints: Allowable values are: `plan`.
* `metadata` - (Optional, List) Global catalog plan metadata.
Nested schema for **metadata**:
	* `plan` - (Optional, List) Metadata controlling Plan related settings.
	Nested schema for **plan**:
		* `allow_internal_users` - (Optional, Boolean) Controls if IBMers are allowed to provision this plan.
		* `bindable` - (Optional, Boolean) Deprecated. Controls the Connections tab on the Resource Details page.
	* `pricing` - (Optional, List) The pricing metadata of this object.
	Nested schema for **pricing**:
		* `origin` - (Optional, String) The source of the pricing information: global_catalog or pricing_catalog.
		  * Constraints: Allowable values are: `global_catalog`, `pricing_catalog`.
		* `type` - (Optional, String) The type of the pricing plan.
		  * Constraints: Allowable values are: `free`, `paid`, `subscription`.
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
* `name` - (Required, String) The programmatic name of this plan.
  * Constraints: The value must match regular expression `/^[a-zA-Z0-9\\-.]+$/`.
* `object_id` - (Optional, String) The desired ID of the global catalog object.
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

* `id` - The unique identifier of the onboarding_catalog_plan.
* `catalog_plan_id` - (String) The ID of a global catalog object.
* `geo_tags` - (List) 
  * Constraints: The list items must match regular expression `/./`. The maximum length is `1000` items. The minimum length is `0` items.
* `url` - (String) The global catalog URL of your product.


## Import

You can import the `ibm_onboarding_catalog_plan` resource by using `id`.
The `id` property can be formed from `product_id`, `catalog_product_id`, and `catalog_plan_id` in the following format:

<pre>
&lt;product_id&gt;/&lt;catalog_product_id&gt;/&lt;catalog_plan_id&gt;
</pre>
* `product_id`: A string. The unique ID of the product.
* `catalog_product_id`: A string. The unique ID of this global catalog product.
* `catalog_plan_id`: A string. The ID of a global catalog object.

# Syntax
<pre>
$ terraform import ibm_onboarding_catalog_plan.onboarding_catalog_plan product_id/catalog_product_id/catalog_plan_id;
</pre>
