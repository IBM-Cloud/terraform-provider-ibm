---
layout: "ibm"
page_title: "IBM : ibm_cbr_rule"
description: |-
  Manages cbr_rule.
subcategory: "Context Based Restrictions"
---

# ibm_cbr_rule

Create, update, and delete cbr_rules with this resource.

## Example Usage

```hcl
resource "ibm_cbr_rule" "cbr_rule_instance" {
  contexts {
		attributes {
			name = "networkZoneId"
			value = "559052eb8f43302824e7ae490c0281eb, bf823d4f45b64ceaa4671bee0479346e"
		}
		attributes {
      		name  = "mfa"
      		value = "LEVEL1"
        }
		attributes {
       		name = "endpointType"
       		value = "private"
    }
  }
  description = "this is an example of rule with one context two zones"
  enforcement_mode = "enabled"
  operations {
		api_types {
			api_type_id = "api_type_id"
		}
  }
  resources {
		attributes {
			name = "accountId"
			value = "12ab34cd56ef78ab90cd12ef34ab56cd"
		}
		attributes {
      		name = "serviceName"
      		value = "network-policy-enabled"
    	}
		tags {
      		name     = "tag_name"
      		value    = "tag_value"
		}
  }
}
```

## Example Usage to create a rule with two contexts

```hcl
resource "ibm_cbr_rule" "cbr_rule_instance" {
  contexts {
		attributes {
			name = "networkZoneId"
			value = "559052eb8f43302824e7ae490c0281eb"
		}
		attributes {
       		name = "endpointType"
       		value = "private"
    	}
  }
  contexts {
		attributes {
			name = "networkZoneId"
			value = "bf823d4f45b64ceaa4671bee0479346e"
		}
		attributes {
       		name = "endpointType"
       		value = "public"
    	}
  }
  description = "this is an example of rule with two contexts"
  enforcement_mode = "enabled"
  operations {
		api_types {
			api_type_id = "api_type_id"
		}
  }
  resources {
		attributes {
			name = "accountId"
			value = "12ab34cd56ef78ab90cd12ef34ab56cd"
		}
		attributes {
      		name = "serviceName"
      		value = "network-policy-enabled"
    	}
		tags {
      		name     = "tag_name"
      		value    = "tag_value"
		}
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `contexts` - (Optional, List) The contexts this rule applies to.
  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
Nested schema for **contexts**:
	* `attributes` - (Required, List) The attributes.
	  * Constraints: The minimum length is `1` item.
	Nested schema for **attributes**:
		* `name` - (Required, String) The attribute name.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.
		* `value` - (Required, String) The attribute value.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[\S\s]+$/`.
* `description` - (Optional, String) The description of the rule.
  * Constraints: The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `/^[\x20-\xFE]*$/`.
* `enforcement_mode` - (Optional, String) The rule enforcement mode: * `enabled` - The restrictions are enforced and reported. This is the default. * `disabled` - The restrictions are disabled. Nothing is enforced or reported. * `report` - The restrictions are evaluated and reported, but not enforced.
  * Constraints: The default value is `enabled`. Allowable values are: `enabled`, `disabled`, `report`.
* `operations` - (Optional, List) The operations this rule applies to.
Nested schema for **operations**:
	* `api_types` - (Required, List) The API types this rule applies to.
	  * Constraints: The maximum length is `100` items. The minimum length is `1` item.
	Nested schema for **api_types**:
		* `api_type_id` - (Required, String)
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_.\-:]+$/`.
* `resources` - (Optional, List) The resources this rule apply to.
  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
Nested schema for **resources**:
	* `attributes` - (Required, List) The resource attributes.
	  * Constraints: The minimum length is `1` item.
	Nested schema for **attributes**:
		* `name` - (Required, String) The attribute name.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_]+$/`.
		* `operator` - (Optional, String) The attribute operator.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.
		* `value` - (Required, String) The attribute value.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[\S\s]+$/`.
	* `tags` - (Optional, List) The optional resource tags.
	  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
	Nested schema for **tags**:
		* `name` - (Required, String) The tag attribute name.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 _.-]+$/`.
		* `operator` - (Optional, String) The attribute operator.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.
		* `value` - (Required, String) The tag attribute value.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 _*?.-]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cbr_rule.
* `created_at` - (String) The time the resource was created.
* `created_by_id` - (String) IAM ID of the user or service which created the resource.
* `crn` - (String) The rule CRN.
* `href` - (String) The href link to the resource.
* `last_modified_at` - (String) The last time the resource was modified.
* `last_modified_by_id` - (String) IAM ID of the user or service which modified the resource.

* `etag` - ETag identifier for cbr_rule.

## Import

You can import the `ibm_cbr_rule` resource by using `id`. The globally unique ID of the rule.

# Syntax
<pre>
$ terraform import ibm_cbr_rule.cbr_rule &lt;id&gt;
</pre>
