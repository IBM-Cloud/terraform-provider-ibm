---
layout: "ibm"
page_title: "IBM : ibm_secret_group"
description: |-
  Manages SecretGroup.
subcategory: "IBM Cloud Secrets Manager Basic API"
---

# ibm_secret_group

Provides a resource for SecretGroup. This allows SecretGroup to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_secret_group" "secret_group" {
  description = "Extended description for this group."
  name = "my-secret-group"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `description` - (Optional, String) An extended description of your secret group.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.
* `name` - (Required, String) The name of your secret group.
  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z][A-Za-z0-9]*(?:_?-?.?[A-Za-z0-9]+)*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the SecretGroup.
* `creation_date` - (String) The date a resource was created. The date format follows RFC 3339.
* `last_update_date` - (String) The date a resource was recently modified. The date format follows RFC 3339.

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

You can import the `ibm_secret_group` resource by using `id`. A v4 UUID identifier.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager)

# Syntax
```
$ terraform import ibm_secret_group.secret_group <id>
```

# Example
```
$ terraform import ibm_secret_group.secret_group b49ad24d-81d4-5ebc-b9b9-b0937d1c84d5
```
