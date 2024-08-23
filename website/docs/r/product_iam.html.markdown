---
layout: "ibm"
page_title: "IBM : ibm_product_iam"
description: |-
  Manages product_iam.
subcategory: "Partner Center Sell"
---

# ibm_product_iam

Provides a resource for product_iam. This allows product_iam to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_product_iam" "product_iam_instance" {
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
  product_id = "product_id"
  supported_anonymous_accesses {
		attributes = { "key" = "anything as a string" }
		roles = [ "roles" ]
  }
  supported_attributes {
		key = "key"
		options {
			operators = [ "operators" ]
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
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `actions` - (Optional, List) Product access management action.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested scheme for **actions**:
	* `description` - (Optional, List) Description for the object.
	Nested scheme for **description**:
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
		* `ja` - (Optional, String) Japan.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `ko` - (Optional, String) Korean.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `pt_br` - (Optional, String) Portuguese, brazilian.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `zh_cn` - (Optional, String) Chines.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `zh_tw` - (Optional, String) Chines taiwan.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `display_name` - (Optional, List) Display name for the object.
	Nested scheme for **display_name**:
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
		* `ja` - (Optional, String) Japan.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `ko` - (Optional, String) Korean.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `pt_br` - (Optional, String) Portuguese, brazilian.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `zh_cn` - (Optional, String) Chines.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `zh_tw` - (Optional, String) Chines taiwan.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `id` - (Optional, String) Unique identifier for the action.
	  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `roles` - (Optional, List) List of roles for the the action.
	  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `additional_policy_scopes` - (Optional, List) List of additional policies.
  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `display_name` - (Optional, List) Display name for the object.
Nested scheme for **display_name**:
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
	* `ja` - (Optional, String) Japan.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
	* `ko` - (Optional, String) Korean.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `pt_br` - (Optional, String) Portuguese, brazilian.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
	* `zh_cn` - (Optional, String) Chines.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `zh_tw` - (Optional, String) Chines taiwan.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
* `enabled` - (Optional, Boolean) Additional policies enabled or disabled.
* `name` - (Optional, String) Parent ids for product access management.
  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `parent_ids` - (Optional, List) List of parent ids for product access management.
  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `product_id` - (Required, Forces new resource, String) The unique ID of the product. This ID can be obtained by calling the list products method and also can be found in Partner Center.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/\\b\\s*([a-f\\d\\\\-]*){1}\\s*/`.
* `supported_anonymous_accesses` - (Optional, List) Support for anonymous accesses.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested scheme for **supported_anonymous_accesses**:
	* `attributes` - (Optional, Map) Support for anonymous accesses.
	* `roles` - (Optional, List) Roles of supported anonymous accesses.
	  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `supported_attributes` - (Optional, List) List of supported attribute.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested scheme for **supported_attributes**:
	* `description` - (Optional, List) Description for the object.
	Nested scheme for **description**:
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
		* `ja` - (Optional, String) Japan.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `ko` - (Optional, String) Korean.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `pt_br` - (Optional, String) Portuguese, brazilian.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `zh_cn` - (Optional, String) Chines.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `zh_tw` - (Optional, String) Chines taiwan.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `display_name` - (Optional, List) Display name for the object.
	Nested scheme for **display_name**:
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
		* `ja` - (Optional, String) Japan.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `ko` - (Optional, String) Korean.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `pt_br` - (Optional, String) Portuguese, brazilian.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `zh_cn` - (Optional, String) Chines.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `zh_tw` - (Optional, String) Chines taiwan.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `key` - (Optional, String) Support attribute key.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `options` - (Optional, List) Support attribute options.
	Nested scheme for **options**:
		* `operators` - (Optional, List) Supported attribute operator.
		  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
	* `ui` - (Optional, List) User interface.
	Nested scheme for **ui**:
		* `input_details` - (Optional, List) Details of the input.
		Nested scheme for **input_details**:
			* `gst` - (Optional, List) Group Security Token.
			Nested scheme for **gst**:
				* `input_option_label` - (Optional, String) Label for potion input.
				  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
				* `query` - (Optional, String) Query to use.
				  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
				* `value_property_name` - (Optional, String) Value of the property name.
				  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
			* `type` - (Optional, String) type of input details.
			  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
			* `url` - (Optional, List) Url data for User interface.
			Nested scheme for **url**:
				* `input_option_label` - (Optional, String) Label options for the user interface url.
				  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
				* `url_endpoint` - (Optional, String) Url itself for the interface.
				  * Constraints: The maximum length is `2083` characters. The minimum length is `0` characters.
			* `values` - (Optional, List) Provided values of input details.
			  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
			Nested scheme for **values**:
				* `display_name` - (Optional, List) Display name for the object.
				Nested scheme for **display_name**:
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
					* `ja` - (Optional, String) Japan.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
					* `ko` - (Optional, String) Korean.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `pt_br` - (Optional, String) Portuguese, brazilian.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
					* `zh_cn` - (Optional, String) Chines.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
					* `zh_tw` - (Optional, String) Chines taiwan.
					  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
				* `value` - (Optional, String) Values of input details.
				  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `input_type` - (Optional, String) Type of the input.
		  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
* `supported_authorization_subjects` - (Optional, List) List of supported authorization subjects.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested scheme for **supported_authorization_subjects**:
	* `attributes` - (Optional, List) Supported authorization subject properties.
	Nested scheme for **attributes**:
		* `service_name` - (Optional, String) Name of the service.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
	* `roles` - (Optional, List) Roles for authorization.
	  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `supported_network` - (Optional, List) Supported networks.
Nested scheme for **supported_network**:
	* `environment_attributes` - (Optional, List) Environment attribute for support.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
	Nested scheme for **environment_attributes**:
		* `key` - (Optional, String) Name of the key.
		  * Constraints: The maximum length is `100` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.
		* `options` - (Optional, List) options for supported networks.
		Nested scheme for **options**:
			* `hidden` - (Optional, Boolean) Should the attribute be shown or not.
		* `values` - (Optional, List) List of values belonging to the key.
		  * Constraints: The list items must match regular expression `/^[ -~\\s]*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `supported_roles` - (Optional, List) Roles you can choose from.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested scheme for **supported_roles**:
	* `description` - (Optional, List) Description for the object.
	Nested scheme for **description**:
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
		* `ja` - (Optional, String) Japan.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `ko` - (Optional, String) Korean.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `pt_br` - (Optional, String) Portuguese, brazilian.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `zh_cn` - (Optional, String) Chines.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `zh_tw` - (Optional, String) Chines taiwan.
		  * Constraints: The maximum length is `20000` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `display_name` - (Optional, List) Display name for the object.
	Nested scheme for **display_name**:
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
		* `ja` - (Optional, String) Japan.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Za-z0-9\\s,.!?;:'"-]+|[ぁ-んァ-ン一-龯、。「」！？\\d\\s]*$/`.
		* `ko` - (Optional, String) Korean.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `pt_br` - (Optional, String) Portuguese, brazilian.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/\\b[\\wàèéìîòóù]+\\b/`.
		* `zh_cn` - (Optional, String) Chines.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
		* `zh_tw` - (Optional, String) Chines taiwan.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/./`.
	* `id` - (Optional, String) Value belonging to the key.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^[ -~\\s]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the product_iam.

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

You can import the `ibm_product_iam` resource by using `name`.
The `name` property can be formed from `product_id`, and `service_name` in the following format:

```
<product_id>/<service_name>
```
* `product_id`: A string. The unique ID of the product. This ID can be obtained by calling the list products method and also can be found in Partner Center.
* `service_name`: A string. The unique programmatic name of the product. This name will be part of the C.R.N.

# Syntax
```
$ terraform import ibm_product_iam.product_iam <product_id>/<service_name>
```
