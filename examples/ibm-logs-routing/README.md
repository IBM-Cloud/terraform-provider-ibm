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
    log_sink_crn = "crn:v1:bluemix:public:logs:eu-de:a/7246b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
    name = "my-log-sink"
    parameters {
      host = "www.example.com"
      port = 443
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
| targets.log_sink_crn | CRN of the Cloud Logs instance to sends logs to | `string` | true |
| targets.name | The name for this target. The name is regionally unique for this tenant. | `string` | true |
| targets.parameters.host | Host name of the log-sink | `string` | true |
| targets.parameters.port | Network port of the log-sink | `integer` | true |

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

## Notes

### The Logs Routing URL can be set in endpoints.json

You can declare the service endpoints in a JSON file and either reference this file in your provider block by using the `endpoints_file_path` argument, or export the path to your file with the `IBMCLOUD_ENDPOINTS_FILE_PATH` or `IC_ENDPOINTS_FILE_PATH` environment variable.
To use the provided endpoints file, set the visibility to either `public` or `pivate` by using the `IC_VISIBILITY` or `IBMCLOUD_VISIBILITY` environment variable, or by setting the `visibility` field in your provider block.

**Example**:

```json
{
    "IBMCLOUD_LOGS_ROUTING_API_ENDPOINT":{
        "public":{
            "us-south":"<endpoint>",
            "us-east":"<endpoint>",
            "eu-gb":"<endpoint>",
            "eu-de":"<endpoint>"
        },
        "private":{
            "us-south":"<endpoint>",
            "us-east":"<endpoint>",
            "eu-gb":"<endpoint>",
            "eu-de":"<endpoint>"
        }
    }
}
```

As of 28 March 2024 the Log Analysis service is deprecated and will no longer be supported as of 30 March 2025.
IBM Cloud Logs will stop supporting `logdna` targets at the same time and no logs will be routed to these type of targets after that date.
You should make sure that you have configured your tenant to direct your logs to another destination before 30 March 2025.
Any `logdna` targets still configured after 30 April 2025 will be removed automatically from your tenant configuration.