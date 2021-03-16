---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_global_load_balancers"
description: |-
  Manages IBM Cloud Internet Services Global Load Balancers resource.
---

# ibm_cis_global_load_balancers

Import the details of an existing IBM Cloud Internet Service global load balancers as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl

data "ibm_cis_global_load_balancers" "test" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
}

```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required, string) The resource crn id of the CIS on which zones were created.
- `domain_id` - (Required, string) The ID of the domain to retrive the load balancers from.

## Attribute Reference

The following attributes are exported:

- `id` - Load balancer ID, domain id and CRN. Ex. id:domain-id:crn
- `glb_id` - Load balancer ID.
- `name` - The DNS name to associate with the load balancer. This can be a hostname, e.g. "www" or the fully qualified name "www.example.com". "example.com" is also accepted.
- `fallback_pool_id` - The pool ID to use when all other pools are detected as unhealthy.
- `default_pool_ids` - A list of pool IDs ordered by their failover priority. Used whenever region/pop pools are not defined.
- `description` - Free text description.
- `proxied` - Whether the hostname gets ibm's origin protection. Defaults to `false`.
- `session_affinity` - Associates all requests coming from an end-user with a single origin. ibm will set a cookie on the initial response to the client, such that consequent requests with the cookie in the request will go to the same origin, so long as it is available.
- `ttl` - Time to live (TTL) of the DNS entry for the IP address returned by this load balancer.
- `enabled` - Indicates if the load balancer is enabled or not.
  Region and pop pools are not currently implemented in this version of the provider.
- `region_pools` - A set containing mappings of region/country codes to a list of pool IDs (ordered by their failover priority) for the given region.
  - `region` - A region code. Multiple entries should not be specified with the same region.
  - `pool_ids` - A list of pool IDs in failover priority to use in the given region.
- `pop_pools` - A set containing mappings of IBM Point-of-Presence (PoP) identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter). This feature is only available to enterprise customers.
  - `pop` - A 3-letter code for the Point-of-Presence.Multiple entries should not be specified with the same PoP.
  - `pool_ids` - A list of pool IDs in failover priority to use for traffic reaching the given PoP.
