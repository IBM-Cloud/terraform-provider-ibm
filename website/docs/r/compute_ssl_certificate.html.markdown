---
layout: "ibm"
page_title: "IBM : compute_ssl_certificate"
sidebar_current: "docs-ibm-resource-compute-ssl-certificate"
description: |-
  Manages IBM Compute SSL Certificate.
---

# ibm\_compute_ssl_certificate

Provides an SSL certificate resource. This allows SSL certificates to be created, updated, and deleted.

For additional details, see the [IBM Cloud Infrastructure (SoftLayer) security certificates docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Security_Certificate).

## Example Usage

In the following example, you can use a certificate on file:

```hcl
resource "ibm_compute_ssl_certificate" "test_cert" {
  certificate = "${file("cert.pem")}"
  private_key = "${file("key.pem")}"
}
```

You can also use an in-line certificate:

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
* `tags` - (Optional, array of strings) Tags associated with the security certificates instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `common_name` - The common name encoded within the certificate. This name is usually a domain name.
* `create_date` - The date the certificate record was created.
* `id` - The ID of the certificate record.
* `key_size` - The size, expressed in number of bits, of the public key represented by the certificate.
* `modify_date` - The date the certificate record was last modified.
* `organization_name` - The organizational name encoded in the certificate.
* `validity_begin` - The UTC timestamp representing the beginning of the certificate's validity.
* `validity_days` - The number of days remaining in the validity period for the certificate.
* `validity_end` - The UTC timestamp representing the end of the certificate's validity period.
