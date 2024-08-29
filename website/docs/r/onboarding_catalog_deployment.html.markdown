---
layout: "ibm"
page_title: "IBM : ibm_onboarding_catalog_deployment"
description: |-
  Manages onboarding_catalog_deployment.
subcategory: "Partner Center Sell"
---

# ibm_onboarding_catalog_deployment

**Note - Intended for internal use only. This resource is strictly experimental and subject to change without any further notice.**

Create, update, and delete onboarding_catalog_deployments with this resource.

## Example Usage

```hcl
resource "ibm_onboarding_catalog_deployment" "onboarding_catalog_deployment_instance" {
  active = true
  catalog_plan_id = "catalog_plan_id"
  catalog_product_id = "catalog_product_id"
  disabled = true
  kind = "deployment"
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
				}
			}
			urls {
				doc_url = "doc_url"
				terms_url = "terms_url"
			}
			hidden = true
			side_by_side_index = 1.0
		}
		service {
			rc_provisionable = true
			iam_compatible = true
		}
  }
  name = "name"
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
  product_id = "product_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `active` - (Required, Boolean) Whether the service is active.
* `catalog_plan_id` - (Required, Forces new resource, String) The unique ID of this global catalog plan.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z\\-_\\d]+$/`.
* `catalog_product_id` - (Required, Forces new resource, String) The unique ID of this global catalog product.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z\\-_\\d]+$/`.
* `disabled` - (Required, Boolean) Determines the global visibility for the catalog entry, and its children. If it is not enabled, all plans are disabled.
* `env` - (Optional, String) The environment to fetch this object from.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z]+$/`.
* `kind` - (Required, String) The kind of the global catalog object.
  * Constraints: Allowable values are: `deployment`.
* `metadata` - (Optional, List) Global catalog deployment metadata.
Nested schema for **metadata**:
	* `rc_compatible` - (Optional, Boolean) Whether the object is compatible with the resource controller service.
	* `service` - (Optional, List) The global catalog metadata of the service.
	Nested schema for **service**:
		* `iam_compatible` - (Optional, Boolean) Whether the service is compatible with the IAM service.
		* `rc_provisionable` - (Optional, Boolean) Whether the service is provisionable by the resource controller service.
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
		* `urls` - (Optional, List) The UI based URLs.
		Nested schema for **urls**:
			* `doc_url` - (Optional, String) The URL for your product documentation.
			* `terms_url` - (Optional, String) The URL for your product's end user license agreement.
* `name` - (Required, String) The programmatic name of this deployment.
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

* `id` - The unique identifier of the onboarding_catalog_deployment.
* `catalog_deployment_id` - (String) The ID of a global catalog object.
* `url` - (String) The global catalog URL of your product.


## Import

You can import the `ibm_onboarding_catalog_deployment` resource by using `id`.
The `id` property can be formed from `product_id`, `catalog_product_id`, `catalog_plan_id`, and `catalog_deployment_id` in the following format:

<pre>
&lt;product_id&gt;/&lt;catalog_product_id&gt;/&lt;catalog_plan_id&gt;/&lt;catalog_deployment_id&gt;
</pre>
* `product_id`: A string. The unique ID of the product.
* `catalog_product_id`: A string. The unique ID of this global catalog product.
* `catalog_plan_id`: A string. The unique ID of this global catalog plan.
* `catalog_deployment_id`: A string. The ID of a global catalog object.

# Syntax
<pre>
$ terraform import ibm_onboarding_catalog_deployment.onboarding_catalog_deployment <catalogproductid>/<planid>/<deploymentid>
</pre>
