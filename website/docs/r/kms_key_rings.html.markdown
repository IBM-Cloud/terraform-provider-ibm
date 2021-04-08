---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-key-rings"
description: |-
  Manages key rings for IBM hs-crypto and kms.
---

# ibm\_kms_key_rings

Provides a resource to manage key rings for hs-crypto and key-protect services. This allows key rings to be created, and deleted. Key rings created through this resource can be used to associate to kms key resource when a standard or a root key gets created or imported.


## Example usage to create a Key Ring and associate a kms key.

```hcl
resource "ibm_resource_instance" "kms_instance" {
  name     = "instance-name"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}
resource "ibm_kms_key_rings" "keyRing" {
  instance_id = ibm_resource_instance.kms_instance.guid
  key_ring_id = "key-ring-id"
}
resource "ibm_kms_key" "key" {
  instance_id = ibm_resource_instance.kp_instance.guid
  key_name       = "key"
  key_ring_id = ibm_kms_key_rings.keyRing.key_ring_id
  standard_key   = false
  payload = "aW1wb3J0ZWQucGF5bG9hZA=="
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, Forces new resource, string) The hs-crypto or key-protect instance guid.
* `key_ring_id` - (Required, Forces new resource, string) The ID that identifies the key ring. Each ID is unique only within the given instance and is not reserved across the Key Protect service.
Constraints: 2 ≤ length ≤ 100, Value must match regular expression ^[a-zA-Z0-9-]*$
* `endpoint_type` - (Optional, Forces new resource, string) The type of the endpoint (public or private) to be used for creating keys.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Unique Identifier for the terraform resource.
* `key_ring_id` - The key ring ID.
