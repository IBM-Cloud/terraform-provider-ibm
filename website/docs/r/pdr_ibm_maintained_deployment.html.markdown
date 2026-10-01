---
layout: "ibm"
page_title: "IBM : ibm_pdr_ibm_maintained_deployment"
description: |-
Manages pdr_ibm_maintained_deployment.
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
--------------------------------------------------------------------------------

# ibm_pdr_ibm_maintained_deployment

Creates ibm maintained Deployment by creating Orchestrator instance in the given PowerVS workspace and configuration. Orchestrator instance can be used to manage multiple virtual servers and ensure continuous availability.

## Example Usage

```hcl
resource "ibm_pdr_ibm_maintained_deployment" "pdr_ibm_maintained_deployment_instance" {
  accept_language = "en-US"
  instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
  standby_redeploy = "true"
  managed_apikey = "2345yujhgfds"
  orchestrator_password = "password"
  orchestrator_ha = true
}
```


### Path Parameters

* `instance_id` - (Required, Forces new resource, String) The unique identifier of the DR service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) Language code used to localize API response messages.
  * Constraints: The maximum length is `50` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous deprovisioning.
  * Constraints: The default value is `true`.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.
* `standby_redeploy` - (Optional, Forces new resource, String) Flag to indicate if standby should be redeployed (must be "true" or "false"). You should use this option only when HA deployment is failed during the initial attempt.
  * Constraints: The maximum length is `5` characters. The minimum length is `4` characters. The value must match regular expression `/^(true|false)$/`.

### Argument Reference

* `managed_apikey` - (Required, Forces new resource, String) API key used to manage the workloads by adding the PowerVS instances to the orchestrator.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.
* `orchestrator_password` - (Required, Forces new resource, String) The password that you can use to access your orchestrator.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[\x20-\x7E]*$/`.
* `orchestrator_ha` - (Optional, Forces new resource, Boolean) Indicates whether the orchestrator High Availability (HA) is enabled for the service instance.

## Attribute Reference

The following attributes are returned depending on the HTTP response status of the POST operation.

### HTTP 200 Response

The following attributes are returned when the operation completes successfully with an HTTP `200` response.

* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `id` - (String) The CRN (Cloud Resource Name) of the DR service instance.
  * Constraints: The maximum length is `512` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[a-zA-Z0-9\\-_]+:public:power-dr-automation:[a-zA-Z0-9\\-_]+:[a-zA-Z0-9\\-_\/]+:[a-zA-Z0-9\\-_]+::$/`.
* `dashboard_url` - (String) URL to the dashboard for managing the DR service instance in IBM Cloud.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^https?:\/\/[a-zA-Z0-9\\-._~:\/?#[\\]@!$&'()*+,;=]+$/`.

### HTTP 201 Response

The following attributes are returned when the operation is accepted for asynchronous processing with an HTTP `201` response.

* `dashboard_url` - (String) URL to the dashboard for monitoring the DR service instance operation in IBM Cloud.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^https?:\/\/[a-zA-Z0-9\\-._~:\/?#[\\]@!$&'()*+,;=]+$/`.
* `operation` - (String) The type of asynchronous operation being performed on the DR service instance (e.g., provision, deprovision, update).
  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.


## Import

You can import the `ibm_pdr_ibm_maintained_deployment` resource by using `id`.

The `id` property can be formed from `instance_id` and the service instance CRN in the following format:

<pre>
&lt;instance_id&gt;/&lt;instance_id&gt;
</pre>

* `instance_id`: A string in the format `123456d3-1122-3344-b67d-4389b44b7bf9`. Service Instance ID.
* `instance_id`: A string in the format `crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::`. The CRN (Cloud Resource Name) of the DR service instance.

# Syntax

<pre>
$ terraform import ibm_pdr_ibm_maintained_deployment.pdr_ibm_maintained_deployment &lt;instance_id&gt;/&lt;instance_id&gt;
</pre>
