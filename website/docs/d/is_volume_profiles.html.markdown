---
layout: "ibm"
page_title: "IBM : Volume profiles"
sidebar_current: "docs-ibm-datasources-is-volume-profiles"
description: |-
  Manages IBM Cloud virtual server volume profiles.
---

# ibm\_is_volume_profiles

Import the details of an existing IBM Cloud virtual server volume profiles as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_volume_profiles" "volprofiles"{
}

```

## Attribute Reference

The following attributes are exported:

* `profiles` - List of all server volume profiles in the region.
  * `name` - The name for this virtual server volume profile.
  * `family` - The family of the virtual server volume profile.

