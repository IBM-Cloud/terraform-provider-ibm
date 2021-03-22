---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: cis_global_load_balancer"
description: |-
  Provides an IBM Cloud Internet Services Global Load Balancer resource.
---

# ibm_cis_global_load_balancer

Provides a IBM CIS Global Load Balancer resource. This sits in front of a number of defined pools of origins, directs traffic to available origins and provides various options for geographically-aware load balancing. This resource is associated with an IBM Cloud Internet Services instance, a CIS Domain resource and CIS Origin pool resources.

## Example Usage

```hcl
# Define a global load balancer which directs traffic to defined origin pools
# In normal usage different pools would be set for data centers/availability zones and/or for different regions
# Within each availability zone or region we can define multiple pools in failover order

resource "ibm_cis_global_load_balancer" "example" {
  cis_id           = ibm_cis.instance.id
  domain_id        = ibm_cis_domain.example.id
  name             = "www.example.com"
  fallback_pool_id = ibm_cis_origin_pool.example.id
  default_pool_ids = [ibm_cis_origin_pool.example.id]
  description      = "example load balancer using geo-balancing"
  proxied          = true
  region_pools{
			region="WEU"
			pool_ids = [ibm_cis_origin_pool.example.id]
		}
	pop_pools{
			pop="LAX"
			pool_ids = [ibm_cis_origin_pool.example.id]
		}
}

resource "ibm_cis_origin_pool" "example" {
  cis_id = ibm_cis.instance.id
  name   = "example-lb-pool"
  origins {
    name    = "example-1"
    address = "192.0.2.1"
    enabled = false
  }
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance
- `domain_id` - (Required,string) The ID of the domain to add the load balancer to.
- `name` - (Required,string) The DNS name to associate with the load balancer. This can be a hostname, e.g. "www" or the fully qualified name "www.example.com". "example.com" is also accepted.
- `fallback_pool_id` - (Required,string) The pool ID to use when all other pools are detected as unhealthy.
- `default_pool_ids` - (Required,map) A list of pool IDs ordered by their failover priority. Used whenever region/pop pools are not defined.
- `description` - (Optional,string) Free text description.
- `proxied` - (Optional,bool) Whether the hostname gets ibm's origin protection. Defaults to `false`.
- `session_affinity` - (Optional,string) Associates all requests coming from an end-user with a single origin. ibm will set a cookie on the initial response to the client, such that consequent requests with the cookie in the request will go to the same origin, so long as it is available.
- `ttl` - (Optional,int) Time to live (TTL) of the DNS entry for the IP address returned by this load balancer.
- `enabled` - (Optional,bool) Indicates if the load balancer is enabled or not.
  Region and pop pools are not currently implemented in this version of the provider.
- `region_pools` - (Optional,set) A set containing mappings of region/country codes to a list of pool IDs (ordered by their failover priority) for the given region.
  - `region` - (Required,string) A region code. Multiple entries should not be specified with the same region.
  - `pool_ids` - (Required,string) A list of pool IDs in failover priority to use in the given region.
- `pop_pools` - (Optional,set) A set containing mappings of IBM Point-of-Presence (PoP) identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter). This feature is only available to enterprise customers.
  - `pop` - (Required,string) A 3-letter code for the Point-of-Presence.Multiple entries should not be specified with the same PoP.
  - `pool_ids` - (Required,string) A list of pool IDs in failover priority to use for traffic reaching the given PoP.

## Attributes Reference

The following attributes are exported:

- `id` - Unique identifier for the global load balancer. Ex. <unique-id>:<domain-id>:<crn>
- `glb_id` - Unique identifier for the global load balancer.
- `created_on` - The RFC3339 timestamp of when the load balancer was created.
- `modified_on` - The RFC3339 timestamp of when the load balancer was last modified.

## Import

The `ibm_cis_global_load_balancer` resource can be imported using the `id`. The ID is formed from the `Global Load Balancer ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated usinga `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Glb ID** is a 32 digit character string of the form: `57d96f0da6ed76251b475971b097205c`. The id of an existing GLB is not avaiable via the UI. It can be retrieved programatically via the CIS API or via the CLI using the CIS command to list the defined GLBs: `ibmcloud cis glbs <domain_id>`

```
$ terraform import ibm_cis_global_load_balancer.myorg <glb_id>:<domain-id>:<crn>

$ terraform import ibm_cis_domain.myorg  57d96f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
