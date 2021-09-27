---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_global_load_balancers"
description: |-
  Manages IBM Cloud Internet Services Global Load Balancers resource.
---

# ibm_cis_global_load_balancers
Retrieve information `24 X 7` availability and performance of your application by using the IBM Cloud Internet Services global Load Balancers. For more information, refer to [CIS global loadbalancer](https://cloud.ibm.com/docs/cis?topic=cis-configure-glb).

## Example usage
The following example retrieves information about an IBM Cloud Internet Services global Load Balancer resource.

```terraform

data "ibm_cis_global_load_balancers" "test" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The resource CRN ID of the CIS on which zones were created.
- `domain_id` - (Required, String) The ID of the domain to retrieve the Load Balancers from.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `description` - (String) Free text description.
- `default_pool_ids` - (String) A list of pool IDs ordered by their failover priority. Used whenever region or pop pools are not defined.
- `enabled` - (String) Indicates if the Load Balancer is enabled or not. Region and pop pools are not currently implemented in this version of the provider.
- `fallback_pool_id` - (String) The pool ID to use when all other pools are detected as unhealthy.
- `glb_id` - (String) The Load Balancer ID.
- `id` - (String) The Load Balancer ID, domain ID and CRN. For example, `id:domain-id:crn`.
- `name` - (String) The DNS name to associate with the Load Balancer. This can be a hostname, for example, `www` or the fully qualified name `www.example.com`, or `example.com`.
- `proxied` - (String) Whether the hostname gets IBM's origin protection. Defaults to **false**.
- `pop_pools` - (String) A set containing mappings of IBM Point-of-Presence (PoP) identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter). This feature is only available to enterprise customers.

  Nested scheme for `pop_pools`:
	- `pop` - (String) A 3-letter code for the Point-of-Presence. Multiple entries should not be specified with the same PoP.
	- `pool_ids` - (String) A list of pool IDs in failover priority to use for traffic reaching the given PoP.
- `region_pools` - (String) A set containing mappings of region or country codes to a list of pool IDs (ordered by their failover priority) for the given region.

  Nested scheme for `region_pools`:
	- `region` - (String) A region code. Multiple entries is not allowed with the same region.
	- `pool_ids` - (String) A list of pool IDs in failover priority to use in the given region.
- `session_affinity` - (String) Associates all requests coming from an end-user with a single origin. IBM will set a cookie on the initial response to the client, such that consequent requests with the cookie in the request will go to the same origin, as long as it is available.
- `steering_policy` - (String) Steering Policy which allows off,geo,random,dynamic_latency.
- `ttl` - (String) Time to live (TTL) of the DNS entry for the IP address returned by this Load Balancer.P.
