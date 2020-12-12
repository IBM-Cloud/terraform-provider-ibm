---
layout: "ibm"
page_title: "IBM: ibm_cis_dns_records_import"
sidebar_current: "docs-ibm-resource-cis-dns-records-import"
description: |-
  Provides a IBM CIS DNS Records Import resource.
---

# ibm_cis_dns_records_import

Provides a IBM CIS DNS Records Import resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to import dns records from file of a domain of a CIS instance

## Example Usage

```hcl
# Import DNS Records of the domain

resource "ibm_cis_dns_records_import" "test" {
	cis_id    = data.ibm_cis.cis.id
	domain_id = data.ibm_cis_domain.cis_domain.domain_id
	file      = "dns_records.txt"
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance.
- `domain_id` - (Required,string) The ID of the domain to import dns records.
- `file` - (Required,ForceNew,string) The dns zone file which contains dns records detail.

## Attributes Reference

The following attributes are exported:

- `id` - The record ID. It is a combination of <`total_records_parsed`>,<`records_added`>,<`file`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
- `total_records_parsed` - The parsed records count from imported file.
- `records_added` - The added records count from imported file.

## Import

The `ibm_cis_dns_records_import` resource can be imported using the `id`. The ID is formed from the zone `file`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character with the prefix of `0:0:`.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **File** is a string of the form: `records.txt`

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

```
$ terraform import ibm_cis_dns_records_import.myorgs <total_records_parsed>:<records_added>:<file>:<domain-id>:<crn>

$ terraform import ibm_cis_dns_records_import.myorgs 0:0:records.txt:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
