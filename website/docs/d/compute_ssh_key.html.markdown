---
layout: "ibm"
page_title: "IBM: ibm_compute_ssh_key"
sidebar_current: "docs-ibm-datasource-compute-ssh-key"
description: |-
  Get information about an IBM Compute SSH key.
---

# ibm\_compute_ssh_key

Import the details of an existing SSH key as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration by using interpolation syntax. 

## Example Usage

```hcl
data "ibm_compute_ssh_key" "public_key" {
    label = "Terraform Public Key"
}
```

The following example shows how you can use this data source to reference the SSH key IDs in the `ibm_compute_vm_instance` resource, since the numeric IDs are often unknown.

```hcl
resource "ibm_compute_vm_instance" "vm1" {
    ...
    ssh_key_ids = ["${data.ibm_compute_ssh_key.public_key.id}"]
    ...
}
```

## Argument Reference

The following arguments are supported:

* `label` - (Required) The label of the SSH key, as it was defined in Bluemix Infrastructure (SoftLayer).
* `most_recent` - (Optional) If more than one SSH key matches the label, you can use this argument to import only the most recent key.

**NOTE**: Terraform will fail if more or less than a single match is returned by the search. Ensure that your label is specific enough to return a single SSH key only, or use the `most_recent` argument.

## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the SSH key.  
* `fingerprint` - Sequence of bytes to authenticate or look up a longer SSH key.
* `public_key` - The public key contents.
* `notes` - Notes stored with the SSH key.
