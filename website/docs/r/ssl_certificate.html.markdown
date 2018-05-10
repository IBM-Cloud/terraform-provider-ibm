---
layout: "ibm"
page_title: "IBM : ssl_certificate"
sidebar_current: "docs-ibm-resource-ssl-certificate"
description: |-
  Manages IBM Compute SSL Certificate.
---

# ibm\_ssl_certificate

Provides an SSL certificate resource. This allows SSL certificates to be requested, and delete request for ssl certificates.

For additional details, see the [IBM Cloud (SoftLayer) security certificates Request docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Security_Certificate/Request).

## Example Usage

In the following example, you can use a certificate on file:

```hcl
resource "ibm_ssl_certificate" "test_cert" {
  csr= "-----BEGIN CERTIFICATE
  REQUEST-----\nMIIC2jCCAcICAQAwgYAxCzAJBgNVBAYTAklOMRMwEQYDVQQIDApNYWhhcmFzaHRh\nMQ0wCwYDVQQHDARQdW5lMRAwDgYDVQQKDAdJQk1QdW5lMQwwCgYDVQQLDANJQk0x\nFDASBgNVBAMMC2libS5wdW5lLmluMRcwFQYJKoZIhvcNAQkBFghyYkBncy5pbjCC\nASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALZTfLDkU/q8ory+a+aBIAB7\nezyMbF4iqEpO3mVmlIfENBRq/MomhTiYID2jVWhPS4KytXa9bANgAG+NN5zMxqgH\nwUVa4fO2Th//o9RdRPZQcC3IGR+b3tpGSV+KxjsW9XV82lQp6l8HxfSBo12PZcwU\nlu53BXMiVKU6ZHSjz6eDF4MGTuPM1QWs/oexb9y6Yqqj2IicrkviJMuKhh0FN6H7\nXT2MHIF8OHT1xhs8w7U1SdUwhBa5dHvj7vqi33DU0p3s8JdXlzE9PUXKuAzXav1D\nngwE1iT5KFHkEFUr7plH5wGsy/y4tnu2xVzfGMyECS9EYmUtOA5M5RpqF5DdXGEC\nAwEAAaAUMBIGCSqGSIb3DQEJAjEFDANnc2wwDQYJKoZIhvcNAQELBQADggEBAGB7\n3lv/6fSn9rgTiHszLL9pOU9ytjOVrhNjFjDzQL73VQ0+Isb7aPHnWrLz4kT9m/60\nmgy/dHsOIF8KRP1LpOs5BYwlstD3Ss57XR8GatnrLMN4lZCjacL6A8RPhwr3x29W\nMyFntvu2caAL4ZpZpWMKHtoemXijCiFXa9Z4pZFBk4V7k0/DIEXEeyYazsSaXeTw\nXr4IFPmk7VS/NkLAht2hRhllN5NHGf/gzTsmgrgKclXtf1Z7EotnDTTIt0dFVtk1\nVX2Z7kvx9/QWbDVhPEz2uOrJnCoAm+0OpQfFc4THcP0uv0Y49B3WUG89mAjlWQKa\nU7hhc8gZ77+eaBQKD6k=\n-----END CERTIFICATE REQUEST-----"
	address1= "abc"
	address2= "xyz"
	city="pune"
	country_code= "IN"
	state="MH"
	postal_code= "411007"
	org_name= "IBMPune"
	phone_no= "9908277823"
	first_name= "pqr"
	ssl_type="SSL_CERTIFICATE_QUICKSSL_PREMIUM_2_YEAR"
	last_name= "zende"
	renewal= true
	server_count= 1
	server_type= "apache2"
	validity_month= 24
	email= "admin@pune.in"
	title= "no need to tittle it"
}
```


## Argument Reference

The following arguments are supported:

* `csr` - (Required, string) The Certificate Signing Request which is specially formatted encrypted message sent from a Secure Sockets Layer (SSL) digital certificate applicant to a certificate authority.
* `ssl_type` - (Required, string) The ssl certificate type.
* `server_type` - (Required, string) The server type for which we are requesting ssl certficate.
* `server_count` - (Required, string) The number of servers with provided server tye .
* `validity_month` - (Required, integer) The validity of ssl certificate in months it should be multiple of 12.
* `address1` - (Required, string) The address of client who is requesting for ssl certificate.
* `address2` - (optional, string) The address of client who is requesting for ssl certificate.
* `city` - (Required, string) The city of client who is requesting for ssl certificate .
* `postal_code` - (Required, integer) The postal code for the city of client.
* `state` - (Required, string) The two letter state code of client who is requesting for ssl certificate. Allowed value for country which doesn't have states is `OT`.
* `country_code` - (Required, string) The two letter country code of client.
* `org_name` - (Required, string) The organization name of client.
* `first_name` - (Required, string) The first name of client.
* `last_name` - (Required, string) The last name of client.
* `title` - (Required, string) The title for ssl certificate request.
* `email` - (Required, string) The email of approver to approve ssl certificate request.
