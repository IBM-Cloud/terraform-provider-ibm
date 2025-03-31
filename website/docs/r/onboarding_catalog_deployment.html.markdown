---
layout: "ibm"
page_title: "IBM : ibm_onboarding_catalog_deployment"
description: |-
  Manages onboarding_catalog_deployment.
subcategory: "Partner Center Sell"
---

# ibm_onboarding_catalog_deployment

**Note - Intended for internal use only. This resource is strictly experimental and subject to change without notice.**

Create, update, and delete onboarding_catalog_deployments with this resource.

## Example Usage

```hcl
resource "ibm_onboarding_catalog_deployment" "onboarding_catalog_deployment_instance" {
  active = true
  catalog_plan_id = ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance.onboarding_catalog_plan_id
  catalog_product_id = ibm_onboarding_catalog_product.onboarding_catalog_product_instance.onboarding_catalog_product_id
  disabled = false
  kind = "deployment"
  metadata {
		rc_compatible = true
		service {
			rc_provisionable = true
			iam_compatible = true
			service_key_supported = true
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
				associations = { "key" = "anything as a string" }
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
		deployment {
			broker {
				name = "name"
				guid = "guid"
			}
			location = "location"
			location_url = "location_url"
			target_crn = "target_crn"
		}
  }
  name = "deployment-eu-de"
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
* `catalog_plan_id` - (Required, Forces new resource, String) The unique ID of this global catalog plan.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^\\S*$/`.
* `catalog_product_id` - (Required, Forces new resource, String) The unique ID of this global catalog product.
  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^\\S*$/`.
* `disabled` - (Required, Boolean) Determines the global visibility for the catalog entry, and its children. If it is not enabled, all plans are disabled.
* `env` - (Optional, String) The environment to fetch this object from.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z_.-]+$/`.
* `kind` - (Required, String) The kind of the global catalog object.
  * Constraints: Allowable values are: `deployment`.
* `metadata` - (Optional, List) Global catalog deployment metadata.
Nested schema for **metadata**:
	* `deployment` - (Optional, List) The global catalog metadata of the deployment.
	Nested schema for **deployment**:
		* `broker` - (Optional, List) The global catalog metadata of the deployment.
		Nested schema for **broker**:
			* `guid` - (Optional, String) Crn or guid of the resource broker.
			  * Constraints: The maximum length is `2000` characters. The minimum length is `2` characters. The value must match regular expression `/^[ -~\\s]*$/`.
			* `name` - (Optional, String) The name of the resource broker.
			  * Constraints: The maximum length is `2000` characters. The minimum length is `2` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `location` - (Optional, String) The global catalog deployment location.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\s]*$/`.
		* `location_url` - (Optional, String) The global catalog deployment URL of location.
		  * Constraints: The maximum length is `2083` characters. The minimum length is `1` character. The value must match regular expression `/^(?!mailto:)(?:(?:http|https|ftp):\/\/)(?:\\S+(?::\\S*)?@)?(?:(?:(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}(?:\\.(?:[0-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))|(?:(?:[a-z\\u00a1-\\uffff0-9]+-?)*[a-z\\u00a1-\\uffff0-9]+)(?:\\.(?:[a-z\\u00a1-\\uffff0-9]+-?)*[a-z\\u00a1-\\uffff0-9]+)*(?:\\.(?:[a-z\\u00a1-\\uffff]{2,})))|localhost)(?::\\d{2,5})?(?:(\/|\\?|#)[^\\s]*)?$/`.
		* `target_crn` - (Optional, String) Region crn.
		  * Constraints: The maximum length is `2000` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\s]*$/`.
	* `rc_compatible` - (Optional, Boolean) Whether the object is compatible with the resource controller service.
	* `service` - (Optional, List) The global catalog metadata of the service.
	Nested schema for **service**:
		* `bindable` - (Computed, Boolean) Deprecated. Controls the Connections tab on the Resource Details page.
		* `iam_compatible` - (Optional, Boolean) Whether the service is compatible with the IAM service.
		* `parameters` - (Optional, List)
		  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
		Nested schema for **parameters**:
			* `associations` - (Optional, Map) A JSON structure to describe the interactions with pricing plans and/or other custom parameters.
			* `description` - (Optional, String) The description of the parameter that is displayed to help users with the value of the parameter.
			  * Constraints: The maximum length is `2000` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\s]*$/`.
			* `displayname` - (Optional, String) The display name for custom service parameters.
			  * Constraints: The maximum length is `2000` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\s]*$/`.
			* `hidden` - (Optional, Boolean) Indicates whether the custom parameters is hidden required or not.
			* `i18n` - (Optional, List) The description for the object.
			Nested schema for **i18n**:
				* `de` - (Optional, List) The translations for custom service parameter display name and description.
				Nested schema for **de**:
					* `description` - (Optional, String) The translations for custom service parameter description.
					  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `displayname` - (Optional, String) The translations for custom service parameter display name.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
				* `en` - (Optional, List) The translations for custom service parameter display name and description.
				Nested schema for **en**:
					* `description` - (Optional, String) The translations for custom service parameter description.
					  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `displayname` - (Optional, String) The translations for custom service parameter display name.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
				* `es` - (Optional, List) The translations for custom service parameter display name and description.
				Nested schema for **es**:
					* `description` - (Optional, String) The translations for custom service parameter description.
					  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `displayname` - (Optional, String) The translations for custom service parameter display name.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
				* `fr` - (Optional, List) The translations for custom service parameter display name and description.
				Nested schema for **fr**:
					* `description` - (Optional, String) The translations for custom service parameter description.
					  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `displayname` - (Optional, String) The translations for custom service parameter display name.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
				* `it` - (Optional, List) The translations for custom service parameter display name and description.
				Nested schema for **it**:
					* `description` - (Optional, String) The translations for custom service parameter description.
					  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `displayname` - (Optional, String) The translations for custom service parameter display name.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
				* `ja` - (Optional, List) The translations for custom service parameter display name and description.
				Nested schema for **ja**:
					* `description` - (Optional, String) The translations for custom service parameter description.
					  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `displayname` - (Optional, String) The translations for custom service parameter display name.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
				* `ko` - (Optional, List) The translations for custom service parameter display name and description.
				Nested schema for **ko**:
					* `description` - (Optional, String) The translations for custom service parameter description.
					  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `displayname` - (Optional, String) The translations for custom service parameter display name.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
				* `pt_br` - (Optional, List) The translations for custom service parameter display name and description.
				Nested schema for **pt_br**:
					* `description` - (Optional, String) The translations for custom service parameter description.
					  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `displayname` - (Optional, String) The translations for custom service parameter display name.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
				* `zh_cn` - (Optional, List) The translations for custom service parameter display name and description.
				Nested schema for **zh_cn**:
					* `description` - (Optional, String) The translations for custom service parameter description.
					  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `displayname` - (Optional, String) The translations for custom service parameter display name.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
				* `zh_tw` - (Optional, List) The translations for custom service parameter display name and description.
				Nested schema for **zh_tw**:
					* `description` - (Optional, String) The translations for custom service parameter description.
					  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `displayname` - (Optional, String) The translations for custom service parameter display name.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
			* `invalidmessage` - (Optional, String) The message that appears when the content of the text box is invalid.
			  * Constraints: The maximum length is `2000` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\s]*$/`.
			* `layout` - (Optional, String) Specifies the layout of check box or radio input types. When unspecified, the default layout is horizontal.
			  * Constraints: The maximum length is `2000` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\s]*$/`.
			* `name` - (Optional, String) The key of the parameter.
			  * Constraints: The maximum length is `2000` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\s]*$/`.
			* `options` - (Optional, List)
			  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
			Nested schema for **options**:
				* `displayname` - (Optional, String) The display name for custom service parameters.
				  * Constraints: The maximum length is `2000` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\s]*$/`.
				* `i18n` - (Optional, List) The description for the object.
				Nested schema for **i18n**:
					* `de` - (Optional, List) The translations for custom service parameter display name and description.
					Nested schema for **de**:
						* `description` - (Optional, String) The translations for custom service parameter description.
						  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
						* `displayname` - (Optional, String) The translations for custom service parameter display name.
						  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `en` - (Optional, List) The translations for custom service parameter display name and description.
					Nested schema for **en**:
						* `description` - (Optional, String) The translations for custom service parameter description.
						  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
						* `displayname` - (Optional, String) The translations for custom service parameter display name.
						  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `es` - (Optional, List) The translations for custom service parameter display name and description.
					Nested schema for **es**:
						* `description` - (Optional, String) The translations for custom service parameter description.
						  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
						* `displayname` - (Optional, String) The translations for custom service parameter display name.
						  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `fr` - (Optional, List) The translations for custom service parameter display name and description.
					Nested schema for **fr**:
						* `description` - (Optional, String) The translations for custom service parameter description.
						  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
						* `displayname` - (Optional, String) The translations for custom service parameter display name.
						  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `it` - (Optional, List) The translations for custom service parameter display name and description.
					Nested schema for **it**:
						* `description` - (Optional, String) The translations for custom service parameter description.
						  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
						* `displayname` - (Optional, String) The translations for custom service parameter display name.
						  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `ja` - (Optional, List) The translations for custom service parameter display name and description.
					Nested schema for **ja**:
						* `description` - (Optional, String) The translations for custom service parameter description.
						  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
						* `displayname` - (Optional, String) The translations for custom service parameter display name.
						  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `ko` - (Optional, List) The translations for custom service parameter display name and description.
					Nested schema for **ko**:
						* `description` - (Optional, String) The translations for custom service parameter description.
						  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
						* `displayname` - (Optional, String) The translations for custom service parameter display name.
						  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `pt_br` - (Optional, List) The translations for custom service parameter display name and description.
					Nested schema for **pt_br**:
						* `description` - (Optional, String) The translations for custom service parameter description.
						  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
						* `displayname` - (Optional, String) The translations for custom service parameter display name.
						  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `zh_cn` - (Optional, List) The translations for custom service parameter display name and description.
					Nested schema for **zh_cn**:
						* `description` - (Optional, String) The translations for custom service parameter description.
						  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
						* `displayname` - (Optional, String) The translations for custom service parameter display name.
						  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `zh_tw` - (Optional, List) The translations for custom service parameter display name and description.
					Nested schema for **zh_tw**:
						* `description` - (Optional, String) The translations for custom service parameter description.
						  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
						* `displayname` - (Optional, String) The translations for custom service parameter display name.
						  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
				* `value` - (Optional, String) The value for custom service parameters.
				  * Constraints: The maximum length is `2000` characters. The minimum length is `1` character. The value must match regular expression `/./`.
			* `options_url` - (Optional, String) The options URL for custom service parameters.
			  * Constraints: The maximum length is `2083` characters. The minimum length is `1` character. The value must match regular expression `/^(?!mailto:)(?:(?:http|https|ftp):\/\/)(?:\\S+(?::\\S*)?@)?(?:(?:(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}(?:\\.(?:[0-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))|(?:(?:[a-z\\u00a1-\\uffff0-9]+-?)*[a-z\\u00a1-\\uffff0-9]+)(?:\\.(?:[a-z\\u00a1-\\uffff0-9]+-?)*[a-z\\u00a1-\\uffff0-9]+)*(?:\\.(?:[a-z\\u00a1-\\uffff]{2,})))|localhost)(?::\\d{2,5})?(?:(\/|\\?|#)[^\\s]*)?$/`.
			* `pattern` - (Optional, String) A regular expression that the value is checked against.
			  * Constraints: The maximum length is `2000` characters. The minimum length is `1` character. The value must match regular expression `/./`.
			* `placeholder` - (Optional, String) The placeholder text for custom parameters.
			  * Constraints: The maximum length is `2000` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\s]*$/`.
			* `readonly` - (Optional, Boolean) A boolean value that indicates whether the value of the parameter is displayed only and cannot be changed by users. The default value is false.
			* `required` - (Optional, Boolean) A boolean value that indicates whether the parameter must be entered in the IBM Cloud user interface.
			* `type` - (Optional, String) The type of custom service parameters.
			  * Constraints: Allowable values are: `text`, `textarea`, `dropdown`, `number`, `password`, `combo`, `checkbox`, `radio`, `multiselect`, `resource_group`, `vcenter_datastore`, `region`, `secret`, `cluster_namespace`.
			* `validation_url` - (Optional, String) The validation URL for custom service parameters.
			  * Constraints: The maximum length is `2083` characters. The minimum length is `1` character. The value must match regular expression `/^(?!mailto:)(?:(?:http|https|ftp):\/\/)(?:\\S+(?::\\S*)?@)?(?:(?:(?:[1-9]\\d?|1\\d\\d|2[01]\\d|22[0-3])(?:\\.(?:1?\\d{1,2}|2[0-4]\\d|25[0-5])){2}(?:\\.(?:[0-9]\\d?|1\\d\\d|2[0-4]\\d|25[0-4]))|(?:(?:[a-z\\u00a1-\\uffff0-9]+-?)*[a-z\\u00a1-\\uffff0-9]+)(?:\\.(?:[a-z\\u00a1-\\uffff0-9]+-?)*[a-z\\u00a1-\\uffff0-9]+)*(?:\\.(?:[a-z\\u00a1-\\uffff]{2,})))|localhost)(?::\\d{2,5})?(?:(\/|\\?|#)[^\\s]*)?$/`.
			* `value` - (Optional, List)
			  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `1000` items. The minimum length is `0` items.
		* `plan_updateable` - (Computed, Boolean) Indicates plan update support and controls the Plan tab on the Resource Details page.
		* `rc_provisionable` - (Optional, Boolean) Whether the service is provisionable by the resource controller service.
		* `service_key_supported` - (Optional, Boolean) Indicates service credentials support and controls the Service Credential tab on Resource Details page.
		* `unique_api_key` - (Computed, Boolean) Indicates whether the deployment uses a unique API key or not.
* `name` - (Required, String) The programmatic name of this deployment.
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
* `tags` - (Optional, List) A list of tags that carry information about your product. These tags can be used to find your product in the IBM Cloud catalog.
  * Constraints: The list items must match regular expression `/^[a-z0-9\\-._]+$/`. The maximum length is `100` items. The minimum length is `0` items.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the onboarding_catalog_deployment.
* `catalog_deployment_id` - (String) The ID of a global catalog object.
* `geo_tags` - (List) 
  * Constraints: The list items must match regular expression `/./`. The maximum length is `1000` items. The minimum length is `0` items.
* `url` - (String) The global catalog URL of your product.


## Import

You can import the `ibm_onboarding_catalog_deployment` resource by using `id`.
The `id` property can be formed from `product_id`, `catalog_product_id`, `catalog_plan_id`, and `catalog_deployment_id` in the following format:

<pre>
product_id/catalog_product_id/catalog_plan_id/catalog_deployment_id</pre>
* `product_id`: A string. The unique ID of the product.
* `catalog_product_id`: A string. The unique ID of this global catalog product.
* `catalog_plan_id`: A string. The unique ID of this global catalog plan.
* `catalog_deployment_id`: A string. The ID of a global catalog object.

# Syntax
<pre>
$ terraform import ibm_onboarding_catalog_deployment.onboarding_catalog_deployment product_id/catalog_product_id/catalog_plan_id/catalog_deployment_id
</pre>
