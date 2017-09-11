---
layout: "ibm"
page_title: "IBM : compute_ssl_certificate"
sidebar_current: "docs-ibm-resource-compute-ssl-certificate"
description: |-
  Manages IBM Compute SSL Certificate.
---

# ibm\_compute_ssl_certificate

Create, update, and destroy [Bluemix Infrastructure (SoftLayer) security certificates](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Security_Certificate).

## Example Usage

Using certificates on file:

```hcl
resource "ibm_compute_ssl_certificate" "test_cert" {
  certificate = "${file("cert.pem")}"
  private_key = "${file("key.pem")}"
}
```

Example with the in-line certificates:

```hcl
resource "ibm_compute_ssl_certificate" "test_cert" {
    certificate = <<EOF
[......] # cert contents
-----END CERTIFICATE-----
    EOF

    private_key = <<EOF
-----BEGIN RSA PRIVATE KEY-----
[......] # cert contents
-----END RSA PRIVATE KEY-----
    EOF
}
```

## Argument Reference

The following arguments are supported:

* `certificate` - (Required, string) The certificate provided publicly to clients requesting identity credentials.
* `intermediate_certificate` - (Optional, string) The certificate from the intermediate certificate authority, or chain certificate, that completes the chain of trust. Required when clients only trust the root certificate.
* `private_key` - (Required, string) The private key in the key/certificate pair.
* `tags` - (Optional, array of strings) Set tags on the security certificates instance.

**NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attributes Reference

The following attributes are exported:

* `common_name` - The common name (usually a domain name) encoded within the certificate.
* `create_date` - The date the certificate record was created.
* `id` - The ID of the certificate record.
* `key_size` - The size (number of bits) of the public key represented by the certificate.
* `modify_date` - The date the certificate record was last modified.
* `organization_name` - The organizational name encoded in the certificate.
* `validity_begin` - The UTC timestamp representing the beginning of the certificate's validity.
* `validity_days` - The number of days remaining in the validity period for the certificate.
* `validity_end` - The UTC timestamp representing the end of the certificate's validity period.
