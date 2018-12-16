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
* `data` - 

