---

layout: "ibm"
page_title: "IBM : dl_gateway"
sidebar_current: "docs-ibm-datasources-dl-gateway"
description: |-
Manages IBM Cloud Infrastructure Directlink gatway.

---

# ibm\_dl_gateway

Import the details of an existing IBM Cloud Infrastructure directlink gateway and its virtual connectionsas a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl

	   data "ibm_dl_gateway" "test_dl_gateway_vc" {
			name = "mygateway"
		 }"

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The unique user-defined name for this gateway.

## Attribute Reference

The following attributes are exported:
* `bgp_asn` - Customer BGP ASN.
* `created_at` - The date and time resource was created.
* `crn` - The CRN (Cloud Resource Name) of this gateway.
* `global` - Gateways with global routing (true) can connect to networks outside their associated region.
* `id` - The unique identifier of this gateway.
* `location_display_name` - Gateway location long name.
* `location_name` - Gateway location.
* `metered` - Metered billing option. When true gateway usage is billed per gigabyte. When false there is no per gigabyte usage charge, instead a flat rate is charged for the gateway.
* `operational_status` - Gateway operational status.
* `resource_group` - Resource group identifier.
* `speed_mbps` - Gateway speed in megabits per second.
* `type` - Gateway type.
* `bgp_base_cidr` - (DEPRECATED) BGP base CIDR is deprecated and no longer recognized the Direct Link APIs. See bgp_cer_cidr and bgp_ibm_cidr fields instead for IP related information. Deprecated field bgp_base_cidr will be removed from the API specificiation after 15-MAR-2021.
* `bgp_cer_cidr` - BGP customer edge router CIDR.
* `bgp_ibm_asn` - IBM BGP ASN.
* `bgp_ibm_cidr` - BGP IBM CIDR.
* `bgp_status` - Gateway BGP status.
* `completion_notice_reject_reason` - Reason for completion notice rejection. Only included on type=dedicated gateways with a rejected completion notice.
* `cross_connect_router` - Cross connect router. Only included on type=dedicated gateways.
* `link_status` - Gateway link status. Only included on type=dedicated gateways.
* `port` - Port Identifier.
* `provider_api_managed` - Indicates whether gateway was created through a provider portal. If true, gateway can only be changed or deleted through the corresponding provider portal.
* `vlan` - VLAN allocated for this gateway. Only set for type=connect gateways created directly through the IBM portal.


* `virtual_connections` - List of the specified gateway's virtual connections
  * `created_at` - The date and time resource was created.
  * `id` - The unique identifier for this virtual connection.Example: ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4
  * `name` - The user-defined name for this virtual connection. Virtualconnection names are unique within a gateway. This is the name of thevirtual connection itself, the network being connected may have its ownname attribute.
  * `status` - Status of the virtual connection.
  Possible values: [pending,attached,approval_pending,rejected,expired,deleting,detached_by_network_pending,detached_by_network]
  Example: attached
  * `type` - Virtual connection type.
  Possible values: [classic,vpc]
  Example: vpc
  * `network_account` - For virtual connections across two different IBM Cloud Accounts network_account indicates the account that owns the target network.Example: 00aa14a2e0fb102c8995ebefff865555
  * `network_id` - Unique identifier of the target network. For type=vpc virtual connections this is the CRN of the target VPC. This field does not apply to type=classic connections.
  Example: crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb
