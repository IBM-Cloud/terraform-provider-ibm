---
layout: "ibm"
page_title: "IBM: ibm_cis_healthchecks"
sidebar_current: "docs-ibm-datasources-cis-healthchecks"
description: |-
Manages IBM Cloud Internet Services Health Check/Monitor resource.
---

# ibm_cis_healthchecks

Import the details of an existing IBM Cloud Internet Service global load balancer health monitor/check as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl

data "ibm_cis_glb_health_checks" "test" {
  cis_id = var.cis_crn
}

```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required, string) The resource crn id of the CIS on which zones were created.

## Attribute Reference

The following attributes are exported:

- `id` - Load balancer monitor ID and CRN. Ex. monitor_id:crn
- `monitor_id` - Load balancer monitor ID.
- `created_on` - The RFC3339 timestamp of when the load balancer monitor was created.
- `modified_on` - The RFC3339 timestamp of when the load balancer monitor was last modified.
- `expected_body` - Reqeusted nody
- `expected_codes` - The expected HTTP response code or code range of the health check. Eg `2xx`
- `method` - The HTTP method to use for the health check.
- `timeout` - The timeout (in seconds) before marking the health check as failed.The Default value is 5.
- `path` - The endpoint path to health check against.
- `interval` - The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations.The Default value is 60.
- `retries` - The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately.The Default value is 2.
- `type` - The protocol to use for the healthcheck. Currently supported protocols are 'HTTP', 'HTTPS' and 'TCP'. The Default value is 'HTTP'.
- `follow_redirects`- Follow redirects if returned by the origin.
- `allow_insecure`- Do not validate the certificate when healthcheck use HTTPS.
- `description` - Free text description.
- `port` - The TCP port to use for the health check.
- `headers` - The health check header
