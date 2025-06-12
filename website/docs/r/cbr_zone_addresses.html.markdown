---
layout: "ibm"
page_title: "IBM : ibm_cbr_zone_addresses"
description: |-
  Manages cbr_zone_addresses.
subcategory: "Context Based Restrictions"
---

# ibm_cbr_zone_addresses

Provides a resource for cbr_zone_addresses. This allows cbr_zone_addresses to be created, updated and deleted.
This resource allows the inclusion of one or more addresses in an existing zone without the need to modify the base cbr_zone resource.

## Example Usage to create a zone addresses resource

```hcl
resource "ibm_cbr_zone_addresses" "cbr_zone_addresses" {
  zone_id = ibm_cbr_zone.cbr_zone.id
  addresses {
    type = "subnet"
    value = "169.24.56.0/24"
  }
  addresses {
    type = "ipRange"
    value = "169.24.22.0-169.24.22.255"
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `zone_id` - (Required, String) The id of the zone in which to include the addresses.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-fA-F0-9]{32}$/`.
* `addresses` - (Required, List) The list of addresses to include in the zone.
  * Constraints: The maximum length is `1000` items. The minimum length is `1` items.
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

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cbr_zone_addresses.

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

You can import the `ibm_cbr_zone_addresses` resource by using `id`. The globally unique ID of the zone addresses.

# Syntax
```
$ terraform import ibm_cbr_zone_addresses.cbr_zone_addresses <id>
```
