---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-key-rings"
description: |-
  Manages key rings for IBM hs-crypto or key-protect.
---

# ibm_kms_key_rings

Retrieve a list of key rings from the hs-crypto or key protect instance. For more information, about retrieving key and key rings, see [Retrieving a key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-retrieve-key).

## Example usage

```terraform
data "ibm_kms_key_rings" "test" {
  instance_id = "guid-of-keyprotect-or hs-crypto-instance"
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `endpoint_type` - (Optional, String) The type of the public endpoint, or private endpoint to be used for creating keys.
- `instance_id` - (Required, String) The key protect instance GUID.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `key_rings` - (List of objects) A list of all key rings in the hs-crypto or key protect instance.

   Nested scheme for `key_rings`:
   - `created_by` - (String) The unique identifier for the resource that created the key ring.
   - `creation_date` - (Timestamp) The date the key ring created. The date format follows `RFC 3339` format.
   - `id` - (String) The unique identifier of the key ring.
