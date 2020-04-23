---
layout: "ibm"
page_title: "IBM: certificate_manager_Certificates"
sidebar_current: "docs-ibm-resource-certificate-manager-certificates"
description: |-
  Lists Certificates of a Certificate Manager Instance
---

# ibm\_certificate_manager_certificates

Imports a read only copy of an existing Certificate Instance resource and lists all the certificates.

## Example Usage

```hcl
data "ibm_resource_instance" "cm" {
    name     = "testname"
    location = "us-south"
    service  = "cloudcerts"
}
data "ibm_certificate_manager_certificates" "certs"{
    certificate_manager_instance_id="${data.ibm_resource_instance.cm.id}"
}
```

## Argument Reference

The following arguments are supported:

* `certificate_manager_instance_id` - (Required,string) The CRN-based service instance ID.

## Attribute Reference

The following attributes are exported:

* `id` - The Id of the Certificate. It is a combination of <`certificate_manager_instance_id`>:<`CertificateID`>
* `name` - The display name for the certificate.
* `domains` -  An array of valid domains for the issued certificate. The first domain is the primary domain. Additional domains are secondary domains.
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
