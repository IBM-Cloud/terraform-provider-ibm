---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: compute_ssh_key"
description: |-
  Manages IBM Compute SSH keys.
---

# ibm_compute_ssh_key

Create, update, and delete an SSH key resource. For more information, about computer SSH key, see [deploying server pools and origins in a single MZR](https://cloud.ibm.com/docs/cloud-infrastructure?topic=cloud-infrastructure-ha-pools-origins).

**Note**

For more information, see the [IBM Cloud Classic Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Security_Ssh_Key).

## Example usage

```terraform
resource "ibm_compute_ssh_key" "test_ssh_key" {
    label = "test_ssh_key_name"
    notes = "test_ssh_key_notes"
    public_key = "ssh-rsa <rsa_public_key>"
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `label`- (Required, String) The descriptive name that is used to identify an SSH key.
- `notes`- (Optional, string) Descriptive text about the SSH key.
- `public_key`- (Required, String) The public SSH key.
- `tags`- (Optional, Array of Strings) Tags associated with the SSH Key instance. **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `fingerprint`- (String) The sequence of bytes to authenticate or look up a longer SSH key.
- `id`- (String )The unique identifier of the new SSH key.
