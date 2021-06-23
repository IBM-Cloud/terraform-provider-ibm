---
subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: ibm_compute_ssh_key"
description: |-
  Get information about an IBM Cloud compute SSH key.
---

# ibm_compute_ssh_key
Retrieve information of an existing SSH key as a read-only data source. For more information, about computer SSH key, see [deploying server pools and origins in a single MZR](https://cloud.ibm.com/docs/cloud-infrastructure?topic=cloud-infrastructure-ha-pools-origins).

## Example usage

```terraform
data "ibm_compute_ssh_key" "public_key" {
    label = "Terraform Public Key"
}
```

The following example shows how you can use this data source to reference the SSH key IDs in the `ibm_compute_vm_instance` resource because the numeric IDs are often unknown.

```terraform
resource "ibm_compute_vm_instance" "vm1" {
  # TF-UPGRADE-TODO: In Terraform v0.10 and earlier, it was sometimes necessary to
  # force an interpolation expression to be interpreted as a list by wrapping it
  # in an extra set of list brackets. That form was supported for compatibility in
  # v0.11, but is no longer supported in Terraform v0.12.
  #
  # If the expression in the following list itself returns a list, remove the
  # brackets to avoid interpretation as a list of lists. If the expression
  # returns a single list item then leave it as-is and remove this TODO comment.
  ssh_key_ids = [data.ibm_compute_ssh_key.public_key.id]
}

```

## Argument reference
Review the argument references that you can specify for your data source.

- `label` - (Required, String) The label of the SSH key.
- `most_recent` - (Optional, Bool) If more than one SSH key matches the label, you can set this argument to **true** to import only the most recent key. **Note** The search must return only one match. More or less than one match causes  Terraform to fail. Ensure that your label is specific enough to return a single SSH key only, or use the `most_recent` argument.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `fingerprint` - (String) The sequence of bytes to authenticate or look up a longer SSH key.
- `id` - (String) The unique identifier of the SSH key.
- `notes` - (String) The notes that are stored with the SSH key.
- `public_key` - (String) The public key contents.
