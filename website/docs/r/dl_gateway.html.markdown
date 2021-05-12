---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_gateway"
description: |-
  Manages IBM Direct Link Gateway.
---

# `ibm_dl_gateway`

Create, update, or delete a Direct Link Gateway by using the Direct Link Gateway resource. For more information, see [about Direct Link](https://cloud.ibm.com/docs/dl?topic=dl-dl-about).


## Example usage to create Direct Link of dedicated type
In the following example, you can create Direct Link of dedicated type:

```
data "ibm_dl_routers" "test_dl_routers" {
		offering_type = "dedicated"
		location_name = "dal10"
	}

resource ibm_dl_gateway test_dl_gateway {
  bgp_asn =  64999
  global = true 
  metered = false
  name = "Gateway1"
  resource_group = "bf823d4f45b64ceaa4671bee0479346e"
  speed_mbps = 1000 
  type =  "dedicated" 
  cross_connect_router = data.ibm_dl_routers.test_dl_routers.cross_connect_routers[0].router_name
  location_name = data.ibm_dl_routers.test_dl_routers.location_name
  customer_name = "Customer1" 
  carrier_name = "Carrier1"

} 
```

## Sample usage to create Direct Link of connect type
In the following example, you can create Direct Link of connect type:


```
data "ibm_dl_ports" "test_ds_dl_ports" {
 
 }
resource "ibm_dl_gateway" "test_dl_connect" {
  bgp_asn =  64999
  global = true
  metered = false
  name = "dl-connect-gw-1"
  speed_mbps = 1000
  type =  "connect"
  port =  data.ibm_dl_ports.test_ds_dl_ports.ports[0].port_id
}
```

## Argument reference
Review the input parameters that you can specify for your resource. 

- `bgp_asn`- (Required, Forces new resource, Integer) The BGP ASN of the gateway to be created. For example, `64999`.
- `bgp_base_cidr` - (Optional, String) (Deprecated) The BGP base CIDR of the gateway to be created. See `bgp_ibm_cidr` and `bgp_cer_cidr` for details on how to create a gateway by using  automatic or explicit IP assignment. Any `bgp_base_cidr` value set will be ignored.
- `bgp_cer_cidr` - (Optional, Forces new resource, String) The BGP customer edge router CIDR. Specify a value within `bgp_base_cidr`.  For auto IP assignment, omit `bgp_cer_cidr` and `bgp_ibm_cidr`. IBM will automatically select values for `bgp_cer_cidr` and `bgp_ibm_cidr`.
- `bgp_ibm_cidr` - (Optional, Forces new resource, String) The BGP IBM CIDR. For auto IP assignment, omit `bgp_cer_cidr` and `bgp_ibm_cidr`. IBM will automatically select values for `bgp_cer_cidr` and `bgp_ibm_cidr`.
- `carrier_name` - (Required, Forces new resource, String) The carrier name is required for `dedicated` type. Constraints are 1 ≤ length ≤ 128, Value must match regular expression ^[a-z][A-Z][0-9][ -_]$. For example, `myCarrierName`.
- `cross_connect_router` - (Required, Forces new resource, String) The cross connect router required for `dedicated` type. For example, `xcr01.dal03`.
- `customer_name` - (Required, Forces new resource, String) The customer name is required for `dedicated` type. Constraints are 1 ≤ length ≤ 128, Value must match regular expression ^[a-z][A-Z][0-9][ -_]$. For example, `newCustomerName`.
- `global`- (Bool) Required-Gateway with global routing as **true** can connect networks outside your associated region.
- `location_name` - (Required, Forces new resource, String) The gateway location is required for `dedicated` type. For example, `dal03`.
- `name` - (Required, String) The unique user-defined name for the gateway. For example, `myGateway`.No.
- `metered`- (Required, Bool) Metered billing option. If set **true** gateway usage is billed per GB. Otherwise, flat rate is charged for the gateway.
- `port` - (Required, Forces new resource, String) The gateway port for type is connect gateways. This parameter is required for Direct Link connect type.
- `resource_group` - (Optional, Forces new resource, String) The resource group. If unspecified, the account's default resource group is used.
- `speed_mbps`- (Required, Integer) The gateway speed in MBPS. For example, `10.254.30.78/30`.
- `type` - (Required, Forces new resource, String) The gateway type, allowed values are `dedicated` and `connect`.

## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `bgp_asn` - (String) The IBM BGP ASN.
- `bgp_status` - (String) The gateway BGP status.
- `completion_notice_reject_reason` - (String) The reason for completion notice rejection.
- `crn` - (String) The CRN of the gateway.
- `created_at` - (String) The date and time resource created.
- `id` - (String) The unique ID of the gateway.
- `location_display_name` - (String) The gateway location long name.
- `link_status` - (String) The gateway link status. You can include only on `type=dedicated` gateways. For example, `down`, `up`.
- `name` - (String) The unique user-defined name for the gateway.
- `operational_status` - (String) The gateway operational status. For gateways pending LOA approval, patch operational_status to the appropriate value to approve or reject its LOA. For example, `loa_accepted`.
- `port` - (String) The gateway port for `type=connect` gateways.
- `provider_api_managed` - (String) Indicates whether gateway changes need to be made via a provider portal.
- `resource_group` - (String) The resource group reference.
- `vlan` - (String) The VLAN allocated for the gateway. You can set only for `type=connect` gateways created directly through the IBM portal.

**Note**
The `Operational_status(Gateway operational status)` and `loa_reject_reason(LOA reject reason)` cannot be updated by using Terraform as the status and reason keeps changing with the different workflow actions.


## Import
The `ibm_dl_gateway` can be imported by using gateway ID. 

**Example**

```
terraform import ibm_dl_gateway.example 5ffda12064634723b079acdb018ef308
```

