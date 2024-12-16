# Examples for IBM Cloud Logs Routing

These examples illustrate how to use the resources and data sources associated with IBM Cloud Logs Routing.

The following resources are supported:
* ibm_logs_router_tenant

The following data sources are supported:
* ibm_logs_router_tenants
* ibm_logs_router_targets

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## IBM Cloud Logs Routing resources

### Resource: ibm_logs_router_tenant

```hcl
resource "ibm_logs_router_tenant" "logs_router_tenant_instance" {
  name = var.logs_router_tenant_name
  region = "us-east"
  targets {
    log_sink_crn = "crn:v1:bluemix:public:logdna:eu-de:a/7246b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
    name = "my-log-sink"
    parameters {
      host = "www.example.com"
      port = 443
      access_credential = "new-credential"
    }
  }
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | The name for this tenant. The name is regionally unique across all tenants in the account. | `string` | true |
| region | The region to onboard this tenant. | `string` | true |
| targets | List of targets. | `list()` | true |
| targets.log_sink_crn | CRN of the Mezmo or Cloud Logs instance to sends logs to | `string` | true |
| targets.name | The name for this target. The name is regionally unique for this tenant. | `string` | true |
| targets.parameters.host | Host name of the log-sink | `string` | true |
| targets.parameters.port | Network port of the log-sink | `integer` | true |
| targets.parameters.access_credential | Secret to connect to the Mezmo log-sink. This is not required for log-sink of type Cloud Logs | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | Time stamp the tenant was originally created. |
| updated_at | Time stamp the tenant was last updated. |
| crn | Cloud resource name of the tenant. |
| etag | Resource version identifier. |
| targets | List of targets. |

## IBM Cloud Logs Routing data sources

### Data source: ibm_logs_router_tenants

```hcl
data "ibm_logs_router_tenants" "logs_router_tenants_instance" {
  name = var.logs_router_tenants_name
  region = "us-east"
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | The name of a tenant. | `string` | true |
| region | The region to query the tenant. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| tenants | List of tenants in the account. |

### Data source: ibm_logs_router_targets

```hcl
data "ibm_logs_router_targets" "logs_router_targets_instance" {
  tenant_id = var.logs_router_targets_tenant_id
  region = "us-east"
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| tenant_id | The instance ID of the tenant. | `` | true |
| region | The region where the tenant for this target exists, | `string` | true |
| name | Optional: Name of the tenant target. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| targets | List of targets of a tenant. |

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |
