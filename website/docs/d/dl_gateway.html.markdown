---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_gateway"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Gateway.
---

# `ibm_dl_gateway`

Import the details of an existing IBM Cloud Infrastructure Direct Link Gateway and its virtual connections. For more information, about IBM Cloud Direct Link, see [Getting started with IBM Cloud Direct Link](https://cloud.ibm.com/docs/dl?topic=dl-get-started-with-ibm-cloud-dl).


## Example usage

```
      data "ibm_dl_gateway" "test_dl_gateway_vc" {
			name = "mygateway"
		 }
```


## Argument reference
Review the Argument reference that you can specify for your resource. 

- `name` - (Required, String) The unique user-defined name for the gateway.


## Attribute reference
Review the attribute reference that you can access after your resource is created. 

- `bgp_asn` - (String) Customer BGP ASN.
- `bgp_base_cidr` - (String) The BGP base CIDR.
- `bgp_cer_cidr` - (String) The BGP customer edge router CIDR.
- `bgp_ibm_asn` - (String) The IBM BGP ASN.
- `bgp_ibm_cidr` - (String) The IBM BGP  CIDR.
- `bgp_status` - (String) The gateway BGP status.
- `created_at` - (String) The date and time resource is created.
- `crn` - (String) The CRN of the gateway.
- `completion_notice_reject_reason` - (String) The reason for completion notice rejection. Only included on a dedicated gateways type with a rejected completion notice.
- `cross_connect_router` - (String) The cross connect router. Only included on a dedicated gateways type.
- `global` - (Bool) Gateway with global routing as **true** can connect networks outside your associated region.
- `id` - (String) The unique identifier of the gateway.
- `location_display_name` - (String) Long name of the gateway location.
- `location_name` - (String) The location name of the gateway.
- `link_status` - (String) The gateway link status. Only included on a dedicated gateways type.
- `metered` - (String) Metered billing option. If set **true** gateway usage is billed per GB. Otherwise, flat rate is charged for the gateway.
- `operational_status` - (String) The gateway operational statu.
- `port` - (Integer) The port identifier.
- `provider_api_managed` - (Bool) Indicates the gateway is created through a provider portal. If set **true**, gateway can only be changed. If set **false**, gateway is deleted through the corresponding provider portal.
- `resource_group` - (String) The resource group identifier.
- `speed_mbps` - (String) The gateway speed in MBPS.
- `type` - (String) The gateway type.
- `virtual_connections` - (String) List of the specified gateway's virtual connections.
	- `created_at` - (String) The creation date and time resource.
	- `id` - (String) The unique identifier of the virtual connection. For example, `ef4dcbtyu1a-fee4-41c7-9e11-9cd99e65c1f4.
	- `name` - (String) The unique user-defined name of the only virtual connection in the gateway.
	- `status` - (String) The status of the virtual connection. Possible values are `pending`,`attached`,`approval_pending`,`rejected`,`expired`,`deleting`,`detached_by_network_pending`,`detached_by_network`.
	- `type` - (String) The virtual connection type. Possible values are `classic`,`vpc`. For example, `vpc`.
	- `network_account` - (String) For virtual connections across two different IBM Cloud accounts. Network_account indicates the account you own the target network. For example, `00aa14a2e0fb102c8995ebefhhhf8655556`
	- `network_id` - (String) The unique identifier of the target network. For type `vpc`, virtual connections is the CRN of the target VPC. This field do not apply for type `classic` connections. For example, `crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2dbb`.
- `vlan` - (String) The VLAN allocated for the gateway. Only set for connect gateways type created directly through the IBM portal.
