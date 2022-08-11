---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_virtual_endpoint_gateway"
description: |-
  Manages IBM virtual endpoint gateway.
---

# ibm_is_virtual_endpoint_gateway
Create, update, or delete a VPC endpoint gateway by using virtual endpoint gateway resource. For more information, about the VPC endpoint gateway, see [creating an endpoint gateway](https://cloud.ibm.com/docs/vpc?topic=vpc-ordering-endpoint-gateway).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
The following example, creates a Virtual Private Endpoint gateway.

```terraform
resource "ibm_is_virtual_endpoint_gateway" "example" {

  name = "example-endpoint-gateway"
  target {
    name          = "ibm-ntp-server"
    resource_type = "provider_infrastructure_service"
  }
  vpc            = ibm_is_vpc.example.id
  resource_group = data.ibm_resource_group.example.id
  security_groups = [ibm_is_security_group.example.id]
}

resource "ibm_is_virtual_endpoint_gateway" "example1" {
  name = "example-endpoint-gateway-1"
  target {
    name          = "ibm-ntp-server"
    resource_type = "provider_infrastructure_service"
  }
  vpc = ibm_is_vpc.example.id
  ips {
    subnet = ibm_is_subnet.example.id
    name   = "example-reserved-ip"
  }
  resource_group = data.ibm_resource_group.example.id
  security_groups = [ibm_is_security_group.example.id]
}

resource "ibm_is_virtual_endpoint_gateway" "example3" {
  name = "example-endpoint-gateway-2"
  target {
    name          = "ibm-ntp-server"
    resource_type = "provider_infrastructure_service"
  }
  vpc = ibm_is_vpc.example.id
  ips {
    subnet = ibm_is_subnet.example.id
    name   = "example-reserved-ip"
  }
  resource_group = data.ibm_resource_group.example.id
  security_groups = [ibm_is_security_group.example.id]
}

resource "ibm_is_virtual_endpoint_gateway" "example4" {
  name = "example-endpoint-gateway-3"
  target {
    crn           = "crn:v1:bluemix:public:cloud-object-storage:global:::endpoint:s3.direct.mil01.cloud-object-storage.appdomain.cloud"
    resource_type = "provider_cloud_service"
  }
  vpc            = ibm_is_vpc.example.id
  resource_group = data.ibm_resource_group.example.id
  security_groups = [ibm_is_security_group.example.id]
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
  
  ~> **NOTE:** `id` and `subnet` are mutually exclusive.

- `resource_group` - (Optional, Forces new resource, String) The resource group ID.
- `security_groups` - (Optional, list) The security groups to use for this endpoint gateway. If unspecified, the VPC's default security group is used.
  **NOTE:** either of `ibm_is_security_group_target` resource or `security_groups` attribute should be used, both can't be use together.
- `tags`- (Optional, Array of Strings) A list of tags associated with the instance.
- `target` - (Required, List) The endpoint gateway target.

  Nested scheme for `target`:
  - `crn` - (Optional, Forces new resource, String) The CRN for this provider cloud service, or the CRN for the user's instance of a provider cloud service.

    **NOTE:** If `crn` is not specified, `name` must be specified. 
  - `name` - (Optional, Forces new resource, String) The endpoint gateway target name.

    **NOTE:** If `name` is not specified, `crn` must be specified. 
  - `resource_type` - (Required, String) The endpoint gateway target resource type. The possible values are `provider_cloud_service`, `provider_infrastructure_service`.
- `vpc` - (Required, Forces new resource, String) The VPC ID.

~> **NOTE:** `ips` configured inline in this resource are not modifiable. Prefer using `ibm_is_virtual_endpoint_gateway_ip` resource to bind/unbind new reserved IPs to endpoint gateways and use the resource `ibm_is_subnet_reserved_ip` to create new reserved IP.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `created_at` - (Timestamp) The created date and time of the endpoint gateway.
- `crn` - (String) The CRN for this endpoint gateway.
- `health_state` - (String) The health state of the endpoint gateway.
- `id` - (String) The unique identifier of the VPE Gateway. The ID is composed of `<gateway_id>`.
- `ips`  (List) The endpoint gateway reserved ips.

  Nested scheme for `ips`:
  - `address` -  The endpoint gateway IPs Address.
  - `id` -  The endpoint gateway resource group IPs ID.
  - `name` -  The endpoint gateway resource group IPs name.
  - `resource_type` -  The endpoint gateway resource group VPC resource type.

- `lifecycle_state` - (String) The lifecycle state of the endpoint gateway.
- `resource_type` - (String) The endpoint gateway resource type.

## Import
The `ibm_is_virtual_endpoint_gateway` resource can be imported by using virtual endpoint gateway ID.

**Example**

```
$ terraform import ibm_is_virtual_endpoint_gateway.example d7bec597-4726-451f-8a63-xxxxsdf345
```
