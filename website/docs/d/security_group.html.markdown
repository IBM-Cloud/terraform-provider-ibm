---
layout: "ibm"
page_title: "IBM: ibm_security_group"
sidebar_current: "docs-ibm-datasource-security-group"
description: |-
  Get information about an IBM Security Group.
---

# ibm\_security_group

Import the details of an existing security group as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_security_group" "allow_ssh" {
    name = "allow_ssh"
}
```

The following example shows how you can use this data source to reference the security group IDs in the `ibm_compute_vm_instance` resource because the numeric IDs are often unknown.

```hcl
resource "ibm_compute_vm_instance" "vm1" {
    ...
    private_security_group_ids = ["${data.ibm_security_group.allow_ssh.id}"]
    ...
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the security group, as it was defined in IBM Cloud Infrastructure (SoftLayer).
* `description` - (Optional, string) The description of the security group, as it was defined in IBM Cloud Infrastructure (SoftLayer).
* `most_recent` - (Optional, boolean) If more than one security group  matches the name and/or description, you can set this argument to `true` to import only the most recent security group.
  **NOTE**: The search must return only one match. More or less than one match causes Terraform to fail. Ensure that your name and description combinations are specific enough to return a single security group  key only, or use the `most_recent` argument.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the security group.
* `description` - The description of the security group.
