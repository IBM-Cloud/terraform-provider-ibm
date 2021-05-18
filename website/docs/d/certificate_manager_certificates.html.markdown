---
subcategory: "Certificate Manager"
layout: "ibm"
page_title: "IBM: certificate_manager_certificates"
description: |-
  Lists certificates of a Certificate Manager instance
---

# `ibm_certificate_manager_certificates`

Retrieve the details of one or lists all certificates that are managed by your Certificate Manager service instance resource. For more information, about Certificate Manager, see [Managing certificates from the dashboard](https://cloud.ibm.com/docs/certificate-manager?topic=certificate-manager-managing-certificates-from-the-dashboard).


## Example usage

```
data "ibm_resource_instance" "cm" {
    name     = "testname"
    location = "us-south"
    service  = "cloudcerts"
}
data "ibm_certificate_manager_certificates" "certs"{
    certificate_manager_instance_id=data.ibm_resource_instance.cm.id
}
```

## Argument reference
Review the input parameters that you can specify for your resource. 

- `certificate_manager_instance_id` - (Required, String) The CRN based of the certificate manager service instance ID.

## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `algorithm` - (String) The Algorithm of a certificate.
- `begins_on` - (String) The creation date of the certificate in UNIX epoch time.
- `domains` - (String) An array of valid domains for the issued certificate. The first domain is the primary domain. extra domains are secondary domains.
- `expires_on` - (String) The expiration date of the certificate in Unix epoch time.
- `has_previous` - (String) Indicates whether a certificate has a previous version.
- `id` - (String) The ID of the certificate that is managed in certificate manager. The ID is composed of `<certificate_manager_instance_ID>:<certificate_ID>`.
- `issuer` - (String) The issuer of the certificate.
- `issuance_info` - (String) The issuance information of a certificate.
	-  `additional_info` - (String) The extra information of a certificate.
	-   `status` - (String) The status of a certificate.
	-   `ordered_on` - (String) The certificate ordered date.
	-   `code` - (String) The code of a certificate.
- `imported` - (String) Indicates whether a certificate has imported or not.
- `key_algorithm` - (String) The key algorithm of a certificate.
- `name` - (String) The display name of the certificate.
- `serial_number` - (String) The serial number of a certificate.
- `status` - (String) The status of a certificate.
