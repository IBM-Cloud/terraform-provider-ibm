---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : "
description: |-
  Manages IBM Cloud Infrastructure Private Domain Name Service GLB Pools.
---

# ibm_dns_glb_pools

Retrieve the details of an existing IBM Cloud infrastructure private DNS Global Load Balancers (glb) pools as a read-only data source. For more information, see [viewing Global Load Balancer events](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-health-check-events).


## Example usage

```terraform
data "ibm_dns_glb_pools" "ds_pdns_glb_pools" {
  instance_id = "resource_instance_guid"
}
```


## Argument reference
Review the argument reference that you can specify for your data source. 

- `instance_id` - (Required, String) The resource GUID of the private DNS service on which zones are created.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `dns_glb_pools` - (List) List of all private DNS Load balancer pools in the IBM Cloud infrastructure.
	
   Nested scheme for `dns_glb_pools`:
   - `description` - (String) The descriptive text of the DNS Load balancer pool.
   - `enable` - (String)  Whether the Load Balancer pool is enabled.
   - `health` - (String) The status of DNS GLB pool's health. Possible values are `DOWN`, `UP`, `DEGRADED`.
   - `healthy_origins_threshold` - (String) The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins falls less than this number, the pool will be marked unhealthy and will failover to the next available pool.
   - `healthcheck_region` - (String) Health check region of VSIs. Allowable values are `us-south`,`us-east`, `eu-gb`, `eu-du`, `au-syd`, `jp-tok`.
   - `healthcheck_subnets` - (String) Health check subnet CRN of VSIs.
   - `origins` (List) The list of origins within the pool. Traffic directed to the pool is balanced across all currently healthy origins, provided the pool itself is healthy.

     Nested scheme for `origins`:
	 - `address` - (String) The address of the origin server. It can be a hostname or an IP address.
	 - `description` - (String) The description of the origin server.
	 - `enabled` - (String) Whether the origin server is enabled.
	 - `health` - (String) Whether the health is **true** or **false**.
	 - `health_failure_reason` - (String) The reason for the health check failure.
	 - `name` - (String) The name of the origin server.
	- `monitor` - (String) The ID of the Load Balancer monitor to be associated to this pool.
	- `name` - (String) The name of the DNS Load balancer pool.
	- `notification_channel` - (String) The webhook URL as a notification channel.
	- `pool_id` - (String) The pool ID.
	- `created_on`- (Timestamp) The time (created On) of the DNS glb pool.
	- `modified_on`- (Timestamp) The time (modified On) of the DNS glb pool.
	
