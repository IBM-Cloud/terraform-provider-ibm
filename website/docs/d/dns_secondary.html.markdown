---
layout: "ibm"
page_title: "IBM: ibm_dns_secondary"
sidebar_current: "docs-ibm-datasource-dns-secondary"
description: |-
  Get information about an IBM DNS secondary resource.
---

# ibm\_dns_secondary

Import the name of an existing dns secondary zone as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_dns_secondary" "secondary_id" {
    name = "test-secondary.com"
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required, string) The name of the secondary zone, as it was defined in IBM Cloud Infrastructure (SoftLayer).

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the secondary.
* `transfer_frequency` - Signifies how often a secondary DNS zone transferred in minutes.
* `master_ip_address` - The IP address of the master name server where a secondary DNS zone is transferred from.
* `status_id` - The current status of a secondary DNS record.
* `status_text` - The textual representation of a secondary DNS zone's status.
