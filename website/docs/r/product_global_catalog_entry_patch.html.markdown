---
layout: "ibm"
page_title: "IBM : ibm_product_global_catalog_entry_patch"
description: |-
  Manages product_global_catalog_entry_patch.
subcategory: "Partner Center Sell"
---

# ibm_product_global_catalog_entry_patch

Provides a resource for product_global_catalog_entry_patch. This allows product_global_catalog_entry_patch to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_product_global_catalog_entry_patch" "product_global_catalog_entry_patch_instance" {
  images {
		feature_image = "feature_image"
		image = "image"
		medium_image = "medium_image"
		small_image = "small_image"
  }
  metadata {
		other = {  }
		rc_compatible = true
		service {
			async_provisioning_supported = true
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
			unique_api_key = true
		}
		ui = {  }
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
  provider {
		email = "email"
		name = "name"
		contact = "contact"
		support_email = "support_email"
		phone = "phone"
  }
  visibility {
		restriction = "restriction"
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `active` - (Optional, Boolean) The status of the gc entry.
* `catalog_crn` - (Optional, String) The cloud resource name in the catalog.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `children_url` - (Optional, String) Url to the children of the entry.
  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
* `created` - (Optional, String) Creation Date of the entry.
* `disabled` - (Optional, Boolean) Is the current entry enabled or not.
* `geo_tags` - (Optional, List) List of geo tags.
  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `images` - (Optional, List) Links to Global Catalog images.
Nested scheme for **images**:
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
Nested scheme for **metadata**:
	* `other` - (Optional, List) Metadata of Global Catalog entry.
	Nested scheme for **other**:
	* `rc_compatible` - (Optional, Boolean) Is compatible with rc.
	* `service` - (Optional, List) Metadata of the service.
	Nested scheme for **service**:
		* `async_provisioning_supported` - (Optional, Boolean) Async provisioning supported or not.
		* `async_unprovisioning_supported` - (Optional, Boolean) Async unprovisioning Supported or not.
		* `bindable` - (Optional, Boolean) Bindable or not.
		* `custom_create_page_hybrid_enabled` - (Optional, Boolean) Custom create page hybrid enabled or not.
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
	* `ui` - (Optional, List) Ui data.
	Nested scheme for **ui**:
* `name` - (Optional, String) The cloud resource name in the catalog.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `overview_ui` - (Optional, List) Ui overview language information.
Nested scheme for **overview_ui**:
	* `de` - (Optional, List) German UI overview languages keys.
	Nested scheme for **de**:
		* `description` - (Optional, String) Description in German.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `display_name` - (Optional, String) Display name in German.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `long_description` - (Optional, String) Long description in German.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
	* `en` - (Optional, List) English UI overview languages keys.
	Nested scheme for **en**:
		* `description` - (Optional, String) Description in English.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `display_name` - (Optional, String) Display name in English.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `long_description` - (Optional, String) Long description in English.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `es` - (Optional, List) Spanish UI overview languages keys.
	Nested scheme for **es**:
		* `description` - (Optional, String) Description in Spanish.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wáéíóúñ]+\\b/`.
		* `display_name` - (Optional, String) Display name in Spanish.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wáéíóúñ]+\\b/`.
		* `long_description` - (Optional, String) Long description in Spanish.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wáéíóúñ]+\\b/`.
	* `fr` - (Optional, List) French UI overview languages keys.
	Nested scheme for **fr**:
		* `description` - (Optional, String) Description in French.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
		* `display_name` - (Optional, String) Display name in French.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
		* `long_description` - (Optional, String) Long description in French.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
	* `it` - (Optional, List) Italian UI overview languages keys.
	Nested scheme for **it**:
		* `description` - (Optional, String) Description in Italian.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `display_name` - (Optional, String) Display name in Italian.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
		* `long_description` - (Optional, String) Long description in Italian.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
	* `ja` - (Optional, List) Japanese UI overview languages keys.
	Nested scheme for **ja**:
		* `description` - (Optional, String) Description in Japanese.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `display_name` - (Optional, String) Display name in Japanese.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `long_description` - (Optional, String) Long description in Japanese.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
	* `ko` - (Optional, List) Korean UI overview languages keys.
	Nested scheme for **ko**:
		* `description` - (Optional, String) Description in Korean.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `display_name` - (Optional, String) Display name in Korean.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `long_description` - (Optional, String) Long description in Korean.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `pt_br` - (Optional, List) Portuguese UI overview languages keys.
	Nested scheme for **pt_br**:
		* `description` - (Optional, String) Description in Portuguese.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `display_name` - (Optional, String) Display name in Portuguese.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `long_description` - (Optional, String) Long description in Portuguese.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
	* `zh_cn` - (Optional, List) Chinese simplified UI overview languages keys.
	Nested scheme for **zh_cn**:
		* `description` - (Optional, String) Description in Chinese simplified.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `display_name` - (Optional, String) Display name in Chinese simplified.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `long_description` - (Optional, String) Long description in Chinese simplified.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `zh_tw` - (Optional, List) Chinese traditional UI overview languages keys.
	Nested scheme for **zh_tw**:
		* `description` - (Optional, String) Description in Chinese traditional.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `display_name` - (Optional, String) Display name in Chinese traditional.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `long_description` - (Optional, String) Long description in Chinese traditional.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
* `pricing_tags` - (Optional, List) List of pricing tags.
  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `1` item.
* `provider` - (Optional, List) Provider for the entry.
Nested scheme for **provider**:
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
* `tags` - (Optional, List) List of pricing tags.
  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `1` item.
* `updated` - (Optional, String) When it was last updated.
* `url` - (Optional, String) Link to the catalog entry.
  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
* `visibility` - (Optional, List) Visibility option for the entry.
Nested scheme for **visibility**:
	* `restriction` - (Optional, String) Visibility restriction for the entry.
	  * Constraints: The maximum length is `512` characters. The minimum length is `5` characters. The value must match regular expression `/^[ -~\\s]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the product_global_catalog_entry_patch.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_product_global_catalog_entry_patch` resource by using `id`. The unique identifier of Catalog entry.

# Syntax
```
$ terraform import ibm_product_global_catalog_entry_patch.product_global_catalog_entry_patch <id>
```
