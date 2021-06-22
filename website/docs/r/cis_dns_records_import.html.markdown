---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_dns_records_import"
description: |-
  Provides a IBM CIS DNS records import resource.
---

# ibm_cis_dns_records_import

Provides an IBM Cloud Internet Services DNS records import resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS domain resource. It allows to import DNS records from file of a domain of a CIS instance. For more information, about CIS DNS records, refer to [managing DNS records](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-managing-dns-records).

## Example usage

```terraform
# Import DNS Records of the domain

resource "ibm_cis_dns_records_import" "test" {
	cis_id    = data.ibm_cis.cis.id
	domain_id = data.ibm_cis_domain.cis_domain.domain_id
	file      = "dns_records.txt"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain to import the DNS records.
- `file` - (Required, Forces new resource, String) The DNS zone file that contains the details of the DNS records.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The record ID. It is a combination of `<total_records_parsed>:<records_added>:<file>:<domain_id>:<cis_id>` attributes concatenated with `:`.
- `records_added` - (String) The added records count from imported file.
- `total_records_parsed`- (Integer) The parsed records count from imported file.


## Import
The `ibm_cis_dns_records_import` resource can be imported by using the ID. The ID is formed from the zone file, the domain ID of the domain and the CRN (Cloud Resource Name)  Concatenated  using a `:` character with the prefix of `0:0:`.

The domain ID and CRN is located on the **Overview** page of the internet services instance under the domain heading of the console, or via by using the `ibmcloud cis` command line commands.

- **File** is a string of the form: `records.txt`

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

**Syntax**

```
$ terraform import ibm_cis_dns_records_import.myorgs <total_records_parsed>:<records_added>:<file>:<domain-id>:<crn>
```
**Example**

```
$ terraform import ibm_cis_dns_records_import.myorgs 0:0:records.txt:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
