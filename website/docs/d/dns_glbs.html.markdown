---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : ibm_dns_glbs"
description: |-
  Manages IBM Cloud Infrastructure private domain name service Global Load Balancers.
---

# ibm_dns_glbs

Retrieve the details of an existing IBM Cloud infrastructure private DNS Global Load Balancers as a read-only data source. For more information, see [working with global Load Balancers](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-global-load-balancers).


## Example usage

```terraform
data "ibm_dns_glbs" "test1" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
}
```

## Argument reference
Review the argument reference that you can specify for your data source. 
 
- `instance_id` - (Required, String) The GUID of the private DNS service instance.
- `zone_id` - (Required, String) The ID of the private DNS zone.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.. 


- `dns_glbs` (List) List of all private DNS Load balancers in the IBM Cloud infrastructure.

   Nested scheme for `dns_glbs`:
   - `az_pools` (List) Map availability zones to the pool ID's.
	 
	 Nested scheme for `az_pools`:
	 - `availability_zone` - (String) The availability zone.
	 - `pools` - (String) List of Load Balancer pools.
   - `created_on`- (Timestamp) The date and time when the Load Balancer was created.
   - `description` - (String) The descriptive text of the DNS Load balancers.
   - `default_pools` - (String) TA list of pool IDs ordered by their failover priority.
   - `fallback_pool` - (String) The pool ID to use when all other pools are detected as unhealthy.
   - `glb_id` - (String) The Load Balancer ID.
   - `health` - (String) Healthy state of the Load Balancer. Possible values are `DOWN`, `UP`, or `DEGRADED`.
   - `modified_on`- (Timestamp) The date and time when the Load Balancer was modified.
   - `name` - (String) The name of the DNS Load balancers.
   - `ttl` - (String) The time to live in second.
