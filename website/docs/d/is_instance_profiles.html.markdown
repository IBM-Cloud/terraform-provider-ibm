---
layout: "ibm"
page_title: "IBM : Instance profiles"
sidebar_current: "docs-ibm-datasources-is-instance-profiles"
description: |-
  Manages IBM Cloud virtual server instance profiles.
---

# ibm\_is_instance_profiles

Import the details of an existing IBM Cloud virtual server instance profiles as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_instance_profiles" "ds_instance_profiles" {
}

```

## Attribute Reference

The following attributes are exported:

* `profiles` - List of all server instance profiles in the region.
  * `name` - The name for this virtual server instance profile.
  * `family` - The family of the virtual server instance profile.
  * `generation` - The platform generation of the virtual server instance profile.

