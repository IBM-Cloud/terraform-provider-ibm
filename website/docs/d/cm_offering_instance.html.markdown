---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_offering_instance"
description: |-
  Get information about cm_offering_instance
---

# ibm\_cm_offering_instance

Provides a read-only data source for cm_offering_instance. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "cm_offering_instance" "cm_offering_instance" {
	instance_identifier = "instance_identifier"
}
```

## Argument Reference

The following arguments are supported:

* `instance_identifier` - (Required, string) Version Instance identifier.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cm_offering_instance.

* `url` - url reference to this object.

* `crn` - platform CRN for this instance.

* `label` - the label for this instance.

* `catalog_id` - Catalog ID this instance was created from.

* `offering_id` - Offering ID this instance was created from.

* `kind_format` - the format this instance has (helm, operator, ova...).

* `version` - The version this instance was installed from (not version id).

* `cluster_id` - Cluster ID.

* `cluster_region` - Cluster region (e.g., us-south).

* `cluster_namespaces` - List of target namespaces to install into.

* `cluster_all_namespaces` - designate to install into all namespaces.

