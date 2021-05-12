---
subcategory: "Certificate Manager"
layout: "ibm"
page_title: "IBM: certificate_manager_certificate"
description: |-
  Reads details of a certificate from a Certificate Manager instance
---

# `ibm_certificate_manager_certificate`

Retrieve the details of an existing certificate instance resource and lists all the certificates. For more information, about Certificate Manager, see [Getting started with Certificate Manager](https://cloud.ibm.com/docs/certificate-manager?topic=certificate-manager-getting-started).


## Example usage

```
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

## Argument reference
Review the input parameters that you can specify for your resource. 
 
- `certificate_manager_instance_id` - (Required, String) The CRN of the certificate manager service instance.
- `name` - (Required, String) The display name for the certificate.

## Attribute reference
Review the output parameters that you can access after your resource is created. 


- `certificate_details` - (String) List of certificates for the provided name.
	- `algorithm` - (String) The algorithm that is used for the certificate. 
	- `begins_on`- (Timestamp) The timestamp when the certificate was created in UNIX epoch time format. 
	- `cert_id` - (String) The CRN based certificate ID. 
    - `domains` - (Array) A list of domains that the certificate is associated with. The first domain is referred to as the primary domain. Any more domains are referred to as secondary domains.
	- `data` - (String) The certificate data.
	   - `content` - (String) The content of certificate data, escaped.
	   - `priv_key` - (String) The private key data, escaped.
	   - `intermediate` - (String) The intermediate certificate data, escaped.
	- `expires_on`- (Date) The date when the certificate expires in UNIX epoch time format.	
	- `has_previous`- (Bool) If set to **true**, the certificate has a previous version. 
	- `issuer` - (String) The issuer of the certificate.
	- `issuance_info` - (List of Objects) The issuance information of the certificate. 
		- `additional_info` - (String) Any more information for the certificate. 
		- `code` - (String) The code of the certificate.
		- `ordered_on`- (Date) The date when the certificate was ordered.
		- `status` - (String) The status of the certificate.
	- `imported`- (Bool) If set to **true**, indicates whether a certificate is imported.
	- `key_algorithm` - (String) The key algorithm of the certificate.
	- `name` - (String) The name of the certificate.
	- `serial_number` - (String) The serial number of the certificate.
	- `status` - (String) The status of the certificate.
- `id` - (String) The ID of the certificate that is managed in Certificate Manager. The ID is composed of `<certificate_manager_instance_ID>:<certificate_ID>`.
