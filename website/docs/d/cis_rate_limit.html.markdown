---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_rate_limit"
description: |-
  Get information on an IBM Cloud Internet Services rate limit resource.
---

# ibm_cis_rate_limit
Retrieve information for a rate limiting rule of an IBM Cloud Internet Services domain. To retrieve information about a rate limiting rule, you must have the enterprise plans for an IBM Cloud Internet Services. For more information, about rate limits, see [Rate limiting](https://cloud.ibm.com/docs/cis?topic=cis-cis-rate-limiting).

## Example usage

```terraform
# Get a rate limit to the domain

data "ibm_cis_rate_limit" "ratelimit" {
    cis_id = data.ibm_cis.cis.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance where you created the rate limiting rule.  
- `domain_id` - (Required, String) The ID of the domain where you created the rate limiting rule.

**Note** 

To get a custom rate limit rule the CIS instance must have an `enterprise` plan.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `action`- (List of actions) A list of actions that you want to perform when incoming requests exceed the specified `threshold`.

   Nested scheme for `action`:
   - `mode` - (String) The type of action that you want to perform. Supported values are `simulate`, `ban`, `challenge`, or `js_challenge`. For more information, about each type, see [Configure response](https://cloud.ibm.com/docs/cis?topic=cis-cis-rate-limiting#rate-limiting-configure-response).
   - `timeout`- (Integer) The time to wait in seconds before the action is performed. The timeout must be equal to or greater than the `period` and is valid only for actions of type `simulate` or `ban`. The value can be between 10 and 86400.
   - `response`- (List of response information) A list of information that you want to return to the client, such as the `content-type` and specific body information. The information provided in this parameter overrides the default HTML error page that is returned to the client. This option is valid only for actions of type `simulate` or `ban`.

     Nested scheme for `response`:
     - `content_type` - (String) The `content-type` of the body that you want to return. Supported values are `text/plain`, `text/xml`, and `application/json`.
     - `body` - (String) The body of the response that you want to return to the client. The information must match the `content_type` that you specified. The value can have a maximum length of 1024.
- `bypass`- (List of bypass criteria) A list of key-value pairs that, when matched, allow the rate limiting rule to be ignored.

  Nested scheme for `bypass`:
  - `name` - (String) The name of the key that you want to apply. Supported values are `url`.
  - `value` - (String) The value of the key that you want to match. When `name` is set to `url`, `value` contains the URL that you want to exclude from the rate limiting rule.
- `correlate`- (List of NAT based rate limits) If provided, NAT-based rate limiting is enabled.

  Nested scheme for `correlate`:
  - `by` - (String) If set to `nat`, NAT-based rate limiting is enabled.
- `description` - (String) The description for your rate limiting rule.
- `disabled`- (Bool) If set to **true**, rate limiting is disabled for the domain.
- `id` - (String) The record ID of the rate limiting rule in the format `<rule_ID>:<domain_ID>:<cis_ID>`.
- `match`- (List of matching rules) A list of characteristics that incoming network traffic must match to be counted toward the `threshold`.

  Nested scheme for `match`:
  - `request`- (List of request characteristics) A list of characteristics that the incoming request must match to be counted toward the `threshold`. If no list is provided, all incoming requests are counted toward the `threshold`.

    Nested scheme for `request`:
    - `methods`(Set of strings) The HTTP methods that the incoming request can use to be counted toward the `threshold`. Supported values are `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `HEAD`, and `ALL`. You can also combine multiple methods and separate them with a comma. For example `POST,PUT`.
    - `schemes`(Set of strings) The scheme of the request that determines the protocol that you want. Supported values are `HTTPS`, `HTTP,HTTPS`, and `ALL`.
    - `url` - (String) The URL that the request uses. Wildcard domains are expanded to match applicable traffic, query strings are not matched. If `*` is returned, the rule is applied to all URLs. The maximum length of this value can be 1024.
  - `response`- (List of HTTP responses) A list of HTTP responses that outgoing packets must match before they can be returned to the client. If an incoming request matches the request criteria, but the response does not match the response criteria, then the request packet is not counted toward the `threshold`. 

    Nested scheme for `response`:
    - `header`- (List of response headers) A list of HTTP response headers that the response packet must match so that the original request is counted toward the `threshold`.

      Nested scheme for `header`:
      - `name` - (String) The name of the HTTP response header.
      - `op` - (String) The operator that applied to your HTTP response header. Supported values are `eq` (equals) and `ne` (not equals).
      - `value` - (String) The value that the HTTP response header must match.
     - `origin_traffic` - (String) The origin traffic.
     - `status` (Set of integers) The HTTP status code that the response must have so that the request is counted toward the `threshold`. The value can be between 100 and 999. If you want to use multiple response codes, you must separate them with a comma, such as `401,403`.
- `period`- (Integer) The period of time in seconds where incoming requests to a domain are counted. If the number of requests exceeds the `threshold`, then connections to the domain are refused. The `period` value can be between 1 and 3600.
- `rule_id` - (String) The ID of the rate limiting rule.
- `threshold`- (Integer) The number of requests received within a specific time period (`period`) before connections to the domain are refused. The threshold value can be between 2 and 1000000.
