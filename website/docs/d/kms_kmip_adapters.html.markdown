---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-kmip-adapters"
description: |-
  Manages kmip adapters for IBM hs-crypto and KMS.
---

# ibm_kms_kmip_adapters
Retrieves a list of KMIP adapters from a Key Protect service instance. The region parameter in the `provider.tf` file must be set. If region parameter is not specified, `us-south` is used by default. If the region in the `provider.tf` file is different from the Key Protect instance, the instance cannot be retrieved by Terraform and the Terraform action fails.
For more information, about KMIP as a whole, see [Using the key management interoperability protocol (KMIP)](https://cloud.ibm.com/docs/key-protect?topic=key-protect-kmip&interface=ui).


## Example usage 
Sample example to retrieve all KMIP adapters in a given KMIP instance

```terraform
resource "ibm_resource_instance" "kms_instance" {
  name     = "instance-name"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}
data "ibm_kms_kmip_adapters" "myadapter" {
    instance_id = ibm_resource_instance.kp_instance.guid
}
```


## Argument reference
Review the argument references that you can specify for your resource. 

- `endpoint_type` - (Optional, String) The type of the public endpoint, or private endpoint to be used for creating keys.
- `instance_id` - (Required, String) The key protect instance GUID.
- `limit` - (Optional, Integer) Limit of how many adapters to be fetched.
- `offset` - (Optional, Integer) Offset of adapters to be fetched.
- `show_total_count` - (Optional, Boolean) Flag to return the count of how many adapters there are in total.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `total_count` - (Integer) If show_total_count is true, this will contain the total number of adapters
- `adapters` - (List of Objects) The IBM-ID of the identity that created the resource

    Nested scheme for `adapters`:
    - `adapter_id` - (String) The UUID of the KMIP adapter.
    - `name` - (String) The name of the KMIP adapter.
    - `description` - (String) The description of the KMIP adapter.
    - `profile` - (String) The profile of the KMIP adapter.
    - `profile_data` - (Map) The profile data of the KMIP adapter.
    - `created_at` - (String) The date the resource was created, in RFC 3339 format
    - `created_by` - (String) The IBM-ID of the identity that created the resource
    - `updated_at` - (String) The date the resource was updated, in RFC 3339 format
    - `updated_by` - (String) The IBM-ID of the identity that updated the resource