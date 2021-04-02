---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_offering_instance"
description: |-
  Manages cm_offering_instance.
---

# ibm\_cm_offering_instance

Provides a resource for cm_offering_instance. This allows cm_offering_instance to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cm_offering_instance" "cm_offering_instance" {
  catalog_id = "catalog_id"
  label = "placeholder"
  kind_format = "operator"
  version = "placeholder"
  cluster_id = "placeholder"
  cluster_region = "placeholder"
  cluster_namespaces = [ "placeholder", "placeholder2" ]
  cluster_all_namespaces = false
}
```

## Argument Reference

The following arguments are supported:

* `label` - (Required, string) the label for this instance.
* `catalog_id` - (Required, string) Catalog ID this instance was created from.
* `offering_id` - (Required, string) Offering ID this instance was created from.
* `kind_format` - (Required, string) the format this instance has (helm, operator, ova...).
* `version` - (Required, string) The version this instance was installed from (not version id).
* `cluster_id` - (Required, string) Cluster ID.
* `cluster_region` - (Required, string) Cluster region (e.g., us-south).
* `cluster_namespaces` - (Required, List) List of target namespaces to install into.
* `cluster_all_namespaces` - (Required, bool) designate to install into all namespaces.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cm_offering_instance.
* `url` - Url reference to this object.
* `crn` - Platform CRN for this instance.
