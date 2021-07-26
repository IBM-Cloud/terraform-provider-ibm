---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : compute_provisioning_hook"
description: |-
  Manages IBM Cloud compute provisioning hook.
---


# ibm_compute_provisioning_hook
Provides provisioning hooks that contains all the information that is needed to add a hook into a server or virtual provision and OS reload. This allows provisioning hooks to be created, updated, and deleted. For more information, about provisioning hook, see [integrate third-party services using hooks](https://cloud.ibm.com/docs/cloud-foundry?topic=cloud-foundry-hooks).

**Note**

For more information, see the [IBM Cloud Classic Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Provisioning_Hook).

## Example usage

```terraform
resource "ibm_compute_provisioning_hook" "test_provisioning_hook" {
    name = "test_provisioning_hook_name"
    uri  = "https://raw.githubusercontent.com/test/slvm/master/test-script.sh"
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `name`- (Required, string) The descriptive name that is used to identify a provisioning hook.
- `tags`- (Optional, array of strings) Tags associated with the provisioning hook instance. **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.
- `uri`- (Required, string) The endpoint from which the script is downloaded or downloaded and executed.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The unique identifier of the new provisioning hook.
