---
subcategory: "Hyper Protect Crypto Services"
layout: "ibm"
page_title: "IBM : Hyper Protect Crypto Services instance"
description: |-
  Get information on an IBM Cloud Hyper Protect Crypto Services instance.
---

# ibm_hpcs

Imports a read only copy of an existing Hyper Protect Crypto Services resource.

## Example usage

```terraform
data "ibm_hpcs" "hpcs_instance" {
  name    = "test"
}
```

## Argument reference

The following arguments are supported:

* `name` - (Required, String) The name used to identify the Hyper Protect Crypto Services instance in the IBM Cloud UI.
* `resource_group_id` - (Optional, String) The ID of the resource group.
* `location` - (Optional, String) The location for this Hyper Protect Crypto Services instance

## Attribute reference

In addition to all arguments above, the following attributes are exported:

* `crn` - (String) The CRN of the Hyper Protect Crypto Services instance.
* `extensions` - (List) The extended metadata as a map associated with the resource instance.
* `guid` - (String) Unique identifier of resource instance.
* `hsm_info` - (List) HSM config of the crypto units.
  Nested scheme for `hsm_info`:
  * `admins` - (List) List of Admins for crypto units.
    Nested scheme for `admins`:
      * `name` - (String) Name of an admin.
      * `ski` - (String) Subject key identifier of the administrator signature key.
  * `current_mk_status` - (String) Status of current master key register.
  * `current_mkvp` - (String) Current master key register verification pattern.
  * `hsm_id` - (String) The HSM ID.
  * `hsm_location` - (String) The HSM Location.
  * `hsm_type` - (String) The HSM Type.
  * `new_mk_status` - (String) Status of new master key register.
  * `new_mkvp` - (String) New master key register verification pattern.
  * `revocation_threshold` - (Integer) Revocation threshold for crypto units.
  * `signature_threshold`- (Integer) Signature threshold for crypto units.
* `failover_units` - (Integer) The number of failover crypto units for your service instance.
* `id` - (String) The unique identifier CRN of this Hyper Protect Crypto Services instance.
* `plan` - (String) The pricing plan for your service instance.
* `service` - (String) The service type `hs-crypto` of an instance.
* `service_endpoints` - (String) The network access to your service instance. Possible values are **public-and-private**, **private-only**.
* `status` - (String) Status of the Hyper Protect Crypto Services instance.
* `units` -(Integer) The number of operational crypto units for your service instance.
