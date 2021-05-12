---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_virtual_connection"
description: |-
  Manages IBM Direct Link Gateway Virtual Connection.
---

# `ibm_dl_virtual_connection`

Create, update, or delete a Direct Link Gateway Virtual Connection by using the Direct Link Gateway resource. For more information, about Direct Link Gateway Virtual Connection, see [Adding virtual connections to a Direct Link gateway](https://cloud.ibm.com/docs/dl?topic=dl-add-virtual-connection).


## Example usage
```
resource "ibm_dl_virtual_connection" "test_dl_gateway_vc"{
		gateway = ibm_dl_gateway.test_dl_gateway.id
		name = "my_dl_vc"
		type = "vpc"
		network_id = ibm_is_vpc.test_dl_vc_vpc.resource_crn   
}  
		
```

## Argument reference
Review the input parameters that you can specify for your resource. 

- `gateway` - (Required, Forces new resource, String) The Direct Link Gateway ID. 
- `name` - (Required, String) The user-defined name for this virtual connection. The virtual connection names are unique within a gateway. This is the name of the virtual connection itself, the network being connected may have its own name attribute. For `type=vpc` virtual connections it is the CRN of the target VPC. This parameter does not apply for `type=classic` connections. For example, `crn:v1:bluemix:public:is:us-east:a/28e4d90ac7504be69447111122223333::vpc:aaa81ac8-5e96-42a0-a4b7-6c2e2d1bb`.
- `network_id` - (Required, Forces new resource, String) Metered billing option. If set **true** gateway usage is billed per GB. Otherwise, flat rate is charged for the gateway.
- `type` - (Required, Forces new resource, String) The type of virtual connection. Allowed values are `classic`,`vpc`.


## Attribute reference
Review the output parameters that you can access after your resource are exported. 

- `created_at` - (String) The date and time resource created.
- `id` - (String) The unique ID of the resource with combination of gateway / virtual_connection_id.
- `virtual_connection_id` - (String) The unique identifier for the Direct Link Gateway virtual connection.
- `status` - (String) The status of the virtual connection. Possible values are `pending`, `attached`, `approval_pending`, `rejected`, `expired`, `deleting`, `detached_by_network_pending`, `detached_by_network`. For example, `attached`.
- `network_account` - (String) The virtual connections across two different IBM Cloud accounts network_account indicates the account that owns the target network. For example, `00aa14a2e0fb102c8995ebeff65555`.

## Import
The `ibm_dl_gateway_vc` can be imported by using Direct Link Gateway ID and Direct Link Gateway virtual connection ID.


**Example**

```
terraform import ibm_dl_virtual_connection.example 
d7bec597-4726-451f-8a53-e62e6f19c32c/cea6651a-bd0a-4438-9f8a-a0770bbf3ebb
```
