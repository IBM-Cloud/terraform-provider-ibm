---
subcategory: "Certificate Manager"
layout: "ibm"
page_title: "IBM: certificate_manager_certificate"
description: |-
  Reads details of a certificate from a Certificate Manager Instance
---

# ibm\_certificate_manager_certificate

Imports a read only copy of an existing Certificate Instance resource and lists all the certificates for the given name.

## Example Usage

```hcl
data "ibm_resource_instance" "cm" {
    name     = "testname"
    location = "us-south"
    service  = "cloudcerts"
}
data "ibm_certificate_manager_certificate" "source_certificate"{
    certificate_manager_instance_id=data.ibm_resource_instance.cm.id
    name = "certificate name"
}
```

## Argument Reference

The following arguments are supported:

* `certificate_manager_instance_id` - (Required,string) The CRN-based service instance ID.
* `name` - (Required,string) The display name for the certificate.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The Id of the Certificate. It is a combination of <`name`>:<`certificate_manager_instance_id`>
* `certificate_details` - List of certificates for the given name.
    * `cert_id` -  The CRN-based Certificate ID.
    * `name` - The display name for the certificate.
    * `domains` -  An array of valid domains for the issued certificate. The first domain is the primary domain. Additional domains are secondary domains.
    * `data` -  The certificate data.
        * `content` -  The content of certificate data, escaped.
        * `priv_key` -  The private key data, escaped.
        * `intermediate` -  The intermediate certificate data, escaped.
    * `issuer` - The issuer of the certificate.
    * `begins_on` - The creation date of the certificate in Unix epoch time.
    * `expires_on` - The expiration date of the certificate in Unix epoch time.
    * `imported` - Indicates whether a certificate has imported or not.
    * `status` - The status of certificate.
    * `has_previous` - Indicates whether a certificate has a previous version.
    * `key_algorithm` - Key Algorithm of a certificate.
    * `algorithm` - Algorithm of a certificate.
    * `serial_number` - The certificate serial number
    * `issuance_info` - Issuance Info of Certificate.
        * `status` - The status of certificate.
        * `ordered_on` - The date the certificate was ordered.
        * `code` - Code of Certificate.
        * `additional_info` - The Additional Info of certificate.
