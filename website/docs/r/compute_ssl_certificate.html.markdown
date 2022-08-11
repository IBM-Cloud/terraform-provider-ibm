---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : compute_ssl_certificate"
description: |-
  Manages IBM Compute SSL certificate.
---

# ibm_compute_ssl_certificate
Create, update, and delete a SSL certificate. For more information, about SSL certificate, see [accessing SSL certificates](https://cloud.ibm.com/docs/ssl-certificates?topic=ssl-certificates-accessing-ssl-certificates).

**Note**

For more information, see the [IBM Cloud Classic Infrastructure (SoftLayer) security certificates docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Security_Certificate).

## Example usage
The example to use a certificate on file:

```terraform
resource "ibm_compute_ssl_certificate" "test_cert" {
  certificate = file("cert.pem")
  private_key = file("key.pem")
}
```

The example to use an in-line certificate:

```terraform
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

## Argument reference
Review the argument references that you can specify for your resource.

- `certificate` - (Required, Forces new resource, String)The certificate provided publicly to clients requesting identity credentials.
- `intermediate_certificate`- (Optional, Forces new resource, String) The certificate from the intermediate certificate authority, or chain certificate, that completes the chain of trust. Required when clients only trust the root certificate.
- `private_key` - (Required, Forces new resource, String)The private key in the key/certificate pair.
- `tags`- (Optional, Array of Strings) Tags associated with the security certificates instance.  **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `common_name`- (String) The common name encoded within the certificate. This name is usually a domain name.
- `create_date`- (String) The date the certificate record was created.
- `id`-The ID of the certificate record.
- `key_size`- (String) The size, expressed in number of bits, of the public key represented by the certificate.
- `modify_date`- (String) The date the certificate record was last modified.
- `organization_name`- (String) The organizational name encoded in the certificate.
- `validity_begin`- (String) The UTC timestamp representing the beginning of the certificate's validity.
- `validity_days`- (String) The number of days remaining in the validity period for the certificate.
- `validity_end`- (String) The UTC timestamp representing the end of the certificate's validity period.

