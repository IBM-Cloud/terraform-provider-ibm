---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: ibm_security_group"
description: |-
  Get information about an IBM Cloud security group.
---

# ibm_security_group
Retrieve information of an existing security group as a read-only data source. For more information, about IBM Cloud security group, see [managing security groups](https://cloud.ibm.com/docs/security-groups?topic=security-groups-managing-sg).

## Example usage
The following example shows how you can use this data source to reference the security group IDs in the `ibm_compute_vm_instance` resource because the numeric IDs are often unknown.


```terraform
data "ibm_security_group" "allow_ssh" {
    name = "allow_ssh"
}

resource "ibm_compute_vm_instance" "vm1" {
  # TF-UPGRADE-TODO: In Terraform v0.10 and earlier, it was sometimes necessary to
  # force an interpolation expression to be interpreted as a list by wrapping it
  # in an extra set of list brackets. That form was supported for compatibility in
  # v0.11, but is no longer supported in Terraform v0.12.
  #
  # If the expression in the following list itself returns a list, remove the
  # brackets to avoid interpretation as a list of lists. If the expression
  # returns a single list item then leave it as-is and remove this TODO comment.
  private_security_group_ids = [data.ibm_security_group.allow_ssh.id]
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `description` - (Optional, String) The description of the security group, as defined in IBM Cloud Classic Infrastructure.
- `most_recent` - (Optional, Bool) If more than one security group has the same name or description, you can set this argument to **true** to import only the most recent security group. **Note**: The search must return only one match, otherwise Terraform fails. Ensure that your name and description combinations are specific to return a single security group key only, or set the **most_recent** argument to **true**.
- `name` - (Required, String) The name of the security group, as defined in IBM Cloud Classic Infrastructure.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `description` - (String) The description of the security group.
- `id` - (String) The unique identifier of the security group.
