---
layout: "ibm"
page_title: "IBM : compute_provisioning_hook"
sidebar_current: "docs-ibm-resource-compute-provisioning-hook"
description: |-
  Manages IBM Compute Provisioning Hook.
---


# ibm\_compute_provisioning_hook

Provides provisioning hooks containing all the information needed to add a hook into a server or virtual provision and OS reload. This allows provisioning hooks to be created, updated, and deleted.

For additional details, see the [IBM Cloud Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Provisioning_Hook).

## Example Usage

```hcl
resource "ibm_compute_provisioning_hook" "test_provisioning_hook" {
    name = "test_provisioning_hook_name"
    uri  = "https://raw.githubusercontent.com/test/slvm/master/test-script.sh"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The descriptive name used to identify a provisioning hook.
* `uri` - (Required, string) The endpoint from which the script is downloaded or downloaded and executed.
* `tags` - (Optional, array of strings) Tags associated with the provisioning hook instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new provisioning hook.
