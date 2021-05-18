---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_offering_instance"
description: |-
  Get information about cm_offering_instance.
---


# `ibm_cm_offering_instance`

Create, modify, or delete an `ibm_cm_offering_instance` data source.  For more information, about managing catalog, refer to [catalog management settings](https://cloud.ibm.com/docs/account?topic=account-account-getting-started).


## Example usage

```
data "cm_offering_instance" "cm_offering_instance" {
	instance_identifier = "instance_identifier"
}
```

## Argument reference
Review the input parameters that you can specify for your data source. 

- `instance_identifier` - (Required, String) The version instance identifier.

## Attribute reference
Review the output parameters that you can access after your data source is created. 

- `catalog_id` - (String) The catalog ID the instance that is created from.
- `cluster_id` - (String) The cluster ID.
- `cluster_region` - (String) The cluster region for example, `us-south`.
- `cluster_namespaces` - (String) The list of target namespaces to install.
- `cluster_all_namespaces` - (String) Designate to install into all namespaces.
- `crn` - (String) The platform CRN for an instance.
- `id` - (String) The unique identifier of the `cm_offering_instance`.
- `kind_format` - (String) The format this instance has such as `helm`, `operator`.
- `label` - (String) The label for an instance.
- `offering_id` - (String) The offering ID the instance that is created from.
- `url` - (String) The URL reference to an object.
- `version` - (String) The version an instance is installed from (but not from the version ID).

