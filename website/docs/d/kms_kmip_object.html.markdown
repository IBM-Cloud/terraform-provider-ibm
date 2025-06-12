---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-kmip-object"
description: |-
  Manages a kmip object for IBM hs-crypto and KMS.
---

# ibm_kms_kmip_adapters
Retrieves a KMIP Object from a Key Protect service instance. The region parameter in the `provider.tf` file must be set. If region parameter is not specified, `us-south` is used by default. If the region in the `provider.tf` file is different from the Key Protect instance, the instance cannot be retrieved by Terraform and the Terraform action fails.
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
data "ibm_kms_kmip_object" "object" {
  instance_id = ibm_resource_instance.kp_instance.guid
  adapter_id = data.ibm_kms_kmip_adapter.myadapter.id
  object_id = "<object-UUID>"
}
```


## Argument reference
Review the argument references that you can specify for your resource. 

- `endpoint_type` - (Optional, String) The type of the public endpoint, or private endpoint to be used for creating keys.
- `instance_id` - (Required, String) The key protect instance GUID.
- `adapter_id` - (Optional, String) The UUID of the KMIP adapter to be fetched. Mutually exclusive argument with `adapter_name`. One has to be given.
- `adapter_name` - (Optional, String) The name of the KMIP adapter to be fetched. Mutually exclusive argument with `adapter_id`. One has to be given.
- `object_id` - (Required, String) The id of the KMIP object to be fetched

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `object_state` - (Integer) The state of the KMIP object
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