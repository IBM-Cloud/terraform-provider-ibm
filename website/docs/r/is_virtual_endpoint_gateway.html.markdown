---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_virtual_endpoint_gateway"
description: |-
  Manages IBM virtual endpoint gateway.
---

# ibm_is_virtual_endpoint_gateway
Create, update, or delete a VPC endpoint gateway by using virtual endpoint gateway resource. For more information, about the VPC endpoint gateway, see [creating an endpoint gateway](https://cloud.ibm.com/docs/vpc?topic=vpc-ordering-endpoint-gateway).

## Example usage
The following example, creates a VPN gateway.

```terraform
resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway1" {

  name = "my-endpoint-gateway-1"
  target {
	name          = "ibm-ntp-server"
    resource_type = "provider_infrastructure_service"
  }
  vpc = ibm_is_vpc.testacc_vpc.id
  resource_group = data.ibm_resource_group.test_acc.id
}

resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway2" {
	name = "my-endpoint-gateway-1"
	target {
	  name          = "ibm-ntp-server"
	  resource_type = "provider_infrastructure_service"
	}
	vpc = ibm_is_vpc.testacc_vpc.id
	ips {
		subnet   = ibm_is_subnet.testacc_subnet.id
		name        = "test-reserved-ip1"
	}
	resource_group = data.ibm_resource_group.test_acc.id
}

resource "ibm_is_virtual_endpoint_gateway" "endpoint_gateway3" {
	name = "my-endpoint-gateway-1"
	target {
	  name          = "ibm-ntp-server"
	  resource_type = "provider_infrastructure_service"
	}
	vpc = ibm_is_vpc.testacc_vpc.id
	ips {
		id   = "0737-5ab3c18e-6f6c-4a69-8f48-20e3456647b5"
	}
	resource_group = data.ibm_resource_group.test_acc.id
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `name` - (Required, Forces new resource, String) The endpoint gateway name.
- `ips`  (Optional, List) The endpoint gateway resource group.

  Nested scheme for `ips`:
  - `id` - (Optional, String) The endpoint gateway resource group IPs ID.
  - `name` - (Optional, String) The endpoint gateway resource group IPs name.
  - `subnet` - (Optional, String) The endpoint gateway resource group subnet ID.
  - `resource_type` - (Required, String) The endpoint gateway resource group VPC resource type.
  
  **NOTE**: `id` and `subnet` are mutually exclusive.

- `resource_group` - (Optional, Forces new resource, String) The resource group ID.
- `tags`- (Optional, Array of Strings) A list of tags associated with the instance.
- `target` - (Required, List) The endpoint gateway target.

  Nested scheme for `target`:
  - `crn` - (Optional, Forces new resource, String) The endpoint gateway target `CRN`. If CRN not specified, `name` must be specified. 
  - `name` - (Required, Forces new resource, String) The endpoint gateway target name.
  - `resource_type` - (Required, String) The endpoint gateway target resource type.
- `vpc` - (Required, Forces new resource, String) The VPC ID.

**NOTE**: `ips` configured inline in this resource are not modifiable. Prefer using `ibm_is_virtual_endpoint_gateway_ip` resource to bind/unbind new reserved IPs to endpoint gateways and use the resource `ibm_is_subnet_reserved_ip` to create new reserved IP.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `created_at` - (Timestamp) The created date and time of the endpoint gateway.
- `health_state` - (String) The health state of the endpoint gateway.
- `id` - (String) The unique identifier of the VPN gateway connection. The ID is composed of `<gateway_id>`.
- `lifecycle_state` - (String) The lifecycle state of the endpoint gateway.
- `resource_type` - (String) The endpoint gateway resource type.


## Import
The `ibm_is_virtual_endpoint_gateway` resource can be imported by using virtual endpoint gateway ID.

**Example**

```
$ terraform import ibm_is_virtual_endpoint_gateway.example d7bec597-4726-451f-8a63-xxxxsdf345
```
