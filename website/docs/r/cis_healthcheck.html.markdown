---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_healthcheck"
description: |-
  Provides a IBM Cloud Internet Services Health Check resource.
---

# ibm_cis_healthcheck
Create, update, or delete an HTTPS health check for your IBM Cloud Internet Services instance. You can configure a health check monitor to actively check the availability of those servers over HTTP(S). For more information, about CIS health check, see [setting up health checks](https://cloud.ibm.com/docs/cis?topic=cis-glb-features-healthchecks).

## Example usage

```terraform
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

## Argument reference
Review the argument references that you can specify for your resource.

- `allow_insecure` - (Optional, Bool) If set to **true**, the certificate is not validated when the health check uses HTTPS. If set to **false**, the certificate is validated, even if the health check uses HTTPS. The default value is **false**.
- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `description` - (Optional, String) A description for your health check.
- `headers` - (Optional, String) The health check headers. Header is not currently supported in this version of the provider.

  Nested scheme for `headers`:
	- `header` - (Optional, String) The value of a header.
	- `values` - (Optional, String) The list of values for a header field. `[expected_body]`, and `[expected_codes]` are required arguments when the type is `HTTP` or `HTTPS`. **Note** Header is not currently supported in this version of the provider.
- `expected_body` - (Required, String) A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy. A null value of “” is allowed to match on any content.
- `expected_codes` - (Required, String) The expected HTTP response code or code range of the health check. For example, 200.
- `follow_redirects` - (Optional, Bool) If set to **true**, a redirect is followed when a redirect is returned by the origin pool. Is set to **false**, redirects from the origin pool are not followed.
- `interval` - (Optional, Integer) The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations. The default value is 60.
- `method` - (Optional, String) The HTTP method to use for the health check. Default: `GET`.
- `timeout` - (Optional, Integer) The timeout in seconds before marking the health check as failed. Default: 5.
- `path` - (Optional, String) The endpoint path to health check against. Default: `/`.
- `port` - (Optional, Integer) The TCP port number that you want to use for the health check.
- `retries` - (Optional, Integer) The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately. Default: 2.
- `type` - (Optional, String) The protocol to use for the health check. Currently supported protocols are `http` and `https`. Default: `http`.


## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the load balancer to monitor.
- `created_on` - (String) The RFC3339 timestamp of when the health check was created.
- `modified_on` - (String) The RFC3339 timestamp of when the health check was last modified.
- `monitor_id` - (String) The load balancer monitor ID.

## Import

The `ibm_cis_health_check` resource can be imported by using the `id`. The ID is formed from the `Healthcheck Id` and the `CRN` (Cloud Resource Name) concatentated usinga `:` character.

The CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading.

- **CRN** The CRN is unique ID of the form: `crn:v1:bluemix:public:internet-svcs:global:a/{IBM-account}:{service-instance}::`

- **Healthcheck ID** The health check ID is a 32 digit character string in the format 1aaaa111111aa11111111111a1a11a1. The ID of a health check is not available via the console. It can be retrieved programmatically via the CIS API or via the command line by running `ibmcloud cis glb-monitors`.

**Syntax**

```
$ terraform import ibm_cis_healthcheck.myorg <id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_healthcheck.myorg 1fc7c3247067ee00856729661c7d58c9:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
