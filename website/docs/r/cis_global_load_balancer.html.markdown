---
layout: "ibm"
page_title: "IBM: cis_global_load_balancer"
sidebar_current: "docs-ibm-resource-cis-global-load-balancer"
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
  cis_id = "${ibm_cis.instance.id}"
  domain_id = "${ibm_cis_domain.example.id}"
  name = "example.com"
  fallback_pool_id = "${ibm_cis_origin_pool.example.id}"
  default_pool_ids = ["${ibm_cis_origin_pool.example.id}"]
  description = "example load balancer using geo-balancing"
  proxied = true
}

resource "ibm_cis_origin_pool" "example" {
  cis_id = "${ibm_cis.instance.id}"
  name = "example-lb-pool"
  origins {
    name = "example-1"
    address = "192.0.2.1"
    enabled = false
  }
}
```

## Argument Reference

The following arguments are supported:

* `cis_id` - (Required) The ID of the CIS service instance
* `domain_id` - (Required) The ID of the domain to add the load balancer to.
* `name` - (Required) The DNS name to associate with the load balancer.
* `fallback_pool_id` - (Required) The pool ID to use when all other pools are detected as unhealthy.
* `default_pool_ids` - (Required) A list of pool IDs ordered by their failover priority. Used whenever region/pop pools are not defined.
* `description` - (Optional) Free text description.
* `proxied` - (Optional) Whether the hostname gets ibm's origin protection. Defaults to `false`.
* `session_affinity` - (Optional) Associates all requests coming from an end-user with a single origin. ibm will set a cookie on the initial response to the client, such that consequent requests with the cookie in the request will go to the same origin, so long as it is available.

Region and pop pools are not currently implemented in this version of the provider. 

## Attributes Reference

The following attributes are exported:

* `id` - Unique identifier for the global load balancer.
* `created_on` - The RFC3339 timestamp of when the load balancer was created.
* `modified_on` - The RFC3339 timestamp of when the load balancer was last modified.

