# Example for IBM Logs Router V1

This example illustrates how to use IBM LogsRouterV1

The following types of resources are supported:

* logs_router_tenant

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## IbmLogsRouterV1 resources

logs_router_tenant resource:

```hcl
resource "logs_router_tenant" "logs_router_tenant_instance" {
  target_type = var.logs_router_tenant_target_type
  target_host = var.logs_router_tenant_target_host
  target_port = var.logs_router_tenant_target_port
  target_instance_crn = var.logs_router_tenant_target_instance_crn
}
```

## IbmLogsRouterV1 data sources

logs_router_tenant data source:

```hcl
data "logs_router_tenant" "logs_router_tenant_instance" {
  tenant_id = ibm_logs_router_tenant.logs_router_tenant_instance.id
}
```

## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| target_type | Type of log-sink. | `string` | true |
| target_host | Host name of log-sink. | `string` | true |
| target_port | Network port of log sink. | `number` | true |
| target_instance_crn | Cloud resource name of the log-sink target instance. | `string` | true |
| tenant_id | The instance ID of the tenant. | `` | true |

## Outputs

| Name | Description |
|------|-------------|
| logs_router_tenant | logs_router_tenant object |
| logs_router_tenant | logs_router_tenant object |
