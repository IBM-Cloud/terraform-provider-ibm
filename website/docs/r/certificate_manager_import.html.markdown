---
layout: "ibm"
page_title: "IBM: certificate_manager_Import"
sidebar_current: "docs-ibm-resource-certificate-manager-import"
description: |-
  Imports and Manages Certificates.
---

# ibm\_certificate_manager

Provides a certificate manager. This allows certificates to be imported, updated, and deleted.

## Example Usage

```hcl
provider "ibm"
{
}
resource "ibm_certificate_manager_import" "cert"{
    certificate_manager_instance_id = "crn_based_instance_id"
     name = "string"
     description="string"
     data = [{
       content = "${file(var.certfile_path)}"
       priv_key = ""
       intermediate = ""
     }],

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
