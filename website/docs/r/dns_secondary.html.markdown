---
layout: "ibm"
page_title: "IBM: dns_secondary"
sidebar_current: "docs-ibm-resource-dns-secondary"
description: |-
  Manages IBM DNS Secondary Zone.
---

# ibm\_dns_secondary

The `ibm_dns_secondary` resource represents a single secondary DNS zone managed on SoftLayer. Each record created within the secondary DNS service defines which zone is transferred, what server it is transferred from, and the frequency that zone transfers occur at. Zone transfers are performed automatically based on the transfer frequency set on the secondary DNS record.

## Example Usage

```hcl
resource "ibm_dns_secondary" "dns-secondary-test" {
    zone_name = "dns-secondary-test.com"
    master_ip_address = "127.0.0.10"
    transfer_frequency = 10
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required, string) The name of the zone that is transferred.
* `transfer_frequency` - (Required, int) Signifies how often a secondary DNS zone should be transferred in minutes.
* `master_ip_address` - (Required, string)  The IP address of the master name server where a secondary DNS zone is transferred from.
* `tags` - (Optional, array of strings) Tags associated with the DNS secondary instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - A secondary zone's internal identifier.
* `status_id` - The current status of a secondary DNS record.
* `status_text` - The textual representation of a secondary DNS zone's status.