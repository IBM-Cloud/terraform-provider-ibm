---
layout: "ibm"
page_title: "IBM : ibm_cm_catalog"
description: |-
  Manages ibm_cm_catalog.
subcategory: "Catalog Management"
---

# ibm_cm_catalog

Provides a resource for ibm_cm_catalog. This allows ibm_cm_catalog to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cm_catalog" "cm_catalog" {
  label = "catalog_label"
  short_description = "catalog description"
  catalog_icon_url = "icon url"
  kind = "offering"
  tags = ["catalog", "tags"]
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `catalog_icon_url` - (Optional, String) URL for an icon associated with this catalog.
* `catalog_banner_url` - (Optional, String) URL for a banner image for this catalog.
* `disabled` - (Optional, Boolean) Denotes whether a catalog is disabled.
* `kind` - (Optional, String) Kind of catalog. Supported kinds are offering and vpe.
* `label` - (Optional, String) Display Name in the requested language.
* `resource_group_id` - (Optional, String) Resource group id the catalog is owned by.
* `short_description` - (Optional, String) Description in the requested language.
* `tags` - (Optional, List) List of tags associated with this catalog.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the ibm_cm_catalog.
* `catalog_icon_url` - (String) The url of the catalog icon.
* `catalog_banner_url` - (String) The url of the catalog banner.
* `created` - (String) The date-time this catalog was created.
* `crn` - (String) CRN associated with the catalog.
* `disabled` - (Boolean) Denotes whether a catalog is disabled.
* `label` - (String) Label of the catalog
* `kind` - (String) Kind of catalog.
* `offerings_url` - (String) URL path to offerings.
* `owning_account` - (String) The account ID of the owning account.
* `resource_group_id` - (String) Resource group id the catalog is owned by.
* `rev` - (String) Cloudant revision.
* `short_description` - (String) Description in the requested language.
* `tags` - (List) List of tags associated with this catalog.
* `updated` - (String) The date-time this catalog was last updated.
* `url` - (String) The url for this specific catalog.

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
export IC_API_KEY="api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="api_key"
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

You can import the `ibm_cm_catalog` resource by using `id`. Unique ID.

# Syntax
```
$ terraform import ibm_cm_catalog.cm_catalog <id>
```
