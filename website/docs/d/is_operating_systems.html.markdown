---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : operating_systems"
description: |-
  Manages IBM Operating Systems.
---

# ibm_is_operating_systems
Retrieve information of an existing IBM Cloud Infrastructure Operating Systems as a read only data source. For more information, about supported Operating System, see [Images](https://cloud.ibm.com/docs/vpc?topic=vpc-about-images).

## Example usage

```terraform
data "ibm_is_operating_systems" "testacc_dsoslist"{
}
```

## Attribute reference
You can access the following attribute references after your data source is created. 

- `operating_systems` - (List) List of all Operating Systems in the IBM Cloud Infrastructure region.

  Nested scheme for `operating_system`:
  - `architecture` - (String) The operating system architecture.
  - `dedicated_host_only` - (String) Images with this operating system can only be used on dedicated hosts or dedicated host groups.
  - `display_name` - (String) A unique, display-friendly name for the operating system.
  - `family` - (String) The name of the software family this operating system belongs to.
  - `href` - (String) The URL for this operating system.
  - `name` - (String) The globally unique name for this operating system.
  - `vendor` - (String) The vendor of the operating system.
  - `version` - (String) The major release version of this operating system.
