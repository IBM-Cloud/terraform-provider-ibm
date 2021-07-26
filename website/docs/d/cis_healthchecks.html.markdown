---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_healthchecks"
description: |-
  Manages IBM Cloud Internet Services health check and monitor resource.
---

# ibm_cis_healthchecks

Retrieve information about an existing IBM Cloud Internet Service Global Load Balancer health monitor and check as a read-only data source. For more information, about CIS health check, see [setting up health checks](https://cloud.ibm.com/docs/cis?topic=cis-glb-features-healthchecks).

## Example usage
The following example retrieves information about an IBM Cloud Internet Services domain.

```terraform

data "ibm_cis_glb_health_checks" "test" {
  cis_id = var.cis_crn
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The resource CRN ID of the IBM Cloud Internet Services on which zones were created.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `allow_insecure` - (String) Do not validate the certificate when health check uses `HTTPS`.
- `created_on` - (String) The RFC3339 timestamp of when the load balancer monitor was created.
- `description` - (String) Free text description.
- `expected_body` - (String) The requested body.
- `expected_codes` - (String) The expected HTTP response code or code range of the health check. For example, `2xx`.
- `follow_redirects` - (String) Follow redirects if returned by the origin.
- `headers` - (String) The health check header.
- `id` - (String) The load balancer monitor ID and CRN. For example, `monitor_id:crn`.
- `interval` - (String) The interval between each health check. Shorter intervals improve failover time, but can increase load on the origins as you check from multiple locations. The default value is `60`.
- `modified_on` - (String) The RFC3339 timestamp of when the load balancer monitor was last modified.
- `monitor_id` - (String) The load balancer monitor ID.
- `method` - (String) The HTTP method to use for the health check.
- `path` - (String) The endpoint path to health check.
- `port` - (String) The TCP port to use for the health check.
- `retries` - (String) The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately. The default value is `2`.
- `timeout` - (String) The timeout (in seconds) before marking the health check as failed. The default value is `5`.
- `type` - (String) The protocol to use for the health check. Currently supported protocols are `HTTP`, `HTTPS`, and `TCP`. The default value is `HTTP`.
