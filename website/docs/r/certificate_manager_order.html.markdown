---
subcategory: "Certificate Manager"
layout: "ibm"
page_title: "IBM: certificate_manager_order"
description: |-
  Orders and manages ordered certificate.
---

# `ibm_certificate_manager_order`

Order, renew, update, or delete a certificate in Certificate Manager. For more information, about an IBM Certificate Manager order, see [Ordering certificates](https://cloud.ibm.com/docs/certificate-manager?topic=certificate-manager-ordering-certificates).


## Example usage
A Example usage to create a Certificate Manager service instance that enables customer managed keys and orders a certificate.


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

resource "ibm_certificate_manager_order" "cert" {
  certificate_manager_instance_id = ibm_resource_instance.cm.id
  name                            = "test"
  description                     = "test description"
  domains                         = ["example.com"]
  rotate_keys                     = false
  domain_validation_method        = "dns-01"
  dns_provider_instance_crn       = ibm_cis.instance.id
}

```

## Timeouts
The following [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) are defined for this resource. 

- **Create**: The ordering of the certificate is considered failed if no response is received for 10 minutes.
- **Update**: The renewal or update of the certificate is considered failed if no response is received for 10 minutes.

## Argument reference
Review the input parameters that you can specify for your resource. 

- `auto_renew_enabled` - (Optional, Bool) Determines the certificate is auto that is renewed. Default is **false**. 
 **Note** 
 With `auto_renew_enabled` as true, certificates are automatically renewed for 31 days. If the certificate expires before 31 days. You can renew by updating `rotate_keys` to renew the certificates automatically.
- `certificate_manager_instance_id` - (Required, Forces new resource, String) The CRN of your Certificate Manager instance.
- `description` - (Optional, String) The description that you want to add to the certificate that you order.
- `domains` (Required, List) A list of valid domains for the issued certificate. The first domain is the primary domain. More domains are secondary domains.Yes.
- `domain_validation_method` - (Optional, String) The domain validation method that you want to use for your domain. The validation method is applied to analyze DNS parameters for your domain and determine the domain health and quality standards that your domain meets. Supported parameters are `dns-01`.
- `dns_provider_instance_crn` - (Optional, String) The CRN based instance ID of the IBM Cloud Internet Services instance that manages the domains. If not present, Certificate Manager assumes that a `v4` or callback URL notifications channel with domain validation exists.
- `key_algorithm` - (Optional, String) The encryption algorithm key that you want to use for your certificate. Supported values are `rsaEncryption 2048 bit`, and `rsaEncryption 4096 bit`. If you do not provide an algorithm, `rsaEncryption 2048 bit` is used by default.
- `name` - (Required, String) The name for the certificate that you want to order.
- `renew_certificate` - (Optional, Bool) Determines the certificate to renew. Default value is **false**.
- `rotate_keys` - (Optional, Bool) Default value is **false**.


## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `algorithm` - (String) The encryption algorithm. Valid values are `sha256WithRSAEncryption`.
- `begins_on` - (String) The creation date of the certificate in UNIX epoch time.
- `expires_on` - (String) The expiration date of the certificate in UNIX epoch time.
- `has_previous`- (Bool) Indicates whether a certificate has a previous version.
- `id` - (String) The ID of the certificate.
- `imported`- (Bool) Indicates whether a certificate was imported or not.
- `issuer` - (String) The issuer of the certificate.
- `status` - (String) The status of certificate. Possible values are `active`, `inactive`, `expired`, `revoked`, `valid`, `pending`, and `failed`.


## Import
The `ibm_certificate_manager_order` resource can be imported by using CRN ID of the certificate. The ID is available in the console as `Certificate CRN` in the certificate details section.

* **ID** is a string of the form: `crn:v1:bluemix:public:cloudcerts:us-south:a/4448261269a14562b839e0a3019ed980:8e80c112-5e48-43f8-8ab9-e198520f62e4:certificate:f543e1907a0020cfe0e883936916b336`.


**Syntax** 

```
terraform import ibm_certificate_manager_order.cert <id>

```
**Example**

```
terraform import ibm_certificate_manager_order.cert crn:v1:bluemix:public:cloudcerts:us-south:a/4448261269a14562b839e0a3019ed980:8e80c112-5e48-43f8-8ab9-e198520f62e4:certificate:f543e1907a0020cfe0e883936916b336
```

