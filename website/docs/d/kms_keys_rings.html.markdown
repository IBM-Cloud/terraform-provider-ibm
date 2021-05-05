---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-key-rings"
description: |-
  Manages key rings for IBM hs-crypto or key-protect.
---

# ibm\_kms_key_rings

Retreives a list of key rings from the hs-crypto or key-protect instance. Import the details of existing keyrings of hs-crypto and kms instance as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_kms_key_rings" "test" {
  instance_id = "guid-of-keyprotect-or hs-crypto-instance"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The keyprotect instance guid.
* `endpoint_type` - (Optional, string) The type of the endpoint (public or private) to be used for fetching keys.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `key_rings` - List of all Key Rings in the IBM hs-crypto or Key-protect instance.
  * `id` - The unique identifier for the key ring
  * `creation_date` - The date the key ring was created. The date format follows RFC 3339.
  * `created_by` - The unique identifier for the resource that created the key ring.