---
layout: "ibm"
page_title: "IBM : ibm_cm_object"
description: |-
  Manages ibm_cm_object.
subcategory: "Catalog Management"
---

# ibm_cm_object

Provides a resource for ibm_cm_object. This allows ibm_cm_object to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cm_object" "cm_object" {
  catalog_id = ibm_cm_catalog.cm_catalog.id
  name = "object_name"
  label = "Object Label"
  kind = "preset_configuration"
  short_description = "short description"
  data = jsonencode(file("data.json"))
  parent_id = "us-south"
  tags = [ "tag1", "tag2" ]
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `catalog_id` - (Required, Forces new resource, String) Catalog identifier.
* `name` - (Required, String) The programmatic name of this object.
* `kind` - (Required, String) Kind of object. Options are "vpe", "preset_configuration", or "proxy_source".
* `parent_id` - (Optional, String) The parent region for this specific object.
* `label` - (Optional, String) Display name in the requested language.
* `tags` - (Optional, List) List of tags associated with this catalog.
* `short_description` - (Optional, String) Short description in the requested language.
* `data` - (Optional, String) Stringified map of object data.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `catalog_name` - (String) The name of the catalog.
* `created` - (String) The date and time this catalog was created.
* `crn` - (String) The crn for this specific object.
* `data` - (String) Stringified map of object data.
* `id` - The unique identifier of the ibm_cm_object.
* `kind` - (String) Kind of object.
* `label` - (String) Display name in the requested language.
* `name` - (String) The programmatic name of this object.
* `parent_id` - (String) The parent region for this specific object.
* `publish` - Publish information.
* Nested scheme for **publish**:
	* `permit_ibm_public_publish` - (Boolean) Is it permitted to request publishing to IBM or Public.
	* `ibm_approved` - (Boolean) Indicates if this offering has been approved for use by all IBMers.
	* `public_approved` - (Boolean) Indicates if this offering has been approved for use by all IBM Cloud users.
	* `portal_approval_record` - (String) The portal's approval record ID.
	* `portal_url` - (String) The portal UI URL.
* `rev` - (String) Cloudant revision.
* `short_description` - (String) Short description in the requested language.
* `state` - Object state.
* Nested scheme for **state**:
	* `current` - (String) one of: new, validated, account-published, ibm-published, public-published.
	* `current_entered` - (String) Date and time of current request.
	* `pending` - (String) one of: new, validated, account-published, ibm-published, public-published.
	* `pending_requested` - (String) Date and time of pending request.
	* `previous` - (String) one of: new, validated, account-published, ibm-published, public-published.
* `tags` - (List) List of tags associated with this catalog.
* `updated` - (String) The data and time this catalog was last updated.
* `url` - (String) The url for this specific object.

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

You can import the `ibm_cm_object` resource by using `id`.
The `id` property is just the `object_id`.

* `object_id`: A string. Object identification.

# Syntax
```
$ terraform import ibm_cm_object.cm_object <object_id>
```
