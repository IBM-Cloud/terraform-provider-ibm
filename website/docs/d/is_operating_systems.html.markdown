---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : operating_systems"
description: |-
  Manages IBM Operating Systems.
---

# ibm\_is_operating_systems

Import the details of an existing IBM Cloud Infrastructure Operating Systems as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.



## Example Usage

```terraform
data "ibm_is_operating_systems" "testacc_dsoslist"{
}
```

## Attribute Reference

The following attributes are exported:
* `operating_systems` - List of all Operating Systems in the IBM Cloud Infrastructure region.
  * `architecture` - The operating system architecture.
  * `dedicated_host_only` - Images with this operating system can only be used on dedicated hosts or dedicated host groups.
  * `display_name` - A unique, display-friendly name for the operating system.
  * `family` - The name of the software family this operating system belongs to.
  * `href` - The URL for this operating system.
  * `name` - The globally unique name for this operating system.
  * `vendor` - The vendor of the operating system.
  * `version` - The major release version of this operating system.