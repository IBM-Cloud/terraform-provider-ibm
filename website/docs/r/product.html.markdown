---
layout: "ibm"
page_title: "IBM : ibm_product"
description: |-
  Manages product.
subcategory: "Partner Center Sell"
---

# ibm_product

Create, update, and delete products with this resource.

## Example Usage

```hcl
resource "ibm_product" "product_instance" {
  deprecate_pending {
		deprecate_date = "2021-01-31T09:44:12Z"
		deprecate_state = "deprecate_state"
		description = "description"
  }
  highlights {
		description = "description"
		description_i18n = {  }
		title = "title"
		title_i18n = {  }
  }
  material_agreement = false
  media {
		caption = "caption"
		caption_i18n = {  }
		thumbnail = "thumbnail"
		type = "image"
		url = "url"
  }
  product_type = "software"
  tax_assessment = "software"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `category` - (Optional, String) A list of values that are used to categorize products in the catalog. By using the Catalogs management CLI plug-in, run the `ibmcloud catalog offering category-options` CLI command to list all possible values.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]+([_]?[a-z0-9]+){0,10}$/`.
* `company` - (Optional, String) The name of your company.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `deprecate_pending` - (Optional, List) The deprecation process of the product is in the pending state.
Nested schema for **deprecate_pending**:
	* `deprecate_date` - (Optional, String) The time when the product was deprecated in standard ISO 8601 format.
	* `deprecate_state` - (Optional, String) The deprecation state of the product.
	  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `description` - (Optional, String) The reason why the product is getting deprecated.
	  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `description` - (Optional, String) The description of the product.
  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `documentation_url` - (Optional, String) The link to the warranted product documentation.
  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
* `editable` - (Optional, Boolean) The product can be edited.
* `highlights` - (Optional, List) The attributes of the product that differentiate it in the market.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **highlights**:
	* `description` - (Optional, String) The description about the features of the product.
	  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `description_i18n` - (Optional, List) The description about the features of the product in translation.
	Nested schema for **description_i18n**:
	* `title` - (Optional, String) The descriptive title for the feature.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `title_i18n` - (Optional, List) The descriptive title for the feature in translation.
	Nested schema for **title_i18n**:
* `icon_url` - (Optional, String) The URL for your company or product logo.
  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
* `keywords` - (Optional, List) The key search terms that are associated with your product.
  * Constraints: The list items must match regular expression `/^[a-z0-9]+([ _:-]?[a-z0-9]+){0,10}$/`. The maximum length is `100` items. The minimum length is `0` items.
* `label` - (Optional, String) The name of the product that you are onboarding. This name is displayed to users when you publish your product in the catalog.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `label_i18n` - (Optional, List) Translated strings for the name of the product.
Nested schema for **label_i18n**:
* `long_description` - (Optional, String) The description about the details of the product. You can use markdown syntax to provide this description.
  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `long_description_i18n` - (Optional, List) Translated strings for describing the details of the product. You can use markdown syntax to provide this description.
Nested schema for **long_description_i18n**:
* `material_agreement` - (Optional, Boolean) The confirmation that your company is authorized to use all materials.
* `media` - (Optional, List) The images or videos that show off the product.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **media**:
	* `caption` - (Required, String) Provide a brief explanation that indicates what the media illustrates. This caption is displayed in the catalog.
	  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `caption_i18n` - (Optional, List) The brief explanation for your images and videos in translation.
	Nested schema for **caption_i18n**:
	* `thumbnail` - (Optional, String) The reduced-size version of your images and videos.
	  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
	* `type` - (Required, String) The type of the media.
	  * Constraints: Allowable values are: `image`, `youtube`, `video_mp_4`, `video_webm`.
	* `url` - (Required, String) The URL that links to the media that shows off the product.
	  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
* `name` - (Optional, String) The name of the product.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `pricing_model` - (Optional, String) The pricing model for your product.
  * Constraints: Allowable values are: `free`, `byol`.
* `product_type` - (Optional, String) The type of the product that you want to onboard to IBM Cloud.
  * Constraints: Allowable values are: `software`, `solution`, `service`.
* `provider_type` - (Optional, List) The group this product's provider is member of.
  * Constraints: Allowable list items are: `ibm_community`, `ibm_third_party`. The maximum length is `2` items. The minimum length is `1` item.
* `short_description_i18n` - (Optional, List) Translated strings for the description of the product.
Nested schema for **short_description_i18n**:
* `tags` - (Optional, List) The keywords and phrases that are associated with your product.
  * Constraints: The list items must match regular expression `/^[a-z0-9]+([ _:-]?[a-z0-9]+){0,10}$/`. The maximum length is `1000` items. The minimum length is `0` items.
* `tax_assessment` - (Optional, String) The tax assessment for your product.
  * Constraints: Allowable values are: `software`, `saas`, `iaas`, `paas`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the product.
* `account_id` - (String) The unique ID for the account in which the product is being onboarded.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-zA-Z0-9]+$/`.
* `catalog_id` - (String) The ID of the private catalog where your products are created.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
* `catalog_offering_id` - (String) The unique ID of the offering in the private catalog.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
* `created_at` - (String) The time when the new product was created in standard ISO 8601 format.
  * Constraints: The maximum length is `29` characters. The minimum length is `13` characters. The value must match regular expression `/^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}\\.\\d{3}Z$/`.
* `publish_state` - (String) The actual publishing state of the product.
  * Constraints: Allowable values are: `deprecated`, `ibm_publish_permitted`, `public_publish_permitted`, `publish_permitted`, `publish_unpermitted`, `published`, `under_deprecation`.
* `published_at` - (String) The time when the new product was published to the IBM Cloud catalog in standard ISO 8601 format.
  * Constraints: The maximum length is `29` characters. The minimum length is `13` characters. The value must match regular expression `/^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}\\.\\d{3}Z$/`.
* `published_to_access_list` - (Boolean) The product is published to an access list. An access list is a list of accounts that your product is potentially shared with.
* `published_to_ibm` - (Boolean) The product is available to all IBMers.
* `published_to_public` - (Boolean) The product is published to the IBM Cloud catalog.
* `updated_at` - (String) The time when the product was updated in standard ISO 8601 format.
  * Constraints: The maximum length is `29` characters. The minimum length is `13` characters. The value must match regular expression `/^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}\\.\\d{3}Z$/`.


## Import

You can import the `ibm_product` resource by using `id`. The ID that uniquely identifies the product in Partner Center. This ID can be found on the Dashboard tab in Partner Center.

# Syntax
```
$ terraform import ibm_product.product <id>
```
