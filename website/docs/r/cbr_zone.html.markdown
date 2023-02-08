---
layout: "ibm"
page_title: "IBM : ibm_cbr_zone"
description: |-
  Manages cbr_zone.
subcategory: "Context Based Restrictions"
---

# ibm_cbr_zone

Provides a resource for cbr_zone. This allows cbr_zone to be created, updated and deleted.

## Example Usage to create a zone with excluded addresses

```hcl
resource "ibm_cbr_zone" "cbr_zone" {
  account_id = "12ab34cd56ef78ab90cd12ef34ab56cd"
  addresses {
    type = "ipAddress"
    value = "169.23.56.234"
  }
  addresses {
    type = "ipRange"
    value = "169.23.22.0-169.23.22.255"
  }
  excluded {
    type  = "ipAddress"
    value = "169.23.22.10"
  }
  excluded {
    type  = "ipAddress"
    value = "169.23.22.11"
  }
  description = "this is an example of zone"
  excluded {
		type = "ipAddress"
		value = "value"
  }
  name = "an example of zone"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `account_id` - (Optional, String) The id of the account owning this zone.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-]+$/`.
* `addresses` - (Optional, List) The list of addresses in the zone.
  * Constraints: The maximum length is `1000` items. The minimum length is `1` item.
Nested scheme for **addresses**:
	* `ref` - (Optional, List) A service reference value.
	Nested scheme for **ref**:
		* `account_id` - (Required, String) The id of the account owning the service.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-]+$/`.
		* `location` - (Optional, String) The location.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-]+$/`.
		* `service_instance` - (Optional, String) The service instance.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-\/]+$/`.
		* `service_name` - (Optional, String) The service name.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-]+$/`.
		* `service_type` - (Optional, String) The service type.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z_]+$/`.
	* `type` - (Optional, String) The type of address.
	  * Constraints: Allowable values are: `ipAddress`, `ipRange`, `subnet`, `vpc`, `serviceRef`.
	* `value` - (Optional, String) The IP address.
	  * Constraints: The maximum length is `45` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9:.]+$/`.
* `description` - (Optional, String) The description of the zone.
  * Constraints: The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `/^[\x20-\xFE]*$/`.
* `excluded` - (Optional, List) The list of excluded addresses in the zone. Only addresses of type `ipAddress`, `ipRange`, and `subnet` can be excluded.
  * Constraints: The maximum length is `1000` items.
Nested scheme for **excluded**:
	* `ref` - (Optional, List) A service reference value.
	Nested scheme for **ref**:
		* `account_id` - (Required, String) The id of the account owning the service.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\-]+$/`.
		* `location` - (Optional, String) The location.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-]+$/`.
		* `service_instance` - (Optional, String) The service instance.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-/]+$/`.
		* `service_name` - (Optional, String) The service name.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z\-]+$/`.
		* `service_type` - (Optional, String) The service type.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-z_]+$/`.
	* `type` - (Optional, String) The type of address.
	  * Constraints: Allowable values are: `ipAddress`, `ipRange`, `subnet`, `vpc`, `serviceRef`.
	* `value` - (Optional, String) The IP address.
	  * Constraints: The maximum length is `45` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9:.]+$/`.
* `name` - (Optional, String) The name of the zone.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 \-_]+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cbr_zone.
* `address_count` - (Integer) The number of addresses in the zone.
* `created_at` - (String) The time the resource was created.
* `created_by_id` - (String) IAM ID of the user or service which created the resource.
* `crn` - (String) The zone CRN.
* `excluded_count` - (Integer) The number of excluded addresses in the zone.
* `href` - (String) The href link to the resource.
* `last_modified_at` - (String) The last time the resource was modified.
* `last_modified_by_id` - (String) IAM ID of the user or service which modified the resource.

* `version` - Version of the cbr_zone.

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

You can import the `ibm_cbr_zone` resource by using `id`. The globally unique ID of the zone.

# Syntax
```
$ terraform import ibm_cbr_zone.cbr_zone <id>
```
