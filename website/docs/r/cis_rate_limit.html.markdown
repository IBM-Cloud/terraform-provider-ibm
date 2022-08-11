---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_rate_limit"
description: |-
  Provides a IBM CIS rate limit resource.
---

# ibm_cis_rate_limit
Create, update, or delete custom rate limits for an IBM Cloud Internet Services domain. For more information, about rate limits, see [Rate limiting](https://cloud.ibm.com/docs/cis?topic=cis-cis-rate-limiting).

## Example usage
The following example shows how you can add a rate limit to an IBM Cloud Internet Services domain.

```terraform
# Add a rate limit to the domain

resource "ibm_cis_rate_limit" "ratelimit" {
    cis_id = data.ibm_cis.cis.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    threshold = 20
    period = 900
    match {
        request {
            url = "*.example.org/path*"
            schemes = ["HTTP", "HTTPS"]
            methods = ["GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"]
        }
        response {
            status = [200, 201, 202, 301, 429]
            origin_traffic = false
        }
    }
    action {
        mode = "ban"
        timeout = 43200
        response {
            content_type = "text/plain"
            body = "custom response body"
        }
    }
    correlate {
        by = "nat"
    }
    disabled = false
    description = "example rate limit for a zone"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `action`- (Required, List) A list of actions that you want to perform when incoming requests exceed the specified `threshold`.

  Nested scheme for `action`:
  - `mode` - (Required, String) The type of action that you want to perform. Supported values are `simulate`, `ban`, `challenge`, or `js_challenge`. For more information, about each type, see [Configure response](https://cloud.ibm.com/docs/cis?topic=cis-cis-rate-limiting#rate-limiting-configure-response).
  - `response`- (Optional, List) A list of information that you want to return to the client, such as the `content-type` and specific body information. The information provided in this parameter overrides the default HTML error page that is returned to the client. You can use this option only for actions of type `simulate` or `ban`.

    Nested scheme for `response`:
    - `content_type` - (Optional, String) The `content-type` of the body that you want to return. Supported values are `text/plain`, `text/xml`, and `application/json`.
    - `body` - (Optional, String) The body of the response that you want to return to the client. The information that you provide must match the `action.response.content_type` that you specified. The value that you enter can have a maximum length of 1024.
  - `timeout` - (Optional, Integer) The time to wait in seconds before the action is performed. The timeout must be equal or greater than the `period` and can be provided only for actions of type `simulate` or `ban`. The value that you enter must be between 10 and 86400.
- `bypass` - (Optional, List) A list of key-value pairs that, when matched, allow the rate limiting rule to be ignored. For example, use this option if you want to ignore the rate limiting for certain URLs.
	
  Nested scheme for `bypass`:
  - `name` - (Optional, String) The name of the key that you want to apply. Supported values are `url`.
  - `value` - (Optional, String) The value of the key that you want to match. When `name` is set to `url`, `value` must be set to the URL that you want to exclude from the rate limiting rule.
- `correlate` - (Optional, List) To enable NAT-based rate limiting.
   
   Nested scheme for `correlate`:
   - `by` - (Optional, String) Enter `nat` to enable NAT-based rate limiting.
- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `disabled` - (Optional, Bool) Set to **true** to disable rate limiting for a domain and **false** to enable rate limiting.
- `description` - (Optional, String) Enter a description for your rate limiting rule.
- `domain_id` - (Required, String) The ID of the domain where you want to add a rate limit.
- `match`- (Optional, List) A list of characteristics that incoming network traffic must match the `threshold` count. 

  Nested scheme for `match`:
  - `request`- (Optional, List) A list of characteristics that the incoming request match the `threshold` count. If this list is not provided, all incoming requests are matched the count of the `threshold`.

    Nested scheme for `request`:
    - `url` - (Optional, String) The URL that the request uses. Wildcard domains are expanded to match applicable traffic, query strings are not matched. You can use `*` to apply the rule to all URLs. The maximum length of this value can be 1024.
    - `schemes` - (Optional, Set of strings) The scheme of the request that determines the protocol that you want. Supported values are `HTTPS`, `HTTP,HTTPS`, and `_ALL_`.
    - `methods` - (Optional, Set of strings) The HTTP methods that the incoming request that match the `threshold` count. Supported values are `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `HEAD`, and `_ALL_`. You can also combine multiple methods and separate them with a comma. For example `POST,PUT`. 
- `response`- (Optional, List) A list of HTTP responses that outgoing packets must match before they can be returned to the client. If an incoming request matches the request criteria, but the response does not match the response criteria, then the request packet is not counted with the `threshold`. 

  Nested scheme for `response`:
  - `header`- (Optional, List) A list of HTTP response headers that the response packet must match so that the original request is matched with the `threshold` count.

    Nested scheme for `header`:
	  - `name` - (Optional, String) The name of the HTTP response header.
	  - `op` - (Optional, String) The operator that you want to apply to your HTTP response header. Supported values are `eq` (equals) and `ne` (not equals).
	  - `value` - (Optional, String) The value that the HTTP response header must match.
   - `origin_traffic` - (Optional, Bool). The origin traffic.
   - `status`- (Optional, Set(Integer))The HTTP status that the response must have so that the request is matched with the `threshold` count. You can specify one (`403`) or multiple (`401,403`) HTTP response codes. The value that you enter must be between 100 and 999.
- `period`- (Required, Integer) The period of time in seconds where incoming requests to a domain are counted. If the number of requests exceeds the `threshold`, then connections to the domain are refused. The `period` value must be between 1 and 3600.    
- `threshold`- (Required, Integer) The number of requests received within a specific time period (`period`) before connections to the domain are refused. The threshold value must be between 2 and 1000000.

**Note**

To create a custom rate limit rule the CIS instance should be a `enterprise` plan

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the rate limiting rule in the format `<rule_ID>:<domain_ID>:<cis_ID>`. .
- `rule_id` - (String) The rate limit rule ID.

## Import

The `ibm_cis_rate_limit` resource can be imported using the `id`. The ID is formed from the `Rate Limit rule ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `bx cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Rate Limit rule ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`.

**Syntax**

```
$ terraform import ibm_cis_rate_limit.ratelimit <rule_id>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_rate_limit.ratelimit 48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
