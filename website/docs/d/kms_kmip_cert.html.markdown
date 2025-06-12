---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-kmip-cert"
description: |-
  Manages a kmip adapter's certificate for IBM hs-crypto and KMS.
---

# ibm_kms_kmip_client_cert
Retrieves a KMIP Client Certificate from a Key Protect service instance based on the certificate name or ID. The region parameter in the `provider.tf` file must be set. If region parameter is not specified, `us-south` is used by default. If the region in the `provider.tf` file is different from the Key Protect instance, the instance cannot be retrieved by Terraform and the Terraform action fails.
For more information, about KMIP as a whole, see [Using the key management interoperability protocol (KMIP)](https://cloud.ibm.com/docs/key-protect?topic=key-protect-kmip&interface=ui).


## Example usage 
Sample example to retrieve a KMIP client certificate as a data source.

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

resource "ibm_kms_kmip_client_cert" "mycert" {
  instance_id = ibm_resource_instance.kp_instance.guid
  adapter_id = ibm_kms_kmip_adapter.myadapter.adapter_id
  certificate = file("${path.module}/certificate.pem")
  name = var.kmip_cert_name
}

data "ibm_kms_kmip_client_cert" "mycert_byname" {
  instance_id = ibm_resource_instance.kp_instance.guid
  adapter_name = ibm_kms_kmip_adapter.myadapter.name
  name = ibm_kms_kmip_client_cert.mycert.name
}

data "ibm_kms_kmip_client_cert" "mycert_byid" {
  instance_id = ibm_resource_instance.kp_instance.guid
  adapter_id = ibm_kms_kmip_adapter.myadapter.adapter_id
  cert_id = ibm_kms_kmip_client_cert.mycert.cert_id
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `endpoint_type` - (Optional, String) The type of the public endpoint, or private endpoint to be used for creating keys.
- `instance_id` - (Required, String) The key protect instance GUID.
- `adapter_id` - (Optional, String) The UUID of the KMIP adapter to be fetched. Mutually exclusive argument with `adapter_name`. One has to be given.
- `adapter_name` - (Optional, String) The name of the KMIP adapter to be fetched. Mutually exclusive argument with `adapter_id`. One has to be given.
- `cert_id` - (Optional, String) The UUID of the KMIP client certificate to be fetched. Mutually exclusive argument with `name`. One has to be given.
- `name` - (Optional, String) The name of the KMIP client certificate to be fetched. Mutually exclusive argument with `cert_id`. One has to be given.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `certificate` - (String) The contents of the KMIP client certificate.
- `created_at` - (String) The date the resource was created, in RFC 3339 format
- `created_by` - (String) The IBM-ID of the identity that created the resource