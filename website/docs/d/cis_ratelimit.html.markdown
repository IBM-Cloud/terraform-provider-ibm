---
layout: "ibm"
page_title: "IBM: ibm_cis_rate_limit"
sidebar_current: "docs-ibm-cis-rate-limit"
description: |-
    Get information on an IBM Cloud Internet Services Rate Limit resource.
---

# ibm\_cis_rate_limit

Imports a read only copy of an existing Internet Services Ratelimiting resource. 

## Example Usage

```hcl
# Get a rate limit to the domain

data "ibm_cis_rate_limit" "ratelimit" {
    cis_id = data.ibm_cis.cis.id
    domain_id = data.ibm_cis_domain.cis_domain.id
}
```

## Argument Reference

The following arguments are supported:

* `cis_id` - (Required,string) The ID of the CIS service instance
* `domain_id` - (Required,string) The ID of the domain to add the Rate Limit rule.

**NOTE:**  To get a custom rate limit rule the CIS instance should be a `enterprise` plan

## Attributes Reference

The following attributes are exported:

* `id` - The record ID. It is a combination of <`rule_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
* `rule_id` - The Rate Limit Rule ID.
* `threshold` - The threshold that triggers the rate limit mitigations, combined with period. Possible range of threshold per period. Min value: 2, max value: 1000000.
* `period` - The time, in seconds, to count matching traffic. If the count exceeds threshold within this period the action is performed. Possible range - Min value:1, max value: 3600.
* `match` - Determines which traffic the rate limiting rule counts towards the threshold.
    * `request` - Matches HTTP requests.If not provided API ll default request to * , [_ALL_], `_ALL_` respectively.
        * `url` -   The URL pattern to match comprised of the host and path, for instance, example.org/path. Wildcards are expanded to match applicable traffic, query strings are not matched. Use * for all traffic to your zone. Max length is 1024.
        * `schemes` -  HTTP Schemes, can be one [HTTPS], both [HTTP,HTTPS] or all [_ALL_].
        * `methods` - HTTP Methods, can be a subset [POST,PUT] or all [_ALL_].Possible Values are `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `HEAD`, `_ALL_`.
    * `response` - Matches HTTP responses before they are returned to the client . If this is defined, then the entire counting of traffic occurs at this stage. 
        * `status` -  HTTP Status codes, can be one [403], many [401,403] or indicate all by not providing this value. Possible range Min value: 100, max value: 999.
        * `origin_traffic` -  Orrigin traffic.
        * `header` - Array of response headers to match. If a response does not meet the header criteria then the request is not counted towards the rate limiting rule. The header matching criteria includes following properties.
            * `name` -  The name of the response header to match.
            * `op` -  The operator when matching, eq means equals, ne means not equals. Possible Values are [`eq`] and [`ne`].
            * `value` -  WThe value of the header, which is exactly matched.
* `action` -  The action performed when the threshold of matched traffic within the period defined is exceeded.
    * `mode` -  The type of action performed. Possible Values are: [`simulate`], [`ban`], [`challenge`], [`js_challenge`].
    * `timeout` -  The time, in seconds, as an integer to perform the mitigation action. Timeout be the same or greater than the period. This field is valid only when mode is [`simulate`] or [`ban`].Possible range - Min value: 10, max value: 86400.
    * `response` - Custom content-type and body to return. This overrides the custom error for the zone. Omission results in the default HTML error page. This field is valid only when mode is [`simulate`] or [`ban`].
        * `content_type` -  The content-type of the body, which must be one of the following: [`text/plain`], [`text/xml`], [`application/json`].
        * `body` -  The body to return. The content here must confirm to the `content_type`. Possible Max length is 10240.
* `disabled` -  Whether this rate limiting rule is currently disabled.
* `description` -  A note that you can use to describe the reason for a rate limiting rule.
* `correlate` - Whether to enable NAT based rate limiting.
    * `by` -  Possible Values: [`nat`].
* `bypass` - Criteria that allows the rate limit to be bypassed. For example, to express that you shouldn’t apply a rate limit to a given set of URLs.
    * `name` -  Possible Value is [`url`].
    * `value` -  The url to bypass.
