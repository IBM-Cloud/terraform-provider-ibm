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
  cis_id         = ibm_cis.instance.id
  expected_body  = "alive"
  expected_codes = "2xx"
  method         = "GET"
  timeout        = 7
  path           = "/health"
  interval       = 60
  retries        = 3
  description    = "example load balancer"
  headers {
		header = "Host"
		values = ["example.com", "example1.com"]
	  }
	headers {
		header = "Host1"
		values = ["example3.com", "example11.com"]
	  }
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance
- `expected_body` - (Optional,string) A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy. A null value of "" is allowed to match on any content.
- `expected_codes` - (Optional,string) The expected HTTP response code or code range of the health check. Eg `2xx`
- `method` - (Optional,string) The HTTP method to use for the health check.
- `timeout` - (Optional,int) The timeout (in seconds) before marking the health check as failed.The Default value is 5.
- `path` - (Optional,string) The endpoint path to health check against.
- `interval` - (Optional,int) The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations.The Default value is 60.
- `retries` - (Optional,int) The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately.The Default value is 2.
- `type` - (Optional,string) The protocol to use for the healthcheck. Currently supported protocols are 'HTTP', 'HTTPS' and 'TCP'. The Default value is 'HTTP'.
- `follow_redirects`-(Optional,bool) Follow redirects if returned by the origin.
- `allow_insecure`-(Optional,bool) Do not validate the certificate when healthcheck use HTTPS.
- `description` - (Optional,string) Free text description.The Default value is false.
- `port` - (Optional,int) The TCP port to use for the health check.
- `headers` - (Optional,string) The health check headers.
  - `header` - (Optional,string) The value of header.
  - `values` - (Optional,string). The List of values for header field.

[`expected_body`],[`expected_codes`] are required aruguments when the type is HTTP or HTTPS.

Header is not currently supported in this version of the provider.

## Attributes Reference

The following attributes are exported:

- `id` - Load balancer monitor ID and CRN. Ex. monitor_id:crn.
- `monitor_id` - Load balancer monitor ID.
- `created_on` - The RFC3339 timestamp of when the load balancer monitor was created.
- `modified_on` - The RFC3339 timestamp of when the load balancer monitor was last modified.

## Import

The `ibm_cis_health_check` resource can be imported using the `id`. The ID is formed from the `Healthcheck Id` and the `CRN` (Cloud Resource Name) concatentated usinga `:` character.

The CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading.

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Healthcheck ID** is a 32 digit character string of the form: `1fc7c3247067ee00856729661c7d58c9`. The id of an existing Healthcheck monitor is not avaiable via the UI. It can be retrieved programmatically via the CIS API or via the CLI using the CIS command to list the defined GLBs: `ibmcloud cis glb-monitors`

```
$ terraform import ibm_cis_healthcheck.myorg <id>:<crn>

$ terraform import ibm_cis_healthcheck.myorg 1fc7c3247067ee00856729661c7d58c9:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
