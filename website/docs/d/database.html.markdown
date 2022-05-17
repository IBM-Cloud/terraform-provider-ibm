---
layout: "ibm"
page_title: "IBM : ibm_database"
description: |-
  Get information about database
subcategory: "The IBM Cloud Databases API"
---

# ibm_database

Provides a read-only data source for database. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_database" "database" {
	id = "id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `id` - (Required, Forces new resource, String) Deployment ID.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the database.
* `admin_usernames` - (Optional, Map) Login name of administration level user.

* `enable_private_endpoints` - (Optional, Boolean) Whether access to this deployment is enabled from IBM Cloud via the IBM Cloud backbone network. This property can be modified by updating this service instance through the Resource Controller API.

* `enable_public_endpoints` - (Optional, Boolean) Whether access to this deployment is enabled from the public internet. This property can be modified by updating this service instance through the Resource Controller API.

* `name` - (Optional, String) Readable name of this deployment.

* `platform` - (Optional, String) Platform for this deployment.

* `platform_options` - (Optional, Map) Platform-specific options for this deployment.

* `type` - (Optional, String) Database type within this deployment.

* `version` - (Optional, String) Version number of the database.

