---
subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_connection"
description: |-
  Manages IBM Transit Gateway Connection.
---

# ibm\_tg_connection

Provides a transit gateway connection resource. This allows transit gateway's connection to be created, and updated and deleted.

## Example Usage

```hcl
resource "ibm_tg_connection" "test_ibm_tg_connection"{
		gateway = ibm_tg_gateway.test_tg_gateway.id
		network_type = "vpc"
		name= "myconnection"
		network_id = ibm_is_vpc.test_tg_vpc.resource_crn
}
  
```

## Argument Reference

The following arguments are supported:
* `gateway` - (Required, Forces new resource, string) The Transit Gateway identifier.
* `name` - (Optional, string) The user-defined name for this transit gateway. If unspecified, the name will be the network name (the name of the VPC in the case of network type 'vpc', and the word Classic, in the case of network type 'classic').
* `network_type` - (Required, Forces new resource, string) Defines what type of network is connected via this connection.Allowable values: [classic,vpc]. Example: vpc
* `network_id` - (Optional,Forces new resource,string) The ID of the network being connected via this connection. This field is required for some types, such as 'vpc'. For network type 'vpc' this is the CRN of the VPC to be connected. This field is required to be unspecified for network type 'classic'. Example: crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b   
* `network_account_id` (Optional,Forces new resource,string) - The ID of the account which owns the network that is being connected. Generally only used if the network is in a different account than the gateway.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the resource.Its combination of gatewayID/connectionID
* `connection_id` - The unique identifier for this Transit Gateway Connection to Network (vpc/classic). 
* `created_at` - The date and time that this connection was created.
* `updated_at` - The date and time that this connection was last updated.
* `status` - What is the current configuration state of this connection
Possible values: [attached,failed,pending,deleting]
* `request_status` - Only visible for cross account connections, this field represents the status of the request to connect the given network between accounts . Possible values: [pending,approved,rejected,expired,detached]

**NOTE** If the the user is provisioning the cross-account gateway/connection the resource doesn't wait for the available status. It goes into provisioning status where the user need to complete the manual approval process

## Import

ibm_tg_connection can be imported using transit gateway id and connection id, eg

```
$ terraform import ibm_tg_connection.example 5ffda12064634723b079acdb018ef308/cea6651a-bd0a-4438-9f8a-a0770bbf3ebb
```