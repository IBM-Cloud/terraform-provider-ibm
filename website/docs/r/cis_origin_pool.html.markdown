---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_origin_pool"
description: |-
  Provides a IBM Cloud Internet Services Origin Pool resource.
---

# ibm_cis_origin_pool
Create, update, or delete an origin pool for your IBM Cloud Internet Services instance. This provides a pool of origins that can be used by a IBM CIS Global Load Balancer. For more information, about CIS origin pool, see [setting up origin pools](https://cloud.ibm.com/docs/cis?topic=cis-glb-features-pools).

## Example usage

```terraform
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

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `check_regions` - (Required, Set) A list of regions (specified by region code) from which to run health checks. If the list is empty, all regions are included, but you must use the enterprise plans. This is the default setting. Region codes can be found on our partner [Cloudflare's website](https://support.cloudflare.com/hc/en-us/articles/115000540888-Load-Balancing-Geographic-Regions).
- `description` - (Optional, String) A description for your origin pool. 
- `enabled`- (Bool) Required-If set to **true**, this pool is enabled and can receive incoming network traffic. Disabled pools do not receive network traffic and are excluded from health checks. Disabling a pool causes any load balancers that use the pool to failover to the next pool (if applicable).
- `name` - (Required, String) A short name (tag) for the pool. Only alphanumeric characters, hyphens, and underscores are allowed.
- `origins`-List of origins-Required-A list of origin servers within this pool. Traffic directed to this pool is balanced across all currently healthy origins, provided the pool itself is healthy.

  Nested scheme for `origins`:
  - `address` - (Required, String) The IPv4 or IPv6 address of the origin server. You can also provide a hostname for the origin that is publicly accessible. Make sure that the hostname resolves to the origin server, and is not proxied by IBM Cloud Internet Services.
  - `enabled` - (Optional, Bool) If set to **true**, the origin sever is enabled within the origin pool. If set to **false**, the origin server is not enabled. Disabled origin servers cannot receive incoming network traffic and are excluded from IBM Cloud Internet Services health checks.
  - `name` - (Required, String) The name of the origin server.
  - `weight` - (Optional, Float) The origin pool weight.
- `minimum_origins` - (Optional, Integer) The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins falls within this number, the pool will be marked unhealthy and we will failover to the next available pool. Default: 1.
- `monitor` - (Optional, String) The ID of the monitor to use for health checking origins within this pool.
- `notification_email` - (Optional, String) The Email address to send health status notifications to. This can be an individual mailbox or a mailing list.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the origin pool.
- `created_on` - (String) The RFC3339 timestamp of when the origin pool was created.
- `health` - (String) The status of the origin pool.
- `healthy` - (String) The status of the origin pool.
- `modified_on` - (String) The RFC3339 timestamp of when the origin pool was last modified.
- `origins`- (List) A list of origin servers that belong to the load balancer pool.
 
  Nested scheme for `origins`:
  - `disabled_at` - (Timestamp) The disabled date and time.
  - `failure_reason` - (String) The reason of failure.
  - `healthy`- (Bool) If set to **true**, the origin server is healthy. If set to **false**, the origin server is not healthy.

## Import
The origin pool can be imported by using the `id`. The ID is formed from the origin pool ID and the CRN (Cloud Resource Name). All values are concatenated with a `:` character.

The CRN can be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the console, or via the `ibmcloud cis` CLI.

- **CRN**: The CRN is a 120 digit character string of the format `crn:v1:bluemix:public:internet-svcs:global:a/1aa1111a1a1111aa1a111111111111aa:11aa111a-11a1-1a11-111a-111aaa11a1a1::` 
- **Origin pool ID**: The origin pool ID is a 32 digit character string in the format 1aaaa111111aa11111111111a1a11a1. The ID of a origin pool is not available via the console. It can be retrieved programmatically via the CIS API or via the command line by running `ibmcloud cis glb-pools`.

**Syntax**

```
$ terraform import ibm_cis_origin_pool.myorg <origin_pool_ID>:<crn>
```

**Example**

```
$ terraform import ibm_cis_origin_pool.myorg 1aaaa111111aa11111111111a1a11a1:crn:v1:bluemix:public:internet-svcs:global:a/1aa1111a1a1111aa1a111111111111aa:11aa111a-11a1-1a11-111a-111aaa11a1a1::
```
