---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : compute_dedicated_host"
description: |-
  Manages IBM Cloud dedicated host.
---

# ibm_compute_dedicated_host
Create, update, and delete a dedicated host resource. For more information, about compute dedicated host, see [dedicated hosts and dedicated instances](https://cloud.ibm.com/docs/virtual-servers?topic=virtual-servers-dedicated-hosts-and-dedicated-instances).

**Note**
For more information, see [SoftLayer API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Virtual_DedicatedHost).

## Example usage
In the following example, you can create a dedicated host:

```terraform
resource "ibm_compute_dedicated_host" "dedicatedhost" {
  hostname        = "host"
  domain          = "example.com"
  router_hostname = "bcr01a.dal09"
  datacenter      = "dal09"
}

```

## Argument reference
Review the argument references that you can specify for your resource. 

- `datacenter` - (Required, Forces new resource, String) The data center in which the dedicated host resides.
- `domain` - (Required, Forces new resource, String) The domain of dedicated host.
- `flavor`- (Optional, Forces new resource, String) The flavor of dedicated host. Default value `56_CORES_X_242_RAM_X_1_4_TB`. [Log in to the IBM-Cloud Infrastructure API to see available flavor types](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/813/getItems.json). Use your API as the password to log in. Log in and find the key called `keyName`.
- `hostname` - (Required, String) The host name of dedicated host.
- `hourly_billing`-  (Optional, Forces new resource,Bool) The billing type for the host. When set to **true**, the dedicated host is billed on hourly usage. Otherwise, the dedicated host is billed monthly. The default value is **true**.
- `router_hostname` - (Required, Forces new resource, String) The hostname of the primary router associated with the dedicated host.
- `wait_time_minutes` - (Optional, Integer)The duration, expressed in minutes to wait for the dedicated host to become available before creation. The default value is `90`.
- `tags`- (Optional, Array of string) Tags associated with the dedicated host.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `cpu_count`- (String) The capacity that the dedicated host's CPU allocation is restricted to.
- `disk_capacity`- (String) The capacity that the dedicated host's disk allocation is restricted to.
- `id`- (String) The unique identifier of the dedicated host.
- `memory_capacity`- (String) The capacity that the dedicated host's memory allocation is restricted to.


## Import
The `ibm_compute_dedicated_host` resource can be imported by using the ID. 

**Example**

```
$ terraform import ibm_compute_dedicated_host.dedicatedhost 238756
```
