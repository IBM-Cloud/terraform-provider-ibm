---
subcategory: "Certificate Manager"
layout: "ibm"
page_title: "IBM: certificate_manager_import"
description: |-
  Imports and manages imported certificate.
---

# `ibm_certificate_manager_import`

Upload or delete a certificate in Certificate Manager. For more information, about IBM Cloud certificate manager, see [Managing certificates](https://cloud.ibm.com/docs/certificate-manager?topic=certificate-manager-managing-certificates-from-the-dashboard).


## Example usage
A Example usage to create a certificate manager service instance that enables customer managed keys and imports a certificate.


```
resource "ibm_resource_instance" "cm" {
  name     = "test"
  location = "us-south"
  plan     = "free"
  service  = "cloudcerts"
  parameters = {
    kms_info = "{-"id-":-"<GUID OF KMS/HPCS INSTANCE>-",-"url-":-"<KMS/HPCS ENDPOINT>-"}",
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


## Argument reference
Review the input parameters that you can specify for your resource. 

- `certificate_manager_instance_id` - (Required, String) The CRN-based service instance ID.
- `description` - (Optional, String) The description of the certificate.
- `name` - (Required, String) The display name for the imported certificate.
- `data`- (Required, Map) The certificate data.
	- `content` - (Required, String) The content of certificate data, escaped.
	- `intermediate` - (Optional, String) The intermediate certificate data, escaped.
  - `priv_key` - (Optional, String) The private key data, escaped.


## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `algorithm` - (String) The encryption algorithm. Valid values are `sha256WithRSAEncryption`. Default value is `sha256WithRSAEncryption`.
- `begins_on` - (String) The creation date of the certificate in UNIX epoch time.
- `expires_on` - (String) The expiration date of the certificate in UNIX epoch time.
- `has_previous`- (Bool) Indicates whether a certificate has a previous version.
- `id` - (String) The ID of the certificate.
- `imported`- (Bool) Indicates whether a certificate was imported or not.
- `issuer` - (String) The issuer of the certificate.
- `key_algorithm` - (String) The key algorithm. Valid values are `rsaEncryption 2048 bit` or `rsaEncryption 4096 bit`. Default value is `rsaEncryption 2048 bit`.
- `status` - (String) The status of certificate. Possible values are `active`, `inactive`, `expired`, `revoked`, `valid`, `pending`, and `failed`.
