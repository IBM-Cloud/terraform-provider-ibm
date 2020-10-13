---
layout: "ibm"
page_title: "IBM: ibm_cis_origin_pool"
sidebar_current: "docs-ibm-cis-origin-pool"
description: |-
  Provides a IBM Cloud Internet Services Origin Pool resource.
---

# ibm_cis_origin_pool

Provides a IBM Cloud Internet Services origin pool resource. This provides a pool of origins that can be used by a IBM CIS Global Load Balancer. This resource is associated with an IBM Cloud Internet Services instance and optionally a CIS Healthcheck monitor resource.

## Example Usage

```hcl
resource "ibm_cis_origin_pool" "example" {
  cis_id = ibm_cis.instance.id
  name   = "example-pool"
  origins {
    name    = "example-1"
    address = "192.0.2.1"
    enabled = false
  }
  origins {
    name    = "example-2"
    address = "192.0.2.2"
    enabled = false
  }
  description        = "example load balancer pool"
  enabled            = false
  minimum_origins    = 1
  notification_email = "someone@example.com"
  check_regions      = ["WEU"]
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance
- `name` - (Required,string) A short name (tag) for the pool. Only alphanumeric characters, hyphens, and underscores are allowed.
- `origins` - (Required,set) The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy origins, provided the pool itself is healthy. It's a complex value. See description below.
  - `name` - (Required,string) A human-identifiable name for the origin.
  - `address` - (Required,string) The IP address (IPv4 or IPv6) of the origin, or the publicly addressable hostname. Hostnames entered here should resolve directly to the origin, and not be a hostname proxied by CIS.
  - `enabled` - (Optional,bool) Whether to enable (the default) this origin within the Pool. Disabled origins will not receive traffic and are excluded from health checks. The origin will only be disabled for the current pool.
  - `weight` - (Optional, float) The origin pool weight.
- `check_regions` - (Required,set) A list of regions (specified by region code) from which to run health checks. Empty means every region (the default), but requires an Enterprise plan. Region codes can be found on our partner Cloudflare's website [here](https://support.cloudflare.com/hc/en-us/articles/115000540888-Load-Balancing-Geographic-Regions).
- `description` - (Optional,string) Free text description.
- `enabled` - (Required,bool) Whether to enable (the default) this pool. Disabled pools will not receive traffic and are excluded from health checks. Disabling a pool will cause any load balancers using it to failover to the next pool (if any).
- `minimum_origins` - (Optional,int) The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins falls below this number, the pool will be marked unhealthy and we will failover to the next available pool. Default: 1.
- `monitor` - (Optional,string) The ID of the Monitor to use for health checking origins within this pool.
- `notification_email` - (Optional,string) The email address to send health status notifications to. This can be an individual mailbox or a mailing list.

## Attributes Reference

The following attributes are exported:

- `id` - ID for this load balancer pool.
- `created_on` - The RFC3339 timestamp of when the load balancer was created.
- `modified_on` - The RFC3339 timestamp of when the load balancer was last modified.
- `health`- The status of the origin pool.
- `healthy` - The status of the origin pool.
- `origins`- The list of origins within this pool.
  - `healthy` - Indicates if pool is healthy or not
  - `disabled_at` - The disabled date and time.
  - `failure_reason` - The reason of failure.

## Import

The `ibm_cis_origin_pool` resource can be imported using the `id`. The ID is formed from the `Origin Pool Id` and the `CRN` (Cloud Resource Name) concatentated usinga `:` character.

The CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading.

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Origin Pool ID** is a 32 digit character string of the form: `000f57b5c42bcff3c02d155c2d58aa97`. The id of an existing pool is not available via the UI. It can be retrieved programmatically via the CIS API or via the CLI using the CIS command to list the defined GLBs: `ibmcloud cis glb-pools`

```
$ terraform import ibm_cis_origin_pool.myorg <origin_pool_id>:<crn>

$ terraform import origin_pool.myorg 000f57b5c42bcff3c02d155c2d58aa97:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
