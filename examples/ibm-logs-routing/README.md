# Examples for IBM Cloud Logs Routing

These examples illustrate how to use the resources and data sources associated with IBM Cloud Logs Routing.

The following resources are supported:
* ibm_logs-router_tenant

The following data sources are supported:
* ibm_logs-router_tenants
* ibm_logs-router_targets

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## IBM Cloud Logs Routing resources

### Resource: ibm_logs-router_tenant

```hcl
resource "ibm_logs-router_tenant" "logs-router_tenant_instance" {
  ibm_api_version = var.logs-router_tenant_ibm_api_version
  name = var.logs-router_tenant_name
  targets {
    log_sink_crn = "crn:v1:bluemix:public:logdna:eu-de:a/7246b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
    name = "my-log-sink"
    parameters {
      host = "www.example.com"
      port = 1
      access_credential = "new-cred"
    }
  }
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| ibm_api_version | Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version. | `string` | true |
| name | The name for this tenant. The name is regionally unique across all tenants in the account. | `string` | true |
| targets | List of targets. | `list()` | true |
| targets.log_sink_crn | CRN of the Mezmo or Cloud Logs instance to sends logs to | `string` | true |
| targets.name | The name for this target. The name is regionally unique for this tenant. | `string` | true |
| targets.parameters.host | Host name of the log-sin | `string` | true |
| targets.parameters.port | Network port of the log-sink | `integer` | true |
| targets.parameters.access_credential | Secret to connect to the Mezmo log sink. This is not required for log sink of type Cloud Logs | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | Time stamp the tenant was originally created. |
| updated_at | Time stamp the tenant was last updated. |
| crn | Cloud resource name of the tenant. |
| etag | Resource version identifier. |

## IBM Cloud Logs Routing data sources

### Data source: ibm_logs-router_tenants

```hcl
data "ibm_logs-router_tenants" "logs-router_tenants_instance" {
  ibm_api_version = var.logs-router_tenants_ibm_api_version
  name = var.logs-router_tenants_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibm_api_version | Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version. | `string` | true |
| name | Optional: The name of a tenant. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| tenants | List of tenants in the account. |

### Data source: ibm_logs-router_targets

```hcl
data "ibm_logs-router_targets" "logs-router_targets_instance" {
  ibm_api_version = var.logs-router_targets_ibm_api_version
  tenant_id = var.logs-router_targets_tenant_id
  name = var.logs-router_targets_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibm_api_version | Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version. | `string` | true |
| tenant_id | The instance ID of the tenant. | `` | true |
| name | Optional: Name of the tenant target. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| targets | List of target of a tenant. |

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
