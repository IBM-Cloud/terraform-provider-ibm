---
layout: "ibm"
page_title: "IBM: ibm_cis_certificate_order"
sidebar_current: "docs-ibm-resource-cis-page-rule"
description: |-
  Provides a IBM CIS Certificate Order resource.
---

# ibm_cis_certificate_order

Provides a IBM CIS Certificate Order resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to order and delete dedicated certificates of a domain of a CIS instance

## Example Usage

```hcl
resource "ibm_cis_certificate_order" "test" {
	cis_id    = data.ibm_cis.cis.id
	domain_id = data.ibm_cis_domain.cis_domain.domain_id
	hosts     = ["example.com"]
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance.
- `domain_id` - (Required,string) The ID of the domain.
- `hosts` - (Required,list(string)) The hosts for which the certificates to be ordered.

## Attributes Reference

The following attributes are exported:

- `id` - The record ID. It is a combination of <`certificate_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
- `certificate_id` - The certificate ID.
- `status` - The certificate status.

## Import

The `ibm_cis_certificate_order` resource can be imported using the `id`. The ID is formed from the `Certificate ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Certificate ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`.

```
$ terraform import ibm_cis_certificate_order.myorg <certificate_id>:<domain-id>:<crn>

$ terraform import ibm_cis_certificate_order.myorg certificate_order 48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
