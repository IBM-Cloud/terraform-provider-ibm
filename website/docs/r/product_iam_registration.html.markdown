---
layout: "ibm"
page_title: "IBM : ibm_product_iam_registration"
description: |-
  Manages product_iam_registration.
subcategory: "Partner Center Sell"
---

# ibm_product_iam_registration

Create, update, and delete product_iam_registrations with this resource.

## Example Usage

```hcl
resource "ibm_product_iam_registration" "product_iam_registration_instance" {
  actions {
		id = "id"
		roles = [ "roles" ]
		description {
			default = "default"
			en = "en"
			de = "de"
			es = "es"
			fr = "fr"
			it = "it"
			ja = "ja"
			ko = "ko"
			pt_br = "pt_br"
			zh_tw = "zh_tw"
			zh_cn = "zh_cn"
		}
		display_name {
			default = "default"
			en = "en"
			de = "de"
			es = "es"
			fr = "fr"
			it = "it"
			ja = "ja"
			ko = "ko"
			pt_br = "pt_br"
			zh_tw = "zh_tw"
			zh_cn = "zh_cn"
		}
		options {
			hidden = true
		}
  }
  display_name {
		default = "default"
		en = "en"
		de = "de"
		es = "es"
		fr = "fr"
		it = "it"
		ja = "ja"
		ko = "ko"
		pt_br = "pt_br"
		zh_tw = "zh_tw"
		zh_cn = "zh_cn"
  }
  resource_hierarchy_attribute {
		key = "key"
		value = "value"
  }
  supported_anonymous_accesses {
		attributes = { "key" = "anything as a string" }
		roles = [ "roles" ]
  }
  supported_attributes {
		key = "key"
		options {
			operators = [ "operators" ]
			hidden = true
			key = "key"
			resource_hierarchy {
				key {
					key = "key"
					value = "value"
				}
				value {
					key = "key"
				}
			}
		}
		display_name {
			default = "default"
			en = "en"
			de = "de"
			es = "es"
			fr = "fr"
			it = "it"
			ja = "ja"
			ko = "ko"
			pt_br = "pt_br"
			zh_tw = "zh_tw"
			zh_cn = "zh_cn"
		}
		description {
			default = "default"
			en = "en"
			de = "de"
			es = "es"
			fr = "fr"
			it = "it"
			ja = "ja"
			ko = "ko"
			pt_br = "pt_br"
			zh_tw = "zh_tw"
			zh_cn = "zh_cn"
		}
		ui {
			input_type = "input_type"
			input_details {
				type = "type"
				values {
					value = "value"
					display_name {
						default = "default"
						en = "en"
						de = "de"
						es = "es"
						fr = "fr"
						it = "it"
						ja = "ja"
						ko = "ko"
						pt_br = "pt_br"
						zh_tw = "zh_tw"
						zh_cn = "zh_cn"
					}
				}
				gst {
					query = "query"
					label_property_name = "label_property_name"
					value_property_name = "value_property_name"
					input_option_label = "input_option_label"
				}
				url {
					url_endpoint = "url_endpoint"
					input_option_label = "input_option_label"
				}
			}
		}
  }
  supported_authorization_subjects {
		attributes {
			service_name = "service_name"
			resource_type = "resource_type"
		}
		roles = [ "roles" ]
  }
  supported_network {
		environment_attributes {
			key = "key"
			values = [ "values" ]
			options {
				hidden = true
			}
		}
  }
  supported_roles {
		id = "id"
		description {
			default = "default"
			en = "en"
			de = "de"
			es = "es"
			fr = "fr"
			it = "it"
			ja = "ja"
			ko = "ko"
			pt_br = "pt_br"
			zh_tw = "zh_tw"
			zh_cn = "zh_cn"
		}
		display_name {
			default = "default"
			en = "en"
			de = "de"
			es = "es"
			fr = "fr"
			it = "it"
			ja = "ja"
			ko = "ko"
			pt_br = "pt_br"
			zh_tw = "zh_tw"
			zh_cn = "zh_cn"
		}
		options {
			access_policy = { "key" = "anything as a string" }
			policy_type = [ "policy_type" ]
		}
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `actions` - (Optional, List) Product access management action.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **actions**:
	* `description` - (Optional, List) Description for the object.
	Nested schema for **description**:
		* `de` - (Optional, String) German.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `default` - (Optional, String) The fallback string for description object.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `en` - (Optional, String) English.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `es` - (Optional, String) Spanish.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wáéíóúñ]+\\b/`.
		* `fr` - (Optional, String) French.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
		* `it` - (Optional, String) Italian.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `ja` - (Optional, String) Japanese.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `ko` - (Optional, String) Korean.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `pt_br` - (Optional, String) Portuguese (Brazil).
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `zh_cn` - (Optional, String) Simplified Chinese.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `zh_tw` - (Optional, String) Traditional Chinese.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `display_name` - (Optional, List) Display name for the object.
	Nested schema for **display_name**:
		* `de` - (Optional, String) German.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `default` - (Optional, String) The fallback string for description object.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `en` - (Optional, String) English.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `es` - (Optional, String) Spanish.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wáéíóúñ]+\\b/`.
		* `fr` - (Optional, String) French.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
		* `it` - (Optional, String) Italian.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `ja` - (Optional, String) Japanese.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `ko` - (Optional, String) Korean.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `pt_br` - (Optional, String) Portuguese (Brazil).
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `zh_cn` - (Optional, String) Simplified Chinese.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `zh_tw` - (Optional, String) Traditional Chinese.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `id` - (Optional, String) Unique identifier for the action.
	  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `options` - (Optional, List) Extra options.
	Nested schema for **options**:
		* `hidden` - (Optional, Boolean)
	* `roles` - (Optional, List) List of roles for the the action.
	  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `additional_policy_scopes` - (Optional, List) List of additional policy scopes.
  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `display_name` - (Optional, List) Display name for the object.
Nested schema for **display_name**:
	* `de` - (Optional, String) German.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
	* `default` - (Optional, String) The fallback string for description object.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `en` - (Optional, String) English.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `es` - (Optional, String) Spanish.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wáéíóúñ]+\\b/`.
	* `fr` - (Optional, String) French.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
	* `it` - (Optional, String) Italian.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
	* `ja` - (Optional, String) Japanese.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
	* `ko` - (Optional, String) Korean.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `pt_br` - (Optional, String) Portuguese (Brazil).
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
	* `zh_cn` - (Optional, String) Simplified Chinese.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `zh_tw` - (Optional, String) Traditional Chinese.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
* `enabled` - (Optional, Boolean) Additional policies enabled or disabled.
* `name` - (Optional, String) IAM programmatic name.
  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `parent_ids` - (Optional, List) List of parent ids for product access management.
  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `resource_hierarchy_attribute` - (Optional, List) Resource hierarchy key value pair for composite services.
Nested schema for **resource_hierarchy_attribute**:
	* `key` - (Optional, String) Resource hierarchy key.
	* `value` - (Optional, String) Resource hierarchy value.
* `service_type` - (Optional, String) 
* `supported_anonymous_accesses` - (Optional, List) Support for anonymous accesses.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **supported_anonymous_accesses**:
	* `attributes` - (Optional, Map) Support for anonymous accesses.
	* `roles` - (Optional, List) Roles of supported anonymous accesses.
	  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `supported_attributes` - (Optional, List) List of supported attribute.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **supported_attributes**:
	* `description` - (Optional, List) Description for the object.
	Nested schema for **description**:
		* `de` - (Optional, String) German.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `default` - (Optional, String) The fallback string for description object.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `en` - (Optional, String) English.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `es` - (Optional, String) Spanish.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wáéíóúñ]+\\b/`.
		* `fr` - (Optional, String) French.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
		* `it` - (Optional, String) Italian.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `ja` - (Optional, String) Japanese.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `ko` - (Optional, String) Korean.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `pt_br` - (Optional, String) Portuguese (Brazil).
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `zh_cn` - (Optional, String) Simplified Chinese.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `zh_tw` - (Optional, String) Traditional Chinese.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `display_name` - (Optional, List) Display name for the object.
	Nested schema for **display_name**:
		* `de` - (Optional, String) German.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `default` - (Optional, String) The fallback string for description object.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `en` - (Optional, String) English.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `es` - (Optional, String) Spanish.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wáéíóúñ]+\\b/`.
		* `fr` - (Optional, String) French.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
		* `it` - (Optional, String) Italian.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `ja` - (Optional, String) Japanese.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `ko` - (Optional, String) Korean.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `pt_br` - (Optional, String) Portuguese (Brazil).
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `zh_cn` - (Optional, String) Simplified Chinese.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `zh_tw` - (Optional, String) Traditional Chinese.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `key` - (Optional, String) Support attribute key.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `options` - (Optional, List) Support attribute options.
	Nested schema for **options**:
		* `hidden` - (Optional, Boolean)
		* `key` - (Optional, String)
		* `operators` - (Optional, List) Supported attribute operator.
		  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
		* `resource_hierarchy` - (Optional, List)
		Nested schema for **resource_hierarchy**:
			* `key` - (Optional, List)
			Nested schema for **key**:
				* `key` - (Optional, String)
				* `value` - (Optional, String)
			* `value` - (Optional, List)
			Nested schema for **value**:
				* `key` - (Optional, String)
	* `ui` - (Optional, List) User interface.
	Nested schema for **ui**:
		* `input_details` - (Optional, List) Details of the input.
		Nested schema for **input_details**:
			* `gst` - (Optional, List) Group Security Token.
			Nested schema for **gst**:
				* `input_option_label` - (Optional, String) Label for potion input.
				  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
				* `label_property_name` - (Optional, String)
				* `query` - (Optional, String) Query to use.
				  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
				* `value_property_name` - (Optional, String) Value of the property name.
				  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
			* `type` - (Optional, String) type of input details.
			  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
			* `url` - (Optional, List) Url data for User interface.
			Nested schema for **url**:
				* `input_option_label` - (Optional, String) Label options for the user interface url.
				  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
				* `url_endpoint` - (Optional, String) Url itself for the interface.
				  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
			* `values` - (Optional, List) Provided values of input details.
			  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
			Nested schema for **values**:
				* `display_name` - (Optional, List) Display name for the object.
				Nested schema for **display_name**:
					* `de` - (Optional, String) German.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
					* `default` - (Optional, String) The fallback string for description object.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
					* `en` - (Optional, String) English.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
					* `es` - (Optional, String) Spanish.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wáéíóúñ]+\\b/`.
					* `fr` - (Optional, String) French.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
					* `it` - (Optional, String) Italian.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
					* `ja` - (Optional, String) Japanese.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
					* `ko` - (Optional, String) Korean.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `pt_br` - (Optional, String) Portuguese (Brazil).
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
					* `zh_cn` - (Optional, String) Simplified Chinese.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `zh_tw` - (Optional, String) Traditional Chinese.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
				* `value` - (Optional, String) Values of input details.
				  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `input_type` - (Optional, String) Type of the input.
		  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `supported_authorization_subjects` - (Optional, List) List of supported authorization subjects.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **supported_authorization_subjects**:
	* `attributes` - (Optional, List) Supported authorization subject properties.
	Nested schema for **attributes**:
		* `resource_type` - (Optional, String) Child of the service.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `service_name` - (Optional, String) Name of the service.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `roles` - (Optional, List) Roles for authorization.
	  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `supported_network` - (Optional, List) Context-based restrictions (CBR).
Nested schema for **supported_network**:
	* `environment_attributes` - (Optional, List) Environment attribute for support.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
	Nested schema for **environment_attributes**:
		* `key` - (Optional, String) Name of the key.
		  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `options` - (Optional, List) options for supported networks.
		Nested schema for **options**:
			* `hidden` - (Optional, Boolean) Should the attribute be shown or not.
		* `values` - (Optional, List) List of values belonging to the key.
		  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `supported_roles` - (Optional, List) Roles you can choose from.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **supported_roles**:
	* `description` - (Optional, List) Description for the object.
	Nested schema for **description**:
		* `de` - (Optional, String) German.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `default` - (Optional, String) The fallback string for description object.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `en` - (Optional, String) English.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `es` - (Optional, String) Spanish.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wáéíóúñ]+\\b/`.
		* `fr` - (Optional, String) French.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
		* `it` - (Optional, String) Italian.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `ja` - (Optional, String) Japanese.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `ko` - (Optional, String) Korean.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `pt_br` - (Optional, String) Portuguese (Brazil).
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `zh_cn` - (Optional, String) Simplified Chinese.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `zh_tw` - (Optional, String) Traditional Chinese.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `display_name` - (Optional, List) Display name for the object.
	Nested schema for **display_name**:
		* `de` - (Optional, String) German.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `default` - (Optional, String) The fallback string for description object.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `en` - (Optional, String) English.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `es` - (Optional, String) Spanish.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wáéíóúñ]+\\b/`.
		* `fr` - (Optional, String) French.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàâçéèêëîïôûùüÿñœæ]+\\b/`.
		* `it` - (Optional, String) Italian.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `ja` - (Optional, String) Japanese.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `ko` - (Optional, String) Korean.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `pt_br` - (Optional, String) Portuguese (Brazil).
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `zh_cn` - (Optional, String) Simplified Chinese.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `zh_tw` - (Optional, String) Traditional Chinese.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `id` - (Optional, String) Value belonging to the key.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `options` - (Optional, List) Supported role options.
	Nested schema for **options**:
		* `access_policy` - (Optional, Map) Role is access managed.
		* `policy_type` - (Optional, List) Policy types where this role could be used.
		  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the product_iam_registration.


## Import

You can import the `ibm_product_iam_registration` resource by using `name`. IAM programmatic name.

# Syntax
<pre>
$ terraform import ibm_product_iam_registration.product_iam_registration &lt;name&gt;
</pre>
