---
layout: "ibm"
page_title: "IBM: certificate_manager_Order_Certificate"
sidebar_current: "docs-ibm-resource-certificate-manager-Order"
description: |-
  Orders and Manages Certificates.
---

# ibm\_certificate_manager

Provides a certificate manager. This allows certificates to be ordered, renewed, updated, and deleted.

## Example Usage

```hcl
resource "ibm_certificate_manager_order" "cert" {
  certificate_manager_instance_id = "${ibm_resource_instance.cm.id}"
  name                            = "test"
  description                     = "test description"
  domains                         = ["example.com"]
  rotate_keys                     = false
  domain_validation_method        = "dns-01"
  dns_provider_instance_crn       = "${ibm_cis.instance.id}"
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
* `rotate_keys` - (Optional,string) Default: False.
* `domain_validation_method` - (Optional,Default,string) Allowable values dns-01]
* `dns_provider_instance_crn` - (Optional,string) The CRN-based instance ID of the IBM Cloud Internet Services instance that manages the domains. If not present, Certificate Manager assumes a v4 or above Callback URL notifications channel with domain validation exists.



## Attribute Reference

The following attributes are exported:

* `id` - The Id of the Certificate
* `issuer` - The issuer of the certificate.
* `begins_on` - The creation date of the certificate in Unix epoch time.
* `expires_on` - The expiration date of the certificate in Unix epoch time.
* `imported` - Indicates whether a certificate has a imported or not.
* `status` - The status of certificate. Possible values: [active,inactive,expired,revoked,valid,pending,failed]
* `has_previous` - Indicates whether a certificate has a previous version.
* `key_algorithm` - Key Algorithm. Allowable values: [rsaEncryption 2048 bit rsaEncryption 4096 bit] Default: [rsaEncryption 2048 bit]
* `algorithm` - Algorithm. Allowable values: [sha256WithRSAEncryption]Default: [sha256WithRSAEncryption]
