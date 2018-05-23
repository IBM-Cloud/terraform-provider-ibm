---
layout: "ibm"
page_title: "IBM: ibm_resource_group"
sidebar_current: "docs-ibm-datasource-reosurce-group"
description: |-
  Get information about an IBM resource Group.
---

# ibm\_resource_group

Import the details of an existing IBM resource Group as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

Import the reosource group by name

```hcl
data "ibm_resource_group" "group" {
  name = "test"
}
```

Import the default resource group

```hcl
data "ibm_resource_group" "group" {
  is_default = "true"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The name of the IBM Cloud resource group. You can retrieve the value by running the `bx resource groups` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).  
  **NOTE**: Conflicts with `is_default`.

* `is_default` - (Optional, boolean) Specifies whether you want to import default resource group. The default value is `false`.  
  **NOTE**: Conflicts with `name`.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the resource group.  
