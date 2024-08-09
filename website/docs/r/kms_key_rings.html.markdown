---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-key-rings"
description: |-
  Manages key rings for IBM hs-crypto and KMS.
---

# ibm_kms_key_rings
Create, modify, or delete a key rings for hs-crypto and key protect services. Key rings created through this resource can be used to associate to KMS key resource when a standard or a root key gets created or imported. For more information, about key management rings, see [creating key rings](https://cloud.ibm.com/docs/key-protect?topic=key-protect-grouping-keys#create-key-ring-api).


## Example usage 
Sample example to provision key ring and associate a key management service key.

```terraform
resource "ibm_resource_instance" "kms_instance" {
  name     = "instance-name"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}
resource "ibm_kms_key_rings" "key_ring" {
  instance_id = ibm_resource_instance.kms_instance.guid
  key_ring_id = "key-ring-id"
}
resource "ibm_kms_key" "key" {
  instance_id = ibm_resource_instance.kp_instance.guid
  key_name       = "key"
  key_ring_id = ibm_kms_key_rings.key_ring.key_ring_id
  standard_key   = false
  payload = "aW1wb3J0ZWQucGF5bG9hZA=="
}
```

~>**Deprecated:**
`force_delete` argument will no longer be supported. Users are advised to remove references to `force_delete` from all `ibm_kms_key_rings` configurations by `July 30th 2025`. New default behavior of deleting a key ring is to move keys with state equals to 5 (destroyed) to the default key ring.

## Argument reference
Review the argument references that you can specify for your resource. 

- `endpoint_type` - (Optional, Forces new resource, String) The type of the public endpoint, or private endpoint to be used for creating keys.
- `instance_id` - (Required, Forces new resource, String) The hs-crypto or key protect instance GUID.
- `key_ring_id` - (Required, Forces new resource, String) The ID that identifies the key ring. Each ID is unique within the given instance and is not reserved across the key protect service. **Constraints** `2 ≤ length ≤ 100`. Value must match regular expression of `^[a-zA-Z0-9-]*$`.
- `force_delete` - (**Deprecated**)(Optional, Bool) If set to **true**, allows force deletion of a key ring. Terraform users are recommended to have this set to **true**. All keys in the key ring are required to be deleted (in state **5**) before this action can be performed. If the key ring to be deleted contains keys, they will be moved to the **default** key ring which requires the **kms.secrets.patch** IAM action.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique ID for the Terraform resource.
- `key_ring_id` - (String) The key ring ID.
