---
layout: "ibm"
page_title: "IBM: ibm_cis_healthcheck"
sidebar_current: "docs-ibm-resource-cis-healthcheck"
description: |-
  Provides a IBM Cloud Internet Services Health Check resource.
---

# ibm_cis_healthcheck

If you're using IBM's Cloud Internet Services Global Load Balancing to load-balance across multiple origin servers or data centers, you can configure a Healthcheck monitor to actively check the availability of those servers over HTTP(S). This resource is associated with an IBM Cloud Internet Services instance. 

## Example Usage

```hcl
resource "ibm_cis_healthcheck" "test" {
  cis_id = "${ibm_cis.instance.id}"
  expected_body = "alive"
  expected_codes = "2xx"
  method = "GET"
  timeout = 7
  path = "/health"
  interval = 60
  retries = 5
  description = "example load balancer"
}
```

## Argument Reference

The following arguments are supported:

* `cis_id` - (Required) The ID of the CIS service instance
* `expected_body` - (Required) A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy. A null value of "" is allowed to match on any content. 
* `expected_codes` - (Required) The expected HTTP response code or code range of the health check. Eg `2xx`
* `method` - (Optional) The HTTP method to use for the health check. Default: "GET".
* `timeout` - (Optional) The timeout (in seconds) before marking the health check as failed. Default: 5.
* `path` - (Optional) The endpoint path to health check against. Default: "/".
* `interval` - (Optional) The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations. Default: 60.
* `retries` - (Optional) The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately. Default: 2.
* `type` - (Optional) The protocol to use for the healthcheck. Currently supported protocols are 'HTTP' and 'HTTPS'. Default: "http".
* `description` - (Optional) Free text description.

Header is not currently supported in this version of the provider. 

## Attributes Reference

The following attributes are exported:

* `id` - Load balancer monitor ID.
* `created_on` - The RFC3339 timestamp of when the load balancer monitor was created.
* `modified_on` - The RFC3339 timestamp of when the load balancer monitor was last modified.

## Import

The `ibm_cis_health_check` resource can be imported using the `id`. The ID is formed from the `Healthcheck Id` and the `CRN` (Cloud Resource Name) concatentated usinga `:` character.  

The CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading. 

* **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

* **Healthcheck ID** is a 32 digit character string of the form: `1fc7c3247067ee00856729661c7d58c9`. The id of an existing Healthcheck monitor is not avaiable via the UI. It can be retrieved programmatically via the CIS API or via the CLI using the CIS command to list the defined GLBs:  `bx cis glb-monitors` 


```
$ terraform import ibm_cis_healthcheck.myorg <healthcheck_id>:<crn>

$ terraform import ibm_cis_healthcheck.myorg 1fc7c3247067ee00856729661c7d58c9:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
