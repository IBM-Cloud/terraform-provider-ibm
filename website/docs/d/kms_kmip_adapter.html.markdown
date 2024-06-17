---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-kmip-adapter"
description: |-
  Manages a kmip adapter for IBM hs-crypto and KMS.
---

# ibm_kms_kmip_adapter
Retrieves a KMIP adapter from a Key Protect service instance based on a given name or ID. The region parameter in the `provider.tf` file must be set. If region parameter is not specified, `us-south` is used by default. If the region in the `provider.tf` file is different from the Key Protect instance, the instance cannot be retrieved by Terraform and the Terraform action fails.
For more information, about KMIP as a whole, see [Using the key management interoperability protocol (KMIP)](https://cloud.ibm.com/docs/key-protect?topic=key-protect-kmip&interface=ui).


## Example usage 
Sample example to retrieve a KMIP adapter as a data source.

```terraform
resource "ibm_resource_instance" "kms_instance" {
  name     = "instance-name"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

resource "ibm_kms_kmip_adapter" "myadapter" {
  instance_id = ibm_resource_instance.kms_instance.guid
  profile = "native_1.0"
  profile_data = {
    "crk_id" = ibm_kms_key.key.key_id
  }
  description = "adding a description"
}

data "ibm_kms_kmip_adapter" "myadapter_byname" {
    instance_id = ibm_resource_instance.kms_instance.guid
    name = ibm_kms_kmip_adapter.myadapter.name
}

data "ibm_kms_kmip_adapter" "myadapter_byid" {
    instance_id = ibm_resource_instance.kms_instance.guid
    adapter_id = ibm_kms_kmip_adapter.myadapter.adapter_id
}

```


## Argument reference
Review the argument references that you can specify for your resource. 

- `endpoint_type` - (Optional, String) The type of the public endpoint, or private endpoint to be used for creating keys.
- `instance_id` - (Required, String) The key protect instance GUID.
- `adapter_id` - (Optional, String) The UUID of the KMIP adapter to be fetched. Mutually exclusive argument with `name`. One has to be given.
- `name` - (Optional, String) The name of the KMIP adapter to be fetched. Mutually exclusive argument with `adapter_id`. One has to be given.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `description` - (String) The description of the KMIP adapter.
- `profile` - (String) The profile of the KMIP adapter.
- `profile_data` - (Map) The profile data of the KMIP adapter.
- `created_at` - (String) The date the resource was created, in RFC 3339 format
- `created_by` - (String) The IBM-ID of the identity that created the resource
- `updated_at` - (String) The date the resource was updated, in RFC 3339 format
- `updated_by` - (String) The IBM-ID of the identity that updated the resource