---

subcategory: "Certificate Manager"
layout: "ibm"
page_title: "IBM: certificate_manager_import"
description: |-
  Imports and Manages Imported Certificate.
---

# ibm\_certificate_manager_import

Imports and manages imported certificate of Certificate Manager Instance.

## Example Usage
This example creates a CMS instance by enabling customer managed keys and imports a certificate.
``` hcl
resource "ibm_resource_instance" "cm" {
  name     = "test"
  location = "us-south"
  plan     = "free"
  service  = "cloudcerts"
  parameters = {
    kms_info = "{\"id\":\"<GUID OF KMS/HPCS INSTANCE>\",\"url\":\"<KMS/HPCS ENDPOINT>\"}",
    tek_id   = "CRN OF KMS/HPCS KEY",
  }
}

resource "ibm_certificate_manager_import" "cert" {
  certificate_manager_instance_id = ibm_resource_instance.cm.id
  name                            = "test"
  description="string"
  data = {
    content = file(var.certfile_path)
  }
}
```

## Argument Reference

The following arguments are supported:

* `certificate_manager_instance_id` - (Required,string) The CRN-based service instance ID.
* `name` - (Required,string) The display name for the imported certificate.
* `data` - (Required,Map) The certificate data.
    * `content` - (Required,string) The content of certificate data, escaped.
    * `priv_key` - (Optional,string) The private key data, escaped.
    * `intermediate` - (Optional,string) The intermediate certificate data, escaped.
* `description` - (Optional),string The optional description for the imported certificate.


## Attribute Reference

The following attributes are exported:

* `id` - The Id of the Certificate
* `issuer` - The issuer of the certificate.
* `begins_on` - The creation date of the certificate in Unix epoch time.
* `expires_on` - The expiration date of the certificate in Unix epoch time.
* `imported` - Indicates whether a certificate has a imported or not.
* `status` - The status of certificate. Possible values: [active,inactive,expired,revoked,valid,pending,failed]
* `has_previous` - Indicates whether a certificate has a previous version.
* `key_algorithm` - Key Algorithm. Allowable values: [rsaEncryption 2048 bit rsaEncryption 4096 bit] Default: [rsaEncryption 2048 bit]
* `algorithm` - Algorithm. Allowable values: [sha256WithRSAEncryption]Default: [sha256WithRSAEncryption]
