---
subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_connection_action"
description: |-
  Perform actions on a connection for a Transit Gateway
---

# ibm_tg_connection_action
Create an action to approve or reject a cross account connection resource. For more information, about Transit Gateway connection, see [adding a cross-account connection](https://cloud.ibm.com/docs/transit-gateway?topic=transit-gateway-adding-cross-account-connections)

## Example usage

```terraform
resource "ibm_tg_connection_action" "test_tg_cross_connection_approval" {
    provider = ibm.account2
    gateway = ibm_tg_gateway.new_tg_gw.id
    connection_id = ibm_tg_conneciton.test_ibm_tg_connection.connection_id
    action = "approve"
}
  
```

## Argument reference
Review the argument references that you can specify for your resource. 
 
- `gateway` - (Required, String) The unique identifier of the gateway.
- `connection_id` - (Required, String) The unique identifier of the gateway connection
- `action` - (Required, String) Whether to approve or reject the cross account connection
