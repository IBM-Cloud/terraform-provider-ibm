---
layout: "ibm"
page_title: "IBM : dl_virtual_connection"
sidebar_current: "docs-ibm-resource-dl-virtual-connection"
description: |-
  Manages IBM Direct Link Gateway Virtual Connection.
---

# ibm\_dl_virtual_connection

Provides a direct link gateway virtual connection resource. This allows direct link gateway virtual connection to be created, and updated and deleted.

## Example Usage

```hcl
resource "ibm_dl_virtual_connection" "test_dl_gateway_vc"{
		gateway = ibm_dl_gateway.test_dl_gateway.id
		name = "my_dl_vc"
		type = "vpc"
		network_id = ibm_is_vpc.test_dl_vc_vpc.resource_crn   }  

```

## Argument Reference

The following arguments are supported:

* `gateway` - (Required, Forces new resource, string) The Direct Link gateway identifier. 
* `name` - (Required,string) The user-defined name for this virtual connection. Virtualconnection names are unique within a gateway. This is the name of thevirtual connection itself, the network being connected may have its ownname attribute.
* `type` - (Required, Forces new resource, string)The type of virtual connection.Allowable values: [classic,vpc]. Example: vpc
* `network_id` -  (Required,Forces new resource, string) Unique identifier of the target network. For type=vpc virtual connections this is the CRN of the target VPC. This field does not apply to type=classic connections.Example: crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bbbbb

## Attribute Reference

The following attributes are exported:

* `created_at` - The date and time resource was created.
* `id` - The unique identifier for this virtual connection.Example: ef4dcb1a-fee4-41c7-9e11-9cd99e65c1f4
* `status` - Status of the virtual connection.
Possible values: [pending,attached,approval_pending,rejected,expired,deleting,detached_by_network_pending,detached_by_network]
Example: attached
* `network_account` - For virtual connections across two different IBM Cloud Accounts network_account indicates the account that owns the target network.Example: 00aa14a2e0fb102c8995ebefff865555

## Import

ibm_dl_gateway_vc can be imported using directlink gateway ID and directlink gateway virtual connection ID, eg

```
$ terraform import ibm_dl_virtual_connection.example 
d7bec597-4726-451f-8a53-e62e6f19c32c/cea6651a-bd0a-4438-9f8a-a0770bbf3ebb
```