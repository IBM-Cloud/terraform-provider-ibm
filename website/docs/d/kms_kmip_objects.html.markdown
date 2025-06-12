---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-kmip-objects"
description: |-
  Manages kmip objects for IBM hs-crypto and KMS.
---

# ibm_kms_kmip_adapters
Retrieves a list of KMIP Objects from a Key Protect service instance for a given KMIP adapter. The region parameter in the `provider.tf` file must be set. If region parameter is not specified, `us-south` is used by default. If the region in the `provider.tf` file is different from the Key Protect instance, the instance cannot be retrieved by Terraform and the Terraform action fails.
For more information, about KMIP as a whole, see [Using the key management interoperability protocol (KMIP)](https://cloud.ibm.com/docs/key-protect?topic=key-protect-kmip&interface=ui).


## Example usage 
Sample example to list KMIP objects in a given adapter

```terraform
resource "ibm_resource_instance" "kms_instance" {
  name     = "instance-name"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}
data "ibm_kms_kmip_adapter" "myadapter" {
    instance_id = ibm_resource_instance.kp_instance.guid
    name = "myadapter"
}
data "ibm_kms_kmip_objects" "objects_list" {
  instance_id = ibm_resource_instance.kp_instance.guid
  adapter_id = data.ibm_kms_kmip_adapter.myadapter.id
  object_state_filter = [1,2,3,4]
}
```


## Argument reference
Review the argument references that you can specify for your resource. 

- `endpoint_type` - (Optional, String) The type of the public endpoint, or private endpoint to be used for creating keys.
- `instance_id` - (Required, String) The key protect instance GUID.
- `limit` - (Optional, Integer) Limit of how many objects to be fetched.
- `offset` - (Optional, Integer) Offset of objects to be fetched.
- `show_total_count` - (Optional, Boolean) Flag to return the count of how many objects there are in total after the filter.
- `object_state_filter` - (Optional, List) A list of integers representing Object States to filter for

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `total_count` - (Integer) If show_total_count is true, this will contain the total number of objects after the State Filter
- `objects` - (List of Objects) The list of KMIP objects in an adapter

    Nested scheme for `objects`:
    - `object_id` - (String) The id of the KMIP Object
    - `object_state` - (Integer) The state of the KMIP object as an enum
    - `object_type` - (Integer) The type of the KMIP object as an enum
    - `created_by` - (String) The IBM-ID of the identity that created the resource
    - `created_at` - (String) The date the resource was created, in RFC 3339 format
    - `created_by_cert_id` - (String) The ID of the certificate that created the object
    - `updated_by` - (String) The IBM-ID of the identity that updated the resource
    - `updated_at` - (String) The date the resource was updated, in RFC 3339 format
    - `updated_by_cert_id` - (String) The ID of the certificate that updated the object
    - `destroyed_by` - (String) The IBM-ID of the identity that destroyed the resource
    - `destroyed_at` - (String) The date the resource was destroyed, in RFC 3339 format
    - `destroyed_by_cert_id` - (String) The ID of the certificate that destroyed the object