---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_gateway"
description: |-
  Manages IBM Direct Link Gateway.
---

# ibm_dl_gateway

Create, update, or delete a Direct Link Gateway by using the Direct Link Gateway resource. For more information, see [about Direct Link](https://cloud.ibm.com/docs/dl?topic=dl-dl-about).


## Example usage to create Direct Link of dedicated type
In the following example, you can create Direct Link of dedicated type:

---
```terraform
data "ibm_dl_routers" "test_dl_routers" {
		offering_type = "dedicated"
		location_name = "dal10"
	}

resource ibm_dl_gateway test_dl_gateway {
  export_route_filters {
     	  action = "deny"
     	  prefix = "150.167.10.0/12"
        ge =19
        le = 29
  }
  import_route_filters {
	     	action = "permit"
	     	prefix = "140.167.10.0/12"
	      ge =17
	      le = 30
  }	
  default_export_route_filter = "permit"
  default_import_route_filter = "deny"
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
  vlan=3965
  #remove_vlan=false

} 
```
---
## Sample usage to create Direct Link of connect type
In the following example, you can create Direct Link of connect type:


---
```terraform
data "ibm_dl_ports" "test_ds_dl_ports" {
 
 }
resource "ibm_dl_gateway" "test_dl_connect" {
  as_prepends {
    length            = 3
    policy            = "import"
    specific_prefixes = ["10.10.9.0/24"]
  }

  as_prepends {
    length            = 3
    policy            = "export"
    specific_prefixes = ["10.10.9.0/24","10.10.10.0/24"]
  }
  export_route_filters {
     	  action = "deny"
     	  prefix = "150.167.10.0/12"
        ge =19
        le = 29
  }
  import_route_filters {
	     	action = "permit"
	     	prefix = "140.167.10.0/12"
	      ge =17
	      le = 30
  }
  default_export_route_filter = "permit"
  default_import_route_filter = "deny"		
  bgp_asn =  64999
  global = true
  metered = false
  name = "dl-connect-gw-1"
  speed_mbps = 1000
  type =  "connect"
  port =  data.ibm_dl_ports.test_ds_dl_ports.ports[0].port_id
}
```
---
## Argument reference
Review the argument reference that you can specify for your resource. 

- `as_prepends` - (Optional, List) List of AS Prepend configuration information
    
    Nested scheme for `as_prepend`:
  - `length` - (Required, Integer ) Number of times the ASN to appended to the AS Path.
  - `policy` - (Required, String) Route type this AS Prepend applies to. Possible values are `import` and `export`.
  - `prefix` - (Optional, Deprecated, String) Comma separated list of prefixes this AS Prepend applies to. Maximum of 10 prefixes. If not specified, this AS Prepend applies to all prefixes. prefix will be deprecated and support will be removed. Use specific_prefixes instead
  - `specific_prefixes` - (Optional, Array of Strings) Array of prefixes this AS Prepend applies to. If this property is absent, the AS Prepend applies to all prefixes.
- `export_route_filters` - (Optional, List) List of Export Route Filter configuration information.
  
  Nested scheme for `export_route_filter`:
  - `action` - (Required, String) Determines whether the  routes that match the prefix-set will be permit or deny
  - `prefix` - (Required, String) IP prefix representing an address and mask length of the prefix-set
  - `ge` - (Optional, Integer) The minimum matching length of the prefix-set
  - `le` - (Optional, Integer) The maximum matching length of the prefix-set
- `import_route_filters` - (Optional, List) List of Import Route Filter configuration information.
   Nested scheme for `import_route_filter`:
   - `action` - (Required, String) Determines whether the  routes that match the prefix-set will be permit or deny
   - `prefix` - (Required, String) IP prefix representing an address and mask length of the prefix-set
   - `ge` - (Optional, Integer) The minimum matching length of the prefix-set
   - `le` - (Optional, Integer) The maximum matching length of the prefix-set
- `authentication_key` - (Optional, String) BGP MD5 authentication key.
- `bfd_interval` - (String) Minimum interval in milliseconds at which the local routing device transmits hello packets and then expects to receive a reply from a neighbor with which it has established a BFD session.
- `bfd_multiplier` - (String) The number of hello packets not received by a neighbor that causes the originating interface to be declared down.
- `bgp_asn`- (Required, Integer) The BGP ASN of the gateway to be created. For example, `64999`.
- `bgp_base_cidr` - (Optional, String) (Deprecated) The BGP base CIDR of the gateway to be created. See `bgp_ibm_cidr` and `bgp_cer_cidr` for details on how to create a gateway by using  automatic or explicit IP assignment. Any `bgp_base_cidr` value set will be ignored.
- `bgp_cer_cidr` - (Optional, String) The BGP customer edge router CIDR. Specify a value within `bgp_base_cidr`.  For auto IP assignment, omit `bgp_cer_cidr` and `bgp_ibm_cidr`. IBM will automatically select values for `bgp_cer_cidr` and `bgp_ibm_cidr`.
- `bgp_ibm_cidr` - (Optional, String) The BGP IBM CIDR. For auto IP assignment, omit `bgp_cer_cidr` and `bgp_ibm_cidr`. IBM will automatically select values for `bgp_cer_cidr` and `bgp_ibm_cidr`.
- `carrier_name` - (Required, Forces new resource, String) The carrier name is required for `dedicated` type. Constraints are 1 ≤ length ≤ 128, Value must match regular expression ^[a-z][A-Z][0-9][ -_]$. For example, `myCarrierName`.
- `connection_mode` - (Optional, String) Type of network connection that you want to bind to your direct link. Allowed values are `direct` and `transit`.
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
- `default_export_route_filter` - (String) The default directional route filter action that applies to routes that do not match any directional route filters. 
- `default_import_route_filter` - (String) The default directional route filter action that applies to routes that do not match any directional route filters. 
- `vlan` - (Optional, Integer) The VLAN allocated for the gateway. You can set only for `type=dedicated` gateways. Allowed vlan range is 2-3967.
- `remove_vlan` - (Optional, Bool) The default value for this attribute is false. Set the value to true, if you want to remove the vlan value set earlier. You can remove vlan only for `type=dedicated` gateways. This attribute value conflicts with `vlan` attribute. You cannot set a `vlan` as well as `remove_vlan` at the same time.  

## Attribute reference
In addition to all argument references list, you can access the following attribute references after your resource is created.
  
- `as_prepends` - (List) List of AS Prepend configuration informationNested scheme for
  - `created_at`- (String) The date and time AS Prepend was created.
  - `id` - (String) The unique identifier for this AS Prepend.
  - `updated_at`- (String) The date and time AS Prepend was updated.
- `bfd_status` - (String) Gateway BFD status
- `bfd_status_updated_at` - (String) Date and time BFD status was updated at
- `bgp_status` - (String) The gateway BGP status.
- `bgp_status_updated_at` - (String) Date and time bgp status was updated.
- `completion_notice_reject_reason` - (String) The reason for completion notice rejection.
- `crn` - (String) The CRN of the gateway.
- `created_at` - (String) The date and time resource created.
- `id` - (String) The unique ID of the gateway.
- `location_display_name` - (String) The gateway location long name.
- `link_status` - (String) The gateway link status. You can include only on `type=dedicated` gateways. For example, `down`, `up`.
- `link_status_updated_at` - (String) Date and time link status was updated.
- `operational_status` - (String) The Gateway operational status. For gateways pending LOA approval, patch `operational_status` to the appropriate value to approve or reject its LOA. For example, `loa_accepted`.
- `provider_api_managed` - (String) Indicates whether gateway changes need to be made via a provider portal.
- `vlan` - (String) The VLAN allocated for the gateway. If the vlan is set by user, then this attribute value is shown only for gateway owners. Otherwise, this attribute value is shown as 0.

**Note**
The `Operational_status(Gateway operational status)` and `loa_reject_reason(LOA reject reason)` cannot be updated by using Terraform as the status and reason keeps changing with the different workflow actions.


## Import
The `ibm_dl_gateway` resource can be imported by using gateway ID.

**Syntax**

---
```
$ terraform import ibm_dl_gateway.example <gateway_ID>
```
---
**Example**

---
```
$ terraform import ibm_dl_gateway.example 5ffda12064634723b079acdb018ef308
```
---
