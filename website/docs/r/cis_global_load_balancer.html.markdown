---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: cis_global_load_balancer"
description: |-
  Provides an IBM Cloud Internet Services Global Load Balancer resource.
---

# ibm_cis_global_load_balancer

Create, update, or delete an IBM CIS Global Load Balancer resource in a number of defined pools of origins, directs traffic to available origins and provides various options for geographically-aware load balancing. This resource is associated with an IBM Cloud Internet Services instance, a CIS Domain resource and CIS Origin pool resources. For more information, about Internet Services GLB, see [GLB concepts](https://cloud.ibm.com/docs/cis?topic=cis-global-load-balancer-glb-concepts).

## Example usage

```terraform
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
  steering_policy = "dynamic_latency"
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


## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain for which you want to add a global load balancer.
- `default_pools_ids` - (Required, String) A list of pool IDs that are ordered by their failover priority.
- `description` - (Optional, String) A description of the global load balancer.
- `enabled` - (Optional, Bool) If set to **true**, the load balancer is enabled and can receive network traffic. If set to **false**, the load balancer is not enabled.
- `fallback_pool_id` - (Required, String) The ID of the pool to use when all other pools are considered unhealthy.
- `name` - (Required, String) The DNS name to associate with the load balancer. This value can be a hostname, like `www`, or the fully qualified domain name, such as `www.example.com`. `example.com` is also accepted.
- `proxied` - (Optional, Bool) Indicates if the host name receives origin protection by IBM Cloud Internet Services. The default value is **false**.
- `pop_pools` - (Optional, Set) A set of mappings of the IBM Point-of-Presence (PoP) identifiers to the list of pool IDs (ordered by their failover priority) for the PoP (datacenter). This feature is only available to the enterprise customers.
  
  Nested scheme for `pop_pools`:
  - `pop` - (Required, String)  Enter a 3-letter code. Should not specify the multiple entries with the same PoP.
  - `pool_ids` - (Required, String)  A list of pool IDs in failover priority to use for the traffic reaching the provided PoP.
- `region_pools` - (Optional, Set) A set of containing mappings of region and country codes to the list of pool of IDs. IDs are ordered by their failover priority.

  Nested scheme for `region_pools`:
  - `region` - (Required, String) Enter a region code. Should not specify the multiple entries with the same region.
  - `pool_ids` - (Required, String) A list of pool IDs in failover priority for the provided region.
- `session_affinity` - (Optional, String) Associates all requests from an end-user with a single origin. IBM sets a cookie on the initial response to the client, so that the consequent requests with the cookie in the request use the same origin, as long as it is available.
- `steering_policy` - (Optional, String) Steering Policy which allows off,geo,random,dynamic_latency.
- `ttl` - (Optional, Integer) The time to live (TTL) in seconds for how long the load balancer must cache a resolved IP address for a DNS entry before the load balancer must look up the IP address again. If your global load balancer is proxied, this value is automatically set and cannot be changed. If your global load balancer is not in proxy, you can enter a value that is 120 or greater.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `created_on` - (Timestamp) The RFC3339 timestamp of when the load balancer was created.
- `glb_id` - (String) The unique identifier for the GLB.
- `id` - (String) The ID of the GLB.
- `modified_on` - (Timestamp) The RFC3339 timestamp of when the load balancer was last modified.
- `name` - (String) The fully qualified domain name that is associated with the Load Balancer.

## Import

The `ibm_cis_global_load_balancer` resource can be imported using the `id`. The ID is formed from the `Global Load Balancer ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated usinga `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` command.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Glb ID** is a 32 digit character string of the form: `57d96f0da6ed76251b475971b097205c`. The id of an existing GLB is not avaiable via the UI. It can be retrieved programatically via the CIS API or via the CLI using the CIS command to list the defined GLBs: `ibmcloud cis glbs <domain_id>`


**Syntax**

```
$ terraform import ibm_cis_global_load_balancer.myorg <glb_id>:<domain-id>:<crn>

```

**Example**

```
$ terraform import ibm_cis_domain.myorg  57d96f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
