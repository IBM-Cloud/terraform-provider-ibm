---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_provider_gateway"
description: |-
  Manages IBM Direct Link Provider Gateway.
---

# ibm_dl_provider_gateway

Create, update, or delete a Direct Link Provider Gateway by using the Direct Link Provider Gateway resource. For more information, about Direct Link Provider Gateway, see [about Direct Link](https://cloud.ibm.com/docs/dl?topic=dl-dl-about#use-case-connect).


## Example usage to create Direct Link of dedicated type
In the following example, you can create Direct Link provider gateway:

```terraform
resource ibm_dl_provider_gateway test_dl_provider_gateway {
  bgp_asn =  64999
  bgp_ibm_cidr =  "169.254.0.29/30"
  bgp_cer_cidr =  "169.254.0.30/30"
  name = "Gateway1"
  speed_mbps = 1000 
  port = "434-c749-4f1d-b190-22"
  customer_account_id = "0c474da-c749-4f1d-b190-2333"
  vlan = 35
} 
```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `bgp_asn`- (Required, Integer) The BGP ASN of the gateway to be created. For example, `64999`.
- `bgp_cer_cidr` - (Optional, String) The BGP customer edge router CIDR. Specify a value within bgp_base_cidr. If bgp_base_cidr is `169.254.0.0/16`, this parameter can exclude and a CIDR is selected automatically. For example, `10.254.30.78/30`.
- `bgp_ibm_cidr` - (Optional, String) The IBM BGP CIDR. Specify a value within bgp_base_cidr. If bgp_base_cidr is `169.254.0.0/16`, this parameter can exclude and a CIDR is selected automatically. For example, `10.254.30.77/30`.
- `customer_account_id` - (Required, Forces new resource, String) The customer IBM Cloud account ID for the new gateway. A gateway object contains the pending create request to be available in the specified account.
- `name` - (Required, String) The unique user-defined name for this gateway. Example: `myGateway`.
- `port` - (Required, Forces new resource, String) The gateway port for type to connect gateway.
- `speed_mbps`- (Required, Integer) The gateway speed in megabits per second. For example, `10.254.30.78/30`.
- `vlan` - (Optional, Integer) VLAN requested for this gateway.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created. 

- `bgp_asn` - (String) The IBM BGP ASN.
- `bgp_status` - (String) The gateway BGP status.
- `crn` - (String) The CRN of the gateway.
- `created_at` - (String) The date and time resource created.
- `customer_account_id` - (String) The customer IBM Cloud account ID for the new gateway. A gateway object contains the pending create request to be available in the specified account.
- `id` - (String) The unique ID of the gateway.
- `name` - (String) The unique user-defined name for the gateway.
- `operational_status` - (String) The gateway operational status. Supported values are`configuring`, `create_pending`, `create_rejected`, `delete_pending`, `provisioned`.
- `port` - (String) The gateway port for `type=connect` gateways.
- `provider_api_managed` - (String) Indicates whether the gateway changes need to be made via a provider portal.
- `vlan` - (String) VLAN requested for this gateway.

## Import
The `ibm_dl_provider_gateway` resource can be imported by using gateway ID. 

**Syntax**

```
$ terraform import ibm_dl_provider_gateway.<gateway_ID>
```

**Example**

```
$ terraform import ibm_dl_provider_gateway.test_dl_provider_gateway 5ffda12064634723b079acdb018ef308
```


