---

subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_provider_gateway"
description: |-
  Manages IBM Direct Link Provider Gateway.
---

# ibm\_dl_provider_gateway

Provides a direct link provider gateway resource. This allows direct link provider gateway to be created, and updated and deleted.

## Example Usage
In the following example, you can create Direct link Provider gateway:

```hcl

resource ibm_dl_provider_gateway test_dl_provider_gateway {
  bgp_asn =  64999
  bgp_ibm_cidr =  "169.254.0.29/30"
  bgp_cer_cidr =  "169.254.0.30/30"
  name = "Gateway1"
  speed_mbps = 1000 
  port = "434-c749-4f1d-b190-22"
  customer_account_id = "0c474da-c749-4f1d-b190-2333"
}   
```

## Argument Reference

The following arguments are supported:

* `bgp_asn` - (Required, Forces new resource, integer) The BGP ASN of the Gateway to be created. Example: 64999
* `name` - (Required, boolean) The unique user-defined name for this gateway. Example: myGateway
* `speed_mbps` - (Required, integer) Gateway speed in megabits per second. Example: 1000
* `bgp_cer_cidr` - (Optional, Forces new resource, string) BGP customer edge router CIDR. Specify a value within bgp_base_cidr. If bgp_base_cidr is 169.254.0.0/16, this field can be ommitted and a CIDR will be selected automatically. Example: 10.254.30.78/30
* `bgp_ibm_cidr` - (Optional, Forces new resource, string) BGP IBM CIDR. Specify a value within bgp_base_cidr. If bgp_base_cidr is 169.254.0.0/16, this field can be ommitted and a CIDR will be selected automatically. Example: 10.254.30.77/30 
* `customer_account_id` - (Required,Forces new resource, string) Customer IBM Cloud account ID for the new gateway. A gateway object containing the pending create request will become available in the specified account.
* `port` - (Required , Forces new resource, string) gateway port



## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of this gateway. 
* `name` - The unique user-defined name for this gateway. 
* `crn` - The CRN (Cloud Resource Name) of this gateway. 
* `created_at` - The date and time resource was created.
* `bgp_asn` - IBM BGP ASN.
* `bgp_status` - Gateway BGP status.
* `customer_account_id` - Customer IBM Cloud account ID for the new gateway. A gateway object containing the pending create request will become available in the specified account.
* `port` - gateway port for type=connect gateways
* `vlan` - VLAN allocated for this gateway. Only set for type=connect gateways created directly through the IBM portal. 
* `provider_api_managed` - Indicates whether gateway changes must be made via a provider portal.
* `operational_status` - Gateway operational status.
Possible values:[configuring, create_pending, create_rejected, delete_pending, provisioned ]

## Import

ibm_dl_provider_gateway can be imported using gateway id, eg

```
$ terraform import ibm_dl_provider_gateway.test_dl_provider_gateway 5ffda12064634723b079acdb018ef308
```
