---
layout: "ibm"
page_title: "IBM : ibm_cbr_rule"
description: |-
  Manages cbr_rule.
subcategory: "Context Based Restrictions"
---

# ibm_cbr_rule

Provides a resource for cbr_rule. This allows cbr_rule to be created, updated and deleted.

## Example Usage to create a rule with one context and two zones

```hcl
resource "ibm_cbr_rule" "cbr_rule" {
  contexts {
		attributes {
			name = "networkZoneId"
			value = "559052eb8f43302824e7ae490c0281eb, bf823d4f45b64ceaa4671bee0479346e"
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
resource "ibm_cbr_rule" "cbr_rule" {
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

Review the argument reference that you can specify for your resource.

* `contexts` - (Optional, List) The contexts this rule applies to.
  * Constraints: The maximum length is `1000` items. The minimum length is `1` item.
Nested scheme for **contexts**:
	* `attributes` - (Required, List) The attributes.
	  * Constraints: The minimum length is `1` item.
	Nested scheme for **attributes**:
		* `name` - (Required, String) The attribute name.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.
		* `value` - (Required, String) The attribute value.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[\S\s]+$/`.
* `description` - (Optional, String) The description of the rule.
  * Constraints: The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `/^[\x20-\xFE]*$/`.
* `enforcement_mode` - (Optional, String) The rule enforcement mode: * `enabled` - The restrictions are enforced and reported. This is the default. * `disabled` - The restrictions are disabled. Nothing is enforced or reported. * `report` - The restrictions are evaluated and reported, but not enforced.
  * Constraints: The default value is `enabled`. Allowable values are: `enabled`, `disabled`, `report`.
* `operations` - (Optional, List) The operations this rule applies to.
Nested scheme for **operations**:
	* `api_types` - (Required, List) The API types this rule applies to.
	  * Constraints: The maximum length is `100` items. The minimum length is `1` item.
	Nested scheme for **api_types**:
		* `api_type_id` - (Required, String)
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_.\-:]+$/`.
* `resources` - (Optional, List) The resources this rule apply to.
  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
Nested scheme for **resources**:
	* `attributes` - (Required, List) The resource attributes.
	  * Constraints: The minimum length is `1` item.
	Nested scheme for **attributes**:
		* `name` - (Required, String) The attribute name.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.
		* `operator` - (Optional, String) The attribute operator.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.
		* `value` - (Required, String) The attribute value.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[\S\s]+$/`.
	* `tags` - (Optional, List) The optional resource tags.
	  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
	Nested scheme for **tags**:
		* `name` - (Required, String) The tag attribute name.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 _.-]+$/`.
		* `operator` - (Optional, String) The attribute operator.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.
		* `value` - (Required, String) The tag attribute value.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 _*?.-]+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cbr_rule.
* `created_at` - (String) The time the resource was created.
* `created_by_id` - (String) IAM ID of the user or service which created the resource.
* `crn` - (String) The rule CRN.
* `href` - (String) The href link to the resource.
* `last_modified_at` - (String) The last time the resource was modified.
* `last_modified_by_id` - (String) IAM ID of the user or service which modified the resource.

* `version` - Version of the cbr_rule.

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

You can import the `ibm_cbr_rule` resource by using `id`. The globally unique ID of the rule.

# Syntax
```
$ terraform import ibm_cbr_rule.cbr_rule <id>
```
