---
layout: "ibm"
page_title: "IBM: certificate_manager_order"
sidebar_current: "docs-ibm-resource-certificate-manager-order"
description: |-
  Orders and Manages Ordered Certificate.
---

# ibm\_certificate_manager_order

Orders and manages ordered certificate of Certificate manager Instance

## Example Usage

This example creates a CMS instance by enabling customer managed keys and orders a certificate.
``` hcl
resource "ibm_resource_instance" "cm" {
  name     = "test"
  location = "us-south"
  plan     = "free"
  service  = "cloudcerts"
  parameters = {
    kms_info = "{\"id\":\"<GUID OF KMS/HPCS INSTANCE>\",\"url\":\"<KMS/HPCS ENDPOINT>\"}",
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

ibm_certificate_manager_order provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for Ordering Certificate.
* `update` - (Default 10 minutes) Used for Renewing/updating Certificate.

## Argument Reference

The following arguments are supported:

* `certificate_manager_instance_id` - (Required,ForceNew,string) The CRN-based service instance ID.
* `name` - (Required,string) The display name for the ordered certificate.
* `description` - (Optional),string The optional description for the ordered certificate.
* `domains` - (Required,ForeceNew,List) An array of valid domains for the issued certificate. The first domain is the primary domain. Additional domains are secondary domains.
* `rotate_keys` - (Optional,bool) Default: False.
* `domain_validation_method` - (Optional,Default,string) Allowable values [dns-01]
* `key_algorithm` - (Optional,Default,string) Key Algorithm. Allowable values: [`rsaEncryption 2048 bit`,`rsaEncryption 4096 bit`] Default: [`rsaEncryption 2048 bit`]
* `dns_provider_instance_crn` - (Optional,string) The CRN-based instance ID of the IBM Cloud Internet Services instance that manages the domains. If not present, Certificate Manager assumes a v4 or above Callback URL notifications channel with domain validation exists.
* `auto_renew_enabled` - (Optional, Default, bool) Determines whether the certificate should be auto renewed. Default: false.
    **NOTE:** With `auto_renew_enabled`, certificates are automatically renewed 31 days before they expire. If your certificate expires in less than 31 days, you must renew it by updating `rotate_keys`. After you do so, your future certificates are renewed automatically.
* `renew_certificate`-(Optional,Default, bool) Determines whether the certificate should be renewed.Default: false.

## Attribute Reference

The following attributes are exported:

* `id` - The Crn Id of the Certificate. It is a combination of instance crn and the certificate id i.e <instance crn>:certificate:<certID>
* `issuer` - The issuer of the certificate.
* `begins_on` - The creation date of the certificate in Unix epoch time.
* `expires_on` - The expiration date of the certificate in Unix epoch time.
* `imported` - Indicates whether a certificate has a imported or not.
* `status` - The status of certificate. Possible values: [active,inactive,expired,revoked,valid,pending,failed]
* `has_previous` - Indicates whether a certificate has a previous version.
* `algorithm` - Algorithm. Allowable values: [sha256WithRSAEncryption]Default: [sha256WithRSAEncryption]

## Import

The `ibm_certificate_manager_order` resource can be imported using the `id`. The ID is the crn id of the certificate.
* **ID** is a string of the form: `crn:v1:bluemix:public:cloudcerts:us-south:a/4448261269a14562b839e0a3019ed980:8e80c112-5e48-43f8-8ab9-e198520f62e4:certificate:f543e1907a0020cfe0e883936916b336`. The id of an existing certificate is also avaiable in the UI as `Certificate CRN` under the certificate details section.

```
$ terraform import ibm_certificate_manager_order.cert <id>

$ terraform import ibm_certificate_manager_order.cert crn:v1:bluemix:public:cloudcerts:us-south:a/4448261269a14562b839e0a3019ed980:8e80c112-5e48-43f8-8ab9-e198520f62e4:certificate:f543e1907a0020cfe0e883936916b336