---
layout: "ibm"
page_title: "IBM : ibm_onboarding_iam_registration"
description: |-
  Manages onboarding_iam_registration.
subcategory: "Partner Center Sell"
---

# ibm_onboarding_iam_registration

Create, update, and delete onboarding_iam_registrations with this resource.

## Example Usage

```hcl
resource "ibm_onboarding_iam_registration" "onboarding_iam_registration_instance" {
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
  enabled = true
  name = "pet-store"
  product_id = ibm_onboarding_product.onboarding_product_instance.id
  resource_hierarchy_attribute {
		key = "key"
		value = "value"
  }
  supported_anonymous_accesses {
		attributes {
			account_id = "account_id"
			service_name = "service_name"
			additional_properties = { "key" = "inner" }
		}
		roles = [ "roles" ]
  }
  supported_attributes {
		key = "key"
		options {
			operators = [ "stringEquals" ]
			hidden = true
			supported_patterns = [ "supported_patterns" ]
			policy_types = [ "access" ]
			is_empty_value_supported = true
			is_string_exists_false_value_supported = true
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
					value_property_name = "value_property_name"
					label_property_name = "label_property_name"
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
			access_policy = true
			policy_type = [ "access" ]
			account_type = "enterprise"
		}
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `actions` - (Optional, List) The product access management action.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **actions**:
	* `description` - (Optional, List) The description for the object.
	Nested schema for **description**:
		* `de` - (Optional, String) German.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `default` - (Optional, String) The fallback string for the description object.
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
	* `display_name` - (Optional, List) The display name of the object.
	Nested schema for **display_name**:
		* `de` - (Optional, String) German.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `default` - (Optional, String) The fallback string for the description object.
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
	* `id` - (Optional, String) The unique identifier for the action.
	  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `options` - (Optional, List) Extra options.
	Nested schema for **options**:
		* `hidden` - (Optional, Boolean) Optional opt-in if action is hidden from customers.
	* `roles` - (Optional, List) The list of roles for the action.
	  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `additional_policy_scopes` - (Optional, List) List of additional policy scopes.
  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `display_name` - (Optional, List) The display name of the object.
Nested schema for **display_name**:
	* `de` - (Optional, String) German.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
	* `default` - (Optional, String) The fallback string for the description object.
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
* `enabled` - (Optional, Boolean) Whether the service is enabled or disabled for IAM.
* `env` - (Optional, String) The environment to fetch this object from.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z_.-]+$/`.
* `name` - (Required, String) The IAM registration name, which must be the programmatic name of the product.
  * Constraints: The value must match regular expression `/^\\S*$/`.
* `parent_ids` - (Optional, List) The list of parent IDs for product access management.
  * Constraints: The list items must match regular expression `/^\\S*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `product_id` - (Required, Forces new resource, String) The unique ID of the product.
  * Constraints: The maximum length is `71` characters. The minimum length is `71` characters. The value must match regular expression `/^[a-zA-Z0-9]{32}:o:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `resource_hierarchy_attribute` - (Optional, List) The resource hierarchy key-value pair for composite services.
Nested schema for **resource_hierarchy_attribute**:
	* `key` - (Optional, String) The resource hierarchy key.
	* `value` - (Optional, String) The resource hierarchy value.
* `service_type` - (Optional, String) The type of the service.
  * Constraints: Allowable values are: `service`, `platform_service`.
* `supported_anonymous_accesses` - (Optional, List) The list of supported anonymous accesses.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **supported_anonymous_accesses**:
	* `attributes` - (Optional, List) The attributes for anonymous accesses.
	Nested schema for **attributes**:
		* `account_id` - (Required, String) An account id.
		* `additional_properties` - (Required, Map) Additional properties the key must come from supported attributes.
		* `service_name` - (Required, String) The name of the service.
	* `roles` - (Optional, List) The roles of supported anonymous accesses.
	  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `supported_attributes` - (Optional, List) The list of supported attributes.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **supported_attributes**:
	* `description` - (Optional, List) The description for the object.
	Nested schema for **description**:
		* `de` - (Optional, String) German.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `default` - (Optional, String) The fallback string for the description object.
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
	* `display_name` - (Optional, List) The display name of the object.
	Nested schema for **display_name**:
		* `de` - (Optional, String) German.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `default` - (Optional, String) The fallback string for the description object.
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
	* `key` - (Optional, String) The supported attribute key.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `options` - (Optional, List) The list of support attribute options.
	Nested schema for **options**:
		* `hidden` - (Optional, Boolean) Optional opt-in if attribute is hidden from customers (customer can still use it if they found out themselves).
		* `is_empty_value_supported` - (Optional, Boolean) Indicate whether the empty value is supported.
		* `is_string_exists_false_value_supported` - (Optional, Boolean) Indicate whether the false value is supported for stringExists operator.
		* `key` - (Optional, String) The name of attribute.
		* `operators` - (Optional, List) The supported attribute operator.
		  * Constraints: Allowable list items are: `stringEquals`, `stringMatch`, `stringEqualsAnyOf`, `stringMatchAnyOf`. The maximum length is `100` items. The minimum length is `0` items.
		* `policy_types` - (Optional, List) The list of policy types.
		  * Constraints: Allowable list items are: `access`, `authorization`. The maximum length is `2` items. The minimum length is `1` item.
		* `resource_hierarchy` - (Optional, List) Resource hierarchy options for composite services.
		Nested schema for **resource_hierarchy**:
			* `key` - (Optional, List) Hierarchy description key.
			Nested schema for **key**:
				* `key` - (Optional, String) Key.
				* `value` - (Optional, String) Value.
			* `value` - (Optional, List) Hierarchy description value.
			Nested schema for **value**:
				* `key` - (Optional, String) Key.
		* `supported_patterns` - (Optional, List) The list of supported patterns.
		  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
	* `ui` - (Optional, List) The user interface.
	Nested schema for **ui**:
		* `input_details` - (Optional, List) The details of the input.
		Nested schema for **input_details**:
			* `gst` - (Optional, List) Required if type is gst.
			Nested schema for **gst**:
				* `input_option_label` - (Optional, String) The label for option input.
				  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
				* `label_property_name` - (Optional, String) One of labelPropertyName or inputOptionLabel is required.
				* `query` - (Optional, String) The query to use.
				  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
				* `value_property_name` - (Optional, String) The value of the property name.
				  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
			* `type` - (Optional, String) They type of the input details.
			  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
			* `url` - (Optional, List) The URL data for user interface.
			Nested schema for **url**:
				* `input_option_label` - (Optional, String) The label options for the user interface URL.
				  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
				* `url_endpoint` - (Optional, String) The URL of the user interface interface.
				  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
			* `values` - (Optional, List) The provided values of input details.
			  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
			Nested schema for **values**:
				* `display_name` - (Optional, List) The display name of the object.
				Nested schema for **display_name**:
					* `de` - (Optional, String) German.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
					* `default` - (Optional, String) The fallback string for the description object.
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
				* `value` - (Optional, String) The values of input details.
				  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `input_type` - (Optional, String) The type of the input.
		  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `supported_authorization_subjects` - (Optional, List) The list of supported authorization subjects.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **supported_authorization_subjects**:
	* `attributes` - (Optional, List) The list of supported authorization subject properties.
	Nested schema for **attributes**:
		* `resource_type` - (Optional, String) The type of the service.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `service_name` - (Optional, String) The name of the service.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `roles` - (Optional, List) The list of roles for authorization.
	  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `supported_network` - (Optional, List) The registration of set of endpoint types that are supported by your service in the `networkType` environment attribute. This constrains the context-based restriction rules specific to the service such that they describe access restrictions on only this set of endpoints.
Nested schema for **supported_network**:
	* `environment_attributes` - (Optional, List) The environment attribute for support.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
	Nested schema for **environment_attributes**:
		* `key` - (Optional, String) The name of the key.
		  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `options` - (Optional, List) The list of options for supported networks.
		Nested schema for **options**:
			* `hidden` - (Optional, Boolean) Whether the attribute is hidden or not.
		* `values` - (Optional, List) The list of values that belong to the key.
		  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `supported_roles` - (Optional, List) The list of roles that you can use to assign access.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **supported_roles**:
	* `description` - (Optional, List) The description for the object.
	Nested schema for **description**:
		* `de` - (Optional, String) German.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `default` - (Optional, String) The fallback string for the description object.
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
	* `display_name` - (Optional, List) The display name of the object.
	Nested schema for **display_name**:
		* `de` - (Optional, String) German.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wäöüß\\d]+\\b/`.
		* `default` - (Optional, String) The fallback string for the description object.
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
	* `id` - (Optional, String) The value belonging to the key.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `options` - (Optional, List) The supported role options.
	Nested schema for **options**:
		* `access_policy` - (Required, Boolean) Optional opt-in to require access control on the role.
		* `account_type` - (Optional, String) Optional opt-in to require checking account type when applying the role.
		  * Constraints: Allowable values are: `enterprise`.
		* `policy_type` - (Optional, List) Optional opt-in to require checking policy type when applying the role.
		  * Constraints: Allowable list items are: `access`, `authorization`, `authorization-delegated`. The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the onboarding_iam_registration.


## Import

You can import the `ibm_onboarding_iam_registration` resource by using `name`.
The `name` property can be formed from `product_id`, and `name` in the following format:

<pre>
&lt;product_id&gt;/&lt;name&gt;
</pre>
* `product_id`: A string. The unique ID of the product.
* `name`: A string in the format `pet-store`. The IAM registration name, which must be the programmatic name of the product.

# Syntax
<pre>
$ terraform import ibm_onboarding_iam_registration.onboarding_iam_registration &lt;product_id&gt;/&lt;name&gt;
</pre>
