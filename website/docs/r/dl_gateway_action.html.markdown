---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_gateway_action"
description: |-
  Manages IBM Direct Link Gateway Action.
---

# ibm_dl_gateway_action

This resource used for provider created Direct Link Connect gateways to approve or reject specific changes initiated from a provider portal For more information, see [about Direct Link](https://cloud.ibm.com/docs/dl?topic=dl-dl-about).


## Sample usage to approve the provider created Direct Link of connect type
In the following example, you can approve the provider created Direct Link of connect type:


---
```terraform
variable "gname" {
  default     = "my-testgw01924"
  description = "gateway Name"
}

data "ibm_dl_provider_ports" "test_ds_dl_ports" {
}
resource ibm_dl_provider_gateway "test_dl_gateway" {
	provider = ibm.packet_fabric
	bgp_asn =  64999
	name = var.gname
	speed_mbps = 1000 
	port = data.ibm_dl_provider_ports.test_ds_dl_ports.ports[0].port_id
	customer_account_id = "Some customer account"			  		                  
}
resource "time_sleep" "wait_dl_connection" {
  create_duration = "1m"
}

data "ibm_dl_gateway" "test_ibm_dl_gateway" {
	provider = ibm.ibm_customer
	name = var.gname
	depends_on = [time_sleep.wait_dl_connection]	
}
resource "ibm_dl_gateway_action" "test_dl_gateway_action" {
	provider = ibm.ibm_customer
	gateway = data.ibm_dl_gateway.test_ibm_dl_gateway.id
	action = "create_gateway_approve"	
	global = true 	
	metered = true
	default_export_route_filter = "permit"
	default_import_route_filter = "deny"
    export_route_filters {
	     	action = "deny"
	     	prefix = "10.10.10.0/24"
	        ge =25
	        e = 29
    }
    import_route_filters {
  	     	action = "permit"
  	     	prefix = "10.10.10.0/24"
  	        ge =25
  	        le = 29
	 }	
	 resource_group = data.ibm_dl_gateway.test_ibm_dl_gateway.resource_group	             
}
```
---
---
```go
variable "gname" {
	default     = "my-testgw01924"
	description = "gateway Name"
}
  
data "ibm_dl_provider_ports" "test_ds_dl_ports" {
}

resource "ibm_dl_provider_gateway" "test_dl_gateway" {
	provider = ibm.packet_fabric
	bgp_asn =  64999
	name = var.gname
	speed_mbps = 1000 
	port = data.ibm_dl_provider_ports.test_ds_dl_ports.ports[0].port_id
	customer_account_id = "Some customer account"			  		                  
}
resource "time_sleep" "wait_dl_connection" {
	create_duration = "1m"
}
data "ibm_dl_gateway" "test_ibm_dl_gateway" {
	provider = ibm.ibm_customer
	name = var.gname
	depends_on = [time_sleep.wait_dl_connection]		
}
resource ibm_dl_gateway_action "test_dl_gateway_action" {
	provider = ibm.ibm_customer
	gateway = data.ibm_dl_gateway.test_ibm_dl_gateway.id
	action = "update_attributes_approve"		             
}
```
---
---
```go
variable "gname" {
	default     = "my-testgw01924"
	description = "gateway Name"
}
data "ibm_dl_provider_ports" "test_ds_dl_ports" {
}
resource ibm_dl_provider_gateway "test_dl_gateway" {
	bgp_asn =  64999
	name = var.gname
	speed_mbps = 1000 
	port = data.ibm_dl_provider_ports.test_ds_dl_ports.ports[0].port_id
	customer_account_id = "Some customer account"		  		                  
}
data "ibm_dl_gateway" "test_ibm_dl_gateway" {
	name = var.gname
	depends_on = [time_sleep.wait_dl_connection]
}
resource ibm_dl_gateway_action "test_dl_gateway_action" {
	gateway = data.ibm_dl_gateway.test_ibm_dl_gateway.id
	action = "delete_gateway_approve"		             
}
```
---
## Argument reference
Review the argument reference that you can specify for your resource. 

- `action` - (Required, String) Approve/reject a pending change request.
- `as_prepends` - (Optional, List) List of AS Prepend configuration information.Applicable only  for create_gateway_approve requests.
    
   Nested scheme for `as_prepend`:
  - `length` - (Required, Integer ) Number of times the ASN to appended to the AS Path.
  - `policy` - (Required, String) Route type this AS Prepend applies to. Possible values are `import` and `export`.
  - `prefix` - (Optional, Deprecated, String) Comma separated list of prefixes this AS Prepend applies to. Maximum of 10 prefixes. If not specified, this AS Prepend applies to all prefixes. prefix will be deprecated and support will be removed. Use specific_prefixes instead
  - `specific_prefixes` - (Optional, Array of Strings) Array of prefixes this AS Prepend applies to. If this property is absent, the AS Prepend applies to all prefixes.
- `export_route_filters` - (Optional, List) List of Export Route Filter configuration information. Applicable only for create_gateway_approve requests.
  
  Nested scheme for `export_route_filter`:
  - `action` - (Required, String) Determines whether the  routes that match the prefix-set will be permit or deny
  - `prefix` - (Required, String) IP prefix representing an address and mask length of the prefix-set
  - `ge` - (Optional, Integer) The minimum matching length of the prefix-set
  - `le` - (Optional, Integer) The maximum matching length of the prefix-set
- `import_route_filters` - (Optional, List) List of Import Route Filter configuration information. Applicable only  for create_gateway_approve requests.
   
   Nested scheme for `import_route_filter`:
   - `action` - (Required, String) Determines whether the  routes that match the prefix-set will be permit or deny
   - `prefix` - (Required, String) IP prefix representing an address and mask length of the prefix-set
   - `ge` - (Optional, Integer) The minimum matching length of the prefix-set
   - `le` - (Optional, Integer) The maximum matching length of the prefix-set
- `authentication_key` - (Optional, String) BGP MD5 authentication key.Applicable only create_gateway_approve for requests.
- `bfd_interval` - (String) Minimum interval in milliseconds at which the local routing device transmits hello packets and then expects to receive a reply from a neighbor with which it has established a BFD session.Applicable only for create_gateway_approve requests.
- `bfd_multiplier` - (String) The number of hello packets not received by a neighbor that causes the originating interface to be declared down.Applicable only for create_gateway_approve requests.
- `connection_mode` - (Optional, String) Type of network connection that you want to bind to your direct link. Allowed values are `direct` and `transit`. Applicable only for create_gateway_approve requests.
- `global`- (Required,Bool) Required-Gateway with global routing as **true** can connect networks outside your associated region.Applicable only for create_gateway_approve requests.
- `metered`- (Required, Bool) Metered billing option. If set **true** gateway usage is billed per GB. Otherwise, flat rate is charged for the gateway. Applicable only for create_gateway_approve requests.
- `resource_group` - (Required, Forces new resource, String) The resource group id. If unspecified, the account's default resource group is used.Applicable only for create_gateway_approve requests.
- `default_export_route_filter` - (Optional,String) The default directional route filter action  that applies to routes that do not match any directional route filters. Applicable only for create_gateway_approve requests.
- `default_import_route_filter` - (Optional,String) The default directional route filter action  that applies to routes that do not match any directional route filters. Applicable only for create_gateway_approve requests.

## Attribute reference
In addition to all argument references list, you can access the following attribute references after your resource is approved.


- `as_prepends` - (List) List of AS Prepend configuration information.
    Nested scheme for `as_prepend`:
    - `created_at`- (String) The date and time AS Prepend was created.
    - `id` - (String) The unique identifier for this AS Prepend.   
    - `updated_at`- (String) The date and time AS Prepend was updated.
- `bgp_asn`- (Required, Integer) The BGP ASN of the gateway to be created. For example, `64999`.
- `bgp_ibm_asn`- (Required, Integer) The IBM BGP ASN.
- `bgp_base_cidr` - (Optional, String) (Deprecated) The BGP base CIDR of the gateway to be created. See `bgp_ibm_cidr` and `bgp_cer_cidr` for details on how to create a gateway by using  automatic or explicit IP assignment. Any `bgp_base_cidr` value set will be ignored.
- `bgp_cer_cidr` - (Optional, String) The BGP customer edge router CIDR. Specify a value within `bgp_base_cidr`.  For auto IP assignment, omit `bgp_cer_cidr` and `bgp_ibm_cidr`. IBM will automatically select values for `bgp_cer_cidr` and `bgp_ibm_cidr`.
- `bgp_ibm_cidr` - (Optional, String) The BGP IBM CIDR. For auto IP assignment, omit `bgp_cer_cidr` and `bgp_ibm_cidr`. IBM will automatically select values for `bgp_cer_cidr` and `bgp_ibm_cidr`.
- `carrier_name` - (Required, Forces new resource, String) The carrier name is required for `dedicated` type. Constraints are 1 ≤ length ≤ 128, Value must match regular expression ^[a-z][A-Z][0-9][ -_]$. For example, `myCarrierName`.
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
- `operational_status` - (String) The gateway operational status. For gateways pending LOA approval, patch operational_status to the appropriate value to approve or reject its LOA. For example, `loa_accepted`.
- `provider_api_managed` - (String) Indicates whether gateway changes need to be made via a provider portal.
- `vlan` - (String) The VLAN allocated for the gateway. You can set only for `type=connect` gateways created directly through the IBM portal.
- `port` - (Required, Forces new resource, String) The gateway port for type is connect gateways. This parameter is required for Direct Link connect type.

**Note**
The `Operational_status(Gateway operational status)` and `loa_reject_reason(LOA reject reason)` cannot be updated by using Terraform as the status and reason keeps changing with the different workflow actions.


## Import
The `ibm_dl_gateway_action` resource can be imported by using gateway ID. 

**Syntax**

---
```
$ terraform import ibm_dl_gateway_action.example <gateway_ID>
```
---
**Example**

---
```
$ terraform import ibm_dl_gateway_action.example 5ffda12064634723b079acdb018ef308
```
---
---
```go
terraform Destroy : because  ibm_dl_gateway_action depends on ibm_dl_provider_gateway follow the below steps to destroythe resources.
 Destroy the ibm_dl_provider_gateway 
 Ex: terraform destroy  -target="ibm_dl_provider_gateway.test_dl_gateway" -target="data.ibm_dl_provider_ports.test_ds_dl_ports"
 Destroy the ibm_dl_gateway_action
 Ex:terraform destroy  -target="ibm_dl_gateway_action.test_dl_gateway_action" 
```
---
