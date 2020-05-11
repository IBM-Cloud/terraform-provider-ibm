---
layout: "ibm"
page_title: "IBM : "
sidebar_current: "docs-ibm-datasources-dns-zones"
description: |-
  Manages IBM Cloud Infrastructure Private Domain Name Service Zones.
---

# ibm\_dns_zones

Import the details of an existing IBM Cloud Infrastructure private domain name service zones as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_dns_zones" "ds_pdnszones" {
  instance_id = "resource_instance_guid"
}

```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The resource instance id of the private DNS on which zones were created.



## Attribute Reference

The following attributes are exported:

* `dns_zones` - List of all private domain name service zones in the IBM Cloud Infrastructure.
  * `zone_id` - The unique identifier of the private DNS zone.
  * `instance_id` - The resource instance id of the private DNS on which zones were created.
  * `description` - The text describing the purpose of the DNS zone.
  * `name` - The name of the DNS zone.
  * `label` - The label of the DNS zone.
  * `created_on` - The created time of the DNS zone.
  * `modified_on` - The modified time of the DNS zone.
  * `state` - The state of the DNS zone.
