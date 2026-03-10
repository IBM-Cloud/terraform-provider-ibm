---
layout: "ibm"
page_title: "IBM : ibm_pdr_get_managed_vm_list"
description: |-
  Get information about pdr_get_managed_vm_list
subcategory: "DrAutomation Service"
---

# ibm_pdr_get_managed_vm_list

Retrieves the list of disaster recovery (DR) managed virtual machines for the specified service instance.

## Example Usage

```hcl
data "ibm_pdr_get_managed_vm_list" "pdr_get_managed_vm_list" {
	instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.(ex., en,it,fr,es,de,ja,ko,pt-BR,zh-HANS,zh-HANT)
* `instance_id` - (Required, Forces new resource, String) ID of the service instance.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_get_managed_vm_list.
* `managed_vms` - (Map) A map where the key is the VM ID and the value is the corresponding ManagedVmDetails object.
