---
layout: "ibm"
page_title: "IBM: ibm_cis_dns_record"
sidebar_current: "docs-ibm-cis-dns-record"
description: |-
  Provides a IBM DNS Record resource.
---

# ibm_cis_dns_record

Provides a IBM DNS Record resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. 

## Example Usage

```hcl
# Add a DNS record to the domain
resource "ibm_cis_dns_record" "example" {
  cis_id = "${ibm_cis.instance.id}"  
  domain_id = "${ibm_cis_domain.example.id}"
  name   = "terraform"
  content  = "192.168.0.11"
  type   = "A"
}
```

## Argument Reference

The following arguments are supported:

* `cis_id` - (Required) The ID of the CIS service instance
* `domain_id` - (Required) The ID of the domain to add the DNS record to.
* `name` - (Required) The name of the record, e.g. "www".
* `type` - (Required) The type of the record. A,AAAA,CNAME,NS,MX,TXT,LOC,SRV,SPF,CAA. 
* `content` - (Optional) The (string) value of the record, e.g. "192.168.127.127". Either this or `data` must be specified
* `data` - (Optional) Map of attributes that constitute the record value. Only for LOC, CCA and SRV record types. Either this or `content` must be specified
* `priority` - (Optional) The priority of the record
* `proxied` - (Optional) Whether the record gets CIS's origin protection; defaults to `false`.

## Attributes Reference

The following attributes are exported:

* `id` - The record ID
* `name` - The FQDN of the record
* `proxiable` - Shows whether this record can be proxied, must be true if setting `proxied=true`
* `created_on` - The RFC3339 timestamp of when the record was created
* `modified_on` - The RFC3339 timestamp of when the record was last modified
* `data` - Map of attributes that constitute the record value.

## Import

The `ibm_cis_dns_record` resource can be imported using the `id`. The ID is formed from the `Dns Record ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.  

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `bx cis` CLI commands.

* **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

* **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

* **Dns Record ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`. The id of an existing DNS record is not avaiable via the UI. It can be retrieved programatically via the CIS API or via the CLI using the CIS command to list the defined DNS recordss:  `bx cis dns-records <domain_id>` 


```
$ terraform import ibm_cis_dns_record.myorg <dns_record_id>:<domain-id>:<crn>

$ terraform import ibm_cis_dns_record.myorg  48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::