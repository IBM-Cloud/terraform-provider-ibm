# This example shows how to create and IBM Cloud Internet Services instance, monitor, pools, global load balancer, DNS Records, Firewall, Rate Limiting

This sample configuration will configure CIS instance, a health check monitor, origin pool, global load balancer, DNS Record, Firewall, Rate Limit. Also see the example  `ibm-website-multi-region` for an example of using CIS with a working website deployed across multiple regions. 

These types of resources are supported:

* [ CIS ](https://cloud.ibm.com/docs/terraform?topic=terraform-cis-resources#cis)
* [ CIS Domain ](https://cloud.ibm.com/docs/terraform?topic=terraform-cis-resources#cis-domain)
* [ CIS Domain Settings ](https://cloud.ibm.com/docs/terraform?topic=terraform-cis-resources#cis-domain-settings)
* [ CIS DNS Record ](https://cloud.ibm.com/docs/terraform?topic=terraform-cis-resources#cis-dns-record)
* [ CIS Firewall ](https://cloud.ibm.com/docs/terraform?topic=terraform-cis-resources#cis-firewall)
* [ CIS GLB ](https://cloud.ibm.com/docs/terraform?topic=terraform-cis-resources#cis-global-lb)
* [ CIS Health Check | Monitor ](https://cloud.ibm.com/docs/terraform?topic=terraform-cis-resources#cis-health)
* [ CIS Origin Pool ](https://cloud.ibm.com/docs/terraform?topic=terraform-cis-resources#cis-origin-pool)
* [ CIS Rate Limit ](https://cloud.ibm.com/docs/terraform?topic=terraform-cis-resources#cis-rate-limit)

## Terraform versions

Terraform 0.12. Pin module version to `~> v1.7`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.29.0`. Branch - `terraform_v0.11.x`.

## Running the configuration 
```shell
terraform init
terraform plan
```

For apply phase

```shell
terraform apply
```

For destroy see notes under **Costs** for how to preserve the CIS service instance to avoid additional billing costs for further instances. Otherwise destroy all resources. 

```shell
terraform destroy
```  

## Costs

This sample uses chargable services and **will** incur costs. Billing for the CIS service instance is pro-rata'd for the remaining duration of the month it is deployed in. Execution of `terraform destroy` will result in deletion of all resources including the CIS service instance. Billing for VSIs and Cloud Load Balancing will terminate on the hour. The billing for the CIS service instance will be pro-rata'd to the end of the month. For each delete and recreate of the environment a new CIS service instance will be created and result in an additional billing instance pro-rata'd for the month. 

To avoid additional CIS service instance costs if the sample confifuration is executed additional times, after creation the `ibm_cis` resource should be removed from the configuration and replaced with a `ibm_cis` data source. All dependent CIS Terraform resource definitions must also be updated to reference the `data source`. 

## CIS Resources

`IBM CLOUD CIS instance`
```hcl
resource "ibm_cis" "web_domain" {
  name              = "web_domain"
  resource_group_id = data.ibm_resource_group.web_group.id
  plan              = "standard"
  location          = "global"
}
```
`Domain settings for IBM CIS instance`
```hcl
resource "ibm_cis_domain_settings" "web_domain" {
  cis_id          = ibm_cis.web_domain.id
  domain_id       = ibm_cis_domain.web_domain.id
  waf             = "on"
  ssl             = "full"
  min_tls_version = "1.2"
}
```
`Adding valid Domain for IBM CIS instance`
```hcl
resource "ibm_cis_domain" "web_domain" {
  cis_id = ibm_cis.web_domain.id
  domain = var.domain
}
```
`CIS GLB Monitor|HealthCheck`
```hcl
resource "ibm_cis_healthcheck" "root" {
  cis_id         = ibm_cis.web_domain.id
  description    = "Websiteroot"
  expected_body  = ""
  expected_codes = "200"
  path           = "/"
}
```
`CIS Origin Pool`
```hcl
resource "ibm_cis_origin_pool" "lon" {
  cis_id        = ibm_cis.web_domain.id
  name          = var.datacenter1
  check_regions = ["WEU"]

  monitor = ibm_cis_healthcheck.root.id

  origins {
    name    = var.datacenter1
    address = "192.0.2.1"
    enabled = true
  }

  description = "LON pool"
  enabled     = true
}
```
`CIS GLB`
```hcl
resource "ibm_cis_global_load_balancer" "web_domain" {
  cis_id           = ibm_cis.web_domain.id
  domain_id        = ibm_cis_domain.web_domain.id
  name             = "${var.dns_name}${var.domain}"
  fallback_pool_id = ibm_cis_origin_pool.lon.id
  default_pool_ids = [ibm_cis_origin_pool.lon.id, ibm_cis_origin_pool.ams.id]
  description      = "Load balancer"
  proxied          = true
  session_affinity = "cookie"
}
```
`CIS DNS Record`
```hcl
resource "ibm_cis_dns_record" "example" {
  cis_id           = ibm_cis.web_domain.id
  domain_id        = ibm_cis_domain.web_domain.id
  name= var.record_name
  type= var.record_type
  content= var.record_content
  proxied=true
}
```
`CIS Firewall`
```hcl
resource "ibm_cis_firewall" "lockdown" {
  cis_id           = ibm_cis.web_domain.id
  domain_id        = ibm_cis_domain.web_domain.id
  firewall_type = var.firewall_type

  lockdown {
    paused = "true"
    urls   = [var.lockdown_url]
   
    configurations {
      target = var.firewall_target
      value  = var.firewall_value
    }
  }
}
```
`Custom Rate Limit rule:`
```hcl
resource "ibm_cis_rate_limit" "ratelimit" {
  cis_id = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.id
  threshold = var.threshold
  period = var.period
  match {
    request {
      url = var.match_request_url
      schemes = var.match_request_schemes
      methods = var.match_request_methods
    }
    response {
      status = var.match_response_status
      origin_traffic = var.match_response_traffic
      header {
        name= var.header1_name
        op= var.header1_op
        value= var.hearder1_value
      }
    }
  }
  action {
    mode = var.action_mode
    timeout = var.action_timeout
    response {
      content_type = var.action_response_content_type
      body = var.action_response_body
    }
  }
  correlate {
    by = var.correlate_by
  }
  disabled = var.disabled
  description = var.description
  bypass {
    name= var.bypass1_name
    value= var.bypass1_value
  }
}
```
`CIS Edge Functions action`
```hcl
resource "ibm_cis_edge_functions_action" "test_action" {
  cis_id      = data.ibm_cis.cis.id
  domain_id   = data.ibm_cis_domain.cis_domain.domain_id
  script_name = "sample-script"
  script      = file("./script.js")
}
```
`CIS Edge Functions trigger`
```hcl
resource "ibm_cis_edge_functions_trigger" "test_trigger" {
  cis_id    = ibm_cis_edge_functions_action.test_action.cis_id
  domain_id = ibm_cis_edge_functions_action.test_action.domain_id
  script    = ibm_cis_edge_functions_action.test_action.script_name
  pattern   = "example.domain.com/*"
}
```

## CIS Data Sources
`CIS Instance`
```hcl
data "ibm_cis" "cis" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = "CISTest"
}
```
`CIS Domain`
```hcl
data "ibm_cis_domain" "cis_domain" {
  cis_id = data.ibm_cis.cis.id
  domain = "cis-terraform.com"
}
```
`CIS Firewall`
```hcl
data "ibm_cis_firewall" "lockdown"{
  cis_id = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.id
  firewall_type = "lockdowns"
}
```
`CIS Rate Limit`
```hcl
data "ibm_cis_rate_limit" "ratelimit" {
  cis_id = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.id
}
```
`CIS Edge Functions action data source`
```hcl
data "ibm_cis_edge_functions_actions" "test_actions" {
  cis_id    = ibm_cis_edge_functions_trigger.test_trigger.cis_id
  domain_id = ibm_cis_edge_functions_trigger.test_trigger.domain_id
}
```
`CIS Edge Functions trigger data source`
```
data "ibm_cis_edge_functions_triggers" "test_triggers" {
  cis_id    = ibm_cis_edge_functions_trigger.test_trigger.cis_id
  domain_id = ibm_cis_edge_functions_trigger.test_trigger.domain_id
}
```
## Dependencies

- User has IAM security rights to create and configure an Internet Services instance
- DNS Domain registration
- [DNS Record CLI](https://cloud.ibm.com/docs/cis-cli-plugin?topic=cis-cli-plugin-cis-cli#dns-record)
- [GLB CLI](https://cloud.ibm.com/docs/cis-cli-plugin?topic=cis-cli-plugin-cis-cli#glb)
- [Firewall CLI](https://cloud.ibm.com/docs/cis-cli-plugin?topic=cis-cli-plugin-cis-cli#firewall)
- To create a custom rate limit rule the CIS instance should be a `enterprise` plan.
- [Rate Limiting Cloud Docs](https://cloud.ibm.com/docs/cis?topic=cis-cis-rate-limiting#rate-limiting-configure-response)
- [Rate Limiting CLI](https://cloud.ibm.com/docs/cis?topic=cis-cli-plugin-cis-cli#ratelimit)
- [Edge Functions CLI](https://cloud.ibm.com/docs/cis?topic=cis-cli-plugin-cis-cli#edge-functions)

## Notes

- Terraform IBM provider (via Terraform 0.12) supports only CIS Firewall - Lockdows
- Terraform IBM provider (via Terraform 0.12) supports only Create a custom rate limiting rule.

## Examples

* [CIS Examples](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-cis)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Configuration 

The following variables need to be set in the `terraform.tf` file before use

* `softlayer_username` is an Infrastructure user name. Go to https://control.bluemix.net/account/user/profile, scroll down, and check API Username.
* `softlayer_api_key` is an Infrastructure API Key. Go to https://control.bluemix.net/account/user/profile, scroll down, and check Authentication Key.
* `ibmcloud_api_key` - An API key for IBM Cloud services. If you don't have one already, go to https://cloud.ibm.com/iam/#/apikeys and create a new key.


Customise the variables in `variables.tf` to your local environment and chosen DNS domain name. 

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| resource\_group | Name of the Resource Group configured resources will be allocated to Default: `Default`| `string` | yes |
| domain | In the DNS Domain for web server registed with the DNS registrar. The DNS domain must be pre-registered with the IBM Cloud [Domain Registration Service](https://cloud.ibm.com/classic/services/domains). | `string` | yes |
| dns\_name |  DNS name (prefix) for website, including '.',e.g. 'www.' Can be "" for website to be at root of domain.| `string` | yes |
| datacenter1 | Name of origin pool in region 1.  Default: `lon2`| `string` | yes |
| datacenter2 | Name of origin pool in region 2.   Default: `ams03`| `string` | yes |
| record\_name | DNS Record Name | `string` | yes |
| record\_type | DNS Record Type | `string` | yes |
| record\_content | DNS Record Content | `string` | yes |
| firewall\_type | Firewall Type | `string` | yes |
| lockdown\_url | Lockdown URL | `string` | yes |
| lockdown\_target | Lockdown Configuration target | `string` | yes |
| lockdown\_value | Lockdown Configuration Value | `string` | yes |
| threshold | Rate Limiting Threshold | `number` | yes |
| period | Rate Limiting Period | `number` | yes |
| match\_request\_url | URL pattern of matching request | `string` | no |
| match\_request\_schemes | HTTP Schemes of matching request. It can be one or many. Example schemes 'HTTP', 'HTTPS'.If not provided API ll default to ALL. | `set(string)` | no |
| match\_request\_methods | HTTP Methos of matching request. It can be one or many. Example methods 'POST', 'PUT'.If not provided API ll default to ALL. | `set(string)` | no |
| match\_response\_status| HTTP Status Codes of matching response. It can be one or many. Example status codes '403', '401 | `set(number)` | no |
| match\_response\_traffic | Origin Traffic of matching response. | `bool` | no |
| header1\_name | The name of the response header to match. | `string` | no |
| header1\_op | The operator when matching. Valid values are 'eq' and 'ne'. | `string` | no |
| hearder1\_value | The value of the header, which is exactly matched. | `string` | no |
| action\_mode | Type of action performed.Valid values are: 'simulate', 'ban', 'challenge', 'js_challenge'. | `string` | yes |
| action\_timeout | The time to perform the mitigation action. Timeout be the same or greater than the period. Required for [`simulate`] and [`ban`] modes. | `number` | no |
| action\_response\_content\_type | Custom content-type and body to return. It must be one of following 'text/plain', 'text/xml', 'application/json'. | `string` | no |
| action\_response\_body | The body to return. The content here must conform to the 'content_type' | `string` | no |
| correlate\_by | Whether to enable NAT based rate limiting. Default - `nat` | `string` | yes |
| disabled | Whether this rate limiting rule is currently disabled. | `string` | no |
| description | A note that you can use to describe the reason for a rate limiting rule. | `string` | no |
| bypass1\_name | bypass URL name. Default - `url` | `string` | no |
| bypass1\_value | bypass URL value | `string` | no |
| script_name | script name | `string` | yes |
| script | script content | `string` | yes |
| pattern | domain name pattern | `string` | yes |

## Outputs

| Name | Description |
|------|-------------|
| web\_dns\_name | Web DNS name. |
| instance\_id | CIS Instance Id |
| domain\_id | Domain Id. It is a combination of `domain_id`:`cis_id`|
| monitor |Monitor Id |
| rate_limit_id | Resource ID. It is a combination of `rule_id`:`domain_id`:`cis_id`|
| edge_functions_action_id | Resource ID. It is combination of `script_name`:`domain_id`:`cis_id`|
| edge_functions_trigger_id | Resource ID. It is combination of `route_id`:`domain_id`:`cis_id`|

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## References

1. [CIS Bluemix-go SDK](https://github.com/IBM-Cloud/bluemix-go/blob/master/api/cis/cisv1/)

